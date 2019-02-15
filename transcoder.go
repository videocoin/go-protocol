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
	instance *stream.Stream
	acc      Account
}

// NewTranscoderClient creates a TranscoderClient instance
func NewTranscoderClient(url string, addr string, keyfilePath string, pwd string) (*TranscoderClient, error) {
	streamAddress := common.HexToAddress(addr)

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	acc := Account{client: client}

	err = acc.loadAccount(keyfilePath, pwd)
	if err != nil {
		return nil, err
	}

	instance, err := stream.NewStream(streamAddress, client)
	if err != nil {
		return nil, err
	}

	t := &TranscoderClient{
		instance: instance,
		acc:      acc,
	}

	return t, nil
}

// SubmitProof sends
func (t *TranscoderClient) SubmitProof(ctx context.Context, bitrate *big.Int, inputChunkID *big.Int, outputChunkID *big.Int) error {
	opt := t.acc.getTxOptions()
	proof := big.NewInt(0)

	// TODO: add checks

	tx, err := t.instance.SubmitProof(opt, bitrate, inputChunkID, proof, outputChunkID)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(nil, t.acc.client, tx)
	if err != nil {
		return err
	}

	return nil
}
