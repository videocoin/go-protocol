package protocol

import (
	"context"
	"math/big"

	"github.com/VideoCoin/go-protocol/abis/streamManager"
	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
	"github.com/VideoCoin/go-videocoin/common"
	"github.com/VideoCoin/go-videocoin/ethclient"
)

// ManagerClient is a wrapper for the stream manager calls.
type ManagerClient struct {
	instance *streamManager.StreamManager
	acc      Account
}

// NewManagerClient creates a ManagerClient instance
func NewManagerClient(url string, addr string, keyfilePath string, pwd string) (*ManagerClient, error) {
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

	m := &ManagerClient{
		instance: managerInstance,
		acc:      acc,
	}

	return m, nil
}

// AddValidator adds a new address to the validator map in the StreamManager smart contract.
func (m *ManagerClient) AddValidator(ctx context.Context, address string) error {
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
func (m *ManagerClient) RemoveValidator(ctx context.Context, address string) error {
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
func (m *ManagerClient) ApproveStreamCreation(streamID *big.Int, chunks []*big.Int) error {
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

// AddInputChunk will addd input chunk ids to the stream contract
func (m *ManagerClient) AddInputChunk() {

}

// AllowRefund will allow the client to refund the escrow for the given stream id.
func (m *ManagerClient) AllowRefund() {

}
