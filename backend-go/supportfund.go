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

// DataStructsFundItem is an auto generated low-level Go binding around an user-defined struct.
type DataStructsFundItem struct {
	Name        string
	Price       *big.Int
	IconId      string
	Count       *big.Int
	AvatarColor string
}

// SupportFundContractMetaData contains all meta data concerning the SupportFundContract contract.
var SupportFundContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dayKey\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"avatarColor\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structDataStructs.FundItem[]\",\"name\":\"newItems\",\"type\":\"tuple[]\"}],\"name\":\"FundItemsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_dayKey\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"avatarColor\",\"type\":\"string\"}],\"internalType\":\"structDataStructs.FundItem[]\",\"name\":\"_newItems\",\"type\":\"tuple[]\"}],\"name\":\"updateFundItemsForDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fundItemsByDay\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"avatarColor\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_day\",\"type\":\"string\"}],\"name\":\"getFundItemsByDay\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"avatarColor\",\"type\":\"string\"}],\"internalType\":\"structDataStructs.FundItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SupportFundContractABI is the input ABI used to generate the binding from.
// Deprecated: Use SupportFundContractMetaData.ABI instead.
var SupportFundContractABI = SupportFundContractMetaData.ABI

// SupportFundContract is an auto generated Go binding around an Ethereum contract.
type SupportFundContract struct {
	SupportFundContractCaller     // Read-only binding to the contract
	SupportFundContractTransactor // Write-only binding to the contract
	SupportFundContractFilterer   // Log filterer for contract events
}

// SupportFundContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SupportFundContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupportFundContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SupportFundContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupportFundContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SupportFundContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupportFundContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SupportFundContractSession struct {
	Contract     *SupportFundContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SupportFundContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SupportFundContractCallerSession struct {
	Contract *SupportFundContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SupportFundContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SupportFundContractTransactorSession struct {
	Contract     *SupportFundContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SupportFundContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SupportFundContractRaw struct {
	Contract *SupportFundContract // Generic contract binding to access the raw methods on
}

// SupportFundContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SupportFundContractCallerRaw struct {
	Contract *SupportFundContractCaller // Generic read-only contract binding to access the raw methods on
}

// SupportFundContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SupportFundContractTransactorRaw struct {
	Contract *SupportFundContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSupportFundContract creates a new instance of SupportFundContract, bound to a specific deployed contract.
func NewSupportFundContract(address common.Address, backend bind.ContractBackend) (*SupportFundContract, error) {
	contract, err := bindSupportFundContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SupportFundContract{SupportFundContractCaller: SupportFundContractCaller{contract: contract}, SupportFundContractTransactor: SupportFundContractTransactor{contract: contract}, SupportFundContractFilterer: SupportFundContractFilterer{contract: contract}}, nil
}

// NewSupportFundContractCaller creates a new read-only instance of SupportFundContract, bound to a specific deployed contract.
func NewSupportFundContractCaller(address common.Address, caller bind.ContractCaller) (*SupportFundContractCaller, error) {
	contract, err := bindSupportFundContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SupportFundContractCaller{contract: contract}, nil
}

// NewSupportFundContractTransactor creates a new write-only instance of SupportFundContract, bound to a specific deployed contract.
func NewSupportFundContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SupportFundContractTransactor, error) {
	contract, err := bindSupportFundContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SupportFundContractTransactor{contract: contract}, nil
}

// NewSupportFundContractFilterer creates a new log filterer instance of SupportFundContract, bound to a specific deployed contract.
func NewSupportFundContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SupportFundContractFilterer, error) {
	contract, err := bindSupportFundContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SupportFundContractFilterer{contract: contract}, nil
}

// bindSupportFundContract binds a generic wrapper to an already deployed contract.
func bindSupportFundContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SupportFundContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SupportFundContract *SupportFundContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SupportFundContract.Contract.SupportFundContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SupportFundContract *SupportFundContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupportFundContract.Contract.SupportFundContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SupportFundContract *SupportFundContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SupportFundContract.Contract.SupportFundContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SupportFundContract *SupportFundContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SupportFundContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SupportFundContract *SupportFundContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupportFundContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SupportFundContract *SupportFundContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SupportFundContract.Contract.contract.Transact(opts, method, params...)
}

