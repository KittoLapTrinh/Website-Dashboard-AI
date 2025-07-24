// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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

// KeyValueContractMetaData contains all meta data concerning the KeyValueContract contract.
var KeyValueContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"newValue\",\"type\":\"int256\"}],\"name\":\"ReturnValueUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"SingleValueUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_newValue\",\"type\":\"int256\"}],\"name\":\"updateReturnData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"}],\"name\":\"updateSingleValueData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getReturnData\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getSingleValueData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"returnData\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"singleValueData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// KeyValueContractABI is the input ABI used to generate the binding from.
// Deprecated: Use KeyValueContractMetaData.ABI instead.
var KeyValueContractABI = KeyValueContractMetaData.ABI

// KeyValueContract is an auto generated Go binding around an Ethereum contract.
type KeyValueContract struct {
	KeyValueContractCaller     // Read-only binding to the contract
	KeyValueContractTransactor // Write-only binding to the contract
	KeyValueContractFilterer   // Log filterer for contract events
}

// KeyValueContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyValueContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyValueContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyValueContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyValueContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyValueContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyValueContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyValueContractSession struct {
	Contract     *KeyValueContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeyValueContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyValueContractCallerSession struct {
	Contract *KeyValueContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// KeyValueContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyValueContractTransactorSession struct {
	Contract     *KeyValueContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// KeyValueContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyValueContractRaw struct {
	Contract *KeyValueContract // Generic contract binding to access the raw methods on
}

// KeyValueContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyValueContractCallerRaw struct {
	Contract *KeyValueContractCaller // Generic read-only contract binding to access the raw methods on
}

// KeyValueContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyValueContractTransactorRaw struct {
	Contract *KeyValueContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyValueContract creates a new instance of KeyValueContract, bound to a specific deployed contract.
func NewKeyValueContract(address common.Address, backend bind.ContractBackend) (*KeyValueContract, error) {
	contract, err := bindKeyValueContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyValueContract{KeyValueContractCaller: KeyValueContractCaller{contract: contract}, KeyValueContractTransactor: KeyValueContractTransactor{contract: contract}, KeyValueContractFilterer: KeyValueContractFilterer{contract: contract}}, nil
}

// NewKeyValueContractCaller creates a new read-only instance of KeyValueContract, bound to a specific deployed contract.
func NewKeyValueContractCaller(address common.Address, caller bind.ContractCaller) (*KeyValueContractCaller, error) {
	contract, err := bindKeyValueContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyValueContractCaller{contract: contract}, nil
}

// NewKeyValueContractTransactor creates a new write-only instance of KeyValueContract, bound to a specific deployed contract.
func NewKeyValueContractTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyValueContractTransactor, error) {
	contract, err := bindKeyValueContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyValueContractTransactor{contract: contract}, nil
}

// NewKeyValueContractFilterer creates a new log filterer instance of KeyValueContract, bound to a specific deployed contract.
func NewKeyValueContractFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyValueContractFilterer, error) {
	contract, err := bindKeyValueContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyValueContractFilterer{contract: contract}, nil
}

// bindKeyValueContract binds a generic wrapper to an already deployed contract.
func bindKeyValueContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KeyValueContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyValueContract *KeyValueContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyValueContract.Contract.KeyValueContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyValueContract *KeyValueContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyValueContract.Contract.KeyValueContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyValueContract *KeyValueContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyValueContract.Contract.KeyValueContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyValueContract *KeyValueContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyValueContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyValueContract *KeyValueContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyValueContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyValueContract *KeyValueContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyValueContract.Contract.contract.Transact(opts, method, params...)
}

