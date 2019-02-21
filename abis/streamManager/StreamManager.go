// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package streamManager

import (
	"math/big"
	"strings"

	ethereum "github.com/VideoCoin/go-videocoin"
	"github.com/VideoCoin/go-videocoin/accounts/abi"
	"github.com/VideoCoin/go-videocoin/accounts/abi/bind"
	"github.com/VideoCoin/go-videocoin/common"
	"github.com/VideoCoin/go-videocoin/core/types"
	"github.com/VideoCoin/go-videocoin/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StreamManagerABI is the input ABI used to generate the binding from.
const StreamManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"streamId\",\"type\":\"uint256\"},{\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"addInputChunkId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"refundAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"allowRefund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"v\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"revokeRefund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"v\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"endStream\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"createStream\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"streamId\",\"type\":\"uint256\"},{\"name\":\"chunks\",\"type\":\"uint256[]\"}],\"name\":\"approveStreamCreation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"name\":\"approved\",\"type\":\"bool\"},{\"name\":\"refund\",\"type\":\"bool\"},{\"name\":\"ended\",\"type\":\"bool\"},{\"name\":\"client\",\"type\":\"address\"},{\"name\":\"stream\",\"type\":\"address\"},{\"name\":\"streamId\",\"type\":\"uint256\"},{\"name\":\"RTMP\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"streamId\",\"type\":\"uint256\"},{\"name\":\"RTMP\",\"type\":\"string\"},{\"name\":\"bitrates\",\"type\":\"uint256[]\"}],\"name\":\"requestStream\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"v\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"StreamRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"StreamApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"streamAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"StreamCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"RefundAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"RefundRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"streamId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"chunkId\",\"type\":\"uint256\"}],\"name\":\"InputChunkAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"streamId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"StreamEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// StreamManager is an auto generated Go binding around an Ethereum contract.
type StreamManager struct {
	StreamManagerCaller     // Read-only binding to the contract
	StreamManagerTransactor // Write-only binding to the contract
	StreamManagerFilterer   // Log filterer for contract events
}

// StreamManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StreamManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StreamManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StreamManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StreamManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StreamManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StreamManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StreamManagerSession struct {
	Contract     *StreamManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StreamManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StreamManagerCallerSession struct {
	Contract *StreamManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StreamManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StreamManagerTransactorSession struct {
	Contract     *StreamManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StreamManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StreamManagerRaw struct {
	Contract *StreamManager // Generic contract binding to access the raw methods on
}

// StreamManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StreamManagerCallerRaw struct {
	Contract *StreamManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StreamManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StreamManagerTransactorRaw struct {
	Contract *StreamManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStreamManager creates a new instance of StreamManager, bound to a specific deployed contract.
func NewStreamManager(address common.Address, backend bind.ContractBackend) (*StreamManager, error) {
	contract, err := bindStreamManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StreamManager{StreamManagerCaller: StreamManagerCaller{contract: contract}, StreamManagerTransactor: StreamManagerTransactor{contract: contract}, StreamManagerFilterer: StreamManagerFilterer{contract: contract}}, nil
}

// NewStreamManagerCaller creates a new read-only instance of StreamManager, bound to a specific deployed contract.
func NewStreamManagerCaller(address common.Address, caller bind.ContractCaller) (*StreamManagerCaller, error) {
	contract, err := bindStreamManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StreamManagerCaller{contract: contract}, nil
}

// NewStreamManagerTransactor creates a new write-only instance of StreamManager, bound to a specific deployed contract.
func NewStreamManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StreamManagerTransactor, error) {
	contract, err := bindStreamManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StreamManagerTransactor{contract: contract}, nil
}

// NewStreamManagerFilterer creates a new log filterer instance of StreamManager, bound to a specific deployed contract.
func NewStreamManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StreamManagerFilterer, error) {
	contract, err := bindStreamManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StreamManagerFilterer{contract: contract}, nil
}

// bindStreamManager binds a generic wrapper to an already deployed contract.
func bindStreamManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StreamManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StreamManager *StreamManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StreamManager.Contract.StreamManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StreamManager *StreamManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StreamManager.Contract.StreamManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StreamManager *StreamManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StreamManager.Contract.StreamManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StreamManager *StreamManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StreamManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StreamManager *StreamManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StreamManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StreamManager *StreamManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StreamManager.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_StreamManager *StreamManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StreamManager.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_StreamManager *StreamManagerSession) IsOwner() (bool, error) {
	return _StreamManager.Contract.IsOwner(&_StreamManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_StreamManager *StreamManagerCallerSession) IsOwner() (bool, error) {
	return _StreamManager.Contract.IsOwner(&_StreamManager.CallOpts)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address v) constant returns(bool)
func (_StreamManager *StreamManagerCaller) IsValidator(opts *bind.CallOpts, v common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StreamManager.contract.Call(opts, out, "isValidator", v)
	return *ret0, err
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address v) constant returns(bool)
func (_StreamManager *StreamManagerSession) IsValidator(v common.Address) (bool, error) {
	return _StreamManager.Contract.IsValidator(&_StreamManager.CallOpts, v)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address v) constant returns(bool)
func (_StreamManager *StreamManagerCallerSession) IsValidator(v common.Address) (bool, error) {
	return _StreamManager.Contract.IsValidator(&_StreamManager.CallOpts, v)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StreamManager *StreamManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StreamManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StreamManager *StreamManagerSession) Owner() (common.Address, error) {
	return _StreamManager.Contract.Owner(&_StreamManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StreamManager *StreamManagerCallerSession) Owner() (common.Address, error) {
	return _StreamManager.Contract.Owner(&_StreamManager.CallOpts)
}

// RefundAllowed is a free data retrieval call binding the contract method 0x0f514717.
//
// Solidity: function refundAllowed(uint256 streamId) constant returns(bool)
func (_StreamManager *StreamManagerCaller) RefundAllowed(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StreamManager.contract.Call(opts, out, "refundAllowed", streamId)
	return *ret0, err
}

// RefundAllowed is a free data retrieval call binding the contract method 0x0f514717.
//
// Solidity: function refundAllowed(uint256 streamId) constant returns(bool)
func (_StreamManager *StreamManagerSession) RefundAllowed(streamId *big.Int) (bool, error) {
	return _StreamManager.Contract.RefundAllowed(&_StreamManager.CallOpts, streamId)
}

// RefundAllowed is a free data retrieval call binding the contract method 0x0f514717.
//
// Solidity: function refundAllowed(uint256 streamId) constant returns(bool)
func (_StreamManager *StreamManagerCallerSession) RefundAllowed(streamId *big.Int) (bool, error) {
	return _StreamManager.Contract.RefundAllowed(&_StreamManager.CallOpts, streamId)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) constant returns(bool approved, bool refund, bool ended, address client, address stream, uint256 streamId, string RTMP)
func (_StreamManager *StreamManagerCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Approved bool
	Refund   bool
	Ended    bool
	Client   common.Address
	Stream   common.Address
	StreamId *big.Int
	RTMP     string
}, error) {
	ret := new(struct {
		Approved bool
		Refund   bool
		Ended    bool
		Client   common.Address
		Stream   common.Address
		StreamId *big.Int
		RTMP     string
	})
	out := ret
	err := _StreamManager.contract.Call(opts, out, "requests", arg0)
	return *ret, err
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) constant returns(bool approved, bool refund, bool ended, address client, address stream, uint256 streamId, string RTMP)
func (_StreamManager *StreamManagerSession) Requests(arg0 *big.Int) (struct {
	Approved bool
	Refund   bool
	Ended    bool
	Client   common.Address
	Stream   common.Address
	StreamId *big.Int
	RTMP     string
}, error) {
	return _StreamManager.Contract.Requests(&_StreamManager.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) constant returns(bool approved, bool refund, bool ended, address client, address stream, uint256 streamId, string RTMP)
func (_StreamManager *StreamManagerCallerSession) Requests(arg0 *big.Int) (struct {
	Approved bool
	Refund   bool
	Ended    bool
	Client   common.Address
	Stream   common.Address
	StreamId *big.Int
	RTMP     string
}, error) {
	return _StreamManager.Contract.Requests(&_StreamManager.CallOpts, arg0)
}

// AddInputChunkId is a paid mutator transaction binding the contract method 0x0a8c4c13.
//
// Solidity: function addInputChunkId(uint256 streamId, uint256 chunkId) returns()
func (_StreamManager *StreamManagerTransactor) AddInputChunkId(opts *bind.TransactOpts, streamId *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _StreamManager.contract.Transact(opts, "addInputChunkId", streamId, chunkId)
}

// AddInputChunkId is a paid mutator transaction binding the contract method 0x0a8c4c13.
//
// Solidity: function addInputChunkId(uint256 streamId, uint256 chunkId) returns()
func (_StreamManager *StreamManagerSession) AddInputChunkId(streamId *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _StreamManager.Contract.AddInputChunkId(&_StreamManager.TransactOpts, streamId, chunkId)
}

// AddInputChunkId is a paid mutator transaction binding the contract method 0x0a8c4c13.
//
// Solidity: function addInputChunkId(uint256 streamId, uint256 chunkId) returns()
func (_StreamManager *StreamManagerTransactorSession) AddInputChunkId(streamId *big.Int, chunkId *big.Int) (*types.Transaction, error) {
	return _StreamManager.Contract.AddInputChunkId(&_StreamManager.TransactOpts, streamId, chunkId)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address v) returns()
func (_StreamManager *StreamManagerTransactor) AddValidator(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _StreamManager.contract.Transact(opts, "addValidator", v)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address v) returns()
func (_StreamManager *StreamManagerSession) AddValidator(v common.Address) (*types.Transaction, error) {
	return _StreamManager.Contract.AddValidator(&_StreamManager.TransactOpts, v)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address v) returns()
func (_StreamManager *StreamManagerTransactorSession) AddValidator(v common.Address) (*types.Transaction, error) {
	return _StreamManager.Contract.AddValidator(&_StreamManager.TransactOpts, v)
}

// AllowRefund is a paid mutator transaction binding the contract method 0x225f6541.
//
// Solidity: function allowRefund(uint256 streamId) returns()
func (_StreamManager *StreamManagerTransactor) AllowRefund(opts *bind.TransactOpts, streamId *big.Int) (*types.Transaction, error) {
	return _StreamManager.contract.Transact(opts, "allowRefund", streamId)
}

// AllowRefund is a paid mutator transaction binding the contract method 0x225f6541.
//
// Solidity: function allowRefund(uint256 streamId) returns()
func (_StreamManager *StreamManagerSession) AllowRefund(streamId *big.Int) (*types.Transaction, error) {
	return _StreamManager.Contract.AllowRefund(&_StreamManager.TransactOpts, streamId)
}

// AllowRefund is a paid mutator transaction binding the contract method 0x225f6541.
//
// Solidity: function allowRefund(uint256 streamId) returns()
func (_StreamManager *StreamManagerTransactorSession) AllowRefund(streamId *big.Int) (*types.Transaction, error) {
	return _StreamManager.Contract.AllowRefund(&_StreamManager.TransactOpts, streamId)
}

// ApproveStreamCreation is a paid mutator transaction binding the contract method 0x70bdecb6.
//
// Solidity: function approveStreamCreation(uint256 streamId, uint256[] chunks) returns()
func (_StreamManager *StreamManagerTransactor) ApproveStreamCreation(opts *bind.TransactOpts, streamId *big.Int, chunks []*big.Int) (*types.Transaction, error) {
	return _StreamManager.contract.Transact(opts, "approveStreamCreation", streamId, chunks)
}

// ApproveStreamCreation is a paid mutator transaction binding the contract method 0x70bdecb6.
//
// Solidity: function approveStreamCreation(uint256 streamId, uint256[] chunks) returns()
func (_StreamManager *StreamManagerSession) ApproveStreamCreation(streamId *big.Int, chunks []*big.Int) (*types.Transaction, error) {
	return _StreamManager.Contract.ApproveStreamCreation(&_StreamManager.TransactOpts, streamId, chunks)
}

// ApproveStreamCreation is a paid mutator transaction binding the contract method 0x70bdecb6.
//
// Solidity: function approveStreamCreation(uint256 streamId, uint256[] chunks) returns()
func (_StreamManager *StreamManagerTransactorSession) ApproveStreamCreation(streamId *big.Int, chunks []*big.Int) (*types.Transaction, error) {
	return _StreamManager.Contract.ApproveStreamCreation(&_StreamManager.TransactOpts, streamId, chunks)
}

// CreateStream is a paid mutator transaction binding the contract method 0x551479dd.
//
// Solidity: function createStream(uint256 streamId) returns(address)
func (_StreamManager *StreamManagerTransactor) CreateStream(opts *bind.TransactOpts, streamId *big.Int) (*types.Transaction, error) {
	return _StreamManager.contract.Transact(opts, "createStream", streamId)
}

// CreateStream is a paid mutator transaction binding the contract method 0x551479dd.
//
// Solidity: function createStream(uint256 streamId) returns(address)
func (_StreamManager *StreamManagerSession) CreateStream(streamId *big.Int) (*types.Transaction, error) {
	return _StreamManager.Contract.CreateStream(&_StreamManager.TransactOpts, streamId)
}

// CreateStream is a paid mutator transaction binding the contract method 0x551479dd.
//
// Solidity: function createStream(uint256 streamId) returns(address)
func (_StreamManager *StreamManagerTransactorSession) CreateStream(streamId *big.Int) (*types.Transaction, error) {
	return _StreamManager.Contract.CreateStream(&_StreamManager.TransactOpts, streamId)
}

// EndStream is a paid mutator transaction binding the contract method 0x50d55afc.
//
// Solidity: function endStream(uint256 streamId) returns()
func (_StreamManager *StreamManagerTransactor) EndStream(opts *bind.TransactOpts, streamId *big.Int) (*types.Transaction, error) {
	return _StreamManager.contract.Transact(opts, "endStream", streamId)
}

// EndStream is a paid mutator transaction binding the contract method 0x50d55afc.
//
// Solidity: function endStream(uint256 streamId) returns()
func (_StreamManager *StreamManagerSession) EndStream(streamId *big.Int) (*types.Transaction, error) {
	return _StreamManager.Contract.EndStream(&_StreamManager.TransactOpts, streamId)
}

// EndStream is a paid mutator transaction binding the contract method 0x50d55afc.
//
// Solidity: function endStream(uint256 streamId) returns()
func (_StreamManager *StreamManagerTransactorSession) EndStream(streamId *big.Int) (*types.Transaction, error) {
	return _StreamManager.Contract.EndStream(&_StreamManager.TransactOpts, streamId)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address v) returns()
func (_StreamManager *StreamManagerTransactor) RemoveValidator(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _StreamManager.contract.Transact(opts, "removeValidator", v)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address v) returns()
func (_StreamManager *StreamManagerSession) RemoveValidator(v common.Address) (*types.Transaction, error) {
	return _StreamManager.Contract.RemoveValidator(&_StreamManager.TransactOpts, v)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address v) returns()
func (_StreamManager *StreamManagerTransactorSession) RemoveValidator(v common.Address) (*types.Transaction, error) {
	return _StreamManager.Contract.RemoveValidator(&_StreamManager.TransactOpts, v)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StreamManager *StreamManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StreamManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StreamManager *StreamManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _StreamManager.Contract.RenounceOwnership(&_StreamManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StreamManager *StreamManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StreamManager.Contract.RenounceOwnership(&_StreamManager.TransactOpts)
}

// RequestStream is a paid mutator transaction binding the contract method 0x8562faa2.
//
// Solidity: function requestStream(uint256 streamId, string RTMP, uint256[] bitrates) returns(uint256)
func (_StreamManager *StreamManagerTransactor) RequestStream(opts *bind.TransactOpts, streamId *big.Int, RTMP string, bitrates []*big.Int) (*types.Transaction, error) {
	return _StreamManager.contract.Transact(opts, "requestStream", streamId, RTMP, bitrates)
}

// RequestStream is a paid mutator transaction binding the contract method 0x8562faa2.
//
// Solidity: function requestStream(uint256 streamId, string RTMP, uint256[] bitrates) returns(uint256)
func (_StreamManager *StreamManagerSession) RequestStream(streamId *big.Int, RTMP string, bitrates []*big.Int) (*types.Transaction, error) {
	return _StreamManager.Contract.RequestStream(&_StreamManager.TransactOpts, streamId, RTMP, bitrates)
}

// RequestStream is a paid mutator transaction binding the contract method 0x8562faa2.
//
// Solidity: function requestStream(uint256 streamId, string RTMP, uint256[] bitrates) returns(uint256)
func (_StreamManager *StreamManagerTransactorSession) RequestStream(streamId *big.Int, RTMP string, bitrates []*big.Int) (*types.Transaction, error) {
	return _StreamManager.Contract.RequestStream(&_StreamManager.TransactOpts, streamId, RTMP, bitrates)
}

// RevokeRefund is a paid mutator transaction binding the contract method 0x4b51438c.
//
// Solidity: function revokeRefund(uint256 streamId) returns()
func (_StreamManager *StreamManagerTransactor) RevokeRefund(opts *bind.TransactOpts, streamId *big.Int) (*types.Transaction, error) {
	return _StreamManager.contract.Transact(opts, "revokeRefund", streamId)
}

// RevokeRefund is a paid mutator transaction binding the contract method 0x4b51438c.
//
// Solidity: function revokeRefund(uint256 streamId) returns()
func (_StreamManager *StreamManagerSession) RevokeRefund(streamId *big.Int) (*types.Transaction, error) {
	return _StreamManager.Contract.RevokeRefund(&_StreamManager.TransactOpts, streamId)
}

// RevokeRefund is a paid mutator transaction binding the contract method 0x4b51438c.
//
// Solidity: function revokeRefund(uint256 streamId) returns()
func (_StreamManager *StreamManagerTransactorSession) RevokeRefund(streamId *big.Int) (*types.Transaction, error) {
	return _StreamManager.Contract.RevokeRefund(&_StreamManager.TransactOpts, streamId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StreamManager *StreamManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StreamManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StreamManager *StreamManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StreamManager.Contract.TransferOwnership(&_StreamManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StreamManager *StreamManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StreamManager.Contract.TransferOwnership(&_StreamManager.TransactOpts, newOwner)
}

// StreamManagerInputChunkAddedIterator is returned from FilterInputChunkAdded and is used to iterate over the raw logs and unpacked data for InputChunkAdded events raised by the StreamManager contract.
type StreamManagerInputChunkAddedIterator struct {
	Event *StreamManagerInputChunkAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StreamManagerInputChunkAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamManagerInputChunkAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StreamManagerInputChunkAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StreamManagerInputChunkAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamManagerInputChunkAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamManagerInputChunkAdded represents a InputChunkAdded event raised by the StreamManager contract.
type StreamManagerInputChunkAdded struct {
	StreamId *big.Int
	ChunkId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInputChunkAdded is a free log retrieval operation binding the contract event 0x79f21a91842ed307e9fade5e30c0d0322ee9f1f64d3918e3d8b2815f02ce85d6.
//
// Solidity: event InputChunkAdded(uint256 indexed streamId, uint256 indexed chunkId)
func (_StreamManager *StreamManagerFilterer) FilterInputChunkAdded(opts *bind.FilterOpts, streamId []*big.Int, chunkId []*big.Int) (*StreamManagerInputChunkAddedIterator, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}
	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}

	logs, sub, err := _StreamManager.contract.FilterLogs(opts, "InputChunkAdded", streamIdRule, chunkIdRule)
	if err != nil {
		return nil, err
	}
	return &StreamManagerInputChunkAddedIterator{contract: _StreamManager.contract, event: "InputChunkAdded", logs: logs, sub: sub}, nil
}

// WatchInputChunkAdded is a free log subscription operation binding the contract event 0x79f21a91842ed307e9fade5e30c0d0322ee9f1f64d3918e3d8b2815f02ce85d6.
//
// Solidity: event InputChunkAdded(uint256 indexed streamId, uint256 indexed chunkId)
func (_StreamManager *StreamManagerFilterer) WatchInputChunkAdded(opts *bind.WatchOpts, sink chan<- *StreamManagerInputChunkAdded, streamId []*big.Int, chunkId []*big.Int) (event.Subscription, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}
	var chunkIdRule []interface{}
	for _, chunkIdItem := range chunkId {
		chunkIdRule = append(chunkIdRule, chunkIdItem)
	}

	logs, sub, err := _StreamManager.contract.WatchLogs(opts, "InputChunkAdded", streamIdRule, chunkIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamManagerInputChunkAdded)
				if err := _StreamManager.contract.UnpackLog(event, "InputChunkAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StreamManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StreamManager contract.
type StreamManagerOwnershipTransferredIterator struct {
	Event *StreamManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StreamManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StreamManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StreamManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamManagerOwnershipTransferred represents a OwnershipTransferred event raised by the StreamManager contract.
type StreamManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StreamManager *StreamManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StreamManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StreamManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StreamManagerOwnershipTransferredIterator{contract: _StreamManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StreamManager *StreamManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StreamManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StreamManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamManagerOwnershipTransferred)
				if err := _StreamManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StreamManagerRefundAllowedIterator is returned from FilterRefundAllowed and is used to iterate over the raw logs and unpacked data for RefundAllowed events raised by the StreamManager contract.
type StreamManagerRefundAllowedIterator struct {
	Event *StreamManagerRefundAllowed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StreamManagerRefundAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamManagerRefundAllowed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StreamManagerRefundAllowed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StreamManagerRefundAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamManagerRefundAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamManagerRefundAllowed represents a RefundAllowed event raised by the StreamManager contract.
type StreamManagerRefundAllowed struct {
	StreamId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRefundAllowed is a free log retrieval operation binding the contract event 0x48030861b818deafd7def8853bcc8a6ec6bab746521a9546b79b8baab82dce6b.
//
// Solidity: event RefundAllowed(uint256 indexed streamId)
func (_StreamManager *StreamManagerFilterer) FilterRefundAllowed(opts *bind.FilterOpts, streamId []*big.Int) (*StreamManagerRefundAllowedIterator, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _StreamManager.contract.FilterLogs(opts, "RefundAllowed", streamIdRule)
	if err != nil {
		return nil, err
	}
	return &StreamManagerRefundAllowedIterator{contract: _StreamManager.contract, event: "RefundAllowed", logs: logs, sub: sub}, nil
}

// WatchRefundAllowed is a free log subscription operation binding the contract event 0x48030861b818deafd7def8853bcc8a6ec6bab746521a9546b79b8baab82dce6b.
//
// Solidity: event RefundAllowed(uint256 indexed streamId)
func (_StreamManager *StreamManagerFilterer) WatchRefundAllowed(opts *bind.WatchOpts, sink chan<- *StreamManagerRefundAllowed, streamId []*big.Int) (event.Subscription, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _StreamManager.contract.WatchLogs(opts, "RefundAllowed", streamIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamManagerRefundAllowed)
				if err := _StreamManager.contract.UnpackLog(event, "RefundAllowed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StreamManagerRefundRevokedIterator is returned from FilterRefundRevoked and is used to iterate over the raw logs and unpacked data for RefundRevoked events raised by the StreamManager contract.
type StreamManagerRefundRevokedIterator struct {
	Event *StreamManagerRefundRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StreamManagerRefundRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamManagerRefundRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StreamManagerRefundRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StreamManagerRefundRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamManagerRefundRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamManagerRefundRevoked represents a RefundRevoked event raised by the StreamManager contract.
type StreamManagerRefundRevoked struct {
	StreamId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRefundRevoked is a free log retrieval operation binding the contract event 0xf12d3e50d9ed05d035c02b9f29e1da20004bce112d4fbe2cc0f0ba272ee5d977.
//
// Solidity: event RefundRevoked(uint256 indexed streamId)
func (_StreamManager *StreamManagerFilterer) FilterRefundRevoked(opts *bind.FilterOpts, streamId []*big.Int) (*StreamManagerRefundRevokedIterator, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _StreamManager.contract.FilterLogs(opts, "RefundRevoked", streamIdRule)
	if err != nil {
		return nil, err
	}
	return &StreamManagerRefundRevokedIterator{contract: _StreamManager.contract, event: "RefundRevoked", logs: logs, sub: sub}, nil
}

// WatchRefundRevoked is a free log subscription operation binding the contract event 0xf12d3e50d9ed05d035c02b9f29e1da20004bce112d4fbe2cc0f0ba272ee5d977.
//
// Solidity: event RefundRevoked(uint256 indexed streamId)
func (_StreamManager *StreamManagerFilterer) WatchRefundRevoked(opts *bind.WatchOpts, sink chan<- *StreamManagerRefundRevoked, streamId []*big.Int) (event.Subscription, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _StreamManager.contract.WatchLogs(opts, "RefundRevoked", streamIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamManagerRefundRevoked)
				if err := _StreamManager.contract.UnpackLog(event, "RefundRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StreamManagerStreamApprovedIterator is returned from FilterStreamApproved and is used to iterate over the raw logs and unpacked data for StreamApproved events raised by the StreamManager contract.
type StreamManagerStreamApprovedIterator struct {
	Event *StreamManagerStreamApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StreamManagerStreamApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamManagerStreamApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StreamManagerStreamApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StreamManagerStreamApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamManagerStreamApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamManagerStreamApproved represents a StreamApproved event raised by the StreamManager contract.
type StreamManagerStreamApproved struct {
	StreamId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStreamApproved is a free log retrieval operation binding the contract event 0x8fd81831397822207cf7571ac8b61f51ae7de628b05492a586bc63c3fd23bf8c.
//
// Solidity: event StreamApproved(uint256 indexed streamId)
func (_StreamManager *StreamManagerFilterer) FilterStreamApproved(opts *bind.FilterOpts, streamId []*big.Int) (*StreamManagerStreamApprovedIterator, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _StreamManager.contract.FilterLogs(opts, "StreamApproved", streamIdRule)
	if err != nil {
		return nil, err
	}
	return &StreamManagerStreamApprovedIterator{contract: _StreamManager.contract, event: "StreamApproved", logs: logs, sub: sub}, nil
}

// WatchStreamApproved is a free log subscription operation binding the contract event 0x8fd81831397822207cf7571ac8b61f51ae7de628b05492a586bc63c3fd23bf8c.
//
// Solidity: event StreamApproved(uint256 indexed streamId)
func (_StreamManager *StreamManagerFilterer) WatchStreamApproved(opts *bind.WatchOpts, sink chan<- *StreamManagerStreamApproved, streamId []*big.Int) (event.Subscription, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _StreamManager.contract.WatchLogs(opts, "StreamApproved", streamIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamManagerStreamApproved)
				if err := _StreamManager.contract.UnpackLog(event, "StreamApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StreamManagerStreamCreatedIterator is returned from FilterStreamCreated and is used to iterate over the raw logs and unpacked data for StreamCreated events raised by the StreamManager contract.
type StreamManagerStreamCreatedIterator struct {
	Event *StreamManagerStreamCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StreamManagerStreamCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamManagerStreamCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StreamManagerStreamCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StreamManagerStreamCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamManagerStreamCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamManagerStreamCreated represents a StreamCreated event raised by the StreamManager contract.
type StreamManagerStreamCreated struct {
	StreamAddress common.Address
	StreamId      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStreamCreated is a free log retrieval operation binding the contract event 0x1bd63527042f119292b792487cee2f3e2f788737aa8ce9c0b5e79a2e17bd6bab.
//
// Solidity: event StreamCreated(address indexed streamAddress, uint256 indexed streamId)
func (_StreamManager *StreamManagerFilterer) FilterStreamCreated(opts *bind.FilterOpts, streamAddress []common.Address, streamId []*big.Int) (*StreamManagerStreamCreatedIterator, error) {

	var streamAddressRule []interface{}
	for _, streamAddressItem := range streamAddress {
		streamAddressRule = append(streamAddressRule, streamAddressItem)
	}
	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _StreamManager.contract.FilterLogs(opts, "StreamCreated", streamAddressRule, streamIdRule)
	if err != nil {
		return nil, err
	}
	return &StreamManagerStreamCreatedIterator{contract: _StreamManager.contract, event: "StreamCreated", logs: logs, sub: sub}, nil
}

// WatchStreamCreated is a free log subscription operation binding the contract event 0x1bd63527042f119292b792487cee2f3e2f788737aa8ce9c0b5e79a2e17bd6bab.
//
// Solidity: event StreamCreated(address indexed streamAddress, uint256 indexed streamId)
func (_StreamManager *StreamManagerFilterer) WatchStreamCreated(opts *bind.WatchOpts, sink chan<- *StreamManagerStreamCreated, streamAddress []common.Address, streamId []*big.Int) (event.Subscription, error) {

	var streamAddressRule []interface{}
	for _, streamAddressItem := range streamAddress {
		streamAddressRule = append(streamAddressRule, streamAddressItem)
	}
	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _StreamManager.contract.WatchLogs(opts, "StreamCreated", streamAddressRule, streamIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamManagerStreamCreated)
				if err := _StreamManager.contract.UnpackLog(event, "StreamCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StreamManagerStreamEndedIterator is returned from FilterStreamEnded and is used to iterate over the raw logs and unpacked data for StreamEnded events raised by the StreamManager contract.
type StreamManagerStreamEndedIterator struct {
	Event *StreamManagerStreamEnded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StreamManagerStreamEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamManagerStreamEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StreamManagerStreamEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StreamManagerStreamEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamManagerStreamEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamManagerStreamEnded represents a StreamEnded event raised by the StreamManager contract.
type StreamManagerStreamEnded struct {
	StreamId *big.Int
	Caller   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStreamEnded is a free log retrieval operation binding the contract event 0xde05e689a2a03aa3267e2f457184c649b080aaa00013ed27d21b85bcb04901ff.
//
// Solidity: event StreamEnded(uint256 indexed streamId, address indexed caller)
func (_StreamManager *StreamManagerFilterer) FilterStreamEnded(opts *bind.FilterOpts, streamId []*big.Int, caller []common.Address) (*StreamManagerStreamEndedIterator, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _StreamManager.contract.FilterLogs(opts, "StreamEnded", streamIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &StreamManagerStreamEndedIterator{contract: _StreamManager.contract, event: "StreamEnded", logs: logs, sub: sub}, nil
}

// WatchStreamEnded is a free log subscription operation binding the contract event 0xde05e689a2a03aa3267e2f457184c649b080aaa00013ed27d21b85bcb04901ff.
//
// Solidity: event StreamEnded(uint256 indexed streamId, address indexed caller)
func (_StreamManager *StreamManagerFilterer) WatchStreamEnded(opts *bind.WatchOpts, sink chan<- *StreamManagerStreamEnded, streamId []*big.Int, caller []common.Address) (event.Subscription, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _StreamManager.contract.WatchLogs(opts, "StreamEnded", streamIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamManagerStreamEnded)
				if err := _StreamManager.contract.UnpackLog(event, "StreamEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StreamManagerStreamRequestedIterator is returned from FilterStreamRequested and is used to iterate over the raw logs and unpacked data for StreamRequested events raised by the StreamManager contract.
type StreamManagerStreamRequestedIterator struct {
	Event *StreamManagerStreamRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StreamManagerStreamRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamManagerStreamRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StreamManagerStreamRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StreamManagerStreamRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamManagerStreamRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamManagerStreamRequested represents a StreamRequested event raised by the StreamManager contract.
type StreamManagerStreamRequested struct {
	Client   common.Address
	StreamId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStreamRequested is a free log retrieval operation binding the contract event 0xcf93cb8f3f726dd083e429df912bf338cc7f82cc3344b3ab5fa960ef0357e321.
//
// Solidity: event StreamRequested(address indexed client, uint256 indexed streamId)
func (_StreamManager *StreamManagerFilterer) FilterStreamRequested(opts *bind.FilterOpts, client []common.Address, streamId []*big.Int) (*StreamManagerStreamRequestedIterator, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _StreamManager.contract.FilterLogs(opts, "StreamRequested", clientRule, streamIdRule)
	if err != nil {
		return nil, err
	}
	return &StreamManagerStreamRequestedIterator{contract: _StreamManager.contract, event: "StreamRequested", logs: logs, sub: sub}, nil
}

// WatchStreamRequested is a free log subscription operation binding the contract event 0xcf93cb8f3f726dd083e429df912bf338cc7f82cc3344b3ab5fa960ef0357e321.
//
// Solidity: event StreamRequested(address indexed client, uint256 indexed streamId)
func (_StreamManager *StreamManagerFilterer) WatchStreamRequested(opts *bind.WatchOpts, sink chan<- *StreamManagerStreamRequested, client []common.Address, streamId []*big.Int) (event.Subscription, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}

	logs, sub, err := _StreamManager.contract.WatchLogs(opts, "StreamRequested", clientRule, streamIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamManagerStreamRequested)
				if err := _StreamManager.contract.UnpackLog(event, "StreamRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StreamManagerValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the StreamManager contract.
type StreamManagerValidatorAddedIterator struct {
	Event *StreamManagerValidatorAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StreamManagerValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamManagerValidatorAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StreamManagerValidatorAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StreamManagerValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamManagerValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamManagerValidatorAdded represents a ValidatorAdded event raised by the StreamManager contract.
type StreamManagerValidatorAdded struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_StreamManager *StreamManagerFilterer) FilterValidatorAdded(opts *bind.FilterOpts, validator []common.Address) (*StreamManagerValidatorAddedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StreamManager.contract.FilterLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StreamManagerValidatorAddedIterator{contract: _StreamManager.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_StreamManager *StreamManagerFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *StreamManagerValidatorAdded, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StreamManager.contract.WatchLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamManagerValidatorAdded)
				if err := _StreamManager.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StreamManagerValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the StreamManager contract.
type StreamManagerValidatorRemovedIterator struct {
	Event *StreamManagerValidatorRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StreamManagerValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StreamManagerValidatorRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StreamManagerValidatorRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StreamManagerValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StreamManagerValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StreamManagerValidatorRemoved represents a ValidatorRemoved event raised by the StreamManager contract.
type StreamManagerValidatorRemoved struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_StreamManager *StreamManagerFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, validator []common.Address) (*StreamManagerValidatorRemovedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StreamManager.contract.FilterLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StreamManagerValidatorRemovedIterator{contract: _StreamManager.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_StreamManager *StreamManagerFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *StreamManagerValidatorRemoved, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StreamManager.contract.WatchLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StreamManagerValidatorRemoved)
				if err := _StreamManager.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
