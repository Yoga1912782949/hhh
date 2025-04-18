// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
)

// AUstDIDsABI is the input ABI used to generate the binding from.
const AUstDIDsABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"credentialHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"CredentialIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"credentialHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"CredentialVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"DIDRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"addIssuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"addVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isIssuer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isVerifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"credentialHash\",\"type\":\"string\"}],\"name\":\"issueCredential\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bsSpecificString\",\"type\":\"string\"}],\"name\":\"registerDID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"removeIssuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"removeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"credentialHash\",\"type\":\"string\"}],\"name\":\"verifyCredential\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AUstDIDs is an auto generated Go binding around a Solidity contract.
type AUstDIDs struct {
	AUstDIDsCaller     // Read-only binding to the contract
	AUstDIDsTransactor // Write-only binding to the contract
	AUstDIDsFilterer   // Log filterer for contract events
}

// AUstDIDsCaller is an auto generated read-only Go binding around a Solidity contract.
type AUstDIDsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AUstDIDsTransactor is an auto generated write-only Go binding around a Solidity contract.
type AUstDIDsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AUstDIDsFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type AUstDIDsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AUstDIDsSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type AUstDIDsSession struct {
	Contract     *AUstDIDs         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AUstDIDsCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type AUstDIDsCallerSession struct {
	Contract *AUstDIDsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AUstDIDsTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type AUstDIDsTransactorSession struct {
	Contract     *AUstDIDsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AUstDIDsRaw is an auto generated low-level Go binding around a Solidity contract.
type AUstDIDsRaw struct {
	Contract *AUstDIDs // Generic contract binding to access the raw methods on
}

// AUstDIDsCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type AUstDIDsCallerRaw struct {
	Contract *AUstDIDsCaller // Generic read-only contract binding to access the raw methods on
}

// AUstDIDsTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type AUstDIDsTransactorRaw struct {
	Contract *AUstDIDsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAUstDIDs creates a new instance of AUstDIDs, bound to a specific deployed contract.
func NewAUstDIDs(address common.Address, backend bind.ContractBackend) (*AUstDIDs, error) {
	contract, err := bindAUstDIDs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AUstDIDs{AUstDIDsCaller: AUstDIDsCaller{contract: contract}, AUstDIDsTransactor: AUstDIDsTransactor{contract: contract}, AUstDIDsFilterer: AUstDIDsFilterer{contract: contract}}, nil
}

// NewAUstDIDsCaller creates a new read-only instance of AUstDIDs, bound to a specific deployed contract.
func NewAUstDIDsCaller(address common.Address, caller bind.ContractCaller) (*AUstDIDsCaller, error) {
	contract, err := bindAUstDIDs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AUstDIDsCaller{contract: contract}, nil
}

// NewAUstDIDsTransactor creates a new write-only instance of AUstDIDs, bound to a specific deployed contract.
func NewAUstDIDsTransactor(address common.Address, transactor bind.ContractTransactor) (*AUstDIDsTransactor, error) {
	contract, err := bindAUstDIDs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AUstDIDsTransactor{contract: contract}, nil
}

// NewAUstDIDsFilterer creates a new log filterer instance of AUstDIDs, bound to a specific deployed contract.
func NewAUstDIDsFilterer(address common.Address, filterer bind.ContractFilterer) (*AUstDIDsFilterer, error) {
	contract, err := bindAUstDIDs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AUstDIDsFilterer{contract: contract}, nil
}

// bindAUstDIDs binds a generic wrapper to an already deployed contract.
func bindAUstDIDs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AUstDIDsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AUstDIDs *AUstDIDsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AUstDIDs.Contract.AUstDIDsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AUstDIDs *AUstDIDsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.AUstDIDsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AUstDIDs *AUstDIDsRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.AUstDIDsTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AUstDIDs *AUstDIDsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AUstDIDs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AUstDIDs *AUstDIDsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AUstDIDs *AUstDIDsTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_AUstDIDs *AUstDIDsCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AUstDIDs.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_AUstDIDs *AUstDIDsSession) Admin() (common.Address, error) {
	return _AUstDIDs.Contract.Admin(&_AUstDIDs.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_AUstDIDs *AUstDIDsCallerSession) Admin() (common.Address, error) {
	return _AUstDIDs.Contract.Admin(&_AUstDIDs.CallOpts)
}

// IsIssuer is a free data retrieval call binding the contract method 0x877b9a67.
//
// Solidity: function isIssuer(address ) constant returns(bool)
func (_AUstDIDs *AUstDIDsCaller) IsIssuer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AUstDIDs.contract.Call(opts, out, "isIssuer", arg0)
	return *ret0, err
}

// IsIssuer is a free data retrieval call binding the contract method 0x877b9a67.
//
// Solidity: function isIssuer(address ) constant returns(bool)
func (_AUstDIDs *AUstDIDsSession) IsIssuer(arg0 common.Address) (bool, error) {
	return _AUstDIDs.Contract.IsIssuer(&_AUstDIDs.CallOpts, arg0)
}

// IsIssuer is a free data retrieval call binding the contract method 0x877b9a67.
//
// Solidity: function isIssuer(address ) constant returns(bool)
func (_AUstDIDs *AUstDIDsCallerSession) IsIssuer(arg0 common.Address) (bool, error) {
	return _AUstDIDs.Contract.IsIssuer(&_AUstDIDs.CallOpts, arg0)
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address ) constant returns(bool)
func (_AUstDIDs *AUstDIDsCaller) IsVerifier(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AUstDIDs.contract.Call(opts, out, "isVerifier", arg0)
	return *ret0, err
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address ) constant returns(bool)
func (_AUstDIDs *AUstDIDsSession) IsVerifier(arg0 common.Address) (bool, error) {
	return _AUstDIDs.Contract.IsVerifier(&_AUstDIDs.CallOpts, arg0)
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address ) constant returns(bool)
func (_AUstDIDs *AUstDIDsCallerSession) IsVerifier(arg0 common.Address) (bool, error) {
	return _AUstDIDs.Contract.IsVerifier(&_AUstDIDs.CallOpts, arg0)
}

// AddIssuer is a paid mutator transaction binding the contract method 0x20694db0.
//
// Solidity: function addIssuer(address issuer) returns()
func (_AUstDIDs *AUstDIDsTransactor) AddIssuer(opts *bind.TransactOpts, issuer common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _AUstDIDs.contract.TransactWithResult(opts, out, "addIssuer", issuer)
	return transaction, receipt, err
}

func (_AUstDIDs *AUstDIDsTransactor) AsyncAddIssuer(handler func(*types.Receipt, error), opts *bind.TransactOpts, issuer common.Address) (*types.Transaction, error) {
	return _AUstDIDs.contract.AsyncTransact(opts, handler, "addIssuer", issuer)
}

// AddIssuer is a paid mutator transaction binding the contract method 0x20694db0.
//
// Solidity: function addIssuer(address issuer) returns()
func (_AUstDIDs *AUstDIDsSession) AddIssuer(issuer common.Address) (*types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.AddIssuer(&_AUstDIDs.TransactOpts, issuer)
}

func (_AUstDIDs *AUstDIDsSession) AsyncAddIssuer(handler func(*types.Receipt, error), issuer common.Address) (*types.Transaction, error) {
	return _AUstDIDs.Contract.AsyncAddIssuer(handler, &_AUstDIDs.TransactOpts, issuer)
}

// AddIssuer is a paid mutator transaction binding the contract method 0x20694db0.
//
// Solidity: function addIssuer(address issuer) returns()
func (_AUstDIDs *AUstDIDsTransactorSession) AddIssuer(issuer common.Address) (*types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.AddIssuer(&_AUstDIDs.TransactOpts, issuer)
}

func (_AUstDIDs *AUstDIDsTransactorSession) AsyncAddIssuer(handler func(*types.Receipt, error), issuer common.Address) (*types.Transaction, error) {
	return _AUstDIDs.Contract.AsyncAddIssuer(handler, &_AUstDIDs.TransactOpts, issuer)
}

// AddVerifier is a paid mutator transaction binding the contract method 0x9000b3d6.
//
// Solidity: function addVerifier(address verifier) returns()
func (_AUstDIDs *AUstDIDsTransactor) AddVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _AUstDIDs.contract.TransactWithResult(opts, out, "addVerifier", verifier)
	return transaction, receipt, err
}

