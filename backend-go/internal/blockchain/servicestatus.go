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

// DataStructsService is an auto generated low-level Go binding around an user-defined struct.
type DataStructsService struct {
	Name     string
	Title    string
	Subtitle string
	Status   uint8
	IconId   string
}

// ServiceStatusContractMetaData contains all meta data concerning the ServiceStatusContract contract.
var ServiceStatusContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumDataStructs.ServiceStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"ServiceStatusUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"enumDataStructs.ServiceStatus\",\"name\":\"_newStatus\",\"type\":\"uint8\"}],\"name\":\"updateServiceStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllServices\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subtitle\",\"type\":\"string\"},{\"internalType\":\"enumDataStructs.ServiceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"}],\"internalType\":\"structDataStructs.Service[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"services\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subtitle\",\"type\":\"string\"},{\"internalType\":\"enumDataStructs.ServiceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ServiceStatusContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ServiceStatusContractMetaData.ABI instead.
var ServiceStatusContractABI = ServiceStatusContractMetaData.ABI

// ServiceStatusContract is an auto generated Go binding around an Ethereum contract.
type ServiceStatusContract struct {
	ServiceStatusContractCaller     // Read-only binding to the contract
	ServiceStatusContractTransactor // Write-only binding to the contract
	ServiceStatusContractFilterer   // Log filterer for contract events
}

// ServiceStatusContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ServiceStatusContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceStatusContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ServiceStatusContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceStatusContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ServiceStatusContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceStatusContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ServiceStatusContractSession struct {
	Contract     *ServiceStatusContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ServiceStatusContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ServiceStatusContractCallerSession struct {
	Contract *ServiceStatusContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ServiceStatusContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ServiceStatusContractTransactorSession struct {
	Contract     *ServiceStatusContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ServiceStatusContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ServiceStatusContractRaw struct {
	Contract *ServiceStatusContract // Generic contract binding to access the raw methods on
}

// ServiceStatusContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ServiceStatusContractCallerRaw struct {
	Contract *ServiceStatusContractCaller // Generic read-only contract binding to access the raw methods on
}

// ServiceStatusContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ServiceStatusContractTransactorRaw struct {
	Contract *ServiceStatusContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewServiceStatusContract creates a new instance of ServiceStatusContract, bound to a specific deployed contract.
func NewServiceStatusContract(address common.Address, backend bind.ContractBackend) (*ServiceStatusContract, error) {
	contract, err := bindServiceStatusContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ServiceStatusContract{ServiceStatusContractCaller: ServiceStatusContractCaller{contract: contract}, ServiceStatusContractTransactor: ServiceStatusContractTransactor{contract: contract}, ServiceStatusContractFilterer: ServiceStatusContractFilterer{contract: contract}}, nil
}

// NewServiceStatusContractCaller creates a new read-only instance of ServiceStatusContract, bound to a specific deployed contract.
func NewServiceStatusContractCaller(address common.Address, caller bind.ContractCaller) (*ServiceStatusContractCaller, error) {
	contract, err := bindServiceStatusContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceStatusContractCaller{contract: contract}, nil
}

// NewServiceStatusContractTransactor creates a new write-only instance of ServiceStatusContract, bound to a specific deployed contract.
func NewServiceStatusContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ServiceStatusContractTransactor, error) {
	contract, err := bindServiceStatusContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceStatusContractTransactor{contract: contract}, nil
}

// NewServiceStatusContractFilterer creates a new log filterer instance of ServiceStatusContract, bound to a specific deployed contract.
func NewServiceStatusContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ServiceStatusContractFilterer, error) {
	contract, err := bindServiceStatusContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ServiceStatusContractFilterer{contract: contract}, nil
}

// bindServiceStatusContract binds a generic wrapper to an already deployed contract.
func bindServiceStatusContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ServiceStatusContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServiceStatusContract *ServiceStatusContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ServiceStatusContract.Contract.ServiceStatusContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServiceStatusContract *ServiceStatusContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceStatusContract.Contract.ServiceStatusContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServiceStatusContract *ServiceStatusContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServiceStatusContract.Contract.ServiceStatusContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServiceStatusContract *ServiceStatusContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ServiceStatusContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServiceStatusContract *ServiceStatusContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceStatusContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServiceStatusContract *ServiceStatusContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServiceStatusContract.Contract.contract.Transact(opts, method, params...)
}

