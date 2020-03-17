#!/bin/bash

abidir=$1

mkdir -p staking
abigen --abi $abidir/StakingManager.abi --pkg staking --type StakingManager --out staking/manager.go
mkdir -p streams
abigen --abi $abidir/StreamManager.abi --pkg streams --type StreamManager --out streams/manager.go
abigen --abi $abidir/Stream.abi --pkg streams --type Stream --out streams/stream.go
