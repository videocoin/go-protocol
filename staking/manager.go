// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

// StakingManagerUnbondingRequest is an auto generated low-level Go binding around an user-defined struct.
type StakingManagerUnbondingRequest struct {
	Transcoder common.Address
	Timestamp  *big.Int
	Amount     *big.Int
}

// StakingManagerABI is the input ABI used to generate the binding from.
const StakingManagerABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minSelfStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transcoderApprovalPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unbondingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashRate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_slashPoolAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"}],\"name\":\"Jailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Rebonded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"}],\"name\":\"TranscoderRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"unbondingID\",\"type\":\"uint256\"}],\"name\":\"UnbondingRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"}],\"name\":\"Unjailed\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transcoderAddr\",\"type\":\"uint256\"}],\"name\":\"collectReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"unbondingTail\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"}],\"name\":\"getSlashableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"getTrancoderSlashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"getTranscoderState\",\"outputs\":[{\"internalType\":\"enumStakingManager.TranscoderState\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unbondingID\",\"type\":\"uint256\"}],\"name\":\"getUnbondingRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingManager.UnbondingRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"isJailed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unbondingID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"rebond\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"registerTranscoder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestUnbonding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setApprovalPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"name\":\"setCapacity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setSelfMinStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setSlashFundAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zone\",\"type\":\"uint256\"}],\"name\":\"setZone\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transcoderApprovalPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transcoders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zone\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"unjail\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unbondingID\",\"type\":\"uint256\"}],\"name\":\"withdrawStake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StakingManager is an auto generated Go binding around an Ethereum contract.
type StakingManager struct {
	StakingManagerCaller     // Read-only binding to the contract
	StakingManagerTransactor // Write-only binding to the contract
	StakingManagerFilterer   // Log filterer for contract events
}

// StakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingManagerSession struct {
	Contract     *StakingManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingManagerCallerSession struct {
	Contract *StakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingManagerTransactorSession struct {
	Contract     *StakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingManagerRaw struct {
	Contract *StakingManager // Generic contract binding to access the raw methods on
}

// StakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingManagerCallerRaw struct {
	Contract *StakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingManagerTransactorRaw struct {
	Contract *StakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingManager creates a new instance of StakingManager, bound to a specific deployed contract.
func NewStakingManager(address common.Address, backend bind.ContractBackend) (*StakingManager, error) {
	contract, err := bindStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// NewStakingManagerCaller creates a new read-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*StakingManagerCaller, error) {
	contract, err := bindStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerCaller{contract: contract}, nil
}

// NewStakingManagerTransactor creates a new write-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingManagerTransactor, error) {
	contract, err := bindStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerTransactor{contract: contract}, nil
}

// NewStakingManagerFilterer creates a new log filterer instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingManagerFilterer, error) {
	contract, err := bindStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingManagerFilterer{contract: contract}, nil
}

// bindStakingManager binds a generic wrapper to an already deployed contract.
func bindStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.StakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transact(opts, method, params...)
}

// Delegators is a free data retrieval call binding the contract method 0x8d23fc61.
//
// Solidity: function delegators(address ) constant returns(uint256 unbondingTail)
func (_StakingManager *StakingManagerCaller) Delegators(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "delegators", arg0)
	return *ret0, err
}

// Delegators is a free data retrieval call binding the contract method 0x8d23fc61.
//
// Solidity: function delegators(address ) constant returns(uint256 unbondingTail)
func (_StakingManager *StakingManagerSession) Delegators(arg0 common.Address) (*big.Int, error) {
	return _StakingManager.Contract.Delegators(&_StakingManager.CallOpts, arg0)
}

// Delegators is a free data retrieval call binding the contract method 0x8d23fc61.
//
// Solidity: function delegators(address ) constant returns(uint256 unbondingTail)
func (_StakingManager *StakingManagerCallerSession) Delegators(arg0 common.Address) (*big.Int, error) {
	return _StakingManager.Contract.Delegators(&_StakingManager.CallOpts, arg0)
}

// GetDelegatorStake is a free data retrieval call binding the contract method 0x5028e2e1.
//
// Solidity: function getDelegatorStake(address transcoderAddr, address delegAddr) constant returns(uint256)
func (_StakingManager *StakingManagerCaller) GetDelegatorStake(opts *bind.CallOpts, transcoderAddr common.Address, delegAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getDelegatorStake", transcoderAddr, delegAddr)
	return *ret0, err
}

// GetDelegatorStake is a free data retrieval call binding the contract method 0x5028e2e1.
//
// Solidity: function getDelegatorStake(address transcoderAddr, address delegAddr) constant returns(uint256)
func (_StakingManager *StakingManagerSession) GetDelegatorStake(transcoderAddr common.Address, delegAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetDelegatorStake(&_StakingManager.CallOpts, transcoderAddr, delegAddr)
}

// GetDelegatorStake is a free data retrieval call binding the contract method 0x5028e2e1.
//
// Solidity: function getDelegatorStake(address transcoderAddr, address delegAddr) constant returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetDelegatorStake(transcoderAddr common.Address, delegAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetDelegatorStake(&_StakingManager.CallOpts, transcoderAddr, delegAddr)
}

// GetSelfStake is a free data retrieval call binding the contract method 0xfece707d.
//
// Solidity: function getSelfStake(address _addr) constant returns(uint256)
func (_StakingManager *StakingManagerCaller) GetSelfStake(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getSelfStake", _addr)
	return *ret0, err
}

// GetSelfStake is a free data retrieval call binding the contract method 0xfece707d.
//
// Solidity: function getSelfStake(address _addr) constant returns(uint256)
func (_StakingManager *StakingManagerSession) GetSelfStake(_addr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetSelfStake(&_StakingManager.CallOpts, _addr)
}

