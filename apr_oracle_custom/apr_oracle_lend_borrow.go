package apr_oracle_custom

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ApreABI is the input ABI used to generate the binding from.
const ApreABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBorrowAPR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getLendDyDxAPR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AAVE\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBorrowFulcrumAPR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getLendAaveAPR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getLendFulcrumAPR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DYDX\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getLendCompoundAPR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getBorrowDyDxAPR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBorrowAaveAPR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Apre is an auto generated Go binding around an Ethereum contract.
type Apre struct {
	ApreCaller     // Read-only binding to the contract
	ApreTransactor // Write-only binding to the contract
	ApreFilterer   // Log filterer for contract events
}

// ApreCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApreSession struct {
	Contract     *Apre             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApreCallerSession struct {
	Contract *ApreCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApreTransactorSession struct {
	Contract     *ApreTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApreRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApreRaw struct {
	Contract *Apre // Generic contract binding to access the raw methods on
}

// ApreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApreCallerRaw struct {
	Contract *ApreCaller // Generic read-only contract binding to access the raw methods on
}

// ApreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApreTransactorRaw struct {
	Contract *ApreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAprCustom creates a new instance of Apre, bound to a specific deployed contract.
func NewAprCustom(address common.Address, backend bind.ContractBackend) (*Apre, error) {
	contract, err := bindApre(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Apre{ApreCaller: ApreCaller{contract: contract}, ApreTransactor: ApreTransactor{contract: contract}, ApreFilterer: ApreFilterer{contract: contract}}, nil
}

// NewApreCaller creates a new read-only instance of Apre, bound to a specific deployed contract.
func NewApreCaller(address common.Address, caller bind.ContractCaller) (*ApreCaller, error) {
	contract, err := bindApre(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApreCaller{contract: contract}, nil
}

// NewApreTransactor creates a new write-only instance of Apre, bound to a specific deployed contract.
func NewApreTransactor(address common.Address, transactor bind.ContractTransactor) (*ApreTransactor, error) {
	contract, err := bindApre(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApreTransactor{contract: contract}, nil
}

// NewApreFilterer creates a new log filterer instance of Apre, bound to a specific deployed contract.
func NewApreFilterer(address common.Address, filterer bind.ContractFilterer) (*ApreFilterer, error) {
	contract, err := bindApre(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApreFilterer{contract: contract}, nil
}

// bindApre binds a generic wrapper to an already deployed contract.
func bindApre(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Apre *ApreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Apre.Contract.ApreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Apre *ApreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apre.Contract.ApreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Apre *ApreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Apre.Contract.ApreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Apre *ApreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Apre.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Apre *ApreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apre.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Apre *ApreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Apre.Contract.contract.Transact(opts, method, params...)
}

// AAVE is a free data retrieval call binding the contract method 0x48ccda3c.
//
// Solidity: function AAVE() constant returns(address)
func (_Apre *ApreCaller) AAVE(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apre.contract.Call(opts, out, "AAVE")
	return *ret0, err
}

// AAVE is a free data retrieval call binding the contract method 0x48ccda3c.
//
// Solidity: function AAVE() constant returns(address)
func (_Apre *ApreSession) AAVE() (common.Address, error) {
	return _Apre.Contract.AAVE(&_Apre.CallOpts)
}

// AAVE is a free data retrieval call binding the contract method 0x48ccda3c.
//
// Solidity: function AAVE() constant returns(address)
func (_Apre *ApreCallerSession) AAVE() (common.Address, error) {
	return _Apre.Contract.AAVE(&_Apre.CallOpts)
}

// DYDX is a free data retrieval call binding the contract method 0xc043fca2.
//
// Solidity: function DYDX() constant returns(address)
func (_Apre *ApreCaller) DYDX(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apre.contract.Call(opts, out, "DYDX")
	return *ret0, err
}

// DYDX is a free data retrieval call binding the contract method 0xc043fca2.
//
// Solidity: function DYDX() constant returns(address)
func (_Apre *ApreSession) DYDX() (common.Address, error) {
	return _Apre.Contract.DYDX(&_Apre.CallOpts)
}

// DYDX is a free data retrieval call binding the contract method 0xc043fca2.
//
// Solidity: function DYDX() constant returns(address)
func (_Apre *ApreCallerSession) DYDX() (common.Address, error) {
	return _Apre.Contract.DYDX(&_Apre.CallOpts)
}

// GetBorrowAPR is a free data retrieval call binding the contract method 0x00e0f7a0.
//
// Solidity: function getBorrowAPR(address token) constant returns(uint256)
func (_Apre *ApreCaller) GetBorrowAPR(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apre.contract.Call(opts, out, "getBorrowAPR", token)
	return *ret0, err
}

// GetBorrowAPR is a free data retrieval call binding the contract method 0x00e0f7a0.
//
// Solidity: function getBorrowAPR(address token) constant returns(uint256)
func (_Apre *ApreSession) GetBorrowAPR(token common.Address) (*big.Int, error) {
	return _Apre.Contract.GetBorrowAPR(&_Apre.CallOpts, token)
}

// GetBorrowAPR is a free data retrieval call binding the contract method 0x00e0f7a0.
//
// Solidity: function getBorrowAPR(address token) constant returns(uint256)
func (_Apre *ApreCallerSession) GetBorrowAPR(token common.Address) (*big.Int, error) {
	return _Apre.Contract.GetBorrowAPR(&_Apre.CallOpts, token)
}

// GetBorrowAaveAPR is a free data retrieval call binding the contract method 0xefb8eba9.
//
// Solidity: function getBorrowAaveAPR(address token) constant returns(uint256)
func (_Apre *ApreCaller) GetBorrowAaveAPR(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apre.contract.Call(opts, out, "getBorrowAaveAPR", token)
	return *ret0, err
}

// GetBorrowAaveAPR is a free data retrieval call binding the contract method 0xefb8eba9.
//
// Solidity: function getBorrowAaveAPR(address token) constant returns(uint256)
func (_Apre *ApreSession) GetBorrowAaveAPR(token common.Address) (*big.Int, error) {
	return _Apre.Contract.GetBorrowAaveAPR(&_Apre.CallOpts, token)
}

// GetBorrowAaveAPR is a free data retrieval call binding the contract method 0xefb8eba9.
//
// Solidity: function getBorrowAaveAPR(address token) constant returns(uint256)
func (_Apre *ApreCallerSession) GetBorrowAaveAPR(token common.Address) (*big.Int, error) {
	return _Apre.Contract.GetBorrowAaveAPR(&_Apre.CallOpts, token)
}

// GetBorrowDyDxAPR is a free data retrieval call binding the contract method 0xdbf30868.
//
// Solidity: function getBorrowDyDxAPR(uint256 marketId) constant returns(uint256)
func (_Apre *ApreCaller) GetBorrowDyDxAPR(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apre.contract.Call(opts, out, "getBorrowDyDxAPR", marketId)
	return *ret0, err
}

// GetBorrowDyDxAPR is a free data retrieval call binding the contract method 0xdbf30868.
//
// Solidity: function getBorrowDyDxAPR(uint256 marketId) constant returns(uint256)
func (_Apre *ApreSession) GetBorrowDyDxAPR(marketId *big.Int) (*big.Int, error) {
	return _Apre.Contract.GetBorrowDyDxAPR(&_Apre.CallOpts, marketId)
}

// GetBorrowDyDxAPR is a free data retrieval call binding the contract method 0xdbf30868.
//
// Solidity: function getBorrowDyDxAPR(uint256 marketId) constant returns(uint256)
func (_Apre *ApreCallerSession) GetBorrowDyDxAPR(marketId *big.Int) (*big.Int, error) {
	return _Apre.Contract.GetBorrowDyDxAPR(&_Apre.CallOpts, marketId)
}

// GetBorrowFulcrumAPR is a free data retrieval call binding the contract method 0x7952c6c6.
//
// Solidity: function getBorrowFulcrumAPR(address token) constant returns(uint256)
func (_Apre *ApreCaller) GetBorrowFulcrumAPR(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apre.contract.Call(opts, out, "getBorrowFulcrumAPR", token)
	return *ret0, err
}

// GetBorrowFulcrumAPR is a free data retrieval call binding the contract method 0x7952c6c6.
//
// Solidity: function getBorrowFulcrumAPR(address token) constant returns(uint256)
func (_Apre *ApreSession) GetBorrowFulcrumAPR(token common.Address) (*big.Int, error) {
	return _Apre.Contract.GetBorrowFulcrumAPR(&_Apre.CallOpts, token)
}

// GetBorrowFulcrumAPR is a free data retrieval call binding the contract method 0x7952c6c6.
//
// Solidity: function getBorrowFulcrumAPR(address token) constant returns(uint256)
func (_Apre *ApreCallerSession) GetBorrowFulcrumAPR(token common.Address) (*big.Int, error) {
	return _Apre.Contract.GetBorrowFulcrumAPR(&_Apre.CallOpts, token)
}

// GetLendAaveAPR is a free data retrieval call binding the contract method 0x7a390efe.
//
// Solidity: function getLendAaveAPR(address token) constant returns(uint256)
func (_Apre *ApreCaller) GetLendAaveAPR(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apre.contract.Call(opts, out, "getLendAaveAPR", token)
	return *ret0, err
}

// GetLendAaveAPR is a free data retrieval call binding the contract method 0x7a390efe.
//
// Solidity: function getLendAaveAPR(address token) constant returns(uint256)
func (_Apre *ApreSession) GetLendAaveAPR(token common.Address) (*big.Int, error) {
	return _Apre.Contract.GetLendAaveAPR(&_Apre.CallOpts, token)
}

// GetLendAaveAPR is a free data retrieval call binding the contract method 0x7a390efe.
//
// Solidity: function getLendAaveAPR(address token) constant returns(uint256)
func (_Apre *ApreCallerSession) GetLendAaveAPR(token common.Address) (*big.Int, error) {
	return _Apre.Contract.GetLendAaveAPR(&_Apre.CallOpts, token)
}

// GetLendCompoundAPR is a free data retrieval call binding the contract method 0xdae01e1c.
//
// Solidity: function getLendCompoundAPR(address token) constant returns(uint256)
func (_Apre *ApreCaller) GetLendCompoundAPR(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apre.contract.Call(opts, out, "getLendCompoundAPR", token)
	return *ret0, err
}

// GetLendCompoundAPR is a free data retrieval call binding the contract method 0xdae01e1c.
//
// Solidity: function getLendCompoundAPR(address token) constant returns(uint256)
func (_Apre *ApreSession) GetLendCompoundAPR(token common.Address) (*big.Int, error) {
	return _Apre.Contract.GetLendCompoundAPR(&_Apre.CallOpts, token)
}

// GetLendCompoundAPR is a free data retrieval call binding the contract method 0xdae01e1c.
//
// Solidity: function getLendCompoundAPR(address token) constant returns(uint256)
func (_Apre *ApreCallerSession) GetLendCompoundAPR(token common.Address) (*big.Int, error) {
	return _Apre.Contract.GetLendCompoundAPR(&_Apre.CallOpts, token)
}

// GetLendDyDxAPR is a free data retrieval call binding the contract method 0x0a076f6d.
//
// Solidity: function getLendDyDxAPR(uint256 marketId) constant returns(uint256)
func (_Apre *ApreCaller) GetLendDyDxAPR(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apre.contract.Call(opts, out, "getLendDyDxAPR", marketId)
	return *ret0, err
}

// GetLendDyDxAPR is a free data retrieval call binding the contract method 0x0a076f6d.
//
// Solidity: function getLendDyDxAPR(uint256 marketId) constant returns(uint256)
func (_Apre *ApreSession) GetLendDyDxAPR(marketId *big.Int) (*big.Int, error) {
	return _Apre.Contract.GetLendDyDxAPR(&_Apre.CallOpts, marketId)
}

// GetLendDyDxAPR is a free data retrieval call binding the contract method 0x0a076f6d.
//
// Solidity: function getLendDyDxAPR(uint256 marketId) constant returns(uint256)
func (_Apre *ApreCallerSession) GetLendDyDxAPR(marketId *big.Int) (*big.Int, error) {
	return _Apre.Contract.GetLendDyDxAPR(&_Apre.CallOpts, marketId)
}

// GetLendFulcrumAPR is a free data retrieval call binding the contract method 0x7a542a26.
//
// Solidity: function getLendFulcrumAPR(address token) constant returns(uint256)
func (_Apre *ApreCaller) GetLendFulcrumAPR(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apre.contract.Call(opts, out, "getLendFulcrumAPR", token)
	return *ret0, err
}

// GetLendFulcrumAPR is a free data retrieval call binding the contract method 0x7a542a26.
//
// Solidity: function getLendFulcrumAPR(address token) constant returns(uint256)
func (_Apre *ApreSession) GetLendFulcrumAPR(token common.Address) (*big.Int, error) {
	return _Apre.Contract.GetLendFulcrumAPR(&_Apre.CallOpts, token)
}

// GetLendFulcrumAPR is a free data retrieval call binding the contract method 0x7a542a26.
//
// Solidity: function getLendFulcrumAPR(address token) constant returns(uint256)
func (_Apre *ApreCallerSession) GetLendFulcrumAPR(token common.Address) (*big.Int, error) {
	return _Apre.Contract.GetLendFulcrumAPR(&_Apre.CallOpts, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Apre *ApreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apre.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Apre *ApreSession) Owner() (common.Address, error) {
	return _Apre.Contract.Owner(&_Apre.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Apre *ApreCallerSession) Owner() (common.Address, error) {
	return _Apre.Contract.Owner(&_Apre.CallOpts)
}
