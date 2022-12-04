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
)

// BallotAgenda is an auto generated low-level Go binding around an user-defined struct.
type BallotAgenda struct {
	NumberOfMiner *big.Int
	Winner        *big.Int
	TokenURI      string
	Status        uint8
	ChairPerson   common.Address
	Candidates    []*big.Int
	Voters        []BallotVoter
}

// BallotVoter is an auto generated low-level Go binding around an user-defined struct.
type BallotVoter struct {
	AgendaId     *big.Int
	Weight       *big.Int
	Voted        *big.Int
	Value        *big.Int
	VoterAddress common.Address
}

// GEMMetaData contains all meta data concerning the GEM contract.
var GEMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"agendaId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winner\",\"type\":\"uint256\"}],\"name\":\"AgendaClose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"agendaId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"chairPerson\",\"type\":\"address\"}],\"name\":\"AgendaOpen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"agendaId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"chairPerson\",\"type\":\"address\"}],\"name\":\"AgendaRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agendaId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"weightArr\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"addressArr\",\"type\":\"address[]\"}],\"name\":\"BatchRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agendaId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"chairPerson\",\"type\":\"address\"}],\"name\":\"Finish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agendaId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voterAddr\",\"type\":\"address\"}],\"name\":\"Register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agendaId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"chairPerson\",\"type\":\"address\"}],\"name\":\"Remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agendaId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voterAddr\",\"type\":\"address\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nom\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"candidates\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"chairPerson\",\"type\":\"address\"}],\"name\":\"setAgenda\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newuri\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agendaId\",\"type\":\"uint256\"}],\"name\":\"showAgenda\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"numberOfMiner\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winner\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"enumBallot.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"chairPerson\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"candidates\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"agendaId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"}],\"internalType\":\"structBallot.Voter[]\",\"name\":\"voters\",\"type\":\"tuple[]\"}],\"internalType\":\"structBallot.Agenda\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GEMABI is the input ABI used to generate the binding from.
// Deprecated: Use GEMMetaData.ABI instead.
var GEMABI = GEMMetaData.ABI

// GEM is an auto generated Go binding around an Ethereum contract.
type GEM struct {
	GEMCaller     // Read-only binding to the contract
	GEMTransactor // Write-only binding to the contract
	GEMFilterer   // Log filterer for contract events
}

