peer chaincode install -p chaincodedev/chaincode/bad -n mycc -v 0
peer chaincode instantiate -n mycc -v 0 -c '{"function":"initProduct","Args":[""]}' -C myc
peer chaincode invoke -n mycc -v 0 -c '{"function":"readProduct","Args":["2015"]}' -C myc

