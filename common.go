package protocol

import (
	"context"
	"io/ioutil"
	"math/big"

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
	addr     common.Address
}

// GetManagerContractAddr returns manager smart contract address
func (m *ManagerContract) GetManagerContractAddr() common.Address {
	return m.addr
}

// Caller defines methods and properties required for account loading.
type Caller struct {
	client *ethclient.Client
	key    *keystore.Key
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
