package protocol

import (
	"context"
	"fmt"
	"math/big"

	"github.com/VideoCoin/go-protocol/abis/streamManager"
	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
	"github.com/VideoCoin/go-videocoin/common"
	"github.com/VideoCoin/go-videocoin/ethclient"
)

// ManagerClient is a wrapper for the stream manager calls.
type ManagerClient struct {
	instance *streamManager.StreamManager
	caller   Caller
	addr     common.Address
}

// GetManagerAccAddr returns manager account address
func (m *ManagerClient) GetManagerAccAddr() common.Address {
	return m.caller.key.Address
}

// GetManagerContractAddr returns manager smart contract address
func (m *ManagerClient) GetManagerContractAddr() common.Address {
	return m.addr
}

// NewManagerClient creates a ManagerClient instance
func NewManagerClient(url string, addr string, keyfilePath string, pwd string) (*ManagerClient, error) {
	managerAddress := common.HexToAddress(addr)

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	caller := Caller{client: client}

	err = caller.loadAccount(keyfilePath, pwd)
	if err != nil {
		return nil, err
	}

	managerInstance, err := streamManager.NewStreamManager(managerAddress, client)
	if err != nil {
		return nil, err
	}

	m := &ManagerClient{
		instance: managerInstance,
		caller:   caller,
		addr:     managerAddress,
	}

	isOwner, err := managerInstance.IsOwner(&bind.CallOpts{From: m.caller.key.Address})
	if err != nil {
		return nil, err
	}

	if !isOwner {
		owner, err := managerInstance.Owner(&bind.CallOpts{})
		err = fmt.Errorf("Account provided %s, is not owner. Real owner is: %s ", m.caller.key.Address.String(), owner.String())
		return nil, err
	}

	return m, nil
}

// AddValidator adds a new address to the validator map in the StreamManager smart contract.
func (m *ManagerClient) AddValidator(ctx context.Context, address string) error {
	opt := m.caller.getTxOptions(0)

	// TODO: check that address is not already validator

	addr := common.HexToAddress(address)

	tx, err := m.instance.AddValidator(opt, addr)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, m.caller.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// RemoveValidator removes an address from the validator map in the StreamManager smart contract.
func (m *ManagerClient) RemoveValidator(ctx context.Context, address string) error {
	opt := m.caller.getTxOptions(0)

	addr := common.HexToAddress(address)

	tx, err := m.instance.RemoveValidator(opt, addr)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, m.caller.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// ApproveStreamCreation approves a user`s stream request.
func (m *ManagerClient) ApproveStreamCreation(ctx context.Context, streamID *big.Int, chunks []*big.Int) error {
	opt := m.caller.getTxOptions(0)

	// TODO: add checks so we can return informative errors when needed

	tx, err := m.instance.ApproveStreamCreation(opt, streamID, chunks)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, m.caller.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// AddInputChunk will add input chunk ids to the stream contract
func (m *ManagerClient) AddInputChunk() {

}

// AllowRefund will allow the client to refund the escrow for the given stream id.
func (m *ManagerClient) AllowRefund(ctx context.Context, streamID *big.Int) error {
	opt := m.caller.getTxOptions(0)

	// TODO: add checks so we can return informative errors when needed

	tx, err := m.instance.AllowRefund(opt, streamID)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, m.caller.client, tx)
	if err != nil {
		return err
	}

	return nil
}
