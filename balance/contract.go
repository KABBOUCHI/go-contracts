// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balance

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

// ResolverBalanceData is an auto generated low-level Go binding around an user-defined struct.
type ResolverBalanceData struct {
	Owner  common.Address
	Tokens []common.Address
}

// ResolverBalanceReturnData is an auto generated low-level Go binding around an user-defined struct.
type ResolverBalanceReturnData struct {
	Owner    common.Address
	Balances []*big.Int
}

// ResolverTokenData is an auto generated low-level Go binding around an user-defined struct.
type ResolverTokenData struct {
	IsToken  bool
	Name     string
	Symbol   string
	Decimals *big.Int
}

// BalanceMetaData contains all meta data concerning the Balance contract.
var BalanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tknAddress\",\"type\":\"address[]\"}],\"name\":\"getAllowances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tknAddress\",\"type\":\"address[]\"}],\"name\":\"getBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getBalancesOfOwners\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"internalType\":\"structResolver.BalanceReturnData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"internalType\":\"structResolver.BalanceData[]\",\"name\":\"datas\",\"type\":\"tuple[]\"}],\"name\":\"getBalancesOfOwnersWithTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"internalType\":\"structResolver.BalanceReturnData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tknAddress\",\"type\":\"address[]\"}],\"name\":\"getTokenDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isToken\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"internalType\":\"structResolver.TokenData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BalanceABI is the input ABI used to generate the binding from.
// Deprecated: Use BalanceMetaData.ABI instead.
var BalanceABI = BalanceMetaData.ABI

// Balance is an auto generated Go binding around an Ethereum contract.
type Balance struct {
	BalanceCaller     // Read-only binding to the contract
	BalanceTransactor // Write-only binding to the contract
	BalanceFilterer   // Log filterer for contract events
}

// BalanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalanceSession struct {
	Contract     *Balance          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalanceCallerSession struct {
	Contract *BalanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BalanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalanceTransactorSession struct {
	Contract     *BalanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BalanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalanceRaw struct {
	Contract *Balance // Generic contract binding to access the raw methods on
}

// BalanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalanceCallerRaw struct {
	Contract *BalanceCaller // Generic read-only contract binding to access the raw methods on
}

// BalanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalanceTransactorRaw struct {
	Contract *BalanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalance creates a new instance of Balance, bound to a specific deployed contract.
func NewBalance(address common.Address, backend bind.ContractBackend) (*Balance, error) {
	contract, err := bindBalance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Balance{BalanceCaller: BalanceCaller{contract: contract}, BalanceTransactor: BalanceTransactor{contract: contract}, BalanceFilterer: BalanceFilterer{contract: contract}}, nil
}

// NewBalanceCaller creates a new read-only instance of Balance, bound to a specific deployed contract.
func NewBalanceCaller(address common.Address, caller bind.ContractCaller) (*BalanceCaller, error) {
	contract, err := bindBalance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceCaller{contract: contract}, nil
}

// NewBalanceTransactor creates a new write-only instance of Balance, bound to a specific deployed contract.
func NewBalanceTransactor(address common.Address, transactor bind.ContractTransactor) (*BalanceTransactor, error) {
	contract, err := bindBalance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceTransactor{contract: contract}, nil
}

// NewBalanceFilterer creates a new log filterer instance of Balance, bound to a specific deployed contract.
func NewBalanceFilterer(address common.Address, filterer bind.ContractFilterer) (*BalanceFilterer, error) {
	contract, err := bindBalance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalanceFilterer{contract: contract}, nil
}

// bindBalance binds a generic wrapper to an already deployed contract.
func bindBalance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balance *BalanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balance.Contract.BalanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balance *BalanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balance.Contract.BalanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balance *BalanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balance.Contract.BalanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balance *BalanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balance *BalanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balance *BalanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balance.Contract.contract.Transact(opts, method, params...)
}

