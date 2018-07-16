//Bad

/*
关闭端口：netstat -anp | grep 3000
kill -9 3655

# 调用命令
peer chaincode install -p chaincodedev/chaincode/bad -n mycc -v 0
peer chaincode instantiate -n mycc -v 0 -c '{"function":"initProduct","Args":[""]}' -C myc
peer chaincode invoke -n mycc -v 0 -c '{"function":"readCopyright","Args":[""]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["readCopyright","2015"]}' -C myc
peer chaincode invoke -n mycc -v 0 -c '{"function":"readCopyright","Args":["2015"]}' -C myc
peer chaincode invoke -n mycc -v 0 -c '{"function":"readCopyright","Args":["2015"]}' -C myc
# 添加Product
peer chaincode invoke -n copyrightforyou -v 0 -c '{"function":"creatCopyright","Args":["2017","888","888","888"]}' -C mychannel
peer chaincode invoke --tls /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"function":"creatCopyright","Args":["2017","2","522423","Beams"]}'

# 更改检查通过与否
peer chaincode invoke -n mycc -v 0 -c '{"function":"checkProduct","Args":["2015","100"]}' -C myc
peer chaincode invoke --tls /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"function":"sellCopyright","Args":["Hello"]}'

# 转让
# peer chaincode invoke -n mycc -v 0 -c '{"function":"transferProduct","Args":["2015","hello"]}' -C myc
peer chaincode invoke --tls /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"function":"transferCopyright","Args":["Hello"]}'

peer chaincode invoke --tls /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"function":"transferCopyright","Args":["2015","1","2","50"]}'


# 捐赠
peer chaincode invoke -n mycc -v 0 -c '{"function":"donateProduct","Args":["2015","hell"]}' -C myc
peer chaincode invoke --tls /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"function":"donateCopyright","Args":["Hello"]}'


# 毁掉Product
peer chaincode invoke -n mycc -v 0 -c '{"function":"destoryProduct","Args":["2019"]}' -C myc
peer chaincode invoke --tls /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"function":"destoryCopyright","Args":["2015"]}'


# 查看产品信息
peer chaincode invoke -n mycc -v 0 -c '{"function":"readCopyright","Args":["2015"]}' -C mychannel
peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n mycc -c '{"function":"readCopyright","Args":["2015"]}'
peer chaincode invoke --tls /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"function":"readCopyright","Args":["2015"]}'

# 查询产品历史信息
peer chaincode invoke -n mycc -v 0 -c '{"function":"getHistoryForProduct","Args":["2015"]}' -C myc
peer chaincode invoke --tls /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"function":"getHistoryForCopyright","Args":["2015"]}'






# 删除网络
docker rm -f $(docker ps -aq)
docker network prune

# 启动自己的链码
T_one:
cd /home/gopath/fabric-samples/chaincode-docker-devmode/
docker-compose -f docker-compose-simple.yaml up

T_two:
docker exec -it chaincode bash
cd sacc
go build
CORE_PEER_ADDRESS=peer:7050 CORE_CHAINCODE_ID_NAME=mycc:0 ./bad

T_three:
docker exec -it cli bash

*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pass "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {
}

type copyright struct {
	CopyrihtHash string `json:"copyrighthash"`
	UserId       string `json:"userid"`
	//Status       string            `json:"status"`
	Which map[string]string `json:"which"`
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (s *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pass.Response {
	_, args := stub.GetFunctionAndParameters()
	return s.initCopyright(stub, args)
}

func (s *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pass.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Printf("invoke is running" + function)

	if function == "initCopyright" {
		return s.initCopyright(stub, args)
	} else if function == "creatCopyright" {
		return s.creatCopyright(stub, args)
	} else if function == "transferCopyright" {
		return s.transferCopyright(stub, args)
	} else if function == "donateCopyright" {
		return s.donateCopyright(stub, args)
	} else if function == "destoryCopyright" {
		return s.destoryCopyright(stub, args)
	} else if function == "readCopyright" {
		return s.readCopyright(stub, args)
	} else if function == "getHistoryForCopyright" {
		return s.getHistoryForCopyright(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func (s *SimpleChaincode) initCopyright(stub shim.ChaincodeStubInterface, args []string) pass.Response {
	//s.Which = make(map[string]int)
	copyrights := []copyright{
		copyright{CopyrihtHash: "2015", UserId: "1" /*, Status: "Active" , Which: {"1": "100"}*/},
		copyright{CopyrihtHash: "2016", UserId: "2" /*, Status: "100" , Which: {"2": "100"}*/},
	}
	//copyrights.Which = make(map[string]string)
	i := 0
	for i < len(copyrights) {
		copyrights[i].Which = make(map[string]string)
		copyrights[i].Which[copyrights[i].UserId] = "100"
		fmt.Println("i is ", i)
		copyrightAsBytes, _ := json.Marshal(copyrights[i])
		stub.PutState(copyrights[i].CopyrihtHash, copyrightAsBytes)
		fmt.Println("Added", copyrights[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SimpleChaincode) creatCopyright(stub shim.ChaincodeStubInterface, args []string) pass.Response {
	var err error

	//	1				2					#5			#6
	//  CopyrihtHash	UserId				#status		#Which
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	fmt.Println("starting add copyright")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	/*if len(args[4]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	}*/

	copyrighthash := strings.ToLower(args[0])
	userid := strings.ToLower(args[1])
	copyrightAsBytes, err := stub.GetState(copyrighthash)
	if err != nil {
		return shim.Error("Failed to get copyright ID" + err.Error())
	} else if copyrightAsBytes != nil {
		fmt.Println("This copyright already exists: " + copyrighthash)
		return shim.Error("This copyright already exists: " + copyrighthash)
	}

	var copyright copyright
	copyright.CopyrihtHash = copyrighthash
	//copyright.Status = "100"
	copyright.UserId = userid
	copyright.Which = map[string]string{userid: "100"}
	//copyright.Which[userid] = "100"
	copyrightJSONasBytes, err := json.Marshal(copyright)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(copyrighthash, copyrightJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end add copyrighthash")
	return shim.Success(nil)
}

func (s *SimpleChaincode) transferCopyright(stub shim.ChaincodeStubInterface, args []string) pass.Response {
	//	0				1				2				3
	//	CopyrihtHash	oldowner		newowner		newstatus
	var err error
	if len(args) < 4 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	copyrighthash := strings.ToLower(args[0])
	oldowner := strings.ToLower(args[1])
	newowner := strings.ToLower(args[2])
	newstatus := args[3]
	newstatus_int, err := strconv.Atoi(newstatus)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("- start transferCopyright ", copyrighthash, oldowner, newowner, newstatus)

	copyrightAsBytes, err := stub.GetState(copyrighthash)
	if err != nil {
		return shim.Error("Failed to get copyright:" + err.Error())
	} else if copyrightAsBytes == nil {
		return shim.Error("copyright does not exist")
	}

	var copyrightOwner copyright
	//copyrightOwner:= copyright{}
	err = json.Unmarshal([]byte(copyrightAsBytes), &copyrightOwner)
	if err != nil {
		return shim.Error("trouble")
	}
	if newstatus_int > 100 {
		return shim.Error("The status can't bigger than 100")
	} else if newstatus_int == 100 {
		copyrightOwner.UserId = newowner
	} else if newstatus_int <= 0 {
		return shim.Error("The status can't smaller than 0")
	} else {
		/*for oldowner := range copyrightOwner.Which {
			//oldstatus_str := copyrightOwner.Which[oldowner]
			temp_oldstatus_int, err := strconv.Atoi(copyrightOwner.Which[oldowner])
			if err != nil {
				return shim.Error(err.Error())
			}
			if newstatus_int > temp_oldstatus_int {
				return shim.Error("Newstatus cna't bigger than oldstatus")
			} else {
				tempstatus_int := temp_oldstatus_int - newstatus_int
				tempstatus_str := strconv.Itoa(tempstatus_int)
				copyrightOwner.Which[oldowner] = tempstatus_str
				copyrightOwner.Which[newowner] = newstatus
			}
		}*/
		value, ok := copyrightOwner.Which[oldowner]
		if ok {
			temp_oldstatus_int, err := strconv.Atoi(value)
			if err != nil {
				//return shim.Error(err.Error())
				return shim.Error("trouble one")
			}
			if newstatus_int > temp_oldstatus_int {
				return shim.Error("Newstatus cna't bigger than oldstatus")
			} else {
				tempstatus_int := temp_oldstatus_int - newstatus_int
				value_new, ok_new := copyrightOwner.Which[newowner]
				if ok_new {
					temp_newstatus_int, err := strconv.Atoi(value_new)
					if err != nil {
						//return shim.Error(err.Error())
						return shim.Error("trouble one")
					}
					newstatus_int += temp_newstatus_int
				}
				if tempstatus_int == 0 {
					delete(copyrightOwner.Which, oldowner)

				} else {
					tempstatus_str := strconv.Itoa(tempstatus_int)
					copyrightOwner.Which[oldowner] = tempstatus_str
				}
				newstatus_str := strconv.Itoa(newstatus_int)
				copyrightOwner.Which[newowner] = newstatus_str
			}
		} else {
			return shim.Error("Key Not Found")
		}
	}

	copyrightJSONasBytes, _ := json.Marshal(copyrightOwner)
	err = stub.PutState(copyrighthash, copyrightJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferCopyright (success)")
	return shim.Success(nil)
}

// need change
func (s *SimpleChaincode) donateCopyright(stub shim.ChaincodeStubInterface, args []string) pass.Response {
	return s.transferCopyright(stub, args)
}

func (s *SimpleChaincode) destoryCopyright(stub shim.ChaincodeStubInterface, args []string) pass.Response {
	var jsonResp string
	var copyrightJSON copyright
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	copyrighthash := strings.ToLower(args[0])

	// to maintain the color~copyright index, we need to read the marble first and get its color
	valAsbytes, err := stub.GetState(copyrighthash) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + copyrighthash + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Marble does not exist: " + copyrighthash + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &copyrightJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + copyrighthash + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.DelState(copyrighthash) //remove the marble from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	return shim.Success(nil)
}

func (s *SimpleChaincode) readCopyright(stub shim.ChaincodeStubInterface, args []string) pass.Response {
	var copyrighthash, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting copyright of the marble to query")
	}

	copyrighthash = strings.ToLower(args[0])
	valAsbytes, err := stub.GetState(copyrighthash) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + copyrighthash + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Copyright does not exist: " + copyrighthash + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

func (s *SimpleChaincode) getHistoryForCopyright(stub shim.ChaincodeStubInterface, args []string) pass.Response {
	//	0
	//	copyrighthash

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	copyrighthash := strings.ToLower(args[0])
	fmt.Printf("- start getHistoryForCopyright: %s\n", copyrighthash)

	resultsIterator, err := stub.GetHistoryForKey(copyrighthash)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForCopyright returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
