package protocol

import (
	"context"
	"fmt"
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
	streams map[string]*StreamContract
}

// NewManagerContract creates a ManagerContract
func NewManagerContract(url string, addr string, keyfilePath string, pwd string) (*ManagerContract, *Caller, error) {
	managerAddress := common.HexToAddress(addr)

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, nil, err
	}

	caller := &Caller{client: client}
	err = caller.loadAccount(keyfilePath, pwd)
	if err != nil {
		return nil, nil, err
	}

	instance, err := streamManager.NewStreamManager(managerAddress, client)
	if err != nil {
		return nil, nil, err
	}

	contract := &ManagerContract{instance, managerAddress, make(map[string]*StreamContract)}

	return contract, caller, nil
}

// IsValidator returns true if address is registerred as validator
func (m *ManagerContract) IsValidator(addr common.Address) (bool, error) {
	isValidator, err := m.instance.IsValidator(&bind.CallOpts{}, addr)
	if err != nil {
		return false, err
	}

	return isValidator, nil
}

// GetStreamContract returns a StreamContract instance for the given stream id.
func (m *ManagerContract) GetStreamContract(streamID *big.Int, client *ethclient.Client) (*StreamContract, error) {
	key := streamID.String()

	if val, ok := m.streams[key]; ok {
		return val, nil
	}

	req, err := m.instance.Requests(&bind.CallOpts{}, streamID)
	if err != nil {
		return nil, err
	}

	streamAddress := req.Stream
	if streamAddress.Big().Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("stream ID: %s does not exist", streamID.String())
	}

	instance, err := stream.NewStream(streamAddress, client)
	if err != nil {
		return nil, fmt.Errorf("could not create stream instance")
	}

	contract := &StreamContract{instance, streamAddress}

	m.streams[key] = contract

	return contract, nil
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
