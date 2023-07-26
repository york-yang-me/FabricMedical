package api

import (
	"chaincode/model"
	"chaincode/pkg/utils"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

// CreateAuthorizing Initiate authorization
func CreateAuthorizing(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// verify parameters
	if len(args) != 4 {
		return shim.Error("the number of parameters does not satisfy")
	}
	objectOfAuthorize := args[0]
	hospital := args[1]
	price := args[2]
	authorizePeriod := args[3]
	if objectOfAuthorize == "" || hospital == "" || price == "" || authorizePeriod == "" {
		return shim.Error("parameters exist empty")
	}
	// parameter data format conversion
	var formattedPrice float64
	if val, err := strconv.ParseFloat(price, 64); err != nil {
		return shim.Error(fmt.Sprintf("error in price parameter format conversion: %s", err))
	} else {
		formattedPrice = val
	}
	var formattedAuthorizePeriod int
	if val, err := strconv.Atoi(authorizePeriod); err != nil {
		return shim.Error(fmt.Sprintf("error in authorizePeriod parameter format conversion: %s", err))
	} else {
		formattedAuthorizePeriod = val
	}
	// determine whether objectOfAuthorize belongs to the hospital
	resultsRealSequence, err := utils.GetStateByPartialCompositeKeys2(stub, model.RealSequenceKey, []string{hospital, objectOfAuthorize})
	if err != nil || len(resultsRealSequence) != 1 {
		return shim.Error(fmt.Sprintf("verify %s belongs to %s failed: %s", objectOfAuthorize, hospital, err))
	}
	var realSequence model.RealSequence
	if err = json.Unmarshal(resultsRealSequence[0], &realSequence); err != nil {
		return shim.Error(fmt.Sprintf("CreateAuthorizing-deserialization error: %s", err))
	}
	// determine whether the record already exists, and cannot initiate sequence repeatedly
	// if Endorsement == true  the dna sequence has been endorsed by gmp
	if realSequence.Endorsement {
		return shim.Error("this dna sequence has been endorsed, not be authorizing repeatedly")
	}
	createTime, _ := stub.GetTxTimestamp()
	authorizing := &model.Authorizing{
		ObjectOfAuthorize: objectOfAuthorize,
		Hospital:          hospital,
		Patient:           "",
		Price:             formattedPrice,
		CreateTime:        time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2023-03-13 15:04:05"),
		AuthorizePeriod:   formattedAuthorizePeriod,
		AuthorizingStatus: model.AuthorizationStatusConstant()["publish"],
	}
	// write into ledger
	if err := utils.WriteLedger(authorizing, stub, model.AuthorizingKey, []string{authorizing.Hospital, authorizing.ObjectOfAuthorize}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	// set the dna sequence into "endorsing"
	realSequence.Endorsement = true
	if err := utils.WriteLedger(realSequence, stub, model.RealSequenceKey, []string{realSequence.Owner, realSequence.RealSequenceID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	// return created information successfully
	authorizingByte, err := json.Marshal(authorizing)
	if err != nil {
		return shim.Error(fmt.Sprintf("error in serialization of successfully created information: %s", err))
	}
	// return successfully
	return shim.Success(authorizingByte)
}

// CreateAppointing Participate in the assignment about dna sequence test
func CreateAppointing(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// verify parameters
	if len(args) != 3 {
		return shim.Error("the number of parameters does not satisfy")
	}
	objectOfAuthorizing := args[0]
	hospital := args[1]
	patient := args[2]
	if objectOfAuthorizing == "" || hospital == "" || patient == "" {
		return shim.Error("parameters exist empty")
	}
	if hospital == patient {
		return shim.Error("hospital can not be the same with patient")
	}
	// According to objectOfAuthorizing and hospital, get the dna sequences which want to test and confirm it
	resultsRealSequence, err := utils.GetStateByPartialCompositeKeys2(stub, model.RealSequenceKey, []string{hospital, objectOfAuthorizing})
	if err != nil || len(resultsRealSequence) != 1 {
		return shim.Error(fmt.Sprintf("according to %s and %s failed to get dna sequences: %s", objectOfAuthorizing, hospital, err))
	}
	// According to objectOfAuthorizing and hospital, get dna sequences information for authorizing
	resultsAuthorizing, err := utils.GetStateByPartialCompositeKeys2(stub, model.AuthorizingKey, []string{hospital, objectOfAuthorizing})
	if err != nil || len(resultsAuthorizing) != 1 {
		return shim.Error(fmt.Sprintf("according to %s and %s failed to get dna sequences: %s", objectOfAuthorizing, hospital, err))
	}
	var authorizing model.Authorizing
	if err = json.Unmarshal(resultsAuthorizing[0], &authorizing); err != nil {
		return shim.Error(fmt.Sprintf("CreateAppointing-deserialization error: %s", err))
	}
	// determine if the dna sequence is authorizing
	if authorizing.AuthorizingStatus != model.AuthorizationStatusConstant()["publish"] {
		return shim.Error("the transaction is not authorized")
	}
	// according to patient get patient information
	resultsAccount, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, []string{patient})
	if err != nil || len(resultsAccount) != 1 {
		return shim.Error(fmt.Sprintf("Failed to verify hospital information%s", err))
	}
	var patientAccount model.Account
	if err = json.Unmarshal(resultsAccount[0], &patientAccount); err != nil {
		return shim.Error(fmt.Sprintf("query patient information-deserialization error: %s", err))
	}
	if patientAccount.UserName == "admin" {
		return shim.Error(fmt.Sprintf("admin can't be appointed%s", err))
	}
	// determine if the balance is sufficient
	if patientAccount.Balance < authorizing.Price {
		return shim.Error(fmt.Sprintf("dna seuqence authorized price is %f, your current balance %f, failed to be appointed", authorizing.Price, patientAccount.Balance))
	}
	// put patient writing into authorizing, modify the transaction status
	authorizing.Patient = patient
	authorizing.AuthorizingStatus = model.AuthorizationStatusConstant()["delivery"]
	if err := utils.WriteLedger(authorizing, stub, model.AuthorizingKey, []string{authorizing.Hospital, authorizing.ObjectOfAuthorize}); err != nil {
		return shim.Error(fmt.Sprintf("faile to put patient writing into authorizing %s", err))
	}
	createTime, _ := stub.GetTxTimestamp()
	// put the transaction into ledger, which is helpful to query for institue
	appointing := &model.Appointing{
		Patient:     patient,
		CreateTime:  time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2023-03-13 15:04:05"),
		Authorizing: authorizing,
	}
	if err := utils.WriteLedger(appointing, stub, model.AppointingKey, []string{appointing.Patient, appointing.CreateTime}); err != nil {
		return shim.Error(fmt.Sprintf("faile to put the transaction into ledger %s", err))
	}
	appointingByte, err := json.Marshal(appointing)
	if err != nil {
		return shim.Error(fmt.Sprintf("error in serialization of successfully created information: %s", err))
	}
	// Authorization successfully, deduct the balance,
	// update the balance of the account ledger,
	// note: the hospital needs to confirm receipt,
	// the money will be transferred to the hospital's account, here first deduct the patient's balance
	patientAccount.Balance -= authorizing.Price
	if err := utils.WriteLedger(patientAccount, stub, model.AccountKey, []string{patientAccount.AccountId}); err != nil {
		return shim.Error(fmt.Sprintf("Failed to debit patient's balance%s", err))
	}
	// return successfully
	return shim.Success(appointingByte)
}

// QueryAuthorizingList query authorization(query all information, and also can check every sequence according to hospital)
func QueryAuthorizingList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var authorizingList []model.Authorizing
	results, err := utils.GetStateByPartialCompositeKeys2(stub, model.AuthorizingKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var authorizing model.Authorizing
			err := json.Unmarshal(v, &authorizing)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryAuthorizingList-deserialization error: %s", err))
			}
			authorizingList = append(authorizingList, authorizing)
		}
	}
	authorizingListByte, err := json.Marshal(authorizingList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryauthorizingList-serialization error:: %s", err))
	}
	return shim.Success(authorizingListByte)
}

