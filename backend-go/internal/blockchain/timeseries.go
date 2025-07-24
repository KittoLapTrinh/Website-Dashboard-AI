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

// DataStructsDataPoint is an auto generated low-level Go binding around an user-defined struct.
type DataStructsDataPoint struct {
	Timestamp *big.Int
	Value     *big.Int
	Change    *big.Int
}

// BlockchainMetaData contains all meta data concerning the Blockchain contract.
var BlockchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"change\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structDataStructs.DataPoint\",\"name\":\"newPoint\",\"type\":\"tuple\"}],\"name\":\"DataPointAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSize\",\"type\":\"uint256\"}],\"name\":\"TimeSeriesReplaced\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getTimeSeriesData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"change\",\"type\":\"int256\"}],\"internalType\":\"structDataStructs.DataPoint[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getTimeSeriesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"getTimeSeriesSlice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"change\",\"type\":\"int256\"}],\"internalType\":\"structDataStructs.DataPoint[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_change\",\"type\":\"int256\"}],\"name\":\"pushDataPoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"change\",\"type\":\"int256\"}],\"internalType\":\"structDataStructs.DataPoint[]\",\"name\":\"_newData\",\"type\":\"tuple[]\"}],\"name\":\"replaceTimeSeries\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"timeSeriesData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"change\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BlockchainABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockchainMetaData.ABI instead.
var BlockchainABI = BlockchainMetaData.ABI

// Blockchain is an auto generated Go binding around an Ethereum contract.
type Blockchain struct {
	BlockchainCaller     // Read-only binding to the contract
	BlockchainTransactor // Write-only binding to the contract
	BlockchainFilterer   // Log filterer for contract events
}

// BlockchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockchainSession struct {
	Contract     *Blockchain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockchainCallerSession struct {
	Contract *BlockchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BlockchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockchainTransactorSession struct {
	Contract     *BlockchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BlockchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockchainRaw struct {
	Contract *Blockchain // Generic contract binding to access the raw methods on
}

// BlockchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockchainCallerRaw struct {
	Contract *BlockchainCaller // Generic read-only contract binding to access the raw methods on
}

// BlockchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockchainTransactorRaw struct {
	Contract *BlockchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockchain creates a new instance of Blockchain, bound to a specific deployed contract.
func NewBlockchain(address common.Address, backend bind.ContractBackend) (*Blockchain, error) {
	contract, err := bindBlockchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blockchain{BlockchainCaller: BlockchainCaller{contract: contract}, BlockchainTransactor: BlockchainTransactor{contract: contract}, BlockchainFilterer: BlockchainFilterer{contract: contract}}, nil
}

// NewBlockchainCaller creates a new read-only instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainCaller(address common.Address, caller bind.ContractCaller) (*BlockchainCaller, error) {
	contract, err := bindBlockchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainCaller{contract: contract}, nil
}

// NewBlockchainTransactor creates a new write-only instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockchainTransactor, error) {
	contract, err := bindBlockchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainTransactor{contract: contract}, nil
}

// NewBlockchainFilterer creates a new log filterer instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockchainFilterer, error) {
	contract, err := bindBlockchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockchainFilterer{contract: contract}, nil
}

// bindBlockchain binds a generic wrapper to an already deployed contract.
func bindBlockchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockchain *BlockchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockchain.Contract.BlockchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockchain *BlockchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.Contract.BlockchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockchain *BlockchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockchain.Contract.BlockchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockchain *BlockchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockchain *BlockchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockchain *BlockchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockchain.Contract.contract.Transact(opts, method, params...)
}

// GetTimeSeriesData is a free data retrieval call binding the contract method 0x7918f328.
//
// Solidity: function getTimeSeriesData(string _key) view returns((uint256,uint256,int256)[])
func (_Blockchain *BlockchainCaller) GetTimeSeriesData(opts *bind.CallOpts, _key string) ([]DataStructsDataPoint, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getTimeSeriesData", _key)

	if err != nil {
		return *new([]DataStructsDataPoint), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataStructsDataPoint)).(*[]DataStructsDataPoint)

	return out0, err

}

// GetTimeSeriesData is a free data retrieval call binding the contract method 0x7918f328.
//
// Solidity: function getTimeSeriesData(string _key) view returns((uint256,uint256,int256)[])
func (_Blockchain *BlockchainSession) GetTimeSeriesData(_key string) ([]DataStructsDataPoint, error) {
	return _Blockchain.Contract.GetTimeSeriesData(&_Blockchain.CallOpts, _key)
}

