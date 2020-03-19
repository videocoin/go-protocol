#!/bin/bash

COMMAND=${COMMAND:-solc}

proto_path=$1
output=$2

$COMMAND --bin -o $output --overwrite --allow-paths $proto_path openzeppelin-solidity=$proto_path/node_modules/openzeppelin-solidity $proto_path/contracts/StakingManager.sol $proto_path/contracts/Stream.sol $proto_path/contracts/StreamManager.sol