// QueryAppointingList query the dna sequences which have been appointed (query all information, and also can check every sequence according to hospital)
func QueryAppointingList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("must be queried by hospital AccountId"))
	}
	var appointingList []model.Appointing
	results, err := utils.GetStateByPartialCompositeKeys2(stub, model.AppointingKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var appointing model.Appointing
			err := json.Unmarshal(v, &appointing)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryAppointingList-deserialization error: %s", err))
			}
			appointingList = append(appointingList, appointing)
		}
	}
	appointingListByte, err := json.Marshal(appointingList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAppointingList-serialization error: %s", err))
	}
	return shim.Success(appointingListByte)
}

// UpdateAuthorizing Update authorization (hospital confirm、patient cancel)
func UpdateAuthorizing(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// verify parameter
	if len(args) != 4 {
		return shim.Error("the number of parameters does not satisfy")
	}
	objectOfAuthorizing := args[0]
	hospital := args[1]
	patient := args[2]
	status := args[3]
	if objectOfAuthorizing == "" || hospital == "" || status == "" {
		return shim.Error("parameters exist empty")
	}
	if hospital == patient {
		return shim.Error("hospital can not be the same with patient")
	}
	// According to objectOfAuthorizing and hospital, get the dna sequence information
	resultsRealSequence, err := utils.GetStateByPartialCompositeKeys2(stub, model.RealSequenceKey, []string{hospital, objectOfAuthorizing})
	if err != nil || len(resultsRealSequence) != 1 {
		return shim.Error(fmt.Sprintf("according to %s and %s, failed to get the dna sequence information: %s", objectOfAuthorizing, hospital, err))
	}
	var realSequence model.RealSequence
	if err = json.Unmarshal(resultsRealSequence[0], &realSequence); err != nil {
		return shim.Error(fmt.Sprintf("UpdateAuthorizing-deserialization error: %s", err))
	}
	// according to objectOfAuthorizing and hospital, get the dna sequences
	resultsAuthorizing, err := utils.GetStateByPartialCompositeKeys2(stub, model.AppointingKey, []string{hospital, objectOfAuthorizing})
	if err != nil || len(resultsAuthorizing) != 1 {
		return shim.Error(fmt.Sprintf("according to %s and %s failed to authorization information: %s", objectOfAuthorizing, hospital, err))
	}
	var authorizing model.Authorizing
	if err = json.Unmarshal(resultsAuthorizing[0], &authorizing); err != nil {
		return shim.Error(fmt.Sprintf("UpdateAuthorizing-deserialization error: %s", err))
	}
	// according to hospital, get patient information about appointing
	var appointing model.Appointing
	// if the current status is "publish", it means no patient appointed
	if authorizing.AuthorizingStatus != model.AuthorizationStatusConstant()["publish"] {
		resultsAppointing, err := utils.GetStateByPartialCompositeKeys2(stub, model.AppointingKey, []string{patient})
		if err != nil || len(resultsAppointing) == 0 {
			return shim.Error(fmt.Sprintf("according to %s, failed to get patient information about authorization: %s", patient, err))
		}
		for _, v := range resultsAppointing {
			if v != nil {
				var s model.Appointing
				err := json.Unmarshal(v, &s)
				if err != nil {
					return shim.Error(fmt.Sprintf("UpdateAppointing-deserialization error: %s", err))
				}
				if s.Authorizing.ObjectOfAuthorize == objectOfAuthorizing && s.Authorizing.Hospital == hospital && s.Patient == patient {
					// determine that the status must be in delivery,
					// in case the dna sequence has already been transacted and is just cancelled.
					if s.Authorizing.AuthorizingStatus == model.AuthorizationStatusConstant()["delivery"] {
						appointing = s
						break
					}
				}
			}
		}
	}
	var data []byte
	// determine the status of authorizing
	switch status {
	case "done":
		// ensure that the sale is in delivery status
		if authorizing.AuthorizingStatus != model.AuthorizationStatusConstant()["delivery"] {
			return shim.Error("the transaction is not in delivery, failed to confirm fund")
		}
		// according to hospital, get hospital information
		resultsHospitalAccount, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, []string{hospital})
		if err != nil || len(resultsHospitalAccount) != 1 {
			return shim.Error(fmt.Sprintf("Failed to verify hospital information %s", err))
		}
		var accountHospital model.Account
		if err = json.Unmarshal(resultsHospitalAccount[0], &accountHospital); err != nil {
			return shim.Error(fmt.Sprintf("query hospital information-deserialization error: %s", err))
		}
		// confirm the payment, add the fund to the hospital's account
		accountHospital.Balance += authorizing.Price
		if err := utils.WriteLedger(accountHospital, stub, model.AccountKey, []string{accountHospital.AccountId}); err != nil {
			return shim.Error(fmt.Sprintf("failed to accpet fund for hospital%s", err))
		}
		// transfer the owner information to the hospital and reset the endorsement status
		realSequence.Owner = hospital
		realSequence.Endorsement = false
		// realSequence.RealSequenceID = stub.GetTxID() // restart updating dna sequence id
		if err := utils.WriteLedger(realSequence, stub, model.RealSequenceKey, []string{realSequence.Owner, realSequence.RealSequenceID}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		// cleanup the original dna sequence information
		if err := utils.DelLedger(stub, model.RealSequenceKey, []string{hospital, objectOfAuthorizing}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		// authorizing status finished, writing into ledger
		authorizing.AuthorizingStatus = model.AuthorizationStatusConstant()["done"]
		authorizing.ObjectOfAuthorize = realSequence.RealSequenceID // renew dna sequence id
		if err := utils.WriteLedger(authorizing, stub, model.AuthorizingKey, []string{authorizing.Hospital, objectOfAuthorizing}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		appointing.Authorizing = authorizing
		if err := utils.WriteLedger(appointing, stub, model.AppointingKey, []string{appointing.Patient, appointing.CreateTime}); err != nil {
			return shim.Error(fmt.Sprintf("failed to write into ledger: %s", err))
		}
		data, err = json.Marshal(appointing)
		if err != nil {
			return shim.Error(fmt.Sprintf("error in serialize transaction: %s", err))
		}
		break
	case "cancelled":
		data, err = closeAuthorizing("cancelled", authorizing, realSequence, appointing, patient, stub)
		if err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		break
	case "expired":
		data, err = closeAuthorizing("expired", authorizing, realSequence, appointing, patient, stub)
		if err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
		break
	default:
		return shim.Error(fmt.Sprintf("%s status not supported", status))
	}
	return shim.Success(data)
}

// closeAuthorizing
// 1、stay in the status of "publish"
// 2、stay in the status of "delivery"
func closeAuthorizing(closeStart string, authorizing model.Authorizing, realSequence model.RealSequence, appointing model.Appointing, patient string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	switch authorizing.AuthorizingStatus {
	case model.AuthorizationStatusConstant()["saleStart"]:
		authorizing.AuthorizingStatus = model.AuthorizationStatusConstant()[closeStart]
		// reset endorsement information
		realSequence.Endorsement = false
		if err := utils.WriteLedger(realSequence, stub, model.RealSequenceKey, []string{realSequence.Owner, realSequence.RealSequenceID}); err != nil {
			return nil, err
		}
		if err := utils.WriteLedger(authorizing, stub, model.AuthorizingKey, []string{authorizing.Hospital, authorizing.ObjectOfAuthorize}); err != nil {
			return nil, err
		}
		data, err := json.Marshal(authorizing)
		if err != nil {
			return nil, err
		}
		return data, nil
	case model.AuthorizationStatusConstant()["delivery"]:
		// get patient information by patient
		resultsPatientAccount, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, []string{patient})
		if err != nil || len(resultsPatientAccount) != 1 {
			return nil, err
		}
		var accountPatient model.Account
		if err = json.Unmarshal(resultsPatientAccount[0], &accountPatient); err != nil {
			return nil, err
		}
		// Cancellation require the fund return to the Patient
		accountPatient.Balance += authorizing.Price
		if err := utils.WriteLedger(accountPatient, stub, model.AccountKey, []string{accountPatient.AccountId}); err != nil {
			return nil, err
		}
		// reset dna sequence endorsement
		realSequence.Endorsement = false
		if err := utils.WriteLedger(realSequence, stub, model.RealSequenceKey, []string{realSequence.Owner, realSequence.RealSequenceID}); err != nil {
			return nil, err
		}
		// Update authorization status
		authorizing.AuthorizingStatus = model.AuthorizationStatusConstant()[closeStart]
		if err := utils.WriteLedger(authorizing, stub, model.AuthorizingKey, []string{authorizing.Hospital, authorizing.ObjectOfAuthorize}); err != nil {
			return nil, err
		}
		appointing.Authorizing = authorizing
		if err := utils.WriteLedger(appointing, stub, model.AppointingKey, []string{appointing.Patient, appointing.CreateTime}); err != nil {
			return nil, err
		}
		data, err := json.Marshal(appointing)
		if err != nil {
			return nil, err
		}
		return data, nil
	default:
		return nil, nil
	}
}