func (_AUstDIDs *AUstDIDsTransactor) AsyncAddVerifier(handler func(*types.Receipt, error), opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _AUstDIDs.contract.AsyncTransact(opts, handler, "addVerifier", verifier)
}

// AddVerifier is a paid mutator transaction binding the contract method 0x9000b3d6.
//
// Solidity: function addVerifier(address verifier) returns()
func (_AUstDIDs *AUstDIDsSession) AddVerifier(verifier common.Address) (*types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.AddVerifier(&_AUstDIDs.TransactOpts, verifier)
}

func (_AUstDIDs *AUstDIDsSession) AsyncAddVerifier(handler func(*types.Receipt, error), verifier common.Address) (*types.Transaction, error) {
	return _AUstDIDs.Contract.AsyncAddVerifier(handler, &_AUstDIDs.TransactOpts, verifier)
}

// AddVerifier is a paid mutator transaction binding the contract method 0x9000b3d6.
//
// Solidity: function addVerifier(address verifier) returns()
func (_AUstDIDs *AUstDIDsTransactorSession) AddVerifier(verifier common.Address) (*types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.AddVerifier(&_AUstDIDs.TransactOpts, verifier)
}

func (_AUstDIDs *AUstDIDsTransactorSession) AsyncAddVerifier(handler func(*types.Receipt, error), verifier common.Address) (*types.Transaction, error) {
	return _AUstDIDs.Contract.AsyncAddVerifier(handler, &_AUstDIDs.TransactOpts, verifier)
}