// GetAllowances is a free data retrieval call binding the contract method 0x49c5f1fb.
//
// Solidity: function getAllowances(address owner, address spender, address[] tknAddress) view returns(uint256[])
func (_Balance *BalanceCaller) GetAllowances(opts *bind.CallOpts, owner common.Address, spender common.Address, tknAddress []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Balance.contract.Call(opts, &out, "getAllowances", owner, spender, tknAddress)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllowances is a free data retrieval call binding the contract method 0x49c5f1fb.
//
// Solidity: function getAllowances(address owner, address spender, address[] tknAddress) view returns(uint256[])
func (_Balance *BalanceSession) GetAllowances(owner common.Address, spender common.Address, tknAddress []common.Address) ([]*big.Int, error) {
	return _Balance.Contract.GetAllowances(&_Balance.CallOpts, owner, spender, tknAddress)
}

// GetAllowances is a free data retrieval call binding the contract method 0x49c5f1fb.
//
// Solidity: function getAllowances(address owner, address spender, address[] tknAddress) view returns(uint256[])
func (_Balance *BalanceCallerSession) GetAllowances(owner common.Address, spender common.Address, tknAddress []common.Address) ([]*big.Int, error) {
	return _Balance.Contract.GetAllowances(&_Balance.CallOpts, owner, spender, tknAddress)
}

// GetBalances is a free data retrieval call binding the contract method 0x6a385ae9.
//
// Solidity: function getBalances(address owner, address[] tknAddress) view returns(uint256[])
func (_Balance *BalanceCaller) GetBalances(opts *bind.CallOpts, owner common.Address, tknAddress []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Balance.contract.Call(opts, &out, "getBalances", owner, tknAddress)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x6a385ae9.
//
// Solidity: function getBalances(address owner, address[] tknAddress) view returns(uint256[])
func (_Balance *BalanceSession) GetBalances(owner common.Address, tknAddress []common.Address) ([]*big.Int, error) {
	return _Balance.Contract.GetBalances(&_Balance.CallOpts, owner, tknAddress)
}

// GetBalances is a free data retrieval call binding the contract method 0x6a385ae9.
//
// Solidity: function getBalances(address owner, address[] tknAddress) view returns(uint256[])
func (_Balance *BalanceCallerSession) GetBalances(owner common.Address, tknAddress []common.Address) ([]*big.Int, error) {
	return _Balance.Contract.GetBalances(&_Balance.CallOpts, owner, tknAddress)
}

// GetBalancesOfOwners is a free data retrieval call binding the contract method 0x2976b796.
//
// Solidity: function getBalancesOfOwners(address[] owners, address[] tokens) view returns((address,uint256[])[])
func (_Balance *BalanceCaller) GetBalancesOfOwners(opts *bind.CallOpts, owners []common.Address, tokens []common.Address) ([]ResolverBalanceReturnData, error) {
	var out []interface{}
	err := _Balance.contract.Call(opts, &out, "getBalancesOfOwners", owners, tokens)

	if err != nil {
		return *new([]ResolverBalanceReturnData), err
	}

	out0 := *abi.ConvertType(out[0], new([]ResolverBalanceReturnData)).(*[]ResolverBalanceReturnData)

	return out0, err

}

// GetBalancesOfOwners is a free data retrieval call binding the contract method 0x2976b796.
//
// Solidity: function getBalancesOfOwners(address[] owners, address[] tokens) view returns((address,uint256[])[])
func (_Balance *BalanceSession) GetBalancesOfOwners(owners []common.Address, tokens []common.Address) ([]ResolverBalanceReturnData, error) {
	return _Balance.Contract.GetBalancesOfOwners(&_Balance.CallOpts, owners, tokens)
}

// GetBalancesOfOwners is a free data retrieval call binding the contract method 0x2976b796.
//
// Solidity: function getBalancesOfOwners(address[] owners, address[] tokens) view returns((address,uint256[])[])
func (_Balance *BalanceCallerSession) GetBalancesOfOwners(owners []common.Address, tokens []common.Address) ([]ResolverBalanceReturnData, error) {
	return _Balance.Contract.GetBalancesOfOwners(&_Balance.CallOpts, owners, tokens)
}

// GetBalancesOfOwnersWithTokens is a free data retrieval call binding the contract method 0x83e9803f.
//
// Solidity: function getBalancesOfOwnersWithTokens((address,address[])[] datas) view returns((address,uint256[])[])
func (_Balance *BalanceCaller) GetBalancesOfOwnersWithTokens(opts *bind.CallOpts, datas []ResolverBalanceData) ([]ResolverBalanceReturnData, error) {
	var out []interface{}
	err := _Balance.contract.Call(opts, &out, "getBalancesOfOwnersWithTokens", datas)

	if err != nil {
		return *new([]ResolverBalanceReturnData), err
	}

	out0 := *abi.ConvertType(out[0], new([]ResolverBalanceReturnData)).(*[]ResolverBalanceReturnData)

	return out0, err

}

// GetBalancesOfOwnersWithTokens is a free data retrieval call binding the contract method 0x83e9803f.
//
// Solidity: function getBalancesOfOwnersWithTokens((address,address[])[] datas) view returns((address,uint256[])[])
func (_Balance *BalanceSession) GetBalancesOfOwnersWithTokens(datas []ResolverBalanceData) ([]ResolverBalanceReturnData, error) {
	return _Balance.Contract.GetBalancesOfOwnersWithTokens(&_Balance.CallOpts, datas)
}

// GetBalancesOfOwnersWithTokens is a free data retrieval call binding the contract method 0x83e9803f.
//
// Solidity: function getBalancesOfOwnersWithTokens((address,address[])[] datas) view returns((address,uint256[])[])
func (_Balance *BalanceCallerSession) GetBalancesOfOwnersWithTokens(datas []ResolverBalanceData) ([]ResolverBalanceReturnData, error) {
	return _Balance.Contract.GetBalancesOfOwnersWithTokens(&_Balance.CallOpts, datas)
}

// GetTokenDetails is a free data retrieval call binding the contract method 0xd2dd54b4.
//
// Solidity: function getTokenDetails(address[] tknAddress) view returns((bool,string,string,uint256)[])
func (_Balance *BalanceCaller) GetTokenDetails(opts *bind.CallOpts, tknAddress []common.Address) ([]ResolverTokenData, error) {
	var out []interface{}
	err := _Balance.contract.Call(opts, &out, "getTokenDetails", tknAddress)

	if err != nil {
		return *new([]ResolverTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new([]ResolverTokenData)).(*[]ResolverTokenData)

	return out0, err

}

// GetTokenDetails is a free data retrieval call binding the contract method 0xd2dd54b4.
//
// Solidity: function getTokenDetails(address[] tknAddress) view returns((bool,string,string,uint256)[])
func (_Balance *BalanceSession) GetTokenDetails(tknAddress []common.Address) ([]ResolverTokenData, error) {
	return _Balance.Contract.GetTokenDetails(&_Balance.CallOpts, tknAddress)
}

// GetTokenDetails is a free data retrieval call binding the contract method 0xd2dd54b4.
//
// Solidity: function getTokenDetails(address[] tknAddress) view returns((bool,string,string,uint256)[])
func (_Balance *BalanceCallerSession) GetTokenDetails(tknAddress []common.Address) ([]ResolverTokenData, error) {
	return _Balance.Contract.GetTokenDetails(&_Balance.CallOpts, tknAddress)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Balance *BalanceCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Balance.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Balance *BalanceSession) Name() (string, error) {
	return _Balance.Contract.Name(&_Balance.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Balance *BalanceCallerSession) Name() (string, error) {
	return _Balance.Contract.Name(&_Balance.CallOpts)
}
