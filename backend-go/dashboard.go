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

// DashboardDataPoint is an auto generated low-level Go binding around an user-defined struct.
type DashboardDataPoint struct {
	Timestamp *big.Int
	Value     *big.Int
	Change    *big.Int
}

// DashboardFundItem is an auto generated low-level Go binding around an user-defined struct.
type DashboardFundItem struct {
	Name        string
	Price       *big.Int
	IconId      string
	Count       *big.Int
	AvatarColor string
}

// DashboardJob is an auto generated low-level Go binding around an user-defined struct.
type DashboardJob struct {
	Foundation string
	Position   string
	Field      string
	Salary     *big.Int
	Form       string
	Trend      uint8
	IconId     string
}

// DashboardService is an auto generated low-level Go binding around an user-defined struct.
type DashboardService struct {
	Name     string
	Title    string
	Subtitle string
	Status   uint8
	IconId   string
}

// DashboardTestimonial is an auto generated low-level Go binding around an user-defined struct.
type DashboardTestimonial struct {
	Quote    string
	AvatarId string
}

// DashboardMetaData contains all meta data concerning the Dashboard contract.
var DashboardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"change\",\"type\":\"int256\"}],\"indexed\":false,\"internalType\":\"structDashboard.DataPoint\",\"name\":\"newPoint\",\"type\":\"tuple\"}],\"name\":\"DataPointAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dayKey\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"avatarColor\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structDashboard.FundItem[]\",\"name\":\"newItems\",\"type\":\"tuple[]\"}],\"name\":\"FundItemsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"foundation\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"salary\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"form\",\"type\":\"string\"},{\"internalType\":\"enumDashboard.JobTrend\",\"name\":\"trend\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structDashboard.Job\",\"name\":\"newJob\",\"type\":\"tuple\"}],\"name\":\"JobAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"newValue\",\"type\":\"int256\"}],\"name\":\"ReturnValueUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumDashboard.ServiceStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"ServiceStatusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"SingleValueUpdated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"foundation\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"salary\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"form\",\"type\":\"string\"},{\"internalType\":\"enumDashboard.JobTrend\",\"name\":\"trend\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"}],\"internalType\":\"structDashboard.Job\",\"name\":\"_newJob\",\"type\":\"tuple\"}],\"name\":\"addJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fundItemsByDay\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"avatarColor\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllJobs\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"foundation\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"salary\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"form\",\"type\":\"string\"},{\"internalType\":\"enumDashboard.JobTrend\",\"name\":\"trend\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"}],\"internalType\":\"structDashboard.Job[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllServices\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subtitle\",\"type\":\"string\"},{\"internalType\":\"enumDashboard.ServiceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"}],\"internalType\":\"structDashboard.Service[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTestimonials\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"quote\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"avatarId\",\"type\":\"string\"}],\"internalType\":\"structDashboard.Testimonial[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_day\",\"type\":\"string\"}],\"name\":\"getFundItemsByDay\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"avatarColor\",\"type\":\"string\"}],\"internalType\":\"structDashboard.FundItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getReturnData\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getSingleValueData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getTimeSeriesData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"change\",\"type\":\"int256\"}],\"internalType\":\"structDashboard.DataPoint[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"foundation\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"salary\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"form\",\"type\":\"string\"},{\"internalType\":\"enumDashboard.JobTrend\",\"name\":\"trend\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_change\",\"type\":\"int256\"}],\"name\":\"pushDataPoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"returnData\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"services\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subtitle\",\"type\":\"string\"},{\"internalType\":\"enumDashboard.ServiceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"singleValueData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"testimonials\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"quote\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"avatarId\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"timeSeriesData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"change\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_dayKey\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"iconId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"avatarColor\",\"type\":\"string\"}],\"internalType\":\"structDashboard.FundItem[]\",\"name\":\"_newItems\",\"type\":\"tuple[]\"}],\"name\":\"updateFundItemsForDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"enumDashboard.ServiceStatus\",\"name\":\"_newStatus\",\"type\":\"uint8\"}],\"name\":\"updateServiceStatusByIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_timeframe\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_viewers\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_returnPct\",\"type\":\"int256\"}],\"name\":\"updateTotalViewersData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DashboardABI is the input ABI used to generate the binding from.
// Deprecated: Use DashboardMetaData.ABI instead.
var DashboardABI = DashboardMetaData.ABI

// Dashboard is an auto generated Go binding around an Ethereum contract.
type Dashboard struct {
	DashboardCaller     // Read-only binding to the contract
	DashboardTransactor // Write-only binding to the contract
	DashboardFilterer   // Log filterer for contract events
}

