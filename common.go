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

func (a *Caller) loadAccount(path string, pwd string) error {
	keyjson, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}

	key, e := keystore.DecryptKey(keyjson, pwd)
	if e != nil {
		return e
	}

	a.key = key

	return nil
}

func (a *Caller) getTxOptions() *bind.TransactOpts {
	opts := bind.NewKeyedTransactor(a.key.PrivateKey)

	return opts
}

// GetAccountBalance return the balance of the account that has been loadedin the caller
func (a *Caller) GetAccountBalance() (*big.Int, error) {
	return a.client.BalanceAt(context.Background(), a.key.Address, nil)

}
