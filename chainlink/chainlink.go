// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package link

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

// LinkMetaData contains all meta data concerning the Link contract.
var LinkMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LinkABI is the input ABI used to generate the binding from.
// Deprecated: Use LinkMetaData.ABI instead.
var LinkABI = LinkMetaData.ABI

// Link is an auto generated Go binding around an Ethereum contract.
type Link struct {
	LinkCaller     // Read-only binding to the contract
	LinkTransactor // Write-only binding to the contract
	LinkFilterer   // Log filterer for contract events
}

// LinkCaller is an auto generated read-only Go binding around an Ethereum contract.
type LinkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LinkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LinkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LinkSession struct {
	Contract     *Link             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LinkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LinkCallerSession struct {
	Contract *LinkCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LinkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LinkTransactorSession struct {
	Contract     *LinkTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LinkRaw is an auto generated low-level Go binding around an Ethereum contract.
type LinkRaw struct {
	Contract *Link // Generic contract binding to access the raw methods on
}

// LinkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LinkCallerRaw struct {
	Contract *LinkCaller // Generic read-only contract binding to access the raw methods on
}

// LinkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LinkTransactorRaw struct {
	Contract *LinkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLink creates a new instance of Link, bound to a specific deployed contract.
func NewLink(address common.Address, backend bind.ContractBackend) (*Link, error) {
	contract, err := bindLink(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Link{LinkCaller: LinkCaller{contract: contract}, LinkTransactor: LinkTransactor{contract: contract}, LinkFilterer: LinkFilterer{contract: contract}}, nil
}

// NewLinkCaller creates a new read-only instance of Link, bound to a specific deployed contract.
func NewLinkCaller(address common.Address, caller bind.ContractCaller) (*LinkCaller, error) {
	contract, err := bindLink(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LinkCaller{contract: contract}, nil
}

// NewLinkTransactor creates a new write-only instance of Link, bound to a specific deployed contract.
func NewLinkTransactor(address common.Address, transactor bind.ContractTransactor) (*LinkTransactor, error) {
	contract, err := bindLink(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LinkTransactor{contract: contract}, nil
}

// NewLinkFilterer creates a new log filterer instance of Link, bound to a specific deployed contract.
func NewLinkFilterer(address common.Address, filterer bind.ContractFilterer) (*LinkFilterer, error) {
	contract, err := bindLink(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LinkFilterer{contract: contract}, nil
}

// bindLink binds a generic wrapper to an already deployed contract.
func bindLink(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LinkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Link *LinkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Link.Contract.LinkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Link *LinkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Link.Contract.LinkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Link *LinkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Link.Contract.LinkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Link *LinkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Link.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Link *LinkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Link.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Link *LinkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Link.Contract.contract.Transact(opts, method, params...)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 roundId) view returns(int256)
func (_Link *LinkCaller) GetAnswer(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Link.contract.Call(opts, &out, "getAnswer", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 roundId) view returns(int256)
func (_Link *LinkSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _Link.Contract.GetAnswer(&_Link.CallOpts, roundId)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 roundId) view returns(int256)
func (_Link *LinkCallerSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _Link.Contract.GetAnswer(&_Link.CallOpts, roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 roundId) view returns(uint256)
func (_Link *LinkCaller) GetTimestamp(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Link.contract.Call(opts, &out, "getTimestamp", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 roundId) view returns(uint256)
func (_Link *LinkSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _Link.Contract.GetTimestamp(&_Link.CallOpts, roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 roundId) view returns(uint256)
func (_Link *LinkCallerSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _Link.Contract.GetTimestamp(&_Link.CallOpts, roundId)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_Link *LinkCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Link.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_Link *LinkSession) LatestAnswer() (*big.Int, error) {
	return _Link.Contract.LatestAnswer(&_Link.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_Link *LinkCallerSession) LatestAnswer() (*big.Int, error) {
	return _Link.Contract.LatestAnswer(&_Link.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_Link *LinkCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Link.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_Link *LinkSession) LatestRound() (*big.Int, error) {
	return _Link.Contract.LatestRound(&_Link.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_Link *LinkCallerSession) LatestRound() (*big.Int, error) {
	return _Link.Contract.LatestRound(&_Link.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_Link *LinkCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Link.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_Link *LinkSession) LatestTimestamp() (*big.Int, error) {
	return _Link.Contract.LatestTimestamp(&_Link.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_Link *LinkCallerSession) LatestTimestamp() (*big.Int, error) {
	return _Link.Contract.LatestTimestamp(&_Link.CallOpts)
}

// LinkAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the Link contract.
type LinkAnswerUpdatedIterator struct {
	Event *LinkAnswerUpdated // Event containing the contract specifics and raw log

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
func (it *LinkAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkAnswerUpdated)
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
		it.Event = new(LinkAnswerUpdated)
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
func (it *LinkAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LinkAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LinkAnswerUpdated represents a AnswerUpdated event raised by the Link contract.
type LinkAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_Link *LinkFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*LinkAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _Link.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &LinkAnswerUpdatedIterator{contract: _Link.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_Link *LinkFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *LinkAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _Link.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LinkAnswerUpdated)
				if err := _Link.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

// ParseAnswerUpdated is a log parse operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_Link *LinkFilterer) ParseAnswerUpdated(log types.Log) (*LinkAnswerUpdated, error) {
	event := new(LinkAnswerUpdated)
	if err := _Link.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LinkNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the Link contract.
type LinkNewRoundIterator struct {
	Event *LinkNewRound // Event containing the contract specifics and raw log

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
func (it *LinkNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkNewRound)
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
		it.Event = new(LinkNewRound)
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
func (it *LinkNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LinkNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LinkNewRound represents a NewRound event raised by the Link contract.
type LinkNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_Link *LinkFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*LinkNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _Link.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &LinkNewRoundIterator{contract: _Link.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_Link *LinkFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *LinkNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _Link.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LinkNewRound)
				if err := _Link.contract.UnpackLog(event, "NewRound", log); err != nil {
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

// ParseNewRound is a log parse operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_Link *LinkFilterer) ParseNewRound(log types.Log) (*LinkNewRound, error) {
	event := new(LinkNewRound)
	if err := _Link.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