// GetSelfStake is a free data retrieval call binding the contract method 0xfece707d.
//
// Solidity: function getSelfStake(address _addr) constant returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetSelfStake(_addr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetSelfStake(&_StakingManager.CallOpts, _addr)
}

// GetSlashableAmount is a free data retrieval call binding the contract method 0x8efc97a1.
//
// Solidity: function getSlashableAmount(address transcoderAddr, address delegatorAddr) constant returns(uint256)
func (_StakingManager *StakingManagerCaller) GetSlashableAmount(opts *bind.CallOpts, transcoderAddr common.Address, delegatorAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getSlashableAmount", transcoderAddr, delegatorAddr)
	return *ret0, err
}

// GetSlashableAmount is a free data retrieval call binding the contract method 0x8efc97a1.
//
// Solidity: function getSlashableAmount(address transcoderAddr, address delegatorAddr) constant returns(uint256)
func (_StakingManager *StakingManagerSession) GetSlashableAmount(transcoderAddr common.Address, delegatorAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetSlashableAmount(&_StakingManager.CallOpts, transcoderAddr, delegatorAddr)
}

// GetSlashableAmount is a free data retrieval call binding the contract method 0x8efc97a1.
//
// Solidity: function getSlashableAmount(address transcoderAddr, address delegatorAddr) constant returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetSlashableAmount(transcoderAddr common.Address, delegatorAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetSlashableAmount(&_StakingManager.CallOpts, transcoderAddr, delegatorAddr)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _addr) constant returns(uint256)
func (_StakingManager *StakingManagerCaller) GetTotalStake(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getTotalStake", _addr)
	return *ret0, err
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _addr) constant returns(uint256)
func (_StakingManager *StakingManagerSession) GetTotalStake(_addr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetTotalStake(&_StakingManager.CallOpts, _addr)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _addr) constant returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetTotalStake(_addr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetTotalStake(&_StakingManager.CallOpts, _addr)
}

// GetTrancoderSlashes is a free data retrieval call binding the contract method 0xec69a1bb.
//
// Solidity: function getTrancoderSlashes(address transcoderAddr) constant returns(uint256)
func (_StakingManager *StakingManagerCaller) GetTrancoderSlashes(opts *bind.CallOpts, transcoderAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getTrancoderSlashes", transcoderAddr)
	return *ret0, err
}

// GetTrancoderSlashes is a free data retrieval call binding the contract method 0xec69a1bb.
//
// Solidity: function getTrancoderSlashes(address transcoderAddr) constant returns(uint256)
func (_StakingManager *StakingManagerSession) GetTrancoderSlashes(transcoderAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetTrancoderSlashes(&_StakingManager.CallOpts, transcoderAddr)
}

// GetTrancoderSlashes is a free data retrieval call binding the contract method 0xec69a1bb.
//
// Solidity: function getTrancoderSlashes(address transcoderAddr) constant returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetTrancoderSlashes(transcoderAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetTrancoderSlashes(&_StakingManager.CallOpts, transcoderAddr)
}

// GetTranscoderState is a free data retrieval call binding the contract method 0xe71824bc.
//
// Solidity: function getTranscoderState(address transcoderAddr) constant returns(uint8)
func (_StakingManager *StakingManagerCaller) GetTranscoderState(opts *bind.CallOpts, transcoderAddr common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getTranscoderState", transcoderAddr)
	return *ret0, err
}

// GetTranscoderState is a free data retrieval call binding the contract method 0xe71824bc.
//
// Solidity: function getTranscoderState(address transcoderAddr) constant returns(uint8)
func (_StakingManager *StakingManagerSession) GetTranscoderState(transcoderAddr common.Address) (uint8, error) {
	return _StakingManager.Contract.GetTranscoderState(&_StakingManager.CallOpts, transcoderAddr)
}

// GetTranscoderState is a free data retrieval call binding the contract method 0xe71824bc.
//
// Solidity: function getTranscoderState(address transcoderAddr) constant returns(uint8)
func (_StakingManager *StakingManagerCallerSession) GetTranscoderState(transcoderAddr common.Address) (uint8, error) {
	return _StakingManager.Contract.GetTranscoderState(&_StakingManager.CallOpts, transcoderAddr)
}

// GetUnbondingRequest is a free data retrieval call binding the contract method 0x9e79b122.
//
// Solidity: function getUnbondingRequest(address delegatorAddr, uint256 unbondingID) constant returns(StakingManagerUnbondingRequest)
func (_StakingManager *StakingManagerCaller) GetUnbondingRequest(opts *bind.CallOpts, delegatorAddr common.Address, unbondingID *big.Int) (StakingManagerUnbondingRequest, error) {
	var (
		ret0 = new(StakingManagerUnbondingRequest)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getUnbondingRequest", delegatorAddr, unbondingID)
	return *ret0, err
}

// GetUnbondingRequest is a free data retrieval call binding the contract method 0x9e79b122.
//
// Solidity: function getUnbondingRequest(address delegatorAddr, uint256 unbondingID) constant returns(StakingManagerUnbondingRequest)
func (_StakingManager *StakingManagerSession) GetUnbondingRequest(delegatorAddr common.Address, unbondingID *big.Int) (StakingManagerUnbondingRequest, error) {
	return _StakingManager.Contract.GetUnbondingRequest(&_StakingManager.CallOpts, delegatorAddr, unbondingID)
}

// GetUnbondingRequest is a free data retrieval call binding the contract method 0x9e79b122.
//
// Solidity: function getUnbondingRequest(address delegatorAddr, uint256 unbondingID) constant returns(StakingManagerUnbondingRequest)
func (_StakingManager *StakingManagerCallerSession) GetUnbondingRequest(delegatorAddr common.Address, unbondingID *big.Int) (StakingManagerUnbondingRequest, error) {
	return _StakingManager.Contract.GetUnbondingRequest(&_StakingManager.CallOpts, delegatorAddr, unbondingID)
}

// IsJailed is a free data retrieval call binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address transcoderAddr) constant returns(bool)
func (_StakingManager *StakingManagerCaller) IsJailed(opts *bind.CallOpts, transcoderAddr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "isJailed", transcoderAddr)
	return *ret0, err
}

