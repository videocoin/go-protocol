package protocol

import (
	"context"
	"fmt"
	"math/big"

	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
	"github.com/VideoCoin/go-videocoin/common"
)

// ManagerClient is a wrapper for the stream manager calls.
type ManagerClient struct {
	ManagerContract
	Caller
}

// NewManagerClient creates a ManagerClient instance
func NewManagerClient(url string, addr string, keyfilePath string, pwd string) (*ManagerClient, error) {
	contract, caller, err := NewManagerContract(url, addr, keyfilePath, pwd)
	if err != nil {
		return nil, err
	}

	m := &ManagerClient{*contract, *caller}

	isOwner, err := m.instance.IsOwner(&bind.CallOpts{From: m.key.Address})
	if err != nil {
		return nil, err
	}

	if !isOwner {
		owner, err := m.instance.Owner(&bind.CallOpts{})
		err = fmt.Errorf("Account provided %s, is not manager. Real manager is: %s ", m.key.Address.String(), owner.String())
		return nil, err
	}

	return m, nil
}

// AddValidator adds a new address to the validator list held by the StreamManager smart contract.
func (m *ManagerClient) AddValidator(ctx context.Context, address string) error {
	addr := common.HexToAddress(address)
	isValidator, err := m.IsValidator(addr)
	if err != nil {
		return err
	}

	if isValidator {
		return fmt.Errorf("address: %s is already a validator", address)
	}

	opt := m.getTxOptions(0)
	tx, err := m.instance.AddValidator(opt, addr)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, m.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// RemoveValidator removes an address from the validator list held by the StreamManager smart contract.
func (m *ManagerClient) RemoveValidator(ctx context.Context, address string) error {
	addr := common.HexToAddress(address)
	isValidator, err := m.IsValidator(addr)
	if err != nil {
		return err
	}

	if !isValidator {
		return fmt.Errorf("address: %s is not a validator", address)
	}

	opt := m.getTxOptions(0)
	tx, err := m.instance.RemoveValidator(opt, addr)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, m.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// ApproveStreamCreation approves a user`s stream request.
func (m *ManagerClient) ApproveStreamCreation(ctx context.Context, streamID *big.Int, chunks []*big.Int) error {
	req, err := m.instance.Requests(&bind.CallOpts{}, streamID)
	if err != nil {
		return err
	}

	if req.Client.Big().Cmp(big.NewInt(0)) == 0 {
		return fmt.Errorf("stream request ID: %s does not exist", streamID.String())
	}

	opt := m.getTxOptions(0)
	tx, err := m.instance.ApproveStreamCreation(opt, streamID, chunks)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, m.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// AddInputChunk will add input chunk ids to the stream contract
func (m *ManagerClient) AddInputChunk(ctx context.Context, streamID *big.Int, chunkID *big.Int) error {
	req, err := m.instance.Requests(&bind.CallOpts{}, streamID)
	if err != nil {
		return err
	}

	if req.Stream.Big().Cmp(big.NewInt(0)) == 0 {
		return fmt.Errorf("stream ID: %s does not exist", streamID.String())
	}

	// add check for already existing input chunkid

	opt := m.getTxOptions(0)
	tx, err := m.instance.AddInputChunkId(opt, streamID, chunkID)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, m.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// AllowRefund will give the client permission to refund the escrow for the given stream id.
func (m *ManagerClient) AllowRefund(ctx context.Context, streamID *big.Int) error {
	req, err := m.instance.Requests(&bind.CallOpts{}, streamID)
	if err != nil {
		return err
	}

	if req.Client.Big().Cmp(big.NewInt(0)) == 0 {
		return fmt.Errorf("stream ID: %s does not exist", streamID.String())
	}

	if req.Refund {
		return fmt.Errorf("refund for stream ID: %s already allowed", streamID.String())
	}

	opt := m.getTxOptions(0)
	tx, err := m.instance.AllowRefund(opt, streamID)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, m.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// RevokeRefund will revoke the client permission to refund the escrow for the given stream id.
func (m *ManagerClient) RevokeRefund(ctx context.Context, streamID *big.Int) error {
	req, err := m.instance.Requests(&bind.CallOpts{}, streamID)
	if err != nil {
		return err
	}

	if req.Client.Big().Cmp(big.NewInt(0)) == 0 {
		return fmt.Errorf("stream ID: %s does not exist", streamID.String())
	}

	opt := m.getTxOptions(0)
	tx, err := m.instance.RevokeRefund(opt, streamID)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, m.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// RefundAllowed queries whether a refund is allowed for the given stream id.
func (m *ManagerClient) RefundAllowed(ctx context.Context, streamID *big.Int) (bool, error) {

	isAllowed, err := m.instance.RefundAllowed(&bind.CallOpts{}, streamID)
	if err != nil {
		return false, err
	}

	return isAllowed, nil
}
