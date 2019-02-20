package protocol

import (
	"context"
	"fmt"
	"math/big"

	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
)

// VerifierClient is a wrapper for the transcoder calls.
type VerifierClient struct {
	ManagerContract
	Caller
}

// NewVerifierClient creates a VerifierClient instance
func NewVerifierClient(url string, addr string, keyfilePath string, pwd string) (*VerifierClient, error) {
	contract, caller, err := NewManagerContract(url, addr, keyfilePath, pwd)
	if err != nil {
		return nil, err
	}

	t := &VerifierClient{*contract, *caller}

	isVerifier, err := t.IsValidator(t.key.Address)
	if err != nil {
		return nil, err
	}

	if !isVerifier {
		return nil, fmt.Errorf("can`t create verifier. address: %s is not a validator", addr)
	}

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
