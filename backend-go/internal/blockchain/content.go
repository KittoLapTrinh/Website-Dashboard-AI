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

// DataStructsContentItem is an auto generated low-level Go binding around an user-defined struct.
type DataStructsContentItem struct {
	Id          *big.Int
	ContentType uint8
	Title       string
	Description string
	ImageUrl    string
	Author      string
	Timestamp   *big.Int
}

// ContentContractMetaData contains all meta data concerning the ContentContract contract.
var ContentContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumDataStructs.ContentType\",\"name\":\"contentType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"author\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structDataStructs.ContentItem\",\"name\":\"newContent\",\"type\":\"tuple\"}],\"name\":\"ContentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"enumDataStructs.ContentType\",\"name\":\"_contentType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_imageUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_author\",\"type\":\"string\"}],\"name\":\"addContent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allContent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumDataStructs.ContentType\",\"name\":\"contentType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"author\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"getRecentContent\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumDataStructs.ContentType\",\"name\":\"contentType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"author\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructs.ContentItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContentContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContentContractMetaData.ABI instead.
var ContentContractABI = ContentContractMetaData.ABI

// ContentContract is an auto generated Go binding around an Ethereum contract.
type ContentContract struct {
	ContentContractCaller     // Read-only binding to the contract
	ContentContractTransactor // Write-only binding to the contract
	ContentContractFilterer   // Log filterer for contract events
}

// ContentContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContentContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContentContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContentContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContentContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContentContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContentContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContentContractSession struct {
	Contract     *ContentContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContentContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContentContractCallerSession struct {
	Contract *ContentContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ContentContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContentContractTransactorSession struct {
	Contract     *ContentContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContentContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContentContractRaw struct {
	Contract *ContentContract // Generic contract binding to access the raw methods on
}

// ContentContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContentContractCallerRaw struct {
	Contract *ContentContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContentContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContentContractTransactorRaw struct {
	Contract *ContentContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContentContract creates a new instance of ContentContract, bound to a specific deployed contract.
func NewContentContract(address common.Address, backend bind.ContractBackend) (*ContentContract, error) {
	contract, err := bindContentContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContentContract{ContentContractCaller: ContentContractCaller{contract: contract}, ContentContractTransactor: ContentContractTransactor{contract: contract}, ContentContractFilterer: ContentContractFilterer{contract: contract}}, nil
}

// NewContentContractCaller creates a new read-only instance of ContentContract, bound to a specific deployed contract.
func NewContentContractCaller(address common.Address, caller bind.ContractCaller) (*ContentContractCaller, error) {
	contract, err := bindContentContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContentContractCaller{contract: contract}, nil
}

// NewContentContractTransactor creates a new write-only instance of ContentContract, bound to a specific deployed contract.
func NewContentContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContentContractTransactor, error) {
	contract, err := bindContentContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContentContractTransactor{contract: contract}, nil
}

// NewContentContractFilterer creates a new log filterer instance of ContentContract, bound to a specific deployed contract.
func NewContentContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContentContractFilterer, error) {
	contract, err := bindContentContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContentContractFilterer{contract: contract}, nil
}

// bindContentContract binds a generic wrapper to an already deployed contract.
func bindContentContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContentContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContentContract *ContentContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContentContract.Contract.ContentContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContentContract *ContentContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContentContract.Contract.ContentContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContentContract *ContentContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContentContract.Contract.ContentContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContentContract *ContentContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContentContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContentContract *ContentContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContentContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContentContract *ContentContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContentContract.Contract.contract.Transact(opts, method, params...)
}