// GetTimeSeriesData is a free data retrieval call binding the contract method 0x7918f328.
//
// Solidity: function getTimeSeriesData(string _key) view returns((uint256,uint256,int256)[])
func (_Blockchain *BlockchainCallerSession) GetTimeSeriesData(_key string) ([]DataStructsDataPoint, error) {
	return _Blockchain.Contract.GetTimeSeriesData(&_Blockchain.CallOpts, _key)
}

// GetTimeSeriesLength is a free data retrieval call binding the contract method 0x4868686d.
//
// Solidity: function getTimeSeriesLength(string _key) view returns(uint256)
func (_Blockchain *BlockchainCaller) GetTimeSeriesLength(opts *bind.CallOpts, _key string) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getTimeSeriesLength", _key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeSeriesLength is a free data retrieval call binding the contract method 0x4868686d.
//
// Solidity: function getTimeSeriesLength(string _key) view returns(uint256)
func (_Blockchain *BlockchainSession) GetTimeSeriesLength(_key string) (*big.Int, error) {
	return _Blockchain.Contract.GetTimeSeriesLength(&_Blockchain.CallOpts, _key)
}

// GetTimeSeriesLength is a free data retrieval call binding the contract method 0x4868686d.
//
// Solidity: function getTimeSeriesLength(string _key) view returns(uint256)
func (_Blockchain *BlockchainCallerSession) GetTimeSeriesLength(_key string) (*big.Int, error) {
	return _Blockchain.Contract.GetTimeSeriesLength(&_Blockchain.CallOpts, _key)
}

// GetTimeSeriesSlice is a free data retrieval call binding the contract method 0x82800399.
//
// Solidity: function getTimeSeriesSlice(string _key, uint256 _startIndex, uint256 _count) view returns((uint256,uint256,int256)[])
func (_Blockchain *BlockchainCaller) GetTimeSeriesSlice(opts *bind.CallOpts, _key string, _startIndex *big.Int, _count *big.Int) ([]DataStructsDataPoint, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getTimeSeriesSlice", _key, _startIndex, _count)

	if err != nil {
		return *new([]DataStructsDataPoint), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataStructsDataPoint)).(*[]DataStructsDataPoint)

	return out0, err

}

// GetTimeSeriesSlice is a free data retrieval call binding the contract method 0x82800399.
//
// Solidity: function getTimeSeriesSlice(string _key, uint256 _startIndex, uint256 _count) view returns((uint256,uint256,int256)[])
func (_Blockchain *BlockchainSession) GetTimeSeriesSlice(_key string, _startIndex *big.Int, _count *big.Int) ([]DataStructsDataPoint, error) {
	return _Blockchain.Contract.GetTimeSeriesSlice(&_Blockchain.CallOpts, _key, _startIndex, _count)
}

// GetTimeSeriesSlice is a free data retrieval call binding the contract method 0x82800399.
//
// Solidity: function getTimeSeriesSlice(string _key, uint256 _startIndex, uint256 _count) view returns((uint256,uint256,int256)[])
func (_Blockchain *BlockchainCallerSession) GetTimeSeriesSlice(_key string, _startIndex *big.Int, _count *big.Int) ([]DataStructsDataPoint, error) {
	return _Blockchain.Contract.GetTimeSeriesSlice(&_Blockchain.CallOpts, _key, _startIndex, _count)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockchain *BlockchainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockchain *BlockchainSession) Owner() (common.Address, error) {
	return _Blockchain.Contract.Owner(&_Blockchain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockchain *BlockchainCallerSession) Owner() (common.Address, error) {
	return _Blockchain.Contract.Owner(&_Blockchain.CallOpts)
}

// TimeSeriesData is a free data retrieval call binding the contract method 0xf65d99aa.
//
// Solidity: function timeSeriesData(string , uint256 ) view returns(uint256 timestamp, uint256 value, int256 change)
func (_Blockchain *BlockchainCaller) TimeSeriesData(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Timestamp *big.Int
	Value     *big.Int
	Change    *big.Int
}, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "timeSeriesData", arg0, arg1)

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
func (_Blockchain *BlockchainSession) TimeSeriesData(arg0 string, arg1 *big.Int) (struct {
	Timestamp *big.Int
	Value     *big.Int
	Change    *big.Int
}, error) {
	return _Blockchain.Contract.TimeSeriesData(&_Blockchain.CallOpts, arg0, arg1)
}