// IssueCredential is a paid mutator transaction binding the contract method 0x51ef0326.
//
// Solidity: function issueCredential(string did, string credentialHash) returns()
func (_AUstDIDs *AUstDIDsTransactor) IssueCredential(opts *bind.TransactOpts, did string, credentialHash string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _AUstDIDs.contract.TransactWithResult(opts, out, "issueCredential", did, credentialHash)
	return transaction, receipt, err
}

func (_AUstDIDs *AUstDIDsTransactor) AsyncIssueCredential(handler func(*types.Receipt, error), opts *bind.TransactOpts, did string, credentialHash string) (*types.Transaction, error) {
	return _AUstDIDs.contract.AsyncTransact(opts, handler, "issueCredential", did, credentialHash)
}

// IssueCredential is a paid mutator transaction binding the contract method 0x51ef0326.
//
// Solidity: function issueCredential(string did, string credentialHash) returns()
func (_AUstDIDs *AUstDIDsSession) IssueCredential(did string, credentialHash string) (*types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.IssueCredential(&_AUstDIDs.TransactOpts, did, credentialHash)
}

func (_AUstDIDs *AUstDIDsSession) AsyncIssueCredential(handler func(*types.Receipt, error), did string, credentialHash string) (*types.Transaction, error) {
	return _AUstDIDs.Contract.AsyncIssueCredential(handler, &_AUstDIDs.TransactOpts, did, credentialHash)
}

// IssueCredential is a paid mutator transaction binding the contract method 0x51ef0326.
//
// Solidity: function issueCredential(string did, string credentialHash) returns()
func (_AUstDIDs *AUstDIDsTransactorSession) IssueCredential(did string, credentialHash string) (*types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.IssueCredential(&_AUstDIDs.TransactOpts, did, credentialHash)
}

func (_AUstDIDs *AUstDIDsTransactorSession) AsyncIssueCredential(handler func(*types.Receipt, error), did string, credentialHash string) (*types.Transaction, error) {
	return _AUstDIDs.Contract.AsyncIssueCredential(handler, &_AUstDIDs.TransactOpts, did, credentialHash)
}

