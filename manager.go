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

	c := &ManagerClient{*contract, *caller}

	isOwner, err := c.instance.IsOwner(&bind.CallOpts{From: c.key.Address})
	if err != nil {
		return nil, err
	}

	if !isOwner {
		owner, err := c.instance.Owner(&bind.CallOpts{})
		err = fmt.Errorf("Account provided %s, is not manager. Real manager is: %s ", c.key.Address.String(), owner.String())
		return nil, err
	}

	return c, nil
}

// AddValidator adds a new address to the validator list held by the StreamManager smart contract.
func (c *ManagerClient) AddValidator(ctx context.Context, address string) error {
	addr := common.HexToAddress(address)
	isValidator, err := c.IsValidator(addr)
	if err != nil {
		return err
	}

	if isValidator {
		return fmt.Errorf("address: %s is already a validator", address)
	}

	opt := c.getTxOptions(0)
	tx, err := c.instance.AddValidator(opt, addr)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, c.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// RemoveValidator removes an address from the validator list held by the StreamManager smart contract.
func (c *ManagerClient) RemoveValidator(ctx context.Context, address string) error {
	addr := common.HexToAddress(address)
	isValidator, err := c.IsValidator(addr)
	if err != nil {
		return err
	}

	if !isValidator {
		return fmt.Errorf("address: %s is not a validator", address)
	}

	opt := c.getTxOptions(0)
	tx, err := c.instance.RemoveValidator(opt, addr)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, c.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// ApproveStreamCreation approves a user`s stream request.
func (c *ManagerClient) ApproveStreamCreation(ctx context.Context, streamID *big.Int, chunks []*big.Int) error {
	req, err := c.instance.Requests(&bind.CallOpts{}, streamID)
	if err != nil {
		return err
	}

	if req.Client.Big().Cmp(big.NewInt(0)) == 0 {
		return fmt.Errorf("stream request ID: %s does not exist", streamID.String())
	}

	opt := c.getTxOptions(0)
	tx, err := c.instance.ApproveStreamCreation(opt, streamID, chunks)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, c.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// AddInputChunk will add input chunk ids to the stream contract
func (c *ManagerClient) AddInputChunk(ctx context.Context, streamID *big.Int, chunkID *big.Int) error {
	req, err := c.instance.Requests(&bind.CallOpts{}, streamID)
	if err != nil {
		return err
	}

	if req.Stream.Big().Cmp(big.NewInt(0)) == 0 {
		return fmt.Errorf("stream ID: %s does not exist", streamID.String())
	}

	// add check for already existing input chunkid

	opt := c.getTxOptions(0)
	tx, err := c.instance.AddInputChunkId(opt, streamID, chunkID)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, c.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// AllowRefund will give the client permission to refund the escrow for the given stream id.
func (c *ManagerClient) AllowRefund(ctx context.Context, streamID *big.Int) error {
	req, err := c.instance.Requests(&bind.CallOpts{}, streamID)
	if err != nil {
		return err
	}

	if req.Client.Big().Cmp(big.NewInt(0)) == 0 {
		return fmt.Errorf("stream ID: %s does not exist", streamID.String())
	}

	if req.Refund {
		return fmt.Errorf("refund for stream ID: %s already allowed", streamID.String())
	}

	opt := c.getTxOptions(0)
	tx, err := c.instance.AllowRefund(opt, streamID)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, c.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// RevokeRefund will revoke the client permission to refund the escrow for the given stream id.
func (c *ManagerClient) RevokeRefund(ctx context.Context, streamID *big.Int) error {
	req, err := c.instance.Requests(&bind.CallOpts{}, streamID)
	if err != nil {
		return err
	}

	if req.Client.Big().Cmp(big.NewInt(0)) == 0 {
		return fmt.Errorf("stream ID: %s does not exist", streamID.String())
	}

	opt := c.getTxOptions(0)
	tx, err := c.instance.RevokeRefund(opt, streamID)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, c.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// RefundAllowed queries whether a refund is allowed for the given stream id.
func (c *ManagerClient) RefundAllowed(ctx context.Context, streamID *big.Int) (bool, error) {

	isAllowed, err := c.instance.RefundAllowed(&bind.CallOpts{}, streamID)
	if err != nil {
		return false, err
	}

	return isAllowed, nil
}
