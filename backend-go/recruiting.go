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

// DataStructsJob is an auto generated low-level Go binding around an user-defined struct.
type DataStructsJob struct {
	Foundation string
	Position   string
	Field      string
	Salary     *big.Int
	Form       string
	Trend      uint8
	IconId     string
}

// RecruitingContractMetaData contains all meta data concerning the RecruitingContract contract.
var RecruitingContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"foundation\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"salary\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"form\",\"type\":\"string\"},{\"internalType\":\"enumDataStructs.JobTrend\",\"name\":\"trend\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"}],\"internalType\":\"structDataStructs.Job\",\"name\":\"_newJob\",\"type\":\"tuple\"}],\"name\":\"addJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"foundation\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"salary\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"form\",\"type\":\"string\"},{\"internalType\":\"enumDataStructs.JobTrend\",\"name\":\"trend\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structDataStructs.Job\",\"name\":\"newJob\",\"type\":\"tuple\"}],\"name\":\"JobAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllJobs\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"foundation\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"salary\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"form\",\"type\":\"string\"},{\"internalType\":\"enumDataStructs.JobTrend\",\"name\":\"trend\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"}],\"internalType\":\"structDataStructs.Job[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"foundation\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"salary\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"form\",\"type\":\"string\"},{\"internalType\":\"enumDataStructs.JobTrend\",\"name\":\"trend\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RecruitingContractABI is the input ABI used to generate the binding from.
// Deprecated: Use RecruitingContractMetaData.ABI instead.
var RecruitingContractABI = RecruitingContractMetaData.ABI

// RecruitingContract is an auto generated Go binding around an Ethereum contract.
type RecruitingContract struct {
	RecruitingContractCaller     // Read-only binding to the contract
	RecruitingContractTransactor // Write-only binding to the contract
	RecruitingContractFilterer   // Log filterer for contract events
}

// RecruitingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecruitingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecruitingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecruitingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecruitingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecruitingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecruitingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecruitingContractSession struct {
	Contract     *RecruitingContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RecruitingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecruitingContractCallerSession struct {
	Contract *RecruitingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// RecruitingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecruitingContractTransactorSession struct {
	Contract     *RecruitingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// RecruitingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecruitingContractRaw struct {
	Contract *RecruitingContract // Generic contract binding to access the raw methods on
}

// RecruitingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecruitingContractCallerRaw struct {
	Contract *RecruitingContractCaller // Generic read-only contract binding to access the raw methods on
}

// RecruitingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecruitingContractTransactorRaw struct {
	Contract *RecruitingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecruitingContract creates a new instance of RecruitingContract, bound to a specific deployed contract.
func NewRecruitingContract(address common.Address, backend bind.ContractBackend) (*RecruitingContract, error) {
	contract, err := bindRecruitingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RecruitingContract{RecruitingContractCaller: RecruitingContractCaller{contract: contract}, RecruitingContractTransactor: RecruitingContractTransactor{contract: contract}, RecruitingContractFilterer: RecruitingContractFilterer{contract: contract}}, nil
}

// NewRecruitingContractCaller creates a new read-only instance of RecruitingContract, bound to a specific deployed contract.
func NewRecruitingContractCaller(address common.Address, caller bind.ContractCaller) (*RecruitingContractCaller, error) {
	contract, err := bindRecruitingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecruitingContractCaller{contract: contract}, nil
}

// NewRecruitingContractTransactor creates a new write-only instance of RecruitingContract, bound to a specific deployed contract.
func NewRecruitingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RecruitingContractTransactor, error) {
	contract, err := bindRecruitingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecruitingContractTransactor{contract: contract}, nil
}

// NewRecruitingContractFilterer creates a new log filterer instance of RecruitingContract, bound to a specific deployed contract.
func NewRecruitingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RecruitingContractFilterer, error) {
	contract, err := bindRecruitingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecruitingContractFilterer{contract: contract}, nil
}

// bindRecruitingContract binds a generic wrapper to an already deployed contract.
func bindRecruitingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RecruitingContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RecruitingContract *RecruitingContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecruitingContract.Contract.RecruitingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RecruitingContract *RecruitingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecruitingContract.Contract.RecruitingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RecruitingContract *RecruitingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecruitingContract.Contract.RecruitingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RecruitingContract *RecruitingContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecruitingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RecruitingContract *RecruitingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecruitingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RecruitingContract *RecruitingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecruitingContract.Contract.contract.Transact(opts, method, params...)
}

