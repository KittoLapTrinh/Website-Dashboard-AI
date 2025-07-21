// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
	_ = abi.ConvertType
)

// DataStructsDataPoint is an auto generated low-level Go binding around an user-defined struct.
type DataStructsDataPoint struct {
	Timestamp *big.Int
	Value     *big.Int
	Change    *big.Int
}

// TimeSeriesContractMetaData contains all meta data concerning the TimeSeriesContract contract.
var TimeSeriesContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"change\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structDataStructs.DataPoint\",\"name\":\"newPoint\",\"type\":\"tuple\"}],\"name\":\"DataPointAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_change\",\"type\":\"int256\"}],\"name\":\"pushDataPoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getTimeSeriesData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"change\",\"type\":\"int256\"}],\"internalType\":\"structDataStructs.DataPoint[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"timeSeriesData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"change\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TimeSeriesContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TimeSeriesContractMetaData.ABI instead.
var TimeSeriesContractABI = TimeSeriesContractMetaData.ABI

// TimeSeriesContract is an auto generated Go binding around an Ethereum contract.
type TimeSeriesContract struct {
	TimeSeriesContractCaller     // Read-only binding to the contract
	TimeSeriesContractTransactor // Write-only binding to the contract
	TimeSeriesContractFilterer   // Log filterer for contract events
}

// TimeSeriesContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimeSeriesContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeSeriesContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimeSeriesContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeSeriesContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimeSeriesContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeSeriesContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimeSeriesContractSession struct {
	Contract     *TimeSeriesContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TimeSeriesContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimeSeriesContractCallerSession struct {
	Contract *TimeSeriesContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TimeSeriesContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimeSeriesContractTransactorSession struct {
	Contract     *TimeSeriesContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TimeSeriesContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimeSeriesContractRaw struct {
	Contract *TimeSeriesContract // Generic contract binding to access the raw methods on
}

// TimeSeriesContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimeSeriesContractCallerRaw struct {
	Contract *TimeSeriesContractCaller // Generic read-only contract binding to access the raw methods on
}

// TimeSeriesContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimeSeriesContractTransactorRaw struct {
	Contract *TimeSeriesContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimeSeriesContract creates a new instance of TimeSeriesContract, bound to a specific deployed contract.
func NewTimeSeriesContract(address common.Address, backend bind.ContractBackend) (*TimeSeriesContract, error) {
	contract, err := bindTimeSeriesContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TimeSeriesContract{TimeSeriesContractCaller: TimeSeriesContractCaller{contract: contract}, TimeSeriesContractTransactor: TimeSeriesContractTransactor{contract: contract}, TimeSeriesContractFilterer: TimeSeriesContractFilterer{contract: contract}}, nil
}

// NewTimeSeriesContractCaller creates a new read-only instance of TimeSeriesContract, bound to a specific deployed contract.
func NewTimeSeriesContractCaller(address common.Address, caller bind.ContractCaller) (*TimeSeriesContractCaller, error) {
	contract, err := bindTimeSeriesContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimeSeriesContractCaller{contract: contract}, nil
}

// NewTimeSeriesContractTransactor creates a new write-only instance of TimeSeriesContract, bound to a specific deployed contract.
func NewTimeSeriesContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TimeSeriesContractTransactor, error) {
	contract, err := bindTimeSeriesContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimeSeriesContractTransactor{contract: contract}, nil
}

// NewTimeSeriesContractFilterer creates a new log filterer instance of TimeSeriesContract, bound to a specific deployed contract.
func NewTimeSeriesContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TimeSeriesContractFilterer, error) {
	contract, err := bindTimeSeriesContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimeSeriesContractFilterer{contract: contract}, nil
}

// bindTimeSeriesContract binds a generic wrapper to an already deployed contract.
func bindTimeSeriesContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TimeSeriesContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimeSeriesContract *TimeSeriesContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimeSeriesContract.Contract.TimeSeriesContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimeSeriesContract *TimeSeriesContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeSeriesContract.Contract.TimeSeriesContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimeSeriesContract *TimeSeriesContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimeSeriesContract.Contract.TimeSeriesContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimeSeriesContract *TimeSeriesContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimeSeriesContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimeSeriesContract *TimeSeriesContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeSeriesContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimeSeriesContract *TimeSeriesContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimeSeriesContract.Contract.contract.Transact(opts, method, params...)
}

