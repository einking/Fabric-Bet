/*
peer chaincode query -n mycc -v 1.0 -C mychannel -c '{"function":"readbet","Args":["000000"]}'
*/
package main

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pass "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {
}

type bet struct {
	UserID string `json:"UserID"`
	Num    string `json:"Num"`
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (s *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pass.Response {
	_, args := stub.GetFunctionAndParameters()
	return s.initbet(stub, args)
}

func (s *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pass.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Printf("invoke is running" + function)

	if function == "initbet" {
		return s.initbet(stub, args)
	} else if function == "creatbet" {
		return s.creatbet(stub, args)
	} else if function == "readbet" {
		return s.readbet(stub, args)
	} else if function == "newbet" {
		return s.newbet(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func (s *SimpleChaincode) initbet(stub shim.ChaincodeStubInterface, args []string) pass.Response {
	rand.Seed(time.Now().Unix())
	seed := rand.Intn(100)
	seedStr := strconv.Itoa(seed)
	var betone bet
	betone.UserID = "000000"
	betone.Num = seedStr
	betAsBytes, _ := json.Marshal(betone)
	stub.PutState(betone.UserID, betAsBytes)

	return shim.Success(nil)
}
func (s *SimpleChaincode) newbet(stub shim.ChaincodeStubInterface, args []string) pass.Response {
	rand.Seed(time.Now().Unix())
	seed := rand.Intn(100)
	seedStr := strconv.Itoa(seed)
	var betone bet
	valAsbytes, _ := stub.GetState("000000")
	err := json.Unmarshal([]byte(valAsbytes), &betone)
	if err != nil {
		return shim.Error(err.Error())
	}
	betone.Num = seedStr
	betAsBytes, _ := json.Marshal(betone)
	stub.PutState(betone.UserID, betAsBytes)

	return shim.Success(nil)
}

func (s *SimpleChaincode) creatbet(stub shim.ChaincodeStubInterface, args []string) pass.Response {
	var err error
	//	1				2
	//  UserID	        Num
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	fmt.Println("starting add bet")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	var betone bet
	UserID := strings.ToLower(args[0])
	Num := strings.ToLower(args[1])
	betAsBytes, err := stub.GetState(UserID)
	if err != nil {
		return shim.Error("Failed to get bet ID" + err.Error())
	} else if betAsBytes != nil {
		err := json.Unmarshal([]byte(betAsBytes), &betone)
		if err != nil {
			return shim.Error(err.Error())
		}
		betone.Num = Num
	} else {
		betone.UserID = UserID
		betone.Num = Num
	}

	betJSONasBytes, err := json.Marshal(betone)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(UserID, betJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end add UserID")
	return shim.Success(nil)
}

func (s *SimpleChaincode) readbet(stub shim.ChaincodeStubInterface, args []string) pass.Response {
	var UserID, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting bet of the marble to query")
	}

	UserID = strings.ToLower(args[0])
	valAsbytes, err := stub.GetState(UserID) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + UserID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"bet does not exist: " + UserID + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}