// DashboardCaller is an auto generated read-only Go binding around an Ethereum contract.
type DashboardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DashboardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DashboardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DashboardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DashboardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DashboardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DashboardSession struct {
	Contract     *Dashboard        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DashboardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DashboardCallerSession struct {
	Contract *DashboardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DashboardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DashboardTransactorSession struct {
	Contract     *DashboardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DashboardRaw is an auto generated low-level Go binding around an Ethereum contract.
type DashboardRaw struct {
	Contract *Dashboard // Generic contract binding to access the raw methods on
}

// DashboardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DashboardCallerRaw struct {
	Contract *DashboardCaller // Generic read-only contract binding to access the raw methods on
}

// DashboardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DashboardTransactorRaw struct {
	Contract *DashboardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDashboard creates a new instance of Dashboard, bound to a specific deployed contract.
func NewDashboard(address common.Address, backend bind.ContractBackend) (*Dashboard, error) {
	contract, err := bindDashboard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dashboard{DashboardCaller: DashboardCaller{contract: contract}, DashboardTransactor: DashboardTransactor{contract: contract}, DashboardFilterer: DashboardFilterer{contract: contract}}, nil
}

// NewDashboardCaller creates a new read-only instance of Dashboard, bound to a specific deployed contract.
func NewDashboardCaller(address common.Address, caller bind.ContractCaller) (*DashboardCaller, error) {
	contract, err := bindDashboard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DashboardCaller{contract: contract}, nil
}

// NewDashboardTransactor creates a new write-only instance of Dashboard, bound to a specific deployed contract.
func NewDashboardTransactor(address common.Address, transactor bind.ContractTransactor) (*DashboardTransactor, error) {
	contract, err := bindDashboard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DashboardTransactor{contract: contract}, nil
}

// NewDashboardFilterer creates a new log filterer instance of Dashboard, bound to a specific deployed contract.
func NewDashboardFilterer(address common.Address, filterer bind.ContractFilterer) (*DashboardFilterer, error) {
	contract, err := bindDashboard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DashboardFilterer{contract: contract}, nil
}

// bindDashboard binds a generic wrapper to an already deployed contract.
func bindDashboard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DashboardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dashboard *DashboardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dashboard.Contract.DashboardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dashboard *DashboardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dashboard.Contract.DashboardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dashboard *DashboardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dashboard.Contract.DashboardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dashboard *DashboardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dashboard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dashboard *DashboardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dashboard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dashboard *DashboardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dashboard.Contract.contract.Transact(opts, method, params...)
}

// FundItemsByDay is a free data retrieval call binding the contract method 0x36a48ee5.
//
// Solidity: function fundItemsByDay(string , uint256 ) view returns(string name, uint256 price, string iconId, uint256 count, string avatarColor)
func (_Dashboard *DashboardCaller) FundItemsByDay(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Name        string
	Price       *big.Int
	IconId      string
	Count       *big.Int
	AvatarColor string
}, error) {
	var out []interface{}
	err := _Dashboard.contract.Call(opts, &out, "fundItemsByDay", arg0, arg1)

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
func (_Dashboard *DashboardSession) FundItemsByDay(arg0 string, arg1 *big.Int) (struct {
	Name        string
	Price       *big.Int
	IconId      string
	Count       *big.Int
	AvatarColor string
}, error) {
	return _Dashboard.Contract.FundItemsByDay(&_Dashboard.CallOpts, arg0, arg1)
}

// FundItemsByDay is a free data retrieval call binding the contract method 0x36a48ee5.
//
// Solidity: function fundItemsByDay(string , uint256 ) view returns(string name, uint256 price, string iconId, uint256 count, string avatarColor)
func (_Dashboard *DashboardCallerSession) FundItemsByDay(arg0 string, arg1 *big.Int) (struct {
	Name        string
	Price       *big.Int
	IconId      string
	Count       *big.Int
	AvatarColor string
}, error) {
	return _Dashboard.Contract.FundItemsByDay(&_Dashboard.CallOpts, arg0, arg1)
}

// GetAllJobs is a free data retrieval call binding the contract method 0x965a79fb.
//
// Solidity: function getAllJobs() view returns((string,string,string,uint256,string,uint8,string)[])
func (_Dashboard *DashboardCaller) GetAllJobs(opts *bind.CallOpts) ([]DashboardJob, error) {
	var out []interface{}
	err := _Dashboard.contract.Call(opts, &out, "getAllJobs")

	if err != nil {
		return *new([]DashboardJob), err
	}

	out0 := *abi.ConvertType(out[0], new([]DashboardJob)).(*[]DashboardJob)

	return out0, err

}

// GetAllJobs is a free data retrieval call binding the contract method 0x965a79fb.
//
// Solidity: function getAllJobs() view returns((string,string,string,uint256,string,uint8,string)[])
func (_Dashboard *DashboardSession) GetAllJobs() ([]DashboardJob, error) {
	return _Dashboard.Contract.GetAllJobs(&_Dashboard.CallOpts)
}

// GetAllJobs is a free data retrieval call binding the contract method 0x965a79fb.
//
// Solidity: function getAllJobs() view returns((string,string,string,uint256,string,uint8,string)[])
func (_Dashboard *DashboardCallerSession) GetAllJobs() ([]DashboardJob, error) {
	return _Dashboard.Contract.GetAllJobs(&_Dashboard.CallOpts)
}

// GetAllServices is a free data retrieval call binding the contract method 0x21fe0f30.
//
// Solidity: function getAllServices() view returns((string,string,string,uint8,string)[])
func (_Dashboard *DashboardCaller) GetAllServices(opts *bind.CallOpts) ([]DashboardService, error) {
	var out []interface{}
	err := _Dashboard.contract.Call(opts, &out, "getAllServices")

	if err != nil {
		return *new([]DashboardService), err
	}

	out0 := *abi.ConvertType(out[0], new([]DashboardService)).(*[]DashboardService)

	return out0, err

}

// GetAllServices is a free data retrieval call binding the contract method 0x21fe0f30.
//
// Solidity: function getAllServices() view returns((string,string,string,uint8,string)[])
func (_Dashboard *DashboardSession) GetAllServices() ([]DashboardService, error) {
	return _Dashboard.Contract.GetAllServices(&_Dashboard.CallOpts)
}

// GetAllServices is a free data retrieval call binding the contract method 0x21fe0f30.
//
// Solidity: function getAllServices() view returns((string,string,string,uint8,string)[])
func (_Dashboard *DashboardCallerSession) GetAllServices() ([]DashboardService, error) {
	return _Dashboard.Contract.GetAllServices(&_Dashboard.CallOpts)
}

// GetAllTestimonials is a free data retrieval call binding the contract method 0x9fe17bd9.
//
// Solidity: function getAllTestimonials() view returns((string,string)[])
func (_Dashboard *DashboardCaller) GetAllTestimonials(opts *bind.CallOpts) ([]DashboardTestimonial, error) {
	var out []interface{}
	err := _Dashboard.contract.Call(opts, &out, "getAllTestimonials")

	if err != nil {
		return *new([]DashboardTestimonial), err
	}

	out0 := *abi.ConvertType(out[0], new([]DashboardTestimonial)).(*[]DashboardTestimonial)

	return out0, err

}

// GetAllTestimonials is a free data retrieval call binding the contract method 0x9fe17bd9.
//
// Solidity: function getAllTestimonials() view returns((string,string)[])
func (_Dashboard *DashboardSession) GetAllTestimonials() ([]DashboardTestimonial, error) {
	return _Dashboard.Contract.GetAllTestimonials(&_Dashboard.CallOpts)
}

// GetAllTestimonials is a free data retrieval call binding the contract method 0x9fe17bd9.
//
// Solidity: function getAllTestimonials() view returns((string,string)[])
func (_Dashboard *DashboardCallerSession) GetAllTestimonials() ([]DashboardTestimonial, error) {
	return _Dashboard.Contract.GetAllTestimonials(&_Dashboard.CallOpts)
}

// GetFundItemsByDay is a free data retrieval call binding the contract method 0x6ca07056.
//
// Solidity: function getFundItemsByDay(string _day) view returns((string,uint256,string,uint256,string)[])
func (_Dashboard *DashboardCaller) GetFundItemsByDay(opts *bind.CallOpts, _day string) ([]DashboardFundItem, error) {
	var out []interface{}
	err := _Dashboard.contract.Call(opts, &out, "getFundItemsByDay", _day)

	if err != nil {
		return *new([]DashboardFundItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]DashboardFundItem)).(*[]DashboardFundItem)

	return out0, err

}