// GetTimeSeriesData is a free data retrieval call binding the contract method 0x7918f328.
//
// Solidity: function getTimeSeriesData(string _key) view returns((uint256,uint256,int256)[])
func (_TimeSeriesContract *TimeSeriesContractCaller) GetTimeSeriesData(opts *bind.CallOpts, _key string) ([]DataStructsDataPoint, error) {
	var out []interface{}
	err := _TimeSeriesContract.contract.Call(opts, &out, "getTimeSeriesData", _key)

	if err != nil {
		return *new([]DataStructsDataPoint), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataStructsDataPoint)).(*[]DataStructsDataPoint)

	return out0, err

}

// GetTimeSeriesData is a free data retrieval call binding the contract method 0x7918f328.
//
// Solidity: function getTimeSeriesData(string _key) view returns((uint256,uint256,int256)[])
func (_TimeSeriesContract *TimeSeriesContractSession) GetTimeSeriesData(_key string) ([]DataStructsDataPoint, error) {
	return _TimeSeriesContract.Contract.GetTimeSeriesData(&_TimeSeriesContract.CallOpts, _key)
}

// GetTimeSeriesData is a free data retrieval call binding the contract method 0x7918f328.
//
// Solidity: function getTimeSeriesData(string _key) view returns((uint256,uint256,int256)[])
func (_TimeSeriesContract *TimeSeriesContractCallerSession) GetTimeSeriesData(_key string) ([]DataStructsDataPoint, error) {
	return _TimeSeriesContract.Contract.GetTimeSeriesData(&_TimeSeriesContract.CallOpts, _key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TimeSeriesContract *TimeSeriesContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TimeSeriesContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TimeSeriesContract *TimeSeriesContractSession) Owner() (common.Address, error) {
	return _TimeSeriesContract.Contract.Owner(&_TimeSeriesContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TimeSeriesContract *TimeSeriesContractCallerSession) Owner() (common.Address, error) {
	return _TimeSeriesContract.Contract.Owner(&_TimeSeriesContract.CallOpts)
}

// TimeSeriesData is a free data retrieval call binding the contract method 0xf65d99aa.
//
// Solidity: function timeSeriesData(string , uint256 ) view returns(uint256 timestamp, uint256 value, int256 change)
func (_TimeSeriesContract *TimeSeriesContractCaller) TimeSeriesData(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Timestamp *big.Int
	Value     *big.Int
	Change    *big.Int
}, error) {
	var out []interface{}
	err := _TimeSeriesContract.contract.Call(opts, &out, "timeSeriesData", arg0, arg1)

	outstruct := new(struct {
		Timestamp *big.Int
		Value     *big.Int
		Change    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Change = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TimeSeriesData is a free data retrieval call binding the contract method 0xf65d99aa.
//
// Solidity: function timeSeriesData(string , uint256 ) view returns(uint256 timestamp, uint256 value, int256 change)
func (_TimeSeriesContract *TimeSeriesContractSession) TimeSeriesData(arg0 string, arg1 *big.Int) (struct {
	Timestamp *big.Int
	Value     *big.Int
	Change    *big.Int
}, error) {
	return _TimeSeriesContract.Contract.TimeSeriesData(&_TimeSeriesContract.CallOpts, arg0, arg1)
}

// TimeSeriesData is a free data retrieval call binding the contract method 0xf65d99aa.
//
// Solidity: function timeSeriesData(string , uint256 ) view returns(uint256 timestamp, uint256 value, int256 change)
func (_TimeSeriesContract *TimeSeriesContractCallerSession) TimeSeriesData(arg0 string, arg1 *big.Int) (struct {
	Timestamp *big.Int
	Value     *big.Int
	Change    *big.Int
}, error) {
	return _TimeSeriesContract.Contract.TimeSeriesData(&_TimeSeriesContract.CallOpts, arg0, arg1)
}

// PushDataPoint is a paid mutator transaction binding the contract method 0x53bfac46.
//
// Solidity: function pushDataPoint(string _key, uint256 _value, int256 _change) returns()
func (_TimeSeriesContract *TimeSeriesContractTransactor) PushDataPoint(opts *bind.TransactOpts, _key string, _value *big.Int, _change *big.Int) (*types.Transaction, error) {
	return _TimeSeriesContract.contract.Transact(opts, "pushDataPoint", _key, _value, _change)
}

// PushDataPoint is a paid mutator transaction binding the contract method 0x53bfac46.
//
// Solidity: function pushDataPoint(string _key, uint256 _value, int256 _change) returns()
func (_TimeSeriesContract *TimeSeriesContractSession) PushDataPoint(_key string, _value *big.Int, _change *big.Int) (*types.Transaction, error) {
	return _TimeSeriesContract.Contract.PushDataPoint(&_TimeSeriesContract.TransactOpts, _key, _value, _change)
}

// PushDataPoint is a paid mutator transaction binding the contract method 0x53bfac46.
//
// Solidity: function pushDataPoint(string _key, uint256 _value, int256 _change) returns()
func (_TimeSeriesContract *TimeSeriesContractTransactorSession) PushDataPoint(_key string, _value *big.Int, _change *big.Int) (*types.Transaction, error) {
	return _TimeSeriesContract.Contract.PushDataPoint(&_TimeSeriesContract.TransactOpts, _key, _value, _change)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TimeSeriesContract *TimeSeriesContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TimeSeriesContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TimeSeriesContract *TimeSeriesContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TimeSeriesContract.Contract.TransferOwnership(&_TimeSeriesContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TimeSeriesContract *TimeSeriesContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TimeSeriesContract.Contract.TransferOwnership(&_TimeSeriesContract.TransactOpts, newOwner)
}

// TimeSeriesContractDataPointAddedIterator is returned from FilterDataPointAdded and is used to iterate over the raw logs and unpacked data for DataPointAdded events raised by the TimeSeriesContract contract.
type TimeSeriesContractDataPointAddedIterator struct {
	Event *TimeSeriesContractDataPointAdded // Event containing the contract specifics and raw log

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
func (it *TimeSeriesContractDataPointAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimeSeriesContractDataPointAdded)
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
		it.Event = new(TimeSeriesContractDataPointAdded)
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
func (it *TimeSeriesContractDataPointAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimeSeriesContractDataPointAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimeSeriesContractDataPointAdded represents a DataPointAdded event raised by the TimeSeriesContract contract.
type TimeSeriesContractDataPointAdded struct {
	Key      string
	NewPoint DataStructsDataPoint
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDataPointAdded is a free log retrieval operation binding the contract event 0x2f533650ae9e5d5131aa36c8d5548a73369e2ea487dc351fcce0fcf93710a6d4.
//
// Solidity: event DataPointAdded(string key, (uint256,uint256,int256) newPoint)
func (_TimeSeriesContract *TimeSeriesContractFilterer) FilterDataPointAdded(opts *bind.FilterOpts) (*TimeSeriesContractDataPointAddedIterator, error) {

	logs, sub, err := _TimeSeriesContract.contract.FilterLogs(opts, "DataPointAdded")
	if err != nil {
		return nil, err
	}
	return &TimeSeriesContractDataPointAddedIterator{contract: _TimeSeriesContract.contract, event: "DataPointAdded", logs: logs, sub: sub}, nil
}

// WatchDataPointAdded is a free log subscription operation binding the contract event 0x2f533650ae9e5d5131aa36c8d5548a73369e2ea487dc351fcce0fcf93710a6d4.
//
// Solidity: event DataPointAdded(string key, (uint256,uint256,int256) newPoint)
func (_TimeSeriesContract *TimeSeriesContractFilterer) WatchDataPointAdded(opts *bind.WatchOpts, sink chan<- *TimeSeriesContractDataPointAdded) (event.Subscription, error) {

	logs, sub, err := _TimeSeriesContract.contract.WatchLogs(opts, "DataPointAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimeSeriesContractDataPointAdded)
				if err := _TimeSeriesContract.contract.UnpackLog(event, "DataPointAdded", log); err != nil {
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

// ParseDataPointAdded is a log parse operation binding the contract event 0x2f533650ae9e5d5131aa36c8d5548a73369e2ea487dc351fcce0fcf93710a6d4.
//
// Solidity: event DataPointAdded(string key, (uint256,uint256,int256) newPoint)
func (_TimeSeriesContract *TimeSeriesContractFilterer) ParseDataPointAdded(log types.Log) (*TimeSeriesContractDataPointAdded, error) {
	event := new(TimeSeriesContractDataPointAdded)
	if err := _TimeSeriesContract.contract.UnpackLog(event, "DataPointAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimeSeriesContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TimeSeriesContract contract.
type TimeSeriesContractOwnershipTransferredIterator struct {
	Event *TimeSeriesContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TimeSeriesContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimeSeriesContractOwnershipTransferred)
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
		it.Event = new(TimeSeriesContractOwnershipTransferred)
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
func (it *TimeSeriesContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimeSeriesContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimeSeriesContractOwnershipTransferred represents a OwnershipTransferred event raised by the TimeSeriesContract contract.
type TimeSeriesContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TimeSeriesContract *TimeSeriesContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TimeSeriesContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TimeSeriesContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TimeSeriesContractOwnershipTransferredIterator{contract: _TimeSeriesContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TimeSeriesContract *TimeSeriesContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TimeSeriesContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TimeSeriesContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimeSeriesContractOwnershipTransferred)
				if err := _TimeSeriesContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TimeSeriesContract *TimeSeriesContractFilterer) ParseOwnershipTransferred(log types.Log) (*TimeSeriesContractOwnershipTransferred, error) {
	event := new(TimeSeriesContractOwnershipTransferred)
	if err := _TimeSeriesContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