// GetAllJobs is a free data retrieval call binding the contract method 0x965a79fb.
//
// Solidity: function getAllJobs() view returns((string,string,string,uint256,string,uint8,string)[])
func (_RecruitingContract *RecruitingContractCaller) GetAllJobs(opts *bind.CallOpts) ([]DataStructsJob, error) {
	var out []interface{}
	err := _RecruitingContract.contract.Call(opts, &out, "getAllJobs")

	if err != nil {
		return *new([]DataStructsJob), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataStructsJob)).(*[]DataStructsJob)

	return out0, err

}

// GetAllJobs is a free data retrieval call binding the contract method 0x965a79fb.
//
// Solidity: function getAllJobs() view returns((string,string,string,uint256,string,uint8,string)[])
func (_RecruitingContract *RecruitingContractSession) GetAllJobs() ([]DataStructsJob, error) {
	return _RecruitingContract.Contract.GetAllJobs(&_RecruitingContract.CallOpts)
}

// GetAllJobs is a free data retrieval call binding the contract method 0x965a79fb.
//
// Solidity: function getAllJobs() view returns((string,string,string,uint256,string,uint8,string)[])
func (_RecruitingContract *RecruitingContractCallerSession) GetAllJobs() ([]DataStructsJob, error) {
	return _RecruitingContract.Contract.GetAllJobs(&_RecruitingContract.CallOpts)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(string foundation, string position, string field, uint256 salary, string form, uint8 trend, string iconId)
func (_RecruitingContract *RecruitingContractCaller) Jobs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Foundation string
	Position   string
	Field      string
	Salary     *big.Int
	Form       string
	Trend      uint8
	IconId     string
}, error) {
	var out []interface{}
	err := _RecruitingContract.contract.Call(opts, &out, "jobs", arg0)

	outstruct := new(struct {
		Foundation string
		Position   string
		Field      string
		Salary     *big.Int
		Form       string
		Trend      uint8
		IconId     string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Foundation = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Position = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Field = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Salary = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Form = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Trend = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.IconId = *abi.ConvertType(out[6], new(string)).(*string)

	return *outstruct, err

}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(string foundation, string position, string field, uint256 salary, string form, uint8 trend, string iconId)
func (_RecruitingContract *RecruitingContractSession) Jobs(arg0 *big.Int) (struct {
	Foundation string
	Position   string
	Field      string
	Salary     *big.Int
	Form       string
	Trend      uint8
	IconId     string
}, error) {
	return _RecruitingContract.Contract.Jobs(&_RecruitingContract.CallOpts, arg0)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(string foundation, string position, string field, uint256 salary, string form, uint8 trend, string iconId)
func (_RecruitingContract *RecruitingContractCallerSession) Jobs(arg0 *big.Int) (struct {
	Foundation string
	Position   string
	Field      string
	Salary     *big.Int
	Form       string
	Trend      uint8
	IconId     string
}, error) {
	return _RecruitingContract.Contract.Jobs(&_RecruitingContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RecruitingContract *RecruitingContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RecruitingContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RecruitingContract *RecruitingContractSession) Owner() (common.Address, error) {
	return _RecruitingContract.Contract.Owner(&_RecruitingContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RecruitingContract *RecruitingContractCallerSession) Owner() (common.Address, error) {
	return _RecruitingContract.Contract.Owner(&_RecruitingContract.CallOpts)
}

// AddJob is a paid mutator transaction binding the contract method 0x29b44e58.
//
// Solidity: function addJob((string,string,string,uint256,string,uint8,string) _newJob) returns()
func (_RecruitingContract *RecruitingContractTransactor) AddJob(opts *bind.TransactOpts, _newJob DataStructsJob) (*types.Transaction, error) {
	return _RecruitingContract.contract.Transact(opts, "addJob", _newJob)
}

// AddJob is a paid mutator transaction binding the contract method 0x29b44e58.
//
// Solidity: function addJob((string,string,string,uint256,string,uint8,string) _newJob) returns()
func (_RecruitingContract *RecruitingContractSession) AddJob(_newJob DataStructsJob) (*types.Transaction, error) {
	return _RecruitingContract.Contract.AddJob(&_RecruitingContract.TransactOpts, _newJob)
}

// AddJob is a paid mutator transaction binding the contract method 0x29b44e58.
//
// Solidity: function addJob((string,string,string,uint256,string,uint8,string) _newJob) returns()
func (_RecruitingContract *RecruitingContractTransactorSession) AddJob(_newJob DataStructsJob) (*types.Transaction, error) {
	return _RecruitingContract.Contract.AddJob(&_RecruitingContract.TransactOpts, _newJob)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RecruitingContract *RecruitingContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RecruitingContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RecruitingContract *RecruitingContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RecruitingContract.Contract.TransferOwnership(&_RecruitingContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RecruitingContract *RecruitingContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RecruitingContract.Contract.TransferOwnership(&_RecruitingContract.TransactOpts, newOwner)
}

// RecruitingContractJobAddedIterator is returned from FilterJobAdded and is used to iterate over the raw logs and unpacked data for JobAdded events raised by the RecruitingContract contract.
type RecruitingContractJobAddedIterator struct {
	Event *RecruitingContractJobAdded // Event containing the contract specifics and raw log

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
func (it *RecruitingContractJobAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecruitingContractJobAdded)
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
		it.Event = new(RecruitingContractJobAdded)
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
func (it *RecruitingContractJobAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecruitingContractJobAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecruitingContractJobAdded represents a JobAdded event raised by the RecruitingContract contract.
type RecruitingContractJobAdded struct {
	NewJob DataStructsJob
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterJobAdded is a free log retrieval operation binding the contract event 0xde0eae4446badd721c880f61c8e29849509fd7ced6259dcc27a46b2affa3329a.
//
// Solidity: event JobAdded((string,string,string,uint256,string,uint8,string) newJob)
func (_RecruitingContract *RecruitingContractFilterer) FilterJobAdded(opts *bind.FilterOpts) (*RecruitingContractJobAddedIterator, error) {

	logs, sub, err := _RecruitingContract.contract.FilterLogs(opts, "JobAdded")
	if err != nil {
		return nil, err
	}
	return &RecruitingContractJobAddedIterator{contract: _RecruitingContract.contract, event: "JobAdded", logs: logs, sub: sub}, nil
}

// WatchJobAdded is a free log subscription operation binding the contract event 0xde0eae4446badd721c880f61c8e29849509fd7ced6259dcc27a46b2affa3329a.
//
// Solidity: event JobAdded((string,string,string,uint256,string,uint8,string) newJob)
func (_RecruitingContract *RecruitingContractFilterer) WatchJobAdded(opts *bind.WatchOpts, sink chan<- *RecruitingContractJobAdded) (event.Subscription, error) {

	logs, sub, err := _RecruitingContract.contract.WatchLogs(opts, "JobAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecruitingContractJobAdded)
				if err := _RecruitingContract.contract.UnpackLog(event, "JobAdded", log); err != nil {
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

// ParseJobAdded is a log parse operation binding the contract event 0xde0eae4446badd721c880f61c8e29849509fd7ced6259dcc27a46b2affa3329a.
//
// Solidity: event JobAdded((string,string,string,uint256,string,uint8,string) newJob)
func (_RecruitingContract *RecruitingContractFilterer) ParseJobAdded(log types.Log) (*RecruitingContractJobAdded, error) {
	event := new(RecruitingContractJobAdded)
	if err := _RecruitingContract.contract.UnpackLog(event, "JobAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecruitingContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RecruitingContract contract.
type RecruitingContractOwnershipTransferredIterator struct {
	Event *RecruitingContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RecruitingContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecruitingContractOwnershipTransferred)
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
		it.Event = new(RecruitingContractOwnershipTransferred)
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
func (it *RecruitingContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecruitingContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecruitingContractOwnershipTransferred represents a OwnershipTransferred event raised by the RecruitingContract contract.
type RecruitingContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RecruitingContract *RecruitingContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RecruitingContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RecruitingContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RecruitingContractOwnershipTransferredIterator{contract: _RecruitingContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RecruitingContract *RecruitingContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RecruitingContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RecruitingContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecruitingContractOwnershipTransferred)
				if err := _RecruitingContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RecruitingContract *RecruitingContractFilterer) ParseOwnershipTransferred(log types.Log) (*RecruitingContractOwnershipTransferred, error) {
	event := new(RecruitingContractOwnershipTransferred)
	if err := _RecruitingContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
