// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pool_address_provider

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

// PoolAddressProviderMetaData contains all meta data concerning the PoolAddressProvider contract.
var PoolAddressProviderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ACLAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ACLManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"AddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldImplementationAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementationAddress\",\"type\":\"address\"}],\"name\":\"AddressSetAsProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"oldMarketId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"newMarketId\",\"type\":\"string\"}],\"name\":\"MarketIdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PoolConfiguratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PoolDataProviderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PriceOracleSentinelUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PriceOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementationAddress\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getACLAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getACLManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolConfigurator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolDataProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracleSentinel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAclAdmin\",\"type\":\"address\"}],\"name\":\"setACLAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAclManager\",\"type\":\"address\"}],\"name\":\"setACLManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newImplementationAddress\",\"type\":\"address\"}],\"name\":\"setAddressAsProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newMarketId\",\"type\":\"string\"}],\"name\":\"setMarketId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPoolConfiguratorImpl\",\"type\":\"address\"}],\"name\":\"setPoolConfiguratorImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDataProvider\",\"type\":\"address\"}],\"name\":\"setPoolDataProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPoolImpl\",\"type\":\"address\"}],\"name\":\"setPoolImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPriceOracleSentinel\",\"type\":\"address\"}],\"name\":\"setPriceOracleSentinel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PoolAddressProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolAddressProviderMetaData.ABI instead.
var PoolAddressProviderABI = PoolAddressProviderMetaData.ABI

// PoolAddressProvider is an auto generated Go binding around an Ethereum contract.
type PoolAddressProvider struct {
	PoolAddressProviderCaller     // Read-only binding to the contract
	PoolAddressProviderTransactor // Write-only binding to the contract
	PoolAddressProviderFilterer   // Log filterer for contract events
}

// PoolAddressProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolAddressProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolAddressProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolAddressProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolAddressProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolAddressProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolAddressProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolAddressProviderSession struct {
	Contract     *PoolAddressProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PoolAddressProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolAddressProviderCallerSession struct {
	Contract *PoolAddressProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PoolAddressProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolAddressProviderTransactorSession struct {
	Contract     *PoolAddressProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PoolAddressProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolAddressProviderRaw struct {
	Contract *PoolAddressProvider // Generic contract binding to access the raw methods on
}

// PoolAddressProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolAddressProviderCallerRaw struct {
	Contract *PoolAddressProviderCaller // Generic read-only contract binding to access the raw methods on
}

// PoolAddressProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolAddressProviderTransactorRaw struct {
	Contract *PoolAddressProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolAddressProvider creates a new instance of PoolAddressProvider, bound to a specific deployed contract.
func NewPoolAddressProvider(address common.Address, backend bind.ContractBackend) (*PoolAddressProvider, error) {
	contract, err := bindPoolAddressProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolAddressProvider{PoolAddressProviderCaller: PoolAddressProviderCaller{contract: contract}, PoolAddressProviderTransactor: PoolAddressProviderTransactor{contract: contract}, PoolAddressProviderFilterer: PoolAddressProviderFilterer{contract: contract}}, nil
}

// NewPoolAddressProviderCaller creates a new read-only instance of PoolAddressProvider, bound to a specific deployed contract.
func NewPoolAddressProviderCaller(address common.Address, caller bind.ContractCaller) (*PoolAddressProviderCaller, error) {
	contract, err := bindPoolAddressProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolAddressProviderCaller{contract: contract}, nil
}

// NewPoolAddressProviderTransactor creates a new write-only instance of PoolAddressProvider, bound to a specific deployed contract.
func NewPoolAddressProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolAddressProviderTransactor, error) {
	contract, err := bindPoolAddressProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolAddressProviderTransactor{contract: contract}, nil
}

// NewPoolAddressProviderFilterer creates a new log filterer instance of PoolAddressProvider, bound to a specific deployed contract.
func NewPoolAddressProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolAddressProviderFilterer, error) {
	contract, err := bindPoolAddressProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolAddressProviderFilterer{contract: contract}, nil
}

// bindPoolAddressProvider binds a generic wrapper to an already deployed contract.
func bindPoolAddressProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolAddressProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolAddressProvider *PoolAddressProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolAddressProvider.Contract.PoolAddressProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolAddressProvider *PoolAddressProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.PoolAddressProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolAddressProvider *PoolAddressProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.PoolAddressProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolAddressProvider *PoolAddressProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolAddressProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolAddressProvider *PoolAddressProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolAddressProvider *PoolAddressProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.contract.Transact(opts, method, params...)
}

// GetACLAdmin is a free data retrieval call binding the contract method 0x0e67178c.
//
// Solidity: function getACLAdmin() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderCaller) GetACLAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolAddressProvider.contract.Call(opts, &out, "getACLAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetACLAdmin is a free data retrieval call binding the contract method 0x0e67178c.
//
// Solidity: function getACLAdmin() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderSession) GetACLAdmin() (common.Address, error) {
	return _PoolAddressProvider.Contract.GetACLAdmin(&_PoolAddressProvider.CallOpts)
}

// GetACLAdmin is a free data retrieval call binding the contract method 0x0e67178c.
//
// Solidity: function getACLAdmin() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderCallerSession) GetACLAdmin() (common.Address, error) {
	return _PoolAddressProvider.Contract.GetACLAdmin(&_PoolAddressProvider.CallOpts)
}

// GetACLManager is a free data retrieval call binding the contract method 0x707cd716.
//
// Solidity: function getACLManager() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderCaller) GetACLManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolAddressProvider.contract.Call(opts, &out, "getACLManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetACLManager is a free data retrieval call binding the contract method 0x707cd716.
//
// Solidity: function getACLManager() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderSession) GetACLManager() (common.Address, error) {
	return _PoolAddressProvider.Contract.GetACLManager(&_PoolAddressProvider.CallOpts)
}