// RegisterDID is a paid mutator transaction binding the contract method 0x570810e2.
//
// Solidity: function registerDID(string bsSpecificString) returns()
func (_AUstDIDs *AUstDIDsTransactor) RegisterDID(opts *bind.TransactOpts, bsSpecificString string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _AUstDIDs.contract.TransactWithResult(opts, out, "registerDID", bsSpecificString)
	return transaction, receipt, err
}

func (_AUstDIDs *AUstDIDsTransactor) AsyncRegisterDID(handler func(*types.Receipt, error), opts *bind.TransactOpts, bsSpecificString string) (*types.Transaction, error) {
	return _AUstDIDs.contract.AsyncTransact(opts, handler, "registerDID", bsSpecificString)
}

// RegisterDID is a paid mutator transaction binding the contract method 0x570810e2.
//
// Solidity: function registerDID(string bsSpecificString) returns()
func (_AUstDIDs *AUstDIDsSession) RegisterDID(bsSpecificString string) (*types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.RegisterDID(&_AUstDIDs.TransactOpts, bsSpecificString)
}

func (_AUstDIDs *AUstDIDsSession) AsyncRegisterDID(handler func(*types.Receipt, error), bsSpecificString string) (*types.Transaction, error) {
	return _AUstDIDs.Contract.AsyncRegisterDID(handler, &_AUstDIDs.TransactOpts, bsSpecificString)
}

// RegisterDID is a paid mutator transaction binding the contract method 0x570810e2.
//
// Solidity: function registerDID(string bsSpecificString) returns()
func (_AUstDIDs *AUstDIDsTransactorSession) RegisterDID(bsSpecificString string) (*types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.RegisterDID(&_AUstDIDs.TransactOpts, bsSpecificString)
}

func (_AUstDIDs *AUstDIDsTransactorSession) AsyncRegisterDID(handler func(*types.Receipt, error), bsSpecificString string) (*types.Transaction, error) {
	return _AUstDIDs.Contract.AsyncRegisterDID(handler, &_AUstDIDs.TransactOpts, bsSpecificString)
}

// RemoveIssuer is a paid mutator transaction binding the contract method 0x47bc7093.
//
// Solidity: function removeIssuer(address issuer) returns()
func (_AUstDIDs *AUstDIDsTransactor) RemoveIssuer(opts *bind.TransactOpts, issuer common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _AUstDIDs.contract.TransactWithResult(opts, out, "removeIssuer", issuer)
	return transaction, receipt, err
}

func (_AUstDIDs *AUstDIDsTransactor) AsyncRemoveIssuer(handler func(*types.Receipt, error), opts *bind.TransactOpts, issuer common.Address) (*types.Transaction, error) {
	return _AUstDIDs.contract.AsyncTransact(opts, handler, "removeIssuer", issuer)
}

// RemoveIssuer is a paid mutator transaction binding the contract method 0x47bc7093.
//
// Solidity: function removeIssuer(address issuer) returns()
func (_AUstDIDs *AUstDIDsSession) RemoveIssuer(issuer common.Address) (*types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.RemoveIssuer(&_AUstDIDs.TransactOpts, issuer)
}

func (_AUstDIDs *AUstDIDsSession) AsyncRemoveIssuer(handler func(*types.Receipt, error), issuer common.Address) (*types.Transaction, error) {
	return _AUstDIDs.Contract.AsyncRemoveIssuer(handler, &_AUstDIDs.TransactOpts, issuer)
}

// RemoveIssuer is a paid mutator transaction binding the contract method 0x47bc7093.
//
// Solidity: function removeIssuer(address issuer) returns()
func (_AUstDIDs *AUstDIDsTransactorSession) RemoveIssuer(issuer common.Address) (*types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.RemoveIssuer(&_AUstDIDs.TransactOpts, issuer)
}