// GetAllServices is a free data retrieval call binding the contract method 0x21fe0f30.
//
// Solidity: function getAllServices() view returns((string,string,string,uint8,string)[])
func (_ServiceStatusContract *ServiceStatusContractCaller) GetAllServices(opts *bind.CallOpts) ([]DataStructsService, error) {
	var out []interface{}
	err := _ServiceStatusContract.contract.Call(opts, &out, "getAllServices")

	if err != nil {
		return *new([]DataStructsService), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataStructsService)).(*[]DataStructsService)

	return out0, err

}

// GetAllServices is a free data retrieval call binding the contract method 0x21fe0f30.
//
// Solidity: function getAllServices() view returns((string,string,string,uint8,string)[])
func (_ServiceStatusContract *ServiceStatusContractSession) GetAllServices() ([]DataStructsService, error) {
	return _ServiceStatusContract.Contract.GetAllServices(&_ServiceStatusContract.CallOpts)
}

// GetAllServices is a free data retrieval call binding the contract method 0x21fe0f30.
//
// Solidity: function getAllServices() view returns((string,string,string,uint8,string)[])
func (_ServiceStatusContract *ServiceStatusContractCallerSession) GetAllServices() ([]DataStructsService, error) {
	return _ServiceStatusContract.Contract.GetAllServices(&_ServiceStatusContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ServiceStatusContract *ServiceStatusContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ServiceStatusContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ServiceStatusContract *ServiceStatusContractSession) Owner() (common.Address, error) {
	return _ServiceStatusContract.Contract.Owner(&_ServiceStatusContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ServiceStatusContract *ServiceStatusContractCallerSession) Owner() (common.Address, error) {
	return _ServiceStatusContract.Contract.Owner(&_ServiceStatusContract.CallOpts)
}

// Services is a free data retrieval call binding the contract method 0xc22c4f43.
//
// Solidity: function services(uint256 ) view returns(string name, string title, string subtitle, uint8 status, string iconId)
func (_ServiceStatusContract *ServiceStatusContractCaller) Services(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name     string
	Title    string
	Subtitle string
	Status   uint8
	IconId   string
}, error) {
	var out []interface{}
	err := _ServiceStatusContract.contract.Call(opts, &out, "services", arg0)

	outstruct := new(struct {
		Name     string
		Title    string
		Subtitle string
		Status   uint8
		IconId   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Title = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Subtitle = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.IconId = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// Services is a free data retrieval call binding the contract method 0xc22c4f43.
//
// Solidity: function services(uint256 ) view returns(string name, string title, string subtitle, uint8 status, string iconId)
func (_ServiceStatusContract *ServiceStatusContractSession) Services(arg0 *big.Int) (struct {
	Name     string
	Title    string
	Subtitle string
	Status   uint8
	IconId   string
}, error) {
	return _ServiceStatusContract.Contract.Services(&_ServiceStatusContract.CallOpts, arg0)
}

// Services is a free data retrieval call binding the contract method 0xc22c4f43.
//
// Solidity: function services(uint256 ) view returns(string name, string title, string subtitle, uint8 status, string iconId)
func (_ServiceStatusContract *ServiceStatusContractCallerSession) Services(arg0 *big.Int) (struct {
	Name     string
	Title    string
	Subtitle string
	Status   uint8
	IconId   string
}, error) {
	return _ServiceStatusContract.Contract.Services(&_ServiceStatusContract.CallOpts, arg0)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ServiceStatusContract *ServiceStatusContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ServiceStatusContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ServiceStatusContract *ServiceStatusContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ServiceStatusContract.Contract.TransferOwnership(&_ServiceStatusContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ServiceStatusContract *ServiceStatusContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ServiceStatusContract.Contract.TransferOwnership(&_ServiceStatusContract.TransactOpts, newOwner)
}

// UpdateServiceStatus is a paid mutator transaction binding the contract method 0x60423fa2.
//
// Solidity: function updateServiceStatus(uint256 _index, uint8 _newStatus) returns()
func (_ServiceStatusContract *ServiceStatusContractTransactor) UpdateServiceStatus(opts *bind.TransactOpts, _index *big.Int, _newStatus uint8) (*types.Transaction, error) {
	return _ServiceStatusContract.contract.Transact(opts, "updateServiceStatus", _index, _newStatus)
}

// UpdateServiceStatus is a paid mutator transaction binding the contract method 0x60423fa2.
//
// Solidity: function updateServiceStatus(uint256 _index, uint8 _newStatus) returns()
func (_ServiceStatusContract *ServiceStatusContractSession) UpdateServiceStatus(_index *big.Int, _newStatus uint8) (*types.Transaction, error) {
	return _ServiceStatusContract.Contract.UpdateServiceStatus(&_ServiceStatusContract.TransactOpts, _index, _newStatus)
}

// UpdateServiceStatus is a paid mutator transaction binding the contract method 0x60423fa2.
//
// Solidity: function updateServiceStatus(uint256 _index, uint8 _newStatus) returns()
func (_ServiceStatusContract *ServiceStatusContractTransactorSession) UpdateServiceStatus(_index *big.Int, _newStatus uint8) (*types.Transaction, error) {
	return _ServiceStatusContract.Contract.UpdateServiceStatus(&_ServiceStatusContract.TransactOpts, _index, _newStatus)
}

// ServiceStatusContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ServiceStatusContract contract.
type ServiceStatusContractOwnershipTransferredIterator struct {
	Event *ServiceStatusContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ServiceStatusContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceStatusContractOwnershipTransferred)
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
		it.Event = new(ServiceStatusContractOwnershipTransferred)
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
func (it *ServiceStatusContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceStatusContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceStatusContractOwnershipTransferred represents a OwnershipTransferred event raised by the ServiceStatusContract contract.
type ServiceStatusContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ServiceStatusContract *ServiceStatusContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ServiceStatusContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ServiceStatusContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ServiceStatusContractOwnershipTransferredIterator{contract: _ServiceStatusContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ServiceStatusContract *ServiceStatusContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ServiceStatusContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ServiceStatusContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceStatusContractOwnershipTransferred)
				if err := _ServiceStatusContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ServiceStatusContract *ServiceStatusContractFilterer) ParseOwnershipTransferred(log types.Log) (*ServiceStatusContractOwnershipTransferred, error) {
	event := new(ServiceStatusContractOwnershipTransferred)
	if err := _ServiceStatusContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceStatusContractServiceStatusUpdatedIterator is returned from FilterServiceStatusUpdated and is used to iterate over the raw logs and unpacked data for ServiceStatusUpdated events raised by the ServiceStatusContract contract.
type ServiceStatusContractServiceStatusUpdatedIterator struct {
	Event *ServiceStatusContractServiceStatusUpdated // Event containing the contract specifics and raw log

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
func (it *ServiceStatusContractServiceStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceStatusContractServiceStatusUpdated)
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
		it.Event = new(ServiceStatusContractServiceStatusUpdated)
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
func (it *ServiceStatusContractServiceStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceStatusContractServiceStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceStatusContractServiceStatusUpdated represents a ServiceStatusUpdated event raised by the ServiceStatusContract contract.
type ServiceStatusContractServiceStatusUpdated struct {
	Index     *big.Int
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterServiceStatusUpdated is a free log retrieval operation binding the contract event 0x5788bd9e26a75de76695c2a82e01198faadcc1c36aa422746f16195be540853a.
//
// Solidity: event ServiceStatusUpdated(uint256 index, uint8 newStatus)
func (_ServiceStatusContract *ServiceStatusContractFilterer) FilterServiceStatusUpdated(opts *bind.FilterOpts) (*ServiceStatusContractServiceStatusUpdatedIterator, error) {

	logs, sub, err := _ServiceStatusContract.contract.FilterLogs(opts, "ServiceStatusUpdated")
	if err != nil {
		return nil, err
	}
	return &ServiceStatusContractServiceStatusUpdatedIterator{contract: _ServiceStatusContract.contract, event: "ServiceStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchServiceStatusUpdated is a free log subscription operation binding the contract event 0x5788bd9e26a75de76695c2a82e01198faadcc1c36aa422746f16195be540853a.
//
// Solidity: event ServiceStatusUpdated(uint256 index, uint8 newStatus)
func (_ServiceStatusContract *ServiceStatusContractFilterer) WatchServiceStatusUpdated(opts *bind.WatchOpts, sink chan<- *ServiceStatusContractServiceStatusUpdated) (event.Subscription, error) {

	logs, sub, err := _ServiceStatusContract.contract.WatchLogs(opts, "ServiceStatusUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceStatusContractServiceStatusUpdated)
				if err := _ServiceStatusContract.contract.UnpackLog(event, "ServiceStatusUpdated", log); err != nil {
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

// ParseServiceStatusUpdated is a log parse operation binding the contract event 0x5788bd9e26a75de76695c2a82e01198faadcc1c36aa422746f16195be540853a.
//
// Solidity: event ServiceStatusUpdated(uint256 index, uint8 newStatus)
func (_ServiceStatusContract *ServiceStatusContractFilterer) ParseServiceStatusUpdated(log types.Log) (*ServiceStatusContractServiceStatusUpdated, error) {
	event := new(ServiceStatusContractServiceStatusUpdated)
	if err := _ServiceStatusContract.contract.UnpackLog(event, "ServiceStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