// GetFundItemsByDay is a free data retrieval call binding the contract method 0x6ca07056.
//
// Solidity: function getFundItemsByDay(string _day) view returns((string,uint256,string,uint256,string)[])
func (_Dashboard *DashboardSession) GetFundItemsByDay(_day string) ([]DashboardFundItem, error) {
	return _Dashboard.Contract.GetFundItemsByDay(&_Dashboard.CallOpts, _day)
}

// GetFundItemsByDay is a free data retrieval call binding the contract method 0x6ca07056.
//
// Solidity: function getFundItemsByDay(string _day) view returns((string,uint256,string,uint256,string)[])
func (_Dashboard *DashboardCallerSession) GetFundItemsByDay(_day string) ([]DashboardFundItem, error) {
	return _Dashboard.Contract.GetFundItemsByDay(&_Dashboard.CallOpts, _day)
}

// GetReturnData is a free data retrieval call binding the contract method 0x56b0bb32.
//
// Solidity: function getReturnData(string _key) view returns(int256)
func (_Dashboard *DashboardCaller) GetReturnData(opts *bind.CallOpts, _key string) (*big.Int, error) {
	var out []interface{}
	err := _Dashboard.contract.Call(opts, &out, "getReturnData", _key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReturnData is a free data retrieval call binding the contract method 0x56b0bb32.
//
// Solidity: function getReturnData(string _key) view returns(int256)
func (_Dashboard *DashboardSession) GetReturnData(_key string) (*big.Int, error) {
	return _Dashboard.Contract.GetReturnData(&_Dashboard.CallOpts, _key)
}

// GetReturnData is a free data retrieval call binding the contract method 0x56b0bb32.
//
// Solidity: function getReturnData(string _key) view returns(int256)
func (_Dashboard *DashboardCallerSession) GetReturnData(_key string) (*big.Int, error) {
	return _Dashboard.Contract.GetReturnData(&_Dashboard.CallOpts, _key)
}

// GetSingleValueData is a free data retrieval call binding the contract method 0x166d8b41.
//
// Solidity: function getSingleValueData(string _key) view returns(uint256)
func (_Dashboard *DashboardCaller) GetSingleValueData(opts *bind.CallOpts, _key string) (*big.Int, error) {
	var out []interface{}
	err := _Dashboard.contract.Call(opts, &out, "getSingleValueData", _key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSingleValueData is a free data retrieval call binding the contract method 0x166d8b41.
//
// Solidity: function getSingleValueData(string _key) view returns(uint256)
func (_Dashboard *DashboardSession) GetSingleValueData(_key string) (*big.Int, error) {
	return _Dashboard.Contract.GetSingleValueData(&_Dashboard.CallOpts, _key)
}

// GetSingleValueData is a free data retrieval call binding the contract method 0x166d8b41.
//
// Solidity: function getSingleValueData(string _key) view returns(uint256)
func (_Dashboard *DashboardCallerSession) GetSingleValueData(_key string) (*big.Int, error) {
	return _Dashboard.Contract.GetSingleValueData(&_Dashboard.CallOpts, _key)
}

// GetTimeSeriesData is a free data retrieval call binding the contract method 0x7918f328.
//
// Solidity: function getTimeSeriesData(string _key) view returns((uint256,uint256,int256)[])
func (_Dashboard *DashboardCaller) GetTimeSeriesData(opts *bind.CallOpts, _key string) ([]DashboardDataPoint, error) {
	var out []interface{}
	err := _Dashboard.contract.Call(opts, &out, "getTimeSeriesData", _key)

	if err != nil {
		return *new([]DashboardDataPoint), err
	}

	out0 := *abi.ConvertType(out[0], new([]DashboardDataPoint)).(*[]DashboardDataPoint)

	return out0, err

}

// GetTimeSeriesData is a free data retrieval call binding the contract method 0x7918f328.
//
// Solidity: function getTimeSeriesData(string _key) view returns((uint256,uint256,int256)[])
func (_Dashboard *DashboardSession) GetTimeSeriesData(_key string) ([]DashboardDataPoint, error) {
	return _Dashboard.Contract.GetTimeSeriesData(&_Dashboard.CallOpts, _key)
}

// GetTimeSeriesData is a free data retrieval call binding the contract method 0x7918f328.
//
// Solidity: function getTimeSeriesData(string _key) view returns((uint256,uint256,int256)[])
func (_Dashboard *DashboardCallerSession) GetTimeSeriesData(_key string) ([]DashboardDataPoint, error) {
	return _Dashboard.Contract.GetTimeSeriesData(&_Dashboard.CallOpts, _key)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(string foundation, string position, string field, uint256 salary, string form, uint8 trend, string iconId)
func (_Dashboard *DashboardCaller) Jobs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Foundation string
	Position   string
	Field      string
	Salary     *big.Int
	Form       string
	Trend      uint8
	IconId     string
}, error) {
	var out []interface{}
	err := _Dashboard.contract.Call(opts, &out, "jobs", arg0)

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
func (_Dashboard *DashboardSession) Jobs(arg0 *big.Int) (struct {
	Foundation string
	Position   string
	Field      string
	Salary     *big.Int
	Form       string
	Trend      uint8
	IconId     string
}, error) {
	return _Dashboard.Contract.Jobs(&_Dashboard.CallOpts, arg0)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(string foundation, string position, string field, uint256 salary, string form, uint8 trend, string iconId)
func (_Dashboard *DashboardCallerSession) Jobs(arg0 *big.Int) (struct {
	Foundation string
	Position   string
	Field      string
	Salary     *big.Int
	Form       string
	Trend      uint8
	IconId     string
}, error) {
	return _Dashboard.Contract.Jobs(&_Dashboard.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dashboard *DashboardCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dashboard.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dashboard *DashboardSession) Owner() (common.Address, error) {
	return _Dashboard.Contract.Owner(&_Dashboard.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dashboard *DashboardCallerSession) Owner() (common.Address, error) {
	return _Dashboard.Contract.Owner(&_Dashboard.CallOpts)
}

// ReturnData is a free data retrieval call binding the contract method 0xe5f92e86.
//
// Solidity: function returnData(string ) view returns(int256)
func (_Dashboard *DashboardCaller) ReturnData(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Dashboard.contract.Call(opts, &out, "returnData", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReturnData is a free data retrieval call binding the contract method 0xe5f92e86.
//
// Solidity: function returnData(string ) view returns(int256)
func (_Dashboard *DashboardSession) ReturnData(arg0 string) (*big.Int, error) {
	return _Dashboard.Contract.ReturnData(&_Dashboard.CallOpts, arg0)
}

// ReturnData is a free data retrieval call binding the contract method 0xe5f92e86.
//
// Solidity: function returnData(string ) view returns(int256)
func (_Dashboard *DashboardCallerSession) ReturnData(arg0 string) (*big.Int, error) {
	return _Dashboard.Contract.ReturnData(&_Dashboard.CallOpts, arg0)
}

// Services is a free data retrieval call binding the contract method 0xc22c4f43.
//
// Solidity: function services(uint256 ) view returns(string name, string title, string subtitle, uint8 status, string iconId)
func (_Dashboard *DashboardCaller) Services(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name     string
	Title    string
	Subtitle string
	Status   uint8
	IconId   string
}, error) {
	var out []interface{}
	err := _Dashboard.contract.Call(opts, &out, "services", arg0)

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
func (_Dashboard *DashboardSession) Services(arg0 *big.Int) (struct {
	Name     string
	Title    string
	Subtitle string
	Status   uint8
	IconId   string
}, error) {
	return _Dashboard.Contract.Services(&_Dashboard.CallOpts, arg0)
}

// Services is a free data retrieval call binding the contract method 0xc22c4f43.
//
// Solidity: function services(uint256 ) view returns(string name, string title, string subtitle, uint8 status, string iconId)
func (_Dashboard *DashboardCallerSession) Services(arg0 *big.Int) (struct {
	Name     string
	Title    string
	Subtitle string
	Status   uint8
	IconId   string
}, error) {
	return _Dashboard.Contract.Services(&_Dashboard.CallOpts, arg0)
}

// SingleValueData is a free data retrieval call binding the contract method 0x5a2ded70.
//
// Solidity: function singleValueData(string ) view returns(uint256)
func (_Dashboard *DashboardCaller) SingleValueData(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Dashboard.contract.Call(opts, &out, "singleValueData", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SingleValueData is a free data retrieval call binding the contract method 0x5a2ded70.
//
// Solidity: function singleValueData(string ) view returns(uint256)
func (_Dashboard *DashboardSession) SingleValueData(arg0 string) (*big.Int, error) {
	return _Dashboard.Contract.SingleValueData(&_Dashboard.CallOpts, arg0)
}

// SingleValueData is a free data retrieval call binding the contract method 0x5a2ded70.
//
// Solidity: function singleValueData(string ) view returns(uint256)
func (_Dashboard *DashboardCallerSession) SingleValueData(arg0 string) (*big.Int, error) {
	return _Dashboard.Contract.SingleValueData(&_Dashboard.CallOpts, arg0)
}

// Testimonials is a free data retrieval call binding the contract method 0x89dc7494.
//
// Solidity: function testimonials(uint256 ) view returns(string quote, string avatarId)
func (_Dashboard *DashboardCaller) Testimonials(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Quote    string
	AvatarId string
}, error) {
	var out []interface{}
	err := _Dashboard.contract.Call(opts, &out, "testimonials", arg0)

	outstruct := new(struct {
		Quote    string
		AvatarId string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Quote = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.AvatarId = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Testimonials is a free data retrieval call binding the contract method 0x89dc7494.
//
// Solidity: function testimonials(uint256 ) view returns(string quote, string avatarId)
func (_Dashboard *DashboardSession) Testimonials(arg0 *big.Int) (struct {
	Quote    string
	AvatarId string
}, error) {
	return _Dashboard.Contract.Testimonials(&_Dashboard.CallOpts, arg0)
}

// Testimonials is a free data retrieval call binding the contract method 0x89dc7494.
//
// Solidity: function testimonials(uint256 ) view returns(string quote, string avatarId)
func (_Dashboard *DashboardCallerSession) Testimonials(arg0 *big.Int) (struct {
	Quote    string
	AvatarId string
}, error) {
	return _Dashboard.Contract.Testimonials(&_Dashboard.CallOpts, arg0)
}

// TimeSeriesData is a free data retrieval call binding the contract method 0xf65d99aa.
//
// Solidity: function timeSeriesData(string , uint256 ) view returns(uint256 timestamp, uint256 value, int256 change)
func (_Dashboard *DashboardCaller) TimeSeriesData(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Timestamp *big.Int
	Value     *big.Int
	Change    *big.Int
}, error) {
	var out []interface{}
	err := _Dashboard.contract.Call(opts, &out, "timeSeriesData", arg0, arg1)

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
func (_Dashboard *DashboardSession) TimeSeriesData(arg0 string, arg1 *big.Int) (struct {
	Timestamp *big.Int
	Value     *big.Int
	Change    *big.Int
}, error) {
	return _Dashboard.Contract.TimeSeriesData(&_Dashboard.CallOpts, arg0, arg1)
}

// TimeSeriesData is a free data retrieval call binding the contract method 0xf65d99aa.
//
// Solidity: function timeSeriesData(string , uint256 ) view returns(uint256 timestamp, uint256 value, int256 change)
func (_Dashboard *DashboardCallerSession) TimeSeriesData(arg0 string, arg1 *big.Int) (struct {
	Timestamp *big.Int
	Value     *big.Int
	Change    *big.Int
}, error) {
	return _Dashboard.Contract.TimeSeriesData(&_Dashboard.CallOpts, arg0, arg1)
}

// AddJob is a paid mutator transaction binding the contract method 0x29b44e58.
//
// Solidity: function addJob((string,string,string,uint256,string,uint8,string) _newJob) returns()
func (_Dashboard *DashboardTransactor) AddJob(opts *bind.TransactOpts, _newJob DashboardJob) (*types.Transaction, error) {
	return _Dashboard.contract.Transact(opts, "addJob", _newJob)
}

// AddJob is a paid mutator transaction binding the contract method 0x29b44e58.
//
// Solidity: function addJob((string,string,string,uint256,string,uint8,string) _newJob) returns()
func (_Dashboard *DashboardSession) AddJob(_newJob DashboardJob) (*types.Transaction, error) {
	return _Dashboard.Contract.AddJob(&_Dashboard.TransactOpts, _newJob)
}

// AddJob is a paid mutator transaction binding the contract method 0x29b44e58.
//
// Solidity: function addJob((string,string,string,uint256,string,uint8,string) _newJob) returns()
func (_Dashboard *DashboardTransactorSession) AddJob(_newJob DashboardJob) (*types.Transaction, error) {
	return _Dashboard.Contract.AddJob(&_Dashboard.TransactOpts, _newJob)
}

// PushDataPoint is a paid mutator transaction binding the contract method 0x53bfac46.
//
// Solidity: function pushDataPoint(string _key, uint256 _value, int256 _change) returns()
func (_Dashboard *DashboardTransactor) PushDataPoint(opts *bind.TransactOpts, _key string, _value *big.Int, _change *big.Int) (*types.Transaction, error) {
	return _Dashboard.contract.Transact(opts, "pushDataPoint", _key, _value, _change)
}

// PushDataPoint is a paid mutator transaction binding the contract method 0x53bfac46.
//
// Solidity: function pushDataPoint(string _key, uint256 _value, int256 _change) returns()
func (_Dashboard *DashboardSession) PushDataPoint(_key string, _value *big.Int, _change *big.Int) (*types.Transaction, error) {
	return _Dashboard.Contract.PushDataPoint(&_Dashboard.TransactOpts, _key, _value, _change)
}

// PushDataPoint is a paid mutator transaction binding the contract method 0x53bfac46.
//
// Solidity: function pushDataPoint(string _key, uint256 _value, int256 _change) returns()
func (_Dashboard *DashboardTransactorSession) PushDataPoint(_key string, _value *big.Int, _change *big.Int) (*types.Transaction, error) {
	return _Dashboard.Contract.PushDataPoint(&_Dashboard.TransactOpts, _key, _value, _change)
}

// UpdateFundItemsForDay is a paid mutator transaction binding the contract method 0xcfb2749d.
//
// Solidity: function updateFundItemsForDay(string _dayKey, (string,uint256,string,uint256,string)[] _newItems) returns()
func (_Dashboard *DashboardTransactor) UpdateFundItemsForDay(opts *bind.TransactOpts, _dayKey string, _newItems []DashboardFundItem) (*types.Transaction, error) {
	return _Dashboard.contract.Transact(opts, "updateFundItemsForDay", _dayKey, _newItems)
}

// UpdateFundItemsForDay is a paid mutator transaction binding the contract method 0xcfb2749d.
//
// Solidity: function updateFundItemsForDay(string _dayKey, (string,uint256,string,uint256,string)[] _newItems) returns()
func (_Dashboard *DashboardSession) UpdateFundItemsForDay(_dayKey string, _newItems []DashboardFundItem) (*types.Transaction, error) {
	return _Dashboard.Contract.UpdateFundItemsForDay(&_Dashboard.TransactOpts, _dayKey, _newItems)
}

// UpdateFundItemsForDay is a paid mutator transaction binding the contract method 0xcfb2749d.
//
// Solidity: function updateFundItemsForDay(string _dayKey, (string,uint256,string,uint256,string)[] _newItems) returns()
func (_Dashboard *DashboardTransactorSession) UpdateFundItemsForDay(_dayKey string, _newItems []DashboardFundItem) (*types.Transaction, error) {
	return _Dashboard.Contract.UpdateFundItemsForDay(&_Dashboard.TransactOpts, _dayKey, _newItems)
}

// UpdateServiceStatusByIndex is a paid mutator transaction binding the contract method 0xeffc7408.
//
// Solidity: function updateServiceStatusByIndex(uint256 _index, uint8 _newStatus) returns()
func (_Dashboard *DashboardTransactor) UpdateServiceStatusByIndex(opts *bind.TransactOpts, _index *big.Int, _newStatus uint8) (*types.Transaction, error) {
	return _Dashboard.contract.Transact(opts, "updateServiceStatusByIndex", _index, _newStatus)
}

// UpdateServiceStatusByIndex is a paid mutator transaction binding the contract method 0xeffc7408.
//
// Solidity: function updateServiceStatusByIndex(uint256 _index, uint8 _newStatus) returns()
func (_Dashboard *DashboardSession) UpdateServiceStatusByIndex(_index *big.Int, _newStatus uint8) (*types.Transaction, error) {
	return _Dashboard.Contract.UpdateServiceStatusByIndex(&_Dashboard.TransactOpts, _index, _newStatus)
}

// UpdateServiceStatusByIndex is a paid mutator transaction binding the contract method 0xeffc7408.
//
// Solidity: function updateServiceStatusByIndex(uint256 _index, uint8 _newStatus) returns()
func (_Dashboard *DashboardTransactorSession) UpdateServiceStatusByIndex(_index *big.Int, _newStatus uint8) (*types.Transaction, error) {
	return _Dashboard.Contract.UpdateServiceStatusByIndex(&_Dashboard.TransactOpts, _index, _newStatus)
}

// UpdateTotalViewersData is a paid mutator transaction binding the contract method 0x9e536ab9.
//
// Solidity: function updateTotalViewersData(string _timeframe, uint256 _viewers, int256 _returnPct) returns()
func (_Dashboard *DashboardTransactor) UpdateTotalViewersData(opts *bind.TransactOpts, _timeframe string, _viewers *big.Int, _returnPct *big.Int) (*types.Transaction, error) {
	return _Dashboard.contract.Transact(opts, "updateTotalViewersData", _timeframe, _viewers, _returnPct)
}

// UpdateTotalViewersData is a paid mutator transaction binding the contract method 0x9e536ab9.
//
// Solidity: function updateTotalViewersData(string _timeframe, uint256 _viewers, int256 _returnPct) returns()
func (_Dashboard *DashboardSession) UpdateTotalViewersData(_timeframe string, _viewers *big.Int, _returnPct *big.Int) (*types.Transaction, error) {
	return _Dashboard.Contract.UpdateTotalViewersData(&_Dashboard.TransactOpts, _timeframe, _viewers, _returnPct)
}

// UpdateTotalViewersData is a paid mutator transaction binding the contract method 0x9e536ab9.
//
// Solidity: function updateTotalViewersData(string _timeframe, uint256 _viewers, int256 _returnPct) returns()
func (_Dashboard *DashboardTransactorSession) UpdateTotalViewersData(_timeframe string, _viewers *big.Int, _returnPct *big.Int) (*types.Transaction, error) {
	return _Dashboard.Contract.UpdateTotalViewersData(&_Dashboard.TransactOpts, _timeframe, _viewers, _returnPct)
}

// DashboardDataPointAddedIterator is returned from FilterDataPointAdded and is used to iterate over the raw logs and unpacked data for DataPointAdded events raised by the Dashboard contract.
type DashboardDataPointAddedIterator struct {
	Event *DashboardDataPointAdded // Event containing the contract specifics and raw log

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
func (it *DashboardDataPointAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DashboardDataPointAdded)
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
		it.Event = new(DashboardDataPointAdded)
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
func (it *DashboardDataPointAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DashboardDataPointAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DashboardDataPointAdded represents a DataPointAdded event raised by the Dashboard contract.
type DashboardDataPointAdded struct {
	Key      string
	NewPoint DashboardDataPoint
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDataPointAdded is a free log retrieval operation binding the contract event 0x2f533650ae9e5d5131aa36c8d5548a73369e2ea487dc351fcce0fcf93710a6d4.
//
// Solidity: event DataPointAdded(string key, (uint256,uint256,int256) newPoint)
func (_Dashboard *DashboardFilterer) FilterDataPointAdded(opts *bind.FilterOpts) (*DashboardDataPointAddedIterator, error) {

	logs, sub, err := _Dashboard.contract.FilterLogs(opts, "DataPointAdded")
	if err != nil {
		return nil, err
	}
	return &DashboardDataPointAddedIterator{contract: _Dashboard.contract, event: "DataPointAdded", logs: logs, sub: sub}, nil
}

// WatchDataPointAdded is a free log subscription operation binding the contract event 0x2f533650ae9e5d5131aa36c8d5548a73369e2ea487dc351fcce0fcf93710a6d4.
//
// Solidity: event DataPointAdded(string key, (uint256,uint256,int256) newPoint)
func (_Dashboard *DashboardFilterer) WatchDataPointAdded(opts *bind.WatchOpts, sink chan<- *DashboardDataPointAdded) (event.Subscription, error) {

	logs, sub, err := _Dashboard.contract.WatchLogs(opts, "DataPointAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DashboardDataPointAdded)
				if err := _Dashboard.contract.UnpackLog(event, "DataPointAdded", log); err != nil {
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
func (_Dashboard *DashboardFilterer) ParseDataPointAdded(log types.Log) (*DashboardDataPointAdded, error) {
	event := new(DashboardDataPointAdded)
	if err := _Dashboard.contract.UnpackLog(event, "DataPointAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DashboardFundItemsUpdatedIterator is returned from FilterFundItemsUpdated and is used to iterate over the raw logs and unpacked data for FundItemsUpdated events raised by the Dashboard contract.
type DashboardFundItemsUpdatedIterator struct {
	Event *DashboardFundItemsUpdated // Event containing the contract specifics and raw log

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
func (it *DashboardFundItemsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DashboardFundItemsUpdated)
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
		it.Event = new(DashboardFundItemsUpdated)
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
func (it *DashboardFundItemsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DashboardFundItemsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DashboardFundItemsUpdated represents a FundItemsUpdated event raised by the Dashboard contract.
type DashboardFundItemsUpdated struct {
	DayKey   string
	NewItems []DashboardFundItem
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFundItemsUpdated is a free log retrieval operation binding the contract event 0xc2449eda7f6620ae8ff0183b4dab7842c5427052d7916956e7e9b252f7d51e06.
//
// Solidity: event FundItemsUpdated(string dayKey, (string,uint256,string,uint256,string)[] newItems)
func (_Dashboard *DashboardFilterer) FilterFundItemsUpdated(opts *bind.FilterOpts) (*DashboardFundItemsUpdatedIterator, error) {

	logs, sub, err := _Dashboard.contract.FilterLogs(opts, "FundItemsUpdated")
	if err != nil {
		return nil, err
	}
	return &DashboardFundItemsUpdatedIterator{contract: _Dashboard.contract, event: "FundItemsUpdated", logs: logs, sub: sub}, nil
}

// WatchFundItemsUpdated is a free log subscription operation binding the contract event 0xc2449eda7f6620ae8ff0183b4dab7842c5427052d7916956e7e9b252f7d51e06.
//
// Solidity: event FundItemsUpdated(string dayKey, (string,uint256,string,uint256,string)[] newItems)
func (_Dashboard *DashboardFilterer) WatchFundItemsUpdated(opts *bind.WatchOpts, sink chan<- *DashboardFundItemsUpdated) (event.Subscription, error) {

	logs, sub, err := _Dashboard.contract.WatchLogs(opts, "FundItemsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DashboardFundItemsUpdated)
				if err := _Dashboard.contract.UnpackLog(event, "FundItemsUpdated", log); err != nil {
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
func (_Dashboard *DashboardFilterer) ParseFundItemsUpdated(log types.Log) (*DashboardFundItemsUpdated, error) {
	event := new(DashboardFundItemsUpdated)
	if err := _Dashboard.contract.UnpackLog(event, "FundItemsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DashboardJobAddedIterator is returned from FilterJobAdded and is used to iterate over the raw logs and unpacked data for JobAdded events raised by the Dashboard contract.
type DashboardJobAddedIterator struct {
	Event *DashboardJobAdded // Event containing the contract specifics and raw log

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
func (it *DashboardJobAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DashboardJobAdded)
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
		it.Event = new(DashboardJobAdded)
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
func (it *DashboardJobAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DashboardJobAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DashboardJobAdded represents a JobAdded event raised by the Dashboard contract.
type DashboardJobAdded struct {
	NewJob DashboardJob
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterJobAdded is a free log retrieval operation binding the contract event 0xde0eae4446badd721c880f61c8e29849509fd7ced6259dcc27a46b2affa3329a.
//
// Solidity: event JobAdded((string,string,string,uint256,string,uint8,string) newJob)
func (_Dashboard *DashboardFilterer) FilterJobAdded(opts *bind.FilterOpts) (*DashboardJobAddedIterator, error) {

	logs, sub, err := _Dashboard.contract.FilterLogs(opts, "JobAdded")
	if err != nil {
		return nil, err
	}
	return &DashboardJobAddedIterator{contract: _Dashboard.contract, event: "JobAdded", logs: logs, sub: sub}, nil
}

// WatchJobAdded is a free log subscription operation binding the contract event 0xde0eae4446badd721c880f61c8e29849509fd7ced6259dcc27a46b2affa3329a.
//
// Solidity: event JobAdded((string,string,string,uint256,string,uint8,string) newJob)
func (_Dashboard *DashboardFilterer) WatchJobAdded(opts *bind.WatchOpts, sink chan<- *DashboardJobAdded) (event.Subscription, error) {

	logs, sub, err := _Dashboard.contract.WatchLogs(opts, "JobAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DashboardJobAdded)
				if err := _Dashboard.contract.UnpackLog(event, "JobAdded", log); err != nil {
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
func (_Dashboard *DashboardFilterer) ParseJobAdded(log types.Log) (*DashboardJobAdded, error) {
	event := new(DashboardJobAdded)
	if err := _Dashboard.contract.UnpackLog(event, "JobAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DashboardReturnValueUpdatedIterator is returned from FilterReturnValueUpdated and is used to iterate over the raw logs and unpacked data for ReturnValueUpdated events raised by the Dashboard contract.
type DashboardReturnValueUpdatedIterator struct {
	Event *DashboardReturnValueUpdated // Event containing the contract specifics and raw log

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
func (it *DashboardReturnValueUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DashboardReturnValueUpdated)
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
		it.Event = new(DashboardReturnValueUpdated)
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
func (it *DashboardReturnValueUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DashboardReturnValueUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DashboardReturnValueUpdated represents a ReturnValueUpdated event raised by the Dashboard contract.
type DashboardReturnValueUpdated struct {
	Key      string
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReturnValueUpdated is a free log retrieval operation binding the contract event 0x67b065c9e8250eb9d63e7fdb9d3ca3c6def37194057697d52d21c0234bcac54e.
//
// Solidity: event ReturnValueUpdated(string key, int256 newValue)
func (_Dashboard *DashboardFilterer) FilterReturnValueUpdated(opts *bind.FilterOpts) (*DashboardReturnValueUpdatedIterator, error) {

	logs, sub, err := _Dashboard.contract.FilterLogs(opts, "ReturnValueUpdated")
	if err != nil {
		return nil, err
	}
	return &DashboardReturnValueUpdatedIterator{contract: _Dashboard.contract, event: "ReturnValueUpdated", logs: logs, sub: sub}, nil
}

// WatchReturnValueUpdated is a free log subscription operation binding the contract event 0x67b065c9e8250eb9d63e7fdb9d3ca3c6def37194057697d52d21c0234bcac54e.
//
// Solidity: event ReturnValueUpdated(string key, int256 newValue)
func (_Dashboard *DashboardFilterer) WatchReturnValueUpdated(opts *bind.WatchOpts, sink chan<- *DashboardReturnValueUpdated) (event.Subscription, error) {

	logs, sub, err := _Dashboard.contract.WatchLogs(opts, "ReturnValueUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DashboardReturnValueUpdated)
				if err := _Dashboard.contract.UnpackLog(event, "ReturnValueUpdated", log); err != nil {
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
func (_Dashboard *DashboardFilterer) ParseReturnValueUpdated(log types.Log) (*DashboardReturnValueUpdated, error) {
	event := new(DashboardReturnValueUpdated)
	if err := _Dashboard.contract.UnpackLog(event, "ReturnValueUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DashboardServiceStatusUpdatedIterator is returned from FilterServiceStatusUpdated and is used to iterate over the raw logs and unpacked data for ServiceStatusUpdated events raised by the Dashboard contract.
type DashboardServiceStatusUpdatedIterator struct {
	Event *DashboardServiceStatusUpdated // Event containing the contract specifics and raw log

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
func (it *DashboardServiceStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DashboardServiceStatusUpdated)
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
		it.Event = new(DashboardServiceStatusUpdated)
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
func (it *DashboardServiceStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DashboardServiceStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DashboardServiceStatusUpdated represents a ServiceStatusUpdated event raised by the Dashboard contract.
type DashboardServiceStatusUpdated struct {
	Index     *big.Int
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterServiceStatusUpdated is a free log retrieval operation binding the contract event 0x5788bd9e26a75de76695c2a82e01198faadcc1c36aa422746f16195be540853a.
//
// Solidity: event ServiceStatusUpdated(uint256 index, uint8 newStatus)
func (_Dashboard *DashboardFilterer) FilterServiceStatusUpdated(opts *bind.FilterOpts) (*DashboardServiceStatusUpdatedIterator, error) {

	logs, sub, err := _Dashboard.contract.FilterLogs(opts, "ServiceStatusUpdated")
	if err != nil {
		return nil, err
	}
	return &DashboardServiceStatusUpdatedIterator{contract: _Dashboard.contract, event: "ServiceStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchServiceStatusUpdated is a free log subscription operation binding the contract event 0x5788bd9e26a75de76695c2a82e01198faadcc1c36aa422746f16195be540853a.
//
// Solidity: event ServiceStatusUpdated(uint256 index, uint8 newStatus)
func (_Dashboard *DashboardFilterer) WatchServiceStatusUpdated(opts *bind.WatchOpts, sink chan<- *DashboardServiceStatusUpdated) (event.Subscription, error) {

	logs, sub, err := _Dashboard.contract.WatchLogs(opts, "ServiceStatusUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DashboardServiceStatusUpdated)
				if err := _Dashboard.contract.UnpackLog(event, "ServiceStatusUpdated", log); err != nil {
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
func (_Dashboard *DashboardFilterer) ParseServiceStatusUpdated(log types.Log) (*DashboardServiceStatusUpdated, error) {
	event := new(DashboardServiceStatusUpdated)
	if err := _Dashboard.contract.UnpackLog(event, "ServiceStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DashboardSingleValueUpdatedIterator is returned from FilterSingleValueUpdated and is used to iterate over the raw logs and unpacked data for SingleValueUpdated events raised by the Dashboard contract.
type DashboardSingleValueUpdatedIterator struct {
	Event *DashboardSingleValueUpdated // Event containing the contract specifics and raw log

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
func (it *DashboardSingleValueUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DashboardSingleValueUpdated)
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
		it.Event = new(DashboardSingleValueUpdated)
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
func (it *DashboardSingleValueUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DashboardSingleValueUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DashboardSingleValueUpdated represents a SingleValueUpdated event raised by the Dashboard contract.
type DashboardSingleValueUpdated struct {
	Key      string
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSingleValueUpdated is a free log retrieval operation binding the contract event 0xfff99ebe4547d53734136b165e336a3484c50aa57fd03b2af7015323b4cb8aea.
//
// Solidity: event SingleValueUpdated(string key, uint256 newValue)
func (_Dashboard *DashboardFilterer) FilterSingleValueUpdated(opts *bind.FilterOpts) (*DashboardSingleValueUpdatedIterator, error) {

	logs, sub, err := _Dashboard.contract.FilterLogs(opts, "SingleValueUpdated")
	if err != nil {
		return nil, err
	}
	return &DashboardSingleValueUpdatedIterator{contract: _Dashboard.contract, event: "SingleValueUpdated", logs: logs, sub: sub}, nil
}

// WatchSingleValueUpdated is a free log subscription operation binding the contract event 0xfff99ebe4547d53734136b165e336a3484c50aa57fd03b2af7015323b4cb8aea.
//
// Solidity: event SingleValueUpdated(string key, uint256 newValue)
func (_Dashboard *DashboardFilterer) WatchSingleValueUpdated(opts *bind.WatchOpts, sink chan<- *DashboardSingleValueUpdated) (event.Subscription, error) {

	logs, sub, err := _Dashboard.contract.WatchLogs(opts, "SingleValueUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DashboardSingleValueUpdated)
				if err := _Dashboard.contract.UnpackLog(event, "SingleValueUpdated", log); err != nil {
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
func (_Dashboard *DashboardFilterer) ParseSingleValueUpdated(log types.Log) (*DashboardSingleValueUpdated, error) {
	event := new(DashboardSingleValueUpdated)
	if err := _Dashboard.contract.UnpackLog(event, "SingleValueUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
