// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave_lending_eth

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AaveLendingEthMetaData contains all meta data concerning the AaveLendingEth contract.
var AaveLendingEthMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interesRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"borrowETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rateMode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"repayETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"permitV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"permitR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"permitS\",\"type\":\"bytes32\"}],\"name\":\"withdrawETHWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveLendingEthABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveLendingEthMetaData.ABI instead.
var AaveLendingEthABI = AaveLendingEthMetaData.ABI

// AaveLendingEth is an auto generated Go binding around an Ethereum contract.
type AaveLendingEth struct {
	AaveLendingEthCaller     // Read-only binding to the contract
	AaveLendingEthTransactor // Write-only binding to the contract
	AaveLendingEthFilterer   // Log filterer for contract events
}

// AaveLendingEthCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveLendingEthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveLendingEthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveLendingEthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveLendingEthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveLendingEthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveLendingEthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveLendingEthSession struct {
	Contract     *AaveLendingEth   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveLendingEthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveLendingEthCallerSession struct {
	Contract *AaveLendingEthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AaveLendingEthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveLendingEthTransactorSession struct {
	Contract     *AaveLendingEthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AaveLendingEthRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveLendingEthRaw struct {
	Contract *AaveLendingEth // Generic contract binding to access the raw methods on
}

// AaveLendingEthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveLendingEthCallerRaw struct {
	Contract *AaveLendingEthCaller // Generic read-only contract binding to access the raw methods on
}

// AaveLendingEthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveLendingEthTransactorRaw struct {
	Contract *AaveLendingEthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveLendingEth creates a new instance of AaveLendingEth, bound to a specific deployed contract.
func NewAaveLendingEth(address common.Address, backend bind.ContractBackend) (*AaveLendingEth, error) {
	contract, err := bindAaveLendingEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveLendingEth{AaveLendingEthCaller: AaveLendingEthCaller{contract: contract}, AaveLendingEthTransactor: AaveLendingEthTransactor{contract: contract}, AaveLendingEthFilterer: AaveLendingEthFilterer{contract: contract}}, nil
}

// NewAaveLendingEthCaller creates a new read-only instance of AaveLendingEth, bound to a specific deployed contract.
func NewAaveLendingEthCaller(address common.Address, caller bind.ContractCaller) (*AaveLendingEthCaller, error) {
	contract, err := bindAaveLendingEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveLendingEthCaller{contract: contract}, nil
}

// NewAaveLendingEthTransactor creates a new write-only instance of AaveLendingEth, bound to a specific deployed contract.
func NewAaveLendingEthTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveLendingEthTransactor, error) {
	contract, err := bindAaveLendingEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveLendingEthTransactor{contract: contract}, nil
}

// NewAaveLendingEthFilterer creates a new log filterer instance of AaveLendingEth, bound to a specific deployed contract.
func NewAaveLendingEthFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveLendingEthFilterer, error) {
	contract, err := bindAaveLendingEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveLendingEthFilterer{contract: contract}, nil
}

// bindAaveLendingEth binds a generic wrapper to an already deployed contract.
func bindAaveLendingEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveLendingEthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveLendingEth *AaveLendingEthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveLendingEth.Contract.AaveLendingEthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveLendingEth *AaveLendingEthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveLendingEth.Contract.AaveLendingEthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveLendingEth *AaveLendingEthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveLendingEth.Contract.AaveLendingEthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveLendingEth *AaveLendingEthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveLendingEth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveLendingEth *AaveLendingEthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveLendingEth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveLendingEth *AaveLendingEthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveLendingEth.Contract.contract.Transact(opts, method, params...)
}

// BorrowETH is a paid mutator transaction binding the contract method 0x66514c97.
//
// Solidity: function borrowETH(address pool, uint256 amount, uint256 interesRateMode, uint16 referralCode) returns()
func (_AaveLendingEth *AaveLendingEthTransactor) BorrowETH(opts *bind.TransactOpts, pool common.Address, amount *big.Int, interesRateMode *big.Int, referralCode uint16) (*types.Transaction, error) {
	return _AaveLendingEth.contract.Transact(opts, "borrowETH", pool, amount, interesRateMode, referralCode)
}

// BorrowETH is a paid mutator transaction binding the contract method 0x66514c97.
//
// Solidity: function borrowETH(address pool, uint256 amount, uint256 interesRateMode, uint16 referralCode) returns()
func (_AaveLendingEth *AaveLendingEthSession) BorrowETH(pool common.Address, amount *big.Int, interesRateMode *big.Int, referralCode uint16) (*types.Transaction, error) {
	return _AaveLendingEth.Contract.BorrowETH(&_AaveLendingEth.TransactOpts, pool, amount, interesRateMode, referralCode)
}

// BorrowETH is a paid mutator transaction binding the contract method 0x66514c97.
//
// Solidity: function borrowETH(address pool, uint256 amount, uint256 interesRateMode, uint16 referralCode) returns()
func (_AaveLendingEth *AaveLendingEthTransactorSession) BorrowETH(pool common.Address, amount *big.Int, interesRateMode *big.Int, referralCode uint16) (*types.Transaction, error) {
	return _AaveLendingEth.Contract.BorrowETH(&_AaveLendingEth.TransactOpts, pool, amount, interesRateMode, referralCode)
}