// GetACLManager is a free data retrieval call binding the contract method 0x707cd716.
//
// Solidity: function getACLManager() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderCallerSession) GetACLManager() (common.Address, error) {
	return _PoolAddressProvider.Contract.GetACLManager(&_PoolAddressProvider.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_PoolAddressProvider *PoolAddressProviderCaller) GetAddress(opts *bind.CallOpts, id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _PoolAddressProvider.contract.Call(opts, &out, "getAddress", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_PoolAddressProvider *PoolAddressProviderSession) GetAddress(id [32]byte) (common.Address, error) {
	return _PoolAddressProvider.Contract.GetAddress(&_PoolAddressProvider.CallOpts, id)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_PoolAddressProvider *PoolAddressProviderCallerSession) GetAddress(id [32]byte) (common.Address, error) {
	return _PoolAddressProvider.Contract.GetAddress(&_PoolAddressProvider.CallOpts, id)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_PoolAddressProvider *PoolAddressProviderCaller) GetMarketId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PoolAddressProvider.contract.Call(opts, &out, "getMarketId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_PoolAddressProvider *PoolAddressProviderSession) GetMarketId() (string, error) {
	return _PoolAddressProvider.Contract.GetMarketId(&_PoolAddressProvider.CallOpts)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_PoolAddressProvider *PoolAddressProviderCallerSession) GetMarketId() (string, error) {
	return _PoolAddressProvider.Contract.GetMarketId(&_PoolAddressProvider.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0x026b1d5f.
//
// Solidity: function getPool() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderCaller) GetPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolAddressProvider.contract.Call(opts, &out, "getPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x026b1d5f.
//
// Solidity: function getPool() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderSession) GetPool() (common.Address, error) {
	return _PoolAddressProvider.Contract.GetPool(&_PoolAddressProvider.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0x026b1d5f.
//
// Solidity: function getPool() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderCallerSession) GetPool() (common.Address, error) {
	return _PoolAddressProvider.Contract.GetPool(&_PoolAddressProvider.CallOpts)
}

// GetPoolConfigurator is a free data retrieval call binding the contract method 0x631adfca.
//
// Solidity: function getPoolConfigurator() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderCaller) GetPoolConfigurator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolAddressProvider.contract.Call(opts, &out, "getPoolConfigurator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolConfigurator is a free data retrieval call binding the contract method 0x631adfca.
//
// Solidity: function getPoolConfigurator() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderSession) GetPoolConfigurator() (common.Address, error) {
	return _PoolAddressProvider.Contract.GetPoolConfigurator(&_PoolAddressProvider.CallOpts)
}

// GetPoolConfigurator is a free data retrieval call binding the contract method 0x631adfca.
//
// Solidity: function getPoolConfigurator() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderCallerSession) GetPoolConfigurator() (common.Address, error) {
	return _PoolAddressProvider.Contract.GetPoolConfigurator(&_PoolAddressProvider.CallOpts)
}

// GetPoolDataProvider is a free data retrieval call binding the contract method 0xe860accb.
//
// Solidity: function getPoolDataProvider() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderCaller) GetPoolDataProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolAddressProvider.contract.Call(opts, &out, "getPoolDataProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolDataProvider is a free data retrieval call binding the contract method 0xe860accb.
//
// Solidity: function getPoolDataProvider() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderSession) GetPoolDataProvider() (common.Address, error) {
	return _PoolAddressProvider.Contract.GetPoolDataProvider(&_PoolAddressProvider.CallOpts)
}

// GetPoolDataProvider is a free data retrieval call binding the contract method 0xe860accb.
//
// Solidity: function getPoolDataProvider() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderCallerSession) GetPoolDataProvider() (common.Address, error) {
	return _PoolAddressProvider.Contract.GetPoolDataProvider(&_PoolAddressProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderCaller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolAddressProvider.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderSession) GetPriceOracle() (common.Address, error) {
	return _PoolAddressProvider.Contract.GetPriceOracle(&_PoolAddressProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderCallerSession) GetPriceOracle() (common.Address, error) {
	return _PoolAddressProvider.Contract.GetPriceOracle(&_PoolAddressProvider.CallOpts)
}

// GetPriceOracleSentinel is a free data retrieval call binding the contract method 0x5eb88d3d.
//
// Solidity: function getPriceOracleSentinel() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderCaller) GetPriceOracleSentinel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolAddressProvider.contract.Call(opts, &out, "getPriceOracleSentinel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracleSentinel is a free data retrieval call binding the contract method 0x5eb88d3d.
//
// Solidity: function getPriceOracleSentinel() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderSession) GetPriceOracleSentinel() (common.Address, error) {
	return _PoolAddressProvider.Contract.GetPriceOracleSentinel(&_PoolAddressProvider.CallOpts)
}

// GetPriceOracleSentinel is a free data retrieval call binding the contract method 0x5eb88d3d.
//
// Solidity: function getPriceOracleSentinel() view returns(address)
func (_PoolAddressProvider *PoolAddressProviderCallerSession) GetPriceOracleSentinel() (common.Address, error) {
	return _PoolAddressProvider.Contract.GetPriceOracleSentinel(&_PoolAddressProvider.CallOpts)
}

// SetACLAdmin is a paid mutator transaction binding the contract method 0x76d84ffc.
//
// Solidity: function setACLAdmin(address newAclAdmin) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactor) SetACLAdmin(opts *bind.TransactOpts, newAclAdmin common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.contract.Transact(opts, "setACLAdmin", newAclAdmin)
}

// SetACLAdmin is a paid mutator transaction binding the contract method 0x76d84ffc.
//
// Solidity: function setACLAdmin(address newAclAdmin) returns()
func (_PoolAddressProvider *PoolAddressProviderSession) SetACLAdmin(newAclAdmin common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetACLAdmin(&_PoolAddressProvider.TransactOpts, newAclAdmin)
}

// SetACLAdmin is a paid mutator transaction binding the contract method 0x76d84ffc.
//
// Solidity: function setACLAdmin(address newAclAdmin) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactorSession) SetACLAdmin(newAclAdmin common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetACLAdmin(&_PoolAddressProvider.TransactOpts, newAclAdmin)
}

// SetACLManager is a paid mutator transaction binding the contract method 0xed301ca9.
//
// Solidity: function setACLManager(address newAclManager) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactor) SetACLManager(opts *bind.TransactOpts, newAclManager common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.contract.Transact(opts, "setACLManager", newAclManager)
}

