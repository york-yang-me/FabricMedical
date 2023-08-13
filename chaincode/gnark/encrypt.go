package gnark

import (
	"encoding/hex"
	"encoding/json"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type uploadInfo struct {
	Hash        string
	Description string
	VerifyKey   string
}

// CreateRealSequenceHash
// args[0] hash  args[1] description args[2] verifyKey
func CreateRealSequenceHash(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	info := uploadInfo{
		Hash:        args[0],
		Description: args[1],
		VerifyKey:   args[2],
	}

	value, err := stub.GetState(info.Hash)
	if err != nil {
		return shim.Error(err.Error())
	}

	if value != nil {
		return shim.Error("hash existed " + info.Hash)
	}

	// record hash value to state database
	m, _ := json.Marshal(info)
	err = stub.PutState(info.Hash, m)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

// UpdateRealSequence
// args[0] hash  args[1] description args[2] proof
func UpdateRealSequence(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	value, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	if value == nil {
		return shim.Error("hash not existed " + args[0])
	}

	info := uploadInfo{}
	err = json.Unmarshal(value, &info)
	if err != nil {
		return shim.Error(err.Error())
	}

	vk, _ := hex.DecodeString(info.VerifyKey)
	witness, _ := hex.DecodeString(args[2])

	// verify if the user has dna sequence
	result, err := verifyProof(info.Hash, vk, witness)
	if err != nil || !result {
		return shim.Error("verify fail")
	}

	info.Description = args[1]
	m, _ := json.Marshal(info)
	err = stub.PutState(info.Hash, m)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func QuerySequenceInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	value, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(value)
}
