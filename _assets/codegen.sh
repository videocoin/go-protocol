#!/bin/bash

abidir=$1
staking=$2
streams=$3

COMMAND=${COMMAND:-abigen}

$COMMAND --abi $abidir/StakingManager.abi --pkg staking --type StakingManager --out $staking/manager.go
$COMMAND --abi $abidir/StreamManager.abi --pkg streams --type StreamManager --out $streams/manager.go
$COMMAND --abi $abidir/Stream.abi --pkg streams --type Stream --out $streams/stream.go
