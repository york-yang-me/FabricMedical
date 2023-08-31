package api

import (
	"chaincode/model"
	"chaincode/pkg/utils"
	"chaincode/verification"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"strconv"
)

// CreateRealSequence create realSequence(admin)
func CreateRealSequence(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// verify parameters
	if len(args) != 4 {
		return shim.Error("the number of parameters are not satisfy")
	}
	accountId := args[0] // accountId is used for verify admin
	owner := args[1]
	totalLength := args[2]
	dnaContents := args[3]
	description := args[4]
	verifyKey := args[5]
	proof := args[6]
	if accountId == "" || owner == "" || totalLength == "" || dnaContents == "" || description == "" {
		return shim.Error("parameters exist null")
	}
	if accountId == owner {
		return shim.Error("the operator should be admin and cannot be the same with everyone")
	}
	// parameter data format conversion
	var formattedTotalLength int
	if val, err := strconv.Atoi(totalLength); err != nil {
		return shim.Error(fmt.Sprintf("error in totalLength parameter format conversion:%s", err))
	} else {
		formattedTotalLength = val
	}
	var formattedDNAContents string
	formattedDNAContents = dnaContents
	// determine if it is an administrator operation
	resultsAccount, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, []string{accountId})
	if err != nil || len(resultsAccount) != 1 {
		return shim.Error(fmt.Sprintf("failed to verify operator permissions%s", err))
	}
	var account model.Account
	if err = json.Unmarshal(resultsAccount[0], &account); err != nil {
		return shim.Error(fmt.Sprintf("check operator information-deserialization error: %s", err))
	}
	if account.UserName != "admin" {
		return shim.Error(fmt.Sprintf("insufficient operator privileges: %s", err))
	}
	// determine if the owner exists
	resultsOwner, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, []string{owner})
	if err != nil || len(resultsOwner) != 1 {
		return shim.Error(fmt.Sprintf("failed to verify owner information %s", err))
	}
	realSequence := &model.RealSequence{
		RealSequenceID:  stub.GetTxID()[:16],
		Owner:           owner,
		Endorsement:     false,
		TotalLength:     formattedTotalLength,
		DNAContentsHash: formattedDNAContents,
		Description:     description,
		VerifyKey:       verifyKey,
		Proof:           proof,
	}
	// write into ledger
	if err := utils.WriteLedger(realSequence, stub, model.RealSequenceKey, []string{realSequence.Owner, realSequence.RealSequenceID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	// return created information successfully
	realSequenceByte, err := json.Marshal(realSequence)
	if err != nil {
		return shim.Error(fmt.Sprintf("error in serialization of successfully created information: %s", err))
	}
	// return successfully
	return shim.Success(realSequenceByte)
}

// QueryRealSequenceList query real dna sequence(query all information, and also can check every sequence according to patient)
func QueryRealSequenceList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var realSequenceList []model.RealSequence
	results, err := utils.GetStateByPartialCompositeKeys2(stub, model.RealSequenceKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var realSequence model.RealSequence
			err := json.Unmarshal(v, &realSequence)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryRealSequenceList-deserialization error: %s", err))
			}
			realSequenceList = append(realSequenceList, realSequence)
		}
	}
	realSequenceListByte, err := json.Marshal(realSequenceList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryRealSequenceList-serialization error: %s", err))
	}
	return shim.Success(realSequenceListByte)
}

// UpdateRealSequence
// args[0] hash  args[1] description args[2] proof
func UpdateRealSequence(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	results, err := utils.GetStateByPartialCompositeKeys2(stub, model.RealSequenceKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range results {
		if v != nil {
			var realSequence model.RealSequence
			err := json.Unmarshal(v, &realSequence)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryRealSequenceList-deserialization error: %s", err))
			}
			if realSequence.DNAContentsHash == args[0] {
				vk, _ := hex.DecodeString(realSequence.VerifyKey)
				witness, _ := hex.DecodeString(args[2])
				// verify if the user has dna sequence
				result, err := verification.VerifyProof(realSequence.DNAContentsHash, vk, witness)
				if err != nil || !result {
					return shim.Error("verify fail")
				}
				realSequence.Description = args[1]
				// write into ledger
				if err := utils.WriteLedger(realSequence, stub, model.RealSequenceKey, []string{realSequence.Owner, realSequence.RealSequenceID}); err != nil {
					return shim.Error(fmt.Sprintf("%s", err))
				}
				return shim.Success(nil)
			}
		}
	}

	return shim.Error(fmt.Sprintf("update failed: %s", err))
}
