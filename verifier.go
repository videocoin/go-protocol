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

// VerifierClient is a wrapper for the transcoder calls.
type VerifierClient struct {
	ManagerContract
	Caller
}

// NewVerifierClient creates a VerifierClient instance
func NewVerifierClient(url string, addr string, keyfilePath string, pwd string) (*VerifierClient, error) {
	streamAddress := common.HexToAddress(addr)

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	caller := Caller{client: client}

	err = caller.loadAccount(keyfilePath, pwd)
	if err != nil {
		return nil, err
	}

	instance, err := streamManager.NewStreamManager(streamAddress, client)
	if err != nil {
		return nil, err
	}

	contract := ManagerContract{instance, streamAddress}

	isVerifier, err := contract.IsValidator(streamAddress)
	if err != nil {
		return nil, err
	}

	if !isVerifier {
		return nil, fmt.Errorf("cant create verifier. address: %s is not a validator", addr)
	}

	// get all streams

	t := &VerifierClient{contract, caller}

	return t, nil
}

// ValidateProof ...
func (v *VerifierClient) ValidateProof(ctx context.Context, streamID *big.Int, bitrate *big.Int, inputChunkID *big.Int) error {
	stream, err := v.GetStreamContract(streamID, v.client)
	if err != nil {
		return err
	}

	// TODO: add checks

	opt := v.getTxOptions(0)
	tx, err := stream.instance.ValidateProof(opt, bitrate, inputChunkID)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, v.client, tx)
	if err != nil {
		return err
	}

	return nil
}
