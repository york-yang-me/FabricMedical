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

func Test_CreateAuthorizing(t *testing.T) {
	stub := initTest(t)
	realSequenceList := checkCreateRealSequence(stub, t)
	// success
	checkInvoke(t, stub, [][]byte{
		[]byte("createAuthorizing"),
		[]byte(realSequenceList[0].RealSequenceID), // dna sequence in authorizing(RealSequenceID in authorizing)
		[]byte(realSequenceList[0].Owner),          // Hospital(Hospital AccountId)
		[]byte("50"),                               // price
		[]byte("30"),                               // the validity of the smart contract(in days)
	})
	// failed to verify objectOfAuthorizing belongs to hospital
	checkInvoke(t, stub, [][]byte{
		[]byte("createAuthorizing"),
		[]byte(realSequenceList[0].RealSequenceID),
		[]byte(realSequenceList[2].Owner),
		[]byte("50"),
		[]byte("30"),
	})
	checkInvoke(t, stub, [][]byte{
		[]byte("createAuthorizing"),
		[]byte("123"),
		[]byte(realSequenceList[0].Owner),
		[]byte("50"),
		[]byte("30"),
	})
	// parameter error
	checkInvoke(t, stub, [][]byte{
		[]byte("createAuthorizing"),
		[]byte(realSequenceList[0].RealSequenceID),
		[]byte(realSequenceList[0].Owner),
		[]byte("50"),
	})
	checkInvoke(t, stub, [][]byte{
		[]byte("createAuthorizing"),
		[]byte(""),
		[]byte(realSequenceList[0].Owner),
		[]byte("50"),
		[]byte("30"),
	})
}

// Test authorization initiation, authorizing and other operations
func Test_QueryAuthorizingList(t *testing.T) {
	stub := initTest(t)
	realSequenceList := checkCreateRealSequence(stub, t)
	fmt.Println(fmt.Sprintf("start to \n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("createAuthorizing"),
		[]byte(realSequenceList[0].RealSequenceID),
		[]byte(realSequenceList[0].Owner),
		[]byte("500000"),
		[]byte("30"),
	}).Payload)))
	fmt.Println(fmt.Sprintf("start to\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("createAuthorizing"),
		[]byte(realSequenceList[2].RealSequenceID),
		[]byte(realSequenceList[2].Owner),
		[]byte("600000"),
		[]byte("40"),
	}).Payload)))
	// query successfully
	fmt.Println(fmt.Sprintf("1、query all \n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAuthorizingList"),
	}).Payload)))
	fmt.Println(fmt.Sprintf("2、query sepcify %s\n%s", realSequenceList[0].Owner, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAuthorizingList"),
		[]byte(realSequenceList[0].Owner),
	}).Payload)))
	// authorizing
	fmt.Println(fmt.Sprintf("3、query %s account balance before authorizing\n%s", realSequenceList[2].Owner, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAccountList"),
		[]byte(realSequenceList[2].Owner),
	}).Payload)))
	fmt.Println(fmt.Sprintf("4、start to authorizing\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("createAppointing"),
		[]byte(realSequenceList[0].RealSequenceID),
		[]byte(realSequenceList[0].Owner),
		[]byte(realSequenceList[2].Owner),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》query %s account balance after authorizing\n%s", realSequenceList[2].Owner, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAccountList"),
		[]byte(realSequenceList[2].Owner),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》hospitall querys institute's authorizing information\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAuthorizingList"),
		[]byte(realSequenceList[0].Owner), // Institute(Institute AccountId)
	}).Payload)))
	fmt.Println(fmt.Sprintf("》institute querys institute's authorizing information\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAppointingList"),
		[]byte(realSequenceList[2].Owner), // Institute(Institute AccountId)
	}).Payload)))
	fmt.Println(fmt.Sprintf("》confirm hospital %s account balance before authorizing\n%s", realSequenceList[0].Owner, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAccountList"),
		[]byte(realSequenceList[0].Owner),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》confirm institute %s account balance before authorizing\n%s", realSequenceList[2].Owner, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAccountList"),
		[]byte(realSequenceList[2].Owner),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》confirm hospital %s with dna sequence before authorizing\n%s", realSequenceList[0].Owner, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryRealSequenceList"),
		[]byte(realSequenceList[0].Owner),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》confirm institute %s with dna sequence before authorizing\n%s", realSequenceList[2].Owner, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryRealSequenceList"),
		[]byte(realSequenceList[2].Owner),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》hospital confirm fund\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("updateAuthorizing"),
		[]byte(realSequenceList[0].RealSequenceID),
		[]byte(realSequenceList[0].Owner), // Hospital(Hospital AccountId)
		[]byte(realSequenceList[2].Owner), // Institute(Institute AccountId)
		[]byte("done"),                    // confirm fund
	}).Payload)))
	fmt.Println(fmt.Sprintf("》confirm hospital %s account balance after authorizing\n%s", realSequenceList[0].Owner, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAccountList"),
		[]byte(realSequenceList[0].Owner),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》confirm institute %s account balance after authorizing\n%s", realSequenceList[2].Owner, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAccountList"),
		[]byte(realSequenceList[2].Owner),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》confirm hospital %s with dna sequence after authorizing\n%s", realSequenceList[0].Owner, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryRealSequenceList"),
		[]byte(realSequenceList[0].Owner),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》confirm institute %s with dna sequence after authorizing\n%s", realSequenceList[2].Owner, string(checkInvoke(t, stub, [][]byte{
		[]byte("queryRealSequenceList"),
		[]byte(realSequenceList[2].Owner),
	}).Payload)))
	fmt.Println(fmt.Sprintf("》hospital query authorizing information(successful)\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAuthorizingList"),
		[]byte(realSequenceList[0].Owner), // Institute(Institute AccountId)
	}).Payload)))
	fmt.Println(fmt.Sprintf("》Institute query authorizing information(successful)\n%s", string(checkInvoke(t, stub, [][]byte{
		[]byte("queryAuthorizingList"),
		[]byte(realSequenceList[2].Owner), // Institute (Institute AccountId)
	}).Payload)))
}