// IsJailed is a free data retrieval call binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address transcoderAddr) constant returns(bool)
func (_StakingManager *StakingManagerSession) IsJailed(transcoderAddr common.Address) (bool, error) {
	return _StakingManager.Contract.IsJailed(&_StakingManager.CallOpts, transcoderAddr)
}

// IsJailed is a free data retrieval call binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address transcoderAddr) constant returns(bool)
func (_StakingManager *StakingManagerCallerSession) IsJailed(transcoderAddr common.Address) (bool, error) {
	return _StakingManager.Contract.IsJailed(&_StakingManager.CallOpts, transcoderAddr)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_StakingManager *StakingManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_StakingManager *StakingManagerSession) IsOwner() (bool, error) {
	return _StakingManager.Contract.IsOwner(&_StakingManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_StakingManager *StakingManagerCallerSession) IsOwner() (bool, error) {
	return _StakingManager.Contract.IsOwner(&_StakingManager.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() constant returns(uint256)
func (_StakingManager *StakingManagerCaller) MinDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "minDelegation")
	return *ret0, err
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() constant returns(uint256)
func (_StakingManager *StakingManagerSession) MinDelegation() (*big.Int, error) {
	return _StakingManager.Contract.MinDelegation(&_StakingManager.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() constant returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MinDelegation() (*big.Int, error) {
	return _StakingManager.Contract.MinDelegation(&_StakingManager.CallOpts)
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() constant returns(uint256)
func (_StakingManager *StakingManagerCaller) MinSelfStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "minSelfStake")
	return *ret0, err
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() constant returns(uint256)
func (_StakingManager *StakingManagerSession) MinSelfStake() (*big.Int, error) {
	return _StakingManager.Contract.MinSelfStake(&_StakingManager.CallOpts)
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() constant returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MinSelfStake() (*big.Int, error) {
	return _StakingManager.Contract.MinSelfStake(&_StakingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StakingManager *StakingManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StakingManager *StakingManagerSession) Owner() (common.Address, error) {
	return _StakingManager.Contract.Owner(&_StakingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StakingManager *StakingManagerCallerSession) Owner() (common.Address, error) {
	return _StakingManager.Contract.Owner(&_StakingManager.CallOpts)
}

// SlashRate is a free data retrieval call binding the contract method 0xe341181d.
//
// Solidity: function slashRate() constant returns(uint256)
func (_StakingManager *StakingManagerCaller) SlashRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "slashRate")
	return *ret0, err
}

// SlashRate is a free data retrieval call binding the contract method 0xe341181d.
//
// Solidity: function slashRate() constant returns(uint256)
func (_StakingManager *StakingManagerSession) SlashRate() (*big.Int, error) {
	return _StakingManager.Contract.SlashRate(&_StakingManager.CallOpts)
}

// SlashRate is a free data retrieval call binding the contract method 0xe341181d.
//
// Solidity: function slashRate() constant returns(uint256)
func (_StakingManager *StakingManagerCallerSession) SlashRate() (*big.Int, error) {
	return _StakingManager.Contract.SlashRate(&_StakingManager.CallOpts)
}

// TranscoderApprovalPeriod is a free data retrieval call binding the contract method 0x3939e608.
//
// Solidity: function transcoderApprovalPeriod() constant returns(uint256)
func (_StakingManager *StakingManagerCaller) TranscoderApprovalPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "transcoderApprovalPeriod")
	return *ret0, err
}

// TranscoderApprovalPeriod is a free data retrieval call binding the contract method 0x3939e608.
//
// Solidity: function transcoderApprovalPeriod() constant returns(uint256)
func (_StakingManager *StakingManagerSession) TranscoderApprovalPeriod() (*big.Int, error) {
	return _StakingManager.Contract.TranscoderApprovalPeriod(&_StakingManager.CallOpts)
}

// TranscoderApprovalPeriod is a free data retrieval call binding the contract method 0x3939e608.
//
// Solidity: function transcoderApprovalPeriod() constant returns(uint256)
func (_StakingManager *StakingManagerCallerSession) TranscoderApprovalPeriod() (*big.Int, error) {
	return _StakingManager.Contract.TranscoderApprovalPeriod(&_StakingManager.CallOpts)
}

// Transcoders is a free data retrieval call binding the contract method 0xe2dc17f6.
//
// Solidity: function transcoders(address ) constant returns(uint256 total, uint256 timestamp, uint256 rewardRate, uint256 rewards, uint256 zone, uint256 capacity, bool jailed)
func (_StakingManager *StakingManagerCaller) Transcoders(opts *bind.CallOpts, arg0 common.Address) (struct {
	Total      *big.Int
	Timestamp  *big.Int
	RewardRate *big.Int
	Rewards    *big.Int
	Zone       *big.Int
	Capacity   *big.Int
	Jailed     bool
}, error) {
	ret := new(struct {
		Total      *big.Int
		Timestamp  *big.Int
		RewardRate *big.Int
		Rewards    *big.Int
		Zone       *big.Int
		Capacity   *big.Int
		Jailed     bool
	})
	out := ret
	err := _StakingManager.contract.Call(opts, out, "transcoders", arg0)
	return *ret, err
}

// Transcoders is a free data retrieval call binding the contract method 0xe2dc17f6.
//
// Solidity: function transcoders(address ) constant returns(uint256 total, uint256 timestamp, uint256 rewardRate, uint256 rewards, uint256 zone, uint256 capacity, bool jailed)
func (_StakingManager *StakingManagerSession) Transcoders(arg0 common.Address) (struct {
	Total      *big.Int
	Timestamp  *big.Int
	RewardRate *big.Int
	Rewards    *big.Int
	Zone       *big.Int
	Capacity   *big.Int
	Jailed     bool
}, error) {
	return _StakingManager.Contract.Transcoders(&_StakingManager.CallOpts, arg0)
}

// Transcoders is a free data retrieval call binding the contract method 0xe2dc17f6.
//
// Solidity: function transcoders(address ) constant returns(uint256 total, uint256 timestamp, uint256 rewardRate, uint256 rewards, uint256 zone, uint256 capacity, bool jailed)
func (_StakingManager *StakingManagerCallerSession) Transcoders(arg0 common.Address) (struct {
	Total      *big.Int
	Timestamp  *big.Int
	RewardRate *big.Int
	Rewards    *big.Int
	Zone       *big.Int
	Capacity   *big.Int
	Jailed     bool
}, error) {
	return _StakingManager.Contract.Transcoders(&_StakingManager.CallOpts, arg0)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_StakingManager *StakingManagerCaller) UnbondingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "unbondingPeriod")
	return *ret0, err
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_StakingManager *StakingManagerSession) UnbondingPeriod() (*big.Int, error) {
	return _StakingManager.Contract.UnbondingPeriod(&_StakingManager.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() constant returns(uint256)
func (_StakingManager *StakingManagerCallerSession) UnbondingPeriod() (*big.Int, error) {
	return _StakingManager.Contract.UnbondingPeriod(&_StakingManager.CallOpts)
}

// CollectReward is a paid mutator transaction binding the contract method 0xbd3507da.
//
// Solidity: function collectReward(uint256 transcoderAddr) returns()
func (_StakingManager *StakingManagerTransactor) CollectReward(opts *bind.TransactOpts, transcoderAddr *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "collectReward", transcoderAddr)
}

// CollectReward is a paid mutator transaction binding the contract method 0xbd3507da.
//
// Solidity: function collectReward(uint256 transcoderAddr) returns()
func (_StakingManager *StakingManagerSession) CollectReward(transcoderAddr *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.CollectReward(&_StakingManager.TransactOpts, transcoderAddr)
}

// CollectReward is a paid mutator transaction binding the contract method 0xbd3507da.
//
// Solidity: function collectReward(uint256 transcoderAddr) returns()
func (_StakingManager *StakingManagerTransactorSession) CollectReward(transcoderAddr *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.CollectReward(&_StakingManager.TransactOpts, transcoderAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address transcoderAddr) returns()
func (_StakingManager *StakingManagerTransactor) Delegate(opts *bind.TransactOpts, transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "delegate", transcoderAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address transcoderAddr) returns()
func (_StakingManager *StakingManagerSession) Delegate(transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Delegate(&_StakingManager.TransactOpts, transcoderAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address transcoderAddr) returns()
func (_StakingManager *StakingManagerTransactorSession) Delegate(transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Delegate(&_StakingManager.TransactOpts, transcoderAddr)
}

// Rebond is a paid mutator transaction binding the contract method 0x91256c6d.
//
// Solidity: function rebond(uint256 unbondingID, address transcoderAddr) returns(bool)
func (_StakingManager *StakingManagerTransactor) Rebond(opts *bind.TransactOpts, unbondingID *big.Int, transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "rebond", unbondingID, transcoderAddr)
}

// Rebond is a paid mutator transaction binding the contract method 0x91256c6d.
//
// Solidity: function rebond(uint256 unbondingID, address transcoderAddr) returns(bool)
func (_StakingManager *StakingManagerSession) Rebond(unbondingID *big.Int, transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Rebond(&_StakingManager.TransactOpts, unbondingID, transcoderAddr)
}

// Rebond is a paid mutator transaction binding the contract method 0x91256c6d.
//
// Solidity: function rebond(uint256 unbondingID, address transcoderAddr) returns(bool)
func (_StakingManager *StakingManagerTransactorSession) Rebond(unbondingID *big.Int, transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Rebond(&_StakingManager.TransactOpts, unbondingID, transcoderAddr)
}

// RegisterTranscoder is a paid mutator transaction binding the contract method 0x399f57c0.
//
// Solidity: function registerTranscoder(uint256 rate) returns()
func (_StakingManager *StakingManagerTransactor) RegisterTranscoder(opts *bind.TransactOpts, rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "registerTranscoder", rate)
}

// RegisterTranscoder is a paid mutator transaction binding the contract method 0x399f57c0.
//
// Solidity: function registerTranscoder(uint256 rate) returns()
func (_StakingManager *StakingManagerSession) RegisterTranscoder(rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RegisterTranscoder(&_StakingManager.TransactOpts, rate)
}

// RegisterTranscoder is a paid mutator transaction binding the contract method 0x399f57c0.
//
// Solidity: function registerTranscoder(uint256 rate) returns()
func (_StakingManager *StakingManagerTransactorSession) RegisterTranscoder(rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RegisterTranscoder(&_StakingManager.TransactOpts, rate)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceOwnership(&_StakingManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceOwnership(&_StakingManager.TransactOpts)
}

// RequestUnbonding is a paid mutator transaction binding the contract method 0x6db28909.
//
// Solidity: function requestUnbonding(address transcoderAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerTransactor) RequestUnbonding(opts *bind.TransactOpts, transcoderAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "requestUnbonding", transcoderAddr, amount)
}

// RequestUnbonding is a paid mutator transaction binding the contract method 0x6db28909.
//
// Solidity: function requestUnbonding(address transcoderAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerSession) RequestUnbonding(transcoderAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RequestUnbonding(&_StakingManager.TransactOpts, transcoderAddr, amount)
}

// RequestUnbonding is a paid mutator transaction binding the contract method 0x6db28909.
//
// Solidity: function requestUnbonding(address transcoderAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerTransactorSession) RequestUnbonding(transcoderAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RequestUnbonding(&_StakingManager.TransactOpts, transcoderAddr, amount)
}

