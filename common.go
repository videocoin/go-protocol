package protocol

import (
	"context"
	"io/ioutil"
	"math/big"

	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
	"github.com/VideoCoin/go-videocoin/accounts/keystore"
	"github.com/VideoCoin/go-videocoin/ethclient"
)

type account interface {
	loadAccount(path string, pwd string) error
	getTxOptions() *bind.TransactOpts
	GetAccountBalance() error
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

func (c *Caller) getTxOptions() *bind.TransactOpts {
	opts := bind.NewKeyedTransactor(c.key.PrivateKey)

	return opts
}

// GetAccountBalance return the balance of the account that has been loadedin the caller
func (c *Caller) GetAccountBalance() (*big.Int, error) {
	return c.client.BalanceAt(context.Background(), c.key.Address, nil)
}
