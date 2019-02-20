package protocol

import (
	"context"
	"math/big"

	"github.com/VideoCoin/go-protocol/abis/stream"
	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
	"github.com/VideoCoin/go-videocoin/common"
	"github.com/VideoCoin/go-videocoin/ethclient"
)

// TranscoderClient is a wrapper for the transcoder calls.
type TranscoderClient struct {
	StreamContract
	Caller
}

// NewTranscoderClient creates a TranscoderClient instance
func NewTranscoderClient(url string, addr string, keyfilePath string, pwd string) (*TranscoderClient, error) {
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

	instance, err := stream.NewStream(streamAddress, client)
	if err != nil {
		return nil, err
	}

	contract := StreamContract{instance, streamAddress}

	t := &TranscoderClient{contract, caller}

	return t, nil
}

// SubmitProof sends
func (t *TranscoderClient) SubmitProof(ctx context.Context, bitrate *big.Int, inputChunkID *big.Int, outputChunkID *big.Int) error {
	opt := t.getTxOptions(0)
	proof := big.NewInt(0)

	// TODO: add checks

	tx, err := t.instance.SubmitProof(opt, bitrate, inputChunkID, proof, outputChunkID)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, t.client, tx)
	if err != nil {
		return err
	}

	return nil
}
