package protocol

import (
	"context"
	"fmt"
	"math/big"

	"github.com/VideoCoin/go-protocol/abis/stream"
	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
)

// UserClient is a wrapper for the User calls.
type UserClient struct {
	ManagerContract
	Caller
}

// NewUserClient creates a UserClient instance
func NewUserClient(url string, addr string, keyfilePath string, pwd string) (*UserClient, error) {
	contract, caller, err := NewManagerContract(url, addr, keyfilePath, pwd)
	if err != nil {
		return nil, err
	}

	u := &UserClient{*contract, *caller}

	return u, nil
}

// RequestStream creates a new stream request from the user with the manager smart contract
func (u *UserClient) RequestStream(ctx context.Context, streamID *big.Int, RTMP string, bitrates []*big.Int) error {
	opt := u.getTxOptions(0)

	// TODO: check that request has not already been submitted

	tx, err := u.instance.RequestStream(opt, streamID, RTMP, bitrates)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, u.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// CreateStream creates a new stream after the user`s request has been approved
func (u *UserClient) CreateStream(ctx context.Context, streamID *big.Int, funds *big.Int) error {
	opt := u.getTxOptions(0)
	opt.Value = funds

	// TODO: check that the request has been approved

	tx, err := u.instance.CreateStream(opt, streamID)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, u.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// ClaimRefund refunds the user the funds that have been escrowed (provided the refund was approved)
func (u *UserClient) ClaimRefund(ctx context.Context, streamID *big.Int) error {
	req, err := u.instance.Requests(&bind.CallOpts{}, streamID)
	if err != nil {
		return err
	}

	if !req.Refund {
		return fmt.Errorf("refund for string ID: %s is not allowed", streamID.String())
	}

	myStream, err := stream.NewStream(req.Stream, u.client)
	if err != nil {
		return err
	}

	opt := u.getTxOptions(0)

	tx, err := myStream.Refund(opt)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, u.client, tx)
	if err != nil {
		return err
	}

	return nil
}