// GetReturnData is a free data retrieval call binding the contract method 0x56b0bb32.
//
// Solidity: function getReturnData(string _key) view returns(int256)
func (_KeyValueContract *KeyValueContractCaller) GetReturnData(opts *bind.CallOpts, _key string) (*big.Int, error) {
	var out []interface{}
	err := _KeyValueContract.contract.Call(opts, &out, "getReturnData", _key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReturnData is a free data retrieval call binding the contract method 0x56b0bb32.
//
// Solidity: function getReturnData(string _key) view returns(int256)
func (_KeyValueContract *KeyValueContractSession) GetReturnData(_key string) (*big.Int, error) {
	return _KeyValueContract.Contract.GetReturnData(&_KeyValueContract.CallOpts, _key)
}

// GetReturnData is a free data retrieval call binding the contract method 0x56b0bb32.
//
// Solidity: function getReturnData(string _key) view returns(int256)
func (_KeyValueContract *KeyValueContractCallerSession) GetReturnData(_key string) (*big.Int, error) {
	return _KeyValueContract.Contract.GetReturnData(&_KeyValueContract.CallOpts, _key)
}

// GetSingleValueData is a free data retrieval call binding the contract method 0x166d8b41.
//
// Solidity: function getSingleValueData(string _key) view returns(uint256)
func (_KeyValueContract *KeyValueContractCaller) GetSingleValueData(opts *bind.CallOpts, _key string) (*big.Int, error) {
	var out []interface{}
	err := _KeyValueContract.contract.Call(opts, &out, "getSingleValueData", _key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSingleValueData is a free data retrieval call binding the contract method 0x166d8b41.
//
// Solidity: function getSingleValueData(string _key) view returns(uint256)
func (_KeyValueContract *KeyValueContractSession) GetSingleValueData(_key string) (*big.Int, error) {
	return _KeyValueContract.Contract.GetSingleValueData(&_KeyValueContract.CallOpts, _key)
}

// GetSingleValueData is a free data retrieval call binding the contract method 0x166d8b41.
//
// Solidity: function getSingleValueData(string _key) view returns(uint256)
func (_KeyValueContract *KeyValueContractCallerSession) GetSingleValueData(_key string) (*big.Int, error) {
	return _KeyValueContract.Contract.GetSingleValueData(&_KeyValueContract.CallOpts, _key)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyValueContract *KeyValueContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeyValueContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyValueContract *KeyValueContractSession) Owner() (common.Address, error) {
	return _KeyValueContract.Contract.Owner(&_KeyValueContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KeyValueContract *KeyValueContractCallerSession) Owner() (common.Address, error) {
	return _KeyValueContract.Contract.Owner(&_KeyValueContract.CallOpts)
}

// ReturnData is a free data retrieval call binding the contract method 0xe5f92e86.
//
// Solidity: function returnData(string ) view returns(int256)
func (_KeyValueContract *KeyValueContractCaller) ReturnData(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _KeyValueContract.contract.Call(opts, &out, "returnData", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReturnData is a free data retrieval call binding the contract method 0xe5f92e86.
//
// Solidity: function returnData(string ) view returns(int256)
func (_KeyValueContract *KeyValueContractSession) ReturnData(arg0 string) (*big.Int, error) {
	return _KeyValueContract.Contract.ReturnData(&_KeyValueContract.CallOpts, arg0)
}

// ReturnData is a free data retrieval call binding the contract method 0xe5f92e86.
//
// Solidity: function returnData(string ) view returns(int256)
func (_KeyValueContract *KeyValueContractCallerSession) ReturnData(arg0 string) (*big.Int, error) {
	return _KeyValueContract.Contract.ReturnData(&_KeyValueContract.CallOpts, arg0)
}

// SingleValueData is a free data retrieval call binding the contract method 0x5a2ded70.
//
// Solidity: function singleValueData(string ) view returns(uint256)
func (_KeyValueContract *KeyValueContractCaller) SingleValueData(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _KeyValueContract.contract.Call(opts, &out, "singleValueData", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SingleValueData is a free data retrieval call binding the contract method 0x5a2ded70.
//
// Solidity: function singleValueData(string ) view returns(uint256)
func (_KeyValueContract *KeyValueContractSession) SingleValueData(arg0 string) (*big.Int, error) {
	return _KeyValueContract.Contract.SingleValueData(&_KeyValueContract.CallOpts, arg0)
}

// SingleValueData is a free data retrieval call binding the contract method 0x5a2ded70.
//
// Solidity: function singleValueData(string ) view returns(uint256)
func (_KeyValueContract *KeyValueContractCallerSession) SingleValueData(arg0 string) (*big.Int, error) {
	return _KeyValueContract.Contract.SingleValueData(&_KeyValueContract.CallOpts, arg0)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyValueContract *KeyValueContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KeyValueContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyValueContract *KeyValueContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KeyValueContract.Contract.TransferOwnership(&_KeyValueContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KeyValueContract *KeyValueContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KeyValueContract.Contract.TransferOwnership(&_KeyValueContract.TransactOpts, newOwner)
}

// UpdateReturnData is a paid mutator transaction binding the contract method 0x4fdc63a2.
//
// Solidity: function updateReturnData(string _key, int256 _newValue) returns()
func (_KeyValueContract *KeyValueContractTransactor) UpdateReturnData(opts *bind.TransactOpts, _key string, _newValue *big.Int) (*types.Transaction, error) {
	return _KeyValueContract.contract.Transact(opts, "updateReturnData", _key, _newValue)
}

// UpdateReturnData is a paid mutator transaction binding the contract method 0x4fdc63a2.
//
// Solidity: function updateReturnData(string _key, int256 _newValue) returns()
func (_KeyValueContract *KeyValueContractSession) UpdateReturnData(_key string, _newValue *big.Int) (*types.Transaction, error) {
	return _KeyValueContract.Contract.UpdateReturnData(&_KeyValueContract.TransactOpts, _key, _newValue)
}

// UpdateReturnData is a paid mutator transaction binding the contract method 0x4fdc63a2.
//
// Solidity: function updateReturnData(string _key, int256 _newValue) returns()
func (_KeyValueContract *KeyValueContractTransactorSession) UpdateReturnData(_key string, _newValue *big.Int) (*types.Transaction, error) {
	return _KeyValueContract.Contract.UpdateReturnData(&_KeyValueContract.TransactOpts, _key, _newValue)
}

// UpdateSingleValueData is a paid mutator transaction binding the contract method 0xa35317ec.
//
// Solidity: function updateSingleValueData(string _key, uint256 _newValue) returns()
func (_KeyValueContract *KeyValueContractTransactor) UpdateSingleValueData(opts *bind.TransactOpts, _key string, _newValue *big.Int) (*types.Transaction, error) {
	return _KeyValueContract.contract.Transact(opts, "updateSingleValueData", _key, _newValue)
}

// UpdateSingleValueData is a paid mutator transaction binding the contract method 0xa35317ec.
//
// Solidity: function updateSingleValueData(string _key, uint256 _newValue) returns()
func (_KeyValueContract *KeyValueContractSession) UpdateSingleValueData(_key string, _newValue *big.Int) (*types.Transaction, error) {
	return _KeyValueContract.Contract.UpdateSingleValueData(&_KeyValueContract.TransactOpts, _key, _newValue)
}

// UpdateSingleValueData is a paid mutator transaction binding the contract method 0xa35317ec.
//
// Solidity: function updateSingleValueData(string _key, uint256 _newValue) returns()
func (_KeyValueContract *KeyValueContractTransactorSession) UpdateSingleValueData(_key string, _newValue *big.Int) (*types.Transaction, error) {
	return _KeyValueContract.Contract.UpdateSingleValueData(&_KeyValueContract.TransactOpts, _key, _newValue)
}

// KeyValueContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KeyValueContract contract.
type KeyValueContractOwnershipTransferredIterator struct {
	Event *KeyValueContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KeyValueContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyValueContractOwnershipTransferred)
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
		it.Event = new(KeyValueContractOwnershipTransferred)
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
func (it *KeyValueContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyValueContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyValueContractOwnershipTransferred represents a OwnershipTransferred event raised by the KeyValueContract contract.
type KeyValueContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KeyValueContract *KeyValueContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KeyValueContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeyValueContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KeyValueContractOwnershipTransferredIterator{contract: _KeyValueContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KeyValueContract *KeyValueContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KeyValueContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KeyValueContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyValueContractOwnershipTransferred)
				if err := _KeyValueContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_KeyValueContract *KeyValueContractFilterer) ParseOwnershipTransferred(log types.Log) (*KeyValueContractOwnershipTransferred, error) {
	event := new(KeyValueContractOwnershipTransferred)
	if err := _KeyValueContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyValueContractReturnValueUpdatedIterator is returned from FilterReturnValueUpdated and is used to iterate over the raw logs and unpacked data for ReturnValueUpdated events raised by the KeyValueContract contract.
type KeyValueContractReturnValueUpdatedIterator struct {
	Event *KeyValueContractReturnValueUpdated // Event containing the contract specifics and raw log

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
func (it *KeyValueContractReturnValueUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyValueContractReturnValueUpdated)
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
		it.Event = new(KeyValueContractReturnValueUpdated)
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
func (it *KeyValueContractReturnValueUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyValueContractReturnValueUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyValueContractReturnValueUpdated represents a ReturnValueUpdated event raised by the KeyValueContract contract.
type KeyValueContractReturnValueUpdated struct {
	Key      string
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReturnValueUpdated is a free log retrieval operation binding the contract event 0x67b065c9e8250eb9d63e7fdb9d3ca3c6def37194057697d52d21c0234bcac54e.
//
// Solidity: event ReturnValueUpdated(string key, int256 newValue)
func (_KeyValueContract *KeyValueContractFilterer) FilterReturnValueUpdated(opts *bind.FilterOpts) (*KeyValueContractReturnValueUpdatedIterator, error) {

	logs, sub, err := _KeyValueContract.contract.FilterLogs(opts, "ReturnValueUpdated")
	if err != nil {
		return nil, err
	}
	return &KeyValueContractReturnValueUpdatedIterator{contract: _KeyValueContract.contract, event: "ReturnValueUpdated", logs: logs, sub: sub}, nil
}

// WatchReturnValueUpdated is a free log subscription operation binding the contract event 0x67b065c9e8250eb9d63e7fdb9d3ca3c6def37194057697d52d21c0234bcac54e.
//
// Solidity: event ReturnValueUpdated(string key, int256 newValue)
func (_KeyValueContract *KeyValueContractFilterer) WatchReturnValueUpdated(opts *bind.WatchOpts, sink chan<- *KeyValueContractReturnValueUpdated) (event.Subscription, error) {

	logs, sub, err := _KeyValueContract.contract.WatchLogs(opts, "ReturnValueUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyValueContractReturnValueUpdated)
				if err := _KeyValueContract.contract.UnpackLog(event, "ReturnValueUpdated", log); err != nil {
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

// ParseReturnValueUpdated is a log parse operation binding the contract event 0x67b065c9e8250eb9d63e7fdb9d3ca3c6def37194057697d52d21c0234bcac54e.
//
// Solidity: event ReturnValueUpdated(string key, int256 newValue)
func (_KeyValueContract *KeyValueContractFilterer) ParseReturnValueUpdated(log types.Log) (*KeyValueContractReturnValueUpdated, error) {
	event := new(KeyValueContractReturnValueUpdated)
	if err := _KeyValueContract.contract.UnpackLog(event, "ReturnValueUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KeyValueContractSingleValueUpdatedIterator is returned from FilterSingleValueUpdated and is used to iterate over the raw logs and unpacked data for SingleValueUpdated events raised by the KeyValueContract contract.
type KeyValueContractSingleValueUpdatedIterator struct {
	Event *KeyValueContractSingleValueUpdated // Event containing the contract specifics and raw log

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
func (it *KeyValueContractSingleValueUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyValueContractSingleValueUpdated)
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
		it.Event = new(KeyValueContractSingleValueUpdated)
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
func (it *KeyValueContractSingleValueUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyValueContractSingleValueUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyValueContractSingleValueUpdated represents a SingleValueUpdated event raised by the KeyValueContract contract.
type KeyValueContractSingleValueUpdated struct {
	Key      string
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSingleValueUpdated is a free log retrieval operation binding the contract event 0xfff99ebe4547d53734136b165e336a3484c50aa57fd03b2af7015323b4cb8aea.
//
// Solidity: event SingleValueUpdated(string key, uint256 newValue)
func (_KeyValueContract *KeyValueContractFilterer) FilterSingleValueUpdated(opts *bind.FilterOpts) (*KeyValueContractSingleValueUpdatedIterator, error) {

	logs, sub, err := _KeyValueContract.contract.FilterLogs(opts, "SingleValueUpdated")
	if err != nil {
		return nil, err
	}
	return &KeyValueContractSingleValueUpdatedIterator{contract: _KeyValueContract.contract, event: "SingleValueUpdated", logs: logs, sub: sub}, nil
}

// WatchSingleValueUpdated is a free log subscription operation binding the contract event 0xfff99ebe4547d53734136b165e336a3484c50aa57fd03b2af7015323b4cb8aea.
//
// Solidity: event SingleValueUpdated(string key, uint256 newValue)
func (_KeyValueContract *KeyValueContractFilterer) WatchSingleValueUpdated(opts *bind.WatchOpts, sink chan<- *KeyValueContractSingleValueUpdated) (event.Subscription, error) {

	logs, sub, err := _KeyValueContract.contract.WatchLogs(opts, "SingleValueUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyValueContractSingleValueUpdated)
				if err := _KeyValueContract.contract.UnpackLog(event, "SingleValueUpdated", log); err != nil {
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

// ParseSingleValueUpdated is a log parse operation binding the contract event 0xfff99ebe4547d53734136b165e336a3484c50aa57fd03b2af7015323b4cb8aea.
//
// Solidity: event SingleValueUpdated(string key, uint256 newValue)
func (_KeyValueContract *KeyValueContractFilterer) ParseSingleValueUpdated(log types.Log) (*KeyValueContractSingleValueUpdated, error) {
	event := new(KeyValueContractSingleValueUpdated)
	if err := _KeyValueContract.contract.UnpackLog(event, "SingleValueUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
