package aprOracle

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

// AprOracleABI is the input ABI used to generate the binding from.
const AprOracleABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getLendDyDxAPR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AAVE\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBorrowCompoundAPR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBorrowFulcrumAPR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getLendAaveAPR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getLendFulcrumAPR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DYDX\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getLendCompoundAPR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getBorrowDyDxAPR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBorrowAaveAPR\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// AprOracle is an auto generated Go binding around an Ethereum contract.
type AprOracle struct {
	AprOracleCaller     // Read-only binding to the contract
	AprOracleTransactor // Write-only binding to the contract
	AprOracleFilterer   // Log filterer for contract events
}

// AprOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type AprOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AprOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AprOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AprOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AprOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AprOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AprOracleSession struct {
	Contract     *AprOracle        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AprOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AprOracleCallerSession struct {
	Contract *AprOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AprOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AprOracleTransactorSession struct {
	Contract     *AprOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AprOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type AprOracleRaw struct {
	Contract *AprOracle // Generic contract binding to access the raw methods on
}

// AprOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AprOracleCallerRaw struct {
	Contract *AprOracleCaller // Generic read-only contract binding to access the raw methods on
}

// AprOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AprOracleTransactorRaw struct {
	Contract *AprOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAprOracle creates a new instance of AprOracle, bound to a specific deployed contract.
func NewAprOracle(address common.Address, backend bind.ContractBackend) (*AprOracle, error) {
	contract, err := bindAprOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AprOracle{AprOracleCaller: AprOracleCaller{contract: contract}, AprOracleTransactor: AprOracleTransactor{contract: contract}, AprOracleFilterer: AprOracleFilterer{contract: contract}}, nil
}

// NewAprOracleCaller creates a new read-only instance of AprOracle, bound to a specific deployed contract.
func NewAprOracleCaller(address common.Address, caller bind.ContractCaller) (*AprOracleCaller, error) {
	contract, err := bindAprOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AprOracleCaller{contract: contract}, nil
}

// NewAprOracleTransactor creates a new write-only instance of AprOracle, bound to a specific deployed contract.
func NewAprOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*AprOracleTransactor, error) {
	contract, err := bindAprOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AprOracleTransactor{contract: contract}, nil
}

// NewAprOracleFilterer creates a new log filterer instance of AprOracle, bound to a specific deployed contract.
func NewAprOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*AprOracleFilterer, error) {
	contract, err := bindAprOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AprOracleFilterer{contract: contract}, nil
}

// bindAprOracle binds a generic wrapper to an already deployed contract.
func bindAprOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AprOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AprOracle *AprOracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AprOracle.Contract.AprOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AprOracle *AprOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AprOracle.Contract.AprOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AprOracle *AprOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AprOracle.Contract.AprOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AprOracle *AprOracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AprOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AprOracle *AprOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AprOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AprOracle *AprOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AprOracle.Contract.contract.Transact(opts, method, params...)
}

// AAVE is a free data retrieval call binding the contract method 0x48ccda3c.
//
// Solidity: function AAVE() view returns(address)
func (_AprOracle *AprOracleCaller) AAVE(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AprOracle.contract.Call(opts, out, "AAVE")
	return *ret0, err
}

// AAVE is a free data retrieval call binding the contract method 0x48ccda3c.
//
// Solidity: function AAVE() view returns(address)
func (_AprOracle *AprOracleSession) AAVE() (common.Address, error) {
	return _AprOracle.Contract.AAVE(&_AprOracle.CallOpts)
}

// AAVE is a free data retrieval call binding the contract method 0x48ccda3c.
//
// Solidity: function AAVE() view returns(address)
func (_AprOracle *AprOracleCallerSession) AAVE() (common.Address, error) {
	return _AprOracle.Contract.AAVE(&_AprOracle.CallOpts)
}

// DYDX is a free data retrieval call binding the contract method 0xc043fca2.
//
// Solidity: function DYDX() view returns(address)
func (_AprOracle *AprOracleCaller) DYDX(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AprOracle.contract.Call(opts, out, "DYDX")
	return *ret0, err
}

// DYDX is a free data retrieval call binding the contract method 0xc043fca2.
//
// Solidity: function DYDX() view returns(address)
func (_AprOracle *AprOracleSession) DYDX() (common.Address, error) {
	return _AprOracle.Contract.DYDX(&_AprOracle.CallOpts)
}

// DYDX is a free data retrieval call binding the contract method 0xc043fca2.
//
// Solidity: function DYDX() view returns(address)
func (_AprOracle *AprOracleCallerSession) DYDX() (common.Address, error) {
	return _AprOracle.Contract.DYDX(&_AprOracle.CallOpts)
}

