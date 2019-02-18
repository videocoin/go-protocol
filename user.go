package protocol

import (
	"context"
	"math/big"

	"github.com/VideoCoin/go-protocol/abis/streamManager"
	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
	"github.com/VideoCoin/go-videocoin/common"
	"github.com/VideoCoin/go-videocoin/ethclient"
)

// UserClient is a wrapper for the User calls.
type UserClient struct {
	instance *streamManager.StreamManager
	caller   Caller
}

// NewUserClient creates a UserClient instance
func NewUserClient(url string, addr string, keyfilePath string, pwd string) (*UserClient, error) {
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

	instance, err := streamManager.NewStreamManager(managerAddress, client)
	if err != nil {
		return nil, err
	}

	t := &UserClient{
		instance: instance,
		caller:   caller,
	}

	return t, nil
}

// RequestStream creates a new stream request from the user with the manager smart contract
func (u *UserClient) RequestStream(ctx context.Context, streamID *big.Int, RTMP string, bitrates []*big.Int) error {
	opt := u.caller.getTxOptions()

	// TODO: check that address is not already validator

	tx, err := u.instance.RequestStream(opt, streamID, RTMP, bitrates)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, u.caller.client, tx)
	if err != nil {
		return err
	}

	return nil
}