// SetACLManager is a paid mutator transaction binding the contract method 0xed301ca9.
//
// Solidity: function setACLManager(address newAclManager) returns()
func (_PoolAddressProvider *PoolAddressProviderSession) SetACLManager(newAclManager common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetACLManager(&_PoolAddressProvider.TransactOpts, newAclManager)
}

// SetACLManager is a paid mutator transaction binding the contract method 0xed301ca9.
//
// Solidity: function setACLManager(address newAclManager) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactorSession) SetACLManager(newAclManager common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetACLManager(&_PoolAddressProvider.TransactOpts, newAclManager)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactor) SetAddress(opts *bind.TransactOpts, id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.contract.Transact(opts, "setAddress", id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_PoolAddressProvider *PoolAddressProviderSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetAddress(&_PoolAddressProvider.TransactOpts, id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactorSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetAddress(&_PoolAddressProvider.TransactOpts, id, newAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address newImplementationAddress) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactor) SetAddressAsProxy(opts *bind.TransactOpts, id [32]byte, newImplementationAddress common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.contract.Transact(opts, "setAddressAsProxy", id, newImplementationAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address newImplementationAddress) returns()
func (_PoolAddressProvider *PoolAddressProviderSession) SetAddressAsProxy(id [32]byte, newImplementationAddress common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetAddressAsProxy(&_PoolAddressProvider.TransactOpts, id, newImplementationAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address newImplementationAddress) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactorSession) SetAddressAsProxy(id [32]byte, newImplementationAddress common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetAddressAsProxy(&_PoolAddressProvider.TransactOpts, id, newImplementationAddress)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string newMarketId) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactor) SetMarketId(opts *bind.TransactOpts, newMarketId string) (*types.Transaction, error) {
	return _PoolAddressProvider.contract.Transact(opts, "setMarketId", newMarketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string newMarketId) returns()
func (_PoolAddressProvider *PoolAddressProviderSession) SetMarketId(newMarketId string) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetMarketId(&_PoolAddressProvider.TransactOpts, newMarketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string newMarketId) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactorSession) SetMarketId(newMarketId string) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetMarketId(&_PoolAddressProvider.TransactOpts, newMarketId)
}

// SetPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xe4ca28b7.
//
// Solidity: function setPoolConfiguratorImpl(address newPoolConfiguratorImpl) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactor) SetPoolConfiguratorImpl(opts *bind.TransactOpts, newPoolConfiguratorImpl common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.contract.Transact(opts, "setPoolConfiguratorImpl", newPoolConfiguratorImpl)
}

// SetPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xe4ca28b7.
//
// Solidity: function setPoolConfiguratorImpl(address newPoolConfiguratorImpl) returns()
func (_PoolAddressProvider *PoolAddressProviderSession) SetPoolConfiguratorImpl(newPoolConfiguratorImpl common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetPoolConfiguratorImpl(&_PoolAddressProvider.TransactOpts, newPoolConfiguratorImpl)
}

// SetPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xe4ca28b7.
//
// Solidity: function setPoolConfiguratorImpl(address newPoolConfiguratorImpl) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactorSession) SetPoolConfiguratorImpl(newPoolConfiguratorImpl common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetPoolConfiguratorImpl(&_PoolAddressProvider.TransactOpts, newPoolConfiguratorImpl)
}

// SetPoolDataProvider is a paid mutator transaction binding the contract method 0xe44e9ed1.
//
// Solidity: function setPoolDataProvider(address newDataProvider) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactor) SetPoolDataProvider(opts *bind.TransactOpts, newDataProvider common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.contract.Transact(opts, "setPoolDataProvider", newDataProvider)
}

// SetPoolDataProvider is a paid mutator transaction binding the contract method 0xe44e9ed1.
//
// Solidity: function setPoolDataProvider(address newDataProvider) returns()
func (_PoolAddressProvider *PoolAddressProviderSession) SetPoolDataProvider(newDataProvider common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetPoolDataProvider(&_PoolAddressProvider.TransactOpts, newDataProvider)
}

// SetPoolDataProvider is a paid mutator transaction binding the contract method 0xe44e9ed1.
//
// Solidity: function setPoolDataProvider(address newDataProvider) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactorSession) SetPoolDataProvider(newDataProvider common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetPoolDataProvider(&_PoolAddressProvider.TransactOpts, newDataProvider)
}

// SetPoolImpl is a paid mutator transaction binding the contract method 0xa1564406.
//
// Solidity: function setPoolImpl(address newPoolImpl) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactor) SetPoolImpl(opts *bind.TransactOpts, newPoolImpl common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.contract.Transact(opts, "setPoolImpl", newPoolImpl)
}

// SetPoolImpl is a paid mutator transaction binding the contract method 0xa1564406.
//
// Solidity: function setPoolImpl(address newPoolImpl) returns()
func (_PoolAddressProvider *PoolAddressProviderSession) SetPoolImpl(newPoolImpl common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetPoolImpl(&_PoolAddressProvider.TransactOpts, newPoolImpl)
}

// SetPoolImpl is a paid mutator transaction binding the contract method 0xa1564406.
//
// Solidity: function setPoolImpl(address newPoolImpl) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactorSession) SetPoolImpl(newPoolImpl common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetPoolImpl(&_PoolAddressProvider.TransactOpts, newPoolImpl)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceOracle) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactor) SetPriceOracle(opts *bind.TransactOpts, newPriceOracle common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.contract.Transact(opts, "setPriceOracle", newPriceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceOracle) returns()
func (_PoolAddressProvider *PoolAddressProviderSession) SetPriceOracle(newPriceOracle common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetPriceOracle(&_PoolAddressProvider.TransactOpts, newPriceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceOracle) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactorSession) SetPriceOracle(newPriceOracle common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetPriceOracle(&_PoolAddressProvider.TransactOpts, newPriceOracle)
}

