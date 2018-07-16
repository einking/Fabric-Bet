#!/bin/bash
#
# Copyright LYH Corp. All Rights Reserved.
#
#

cd bad
CORE_PEER_ADcomposeDRESS=peer:7050 CORE_CHAINCODE_ID_NAME=mycc:0 ./bad 2>&1


