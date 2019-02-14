package protocol

import (
	"io/ioutil"
	"math/big"

	"github.com/VideoCoin/go-protocol/abis/streamManager"
	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
	"github.com/VideoCoin/go-videocoin/accounts/keystore"
	"github.com/VideoCoin/go-videocoin/common"
	"github.com/VideoCoin/go-videocoin/ethclient"
)

// Manager is a wrapper for the stream manager calls.
type Manager struct {
	instance *streamManager.StreamManager
	client   *ethclient.Client
	key      *keystore.Key
}

// New creates a Manager instance
func New(url string, addr string, keyfilePath string, pwd string) (*Manager, error) {
	managerAddress := common.HexToAddress(addr)

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	managerInstance, err := streamManager.NewStreamManager(managerAddress, client)
	if err != nil {
		return nil, err
	}

	m := &Manager{
		instance: managerInstance,
		client:   client,
	}

	err = m.loadAccount(keyfilePath, pwd)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// AddValidator adds a new address to the validator map in the StreamManager smart contract.
func (m *Manager) AddValidator() {

}

// RemoveValidator removes an address from the validator map in the StreamManager smart contract.
func (m *Manager) RemoveValidator() {

}

// ApproveStreamCreation approves a user`s stream request.
func (m *Manager) ApproveStreamCreation(streamID *big.Int, chunks []*big.Int) error {
	opt := m.getTxOptions()

	tx, err := m.instance.ApproveStreamCreation(opt, streamID, chunks)
	if err != nil {
		return err
	}

	_ = tx

	return nil
}

// addInputChunk
func (m *Manager) addInputChunk() {

}

func (m *Manager) loadAccount(path string, pwd string) error {
	keyjson, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}

	key, e := keystore.DecryptKey(keyjson, pwd)
	if e != nil {
		return e
	}

	m.key = key

	return nil
}

func (m *Manager) getTxOptions() *bind.TransactOpts {
	opts := bind.NewKeyedTransactor(m.key.PrivateKey)

	return opts
}