// SetPriceOracleSentinel is a paid mutator transaction binding the contract method 0x74944cec.
//
// Solidity: function setPriceOracleSentinel(address newPriceOracleSentinel) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactor) SetPriceOracleSentinel(opts *bind.TransactOpts, newPriceOracleSentinel common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.contract.Transact(opts, "setPriceOracleSentinel", newPriceOracleSentinel)
}

// SetPriceOracleSentinel is a paid mutator transaction binding the contract method 0x74944cec.
//
// Solidity: function setPriceOracleSentinel(address newPriceOracleSentinel) returns()
func (_PoolAddressProvider *PoolAddressProviderSession) SetPriceOracleSentinel(newPriceOracleSentinel common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetPriceOracleSentinel(&_PoolAddressProvider.TransactOpts, newPriceOracleSentinel)
}

// SetPriceOracleSentinel is a paid mutator transaction binding the contract method 0x74944cec.
//
// Solidity: function setPriceOracleSentinel(address newPriceOracleSentinel) returns()
func (_PoolAddressProvider *PoolAddressProviderTransactorSession) SetPriceOracleSentinel(newPriceOracleSentinel common.Address) (*types.Transaction, error) {
	return _PoolAddressProvider.Contract.SetPriceOracleSentinel(&_PoolAddressProvider.TransactOpts, newPriceOracleSentinel)
}