// GetBorrowAaveAPR is a free data retrieval call binding the contract method 0xefb8eba9.
//
// Solidity: function getBorrowAaveAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleCaller) GetBorrowAaveAPR(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AprOracle.contract.Call(opts, out, "getBorrowAaveAPR", token)
	return *ret0, err
}

// GetBorrowAaveAPR is a free data retrieval call binding the contract method 0xefb8eba9.
//
// Solidity: function getBorrowAaveAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleSession) GetBorrowAaveAPR(token common.Address) (*big.Int, error) {
	return _AprOracle.Contract.GetBorrowAaveAPR(&_AprOracle.CallOpts, token)
}

// GetBorrowAaveAPR is a free data retrieval call binding the contract method 0xefb8eba9.
//
// Solidity: function getBorrowAaveAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleCallerSession) GetBorrowAaveAPR(token common.Address) (*big.Int, error) {
	return _AprOracle.Contract.GetBorrowAaveAPR(&_AprOracle.CallOpts, token)
}

// GetBorrowCompoundAPR is a free data retrieval call binding the contract method 0x5e14f079.
//
// Solidity: function getBorrowCompoundAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleCaller) GetBorrowCompoundAPR(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AprOracle.contract.Call(opts, out, "getBorrowCompoundAPR", token)
	return *ret0, err
}

// GetBorrowCompoundAPR is a free data retrieval call binding the contract method 0x5e14f079.
//
// Solidity: function getBorrowCompoundAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleSession) GetBorrowCompoundAPR(token common.Address) (*big.Int, error) {
	return _AprOracle.Contract.GetBorrowCompoundAPR(&_AprOracle.CallOpts, token)
}

// GetBorrowCompoundAPR is a free data retrieval call binding the contract method 0x5e14f079.
//
// Solidity: function getBorrowCompoundAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleCallerSession) GetBorrowCompoundAPR(token common.Address) (*big.Int, error) {
	return _AprOracle.Contract.GetBorrowCompoundAPR(&_AprOracle.CallOpts, token)
}

// GetBorrowDyDxAPR is a free data retrieval call binding the contract method 0xdbf30868.
//
// Solidity: function getBorrowDyDxAPR(uint256 marketId) view returns(uint256)
func (_AprOracle *AprOracleCaller) GetBorrowDyDxAPR(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AprOracle.contract.Call(opts, out, "getBorrowDyDxAPR", marketId)
	return *ret0, err
}

// GetBorrowDyDxAPR is a free data retrieval call binding the contract method 0xdbf30868.
//
// Solidity: function getBorrowDyDxAPR(uint256 marketId) view returns(uint256)
func (_AprOracle *AprOracleSession) GetBorrowDyDxAPR(marketId *big.Int) (*big.Int, error) {
	return _AprOracle.Contract.GetBorrowDyDxAPR(&_AprOracle.CallOpts, marketId)
}

// GetBorrowDyDxAPR is a free data retrieval call binding the contract method 0xdbf30868.
//
// Solidity: function getBorrowDyDxAPR(uint256 marketId) view returns(uint256)
func (_AprOracle *AprOracleCallerSession) GetBorrowDyDxAPR(marketId *big.Int) (*big.Int, error) {
	return _AprOracle.Contract.GetBorrowDyDxAPR(&_AprOracle.CallOpts, marketId)
}

// GetBorrowFulcrumAPR is a free data retrieval call binding the contract method 0x7952c6c6.
//
// Solidity: function getBorrowFulcrumAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleCaller) GetBorrowFulcrumAPR(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AprOracle.contract.Call(opts, out, "getBorrowFulcrumAPR", token)
	return *ret0, err
}

// GetBorrowFulcrumAPR is a free data retrieval call binding the contract method 0x7952c6c6.
//
// Solidity: function getBorrowFulcrumAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleSession) GetBorrowFulcrumAPR(token common.Address) (*big.Int, error) {
	return _AprOracle.Contract.GetBorrowFulcrumAPR(&_AprOracle.CallOpts, token)
}

// GetBorrowFulcrumAPR is a free data retrieval call binding the contract method 0x7952c6c6.
//
// Solidity: function getBorrowFulcrumAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleCallerSession) GetBorrowFulcrumAPR(token common.Address) (*big.Int, error) {
	return _AprOracle.Contract.GetBorrowFulcrumAPR(&_AprOracle.CallOpts, token)
}