// TimeSeriesData is a free data retrieval call binding the contract method 0xf65d99aa.
//
// Solidity: function timeSeriesData(string , uint256 ) view returns(uint256 timestamp, uint256 value, int256 change)
func (_Blockchain *BlockchainCallerSession) TimeSeriesData(arg0 string, arg1 *big.Int) (struct {
	Timestamp *big.Int
	Value     *big.Int
	Change    *big.Int
}, error) {
	return _Blockchain.Contract.TimeSeriesData(&_Blockchain.CallOpts, arg0, arg1)
}

// PushDataPoint is a paid mutator transaction binding the contract method 0x53bfac46.
//
// Solidity: function pushDataPoint(string _key, uint256 _value, int256 _change) returns()
func (_Blockchain *BlockchainTransactor) PushDataPoint(opts *bind.TransactOpts, _key string, _value *big.Int, _change *big.Int) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "pushDataPoint", _key, _value, _change)
}

// PushDataPoint is a paid mutator transaction binding the contract method 0x53bfac46.
//
// Solidity: function pushDataPoint(string _key, uint256 _value, int256 _change) returns()
func (_Blockchain *BlockchainSession) PushDataPoint(_key string, _value *big.Int, _change *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.PushDataPoint(&_Blockchain.TransactOpts, _key, _value, _change)
}

// PushDataPoint is a paid mutator transaction binding the contract method 0x53bfac46.
//
// Solidity: function pushDataPoint(string _key, uint256 _value, int256 _change) returns()
func (_Blockchain *BlockchainTransactorSession) PushDataPoint(_key string, _value *big.Int, _change *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.PushDataPoint(&_Blockchain.TransactOpts, _key, _value, _change)
}

// ReplaceTimeSeries is a paid mutator transaction binding the contract method 0xa2a0ff2c.
//
// Solidity: function replaceTimeSeries(string _key, (uint256,uint256,int256)[] _newData) returns()
func (_Blockchain *BlockchainTransactor) ReplaceTimeSeries(opts *bind.TransactOpts, _key string, _newData []DataStructsDataPoint) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "replaceTimeSeries", _key, _newData)
}

// ReplaceTimeSeries is a paid mutator transaction binding the contract method 0xa2a0ff2c.
//
// Solidity: function replaceTimeSeries(string _key, (uint256,uint256,int256)[] _newData) returns()
func (_Blockchain *BlockchainSession) ReplaceTimeSeries(_key string, _newData []DataStructsDataPoint) (*types.Transaction, error) {
	return _Blockchain.Contract.ReplaceTimeSeries(&_Blockchain.TransactOpts, _key, _newData)
}

