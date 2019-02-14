package protocol

import (
	"context"
	"math/big"

	"github.com/VideoCoin/go-protocol/abis/streamManager"
	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
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

	acc := Account{client: client}

	err = acc.loadAccount(keyfilePath, pwd)
	if err != nil {
		return nil, err
	}

	managerInstance, err := streamManager.NewStreamManager(managerAddress, client)
	if err != nil {
		return nil, err
	}

	m := &Manager{
		instance: managerInstance,
		acc:      acc,
	}

	return m, nil
}

// AddValidator adds a new address to the validator map in the StreamManager smart contract.
func (m *Manager) AddValidator(ctx context.Context, address string) error {
	opt := m.acc.getTxOptions()

	// TODO: check that address is not already validator

	addr := common.HexToAddress(address)

	tx, err := m.instance.AddValidator(opt, addr)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(nil, m.acc.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// RemoveValidator removes an address from the validator map in the StreamManager smart contract.
func (m *Manager) RemoveValidator(ctx context.Context, address string) error {
	opt := m.acc.getTxOptions()

	addr := common.HexToAddress(address)

	tx, err := m.instance.RemoveValidator(opt, addr)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(nil, m.acc.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// ApproveStreamCreation approves a user`s stream request.
func (m *Manager) ApproveStreamCreation(streamID *big.Int, chunks []*big.Int) error {
	opt := m.acc.getTxOptions()

	// TODO: add checks so we can return informative errors when needed

	tx, err := m.instance.ApproveStreamCreation(opt, streamID, chunks)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(nil, m.acc.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// addInputChunk
func (m *Manager) addInputChunk() {

}
