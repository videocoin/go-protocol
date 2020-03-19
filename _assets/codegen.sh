#!/bin/bash

abidir=$1
bindir=$2
staking=$3
streams=$4

COMMAND=${COMMAND:-abigen}

$COMMAND --bin $bindir/StakingManager.bin --abi $abidir/StakingManager.abi --pkg staking --type StakingManager --out $staking/manager.go
$COMMAND --bin $bindir/StreamManager.bin --abi $abidir/StreamManager.abi --pkg streams --type StreamManager --out $streams/manager.go
$COMMAND --bin $bindir/Stream.bin --abi $abidir/Stream.abi --pkg streams --type Stream --out $streams/stream.go
