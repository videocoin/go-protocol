package protocol

import (
	"context"
	"math/big"

	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
)

// TranscoderClient is a wrapper for the transcoder calls.
type TranscoderClient struct {
	ManagerContract
	Caller
}

// NewTranscoderClient creates a TranscoderClient instance
func NewTranscoderClient(url string, addr string, keyfilePath string, pwd string) (*TranscoderClient, error) {
	contract, caller, err := NewManagerContract(url, addr, keyfilePath, pwd)
	if err != nil {
		return nil, err
	}

	c := &TranscoderClient{*contract, *caller}

	return c, nil
}

// SubmitProof sends
func (c *TranscoderClient) SubmitProof(ctx context.Context, streamID *big.Int, bitrate *big.Int, inputChunkID *big.Int, outputChunkID *big.Int) error {
	stream, err := c.GetStreamContract(streamID, c.client)
	if err != nil {
		return err
	}

	proof := big.NewInt(0)

	// TODO: add checks

	opt := c.getTxOptions(0)
	tx, err := stream.instance.SubmitProof(opt, bitrate, inputChunkID, proof, outputChunkID)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, c.client, tx)
	if err != nil {
		return err
	}

	return nil
}