// SetApprovalPeriod is a paid mutator transaction binding the contract method 0x2c9f0f2e.
//
// Solidity: function setApprovalPeriod(uint256 period) returns()
func (_StakingManager *StakingManagerTransactor) SetApprovalPeriod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setApprovalPeriod", period)
}

// SetApprovalPeriod is a paid mutator transaction binding the contract method 0x2c9f0f2e.
//
// Solidity: function setApprovalPeriod(uint256 period) returns()
func (_StakingManager *StakingManagerSession) SetApprovalPeriod(period *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetApprovalPeriod(&_StakingManager.TransactOpts, period)
}

// SetApprovalPeriod is a paid mutator transaction binding the contract method 0x2c9f0f2e.
//
// Solidity: function setApprovalPeriod(uint256 period) returns()
func (_StakingManager *StakingManagerTransactorSession) SetApprovalPeriod(period *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetApprovalPeriod(&_StakingManager.TransactOpts, period)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x503074ef.
//
// Solidity: function setCapacity(address addr, uint256 capacity) returns()
func (_StakingManager *StakingManagerTransactor) SetCapacity(opts *bind.TransactOpts, addr common.Address, capacity *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setCapacity", addr, capacity)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x503074ef.
//
// Solidity: function setCapacity(address addr, uint256 capacity) returns()
func (_StakingManager *StakingManagerSession) SetCapacity(addr common.Address, capacity *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetCapacity(&_StakingManager.TransactOpts, addr, capacity)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x503074ef.
//
// Solidity: function setCapacity(address addr, uint256 capacity) returns()
func (_StakingManager *StakingManagerTransactorSession) SetCapacity(addr common.Address, capacity *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetCapacity(&_StakingManager.TransactOpts, addr, capacity)
}

// SetSelfMinStake is a paid mutator transaction binding the contract method 0x26e348ba.
//
// Solidity: function setSelfMinStake(uint256 amount) returns()
func (_StakingManager *StakingManagerTransactor) SetSelfMinStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setSelfMinStake", amount)
}

// SetSelfMinStake is a paid mutator transaction binding the contract method 0x26e348ba.
//
// Solidity: function setSelfMinStake(uint256 amount) returns()
func (_StakingManager *StakingManagerSession) SetSelfMinStake(amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSelfMinStake(&_StakingManager.TransactOpts, amount)
}

// SetSelfMinStake is a paid mutator transaction binding the contract method 0x26e348ba.
//
// Solidity: function setSelfMinStake(uint256 amount) returns()
func (_StakingManager *StakingManagerTransactorSession) SetSelfMinStake(amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSelfMinStake(&_StakingManager.TransactOpts, amount)
}

// SetSlashFundAddress is a paid mutator transaction binding the contract method 0xa79e7263.
//
// Solidity: function setSlashFundAddress(address addr) returns()
func (_StakingManager *StakingManagerTransactor) SetSlashFundAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setSlashFundAddress", addr)
}

// SetSlashFundAddress is a paid mutator transaction binding the contract method 0xa79e7263.
//
// Solidity: function setSlashFundAddress(address addr) returns()
func (_StakingManager *StakingManagerSession) SetSlashFundAddress(addr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSlashFundAddress(&_StakingManager.TransactOpts, addr)
}

// SetSlashFundAddress is a paid mutator transaction binding the contract method 0xa79e7263.
//
// Solidity: function setSlashFundAddress(address addr) returns()
func (_StakingManager *StakingManagerTransactorSession) SetSlashFundAddress(addr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSlashFundAddress(&_StakingManager.TransactOpts, addr)
}

// SetZone is a paid mutator transaction binding the contract method 0x96fb4a9c.
//
// Solidity: function setZone(address addr, uint256 zone) returns()
func (_StakingManager *StakingManagerTransactor) SetZone(opts *bind.TransactOpts, addr common.Address, zone *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setZone", addr, zone)
}

// SetZone is a paid mutator transaction binding the contract method 0x96fb4a9c.
//
// Solidity: function setZone(address addr, uint256 zone) returns()
func (_StakingManager *StakingManagerSession) SetZone(addr common.Address, zone *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetZone(&_StakingManager.TransactOpts, addr, zone)
}

// SetZone is a paid mutator transaction binding the contract method 0x96fb4a9c.
//
// Solidity: function setZone(address addr, uint256 zone) returns()
func (_StakingManager *StakingManagerTransactorSession) SetZone(addr common.Address, zone *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetZone(&_StakingManager.TransactOpts, addr, zone)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address addr) returns(bool)
func (_StakingManager *StakingManagerTransactor) Slash(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "slash", addr)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address addr) returns(bool)
func (_StakingManager *StakingManagerSession) Slash(addr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Slash(&_StakingManager.TransactOpts, addr)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address addr) returns(bool)
func (_StakingManager *StakingManagerTransactorSession) Slash(addr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Slash(&_StakingManager.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.TransferOwnership(&_StakingManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.TransferOwnership(&_StakingManager.TransactOpts, newOwner)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address transcoderAddr) returns()
func (_StakingManager *StakingManagerTransactor) Unjail(opts *bind.TransactOpts, transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "unjail", transcoderAddr)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address transcoderAddr) returns()
func (_StakingManager *StakingManagerSession) Unjail(transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Unjail(&_StakingManager.TransactOpts, transcoderAddr)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address transcoderAddr) returns()
func (_StakingManager *StakingManagerTransactorSession) Unjail(transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Unjail(&_StakingManager.TransactOpts, transcoderAddr)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 unbondingID) returns(bool)
func (_StakingManager *StakingManagerTransactor) WithdrawStake(opts *bind.TransactOpts, unbondingID *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "withdrawStake", unbondingID)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 unbondingID) returns(bool)
func (_StakingManager *StakingManagerSession) WithdrawStake(unbondingID *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawStake(&_StakingManager.TransactOpts, unbondingID)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0x25d5971f.
//
// Solidity: function withdrawStake(uint256 unbondingID) returns(bool)
func (_StakingManager *StakingManagerTransactorSession) WithdrawStake(unbondingID *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawStake(&_StakingManager.TransactOpts, unbondingID)
}

// StakingManagerDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the StakingManager contract.
type StakingManagerDelegatedIterator struct {
	Event *StakingManagerDelegated // Event containing the contract specifics and raw log

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
func (it *StakingManagerDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerDelegated)
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
		it.Event = new(StakingManagerDelegated)
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
func (it *StakingManagerDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerDelegated represents a Delegated event raised by the StakingManager contract.
type StakingManagerDelegated struct {
	Transcoder common.Address
	Delegator  common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed transcoder, address indexed delegator, uint256 indexed amount)
func (_StakingManager *StakingManagerFilterer) FilterDelegated(opts *bind.FilterOpts, transcoder []common.Address, delegator []common.Address, amount []*big.Int) (*StakingManagerDelegatedIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Delegated", transcoderRule, delegatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerDelegatedIterator{contract: _StakingManager.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed transcoder, address indexed delegator, uint256 indexed amount)
func (_StakingManager *StakingManagerFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *StakingManagerDelegated, transcoder []common.Address, delegator []common.Address, amount []*big.Int) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Delegated", transcoderRule, delegatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerDelegated)
				if err := _StakingManager.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed transcoder, address indexed delegator, uint256 indexed amount)
func (_StakingManager *StakingManagerFilterer) ParseDelegated(log types.Log) (*StakingManagerDelegated, error) {
	event := new(StakingManagerDelegated)
	if err := _StakingManager.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerJailedIterator is returned from FilterJailed and is used to iterate over the raw logs and unpacked data for Jailed events raised by the StakingManager contract.
type StakingManagerJailedIterator struct {
	Event *StakingManagerJailed // Event containing the contract specifics and raw log

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
func (it *StakingManagerJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerJailed)
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
		it.Event = new(StakingManagerJailed)
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
func (it *StakingManagerJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerJailed represents a Jailed event raised by the StakingManager contract.
type StakingManagerJailed struct {
	Transcoder common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterJailed is a free log retrieval operation binding the contract event 0x519ec2af7b403e5bfa116afc87904cd6aa3e97a09cae81b522551191195674e7.
//
// Solidity: event Jailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) FilterJailed(opts *bind.FilterOpts, transcoder []common.Address) (*StakingManagerJailedIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Jailed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerJailedIterator{contract: _StakingManager.contract, event: "Jailed", logs: logs, sub: sub}, nil
}

// WatchJailed is a free log subscription operation binding the contract event 0x519ec2af7b403e5bfa116afc87904cd6aa3e97a09cae81b522551191195674e7.
//
// Solidity: event Jailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) WatchJailed(opts *bind.WatchOpts, sink chan<- *StakingManagerJailed, transcoder []common.Address) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Jailed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerJailed)
				if err := _StakingManager.contract.UnpackLog(event, "Jailed", log); err != nil {
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

// ParseJailed is a log parse operation binding the contract event 0x519ec2af7b403e5bfa116afc87904cd6aa3e97a09cae81b522551191195674e7.
//
// Solidity: event Jailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) ParseJailed(log types.Log) (*StakingManagerJailed, error) {
	event := new(StakingManagerJailed)
	if err := _StakingManager.contract.UnpackLog(event, "Jailed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingManager contract.
type StakingManagerOwnershipTransferredIterator struct {
	Event *StakingManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerOwnershipTransferred)
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
		it.Event = new(StakingManagerOwnershipTransferred)
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
func (it *StakingManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerOwnershipTransferred represents a OwnershipTransferred event raised by the StakingManager contract.
type StakingManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerOwnershipTransferredIterator{contract: _StakingManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerOwnershipTransferred)
				if err := _StakingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) ParseOwnershipTransferred(log types.Log) (*StakingManagerOwnershipTransferred, error) {
	event := new(StakingManagerOwnershipTransferred)
	if err := _StakingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerRebondedIterator is returned from FilterRebonded and is used to iterate over the raw logs and unpacked data for Rebonded events raised by the StakingManager contract.
type StakingManagerRebondedIterator struct {
	Event *StakingManagerRebonded // Event containing the contract specifics and raw log

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
func (it *StakingManagerRebondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerRebonded)
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
		it.Event = new(StakingManagerRebonded)
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
func (it *StakingManagerRebondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerRebondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerRebonded represents a Rebonded event raised by the StakingManager contract.
type StakingManagerRebonded struct {
	Delegator  common.Address
	Transcoder common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRebonded is a free log retrieval operation binding the contract event 0x870b8cd1f69783c8a776e39ac19c4601d05a03e14275ccd7356b9620f328747d.
//
// Solidity: event Rebonded(address indexed delegator, address indexed transcoder, uint256 indexed amount)
func (_StakingManager *StakingManagerFilterer) FilterRebonded(opts *bind.FilterOpts, delegator []common.Address, transcoder []common.Address, amount []*big.Int) (*StakingManagerRebondedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Rebonded", delegatorRule, transcoderRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerRebondedIterator{contract: _StakingManager.contract, event: "Rebonded", logs: logs, sub: sub}, nil
}

// WatchRebonded is a free log subscription operation binding the contract event 0x870b8cd1f69783c8a776e39ac19c4601d05a03e14275ccd7356b9620f328747d.
//
// Solidity: event Rebonded(address indexed delegator, address indexed transcoder, uint256 indexed amount)
func (_StakingManager *StakingManagerFilterer) WatchRebonded(opts *bind.WatchOpts, sink chan<- *StakingManagerRebonded, delegator []common.Address, transcoder []common.Address, amount []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Rebonded", delegatorRule, transcoderRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerRebonded)
				if err := _StakingManager.contract.UnpackLog(event, "Rebonded", log); err != nil {
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

// ParseRebonded is a log parse operation binding the contract event 0x870b8cd1f69783c8a776e39ac19c4601d05a03e14275ccd7356b9620f328747d.
//
// Solidity: event Rebonded(address indexed delegator, address indexed transcoder, uint256 indexed amount)
func (_StakingManager *StakingManagerFilterer) ParseRebonded(log types.Log) (*StakingManagerRebonded, error) {
	event := new(StakingManagerRebonded)
	if err := _StakingManager.contract.UnpackLog(event, "Rebonded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the StakingManager contract.
type StakingManagerSlashedIterator struct {
	Event *StakingManagerSlashed // Event containing the contract specifics and raw log

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
func (it *StakingManagerSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerSlashed)
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
		it.Event = new(StakingManagerSlashed)
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
func (it *StakingManagerSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerSlashed represents a Slashed event raised by the StakingManager contract.
type StakingManagerSlashed struct {
	Transcoder common.Address
	Rate       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed transcoder, uint256 indexed rate)
func (_StakingManager *StakingManagerFilterer) FilterSlashed(opts *bind.FilterOpts, transcoder []common.Address, rate []*big.Int) (*StakingManagerSlashedIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}
	var rateRule []interface{}
	for _, rateItem := range rate {
		rateRule = append(rateRule, rateItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Slashed", transcoderRule, rateRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerSlashedIterator{contract: _StakingManager.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed transcoder, uint256 indexed rate)
func (_StakingManager *StakingManagerFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *StakingManagerSlashed, transcoder []common.Address, rate []*big.Int) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}
	var rateRule []interface{}
	for _, rateItem := range rate {
		rateRule = append(rateRule, rateItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Slashed", transcoderRule, rateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerSlashed)
				if err := _StakingManager.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed transcoder, uint256 indexed rate)
func (_StakingManager *StakingManagerFilterer) ParseSlashed(log types.Log) (*StakingManagerSlashed, error) {
	event := new(StakingManagerSlashed)
	if err := _StakingManager.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerStakeWithdrawalIterator is returned from FilterStakeWithdrawal and is used to iterate over the raw logs and unpacked data for StakeWithdrawal events raised by the StakingManager contract.
type StakingManagerStakeWithdrawalIterator struct {
	Event *StakingManagerStakeWithdrawal // Event containing the contract specifics and raw log

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
func (it *StakingManagerStakeWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerStakeWithdrawal)
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
		it.Event = new(StakingManagerStakeWithdrawal)
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
func (it *StakingManagerStakeWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerStakeWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerStakeWithdrawal represents a StakeWithdrawal event raised by the StakingManager contract.
type StakingManagerStakeWithdrawal struct {
	Delegator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawal is a free log retrieval operation binding the contract event 0xf0ed97f7b968f9d8268bc8d104a11b3586ceeadd0e0af5f73769e2b479f9d0ae.
//
// Solidity: event StakeWithdrawal(address indexed delegator, uint256 indexed amount)
func (_StakingManager *StakingManagerFilterer) FilterStakeWithdrawal(opts *bind.FilterOpts, delegator []common.Address, amount []*big.Int) (*StakingManagerStakeWithdrawalIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "StakeWithdrawal", delegatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerStakeWithdrawalIterator{contract: _StakingManager.contract, event: "StakeWithdrawal", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawal is a free log subscription operation binding the contract event 0xf0ed97f7b968f9d8268bc8d104a11b3586ceeadd0e0af5f73769e2b479f9d0ae.
//
// Solidity: event StakeWithdrawal(address indexed delegator, uint256 indexed amount)
func (_StakingManager *StakingManagerFilterer) WatchStakeWithdrawal(opts *bind.WatchOpts, sink chan<- *StakingManagerStakeWithdrawal, delegator []common.Address, amount []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "StakeWithdrawal", delegatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerStakeWithdrawal)
				if err := _StakingManager.contract.UnpackLog(event, "StakeWithdrawal", log); err != nil {
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

// ParseStakeWithdrawal is a log parse operation binding the contract event 0xf0ed97f7b968f9d8268bc8d104a11b3586ceeadd0e0af5f73769e2b479f9d0ae.
//
// Solidity: event StakeWithdrawal(address indexed delegator, uint256 indexed amount)
func (_StakingManager *StakingManagerFilterer) ParseStakeWithdrawal(log types.Log) (*StakingManagerStakeWithdrawal, error) {
	event := new(StakingManagerStakeWithdrawal)
	if err := _StakingManager.contract.UnpackLog(event, "StakeWithdrawal", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerTranscoderRegisteredIterator is returned from FilterTranscoderRegistered and is used to iterate over the raw logs and unpacked data for TranscoderRegistered events raised by the StakingManager contract.
type StakingManagerTranscoderRegisteredIterator struct {
	Event *StakingManagerTranscoderRegistered // Event containing the contract specifics and raw log

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
func (it *StakingManagerTranscoderRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerTranscoderRegistered)
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
		it.Event = new(StakingManagerTranscoderRegistered)
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
func (it *StakingManagerTranscoderRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerTranscoderRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerTranscoderRegistered represents a TranscoderRegistered event raised by the StakingManager contract.
type StakingManagerTranscoderRegistered struct {
	Transcoder common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTranscoderRegistered is a free log retrieval operation binding the contract event 0x6fbcf0f12b438f90175bebf725f86a4a74d12525d5d2c144a68e400696bce58b.
//
// Solidity: event TranscoderRegistered(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) FilterTranscoderRegistered(opts *bind.FilterOpts, transcoder []common.Address) (*StakingManagerTranscoderRegisteredIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "TranscoderRegistered", transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerTranscoderRegisteredIterator{contract: _StakingManager.contract, event: "TranscoderRegistered", logs: logs, sub: sub}, nil
}

// WatchTranscoderRegistered is a free log subscription operation binding the contract event 0x6fbcf0f12b438f90175bebf725f86a4a74d12525d5d2c144a68e400696bce58b.
//
// Solidity: event TranscoderRegistered(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) WatchTranscoderRegistered(opts *bind.WatchOpts, sink chan<- *StakingManagerTranscoderRegistered, transcoder []common.Address) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "TranscoderRegistered", transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerTranscoderRegistered)
				if err := _StakingManager.contract.UnpackLog(event, "TranscoderRegistered", log); err != nil {
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

// ParseTranscoderRegistered is a log parse operation binding the contract event 0x6fbcf0f12b438f90175bebf725f86a4a74d12525d5d2c144a68e400696bce58b.
//
// Solidity: event TranscoderRegistered(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) ParseTranscoderRegistered(log types.Log) (*StakingManagerTranscoderRegistered, error) {
	event := new(StakingManagerTranscoderRegistered)
	if err := _StakingManager.contract.UnpackLog(event, "TranscoderRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerUnbondingRequestedIterator is returned from FilterUnbondingRequested and is used to iterate over the raw logs and unpacked data for UnbondingRequested events raised by the StakingManager contract.
type StakingManagerUnbondingRequestedIterator struct {
	Event *StakingManagerUnbondingRequested // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnbondingRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnbondingRequested)
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
		it.Event = new(StakingManagerUnbondingRequested)
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
func (it *StakingManagerUnbondingRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnbondingRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnbondingRequested represents a UnbondingRequested event raised by the StakingManager contract.
type StakingManagerUnbondingRequested struct {
	UnbondingID *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnbondingRequested is a free log retrieval operation binding the contract event 0xcf7436a5022e8b877fe3f9116467e1ee036494fe62586e9c00395e79c6dcbdbb.
//
// Solidity: event UnbondingRequested(uint256 indexed unbondingID)
func (_StakingManager *StakingManagerFilterer) FilterUnbondingRequested(opts *bind.FilterOpts, unbondingID []*big.Int) (*StakingManagerUnbondingRequestedIterator, error) {

	var unbondingIDRule []interface{}
	for _, unbondingIDItem := range unbondingID {
		unbondingIDRule = append(unbondingIDRule, unbondingIDItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "UnbondingRequested", unbondingIDRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnbondingRequestedIterator{contract: _StakingManager.contract, event: "UnbondingRequested", logs: logs, sub: sub}, nil
}

// WatchUnbondingRequested is a free log subscription operation binding the contract event 0xcf7436a5022e8b877fe3f9116467e1ee036494fe62586e9c00395e79c6dcbdbb.
//
// Solidity: event UnbondingRequested(uint256 indexed unbondingID)
func (_StakingManager *StakingManagerFilterer) WatchUnbondingRequested(opts *bind.WatchOpts, sink chan<- *StakingManagerUnbondingRequested, unbondingID []*big.Int) (event.Subscription, error) {

	var unbondingIDRule []interface{}
	for _, unbondingIDItem := range unbondingID {
		unbondingIDRule = append(unbondingIDRule, unbondingIDItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "UnbondingRequested", unbondingIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnbondingRequested)
				if err := _StakingManager.contract.UnpackLog(event, "UnbondingRequested", log); err != nil {
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

// ParseUnbondingRequested is a log parse operation binding the contract event 0xcf7436a5022e8b877fe3f9116467e1ee036494fe62586e9c00395e79c6dcbdbb.
//
// Solidity: event UnbondingRequested(uint256 indexed unbondingID)
func (_StakingManager *StakingManagerFilterer) ParseUnbondingRequested(log types.Log) (*StakingManagerUnbondingRequested, error) {
	event := new(StakingManagerUnbondingRequested)
	if err := _StakingManager.contract.UnpackLog(event, "UnbondingRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerUnjailedIterator is returned from FilterUnjailed and is used to iterate over the raw logs and unpacked data for Unjailed events raised by the StakingManager contract.
type StakingManagerUnjailedIterator struct {
	Event *StakingManagerUnjailed // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnjailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnjailed)
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
		it.Event = new(StakingManagerUnjailed)
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
func (it *StakingManagerUnjailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnjailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnjailed represents a Unjailed event raised by the StakingManager contract.
type StakingManagerUnjailed struct {
	Transcoder common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnjailed is a free log retrieval operation binding the contract event 0xfa5039497ad9ba11f0eb5239b2614e925541bbcc0cf3476dd68e1927c86d33ff.
//
// Solidity: event Unjailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) FilterUnjailed(opts *bind.FilterOpts, transcoder []common.Address) (*StakingManagerUnjailedIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Unjailed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnjailedIterator{contract: _StakingManager.contract, event: "Unjailed", logs: logs, sub: sub}, nil
}

// WatchUnjailed is a free log subscription operation binding the contract event 0xfa5039497ad9ba11f0eb5239b2614e925541bbcc0cf3476dd68e1927c86d33ff.
//
// Solidity: event Unjailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) WatchUnjailed(opts *bind.WatchOpts, sink chan<- *StakingManagerUnjailed, transcoder []common.Address) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Unjailed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnjailed)
				if err := _StakingManager.contract.UnpackLog(event, "Unjailed", log); err != nil {
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

// ParseUnjailed is a log parse operation binding the contract event 0xfa5039497ad9ba11f0eb5239b2614e925541bbcc0cf3476dd68e1927c86d33ff.
//
// Solidity: event Unjailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) ParseUnjailed(log types.Log) (*StakingManagerUnjailed, error) {
	event := new(StakingManagerUnjailed)
	if err := _StakingManager.contract.UnpackLog(event, "Unjailed", log); err != nil {
		return nil, err
	}
	return event, nil
}