// ReplaceTimeSeries is a paid mutator transaction binding the contract method 0xa2a0ff2c.
//
// Solidity: function replaceTimeSeries(string _key, (uint256,uint256,int256)[] _newData) returns()
func (_Blockchain *BlockchainTransactorSession) ReplaceTimeSeries(_key string, _newData []DataStructsDataPoint) (*types.Transaction, error) {
	return _Blockchain.Contract.ReplaceTimeSeries(&_Blockchain.TransactOpts, _key, _newData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockchain *BlockchainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockchain *BlockchainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.TransferOwnership(&_Blockchain.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockchain *BlockchainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.TransferOwnership(&_Blockchain.TransactOpts, newOwner)
}

// BlockchainDataPointAddedIterator is returned from FilterDataPointAdded and is used to iterate over the raw logs and unpacked data for DataPointAdded events raised by the Blockchain contract.
type BlockchainDataPointAddedIterator struct {
	Event *BlockchainDataPointAdded // Event containing the contract specifics and raw log

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
func (it *BlockchainDataPointAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainDataPointAdded)
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
		it.Event = new(BlockchainDataPointAdded)
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
func (it *BlockchainDataPointAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainDataPointAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainDataPointAdded represents a DataPointAdded event raised by the Blockchain contract.
type BlockchainDataPointAdded struct {
	Key      string
	NewPoint DataStructsDataPoint
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDataPointAdded is a free log retrieval operation binding the contract event 0x2f533650ae9e5d5131aa36c8d5548a73369e2ea487dc351fcce0fcf93710a6d4.
//
// Solidity: event DataPointAdded(string key, (uint256,uint256,int256) newPoint)
func (_Blockchain *BlockchainFilterer) FilterDataPointAdded(opts *bind.FilterOpts) (*BlockchainDataPointAddedIterator, error) {

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "DataPointAdded")
	if err != nil {
		return nil, err
	}
	return &BlockchainDataPointAddedIterator{contract: _Blockchain.contract, event: "DataPointAdded", logs: logs, sub: sub}, nil
}

// WatchDataPointAdded is a free log subscription operation binding the contract event 0x2f533650ae9e5d5131aa36c8d5548a73369e2ea487dc351fcce0fcf93710a6d4.
//
// Solidity: event DataPointAdded(string key, (uint256,uint256,int256) newPoint)
func (_Blockchain *BlockchainFilterer) WatchDataPointAdded(opts *bind.WatchOpts, sink chan<- *BlockchainDataPointAdded) (event.Subscription, error) {

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "DataPointAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainDataPointAdded)
				if err := _Blockchain.contract.UnpackLog(event, "DataPointAdded", log); err != nil {
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
func (_Blockchain *BlockchainFilterer) ParseDataPointAdded(log types.Log) (*BlockchainDataPointAdded, error) {
	event := new(BlockchainDataPointAdded)
	if err := _Blockchain.contract.UnpackLog(event, "DataPointAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Blockchain contract.
type BlockchainOwnershipTransferredIterator struct {
	Event *BlockchainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlockchainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainOwnershipTransferred)
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
		it.Event = new(BlockchainOwnershipTransferred)
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
func (it *BlockchainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainOwnershipTransferred represents a OwnershipTransferred event raised by the Blockchain contract.
type BlockchainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blockchain *BlockchainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlockchainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainOwnershipTransferredIterator{contract: _Blockchain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blockchain *BlockchainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlockchainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainOwnershipTransferred)
				if err := _Blockchain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Blockchain *BlockchainFilterer) ParseOwnershipTransferred(log types.Log) (*BlockchainOwnershipTransferred, error) {
	event := new(BlockchainOwnershipTransferred)
	if err := _Blockchain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainTimeSeriesReplacedIterator is returned from FilterTimeSeriesReplaced and is used to iterate over the raw logs and unpacked data for TimeSeriesReplaced events raised by the Blockchain contract.
type BlockchainTimeSeriesReplacedIterator struct {
	Event *BlockchainTimeSeriesReplaced // Event containing the contract specifics and raw log

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
func (it *BlockchainTimeSeriesReplacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainTimeSeriesReplaced)
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
		it.Event = new(BlockchainTimeSeriesReplaced)
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
func (it *BlockchainTimeSeriesReplacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainTimeSeriesReplacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainTimeSeriesReplaced represents a TimeSeriesReplaced event raised by the Blockchain contract.
type BlockchainTimeSeriesReplaced struct {
	Key     string
	NewSize *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTimeSeriesReplaced is a free log retrieval operation binding the contract event 0xce9f244d6a6ca82290b92e737914a33cf9ba8c857e20fbe458ca9be02af6fd02.
//
// Solidity: event TimeSeriesReplaced(string key, uint256 newSize)
func (_Blockchain *BlockchainFilterer) FilterTimeSeriesReplaced(opts *bind.FilterOpts) (*BlockchainTimeSeriesReplacedIterator, error) {

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "TimeSeriesReplaced")
	if err != nil {
		return nil, err
	}
	return &BlockchainTimeSeriesReplacedIterator{contract: _Blockchain.contract, event: "TimeSeriesReplaced", logs: logs, sub: sub}, nil
}

// WatchTimeSeriesReplaced is a free log subscription operation binding the contract event 0xce9f244d6a6ca82290b92e737914a33cf9ba8c857e20fbe458ca9be02af6fd02.
//
// Solidity: event TimeSeriesReplaced(string key, uint256 newSize)
func (_Blockchain *BlockchainFilterer) WatchTimeSeriesReplaced(opts *bind.WatchOpts, sink chan<- *BlockchainTimeSeriesReplaced) (event.Subscription, error) {

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "TimeSeriesReplaced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainTimeSeriesReplaced)
				if err := _Blockchain.contract.UnpackLog(event, "TimeSeriesReplaced", log); err != nil {
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

// ParseTimeSeriesReplaced is a log parse operation binding the contract event 0xce9f244d6a6ca82290b92e737914a33cf9ba8c857e20fbe458ca9be02af6fd02.
//
// Solidity: event TimeSeriesReplaced(string key, uint256 newSize)
func (_Blockchain *BlockchainFilterer) ParseTimeSeriesReplaced(log types.Log) (*BlockchainTimeSeriesReplaced, error) {
	event := new(BlockchainTimeSeriesReplaced)
	if err := _Blockchain.contract.UnpackLog(event, "TimeSeriesReplaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
