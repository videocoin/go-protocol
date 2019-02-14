package protocol

import (
	"github.com/VideoCoin/go-protocol/abis/streamManager"
	"github.com/VideoCoin/go-videocoin/common"
	"github.com/VideoCoin/go-videocoin/ethclient"
)

// Manager is a wrapper for the stream manager calls.
type Manager struct {
	keyjson  []byte
	instance *streamManager.StreamManager
}

// New creates ...
func New(addr string, url string) (*Manager, error) {
	managerAddress := common.HexToAddress(addr)

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	managerInstance, err := streamManager.NewStreamManager(managerAddress, client)
	if err != nil {
		return nil, err
	}

	return &Manager{instance: managerInstance}, nil
}

// AddValidator adds a new address to the validator map in the StreamManager smart contract.
func AddValidator() {

}

// RemoveValidator removes an address from the validator map in the StreamManager smart contract.
func RemoveValidator() {

}

// ApproveStreamCreation approves a user`s stream request.
func ApproveStreamCreation() {

}

// addInputChunk
func addInputChunk() {

}
