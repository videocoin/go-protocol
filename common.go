package protocol

import (
	"context"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/VideoCoin/go-protocol/abis/stream"
	"github.com/VideoCoin/go-protocol/abis/streamManager"
	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
	"github.com/VideoCoin/go-videocoin/accounts/keystore"
	"github.com/VideoCoin/go-videocoin/common"
	"github.com/VideoCoin/go-videocoin/ethclient"
)

type account interface {
	loadAccount(path string, pwd string) error
	getTxOptions(gasLimit int) *bind.TransactOpts
	GetAccountBalance() error
}

// ManagerContract wraps the manager smart contract & some of its methods
type ManagerContract struct {
	instance *streamManager.StreamManager
	common.Address
}

// IsValidator returns true if address is registerred as validator
func (m *ManagerContract) IsValidator() (bool, error) {
	isValidator, err := m.instance.IsValidator(&bind.CallOpts{}, m.Address)
	if err != nil {
		return false, err
	}

	return isValidator, nil
}

// GetStreams returns a list of all streams registered with the manager.
func (m *ManagerContract) GetStreams(fromBlock uint64) error {
	opts := bind.FilterOpts{
		Start: fromBlock,
	}

	iterator, err := m.instance.FilterStreamRequested(&opts, nil, nil)
	if err != nil {
		return err
	}

	defer iterator.Close()

	for iterator.Next() {
		event := iterator.Event
		log.Print("Clients requesting stream: " + event.Client.String())
	}

	return nil
}

// StreamContract wraps a stream smart contract & some of its methods
type StreamContract struct {
	instance *stream.Stream
	common.Address
}

// Caller implements methods and properties required for account loading.
type Caller struct {
	client *ethclient.Client
	key    *keystore.Key
}

// GetCallerAddr returns manager smart contract address
func (c *Caller) GetCallerAddr() common.Address {
	return c.key.Address
}

func (c *Caller) loadAccount(path string, pwd string) error {
	keyjson, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}

	key, e := keystore.DecryptKey(keyjson, pwd)
	if e != nil {
		return e
	}

	c.key = key

	return nil
}

func (c *Caller) getTxOptions(gasLimit int) *bind.TransactOpts {
	opts := bind.NewKeyedTransactor(c.key.PrivateKey)
	if gasLimit == 0 {
		opts.GasLimit = uint64(3000000)
	}

	return opts
}

// GetAccountBalance return the balance of the account that has been loadedin the caller
func (c *Caller) GetAccountBalance() (*big.Int, error) {
	return c.client.BalanceAt(context.Background(), c.key.Address, nil)
}