// PoolAddressProviderACLAdminUpdatedIterator is returned from FilterACLAdminUpdated and is used to iterate over the raw logs and unpacked data for ACLAdminUpdated events raised by the PoolAddressProvider contract.
type PoolAddressProviderACLAdminUpdatedIterator struct {
	Event *PoolAddressProviderACLAdminUpdated // Event containing the contract specifics and raw log

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
func (it *PoolAddressProviderACLAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAddressProviderACLAdminUpdated)
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
		it.Event = new(PoolAddressProviderACLAdminUpdated)
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
func (it *PoolAddressProviderACLAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAddressProviderACLAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAddressProviderACLAdminUpdated represents a ACLAdminUpdated event raised by the PoolAddressProvider contract.
type PoolAddressProviderACLAdminUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterACLAdminUpdated is a free log retrieval operation binding the contract event 0xe9cf53972264dc95304fd424458745019ddfca0e37ae8f703d74772c41ad115b.
//
// Solidity: event ACLAdminUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) FilterACLAdminUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*PoolAddressProviderACLAdminUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.FilterLogs(opts, "ACLAdminUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &PoolAddressProviderACLAdminUpdatedIterator{contract: _PoolAddressProvider.contract, event: "ACLAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchACLAdminUpdated is a free log subscription operation binding the contract event 0xe9cf53972264dc95304fd424458745019ddfca0e37ae8f703d74772c41ad115b.
//
// Solidity: event ACLAdminUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) WatchACLAdminUpdated(opts *bind.WatchOpts, sink chan<- *PoolAddressProviderACLAdminUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.WatchLogs(opts, "ACLAdminUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAddressProviderACLAdminUpdated)
				if err := _PoolAddressProvider.contract.UnpackLog(event, "ACLAdminUpdated", log); err != nil {
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

// ParseACLAdminUpdated is a log parse operation binding the contract event 0xe9cf53972264dc95304fd424458745019ddfca0e37ae8f703d74772c41ad115b.
//
// Solidity: event ACLAdminUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) ParseACLAdminUpdated(log types.Log) (*PoolAddressProviderACLAdminUpdated, error) {
	event := new(PoolAddressProviderACLAdminUpdated)
	if err := _PoolAddressProvider.contract.UnpackLog(event, "ACLAdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAddressProviderACLManagerUpdatedIterator is returned from FilterACLManagerUpdated and is used to iterate over the raw logs and unpacked data for ACLManagerUpdated events raised by the PoolAddressProvider contract.
type PoolAddressProviderACLManagerUpdatedIterator struct {
	Event *PoolAddressProviderACLManagerUpdated // Event containing the contract specifics and raw log

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
func (it *PoolAddressProviderACLManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAddressProviderACLManagerUpdated)
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
		it.Event = new(PoolAddressProviderACLManagerUpdated)
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
func (it *PoolAddressProviderACLManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAddressProviderACLManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAddressProviderACLManagerUpdated represents a ACLManagerUpdated event raised by the PoolAddressProvider contract.
type PoolAddressProviderACLManagerUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterACLManagerUpdated is a free log retrieval operation binding the contract event 0xb30efa04327bb8a537d61cc1e5c48095345ad18ef7cc04e6bacf7dfb6caaf507.
//
// Solidity: event ACLManagerUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) FilterACLManagerUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*PoolAddressProviderACLManagerUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.FilterLogs(opts, "ACLManagerUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &PoolAddressProviderACLManagerUpdatedIterator{contract: _PoolAddressProvider.contract, event: "ACLManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchACLManagerUpdated is a free log subscription operation binding the contract event 0xb30efa04327bb8a537d61cc1e5c48095345ad18ef7cc04e6bacf7dfb6caaf507.
//
// Solidity: event ACLManagerUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) WatchACLManagerUpdated(opts *bind.WatchOpts, sink chan<- *PoolAddressProviderACLManagerUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.WatchLogs(opts, "ACLManagerUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAddressProviderACLManagerUpdated)
				if err := _PoolAddressProvider.contract.UnpackLog(event, "ACLManagerUpdated", log); err != nil {
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

// ParseACLManagerUpdated is a log parse operation binding the contract event 0xb30efa04327bb8a537d61cc1e5c48095345ad18ef7cc04e6bacf7dfb6caaf507.
//
// Solidity: event ACLManagerUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) ParseACLManagerUpdated(log types.Log) (*PoolAddressProviderACLManagerUpdated, error) {
	event := new(PoolAddressProviderACLManagerUpdated)
	if err := _PoolAddressProvider.contract.UnpackLog(event, "ACLManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAddressProviderAddressSetIterator is returned from FilterAddressSet and is used to iterate over the raw logs and unpacked data for AddressSet events raised by the PoolAddressProvider contract.
type PoolAddressProviderAddressSetIterator struct {
	Event *PoolAddressProviderAddressSet // Event containing the contract specifics and raw log

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
func (it *PoolAddressProviderAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAddressProviderAddressSet)
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
		it.Event = new(PoolAddressProviderAddressSet)
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
func (it *PoolAddressProviderAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAddressProviderAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAddressProviderAddressSet represents a AddressSet event raised by the PoolAddressProvider contract.
type PoolAddressProviderAddressSet struct {
	Id         [32]byte
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressSet is a free log retrieval operation binding the contract event 0x9ef0e8c8e52743bb38b83b17d9429141d494b8041ca6d616a6c77cebae9cd8b7.
//
// Solidity: event AddressSet(bytes32 indexed id, address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) FilterAddressSet(opts *bind.FilterOpts, id [][32]byte, oldAddress []common.Address, newAddress []common.Address) (*PoolAddressProviderAddressSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.FilterLogs(opts, "AddressSet", idRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &PoolAddressProviderAddressSetIterator{contract: _PoolAddressProvider.contract, event: "AddressSet", logs: logs, sub: sub}, nil
}

// WatchAddressSet is a free log subscription operation binding the contract event 0x9ef0e8c8e52743bb38b83b17d9429141d494b8041ca6d616a6c77cebae9cd8b7.
//
// Solidity: event AddressSet(bytes32 indexed id, address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) WatchAddressSet(opts *bind.WatchOpts, sink chan<- *PoolAddressProviderAddressSet, id [][32]byte, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.WatchLogs(opts, "AddressSet", idRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAddressProviderAddressSet)
				if err := _PoolAddressProvider.contract.UnpackLog(event, "AddressSet", log); err != nil {
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

// ParseAddressSet is a log parse operation binding the contract event 0x9ef0e8c8e52743bb38b83b17d9429141d494b8041ca6d616a6c77cebae9cd8b7.
//
// Solidity: event AddressSet(bytes32 indexed id, address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) ParseAddressSet(log types.Log) (*PoolAddressProviderAddressSet, error) {
	event := new(PoolAddressProviderAddressSet)
	if err := _PoolAddressProvider.contract.UnpackLog(event, "AddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAddressProviderAddressSetAsProxyIterator is returned from FilterAddressSetAsProxy and is used to iterate over the raw logs and unpacked data for AddressSetAsProxy events raised by the PoolAddressProvider contract.
type PoolAddressProviderAddressSetAsProxyIterator struct {
	Event *PoolAddressProviderAddressSetAsProxy // Event containing the contract specifics and raw log

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
func (it *PoolAddressProviderAddressSetAsProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAddressProviderAddressSetAsProxy)
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
		it.Event = new(PoolAddressProviderAddressSetAsProxy)
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
func (it *PoolAddressProviderAddressSetAsProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAddressProviderAddressSetAsProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAddressProviderAddressSetAsProxy represents a AddressSetAsProxy event raised by the PoolAddressProvider contract.
type PoolAddressProviderAddressSetAsProxy struct {
	Id                       [32]byte
	ProxyAddress             common.Address
	OldImplementationAddress common.Address
	NewImplementationAddress common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterAddressSetAsProxy is a free log retrieval operation binding the contract event 0x3bbd45b5429b385e3fb37ad5cd1cd1435a3c8ec32196c7937597365a3fd3e99c.
//
// Solidity: event AddressSetAsProxy(bytes32 indexed id, address indexed proxyAddress, address oldImplementationAddress, address indexed newImplementationAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) FilterAddressSetAsProxy(opts *bind.FilterOpts, id [][32]byte, proxyAddress []common.Address, newImplementationAddress []common.Address) (*PoolAddressProviderAddressSetAsProxyIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proxyAddressRule []interface{}
	for _, proxyAddressItem := range proxyAddress {
		proxyAddressRule = append(proxyAddressRule, proxyAddressItem)
	}

	var newImplementationAddressRule []interface{}
	for _, newImplementationAddressItem := range newImplementationAddress {
		newImplementationAddressRule = append(newImplementationAddressRule, newImplementationAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.FilterLogs(opts, "AddressSetAsProxy", idRule, proxyAddressRule, newImplementationAddressRule)
	if err != nil {
		return nil, err
	}
	return &PoolAddressProviderAddressSetAsProxyIterator{contract: _PoolAddressProvider.contract, event: "AddressSetAsProxy", logs: logs, sub: sub}, nil
}

// WatchAddressSetAsProxy is a free log subscription operation binding the contract event 0x3bbd45b5429b385e3fb37ad5cd1cd1435a3c8ec32196c7937597365a3fd3e99c.
//
// Solidity: event AddressSetAsProxy(bytes32 indexed id, address indexed proxyAddress, address oldImplementationAddress, address indexed newImplementationAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) WatchAddressSetAsProxy(opts *bind.WatchOpts, sink chan<- *PoolAddressProviderAddressSetAsProxy, id [][32]byte, proxyAddress []common.Address, newImplementationAddress []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proxyAddressRule []interface{}
	for _, proxyAddressItem := range proxyAddress {
		proxyAddressRule = append(proxyAddressRule, proxyAddressItem)
	}

	var newImplementationAddressRule []interface{}
	for _, newImplementationAddressItem := range newImplementationAddress {
		newImplementationAddressRule = append(newImplementationAddressRule, newImplementationAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.WatchLogs(opts, "AddressSetAsProxy", idRule, proxyAddressRule, newImplementationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAddressProviderAddressSetAsProxy)
				if err := _PoolAddressProvider.contract.UnpackLog(event, "AddressSetAsProxy", log); err != nil {
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

// ParseAddressSetAsProxy is a log parse operation binding the contract event 0x3bbd45b5429b385e3fb37ad5cd1cd1435a3c8ec32196c7937597365a3fd3e99c.
//
// Solidity: event AddressSetAsProxy(bytes32 indexed id, address indexed proxyAddress, address oldImplementationAddress, address indexed newImplementationAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) ParseAddressSetAsProxy(log types.Log) (*PoolAddressProviderAddressSetAsProxy, error) {
	event := new(PoolAddressProviderAddressSetAsProxy)
	if err := _PoolAddressProvider.contract.UnpackLog(event, "AddressSetAsProxy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAddressProviderMarketIdSetIterator is returned from FilterMarketIdSet and is used to iterate over the raw logs and unpacked data for MarketIdSet events raised by the PoolAddressProvider contract.
type PoolAddressProviderMarketIdSetIterator struct {
	Event *PoolAddressProviderMarketIdSet // Event containing the contract specifics and raw log

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
func (it *PoolAddressProviderMarketIdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAddressProviderMarketIdSet)
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
		it.Event = new(PoolAddressProviderMarketIdSet)
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
func (it *PoolAddressProviderMarketIdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAddressProviderMarketIdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAddressProviderMarketIdSet represents a MarketIdSet event raised by the PoolAddressProvider contract.
type PoolAddressProviderMarketIdSet struct {
	OldMarketId common.Hash
	NewMarketId common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketIdSet is a free log retrieval operation binding the contract event 0xe685c8cdecc6030c45030fd54778812cb84ed8e4467c38294403d68ba7860823.
//
// Solidity: event MarketIdSet(string indexed oldMarketId, string indexed newMarketId)
func (_PoolAddressProvider *PoolAddressProviderFilterer) FilterMarketIdSet(opts *bind.FilterOpts, oldMarketId []string, newMarketId []string) (*PoolAddressProviderMarketIdSetIterator, error) {

	var oldMarketIdRule []interface{}
	for _, oldMarketIdItem := range oldMarketId {
		oldMarketIdRule = append(oldMarketIdRule, oldMarketIdItem)
	}
	var newMarketIdRule []interface{}
	for _, newMarketIdItem := range newMarketId {
		newMarketIdRule = append(newMarketIdRule, newMarketIdItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.FilterLogs(opts, "MarketIdSet", oldMarketIdRule, newMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &PoolAddressProviderMarketIdSetIterator{contract: _PoolAddressProvider.contract, event: "MarketIdSet", logs: logs, sub: sub}, nil
}

// WatchMarketIdSet is a free log subscription operation binding the contract event 0xe685c8cdecc6030c45030fd54778812cb84ed8e4467c38294403d68ba7860823.
//
// Solidity: event MarketIdSet(string indexed oldMarketId, string indexed newMarketId)
func (_PoolAddressProvider *PoolAddressProviderFilterer) WatchMarketIdSet(opts *bind.WatchOpts, sink chan<- *PoolAddressProviderMarketIdSet, oldMarketId []string, newMarketId []string) (event.Subscription, error) {

	var oldMarketIdRule []interface{}
	for _, oldMarketIdItem := range oldMarketId {
		oldMarketIdRule = append(oldMarketIdRule, oldMarketIdItem)
	}
	var newMarketIdRule []interface{}
	for _, newMarketIdItem := range newMarketId {
		newMarketIdRule = append(newMarketIdRule, newMarketIdItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.WatchLogs(opts, "MarketIdSet", oldMarketIdRule, newMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAddressProviderMarketIdSet)
				if err := _PoolAddressProvider.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
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

// ParseMarketIdSet is a log parse operation binding the contract event 0xe685c8cdecc6030c45030fd54778812cb84ed8e4467c38294403d68ba7860823.
//
// Solidity: event MarketIdSet(string indexed oldMarketId, string indexed newMarketId)
func (_PoolAddressProvider *PoolAddressProviderFilterer) ParseMarketIdSet(log types.Log) (*PoolAddressProviderMarketIdSet, error) {
	event := new(PoolAddressProviderMarketIdSet)
	if err := _PoolAddressProvider.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAddressProviderPoolConfiguratorUpdatedIterator is returned from FilterPoolConfiguratorUpdated and is used to iterate over the raw logs and unpacked data for PoolConfiguratorUpdated events raised by the PoolAddressProvider contract.
type PoolAddressProviderPoolConfiguratorUpdatedIterator struct {
	Event *PoolAddressProviderPoolConfiguratorUpdated // Event containing the contract specifics and raw log

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
func (it *PoolAddressProviderPoolConfiguratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAddressProviderPoolConfiguratorUpdated)
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
		it.Event = new(PoolAddressProviderPoolConfiguratorUpdated)
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
func (it *PoolAddressProviderPoolConfiguratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAddressProviderPoolConfiguratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAddressProviderPoolConfiguratorUpdated represents a PoolConfiguratorUpdated event raised by the PoolAddressProvider contract.
type PoolAddressProviderPoolConfiguratorUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolConfiguratorUpdated is a free log retrieval operation binding the contract event 0x8932892569eba59c8382a089d9b732d1f49272878775235761a2a6b0309cd465.
//
// Solidity: event PoolConfiguratorUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) FilterPoolConfiguratorUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*PoolAddressProviderPoolConfiguratorUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.FilterLogs(opts, "PoolConfiguratorUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &PoolAddressProviderPoolConfiguratorUpdatedIterator{contract: _PoolAddressProvider.contract, event: "PoolConfiguratorUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolConfiguratorUpdated is a free log subscription operation binding the contract event 0x8932892569eba59c8382a089d9b732d1f49272878775235761a2a6b0309cd465.
//
// Solidity: event PoolConfiguratorUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) WatchPoolConfiguratorUpdated(opts *bind.WatchOpts, sink chan<- *PoolAddressProviderPoolConfiguratorUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.WatchLogs(opts, "PoolConfiguratorUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAddressProviderPoolConfiguratorUpdated)
				if err := _PoolAddressProvider.contract.UnpackLog(event, "PoolConfiguratorUpdated", log); err != nil {
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

// ParsePoolConfiguratorUpdated is a log parse operation binding the contract event 0x8932892569eba59c8382a089d9b732d1f49272878775235761a2a6b0309cd465.
//
// Solidity: event PoolConfiguratorUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) ParsePoolConfiguratorUpdated(log types.Log) (*PoolAddressProviderPoolConfiguratorUpdated, error) {
	event := new(PoolAddressProviderPoolConfiguratorUpdated)
	if err := _PoolAddressProvider.contract.UnpackLog(event, "PoolConfiguratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAddressProviderPoolDataProviderUpdatedIterator is returned from FilterPoolDataProviderUpdated and is used to iterate over the raw logs and unpacked data for PoolDataProviderUpdated events raised by the PoolAddressProvider contract.
type PoolAddressProviderPoolDataProviderUpdatedIterator struct {
	Event *PoolAddressProviderPoolDataProviderUpdated // Event containing the contract specifics and raw log

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
func (it *PoolAddressProviderPoolDataProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAddressProviderPoolDataProviderUpdated)
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
		it.Event = new(PoolAddressProviderPoolDataProviderUpdated)
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
func (it *PoolAddressProviderPoolDataProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAddressProviderPoolDataProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAddressProviderPoolDataProviderUpdated represents a PoolDataProviderUpdated event raised by the PoolAddressProvider contract.
type PoolAddressProviderPoolDataProviderUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolDataProviderUpdated is a free log retrieval operation binding the contract event 0xc853974cfbf81487a14a23565917bee63f527853bcb5fa54f2ae1cdf8a38356d.
//
// Solidity: event PoolDataProviderUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) FilterPoolDataProviderUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*PoolAddressProviderPoolDataProviderUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.FilterLogs(opts, "PoolDataProviderUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &PoolAddressProviderPoolDataProviderUpdatedIterator{contract: _PoolAddressProvider.contract, event: "PoolDataProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolDataProviderUpdated is a free log subscription operation binding the contract event 0xc853974cfbf81487a14a23565917bee63f527853bcb5fa54f2ae1cdf8a38356d.
//
// Solidity: event PoolDataProviderUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) WatchPoolDataProviderUpdated(opts *bind.WatchOpts, sink chan<- *PoolAddressProviderPoolDataProviderUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.WatchLogs(opts, "PoolDataProviderUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAddressProviderPoolDataProviderUpdated)
				if err := _PoolAddressProvider.contract.UnpackLog(event, "PoolDataProviderUpdated", log); err != nil {
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

// ParsePoolDataProviderUpdated is a log parse operation binding the contract event 0xc853974cfbf81487a14a23565917bee63f527853bcb5fa54f2ae1cdf8a38356d.
//
// Solidity: event PoolDataProviderUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) ParsePoolDataProviderUpdated(log types.Log) (*PoolAddressProviderPoolDataProviderUpdated, error) {
	event := new(PoolAddressProviderPoolDataProviderUpdated)
	if err := _PoolAddressProvider.contract.UnpackLog(event, "PoolDataProviderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAddressProviderPoolUpdatedIterator is returned from FilterPoolUpdated and is used to iterate over the raw logs and unpacked data for PoolUpdated events raised by the PoolAddressProvider contract.
type PoolAddressProviderPoolUpdatedIterator struct {
	Event *PoolAddressProviderPoolUpdated // Event containing the contract specifics and raw log

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
func (it *PoolAddressProviderPoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAddressProviderPoolUpdated)
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
		it.Event = new(PoolAddressProviderPoolUpdated)
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
func (it *PoolAddressProviderPoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAddressProviderPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAddressProviderPoolUpdated represents a PoolUpdated event raised by the PoolAddressProvider contract.
type PoolAddressProviderPoolUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolUpdated is a free log retrieval operation binding the contract event 0x90affc163f1a2dfedcd36aa02ed992eeeba8100a4014f0b4cdc20ea265a66627.
//
// Solidity: event PoolUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) FilterPoolUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*PoolAddressProviderPoolUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.FilterLogs(opts, "PoolUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &PoolAddressProviderPoolUpdatedIterator{contract: _PoolAddressProvider.contract, event: "PoolUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolUpdated is a free log subscription operation binding the contract event 0x90affc163f1a2dfedcd36aa02ed992eeeba8100a4014f0b4cdc20ea265a66627.
//
// Solidity: event PoolUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) WatchPoolUpdated(opts *bind.WatchOpts, sink chan<- *PoolAddressProviderPoolUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.WatchLogs(opts, "PoolUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAddressProviderPoolUpdated)
				if err := _PoolAddressProvider.contract.UnpackLog(event, "PoolUpdated", log); err != nil {
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

// ParsePoolUpdated is a log parse operation binding the contract event 0x90affc163f1a2dfedcd36aa02ed992eeeba8100a4014f0b4cdc20ea265a66627.
//
// Solidity: event PoolUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) ParsePoolUpdated(log types.Log) (*PoolAddressProviderPoolUpdated, error) {
	event := new(PoolAddressProviderPoolUpdated)
	if err := _PoolAddressProvider.contract.UnpackLog(event, "PoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAddressProviderPriceOracleSentinelUpdatedIterator is returned from FilterPriceOracleSentinelUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleSentinelUpdated events raised by the PoolAddressProvider contract.
type PoolAddressProviderPriceOracleSentinelUpdatedIterator struct {
	Event *PoolAddressProviderPriceOracleSentinelUpdated // Event containing the contract specifics and raw log

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
func (it *PoolAddressProviderPriceOracleSentinelUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAddressProviderPriceOracleSentinelUpdated)
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
		it.Event = new(PoolAddressProviderPriceOracleSentinelUpdated)
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
func (it *PoolAddressProviderPriceOracleSentinelUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAddressProviderPriceOracleSentinelUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAddressProviderPriceOracleSentinelUpdated represents a PriceOracleSentinelUpdated event raised by the PoolAddressProvider contract.
type PoolAddressProviderPriceOracleSentinelUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleSentinelUpdated is a free log retrieval operation binding the contract event 0x5326514eeca90494a14bedabcff812a0e683029ee85d1e23824d44fd14cd6ae7.
//
// Solidity: event PriceOracleSentinelUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) FilterPriceOracleSentinelUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*PoolAddressProviderPriceOracleSentinelUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.FilterLogs(opts, "PriceOracleSentinelUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &PoolAddressProviderPriceOracleSentinelUpdatedIterator{contract: _PoolAddressProvider.contract, event: "PriceOracleSentinelUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleSentinelUpdated is a free log subscription operation binding the contract event 0x5326514eeca90494a14bedabcff812a0e683029ee85d1e23824d44fd14cd6ae7.
//
// Solidity: event PriceOracleSentinelUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) WatchPriceOracleSentinelUpdated(opts *bind.WatchOpts, sink chan<- *PoolAddressProviderPriceOracleSentinelUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.WatchLogs(opts, "PriceOracleSentinelUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAddressProviderPriceOracleSentinelUpdated)
				if err := _PoolAddressProvider.contract.UnpackLog(event, "PriceOracleSentinelUpdated", log); err != nil {
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

// ParsePriceOracleSentinelUpdated is a log parse operation binding the contract event 0x5326514eeca90494a14bedabcff812a0e683029ee85d1e23824d44fd14cd6ae7.
//
// Solidity: event PriceOracleSentinelUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) ParsePriceOracleSentinelUpdated(log types.Log) (*PoolAddressProviderPriceOracleSentinelUpdated, error) {
	event := new(PoolAddressProviderPriceOracleSentinelUpdated)
	if err := _PoolAddressProvider.contract.UnpackLog(event, "PriceOracleSentinelUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAddressProviderPriceOracleUpdatedIterator is returned from FilterPriceOracleUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleUpdated events raised by the PoolAddressProvider contract.
type PoolAddressProviderPriceOracleUpdatedIterator struct {
	Event *PoolAddressProviderPriceOracleUpdated // Event containing the contract specifics and raw log

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
func (it *PoolAddressProviderPriceOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAddressProviderPriceOracleUpdated)
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
		it.Event = new(PoolAddressProviderPriceOracleUpdated)
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
func (it *PoolAddressProviderPriceOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAddressProviderPriceOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAddressProviderPriceOracleUpdated represents a PriceOracleUpdated event raised by the PoolAddressProvider contract.
type PoolAddressProviderPriceOracleUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleUpdated is a free log retrieval operation binding the contract event 0x56b5f80d8cac1479698aa7d01605fd6111e90b15fc4d2b377417f46034876cbd.
//
// Solidity: event PriceOracleUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) FilterPriceOracleUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*PoolAddressProviderPriceOracleUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.FilterLogs(opts, "PriceOracleUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &PoolAddressProviderPriceOracleUpdatedIterator{contract: _PoolAddressProvider.contract, event: "PriceOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleUpdated is a free log subscription operation binding the contract event 0x56b5f80d8cac1479698aa7d01605fd6111e90b15fc4d2b377417f46034876cbd.
//
// Solidity: event PriceOracleUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) WatchPriceOracleUpdated(opts *bind.WatchOpts, sink chan<- *PoolAddressProviderPriceOracleUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.WatchLogs(opts, "PriceOracleUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAddressProviderPriceOracleUpdated)
				if err := _PoolAddressProvider.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
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

// ParsePriceOracleUpdated is a log parse operation binding the contract event 0x56b5f80d8cac1479698aa7d01605fd6111e90b15fc4d2b377417f46034876cbd.
//
// Solidity: event PriceOracleUpdated(address indexed oldAddress, address indexed newAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) ParsePriceOracleUpdated(log types.Log) (*PoolAddressProviderPriceOracleUpdated, error) {
	event := new(PoolAddressProviderPriceOracleUpdated)
	if err := _PoolAddressProvider.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAddressProviderProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the PoolAddressProvider contract.
type PoolAddressProviderProxyCreatedIterator struct {
	Event *PoolAddressProviderProxyCreated // Event containing the contract specifics and raw log

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
func (it *PoolAddressProviderProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAddressProviderProxyCreated)
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
		it.Event = new(PoolAddressProviderProxyCreated)
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
func (it *PoolAddressProviderProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAddressProviderProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAddressProviderProxyCreated represents a ProxyCreated event raised by the PoolAddressProvider contract.
type PoolAddressProviderProxyCreated struct {
	Id                    [32]byte
	ProxyAddress          common.Address
	ImplementationAddress common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x4a465a9bd819d9662563c1e11ae958f8109e437e7f4bf1c6ef0b9a7b3f35d478.
//
// Solidity: event ProxyCreated(bytes32 indexed id, address indexed proxyAddress, address indexed implementationAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) FilterProxyCreated(opts *bind.FilterOpts, id [][32]byte, proxyAddress []common.Address, implementationAddress []common.Address) (*PoolAddressProviderProxyCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proxyAddressRule []interface{}
	for _, proxyAddressItem := range proxyAddress {
		proxyAddressRule = append(proxyAddressRule, proxyAddressItem)
	}
	var implementationAddressRule []interface{}
	for _, implementationAddressItem := range implementationAddress {
		implementationAddressRule = append(implementationAddressRule, implementationAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.FilterLogs(opts, "ProxyCreated", idRule, proxyAddressRule, implementationAddressRule)
	if err != nil {
		return nil, err
	}
	return &PoolAddressProviderProxyCreatedIterator{contract: _PoolAddressProvider.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x4a465a9bd819d9662563c1e11ae958f8109e437e7f4bf1c6ef0b9a7b3f35d478.
//
// Solidity: event ProxyCreated(bytes32 indexed id, address indexed proxyAddress, address indexed implementationAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *PoolAddressProviderProxyCreated, id [][32]byte, proxyAddress []common.Address, implementationAddress []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proxyAddressRule []interface{}
	for _, proxyAddressItem := range proxyAddress {
		proxyAddressRule = append(proxyAddressRule, proxyAddressItem)
	}
	var implementationAddressRule []interface{}
	for _, implementationAddressItem := range implementationAddress {
		implementationAddressRule = append(implementationAddressRule, implementationAddressItem)
	}

	logs, sub, err := _PoolAddressProvider.contract.WatchLogs(opts, "ProxyCreated", idRule, proxyAddressRule, implementationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAddressProviderProxyCreated)
				if err := _PoolAddressProvider.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
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

// ParseProxyCreated is a log parse operation binding the contract event 0x4a465a9bd819d9662563c1e11ae958f8109e437e7f4bf1c6ef0b9a7b3f35d478.
//
// Solidity: event ProxyCreated(bytes32 indexed id, address indexed proxyAddress, address indexed implementationAddress)
func (_PoolAddressProvider *PoolAddressProviderFilterer) ParseProxyCreated(log types.Log) (*PoolAddressProviderProxyCreated, error) {
	event := new(PoolAddressProviderProxyCreated)
	if err := _PoolAddressProvider.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