func (_AUstDIDs *AUstDIDsTransactorSession) AsyncRemoveIssuer(handler func(*types.Receipt, error), issuer common.Address) (*types.Transaction, error) {
	return _AUstDIDs.Contract.AsyncRemoveIssuer(handler, &_AUstDIDs.TransactOpts, issuer)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_AUstDIDs *AUstDIDsTransactor) RemoveVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _AUstDIDs.contract.TransactWithResult(opts, out, "removeVerifier", verifier)
	return transaction, receipt, err
}

func (_AUstDIDs *AUstDIDsTransactor) AsyncRemoveVerifier(handler func(*types.Receipt, error), opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _AUstDIDs.contract.AsyncTransact(opts, handler, "removeVerifier", verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_AUstDIDs *AUstDIDsSession) RemoveVerifier(verifier common.Address) (*types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.RemoveVerifier(&_AUstDIDs.TransactOpts, verifier)
}

func (_AUstDIDs *AUstDIDsSession) AsyncRemoveVerifier(handler func(*types.Receipt, error), verifier common.Address) (*types.Transaction, error) {
	return _AUstDIDs.Contract.AsyncRemoveVerifier(handler, &_AUstDIDs.TransactOpts, verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_AUstDIDs *AUstDIDsTransactorSession) RemoveVerifier(verifier common.Address) (*types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.RemoveVerifier(&_AUstDIDs.TransactOpts, verifier)
}

func (_AUstDIDs *AUstDIDsTransactorSession) AsyncRemoveVerifier(handler func(*types.Receipt, error), verifier common.Address) (*types.Transaction, error) {
	return _AUstDIDs.Contract.AsyncRemoveVerifier(handler, &_AUstDIDs.TransactOpts, verifier)
}

// VerifyCredential is a paid mutator transaction binding the contract method 0x1230abb5.
//
// Solidity: function verifyCredential(string did, string credentialHash) returns(bool)
func (_AUstDIDs *AUstDIDsTransactor) VerifyCredential(opts *bind.TransactOpts, did string, credentialHash string) (bool, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	transaction, receipt, err := _AUstDIDs.contract.TransactWithResult(opts, out, "verifyCredential", did, credentialHash)
	return *ret0, transaction, receipt, err
}

func (_AUstDIDs *AUstDIDsTransactor) AsyncVerifyCredential(handler func(*types.Receipt, error), opts *bind.TransactOpts, did string, credentialHash string) (*types.Transaction, error) {
	return _AUstDIDs.contract.AsyncTransact(opts, handler, "verifyCredential", did, credentialHash)
}

// VerifyCredential is a paid mutator transaction binding the contract method 0x1230abb5.
//
// Solidity: function verifyCredential(string did, string credentialHash) returns(bool)
func (_AUstDIDs *AUstDIDsSession) VerifyCredential(did string, credentialHash string) (bool, *types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.VerifyCredential(&_AUstDIDs.TransactOpts, did, credentialHash)
}

func (_AUstDIDs *AUstDIDsSession) AsyncVerifyCredential(handler func(*types.Receipt, error), did string, credentialHash string) (*types.Transaction, error) {
	return _AUstDIDs.Contract.AsyncVerifyCredential(handler, &_AUstDIDs.TransactOpts, did, credentialHash)
}

// VerifyCredential is a paid mutator transaction binding the contract method 0x1230abb5.
//
// Solidity: function verifyCredential(string did, string credentialHash) returns(bool)
func (_AUstDIDs *AUstDIDsTransactorSession) VerifyCredential(did string, credentialHash string) (bool, *types.Transaction, *types.Receipt, error) {
	return _AUstDIDs.Contract.VerifyCredential(&_AUstDIDs.TransactOpts, did, credentialHash)
}

func (_AUstDIDs *AUstDIDsTransactorSession) AsyncVerifyCredential(handler func(*types.Receipt, error), did string, credentialHash string) (*types.Transaction, error) {
	return _AUstDIDs.Contract.AsyncVerifyCredential(handler, &_AUstDIDs.TransactOpts, did, credentialHash)
}

// AUstDIDsCredentialIssued represents a CredentialIssued event raised by the AUstDIDs contract.
type AUstDIDsCredentialIssued struct {
	Did            common.Hash
	CredentialHash string
	Issuer         common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// WatchCredentialIssued is a free log subscription operation binding the contract event 0xbced6383e2dd56a440d8755db8d2520393bede4e079e12a544d29aa065174711.
//
// Solidity: event CredentialIssued(string indexed did, string credentialHash, address issuer)
func (_AUstDIDs *AUstDIDsFilterer) WatchCredentialIssued(fromBlock *uint64, handler func(int, []types.Log), did string) (string, error) {
	return _AUstDIDs.contract.WatchLogs(fromBlock, handler, "CredentialIssued", did)
}

func (_AUstDIDs *AUstDIDsFilterer) WatchAllCredentialIssued(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _AUstDIDs.contract.WatchLogs(fromBlock, handler, "CredentialIssued")
}

// ParseCredentialIssued is a log parse operation binding the contract event 0xbced6383e2dd56a440d8755db8d2520393bede4e079e12a544d29aa065174711.
//
// Solidity: event CredentialIssued(string indexed did, string credentialHash, address issuer)
func (_AUstDIDs *AUstDIDsFilterer) ParseCredentialIssued(log types.Log) (*AUstDIDsCredentialIssued, error) {
	event := new(AUstDIDsCredentialIssued)
	if err := _AUstDIDs.contract.UnpackLog(event, "CredentialIssued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchCredentialIssued is a free log subscription operation binding the contract event 0xbced6383e2dd56a440d8755db8d2520393bede4e079e12a544d29aa065174711.
//
// Solidity: event CredentialIssued(string indexed did, string credentialHash, address issuer)
func (_AUstDIDs *AUstDIDsSession) WatchCredentialIssued(fromBlock *uint64, handler func(int, []types.Log), did string) (string, error) {
	return _AUstDIDs.Contract.WatchCredentialIssued(fromBlock, handler, did)
}

func (_AUstDIDs *AUstDIDsSession) WatchAllCredentialIssued(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _AUstDIDs.Contract.WatchAllCredentialIssued(fromBlock, handler)
}

// ParseCredentialIssued is a log parse operation binding the contract event 0xbced6383e2dd56a440d8755db8d2520393bede4e079e12a544d29aa065174711.
//
// Solidity: event CredentialIssued(string indexed did, string credentialHash, address issuer)
func (_AUstDIDs *AUstDIDsSession) ParseCredentialIssued(log types.Log) (*AUstDIDsCredentialIssued, error) {
	return _AUstDIDs.Contract.ParseCredentialIssued(log)
}

// AUstDIDsCredentialVerified represents a CredentialVerified event raised by the AUstDIDs contract.
type AUstDIDsCredentialVerified struct {
	Did            common.Hash
	CredentialHash string
	Verifier       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// WatchCredentialVerified is a free log subscription operation binding the contract event 0xc850fcf4ae207de6eec1a515de48cc2b5071d284fdaf7b3445d16c5ce89a2606.
//
// Solidity: event CredentialVerified(string indexed did, string credentialHash, address verifier)
func (_AUstDIDs *AUstDIDsFilterer) WatchCredentialVerified(fromBlock *uint64, handler func(int, []types.Log), did string) (string, error) {
	return _AUstDIDs.contract.WatchLogs(fromBlock, handler, "CredentialVerified", did)
}

func (_AUstDIDs *AUstDIDsFilterer) WatchAllCredentialVerified(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _AUstDIDs.contract.WatchLogs(fromBlock, handler, "CredentialVerified")
}

// ParseCredentialVerified is a log parse operation binding the contract event 0xc850fcf4ae207de6eec1a515de48cc2b5071d284fdaf7b3445d16c5ce89a2606.
//
// Solidity: event CredentialVerified(string indexed did, string credentialHash, address verifier)
func (_AUstDIDs *AUstDIDsFilterer) ParseCredentialVerified(log types.Log) (*AUstDIDsCredentialVerified, error) {
	event := new(AUstDIDsCredentialVerified)
	if err := _AUstDIDs.contract.UnpackLog(event, "CredentialVerified", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchCredentialVerified is a free log subscription operation binding the contract event 0xc850fcf4ae207de6eec1a515de48cc2b5071d284fdaf7b3445d16c5ce89a2606.
//
// Solidity: event CredentialVerified(string indexed did, string credentialHash, address verifier)
func (_AUstDIDs *AUstDIDsSession) WatchCredentialVerified(fromBlock *uint64, handler func(int, []types.Log), did string) (string, error) {
	return _AUstDIDs.Contract.WatchCredentialVerified(fromBlock, handler, did)
}

func (_AUstDIDs *AUstDIDsSession) WatchAllCredentialVerified(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _AUstDIDs.Contract.WatchAllCredentialVerified(fromBlock, handler)
}

// ParseCredentialVerified is a log parse operation binding the contract event 0xc850fcf4ae207de6eec1a515de48cc2b5071d284fdaf7b3445d16c5ce89a2606.
//
// Solidity: event CredentialVerified(string indexed did, string credentialHash, address verifier)
func (_AUstDIDs *AUstDIDsSession) ParseCredentialVerified(log types.Log) (*AUstDIDsCredentialVerified, error) {
	return _AUstDIDs.Contract.ParseCredentialVerified(log)
}

// AUstDIDsDIDRegistered represents a DIDRegistered event raised by the AUstDIDs contract.
type AUstDIDsDIDRegistered struct {
	Owner common.Address
	Did   string
	Raw   types.Log // Blockchain specific contextual infos
}

// WatchDIDRegistered is a free log subscription operation binding the contract event 0xbfab59a7af78dfc896c3096604749fcee2db2e850297de42de5e6369f54bd3fb.
//
// Solidity: event DIDRegistered(address indexed owner, string did)
func (_AUstDIDs *AUstDIDsFilterer) WatchDIDRegistered(fromBlock *uint64, handler func(int, []types.Log), owner common.Address) (string, error) {
	return _AUstDIDs.contract.WatchLogs(fromBlock, handler, "DIDRegistered", owner)
}

func (_AUstDIDs *AUstDIDsFilterer) WatchAllDIDRegistered(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _AUstDIDs.contract.WatchLogs(fromBlock, handler, "DIDRegistered")
}

// ParseDIDRegistered is a log parse operation binding the contract event 0xbfab59a7af78dfc896c3096604749fcee2db2e850297de42de5e6369f54bd3fb.
//
// Solidity: event DIDRegistered(address indexed owner, string did)
func (_AUstDIDs *AUstDIDsFilterer) ParseDIDRegistered(log types.Log) (*AUstDIDsDIDRegistered, error) {
	event := new(AUstDIDsDIDRegistered)
	if err := _AUstDIDs.contract.UnpackLog(event, "DIDRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchDIDRegistered is a free log subscription operation binding the contract event 0xbfab59a7af78dfc896c3096604749fcee2db2e850297de42de5e6369f54bd3fb.
//
// Solidity: event DIDRegistered(address indexed owner, string did)
func (_AUstDIDs *AUstDIDsSession) WatchDIDRegistered(fromBlock *uint64, handler func(int, []types.Log), owner common.Address) (string, error) {
	return _AUstDIDs.Contract.WatchDIDRegistered(fromBlock, handler, owner)
}

func (_AUstDIDs *AUstDIDsSession) WatchAllDIDRegistered(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _AUstDIDs.Contract.WatchAllDIDRegistered(fromBlock, handler)
}

// ParseDIDRegistered is a log parse operation binding the contract event 0xbfab59a7af78dfc896c3096604749fcee2db2e850297de42de5e6369f54bd3fb.
//
// Solidity: event DIDRegistered(address indexed owner, string did)
func (_AUstDIDs *AUstDIDsSession) ParseDIDRegistered(log types.Log) (*AUstDIDsDIDRegistered, error) {
	return _AUstDIDs.Contract.ParseDIDRegistered(log)
}
