package protocol

import (
	"math/big"

	"github.com/VideoCoin/go-protocol/abis/streamManager"
	"github.com/VideoCoin/go-videocoin/common"
	"github.com/VideoCoin/go-videocoin/ethclient"
)

// Manager is a wrapper for the stream manager calls.
type Manager struct {
	instance *streamManager.StreamManager
	acc      Account
}

// New creates a Manager instance
func New(url string, addr string, keyfilePath string, pwd string) (*Manager, error) {
	managerAddress := common.HexToAddress(addr)

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	managerInstance, err := streamManager.NewStreamManager(managerAddress, client)
	if err != nil {
		return nil, err
	}

	m := &Manager{
		instance: managerInstance,
		acc:      Account{client: client},
	}

	err = m.acc.loadAccount(keyfilePath, pwd)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// AddValidator adds a new address to the validator map in the StreamManager smart contract.
func (m *Manager) AddValidator() {

}

// RemoveValidator removes an address from the validator map in the StreamManager smart contract.
func (m *Manager) RemoveValidator() {

}

// ApproveStreamCreation approves a user`s stream request.
func (m *Manager) ApproveStreamCreation(streamID *big.Int, chunks []*big.Int) error {
	opt := m.acc.getTxOptions()

	tx, err := m.instance.ApproveStreamCreation(opt, streamID, chunks)
	if err != nil {
		return err
	}

	_ = tx

	return nil
}

// addInputChunk
func (m *Manager) addInputChunk() {

}