// GetLendAaveAPR is a free data retrieval call binding the contract method 0x7a390efe.
//
// Solidity: function getLendAaveAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleCaller) GetLendAaveAPR(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AprOracle.contract.Call(opts, out, "getLendAaveAPR", token)
	return *ret0, err
}

// GetLendAaveAPR is a free data retrieval call binding the contract method 0x7a390efe.
//
// Solidity: function getLendAaveAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleSession) GetLendAaveAPR(token common.Address) (*big.Int, error) {
	return _AprOracle.Contract.GetLendAaveAPR(&_AprOracle.CallOpts, token)
}

// GetLendAaveAPR is a free data retrieval call binding the contract method 0x7a390efe.
//
// Solidity: function getLendAaveAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleCallerSession) GetLendAaveAPR(token common.Address) (*big.Int, error) {
	return _AprOracle.Contract.GetLendAaveAPR(&_AprOracle.CallOpts, token)
}

// GetLendCompoundAPR is a free data retrieval call binding the contract method 0xdae01e1c.
//
// Solidity: function getLendCompoundAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleCaller) GetLendCompoundAPR(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AprOracle.contract.Call(opts, out, "getLendCompoundAPR", token)
	return *ret0, err
}

// GetLendCompoundAPR is a free data retrieval call binding the contract method 0xdae01e1c.
//
// Solidity: function getLendCompoundAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleSession) GetLendCompoundAPR(token common.Address) (*big.Int, error) {
	return _AprOracle.Contract.GetLendCompoundAPR(&_AprOracle.CallOpts, token)
}

// GetLendCompoundAPR is a free data retrieval call binding the contract method 0xdae01e1c.
//
// Solidity: function getLendCompoundAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleCallerSession) GetLendCompoundAPR(token common.Address) (*big.Int, error) {
	return _AprOracle.Contract.GetLendCompoundAPR(&_AprOracle.CallOpts, token)
}

// GetLendDyDxAPR is a free data retrieval call binding the contract method 0x0a076f6d.
//
// Solidity: function getLendDyDxAPR(uint256 marketId) view returns(uint256)
func (_AprOracle *AprOracleCaller) GetLendDyDxAPR(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AprOracle.contract.Call(opts, out, "getLendDyDxAPR", marketId)
	return *ret0, err
}

// GetLendDyDxAPR is a free data retrieval call binding the contract method 0x0a076f6d.
//
// Solidity: function getLendDyDxAPR(uint256 marketId) view returns(uint256)
func (_AprOracle *AprOracleSession) GetLendDyDxAPR(marketId *big.Int) (*big.Int, error) {
	return _AprOracle.Contract.GetLendDyDxAPR(&_AprOracle.CallOpts, marketId)
}

// GetLendDyDxAPR is a free data retrieval call binding the contract method 0x0a076f6d.
//
// Solidity: function getLendDyDxAPR(uint256 marketId) view returns(uint256)
func (_AprOracle *AprOracleCallerSession) GetLendDyDxAPR(marketId *big.Int) (*big.Int, error) {
	return _AprOracle.Contract.GetLendDyDxAPR(&_AprOracle.CallOpts, marketId)
}

// GetLendFulcrumAPR is a free data retrieval call binding the contract method 0x7a542a26.
//
// Solidity: function getLendFulcrumAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleCaller) GetLendFulcrumAPR(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AprOracle.contract.Call(opts, out, "getLendFulcrumAPR", token)
	return *ret0, err
}

// GetLendFulcrumAPR is a free data retrieval call binding the contract method 0x7a542a26.
//
// Solidity: function getLendFulcrumAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleSession) GetLendFulcrumAPR(token common.Address) (*big.Int, error) {
	return _AprOracle.Contract.GetLendFulcrumAPR(&_AprOracle.CallOpts, token)
}

// GetLendFulcrumAPR is a free data retrieval call binding the contract method 0x7a542a26.
//
// Solidity: function getLendFulcrumAPR(address token) view returns(uint256)
func (_AprOracle *AprOracleCallerSession) GetLendFulcrumAPR(token common.Address) (*big.Int, error) {
	return _AprOracle.Contract.GetLendFulcrumAPR(&_AprOracle.CallOpts, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AprOracle *AprOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AprOracle.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AprOracle *AprOracleSession) Owner() (common.Address, error) {
	return _AprOracle.Contract.Owner(&_AprOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AprOracle *AprOracleCallerSession) Owner() (common.Address, error) {
	return _AprOracle.Contract.Owner(&_AprOracle.CallOpts)
}
