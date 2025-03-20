// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// Grant is an auto generated low-level Go binding around an user-defined struct.
type Grant struct {
	ClaimTime     uint64
	ClaimInterval uint64
	ClaimLimit    *big.Int
	DustBalance   *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"grantTimeUsed\",\"type\":\"uint64\"}],\"name\":\"GrantClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lockTime\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lockTimeout\",\"type\":\"uint64\"}],\"name\":\"GrantLock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"interval\",\"type\":\"uint64\"}],\"name\":\"GrantUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GRANT_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"requestAmount\",\"type\":\"uint256\"}],\"name\":\"_calculateClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"claimAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"newClaimTime\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"newDustBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"usedTime\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"claimTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"claimTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"interval\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"createGrant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"}],\"name\":\"getClaimableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"}],\"name\":\"getGrant\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"claimTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"claimInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"claimLimit\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"dustBalance\",\"type\":\"uint256\"}],\"internalType\":\"structGrant\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"}],\"name\":\"getGrantLockTime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"}],\"name\":\"getGrantName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"}],\"name\":\"getGrantTotalClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGrants\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"claimTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"claimInterval\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"claimLimit\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"dustBalance\",\"type\":\"uint256\"}],\"internalType\":\"structGrant[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"getManagerCooldown\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getManagerGrantLimits\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lockTime\",\"type\":\"uint64\"}],\"name\":\"lockGrant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"}],\"name\":\"notifyGrantTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"}],\"name\":\"removeGrant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"renameGrant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"rescueCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"lockTime\",\"type\":\"uint32\"}],\"name\":\"setClaimTransferLockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"interval\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"cooldown\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"cooldownLock\",\"type\":\"uint32\"}],\"name\":\"setManagerGrantLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setProxyManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"transferGrant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"grantId\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"interval\",\"type\":\"uint64\"}],\"name\":\"updateGrant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Contract *ContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Contract *ContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Contract.Contract.DEFAULTADMINROLE(&_Contract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Contract *ContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Contract.Contract.DEFAULTADMINROLE(&_Contract.CallOpts)
}

// GRANTMANAGERROLE is a free data retrieval call binding the contract method 0x90270417.
//
// Solidity: function GRANT_MANAGER_ROLE() view returns(bytes32)
func (_Contract *ContractCaller) GRANTMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "GRANT_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GRANTMANAGERROLE is a free data retrieval call binding the contract method 0x90270417.
//
// Solidity: function GRANT_MANAGER_ROLE() view returns(bytes32)
func (_Contract *ContractSession) GRANTMANAGERROLE() ([32]byte, error) {
	return _Contract.Contract.GRANTMANAGERROLE(&_Contract.CallOpts)
}

// GRANTMANAGERROLE is a free data retrieval call binding the contract method 0x90270417.
//
// Solidity: function GRANT_MANAGER_ROLE() view returns(bytes32)
func (_Contract *ContractCallerSession) GRANTMANAGERROLE() ([32]byte, error) {
	return _Contract.Contract.GRANTMANAGERROLE(&_Contract.CallOpts)
}

