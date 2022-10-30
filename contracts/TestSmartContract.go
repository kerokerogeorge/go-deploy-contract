// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storage

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

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getMood\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_mood\",\"type\":\"string\"}],\"name\":\"setMood\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610762806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80632e64cec1146100515780635f3cbff51461006f5780636057361d1461008b5780639d0c1397146100a7575b600080fd5b6100596100c5565b6040516100669190610196565b60405180910390f35b6100896004803603810190610084919061030b565b6100ce565b005b6100a560048036038101906100a09190610380565b6100e1565b005b6100af6100eb565b6040516100bc919061042c565b60405180910390f35b60008054905090565b80600190816100dd919061065a565b5050565b8060008190555050565b6060600180546100fa9061047d565b80601f01602080910402602001604051908101604052809291908181526020018280546101269061047d565b80156101735780601f1061014857610100808354040283529160200191610173565b820191906000526020600020905b81548152906001019060200180831161015657829003601f168201915b5050505050905090565b6000819050919050565b6101908161017d565b82525050565b60006020820190506101ab6000830184610187565b92915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610218826101cf565b810181811067ffffffffffffffff82111715610237576102366101e0565b5b80604052505050565b600061024a6101b1565b9050610256828261020f565b919050565b600067ffffffffffffffff821115610276576102756101e0565b5b61027f826101cf565b9050602081019050919050565b82818337600083830152505050565b60006102ae6102a98461025b565b610240565b9050828152602081018484840111156102ca576102c96101ca565b5b6102d584828561028c565b509392505050565b600082601f8301126102f2576102f16101c5565b5b813561030284826020860161029b565b91505092915050565b600060208284031215610321576103206101bb565b5b600082013567ffffffffffffffff81111561033f5761033e6101c0565b5b61034b848285016102dd565b91505092915050565b61035d8161017d565b811461036857600080fd5b50565b60008135905061037a81610354565b92915050565b600060208284031215610396576103956101bb565b5b60006103a48482850161036b565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156103e75780820151818401526020810190506103cc565b60008484015250505050565b60006103fe826103ad565b61040881856103b8565b93506104188185602086016103c9565b610421816101cf565b840191505092915050565b6000602082019050818103600083015261044681846103f3565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061049557607f821691505b6020821081036104a8576104a761044e565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026105107fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826104d3565b61051a86836104d3565b95508019841693508086168417925050509392505050565b6000819050919050565b600061055761055261054d8461017d565b610532565b61017d565b9050919050565b6000819050919050565b6105718361053c565b61058561057d8261055e565b8484546104e0565b825550505050565b600090565b61059a61058d565b6105a5818484610568565b505050565b5b818110156105c9576105be600082610592565b6001810190506105ab565b5050565b601f82111561060e576105df816104ae565b6105e8846104c3565b810160208510156105f7578190505b61060b610603856104c3565b8301826105aa565b50505b505050565b600082821c905092915050565b600061063160001984600802610613565b1980831691505092915050565b600061064a8383610620565b9150826002028217905092915050565b610663826103ad565b67ffffffffffffffff81111561067c5761067b6101e0565b5b610686825461047d565b6106918282856105cd565b600060209050601f8311600181146106c457600084156106b2578287015190505b6106bc858261063e565b865550610724565b601f1984166106d2866104ae565b60005b828110156106fa578489015182556001820191506020850194506020810190506106d5565b868310156107175784890151610713601f891682610620565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220d940fe0c70d45da4e5cf96a63909f3bdb816f641ed72dd7e7ce1ae45781e2efc64736f6c63430008110033",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// StorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageMetaData.Bin instead.
var StorageBin = StorageMetaData.Bin

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// GetMood is a free data retrieval call binding the contract method 0x9d0c1397.
//
// Solidity: function getMood() view returns(string)
func (_Storage *StorageCaller) GetMood(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getMood")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMood is a free data retrieval call binding the contract method 0x9d0c1397.
//
// Solidity: function getMood() view returns(string)
func (_Storage *StorageSession) GetMood() (string, error) {
	return _Storage.Contract.GetMood(&_Storage.CallOpts)
}

// GetMood is a free data retrieval call binding the contract method 0x9d0c1397.
//
// Solidity: function getMood() view returns(string)
func (_Storage *StorageCallerSession) GetMood() (string, error) {
	return _Storage.Contract.GetMood(&_Storage.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Storage *StorageCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Storage *StorageSession) Retrieve() (*big.Int, error) {
	return _Storage.Contract.Retrieve(&_Storage.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Storage *StorageCallerSession) Retrieve() (*big.Int, error) {
	return _Storage.Contract.Retrieve(&_Storage.CallOpts)
}

// SetMood is a paid mutator transaction binding the contract method 0x5f3cbff5.
//
// Solidity: function setMood(string _mood) returns()
func (_Storage *StorageTransactor) SetMood(opts *bind.TransactOpts, _mood string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setMood", _mood)
}

// SetMood is a paid mutator transaction binding the contract method 0x5f3cbff5.
//
// Solidity: function setMood(string _mood) returns()
func (_Storage *StorageSession) SetMood(_mood string) (*types.Transaction, error) {
	return _Storage.Contract.SetMood(&_Storage.TransactOpts, _mood)
}

// SetMood is a paid mutator transaction binding the contract method 0x5f3cbff5.
//
// Solidity: function setMood(string _mood) returns()
func (_Storage *StorageTransactorSession) SetMood(_mood string) (*types.Transaction, error) {
	return _Storage.Contract.SetMood(&_Storage.TransactOpts, _mood)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 number) returns()
func (_Storage *StorageTransactor) Store(opts *bind.TransactOpts, number *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "store", number)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 number) returns()
func (_Storage *StorageSession) Store(number *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Store(&_Storage.TransactOpts, number)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 number) returns()
func (_Storage *StorageTransactorSession) Store(number *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Store(&_Storage.TransactOpts, number)
}