// FundItemsByDay is a free data retrieval call binding the contract method 0x36a48ee5.
//
// Solidity: function fundItemsByDay(string , uint256 ) view returns(string name, uint256 price, string iconId, uint256 count, string avatarColor)
func (_SupportFundContract *SupportFundContractCaller) FundItemsByDay(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Name        string
	Price       *big.Int
	IconId      string
	Count       *big.Int
	AvatarColor string
}, error) {
	var out []interface{}
	err := _SupportFundContract.contract.Call(opts, &out, "fundItemsByDay", arg0, arg1)

	outstruct := new(struct {
		Name        string
		Price       *big.Int
		IconId      string
		Count       *big.Int
		AvatarColor string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IconId = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Count = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AvatarColor = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// FundItemsByDay is a free data retrieval call binding the contract method 0x36a48ee5.
//
// Solidity: function fundItemsByDay(string , uint256 ) view returns(string name, uint256 price, string iconId, uint256 count, string avatarColor)
func (_SupportFundContract *SupportFundContractSession) FundItemsByDay(arg0 string, arg1 *big.Int) (struct {
	Name        string
	Price       *big.Int
	IconId      string
	Count       *big.Int
	AvatarColor string
}, error) {
	return _SupportFundContract.Contract.FundItemsByDay(&_SupportFundContract.CallOpts, arg0, arg1)
}

// FundItemsByDay is a free data retrieval call binding the contract method 0x36a48ee5.
//
// Solidity: function fundItemsByDay(string , uint256 ) view returns(string name, uint256 price, string iconId, uint256 count, string avatarColor)
func (_SupportFundContract *SupportFundContractCallerSession) FundItemsByDay(arg0 string, arg1 *big.Int) (struct {
	Name        string
	Price       *big.Int
	IconId      string
	Count       *big.Int
	AvatarColor string
}, error) {
	return _SupportFundContract.Contract.FundItemsByDay(&_SupportFundContract.CallOpts, arg0, arg1)
}

// GetFundItemsByDay is a free data retrieval call binding the contract method 0x6ca07056.
//
// Solidity: function getFundItemsByDay(string _day) view returns((string,uint256,string,uint256,string)[])
func (_SupportFundContract *SupportFundContractCaller) GetFundItemsByDay(opts *bind.CallOpts, _day string) ([]DataStructsFundItem, error) {
	var out []interface{}
	err := _SupportFundContract.contract.Call(opts, &out, "getFundItemsByDay", _day)

	if err != nil {
		return *new([]DataStructsFundItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataStructsFundItem)).(*[]DataStructsFundItem)

	return out0, err

}

// GetFundItemsByDay is a free data retrieval call binding the contract method 0x6ca07056.
//
// Solidity: function getFundItemsByDay(string _day) view returns((string,uint256,string,uint256,string)[])
func (_SupportFundContract *SupportFundContractSession) GetFundItemsByDay(_day string) ([]DataStructsFundItem, error) {
	return _SupportFundContract.Contract.GetFundItemsByDay(&_SupportFundContract.CallOpts, _day)
}

// GetFundItemsByDay is a free data retrieval call binding the contract method 0x6ca07056.
//
// Solidity: function getFundItemsByDay(string _day) view returns((string,uint256,string,uint256,string)[])
func (_SupportFundContract *SupportFundContractCallerSession) GetFundItemsByDay(_day string) ([]DataStructsFundItem, error) {
	return _SupportFundContract.Contract.GetFundItemsByDay(&_SupportFundContract.CallOpts, _day)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SupportFundContract *SupportFundContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SupportFundContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SupportFundContract *SupportFundContractSession) Owner() (common.Address, error) {
	return _SupportFundContract.Contract.Owner(&_SupportFundContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SupportFundContract *SupportFundContractCallerSession) Owner() (common.Address, error) {
	return _SupportFundContract.Contract.Owner(&_SupportFundContract.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SupportFundContract *SupportFundContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SupportFundContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SupportFundContract *SupportFundContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SupportFundContract.Contract.TransferOwnership(&_SupportFundContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SupportFundContract *SupportFundContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SupportFundContract.Contract.TransferOwnership(&_SupportFundContract.TransactOpts, newOwner)
}

// UpdateFundItemsForDay is a paid mutator transaction binding the contract method 0xcfb2749d.
//
// Solidity: function updateFundItemsForDay(string _dayKey, (string,uint256,string,uint256,string)[] _newItems) returns()
func (_SupportFundContract *SupportFundContractTransactor) UpdateFundItemsForDay(opts *bind.TransactOpts, _dayKey string, _newItems []DataStructsFundItem) (*types.Transaction, error) {
	return _SupportFundContract.contract.Transact(opts, "updateFundItemsForDay", _dayKey, _newItems)
}

// UpdateFundItemsForDay is a paid mutator transaction binding the contract method 0xcfb2749d.
//
// Solidity: function updateFundItemsForDay(string _dayKey, (string,uint256,string,uint256,string)[] _newItems) returns()
func (_SupportFundContract *SupportFundContractSession) UpdateFundItemsForDay(_dayKey string, _newItems []DataStructsFundItem) (*types.Transaction, error) {
	return _SupportFundContract.Contract.UpdateFundItemsForDay(&_SupportFundContract.TransactOpts, _dayKey, _newItems)
}

// UpdateFundItemsForDay is a paid mutator transaction binding the contract method 0xcfb2749d.
//
// Solidity: function updateFundItemsForDay(string _dayKey, (string,uint256,string,uint256,string)[] _newItems) returns()
func (_SupportFundContract *SupportFundContractTransactorSession) UpdateFundItemsForDay(_dayKey string, _newItems []DataStructsFundItem) (*types.Transaction, error) {
	return _SupportFundContract.Contract.UpdateFundItemsForDay(&_SupportFundContract.TransactOpts, _dayKey, _newItems)
}

// SupportFundContractFundItemsUpdatedIterator is returned from FilterFundItemsUpdated and is used to iterate over the raw logs and unpacked data for FundItemsUpdated events raised by the SupportFundContract contract.
type SupportFundContractFundItemsUpdatedIterator struct {
	Event *SupportFundContractFundItemsUpdated // Event containing the contract specifics and raw log

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
func (it *SupportFundContractFundItemsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupportFundContractFundItemsUpdated)
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
		it.Event = new(SupportFundContractFundItemsUpdated)
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
func (it *SupportFundContractFundItemsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupportFundContractFundItemsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupportFundContractFundItemsUpdated represents a FundItemsUpdated event raised by the SupportFundContract contract.
type SupportFundContractFundItemsUpdated struct {
	DayKey   string
	NewItems []DataStructsFundItem
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFundItemsUpdated is a free log retrieval operation binding the contract event 0xc2449eda7f6620ae8ff0183b4dab7842c5427052d7916956e7e9b252f7d51e06.
//
// Solidity: event FundItemsUpdated(string dayKey, (string,uint256,string,uint256,string)[] newItems)
func (_SupportFundContract *SupportFundContractFilterer) FilterFundItemsUpdated(opts *bind.FilterOpts) (*SupportFundContractFundItemsUpdatedIterator, error) {

	logs, sub, err := _SupportFundContract.contract.FilterLogs(opts, "FundItemsUpdated")
	if err != nil {
		return nil, err
	}
	return &SupportFundContractFundItemsUpdatedIterator{contract: _SupportFundContract.contract, event: "FundItemsUpdated", logs: logs, sub: sub}, nil
}

// WatchFundItemsUpdated is a free log subscription operation binding the contract event 0xc2449eda7f6620ae8ff0183b4dab7842c5427052d7916956e7e9b252f7d51e06.
//
// Solidity: event FundItemsUpdated(string dayKey, (string,uint256,string,uint256,string)[] newItems)
func (_SupportFundContract *SupportFundContractFilterer) WatchFundItemsUpdated(opts *bind.WatchOpts, sink chan<- *SupportFundContractFundItemsUpdated) (event.Subscription, error) {

	logs, sub, err := _SupportFundContract.contract.WatchLogs(opts, "FundItemsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupportFundContractFundItemsUpdated)
				if err := _SupportFundContract.contract.UnpackLog(event, "FundItemsUpdated", log); err != nil {
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

// ParseFundItemsUpdated is a log parse operation binding the contract event 0xc2449eda7f6620ae8ff0183b4dab7842c5427052d7916956e7e9b252f7d51e06.
//
// Solidity: event FundItemsUpdated(string dayKey, (string,uint256,string,uint256,string)[] newItems)
func (_SupportFundContract *SupportFundContractFilterer) ParseFundItemsUpdated(log types.Log) (*SupportFundContractFundItemsUpdated, error) {
	event := new(SupportFundContractFundItemsUpdated)
	if err := _SupportFundContract.contract.UnpackLog(event, "FundItemsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupportFundContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SupportFundContract contract.
type SupportFundContractOwnershipTransferredIterator struct {
	Event *SupportFundContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SupportFundContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupportFundContractOwnershipTransferred)
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
		it.Event = new(SupportFundContractOwnershipTransferred)
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
func (it *SupportFundContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupportFundContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupportFundContractOwnershipTransferred represents a OwnershipTransferred event raised by the SupportFundContract contract.
type SupportFundContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SupportFundContract *SupportFundContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SupportFundContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SupportFundContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SupportFundContractOwnershipTransferredIterator{contract: _SupportFundContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SupportFundContract *SupportFundContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SupportFundContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SupportFundContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupportFundContractOwnershipTransferred)
				if err := _SupportFundContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SupportFundContract *SupportFundContractFilterer) ParseOwnershipTransferred(log types.Log) (*SupportFundContractOwnershipTransferred, error) {
	event := new(SupportFundContractOwnershipTransferred)
	if err := _SupportFundContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
