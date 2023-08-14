package main

import (
	"bytes"
	"chaincode/model"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"

	"github.com/hyperledger/fabric-chaincode-go/shimtest"
)

func initTest(t *testing.T) *shimtest.MockStub {
	scc := new(BlockChainRealSequence)
	stub := shimtest.NewMockStub("ex01", scc)
	checkInit(t, stub, [][]byte{[]byte("init")})
	return stub
}

func checkInit(t *testing.T, stub *shimtest.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shimtest.MockStub, args [][]byte) pb.Response {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
	return res
}

// test initial chaincode
func TestBlockChainRealSequence_Init(t *testing.T) {
	initTest(t)
}

// test getting account information
func Test_QueryAccountList(t *testing.T) {
	stub := initTest(t)
	fmt.Println(fmt.Sprintf("1、test to get all data\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryAccountList"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("2、test to obtain multiple data\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryAccountList"),
			[]byte("5feceb66ffc8"),
			[]byte("6b86b273ff34"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("3、test to get single data\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryAccountList"),
			[]byte("4e07408562be"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("4、test to invalid data\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryAccountList"),
			[]byte("0"),
		}).Payload)))
}

// Test create dna sequences
func Test_CreateRealSequences(t *testing.T) {
	stub := initTest(t)
	// success
	checkInvoke(t, stub, [][]byte{
		[]byte("createRealSequence"),
		[]byte("5feceb66ffc8"), // operator
		[]byte("6b86b273ff34"), // owner
		[]byte("1000"),         // total length
		[]byte("abschsw"),      // DNA contents
	})
	// insufficient authority
	checkInvoke(t, stub, [][]byte{
		[]byte("createRealSequence"),
		[]byte("6b86b273ff34"), // operator
		[]byte("4e07408562be"), // owner
		[]byte("50"),           // total length
		[]byte("abcdef"),       // DNA contents
	})
	// operator should be admin
	checkInvoke(t, stub, [][]byte{
		[]byte("createRealSequence"),
		[]byte("5feceb66ffc8"), //operator
		[]byte("5feceb66ffc8"), //owner
		[]byte("100"),          //total length
		[]byte("adqwqdsa"),     //DNA contents
	})
	// failed to verify Owner information
	checkInvoke(t, stub, [][]byte{
		[]byte("createRealSequence"),
		[]byte("5feceb66ffc8"),    //operator
		[]byte("6b86b273ff34555"), //owner
		[]byte("1000"),            //total length
		[]byte("sadsadasda"),      //DNA contents
	})
	// insufficient the number of parameters
	checkInvoke(t, stub, [][]byte{
		[]byte("createRealSequence"),
		[]byte("5feceb66ffc8"), //operator
		[]byte("6b86b273ff34"), //owner
		[]byte("500"),          //total length
	})
	// parameter format conversion error
	checkInvoke(t, stub, [][]byte{
		[]byte("createRealSequence"),
		[]byte("5feceb66ffc8"), //operator
		[]byte("6b86b273ff34"), //owner
		[]byte("50"),           //total length
		[]byte("daadsa"),       //DNA contents
	})
}

func checkCreateRealSequence(stub *shimtest.MockStub, t *testing.T) []model.RealSequence {
	var realSequenceList []model.RealSequence
	var realSequence model.RealSequence
	// success
	resp1 := checkInvoke(t, stub, [][]byte{
		[]byte("createRealSequence"),
		[]byte("5feceb66ffc8"), //operator
		[]byte("6b86b273ff34"), //owner
		[]byte("50"),           //total length
		[]byte("30"),           //DNA contents
	})
	resp2 := checkInvoke(t, stub, [][]byte{
		[]byte("createRealSequence"),
		[]byte("5feceb66ffc8"), //operator
		[]byte("6b86b273ff34"), //owner
		[]byte("80"),           //total length
		[]byte("60.8"),         //DNA contents
	})
	resp3 := checkInvoke(t, stub, [][]byte{
		[]byte("createRealSequence"),
		[]byte("5feceb66ffc8"), //operator
		[]byte("4e07408562be"), //owner
		[]byte("60"),           //total length
		[]byte("40"),           //DNA contents
	})
	resp4 := checkInvoke(t, stub, [][]byte{
		[]byte("createRealSequence"),
		[]byte("5feceb66ffc8"), //operator
		[]byte("ef2d127de37b"), //owner
		[]byte("80"),           //total length
		[]byte("60"),           //DNA contents
	})
	json.Unmarshal(bytes.NewBuffer(resp1.Payload).Bytes(), &realSequence)
	realSequenceList = append(realSequenceList, realSequence)
	json.Unmarshal(bytes.NewBuffer(resp2.Payload).Bytes(), &realSequence)
	realSequenceList = append(realSequenceList, realSequence)
	json.Unmarshal(bytes.NewBuffer(resp3.Payload).Bytes(), &realSequence)
	realSequenceList = append(realSequenceList, realSequence)
	json.Unmarshal(bytes.NewBuffer(resp4.Payload).Bytes(), &realSequence)
	realSequenceList = append(realSequenceList, realSequence)
	return realSequenceList
}

func Test_QueryRealSequenceList(t *testing.T) {
	stub := initTest(t)
	realSequenceList := checkCreateRealSequence(stub, t)

	fmt.Println(fmt.Sprintf("1、test to get all data\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryRealSequenceList"),
		}).Payload)))
	fmt.Println(fmt.Sprintf("2、test to get sepcified data\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryRealSequenceList"),
			[]byte(realSequenceList[0].Owner),
			[]byte(realSequenceList[0].RealSequenceID),
		}).Payload)))
	fmt.Println(fmt.Sprintf("3、test to get invalid data\n%s",
		string(checkInvoke(t, stub, [][]byte{
			[]byte("queryRealSequenceList"),
			[]byte("0"),
		}).Payload)))
}
