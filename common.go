package protocol

import (
	"io/ioutil"

	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
	"github.com/VideoCoin/go-videocoin/accounts/keystore"
	"github.com/VideoCoin/go-videocoin/ethclient"
)

type accountUtils interface {
	loadAccount(path string, pwd string) error
	getTxOptions() *bind.TransactOpts
}

// Account defines methods and properties required for account loading.
type Account struct {
	client *ethclient.Client
	key    *keystore.Key
}

func (a *Account) loadAccount(path string, pwd string) error {
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

func (a *Account) getTxOptions() *bind.TransactOpts {
	opts := bind.NewKeyedTransactor(a.key.PrivateKey)

	return opts
}