// GEMCaller is an auto generated read-only Go binding around an Ethereum contract.
type GEMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GEMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GEMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GEMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GEMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GEMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GEMSession struct {
	Contract     *GEM              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GEMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GEMCallerSession struct {
	Contract *GEMCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GEMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GEMTransactorSession struct {
	Contract     *GEMTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GEMRaw is an auto generated low-level Go binding around an Ethereum contract.
type GEMRaw struct {
	Contract *GEM // Generic contract binding to access the raw methods on
}

// GEMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GEMCallerRaw struct {
	Contract *GEMCaller // Generic read-only contract binding to access the raw methods on
}

// GEMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GEMTransactorRaw struct {
	Contract *GEMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGEM creates a new instance of GEM, bound to a specific deployed contract.
func NewGEM(address common.Address, backend bind.ContractBackend) (*GEM, error) {
	contract, err := bindGEM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GEM{GEMCaller: GEMCaller{contract: contract}, GEMTransactor: GEMTransactor{contract: contract}, GEMFilterer: GEMFilterer{contract: contract}}, nil
}

// NewGEMCaller creates a new read-only instance of GEM, bound to a specific deployed contract.
func NewGEMCaller(address common.Address, caller bind.ContractCaller) (*GEMCaller, error) {
	contract, err := bindGEM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GEMCaller{contract: contract}, nil
}

// NewGEMTransactor creates a new write-only instance of GEM, bound to a specific deployed contract.
func NewGEMTransactor(address common.Address, transactor bind.ContractTransactor) (*GEMTransactor, error) {
	contract, err := bindGEM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GEMTransactor{contract: contract}, nil
}

// NewGEMFilterer creates a new log filterer instance of GEM, bound to a specific deployed contract.
func NewGEMFilterer(address common.Address, filterer bind.ContractFilterer) (*GEMFilterer, error) {
	contract, err := bindGEM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GEMFilterer{contract: contract}, nil
}

// bindGEM binds a generic wrapper to an already deployed contract.
func bindGEM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GEMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GEM *GEMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GEM.Contract.GEMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GEM *GEMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GEM.Contract.GEMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GEM *GEMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GEM.Contract.GEMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GEM *GEMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GEM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GEM *GEMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GEM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GEM *GEMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GEM.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_GEM *GEMCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GEM.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_GEM *GEMSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _GEM.Contract.BalanceOf(&_GEM.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_GEM *GEMCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _GEM.Contract.BalanceOf(&_GEM.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_GEM *GEMCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _GEM.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_GEM *GEMSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _GEM.Contract.BalanceOfBatch(&_GEM.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_GEM *GEMCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _GEM.Contract.BalanceOfBatch(&_GEM.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_GEM *GEMCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _GEM.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_GEM *GEMSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _GEM.Contract.IsApprovedForAll(&_GEM.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_GEM *GEMCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _GEM.Contract.IsApprovedForAll(&_GEM.CallOpts, account, operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GEM *GEMCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GEM.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GEM *GEMSession) Owner() (common.Address, error) {
	return _GEM.Contract.Owner(&_GEM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GEM *GEMCallerSession) Owner() (common.Address, error) {
	return _GEM.Contract.Owner(&_GEM.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GEM *GEMCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GEM.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GEM *GEMSession) ProxiableUUID() ([32]byte, error) {
	return _GEM.Contract.ProxiableUUID(&_GEM.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GEM *GEMCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GEM.Contract.ProxiableUUID(&_GEM.CallOpts)
}

// ShowAgenda is a free data retrieval call binding the contract method 0xbbfcc856.
//
// Solidity: function showAgenda(uint256 agendaId) view returns((uint256,uint256,string,uint8,address,uint256[],(uint256,uint256,uint256,uint256,address)[]))
func (_GEM *GEMCaller) ShowAgenda(opts *bind.CallOpts, agendaId *big.Int) (BallotAgenda, error) {
	var out []interface{}
	err := _GEM.contract.Call(opts, &out, "showAgenda", agendaId)

	if err != nil {
		return *new(BallotAgenda), err
	}

	out0 := *abi.ConvertType(out[0], new(BallotAgenda)).(*BallotAgenda)

	return out0, err

}

// ShowAgenda is a free data retrieval call binding the contract method 0xbbfcc856.
//
// Solidity: function showAgenda(uint256 agendaId) view returns((uint256,uint256,string,uint8,address,uint256[],(uint256,uint256,uint256,uint256,address)[]))
func (_GEM *GEMSession) ShowAgenda(agendaId *big.Int) (BallotAgenda, error) {
	return _GEM.Contract.ShowAgenda(&_GEM.CallOpts, agendaId)
}

// ShowAgenda is a free data retrieval call binding the contract method 0xbbfcc856.
//
// Solidity: function showAgenda(uint256 agendaId) view returns((uint256,uint256,string,uint8,address,uint256[],(uint256,uint256,uint256,uint256,address)[]))
func (_GEM *GEMCallerSession) ShowAgenda(agendaId *big.Int) (BallotAgenda, error) {
	return _GEM.Contract.ShowAgenda(&_GEM.CallOpts, agendaId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GEM *GEMCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GEM.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GEM *GEMSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GEM.Contract.SupportsInterface(&_GEM.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GEM *GEMCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GEM.Contract.SupportsInterface(&_GEM.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_GEM *GEMCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _GEM.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_GEM *GEMSession) Uri(arg0 *big.Int) (string, error) {
	return _GEM.Contract.Uri(&_GEM.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_GEM *GEMCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _GEM.Contract.Uri(&_GEM.CallOpts, arg0)
}

// BatchRegister is a paid mutator transaction binding the contract method 0xf2ceb3f0.
//
// Solidity: function BatchRegister(uint256 agendaId, uint256[] weightArr, address[] addressArr) returns()
func (_GEM *GEMTransactor) BatchRegister(opts *bind.TransactOpts, agendaId *big.Int, weightArr []*big.Int, addressArr []common.Address) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "BatchRegister", agendaId, weightArr, addressArr)
}

// BatchRegister is a paid mutator transaction binding the contract method 0xf2ceb3f0.
//
// Solidity: function BatchRegister(uint256 agendaId, uint256[] weightArr, address[] addressArr) returns()
func (_GEM *GEMSession) BatchRegister(agendaId *big.Int, weightArr []*big.Int, addressArr []common.Address) (*types.Transaction, error) {
	return _GEM.Contract.BatchRegister(&_GEM.TransactOpts, agendaId, weightArr, addressArr)
}

// BatchRegister is a paid mutator transaction binding the contract method 0xf2ceb3f0.
//
// Solidity: function BatchRegister(uint256 agendaId, uint256[] weightArr, address[] addressArr) returns()
func (_GEM *GEMTransactorSession) BatchRegister(agendaId *big.Int, weightArr []*big.Int, addressArr []common.Address) (*types.Transaction, error) {
	return _GEM.Contract.BatchRegister(&_GEM.TransactOpts, agendaId, weightArr, addressArr)
}

// Finish is a paid mutator transaction binding the contract method 0xcd6c57d9.
//
// Solidity: function Finish(uint256 agendaId, address chairPerson) returns()
func (_GEM *GEMTransactor) Finish(opts *bind.TransactOpts, agendaId *big.Int, chairPerson common.Address) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "Finish", agendaId, chairPerson)
}

// Finish is a paid mutator transaction binding the contract method 0xcd6c57d9.
//
// Solidity: function Finish(uint256 agendaId, address chairPerson) returns()
func (_GEM *GEMSession) Finish(agendaId *big.Int, chairPerson common.Address) (*types.Transaction, error) {
	return _GEM.Contract.Finish(&_GEM.TransactOpts, agendaId, chairPerson)
}

// Finish is a paid mutator transaction binding the contract method 0xcd6c57d9.
//
// Solidity: function Finish(uint256 agendaId, address chairPerson) returns()
func (_GEM *GEMTransactorSession) Finish(agendaId *big.Int, chairPerson common.Address) (*types.Transaction, error) {
	return _GEM.Contract.Finish(&_GEM.TransactOpts, agendaId, chairPerson)
}

// Register is a paid mutator transaction binding the contract method 0x8c0cb329.
//
// Solidity: function Register(uint256 agendaId, uint256 weight, address voterAddr) returns()
func (_GEM *GEMTransactor) Register(opts *bind.TransactOpts, agendaId *big.Int, weight *big.Int, voterAddr common.Address) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "Register", agendaId, weight, voterAddr)
}

// Register is a paid mutator transaction binding the contract method 0x8c0cb329.
//
// Solidity: function Register(uint256 agendaId, uint256 weight, address voterAddr) returns()
func (_GEM *GEMSession) Register(agendaId *big.Int, weight *big.Int, voterAddr common.Address) (*types.Transaction, error) {
	return _GEM.Contract.Register(&_GEM.TransactOpts, agendaId, weight, voterAddr)
}

// Register is a paid mutator transaction binding the contract method 0x8c0cb329.
//
// Solidity: function Register(uint256 agendaId, uint256 weight, address voterAddr) returns()
func (_GEM *GEMTransactorSession) Register(agendaId *big.Int, weight *big.Int, voterAddr common.Address) (*types.Transaction, error) {
	return _GEM.Contract.Register(&_GEM.TransactOpts, agendaId, weight, voterAddr)
}

// Remove is a paid mutator transaction binding the contract method 0xe52b7042.
//
// Solidity: function Remove(uint256 agendaId, address chairPerson) returns()
func (_GEM *GEMTransactor) Remove(opts *bind.TransactOpts, agendaId *big.Int, chairPerson common.Address) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "Remove", agendaId, chairPerson)
}

// Remove is a paid mutator transaction binding the contract method 0xe52b7042.
//
// Solidity: function Remove(uint256 agendaId, address chairPerson) returns()
func (_GEM *GEMSession) Remove(agendaId *big.Int, chairPerson common.Address) (*types.Transaction, error) {
	return _GEM.Contract.Remove(&_GEM.TransactOpts, agendaId, chairPerson)
}

// Remove is a paid mutator transaction binding the contract method 0xe52b7042.
//
// Solidity: function Remove(uint256 agendaId, address chairPerson) returns()
func (_GEM *GEMTransactorSession) Remove(agendaId *big.Int, chairPerson common.Address) (*types.Transaction, error) {
	return _GEM.Contract.Remove(&_GEM.TransactOpts, agendaId, chairPerson)
}

// Vote is a paid mutator transaction binding the contract method 0x7c2f48a4.
//
// Solidity: function Vote(uint256 agendaId, uint256 value, address voterAddr) returns()
func (_GEM *GEMTransactor) Vote(opts *bind.TransactOpts, agendaId *big.Int, value *big.Int, voterAddr common.Address) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "Vote", agendaId, value, voterAddr)
}

// Vote is a paid mutator transaction binding the contract method 0x7c2f48a4.
//
// Solidity: function Vote(uint256 agendaId, uint256 value, address voterAddr) returns()
func (_GEM *GEMSession) Vote(agendaId *big.Int, value *big.Int, voterAddr common.Address) (*types.Transaction, error) {
	return _GEM.Contract.Vote(&_GEM.TransactOpts, agendaId, value, voterAddr)
}

// Vote is a paid mutator transaction binding the contract method 0x7c2f48a4.
//
// Solidity: function Vote(uint256 agendaId, uint256 value, address voterAddr) returns()
func (_GEM *GEMTransactorSession) Vote(agendaId *big.Int, value *big.Int, voterAddr common.Address) (*types.Transaction, error) {
	return _GEM.Contract.Vote(&_GEM.TransactOpts, agendaId, value, voterAddr)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_GEM *GEMTransactor) Burn(opts *bind.TransactOpts, account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "burn", account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_GEM *GEMSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _GEM.Contract.Burn(&_GEM.TransactOpts, account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_GEM *GEMTransactorSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _GEM.Contract.Burn(&_GEM.TransactOpts, account, id, value)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_GEM *GEMTransactor) BurnBatch(opts *bind.TransactOpts, account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "burnBatch", account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_GEM *GEMSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _GEM.Contract.BurnBatch(&_GEM.TransactOpts, account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_GEM *GEMTransactorSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _GEM.Contract.BurnBatch(&_GEM.TransactOpts, account, ids, values)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GEM *GEMTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GEM *GEMSession) Initialize() (*types.Transaction, error) {
	return _GEM.Contract.Initialize(&_GEM.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GEM *GEMTransactorSession) Initialize() (*types.Transaction, error) {
	return _GEM.Contract.Initialize(&_GEM.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address account, uint256 id, uint256 amount, bytes data) returns()
func (_GEM *GEMTransactor) Mint(opts *bind.TransactOpts, account common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "mint", account, id, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address account, uint256 id, uint256 amount, bytes data) returns()
func (_GEM *GEMSession) Mint(account common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GEM.Contract.Mint(&_GEM.TransactOpts, account, id, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address account, uint256 id, uint256 amount, bytes data) returns()
func (_GEM *GEMTransactorSession) Mint(account common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GEM.Contract.Mint(&_GEM.TransactOpts, account, id, amount, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GEM *GEMTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GEM *GEMSession) RenounceOwnership() (*types.Transaction, error) {
	return _GEM.Contract.RenounceOwnership(&_GEM.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GEM *GEMTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GEM.Contract.RenounceOwnership(&_GEM.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_GEM *GEMTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_GEM *GEMSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _GEM.Contract.SafeBatchTransferFrom(&_GEM.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_GEM *GEMTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _GEM.Contract.SafeBatchTransferFrom(&_GEM.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_GEM *GEMTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_GEM *GEMSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GEM.Contract.SafeTransferFrom(&_GEM.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_GEM *GEMTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _GEM.Contract.SafeTransferFrom(&_GEM.TransactOpts, from, to, id, amount, data)
}

// SetAgenda is a paid mutator transaction binding the contract method 0x72f3a2af.
//
// Solidity: function setAgenda(uint256 nom, uint256[] candidates, string tokenURI, address chairPerson) returns()
func (_GEM *GEMTransactor) SetAgenda(opts *bind.TransactOpts, nom *big.Int, candidates []*big.Int, tokenURI string, chairPerson common.Address) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "setAgenda", nom, candidates, tokenURI, chairPerson)
}

// SetAgenda is a paid mutator transaction binding the contract method 0x72f3a2af.
//
// Solidity: function setAgenda(uint256 nom, uint256[] candidates, string tokenURI, address chairPerson) returns()
func (_GEM *GEMSession) SetAgenda(nom *big.Int, candidates []*big.Int, tokenURI string, chairPerson common.Address) (*types.Transaction, error) {
	return _GEM.Contract.SetAgenda(&_GEM.TransactOpts, nom, candidates, tokenURI, chairPerson)
}

// SetAgenda is a paid mutator transaction binding the contract method 0x72f3a2af.
//
// Solidity: function setAgenda(uint256 nom, uint256[] candidates, string tokenURI, address chairPerson) returns()
func (_GEM *GEMTransactorSession) SetAgenda(nom *big.Int, candidates []*big.Int, tokenURI string, chairPerson common.Address) (*types.Transaction, error) {
	return _GEM.Contract.SetAgenda(&_GEM.TransactOpts, nom, candidates, tokenURI, chairPerson)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_GEM *GEMTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_GEM *GEMSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _GEM.Contract.SetApprovalForAll(&_GEM.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_GEM *GEMTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _GEM.Contract.SetApprovalForAll(&_GEM.TransactOpts, operator, approved)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newuri) returns()
func (_GEM *GEMTransactor) SetURI(opts *bind.TransactOpts, newuri string) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "setURI", newuri)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newuri) returns()
func (_GEM *GEMSession) SetURI(newuri string) (*types.Transaction, error) {
	return _GEM.Contract.SetURI(&_GEM.TransactOpts, newuri)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string newuri) returns()
func (_GEM *GEMTransactorSession) SetURI(newuri string) (*types.Transaction, error) {
	return _GEM.Contract.SetURI(&_GEM.TransactOpts, newuri)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GEM *GEMTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GEM *GEMSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GEM.Contract.TransferOwnership(&_GEM.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GEM *GEMTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GEM.Contract.TransferOwnership(&_GEM.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GEM *GEMTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GEM *GEMSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GEM.Contract.UpgradeTo(&_GEM.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GEM *GEMTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GEM.Contract.UpgradeTo(&_GEM.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GEM *GEMTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GEM.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GEM *GEMSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GEM.Contract.UpgradeToAndCall(&_GEM.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GEM *GEMTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GEM.Contract.UpgradeToAndCall(&_GEM.TransactOpts, newImplementation, data)
}

// GEMAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the GEM contract.
type GEMAdminChangedIterator struct {
	Event *GEMAdminChanged // Event containing the contract specifics and raw log

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
func (it *GEMAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GEMAdminChanged)
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
		it.Event = new(GEMAdminChanged)
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
func (it *GEMAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GEMAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GEMAdminChanged represents a AdminChanged event raised by the GEM contract.
type GEMAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GEM *GEMFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*GEMAdminChangedIterator, error) {

	logs, sub, err := _GEM.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &GEMAdminChangedIterator{contract: _GEM.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GEM *GEMFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *GEMAdminChanged) (event.Subscription, error) {

	logs, sub, err := _GEM.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GEMAdminChanged)
				if err := _GEM.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GEM *GEMFilterer) ParseAdminChanged(log types.Log) (*GEMAdminChanged, error) {
	event := new(GEMAdminChanged)
	if err := _GEM.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GEMAgendaCloseIterator is returned from FilterAgendaClose and is used to iterate over the raw logs and unpacked data for AgendaClose events raised by the GEM contract.
type GEMAgendaCloseIterator struct {
	Event *GEMAgendaClose // Event containing the contract specifics and raw log

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
func (it *GEMAgendaCloseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GEMAgendaClose)
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
		it.Event = new(GEMAgendaClose)
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
func (it *GEMAgendaCloseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GEMAgendaCloseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GEMAgendaClose represents a AgendaClose event raised by the GEM contract.
type GEMAgendaClose struct {
	AgendaId *big.Int
	Winner   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAgendaClose is a free log retrieval operation binding the contract event 0x67041e01004ba8929eb4a065275e8e631a725bca06aa6e39adc31171c60e34a7.
//
// Solidity: event AgendaClose(uint256 agendaId, uint256 winner)
func (_GEM *GEMFilterer) FilterAgendaClose(opts *bind.FilterOpts) (*GEMAgendaCloseIterator, error) {

	logs, sub, err := _GEM.contract.FilterLogs(opts, "AgendaClose")
	if err != nil {
		return nil, err
	}
	return &GEMAgendaCloseIterator{contract: _GEM.contract, event: "AgendaClose", logs: logs, sub: sub}, nil
}

// WatchAgendaClose is a free log subscription operation binding the contract event 0x67041e01004ba8929eb4a065275e8e631a725bca06aa6e39adc31171c60e34a7.
//
// Solidity: event AgendaClose(uint256 agendaId, uint256 winner)
func (_GEM *GEMFilterer) WatchAgendaClose(opts *bind.WatchOpts, sink chan<- *GEMAgendaClose) (event.Subscription, error) {

	logs, sub, err := _GEM.contract.WatchLogs(opts, "AgendaClose")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GEMAgendaClose)
				if err := _GEM.contract.UnpackLog(event, "AgendaClose", log); err != nil {
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

// ParseAgendaClose is a log parse operation binding the contract event 0x67041e01004ba8929eb4a065275e8e631a725bca06aa6e39adc31171c60e34a7.
//
// Solidity: event AgendaClose(uint256 agendaId, uint256 winner)
func (_GEM *GEMFilterer) ParseAgendaClose(log types.Log) (*GEMAgendaClose, error) {
	event := new(GEMAgendaClose)
	if err := _GEM.contract.UnpackLog(event, "AgendaClose", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GEMAgendaOpenIterator is returned from FilterAgendaOpen and is used to iterate over the raw logs and unpacked data for AgendaOpen events raised by the GEM contract.
type GEMAgendaOpenIterator struct {
	Event *GEMAgendaOpen // Event containing the contract specifics and raw log

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
func (it *GEMAgendaOpenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GEMAgendaOpen)
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
		it.Event = new(GEMAgendaOpen)
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
func (it *GEMAgendaOpenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GEMAgendaOpenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GEMAgendaOpen represents a AgendaOpen event raised by the GEM contract.
type GEMAgendaOpen struct {
	AgendaId    *big.Int
	ChairPerson common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAgendaOpen is a free log retrieval operation binding the contract event 0xa247795eeea00a81689afcbbd7727dd7c79641f141156ecba0bac4d81c0654a2.
//
// Solidity: event AgendaOpen(uint256 agendaId, address chairPerson)
func (_GEM *GEMFilterer) FilterAgendaOpen(opts *bind.FilterOpts) (*GEMAgendaOpenIterator, error) {

	logs, sub, err := _GEM.contract.FilterLogs(opts, "AgendaOpen")
	if err != nil {
		return nil, err
	}
	return &GEMAgendaOpenIterator{contract: _GEM.contract, event: "AgendaOpen", logs: logs, sub: sub}, nil
}

// WatchAgendaOpen is a free log subscription operation binding the contract event 0xa247795eeea00a81689afcbbd7727dd7c79641f141156ecba0bac4d81c0654a2.
//
// Solidity: event AgendaOpen(uint256 agendaId, address chairPerson)
func (_GEM *GEMFilterer) WatchAgendaOpen(opts *bind.WatchOpts, sink chan<- *GEMAgendaOpen) (event.Subscription, error) {

	logs, sub, err := _GEM.contract.WatchLogs(opts, "AgendaOpen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GEMAgendaOpen)
				if err := _GEM.contract.UnpackLog(event, "AgendaOpen", log); err != nil {
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

// ParseAgendaOpen is a log parse operation binding the contract event 0xa247795eeea00a81689afcbbd7727dd7c79641f141156ecba0bac4d81c0654a2.
//
// Solidity: event AgendaOpen(uint256 agendaId, address chairPerson)
func (_GEM *GEMFilterer) ParseAgendaOpen(log types.Log) (*GEMAgendaOpen, error) {
	event := new(GEMAgendaOpen)
	if err := _GEM.contract.UnpackLog(event, "AgendaOpen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GEMAgendaRemovedIterator is returned from FilterAgendaRemoved and is used to iterate over the raw logs and unpacked data for AgendaRemoved events raised by the GEM contract.
type GEMAgendaRemovedIterator struct {
	Event *GEMAgendaRemoved // Event containing the contract specifics and raw log

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
func (it *GEMAgendaRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GEMAgendaRemoved)
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
		it.Event = new(GEMAgendaRemoved)
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
func (it *GEMAgendaRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GEMAgendaRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GEMAgendaRemoved represents a AgendaRemoved event raised by the GEM contract.
type GEMAgendaRemoved struct {
	AgendaId    *big.Int
	ChairPerson common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAgendaRemoved is a free log retrieval operation binding the contract event 0x33d6926d6e9005876915395283137e084530010df2227be5855c0c949d19b3c9.
//
// Solidity: event AgendaRemoved(uint256 agendaId, address chairPerson)
func (_GEM *GEMFilterer) FilterAgendaRemoved(opts *bind.FilterOpts) (*GEMAgendaRemovedIterator, error) {

	logs, sub, err := _GEM.contract.FilterLogs(opts, "AgendaRemoved")
	if err != nil {
		return nil, err
	}
	return &GEMAgendaRemovedIterator{contract: _GEM.contract, event: "AgendaRemoved", logs: logs, sub: sub}, nil
}

// WatchAgendaRemoved is a free log subscription operation binding the contract event 0x33d6926d6e9005876915395283137e084530010df2227be5855c0c949d19b3c9.
//
// Solidity: event AgendaRemoved(uint256 agendaId, address chairPerson)
func (_GEM *GEMFilterer) WatchAgendaRemoved(opts *bind.WatchOpts, sink chan<- *GEMAgendaRemoved) (event.Subscription, error) {

	logs, sub, err := _GEM.contract.WatchLogs(opts, "AgendaRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GEMAgendaRemoved)
				if err := _GEM.contract.UnpackLog(event, "AgendaRemoved", log); err != nil {
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

// ParseAgendaRemoved is a log parse operation binding the contract event 0x33d6926d6e9005876915395283137e084530010df2227be5855c0c949d19b3c9.
//
// Solidity: event AgendaRemoved(uint256 agendaId, address chairPerson)
func (_GEM *GEMFilterer) ParseAgendaRemoved(log types.Log) (*GEMAgendaRemoved, error) {
	event := new(GEMAgendaRemoved)
	if err := _GEM.contract.UnpackLog(event, "AgendaRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GEMApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the GEM contract.
type GEMApprovalForAllIterator struct {
	Event *GEMApprovalForAll // Event containing the contract specifics and raw log

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
func (it *GEMApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GEMApprovalForAll)
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
		it.Event = new(GEMApprovalForAll)
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
func (it *GEMApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GEMApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GEMApprovalForAll represents a ApprovalForAll event raised by the GEM contract.
type GEMApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_GEM *GEMFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*GEMApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _GEM.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &GEMApprovalForAllIterator{contract: _GEM.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_GEM *GEMFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *GEMApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _GEM.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GEMApprovalForAll)
				if err := _GEM.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_GEM *GEMFilterer) ParseApprovalForAll(log types.Log) (*GEMApprovalForAll, error) {
	event := new(GEMApprovalForAll)
	if err := _GEM.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GEMBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the GEM contract.
type GEMBeaconUpgradedIterator struct {
	Event *GEMBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *GEMBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GEMBeaconUpgraded)
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
		it.Event = new(GEMBeaconUpgraded)
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
func (it *GEMBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GEMBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GEMBeaconUpgraded represents a BeaconUpgraded event raised by the GEM contract.
type GEMBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GEM *GEMFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*GEMBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _GEM.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &GEMBeaconUpgradedIterator{contract: _GEM.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GEM *GEMFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *GEMBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _GEM.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GEMBeaconUpgraded)
				if err := _GEM.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GEM *GEMFilterer) ParseBeaconUpgraded(log types.Log) (*GEMBeaconUpgraded, error) {
	event := new(GEMBeaconUpgraded)
	if err := _GEM.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GEMInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GEM contract.
type GEMInitializedIterator struct {
	Event *GEMInitialized // Event containing the contract specifics and raw log

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
func (it *GEMInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GEMInitialized)
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
		it.Event = new(GEMInitialized)
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
func (it *GEMInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GEMInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GEMInitialized represents a Initialized event raised by the GEM contract.
type GEMInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GEM *GEMFilterer) FilterInitialized(opts *bind.FilterOpts) (*GEMInitializedIterator, error) {

	logs, sub, err := _GEM.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GEMInitializedIterator{contract: _GEM.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GEM *GEMFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GEMInitialized) (event.Subscription, error) {

	logs, sub, err := _GEM.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GEMInitialized)
				if err := _GEM.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GEM *GEMFilterer) ParseInitialized(log types.Log) (*GEMInitialized, error) {
	event := new(GEMInitialized)
	if err := _GEM.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GEMOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GEM contract.
type GEMOwnershipTransferredIterator struct {
	Event *GEMOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GEMOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GEMOwnershipTransferred)
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
		it.Event = new(GEMOwnershipTransferred)
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
func (it *GEMOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GEMOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GEMOwnershipTransferred represents a OwnershipTransferred event raised by the GEM contract.
type GEMOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GEM *GEMFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GEMOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GEM.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GEMOwnershipTransferredIterator{contract: _GEM.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GEM *GEMFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GEMOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GEM.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GEMOwnershipTransferred)
				if err := _GEM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GEM *GEMFilterer) ParseOwnershipTransferred(log types.Log) (*GEMOwnershipTransferred, error) {
	event := new(GEMOwnershipTransferred)
	if err := _GEM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GEMTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the GEM contract.
type GEMTransferBatchIterator struct {
	Event *GEMTransferBatch // Event containing the contract specifics and raw log

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
func (it *GEMTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GEMTransferBatch)
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
		it.Event = new(GEMTransferBatch)
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
func (it *GEMTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GEMTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GEMTransferBatch represents a TransferBatch event raised by the GEM contract.
type GEMTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_GEM *GEMFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*GEMTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GEM.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GEMTransferBatchIterator{contract: _GEM.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_GEM *GEMFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *GEMTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GEM.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GEMTransferBatch)
				if err := _GEM.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_GEM *GEMFilterer) ParseTransferBatch(log types.Log) (*GEMTransferBatch, error) {
	event := new(GEMTransferBatch)
	if err := _GEM.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GEMTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the GEM contract.
type GEMTransferSingleIterator struct {
	Event *GEMTransferSingle // Event containing the contract specifics and raw log

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
func (it *GEMTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GEMTransferSingle)
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
		it.Event = new(GEMTransferSingle)
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
func (it *GEMTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GEMTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GEMTransferSingle represents a TransferSingle event raised by the GEM contract.
type GEMTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_GEM *GEMFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*GEMTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GEM.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GEMTransferSingleIterator{contract: _GEM.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_GEM *GEMFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *GEMTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GEM.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GEMTransferSingle)
				if err := _GEM.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_GEM *GEMFilterer) ParseTransferSingle(log types.Log) (*GEMTransferSingle, error) {
	event := new(GEMTransferSingle)
	if err := _GEM.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GEMURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the GEM contract.
type GEMURIIterator struct {
	Event *GEMURI // Event containing the contract specifics and raw log

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
func (it *GEMURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GEMURI)
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
		it.Event = new(GEMURI)
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
func (it *GEMURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GEMURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GEMURI represents a URI event raised by the GEM contract.
type GEMURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_GEM *GEMFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*GEMURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _GEM.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &GEMURIIterator{contract: _GEM.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_GEM *GEMFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *GEMURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _GEM.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GEMURI)
				if err := _GEM.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_GEM *GEMFilterer) ParseURI(log types.Log) (*GEMURI, error) {
	event := new(GEMURI)
	if err := _GEM.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GEMUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GEM contract.
type GEMUpgradedIterator struct {
	Event *GEMUpgraded // Event containing the contract specifics and raw log

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
func (it *GEMUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GEMUpgraded)
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
		it.Event = new(GEMUpgraded)
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
func (it *GEMUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GEMUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GEMUpgraded represents a Upgraded event raised by the GEM contract.
type GEMUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GEM *GEMFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GEMUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GEM.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GEMUpgradedIterator{contract: _GEM.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GEM *GEMFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GEMUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GEM.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GEMUpgraded)
				if err := _GEM.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GEM *GEMFilterer) ParseUpgraded(log types.Log) (*GEMUpgraded, error) {
	event := new(GEMUpgraded)
	if err := _GEM.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
