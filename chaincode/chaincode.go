package main

import (
	"chaincode/api"
	"chaincode/gnark"
	"chaincode/model"
	"chaincode/pkg/utils"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type BlockChainRealSequence struct {
}

// Init chaincode
func (t *BlockChainRealSequence) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("initial chaincode")
	//initial default chaincode
	var accountIds = [6]string{
		"5feceb66ffc8",
		"6b86b273ff34",
		"d4735e3a265e",
		"4e07408562be",
		"4b227777d4dd",
		"ef2d127de37b",
	}
	var userNames = [6]string{"admin", "hospital_1", "hospital_2", "patient_1", "patient_2", "patient_3"}
	var balances = [6]float64{0, 5000000, 5000000, 5000000, 5000000, 5000000}
	//initial accounts' data
	for i, val := range accountIds {
		account := &model.Account{
			AccountId: val,
			UserName:  userNames[i],
			Balance:   balances[i],
		}
		// write into ledger
		if err := utils.WriteLedger(account, stub, model.AccountKey, []string{val}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}
	return shim.Success(nil)
}

// Invoke realize Invoke interface callback smart contract
func (t *BlockChainRealSequence) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "hello":
		return api.Hello(stub, args)
	case "queryAccountList":
		return api.QueryAccountList(stub, args)
	case "createRealSequence":
		return api.CreateRealSequence(stub, args)
	case "queryRealSequenceList":
		return api.QueryRealSequenceList(stub, args)
	case "createRealSequenceHash":
		return gnark.CreateRealSequenceHash(stub, args)
	case "updateRealSequence":
		return gnark.UpdateRealSequence(stub, args)
	case "querySequenceInfo":
		return gnark.QuerySequenceInfo(stub, args)
	default:
		return shim.Error(fmt.Sprintf("no such function: %s", funcName))
	}
}

func main() {
	timeLocal, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = timeLocal
	err = shim.Start(new(BlockChainRealSequence))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