// AllContent is a free data retrieval call binding the contract method 0x07d1b2fd.
//
// Solidity: function allContent(uint256 ) view returns(uint256 id, uint8 contentType, string title, string description, string imageUrl, string author, uint256 timestamp)
func (_ContentContract *ContentContractCaller) AllContent(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id          *big.Int
	ContentType uint8
	Title       string
	Description string
	ImageUrl    string
	Author      string
	Timestamp   *big.Int
}, error) {
	var out []interface{}
	err := _ContentContract.contract.Call(opts, &out, "allContent", arg0)

	outstruct := new(struct {
		Id          *big.Int
		ContentType uint8
		Title       string
		Description string
		ImageUrl    string
		Author      string
		Timestamp   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ContentType = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Title = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.ImageUrl = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Author = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AllContent is a free data retrieval call binding the contract method 0x07d1b2fd.
//
// Solidity: function allContent(uint256 ) view returns(uint256 id, uint8 contentType, string title, string description, string imageUrl, string author, uint256 timestamp)
func (_ContentContract *ContentContractSession) AllContent(arg0 *big.Int) (struct {
	Id          *big.Int
	ContentType uint8
	Title       string
	Description string
	ImageUrl    string
	Author      string
	Timestamp   *big.Int
}, error) {
	return _ContentContract.Contract.AllContent(&_ContentContract.CallOpts, arg0)
}

// AllContent is a free data retrieval call binding the contract method 0x07d1b2fd.
//
// Solidity: function allContent(uint256 ) view returns(uint256 id, uint8 contentType, string title, string description, string imageUrl, string author, uint256 timestamp)
func (_ContentContract *ContentContractCallerSession) AllContent(arg0 *big.Int) (struct {
	Id          *big.Int
	ContentType uint8
	Title       string
	Description string
	ImageUrl    string
	Author      string
	Timestamp   *big.Int
}, error) {
	return _ContentContract.Contract.AllContent(&_ContentContract.CallOpts, arg0)
}

// GetRecentContent is a free data retrieval call binding the contract method 0xff88834b.
//
// Solidity: function getRecentContent(uint256 _count) view returns((uint256,uint8,string,string,string,string,uint256)[])
func (_ContentContract *ContentContractCaller) GetRecentContent(opts *bind.CallOpts, _count *big.Int) ([]DataStructsContentItem, error) {
	var out []interface{}
	err := _ContentContract.contract.Call(opts, &out, "getRecentContent", _count)

	if err != nil {
		return *new([]DataStructsContentItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataStructsContentItem)).(*[]DataStructsContentItem)

	return out0, err

}

// GetRecentContent is a free data retrieval call binding the contract method 0xff88834b.
//
// Solidity: function getRecentContent(uint256 _count) view returns((uint256,uint8,string,string,string,string,uint256)[])
func (_ContentContract *ContentContractSession) GetRecentContent(_count *big.Int) ([]DataStructsContentItem, error) {
	return _ContentContract.Contract.GetRecentContent(&_ContentContract.CallOpts, _count)
}

// GetRecentContent is a free data retrieval call binding the contract method 0xff88834b.
//
// Solidity: function getRecentContent(uint256 _count) view returns((uint256,uint8,string,string,string,string,uint256)[])
func (_ContentContract *ContentContractCallerSession) GetRecentContent(_count *big.Int) ([]DataStructsContentItem, error) {
	return _ContentContract.Contract.GetRecentContent(&_ContentContract.CallOpts, _count)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContentContract *ContentContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContentContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContentContract *ContentContractSession) Owner() (common.Address, error) {
	return _ContentContract.Contract.Owner(&_ContentContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContentContract *ContentContractCallerSession) Owner() (common.Address, error) {
	return _ContentContract.Contract.Owner(&_ContentContract.CallOpts)
}

// AddContent is a paid mutator transaction binding the contract method 0x4b525d46.
//
// Solidity: function addContent(uint8 _contentType, string _title, string _description, string _imageUrl, string _author) returns()
func (_ContentContract *ContentContractTransactor) AddContent(opts *bind.TransactOpts, _contentType uint8, _title string, _description string, _imageUrl string, _author string) (*types.Transaction, error) {
	return _ContentContract.contract.Transact(opts, "addContent", _contentType, _title, _description, _imageUrl, _author)
}

// AddContent is a paid mutator transaction binding the contract method 0x4b525d46.
//
// Solidity: function addContent(uint8 _contentType, string _title, string _description, string _imageUrl, string _author) returns()
func (_ContentContract *ContentContractSession) AddContent(_contentType uint8, _title string, _description string, _imageUrl string, _author string) (*types.Transaction, error) {
	return _ContentContract.Contract.AddContent(&_ContentContract.TransactOpts, _contentType, _title, _description, _imageUrl, _author)
}

// AddContent is a paid mutator transaction binding the contract method 0x4b525d46.
//
// Solidity: function addContent(uint8 _contentType, string _title, string _description, string _imageUrl, string _author) returns()
func (_ContentContract *ContentContractTransactorSession) AddContent(_contentType uint8, _title string, _description string, _imageUrl string, _author string) (*types.Transaction, error) {
	return _ContentContract.Contract.AddContent(&_ContentContract.TransactOpts, _contentType, _title, _description, _imageUrl, _author)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContentContract *ContentContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContentContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContentContract *ContentContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContentContract.Contract.TransferOwnership(&_ContentContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContentContract *ContentContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContentContract.Contract.TransferOwnership(&_ContentContract.TransactOpts, newOwner)
}

// ContentContractContentAddedIterator is returned from FilterContentAdded and is used to iterate over the raw logs and unpacked data for ContentAdded events raised by the ContentContract contract.
type ContentContractContentAddedIterator struct {
	Event *ContentContractContentAdded // Event containing the contract specifics and raw log

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
func (it *ContentContractContentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentContractContentAdded)
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
		it.Event = new(ContentContractContentAdded)
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
func (it *ContentContractContentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentContractContentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentContractContentAdded represents a ContentAdded event raised by the ContentContract contract.
type ContentContractContentAdded struct {
	NewContent DataStructsContentItem
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterContentAdded is a free log retrieval operation binding the contract event 0xf6f6b02c8aafe26c9c46c27302705d772ddcb0f7e4698765be0c3c4de35619f2.
//
// Solidity: event ContentAdded((uint256,uint8,string,string,string,string,uint256) newContent)
func (_ContentContract *ContentContractFilterer) FilterContentAdded(opts *bind.FilterOpts) (*ContentContractContentAddedIterator, error) {

	logs, sub, err := _ContentContract.contract.FilterLogs(opts, "ContentAdded")
	if err != nil {
		return nil, err
	}
	return &ContentContractContentAddedIterator{contract: _ContentContract.contract, event: "ContentAdded", logs: logs, sub: sub}, nil
}

// WatchContentAdded is a free log subscription operation binding the contract event 0xf6f6b02c8aafe26c9c46c27302705d772ddcb0f7e4698765be0c3c4de35619f2.
//
// Solidity: event ContentAdded((uint256,uint8,string,string,string,string,uint256) newContent)
func (_ContentContract *ContentContractFilterer) WatchContentAdded(opts *bind.WatchOpts, sink chan<- *ContentContractContentAdded) (event.Subscription, error) {

	logs, sub, err := _ContentContract.contract.WatchLogs(opts, "ContentAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentContractContentAdded)
				if err := _ContentContract.contract.UnpackLog(event, "ContentAdded", log); err != nil {
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

// ParseContentAdded is a log parse operation binding the contract event 0xf6f6b02c8aafe26c9c46c27302705d772ddcb0f7e4698765be0c3c4de35619f2.
//
// Solidity: event ContentAdded((uint256,uint8,string,string,string,string,uint256) newContent)
func (_ContentContract *ContentContractFilterer) ParseContentAdded(log types.Log) (*ContentContractContentAdded, error) {
	event := new(ContentContractContentAdded)
	if err := _ContentContract.contract.UnpackLog(event, "ContentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContentContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContentContract contract.
type ContentContractOwnershipTransferredIterator struct {
	Event *ContentContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContentContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentContractOwnershipTransferred)
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
		it.Event = new(ContentContractOwnershipTransferred)
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
func (it *ContentContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentContractOwnershipTransferred represents a OwnershipTransferred event raised by the ContentContract contract.
type ContentContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContentContract *ContentContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContentContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContentContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContentContractOwnershipTransferredIterator{contract: _ContentContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContentContract *ContentContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContentContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContentContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentContractOwnershipTransferred)
				if err := _ContentContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContentContract *ContentContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContentContractOwnershipTransferred, error) {
	event := new(ContentContractOwnershipTransferred)
	if err := _ContentContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