// DepositETH is a paid mutator transaction binding the contract method 0x474cf53d.
//
// Solidity: function depositETH(address pool, address onBehalfOf, uint16 referralCode) payable returns()
func (_AaveLendingEth *AaveLendingEthTransactor) DepositETH(opts *bind.TransactOpts, pool common.Address, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveLendingEth.contract.Transact(opts, "depositETH", pool, onBehalfOf, referralCode)
}

// DepositETH is a paid mutator transaction binding the contract method 0x474cf53d.
//
// Solidity: function depositETH(address pool, address onBehalfOf, uint16 referralCode) payable returns()
func (_AaveLendingEth *AaveLendingEthSession) DepositETH(pool common.Address, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveLendingEth.Contract.DepositETH(&_AaveLendingEth.TransactOpts, pool, onBehalfOf, referralCode)
}

// DepositETH is a paid mutator transaction binding the contract method 0x474cf53d.
//
// Solidity: function depositETH(address pool, address onBehalfOf, uint16 referralCode) payable returns()
func (_AaveLendingEth *AaveLendingEthTransactorSession) DepositETH(pool common.Address, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveLendingEth.Contract.DepositETH(&_AaveLendingEth.TransactOpts, pool, onBehalfOf, referralCode)
}

// RepayETH is a paid mutator transaction binding the contract method 0x02c5fcf8.
//
// Solidity: function repayETH(address pool, uint256 amount, uint256 rateMode, address onBehalfOf) payable returns()
func (_AaveLendingEth *AaveLendingEthTransactor) RepayETH(opts *bind.TransactOpts, pool common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveLendingEth.contract.Transact(opts, "repayETH", pool, amount, rateMode, onBehalfOf)
}

// RepayETH is a paid mutator transaction binding the contract method 0x02c5fcf8.
//
// Solidity: function repayETH(address pool, uint256 amount, uint256 rateMode, address onBehalfOf) payable returns()
func (_AaveLendingEth *AaveLendingEthSession) RepayETH(pool common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveLendingEth.Contract.RepayETH(&_AaveLendingEth.TransactOpts, pool, amount, rateMode, onBehalfOf)
}

// RepayETH is a paid mutator transaction binding the contract method 0x02c5fcf8.
//
// Solidity: function repayETH(address pool, uint256 amount, uint256 rateMode, address onBehalfOf) payable returns()
func (_AaveLendingEth *AaveLendingEthTransactorSession) RepayETH(pool common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveLendingEth.Contract.RepayETH(&_AaveLendingEth.TransactOpts, pool, amount, rateMode, onBehalfOf)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x80500d20.
//
// Solidity: function withdrawETH(address pool, uint256 amount, address onBehalfOf) returns()
func (_AaveLendingEth *AaveLendingEthTransactor) WithdrawETH(opts *bind.TransactOpts, pool common.Address, amount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveLendingEth.contract.Transact(opts, "withdrawETH", pool, amount, onBehalfOf)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x80500d20.
//
// Solidity: function withdrawETH(address pool, uint256 amount, address onBehalfOf) returns()
func (_AaveLendingEth *AaveLendingEthSession) WithdrawETH(pool common.Address, amount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveLendingEth.Contract.WithdrawETH(&_AaveLendingEth.TransactOpts, pool, amount, onBehalfOf)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x80500d20.
//
// Solidity: function withdrawETH(address pool, uint256 amount, address onBehalfOf) returns()
func (_AaveLendingEth *AaveLendingEthTransactorSession) WithdrawETH(pool common.Address, amount *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveLendingEth.Contract.WithdrawETH(&_AaveLendingEth.TransactOpts, pool, amount, onBehalfOf)
}

// WithdrawETHWithPermit is a paid mutator transaction binding the contract method 0xd4c40b6c.
//
// Solidity: function withdrawETHWithPermit(address pool, uint256 amount, address to, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_AaveLendingEth *AaveLendingEthTransactor) WithdrawETHWithPermit(opts *bind.TransactOpts, pool common.Address, amount *big.Int, to common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AaveLendingEth.contract.Transact(opts, "withdrawETHWithPermit", pool, amount, to, deadline, permitV, permitR, permitS)
}

// WithdrawETHWithPermit is a paid mutator transaction binding the contract method 0xd4c40b6c.
//
// Solidity: function withdrawETHWithPermit(address pool, uint256 amount, address to, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_AaveLendingEth *AaveLendingEthSession) WithdrawETHWithPermit(pool common.Address, amount *big.Int, to common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AaveLendingEth.Contract.WithdrawETHWithPermit(&_AaveLendingEth.TransactOpts, pool, amount, to, deadline, permitV, permitR, permitS)
}

// WithdrawETHWithPermit is a paid mutator transaction binding the contract method 0xd4c40b6c.
//
// Solidity: function withdrawETHWithPermit(address pool, uint256 amount, address to, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_AaveLendingEth *AaveLendingEthTransactorSession) WithdrawETHWithPermit(pool common.Address, amount *big.Int, to common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AaveLendingEth.Contract.WithdrawETHWithPermit(&_AaveLendingEth.TransactOpts, pool, amount, to, deadline, permitV, permitR, permitS)
}
