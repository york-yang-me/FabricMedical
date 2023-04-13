package utils

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

// WriteLedger write data into ledger
func WriteLedger(obj interface{}, stub shim.ChaincodeStubInterface, objectType string, keys []string) error {
	//create composite primary key
	var key string
	if val, err := stub.CreateCompositeKey(objectType, keys); err != nil {
		return errors.New(fmt.Sprintf("%s-Error creating composite primary key %s", objectType, err))
	} else {
		key = val
	}
	bytes, err := json.Marshal(obj)
	if err != nil {
		return errors.New(fmt.Sprintf("%s-Failed to serialize json data: %s", objectType, err))
	}
	//write data into blockchain ledger
	if err := stub.PutState(key, bytes); err != nil {
		return errors.New(fmt.Sprintf("%s-Failed to write into blockchain ledger: %s", objectType, err))
	}
	return nil
}

// DelLedger delete ledger
func DelLedger(stub shim.ChaincodeStubInterface, objectType string, keys []string) error {
	// create composite primary key
	var key string
	if val, err := stub.CreateCompositeKey(objectType, keys); err != nil {
		return errors.New(fmt.Sprintf("%s-Failed to serialize json data %s", objectType, err))
	} else {
		key = val
	}
	// write into blockchain ledger
	if err := stub.DelState(key); err != nil {
		return errors.New(fmt.Sprintf("%s-Failed to delete blockchain ledger: %s", objectType, err))
	}
	return nil
}

// GetStateByPartialCompositeKeys Query data based on compound primary key(suitable for all/multiple/single data)
// splitting keys into queries
func GetStateByPartialCompositeKeys(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results [][]byte, err error) {
	if len(keys) == 0 {
		// if key length == 0, return all data
		// find the relevant data from the blockchain by primary key,
		// which is equivalent to a fuzzy query on the primary key
		resultIterator, err := stub.GetStateByPartialCompositeKey(objectType, keys)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s-Failed to get all data: %s", objectType, err))
		}
		defer resultIterator.Close()

		// check if the return data is empty ==0 return empty array !=0 traverse data
		for resultIterator.HasNext() {
			val, err := resultIterator.Next()
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-Failed to results: %s", objectType, err))
			}

			results = append(results, val.GetValue())
		}
	} else {
		// if keys length != 0 return data for querying
		for _, v := range keys {
			// create composite key
			key, err := stub.CreateCompositeKey(objectType, []string{v})
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-Failed to create composite keys: %s", objectType, err))
			}
			// failed to get data from ledger
			bytes, err := stub.GetState(key)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("%s-Failed to get data: %s", objectType, err))
			}

			if bytes != nil {
				results = append(results, bytes)
			}
		}
	}

	return results, nil
}

// GetStateByPartialCompositeKeys2 Query data based on compound primary key (suitable for getting all or specified data)
func GetStateByPartialCompositeKeys2(stub shim.ChaincodeStubInterface, objectType string, keys []string) (results [][]byte, err error) {
	// Find relevant data from the blockchain by primary key == a fuzzy query on the primary key
	resultIterator, err := stub.GetStateByPartialCompositeKey(objectType, keys)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s-Failed to get all data: %s", objectType, err))
	}
	defer resultIterator.Close()

	// check if the data is Null
	for resultIterator.HasNext() {
		val, err := resultIterator.Next()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("%s-Failed to return the data: %s", objectType, err))
		}

		results = append(results, val.GetValue())
	}
	return results, nil
}