// CalculateClaim is a free data retrieval call binding the contract method 0xfd264ad6.
//
// Solidity: function _calculateClaim(uint64 grantId, uint256 requestAmount) view returns(uint256 claimAmount, uint64 newClaimTime, uint256 newDustBalance, uint64 usedTime)
func (_Contract *ContractCaller) CalculateClaim(opts *bind.CallOpts, grantId uint64, requestAmount *big.Int) (struct {
	ClaimAmount    *big.Int
	NewClaimTime   uint64
	NewDustBalance *big.Int
	UsedTime       uint64
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_calculateClaim", grantId, requestAmount)

	outstruct := new(struct {
		ClaimAmount    *big.Int
		NewClaimTime   uint64
		NewDustBalance *big.Int
		UsedTime       uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ClaimAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NewClaimTime = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.NewDustBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UsedTime = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// CalculateClaim is a free data retrieval call binding the contract method 0xfd264ad6.
//
// Solidity: function _calculateClaim(uint64 grantId, uint256 requestAmount) view returns(uint256 claimAmount, uint64 newClaimTime, uint256 newDustBalance, uint64 usedTime)
func (_Contract *ContractSession) CalculateClaim(grantId uint64, requestAmount *big.Int) (struct {
	ClaimAmount    *big.Int
	NewClaimTime   uint64
	NewDustBalance *big.Int
	UsedTime       uint64
}, error) {
	return _Contract.Contract.CalculateClaim(&_Contract.CallOpts, grantId, requestAmount)
}

// CalculateClaim is a free data retrieval call binding the contract method 0xfd264ad6.
//
// Solidity: function _calculateClaim(uint64 grantId, uint256 requestAmount) view returns(uint256 claimAmount, uint64 newClaimTime, uint256 newDustBalance, uint64 usedTime)
func (_Contract *ContractCallerSession) CalculateClaim(grantId uint64, requestAmount *big.Int) (struct {
	ClaimAmount    *big.Int
	NewClaimTime   uint64
	NewDustBalance *big.Int
	UsedTime       uint64
}, error) {
	return _Contract.Contract.CalculateClaim(&_Contract.CallOpts, grantId, requestAmount)
}

// GetClaimableBalance is a free data retrieval call binding the contract method 0x939a3faf.
//
// Solidity: function getClaimableBalance(uint64 grantId) view returns(uint256)
func (_Contract *ContractCaller) GetClaimableBalance(opts *bind.CallOpts, grantId uint64) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getClaimableBalance", grantId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableBalance is a free data retrieval call binding the contract method 0x939a3faf.
//
// Solidity: function getClaimableBalance(uint64 grantId) view returns(uint256)
func (_Contract *ContractSession) GetClaimableBalance(grantId uint64) (*big.Int, error) {
	return _Contract.Contract.GetClaimableBalance(&_Contract.CallOpts, grantId)
}

// GetClaimableBalance is a free data retrieval call binding the contract method 0x939a3faf.
//
// Solidity: function getClaimableBalance(uint64 grantId) view returns(uint256)
func (_Contract *ContractCallerSession) GetClaimableBalance(grantId uint64) (*big.Int, error) {
	return _Contract.Contract.GetClaimableBalance(&_Contract.CallOpts, grantId)
}

// GetClaimableBalance0 is a free data retrieval call binding the contract method 0xe0bd3015.
//
// Solidity: function getClaimableBalance() view returns(uint256)
func (_Contract *ContractCaller) GetClaimableBalance0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getClaimableBalance0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableBalance0 is a free data retrieval call binding the contract method 0xe0bd3015.
//
// Solidity: function getClaimableBalance() view returns(uint256)
func (_Contract *ContractSession) GetClaimableBalance0() (*big.Int, error) {
	return _Contract.Contract.GetClaimableBalance0(&_Contract.CallOpts)
}

// GetClaimableBalance0 is a free data retrieval call binding the contract method 0xe0bd3015.
//
// Solidity: function getClaimableBalance() view returns(uint256)
func (_Contract *ContractCallerSession) GetClaimableBalance0() (*big.Int, error) {
	return _Contract.Contract.GetClaimableBalance0(&_Contract.CallOpts)
}

// GetGrant is a free data retrieval call binding the contract method 0x2f274411.
//
// Solidity: function getGrant(uint64 grantId) view returns((uint64,uint64,uint128,uint256))
func (_Contract *ContractCaller) GetGrant(opts *bind.CallOpts, grantId uint64) (Grant, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getGrant", grantId)

	if err != nil {
		return *new(Grant), err
	}

	out0 := *abi.ConvertType(out[0], new(Grant)).(*Grant)

	return out0, err

}

// GetGrant is a free data retrieval call binding the contract method 0x2f274411.
//
// Solidity: function getGrant(uint64 grantId) view returns((uint64,uint64,uint128,uint256))
func (_Contract *ContractSession) GetGrant(grantId uint64) (Grant, error) {
	return _Contract.Contract.GetGrant(&_Contract.CallOpts, grantId)
}

// GetGrant is a free data retrieval call binding the contract method 0x2f274411.
//
// Solidity: function getGrant(uint64 grantId) view returns((uint64,uint64,uint128,uint256))
func (_Contract *ContractCallerSession) GetGrant(grantId uint64) (Grant, error) {
	return _Contract.Contract.GetGrant(&_Contract.CallOpts, grantId)
}

// GetGrantLockTime is a free data retrieval call binding the contract method 0x346a5bf3.
//
// Solidity: function getGrantLockTime(uint64 grantId) view returns(uint64)
func (_Contract *ContractCaller) GetGrantLockTime(opts *bind.CallOpts, grantId uint64) (uint64, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getGrantLockTime", grantId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetGrantLockTime is a free data retrieval call binding the contract method 0x346a5bf3.
//
// Solidity: function getGrantLockTime(uint64 grantId) view returns(uint64)
func (_Contract *ContractSession) GetGrantLockTime(grantId uint64) (uint64, error) {
	return _Contract.Contract.GetGrantLockTime(&_Contract.CallOpts, grantId)
}

// GetGrantLockTime is a free data retrieval call binding the contract method 0x346a5bf3.
//
// Solidity: function getGrantLockTime(uint64 grantId) view returns(uint64)
func (_Contract *ContractCallerSession) GetGrantLockTime(grantId uint64) (uint64, error) {
	return _Contract.Contract.GetGrantLockTime(&_Contract.CallOpts, grantId)
}

// GetGrantName is a free data retrieval call binding the contract method 0xf364bc3c.
//
// Solidity: function getGrantName(uint64 grantId) view returns(bytes32)
func (_Contract *ContractCaller) GetGrantName(opts *bind.CallOpts, grantId uint64) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getGrantName", grantId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetGrantName is a free data retrieval call binding the contract method 0xf364bc3c.
//
// Solidity: function getGrantName(uint64 grantId) view returns(bytes32)
func (_Contract *ContractSession) GetGrantName(grantId uint64) ([32]byte, error) {
	return _Contract.Contract.GetGrantName(&_Contract.CallOpts, grantId)
}

// GetGrantName is a free data retrieval call binding the contract method 0xf364bc3c.
//
// Solidity: function getGrantName(uint64 grantId) view returns(bytes32)
func (_Contract *ContractCallerSession) GetGrantName(grantId uint64) ([32]byte, error) {
	return _Contract.Contract.GetGrantName(&_Contract.CallOpts, grantId)
}

// GetGrantTotalClaimed is a free data retrieval call binding the contract method 0xc46a5c9f.
//
// Solidity: function getGrantTotalClaimed(uint64 grantId) view returns(uint256)
func (_Contract *ContractCaller) GetGrantTotalClaimed(opts *bind.CallOpts, grantId uint64) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getGrantTotalClaimed", grantId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGrantTotalClaimed is a free data retrieval call binding the contract method 0xc46a5c9f.
//
// Solidity: function getGrantTotalClaimed(uint64 grantId) view returns(uint256)
func (_Contract *ContractSession) GetGrantTotalClaimed(grantId uint64) (*big.Int, error) {
	return _Contract.Contract.GetGrantTotalClaimed(&_Contract.CallOpts, grantId)
}

// GetGrantTotalClaimed is a free data retrieval call binding the contract method 0xc46a5c9f.
//
// Solidity: function getGrantTotalClaimed(uint64 grantId) view returns(uint256)
func (_Contract *ContractCallerSession) GetGrantTotalClaimed(grantId uint64) (*big.Int, error) {
	return _Contract.Contract.GetGrantTotalClaimed(&_Contract.CallOpts, grantId)
}

// GetGrants is a free data retrieval call binding the contract method 0x158c75e5.
//
// Solidity: function getGrants() view returns((uint64,uint64,uint128,uint256)[])
func (_Contract *ContractCaller) GetGrants(opts *bind.CallOpts) ([]Grant, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getGrants")

	if err != nil {
		return *new([]Grant), err
	}

	out0 := *abi.ConvertType(out[0], new([]Grant)).(*[]Grant)

	return out0, err

}

// GetGrants is a free data retrieval call binding the contract method 0x158c75e5.
//
// Solidity: function getGrants() view returns((uint64,uint64,uint128,uint256)[])
func (_Contract *ContractSession) GetGrants() ([]Grant, error) {
	return _Contract.Contract.GetGrants(&_Contract.CallOpts)
}

// GetGrants is a free data retrieval call binding the contract method 0x158c75e5.
//
// Solidity: function getGrants() view returns((uint64,uint64,uint128,uint256)[])
func (_Contract *ContractCallerSession) GetGrants() ([]Grant, error) {
	return _Contract.Contract.GetGrants(&_Contract.CallOpts)
}

// GetManagerCooldown is a free data retrieval call binding the contract method 0x5f2f2697.
//
// Solidity: function getManagerCooldown(address manager) view returns(uint64)
func (_Contract *ContractCaller) GetManagerCooldown(opts *bind.CallOpts, manager common.Address) (uint64, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getManagerCooldown", manager)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetManagerCooldown is a free data retrieval call binding the contract method 0x5f2f2697.
//
// Solidity: function getManagerCooldown(address manager) view returns(uint64)
func (_Contract *ContractSession) GetManagerCooldown(manager common.Address) (uint64, error) {
	return _Contract.Contract.GetManagerCooldown(&_Contract.CallOpts, manager)
}

// GetManagerCooldown is a free data retrieval call binding the contract method 0x5f2f2697.
//
// Solidity: function getManagerCooldown(address manager) view returns(uint64)
func (_Contract *ContractCallerSession) GetManagerCooldown(manager common.Address) (uint64, error) {
	return _Contract.Contract.GetManagerCooldown(&_Contract.CallOpts, manager)
}

// GetManagerGrantLimits is a free data retrieval call binding the contract method 0x56e239f1.
//
// Solidity: function getManagerGrantLimits() view returns(uint128, uint64, uint32, uint32)
func (_Contract *ContractCaller) GetManagerGrantLimits(opts *bind.CallOpts) (*big.Int, uint64, uint32, uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getManagerGrantLimits")

	if err != nil {
		return *new(*big.Int), *new(uint64), *new(uint32), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(uint32)).(*uint32)
	out3 := *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return out0, out1, out2, out3, err

}

// GetManagerGrantLimits is a free data retrieval call binding the contract method 0x56e239f1.
//
// Solidity: function getManagerGrantLimits() view returns(uint128, uint64, uint32, uint32)
func (_Contract *ContractSession) GetManagerGrantLimits() (*big.Int, uint64, uint32, uint32, error) {
	return _Contract.Contract.GetManagerGrantLimits(&_Contract.CallOpts)
}

// GetManagerGrantLimits is a free data retrieval call binding the contract method 0x56e239f1.
//
// Solidity: function getManagerGrantLimits() view returns(uint128, uint64, uint32, uint32)
func (_Contract *ContractCallerSession) GetManagerGrantLimits() (*big.Int, uint64, uint32, uint32, error) {
	return _Contract.Contract.GetManagerGrantLimits(&_Contract.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Contract *ContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Contract *ContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Contract.Contract.GetRoleAdmin(&_Contract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Contract *ContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Contract.Contract.GetRoleAdmin(&_Contract.CallOpts, role)
}

// GetVaultToken is a free data retrieval call binding the contract method 0x76d0c864.
//
// Solidity: function getVaultToken() view returns(address)
func (_Contract *ContractCaller) GetVaultToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getVaultToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVaultToken is a free data retrieval call binding the contract method 0x76d0c864.
//
// Solidity: function getVaultToken() view returns(address)
func (_Contract *ContractSession) GetVaultToken() (common.Address, error) {
	return _Contract.Contract.GetVaultToken(&_Contract.CallOpts)
}

// GetVaultToken is a free data retrieval call binding the contract method 0x76d0c864.
//
// Solidity: function getVaultToken() view returns(address)
func (_Contract *ContractCallerSession) GetVaultToken() (common.Address, error) {
	return _Contract.Contract.GetVaultToken(&_Contract.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Contract *ContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Contract *ContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Contract.Contract.HasRole(&_Contract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Contract *ContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Contract.Contract.HasRole(&_Contract.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contract *ContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contract *ContractSession) Paused() (bool, error) {
	return _Contract.Contract.Paused(&_Contract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Contract *ContractCallerSession) Paused() (bool, error) {
	return _Contract.Contract.Paused(&_Contract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 amount) returns(uint256)
func (_Contract *ContractTransactor) Claim(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claim", amount)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 amount) returns(uint256)
func (_Contract *ContractSession) Claim(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Claim(&_Contract.TransactOpts, amount)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 amount) returns(uint256)
func (_Contract *ContractTransactorSession) Claim(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Claim(&_Contract.TransactOpts, amount)
}

// Claim0 is a paid mutator transaction binding the contract method 0x503914db.
//
// Solidity: function claim(uint64 grantId, uint256 amount) returns(uint256)
func (_Contract *ContractTransactor) Claim0(opts *bind.TransactOpts, grantId uint64, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claim0", grantId, amount)
}

// Claim0 is a paid mutator transaction binding the contract method 0x503914db.
//
// Solidity: function claim(uint64 grantId, uint256 amount) returns(uint256)
func (_Contract *ContractSession) Claim0(grantId uint64, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Claim0(&_Contract.TransactOpts, grantId, amount)
}

// Claim0 is a paid mutator transaction binding the contract method 0x503914db.
//
// Solidity: function claim(uint64 grantId, uint256 amount) returns(uint256)
func (_Contract *ContractTransactorSession) Claim0(grantId uint64, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Claim0(&_Contract.TransactOpts, grantId, amount)
}

// ClaimTo is a paid mutator transaction binding the contract method 0x1fca9342.
//
// Solidity: function claimTo(uint64 grantId, uint256 amount, address target) returns(uint256)
func (_Contract *ContractTransactor) ClaimTo(opts *bind.TransactOpts, grantId uint64, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimTo", grantId, amount, target)
}

// ClaimTo is a paid mutator transaction binding the contract method 0x1fca9342.
//
// Solidity: function claimTo(uint64 grantId, uint256 amount, address target) returns(uint256)
func (_Contract *ContractSession) ClaimTo(grantId uint64, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ClaimTo(&_Contract.TransactOpts, grantId, amount, target)
}

// ClaimTo is a paid mutator transaction binding the contract method 0x1fca9342.
//
// Solidity: function claimTo(uint64 grantId, uint256 amount, address target) returns(uint256)
func (_Contract *ContractTransactorSession) ClaimTo(grantId uint64, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ClaimTo(&_Contract.TransactOpts, grantId, amount, target)
}

// ClaimTo0 is a paid mutator transaction binding the contract method 0x30e1198b.
//
// Solidity: function claimTo(uint256 amount, address target) returns(uint256)
func (_Contract *ContractTransactor) ClaimTo0(opts *bind.TransactOpts, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimTo0", amount, target)
}

// ClaimTo0 is a paid mutator transaction binding the contract method 0x30e1198b.
//
// Solidity: function claimTo(uint256 amount, address target) returns(uint256)
func (_Contract *ContractSession) ClaimTo0(amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ClaimTo0(&_Contract.TransactOpts, amount, target)
}

// ClaimTo0 is a paid mutator transaction binding the contract method 0x30e1198b.
//
// Solidity: function claimTo(uint256 amount, address target) returns(uint256)
func (_Contract *ContractTransactorSession) ClaimTo0(amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ClaimTo0(&_Contract.TransactOpts, amount, target)
}

// CreateGrant is a paid mutator transaction binding the contract method 0x5be61f5e.
//
// Solidity: function createGrant(address addr, uint128 amount, uint64 interval, bytes32 name) returns()
func (_Contract *ContractTransactor) CreateGrant(opts *bind.TransactOpts, addr common.Address, amount *big.Int, interval uint64, name [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createGrant", addr, amount, interval, name)
}

// CreateGrant is a paid mutator transaction binding the contract method 0x5be61f5e.
//
// Solidity: function createGrant(address addr, uint128 amount, uint64 interval, bytes32 name) returns()
func (_Contract *ContractSession) CreateGrant(addr common.Address, amount *big.Int, interval uint64, name [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.CreateGrant(&_Contract.TransactOpts, addr, amount, interval, name)
}

// CreateGrant is a paid mutator transaction binding the contract method 0x5be61f5e.
//
// Solidity: function createGrant(address addr, uint128 amount, uint64 interval, bytes32 name) returns()
func (_Contract *ContractTransactorSession) CreateGrant(addr common.Address, amount *big.Int, interval uint64, name [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.CreateGrant(&_Contract.TransactOpts, addr, amount, interval, name)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Contract *ContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Contract *ContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.GrantRole(&_Contract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Contract *ContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.GrantRole(&_Contract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address tokenAddr) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, tokenAddr common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", tokenAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address tokenAddr) returns()
func (_Contract *ContractSession) Initialize(tokenAddr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, tokenAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address tokenAddr) returns()
func (_Contract *ContractTransactorSession) Initialize(tokenAddr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, tokenAddr)
}

// LockGrant is a paid mutator transaction binding the contract method 0x23d2dad7.
//
// Solidity: function lockGrant(uint64 grantId, uint64 lockTime) returns()
func (_Contract *ContractTransactor) LockGrant(opts *bind.TransactOpts, grantId uint64, lockTime uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "lockGrant", grantId, lockTime)
}

// LockGrant is a paid mutator transaction binding the contract method 0x23d2dad7.
//
// Solidity: function lockGrant(uint64 grantId, uint64 lockTime) returns()
func (_Contract *ContractSession) LockGrant(grantId uint64, lockTime uint64) (*types.Transaction, error) {
	return _Contract.Contract.LockGrant(&_Contract.TransactOpts, grantId, lockTime)
}

// LockGrant is a paid mutator transaction binding the contract method 0x23d2dad7.
//
// Solidity: function lockGrant(uint64 grantId, uint64 lockTime) returns()
func (_Contract *ContractTransactorSession) LockGrant(grantId uint64, lockTime uint64) (*types.Transaction, error) {
	return _Contract.Contract.LockGrant(&_Contract.TransactOpts, grantId, lockTime)
}

// NotifyGrantTransfer is a paid mutator transaction binding the contract method 0x540b755c.
//
// Solidity: function notifyGrantTransfer(uint64 grantId) returns()
func (_Contract *ContractTransactor) NotifyGrantTransfer(opts *bind.TransactOpts, grantId uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "notifyGrantTransfer", grantId)
}

// NotifyGrantTransfer is a paid mutator transaction binding the contract method 0x540b755c.
//
// Solidity: function notifyGrantTransfer(uint64 grantId) returns()
func (_Contract *ContractSession) NotifyGrantTransfer(grantId uint64) (*types.Transaction, error) {
	return _Contract.Contract.NotifyGrantTransfer(&_Contract.TransactOpts, grantId)
}

// NotifyGrantTransfer is a paid mutator transaction binding the contract method 0x540b755c.
//
// Solidity: function notifyGrantTransfer(uint64 grantId) returns()
func (_Contract *ContractTransactorSession) NotifyGrantTransfer(grantId uint64) (*types.Transaction, error) {
	return _Contract.Contract.NotifyGrantTransfer(&_Contract.TransactOpts, grantId)
}

// RemoveGrant is a paid mutator transaction binding the contract method 0x362e7e13.
//
// Solidity: function removeGrant(uint64 grantId) returns()
func (_Contract *ContractTransactor) RemoveGrant(opts *bind.TransactOpts, grantId uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "removeGrant", grantId)
}

// RemoveGrant is a paid mutator transaction binding the contract method 0x362e7e13.
//
// Solidity: function removeGrant(uint64 grantId) returns()
func (_Contract *ContractSession) RemoveGrant(grantId uint64) (*types.Transaction, error) {
	return _Contract.Contract.RemoveGrant(&_Contract.TransactOpts, grantId)
}

// RemoveGrant is a paid mutator transaction binding the contract method 0x362e7e13.
//
// Solidity: function removeGrant(uint64 grantId) returns()
func (_Contract *ContractTransactorSession) RemoveGrant(grantId uint64) (*types.Transaction, error) {
	return _Contract.Contract.RemoveGrant(&_Contract.TransactOpts, grantId)
}

// RenameGrant is a paid mutator transaction binding the contract method 0xa0000259.
//
// Solidity: function renameGrant(uint64 grantId, bytes32 name) returns()
func (_Contract *ContractTransactor) RenameGrant(opts *bind.TransactOpts, grantId uint64, name [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renameGrant", grantId, name)
}

// RenameGrant is a paid mutator transaction binding the contract method 0xa0000259.
//
// Solidity: function renameGrant(uint64 grantId, bytes32 name) returns()
func (_Contract *ContractSession) RenameGrant(grantId uint64, name [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.RenameGrant(&_Contract.TransactOpts, grantId, name)
}

// RenameGrant is a paid mutator transaction binding the contract method 0xa0000259.
//
// Solidity: function renameGrant(uint64 grantId, bytes32 name) returns()
func (_Contract *ContractTransactorSession) RenameGrant(grantId uint64, name [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.RenameGrant(&_Contract.TransactOpts, grantId, name)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Contract *ContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Contract *ContractSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RenounceRole(&_Contract.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Contract *ContractTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RenounceRole(&_Contract.TransactOpts, role, account)
}

// RescueCall is a paid mutator transaction binding the contract method 0x96dfe5de.
//
// Solidity: function rescueCall(address addr, uint256 amount, bytes data) returns()
func (_Contract *ContractTransactor) RescueCall(opts *bind.TransactOpts, addr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "rescueCall", addr, amount, data)
}

// RescueCall is a paid mutator transaction binding the contract method 0x96dfe5de.
//
// Solidity: function rescueCall(address addr, uint256 amount, bytes data) returns()
func (_Contract *ContractSession) RescueCall(addr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.RescueCall(&_Contract.TransactOpts, addr, amount, data)
}

// RescueCall is a paid mutator transaction binding the contract method 0x96dfe5de.
//
// Solidity: function rescueCall(address addr, uint256 amount, bytes data) returns()
func (_Contract *ContractTransactorSession) RescueCall(addr common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.RescueCall(&_Contract.TransactOpts, addr, amount, data)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Contract *ContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Contract *ContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RevokeRole(&_Contract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Contract *ContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RevokeRole(&_Contract.TransactOpts, role, account)
}

// SetClaimTransferLockTime is a paid mutator transaction binding the contract method 0xa6a1cb4c.
//
// Solidity: function setClaimTransferLockTime(uint32 lockTime) returns()
func (_Contract *ContractTransactor) SetClaimTransferLockTime(opts *bind.TransactOpts, lockTime uint32) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setClaimTransferLockTime", lockTime)
}

// SetClaimTransferLockTime is a paid mutator transaction binding the contract method 0xa6a1cb4c.
//
// Solidity: function setClaimTransferLockTime(uint32 lockTime) returns()
func (_Contract *ContractSession) SetClaimTransferLockTime(lockTime uint32) (*types.Transaction, error) {
	return _Contract.Contract.SetClaimTransferLockTime(&_Contract.TransactOpts, lockTime)
}

// SetClaimTransferLockTime is a paid mutator transaction binding the contract method 0xa6a1cb4c.
//
// Solidity: function setClaimTransferLockTime(uint32 lockTime) returns()
func (_Contract *ContractTransactorSession) SetClaimTransferLockTime(lockTime uint32) (*types.Transaction, error) {
	return _Contract.Contract.SetClaimTransferLockTime(&_Contract.TransactOpts, lockTime)
}

// SetManagerGrantLimits is a paid mutator transaction binding the contract method 0x08cf0ebb.
//
// Solidity: function setManagerGrantLimits(uint128 amount, uint64 interval, uint32 cooldown, uint32 cooldownLock) returns()
func (_Contract *ContractTransactor) SetManagerGrantLimits(opts *bind.TransactOpts, amount *big.Int, interval uint64, cooldown uint32, cooldownLock uint32) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setManagerGrantLimits", amount, interval, cooldown, cooldownLock)
}

// SetManagerGrantLimits is a paid mutator transaction binding the contract method 0x08cf0ebb.
//
// Solidity: function setManagerGrantLimits(uint128 amount, uint64 interval, uint32 cooldown, uint32 cooldownLock) returns()
func (_Contract *ContractSession) SetManagerGrantLimits(amount *big.Int, interval uint64, cooldown uint32, cooldownLock uint32) (*types.Transaction, error) {
	return _Contract.Contract.SetManagerGrantLimits(&_Contract.TransactOpts, amount, interval, cooldown, cooldownLock)
}

// SetManagerGrantLimits is a paid mutator transaction binding the contract method 0x08cf0ebb.
//
// Solidity: function setManagerGrantLimits(uint128 amount, uint64 interval, uint32 cooldown, uint32 cooldownLock) returns()
func (_Contract *ContractTransactorSession) SetManagerGrantLimits(amount *big.Int, interval uint64, cooldown uint32, cooldownLock uint32) (*types.Transaction, error) {
	return _Contract.Contract.SetManagerGrantLimits(&_Contract.TransactOpts, amount, interval, cooldown, cooldownLock)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_Contract *ContractTransactor) SetPaused(opts *bind.TransactOpts, paused bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPaused", paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_Contract *ContractSession) SetPaused(paused bool) (*types.Transaction, error) {
	return _Contract.Contract.SetPaused(&_Contract.TransactOpts, paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool paused) returns()
func (_Contract *ContractTransactorSession) SetPaused(paused bool) (*types.Transaction, error) {
	return _Contract.Contract.SetPaused(&_Contract.TransactOpts, paused)
}

// SetProxyManager is a paid mutator transaction binding the contract method 0xfe7f3505.
//
// Solidity: function setProxyManager(address manager) returns()
func (_Contract *ContractTransactor) SetProxyManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setProxyManager", manager)
}

// SetProxyManager is a paid mutator transaction binding the contract method 0xfe7f3505.
//
// Solidity: function setProxyManager(address manager) returns()
func (_Contract *ContractSession) SetProxyManager(manager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetProxyManager(&_Contract.TransactOpts, manager)
}

// SetProxyManager is a paid mutator transaction binding the contract method 0xfe7f3505.
//
// Solidity: function setProxyManager(address manager) returns()
func (_Contract *ContractTransactorSession) SetProxyManager(manager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetProxyManager(&_Contract.TransactOpts, manager)
}

// TransferGrant is a paid mutator transaction binding the contract method 0xdc0533ef.
//
// Solidity: function transferGrant(uint64 grantId, address addr) returns()
func (_Contract *ContractTransactor) TransferGrant(opts *bind.TransactOpts, grantId uint64, addr common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferGrant", grantId, addr)
}

// TransferGrant is a paid mutator transaction binding the contract method 0xdc0533ef.
//
// Solidity: function transferGrant(uint64 grantId, address addr) returns()
func (_Contract *ContractSession) TransferGrant(grantId uint64, addr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferGrant(&_Contract.TransactOpts, grantId, addr)
}

// TransferGrant is a paid mutator transaction binding the contract method 0xdc0533ef.
//
// Solidity: function transferGrant(uint64 grantId, address addr) returns()
func (_Contract *ContractTransactorSession) TransferGrant(grantId uint64, addr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferGrant(&_Contract.TransactOpts, grantId, addr)
}

// UpdateGrant is a paid mutator transaction binding the contract method 0x8a1a14eb.
//
// Solidity: function updateGrant(uint64 grantId, uint128 amount, uint64 interval) returns()
func (_Contract *ContractTransactor) UpdateGrant(opts *bind.TransactOpts, grantId uint64, amount *big.Int, interval uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateGrant", grantId, amount, interval)
}

// UpdateGrant is a paid mutator transaction binding the contract method 0x8a1a14eb.
//
// Solidity: function updateGrant(uint64 grantId, uint128 amount, uint64 interval) returns()
func (_Contract *ContractSession) UpdateGrant(grantId uint64, amount *big.Int, interval uint64) (*types.Transaction, error) {
	return _Contract.Contract.UpdateGrant(&_Contract.TransactOpts, grantId, amount, interval)
}

// UpdateGrant is a paid mutator transaction binding the contract method 0x8a1a14eb.
//
// Solidity: function updateGrant(uint64 grantId, uint128 amount, uint64 interval) returns()
func (_Contract *ContractTransactorSession) UpdateGrant(grantId uint64, amount *big.Int, interval uint64) (*types.Transaction, error) {
	return _Contract.Contract.UpdateGrant(&_Contract.TransactOpts, grantId, amount, interval)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactorSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// ContractGrantClaimIterator is returned from FilterGrantClaim and is used to iterate over the raw logs and unpacked data for GrantClaim events raised by the Contract contract.
type ContractGrantClaimIterator struct {
	Event *ContractGrantClaim // Event containing the contract specifics and raw log

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
func (it *ContractGrantClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractGrantClaim)
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
		it.Event = new(ContractGrantClaim)
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
func (it *ContractGrantClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractGrantClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractGrantClaim represents a GrantClaim event raised by the Contract contract.
type ContractGrantClaim struct {
	GrantId       uint64
	To            common.Address
	Amount        *big.Int
	GrantTimeUsed uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGrantClaim is a free log retrieval operation binding the contract event 0x9bc6681351310ea1f31aebc395b080cba707096cdb76f0a969c1c2e1a4c9339a.
//
// Solidity: event GrantClaim(uint64 indexed grantId, address indexed to, uint256 amount, uint64 grantTimeUsed)
func (_Contract *ContractFilterer) FilterGrantClaim(opts *bind.FilterOpts, grantId []uint64, to []common.Address) (*ContractGrantClaimIterator, error) {

	var grantIdRule []interface{}
	for _, grantIdItem := range grantId {
		grantIdRule = append(grantIdRule, grantIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "GrantClaim", grantIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractGrantClaimIterator{contract: _Contract.contract, event: "GrantClaim", logs: logs, sub: sub}, nil
}

// WatchGrantClaim is a free log subscription operation binding the contract event 0x9bc6681351310ea1f31aebc395b080cba707096cdb76f0a969c1c2e1a4c9339a.
//
// Solidity: event GrantClaim(uint64 indexed grantId, address indexed to, uint256 amount, uint64 grantTimeUsed)
func (_Contract *ContractFilterer) WatchGrantClaim(opts *bind.WatchOpts, sink chan<- *ContractGrantClaim, grantId []uint64, to []common.Address) (event.Subscription, error) {

	var grantIdRule []interface{}
	for _, grantIdItem := range grantId {
		grantIdRule = append(grantIdRule, grantIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "GrantClaim", grantIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractGrantClaim)
				if err := _Contract.contract.UnpackLog(event, "GrantClaim", log); err != nil {
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

// ParseGrantClaim is a log parse operation binding the contract event 0x9bc6681351310ea1f31aebc395b080cba707096cdb76f0a969c1c2e1a4c9339a.
//
// Solidity: event GrantClaim(uint64 indexed grantId, address indexed to, uint256 amount, uint64 grantTimeUsed)
func (_Contract *ContractFilterer) ParseGrantClaim(log types.Log) (*ContractGrantClaim, error) {
	event := new(ContractGrantClaim)
	if err := _Contract.contract.UnpackLog(event, "GrantClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractGrantLockIterator is returned from FilterGrantLock and is used to iterate over the raw logs and unpacked data for GrantLock events raised by the Contract contract.
type ContractGrantLockIterator struct {
	Event *ContractGrantLock // Event containing the contract specifics and raw log

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
func (it *ContractGrantLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractGrantLock)
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
		it.Event = new(ContractGrantLock)
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
func (it *ContractGrantLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractGrantLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractGrantLock represents a GrantLock event raised by the Contract contract.
type ContractGrantLock struct {
	GrantId     uint64
	LockTime    uint64
	LockTimeout uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGrantLock is a free log retrieval operation binding the contract event 0x86eb882c8f52dab9f2ab23b63cd1c3d92e01df6aac65bd2e7c22a58c7cf6efb7.
//
// Solidity: event GrantLock(uint64 indexed grantId, uint64 lockTime, uint64 lockTimeout)
func (_Contract *ContractFilterer) FilterGrantLock(opts *bind.FilterOpts, grantId []uint64) (*ContractGrantLockIterator, error) {

	var grantIdRule []interface{}
	for _, grantIdItem := range grantId {
		grantIdRule = append(grantIdRule, grantIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "GrantLock", grantIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractGrantLockIterator{contract: _Contract.contract, event: "GrantLock", logs: logs, sub: sub}, nil
}

// WatchGrantLock is a free log subscription operation binding the contract event 0x86eb882c8f52dab9f2ab23b63cd1c3d92e01df6aac65bd2e7c22a58c7cf6efb7.
//
// Solidity: event GrantLock(uint64 indexed grantId, uint64 lockTime, uint64 lockTimeout)
func (_Contract *ContractFilterer) WatchGrantLock(opts *bind.WatchOpts, sink chan<- *ContractGrantLock, grantId []uint64) (event.Subscription, error) {

	var grantIdRule []interface{}
	for _, grantIdItem := range grantId {
		grantIdRule = append(grantIdRule, grantIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "GrantLock", grantIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractGrantLock)
				if err := _Contract.contract.UnpackLog(event, "GrantLock", log); err != nil {
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

// ParseGrantLock is a log parse operation binding the contract event 0x86eb882c8f52dab9f2ab23b63cd1c3d92e01df6aac65bd2e7c22a58c7cf6efb7.
//
// Solidity: event GrantLock(uint64 indexed grantId, uint64 lockTime, uint64 lockTimeout)
func (_Contract *ContractFilterer) ParseGrantLock(log types.Log) (*ContractGrantLock, error) {
	event := new(ContractGrantLock)
	if err := _Contract.contract.UnpackLog(event, "GrantLock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractGrantUpdateIterator is returned from FilterGrantUpdate and is used to iterate over the raw logs and unpacked data for GrantUpdate events raised by the Contract contract.
type ContractGrantUpdateIterator struct {
	Event *ContractGrantUpdate // Event containing the contract specifics and raw log

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
func (it *ContractGrantUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractGrantUpdate)
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
		it.Event = new(ContractGrantUpdate)
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
func (it *ContractGrantUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractGrantUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractGrantUpdate represents a GrantUpdate event raised by the Contract contract.
type ContractGrantUpdate struct {
	GrantId  uint64
	Amount   *big.Int
	Interval uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGrantUpdate is a free log retrieval operation binding the contract event 0x3376546f4371ebddddc9b82197e907e12636954e240f592c9fe99fb6dfcf9e7a.
//
// Solidity: event GrantUpdate(uint64 indexed grantId, uint128 amount, uint64 interval)
func (_Contract *ContractFilterer) FilterGrantUpdate(opts *bind.FilterOpts, grantId []uint64) (*ContractGrantUpdateIterator, error) {

	var grantIdRule []interface{}
	for _, grantIdItem := range grantId {
		grantIdRule = append(grantIdRule, grantIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "GrantUpdate", grantIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractGrantUpdateIterator{contract: _Contract.contract, event: "GrantUpdate", logs: logs, sub: sub}, nil
}

// WatchGrantUpdate is a free log subscription operation binding the contract event 0x3376546f4371ebddddc9b82197e907e12636954e240f592c9fe99fb6dfcf9e7a.
//
// Solidity: event GrantUpdate(uint64 indexed grantId, uint128 amount, uint64 interval)
func (_Contract *ContractFilterer) WatchGrantUpdate(opts *bind.WatchOpts, sink chan<- *ContractGrantUpdate, grantId []uint64) (event.Subscription, error) {

	var grantIdRule []interface{}
	for _, grantIdItem := range grantId {
		grantIdRule = append(grantIdRule, grantIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "GrantUpdate", grantIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractGrantUpdate)
				if err := _Contract.contract.UnpackLog(event, "GrantUpdate", log); err != nil {
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

// ParseGrantUpdate is a log parse operation binding the contract event 0x3376546f4371ebddddc9b82197e907e12636954e240f592c9fe99fb6dfcf9e7a.
//
// Solidity: event GrantUpdate(uint64 indexed grantId, uint128 amount, uint64 interval)
func (_Contract *ContractFilterer) ParseGrantUpdate(log types.Log) (*ContractGrantUpdate, error) {
	event := new(ContractGrantUpdate)
	if err := _Contract.contract.UnpackLog(event, "GrantUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Contract contract.
type ContractPausedIterator struct {
	Event *ContractPaused // Event containing the contract specifics and raw log

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
func (it *ContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPaused)
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
		it.Event = new(ContractPaused)
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
func (it *ContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPaused represents a Paused event raised by the Contract contract.
type ContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Contract *ContractFilterer) FilterPaused(opts *bind.FilterOpts) (*ContractPausedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ContractPausedIterator{contract: _Contract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Contract *ContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractPaused) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPaused)
				if err := _Contract.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Contract *ContractFilterer) ParsePaused(log types.Log) (*ContractPaused, error) {
	event := new(ContractPaused)
	if err := _Contract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Contract contract.
type ContractRoleAdminChangedIterator struct {
	Event *ContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRoleAdminChanged)
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
		it.Event = new(ContractRoleAdminChanged)
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
func (it *ContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRoleAdminChanged represents a RoleAdminChanged event raised by the Contract contract.
type ContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Contract *ContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ContractRoleAdminChangedIterator{contract: _Contract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Contract *ContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRoleAdminChanged)
				if err := _Contract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Contract *ContractFilterer) ParseRoleAdminChanged(log types.Log) (*ContractRoleAdminChanged, error) {
	event := new(ContractRoleAdminChanged)
	if err := _Contract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Contract contract.
type ContractRoleGrantedIterator struct {
	Event *ContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *ContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRoleGranted)
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
		it.Event = new(ContractRoleGranted)
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
func (it *ContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRoleGranted represents a RoleGranted event raised by the Contract contract.
type ContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contract *ContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractRoleGrantedIterator{contract: _Contract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contract *ContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRoleGranted)
				if err := _Contract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contract *ContractFilterer) ParseRoleGranted(log types.Log) (*ContractRoleGranted, error) {
	event := new(ContractRoleGranted)
	if err := _Contract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Contract contract.
type ContractRoleRevokedIterator struct {
	Event *ContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRoleRevoked)
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
		it.Event = new(ContractRoleRevoked)
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
func (it *ContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRoleRevoked represents a RoleRevoked event raised by the Contract contract.
type ContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contract *ContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractRoleRevokedIterator{contract: _Contract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contract *ContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRoleRevoked)
				if err := _Contract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contract *ContractFilterer) ParseRoleRevoked(log types.Log) (*ContractRoleRevoked, error) {
	event := new(ContractRoleRevoked)
	if err := _Contract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Contract contract.
type ContractUnpausedIterator struct {
	Event *ContractUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUnpaused)
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
		it.Event = new(ContractUnpaused)
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
func (it *ContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUnpaused represents a Unpaused event raised by the Contract contract.
type ContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Contract *ContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ContractUnpausedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ContractUnpausedIterator{contract: _Contract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Contract *ContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUnpaused)
				if err := _Contract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Contract *ContractFilterer) ParseUnpaused(log types.Log) (*ContractUnpaused, error) {
	event := new(ContractUnpaused)
	if err := _Contract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
