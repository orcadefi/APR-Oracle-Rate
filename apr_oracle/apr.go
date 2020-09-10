// This file is a generated binding and any manual changes will be lost.
// Smart Contract address 0x97FF4A1b787ADe6b94cca95b61F79417c673331D (APROracle)
package apr_oracle

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

// AprABI is the input ABI used to generate the binding from.
const AprABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"AAVE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ABAT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ADAI\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AKNC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ALEND\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ALINK\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AMANA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AMKR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AREP\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ASNX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ASUSD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ATUSD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AUSDC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AUSDT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AWBTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"AZRX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CBAT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CDAI\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CREP\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CSAI\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CUSDC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CWBTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CZRX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DYDX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IBAT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IDAI\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IKNC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ILINK\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IREP\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ISAI\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ISUSD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IUSDC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IWBTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IZRX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getABATAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getADAIAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAETHAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAKNCAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getALENDAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getALINKAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAMANAAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAMKRAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAREPAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getASNXAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getASUSDAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getATUSDAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAUSDCAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAUSDTAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAWBTCAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAZRXAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getAaveAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAaveCore\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllAaveAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"aDAI\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aTUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aUSDC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aUSDT\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aSUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aBAT\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aLINK\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aKNC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aREP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aZRX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aSNX\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllCompoundAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cDAI\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cBAT\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cREP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cSAI\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cUSDC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cWBTC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cZRC\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllDyDxAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dSAI\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dUSDC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dDAI\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllFulcrumAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"iZRX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iREP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iKNC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iWBTC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iUSDC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iSAI\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iDAI\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iLINK\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"iSUSD\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCBATAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCDAIAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCETHAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCREPAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCSAIAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCUSDCAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCWBTCAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCZRCAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getCompoundAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"marketId\",\"type\":\"uint256\"}],\"name\":\"getDyDxAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDyDxDAIAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDyDxETHAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDyDxSAIAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDyDxUSDCAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFulcrumAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIDAIAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIETHAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIKNCAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getILINKAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIREPAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getISAIAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getISUSDAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIUSDCAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIWBTCAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIZRXAPR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"inCaseETHGetsStuck\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_TokenAddress\",\"type\":\"address\"}],\"name\":\"inCaseTokenGetsStuck\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"recommendDAI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"recommendETH\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"recommendUSDC\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_liquidity\",\"type\":\"uint256\"}],\"name\":\"setLiquidity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_AAVE\",\"type\":\"address\"}],\"name\":\"set_new_AAVE\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_ABAT\",\"type\":\"address\"}],\"name\":\"set_new_ABAT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_ADAI\",\"type\":\"address\"}],\"name\":\"set_new_ADAI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_AETH\",\"type\":\"address\"}],\"name\":\"set_new_AETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_AKNC\",\"type\":\"address\"}],\"name\":\"set_new_AKNC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_ALEND\",\"type\":\"address\"}],\"name\":\"set_new_ALEND\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_ALINK\",\"type\":\"address\"}],\"name\":\"set_new_ALINK\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_AMANA\",\"type\":\"address\"}],\"name\":\"set_new_AMANA\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_AMKR\",\"type\":\"address\"}],\"name\":\"set_new_AMKR\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_AREP\",\"type\":\"address\"}],\"name\":\"set_new_AREP\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_ASNX\",\"type\":\"address\"}],\"name\":\"set_new_ASNX\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_ASUSD\",\"type\":\"address\"}],\"name\":\"set_new_ASUSD\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_ATUSD\",\"type\":\"address\"}],\"name\":\"set_new_ATUSD\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_AUSDC\",\"type\":\"address\"}],\"name\":\"set_new_AUSDC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_AUSDT\",\"type\":\"address\"}],\"name\":\"set_new_AUSDT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_AWBTC\",\"type\":\"address\"}],\"name\":\"set_new_AWBTC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_AZRX\",\"type\":\"address\"}],\"name\":\"set_new_AZRX\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_CBAT\",\"type\":\"address\"}],\"name\":\"set_new_CBAT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_CDAI\",\"type\":\"address\"}],\"name\":\"set_new_CDAI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_CETH\",\"type\":\"address\"}],\"name\":\"set_new_CETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_CREP\",\"type\":\"address\"}],\"name\":\"set_new_CREP\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_CSAI\",\"type\":\"address\"}],\"name\":\"set_new_CSAI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_CUSDC\",\"type\":\"address\"}],\"name\":\"set_new_CUSDC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_CWBTC\",\"type\":\"address\"}],\"name\":\"set_new_CWBTC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_CZRX\",\"type\":\"address\"}],\"name\":\"set_new_CZRX\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_DYDX\",\"type\":\"address\"}],\"name\":\"set_new_DYDX\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_IDAI\",\"type\":\"address\"}],\"name\":\"set_new_IDAI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_IETH\",\"type\":\"address\"}],\"name\":\"set_new_IETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_IKNC\",\"type\":\"address\"}],\"name\":\"set_new_IKNC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_ILINK\",\"type\":\"address\"}],\"name\":\"set_new_ILINK\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_IREP\",\"type\":\"address\"}],\"name\":\"set_new_IREP\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_ISAI\",\"type\":\"address\"}],\"name\":\"set_new_ISAI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_ISUSD\",\"type\":\"address\"}],\"name\":\"set_new_ISUSD\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_IUSDC\",\"type\":\"address\"}],\"name\":\"set_new_IUSDC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_IWBTC\",\"type\":\"address\"}],\"name\":\"set_new_IWBTC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_IZRX\",\"type\":\"address\"}],\"name\":\"set_new_IZRX\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AprBin is the compiled bytecode used for deploying new contracts.
var AprBin = "0x6080604052670de0b6b3a76400006001553480156200001d57600080fd5b506200002e62000d1f60201b60201c565b6000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a333600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550731e0447b19bb6ecfdae1e4ae1694b0c3659614e4e600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507324a42fd28c976a61df5d00d0599c34c4f90748c8600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550735d3a536e4d6dbd6114cc1ead35777bab948e3643600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550736c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550734ddc2d193948926d02f9b1fe9e1daa0718270ed5600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073158079ee67fce2f58472a96584a73c7ab9ac95c1600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073f5dce57282a584d2746faf1593d3121fcac444dc600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507339aa39c021dfbae8fac545936693ac917d5e7563600c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073c11b1268c1a384e55c48c2391d8d480264a3a7f4600d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073b3319f5d18bc0d84dd1b4825dcde5d5f7266d407600e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073a7eb2bc82df18013ecc2a6c533fc29446442edee600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073bd56e9477fc6997609cf45f84795efbdac642ff1601060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550731cc9567ea2eb740824a45f8026ccf8e46973234d601160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073ba9262578efef8b3aff7f60cd629d6cc8859c8b5601360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073f013406a0b1d544238083df0b93ad0d2cbe0f65f601460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507377f973fcaf871459aa58cd81881ce453759281bc601560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507314094949152eddbfcd073717200da82fed8dc960601660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073493c57c4763932315a328269e1adad09653b9081601760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550731d496da96caf6b518b133736beca85d5c4f9cbc5601860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507349f4592e641820e928f9919ef4abd92a719b4b49601960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550736b175474e89094c44da98b954eedeac495271d0f601a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506e085d4780b73119b644ae5ecd22b376601b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48601c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073dac17f958d2ee523a2206206994597c13d831ec7601d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507357ab1ec28d129707052df4df418d58a2d46d5f51601e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507380fb784b7ed66730e8b1dbd9820afd29931aab03601f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730d8775f648430679a709e98d2b0cb6250d2887ef602060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee602160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073514910771af9ca656af840dff83e8264ecf986ca602260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073dd974d5c2e2928dea5f71b9825b8b646686bd200602360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550731985365e9f78359a9b6ad760e32412f4a445e862602460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550739f8f72aa9304c8b593d555f12ef6589cc3a579a2602560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730f5d2fb29fb7d3cfee444a200298f468908cc942602660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073e41d2489571d322189246dafa5ebde1f4699f498602760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073c011a73ee8576fb46f5e1c5751ca3b9fe0af2a6f602860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550732260fac5e5542a773aa44fbcfedf7c193bc2c599602960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062000d27565b600033905090565b6159388062000d376000396000f3fe608060405234801561001057600080fd5b50600436106107d15760003560e01c806378226b4111610400578063b548780111610220578063d2db39fe11610130578063e9885aeb116100c3578063f53a062811610092578063f53a06281461171c578063f60a15ed14611738578063f82ce27d14611742578063f8b5fff014611760578063ff3c269e1461177e576107d1565b8063e9885aeb146116aa578063eb22a022146116c6578063f2fde38b146116e4578063f4069cba14611700576107d1565b8063dae180b8116100ff578063dae180b814611634578063e1f2aeff14611652578063e49f09e414611670578063e8d791481461168c576107d1565b8063d2db39fe146115c0578063d6a79ada146115dc578063d9c75e37146115f8578063dada70dc14611616576107d1565b8063bfd5b402116101b3578063c6882b7411610182578063c6882b741461152e578063c6b0d25f1461154c578063c9f8217014611568578063ccc4168114611586578063cda08710146115a2576107d1565b8063bfd5b402146114ba578063c043fca2146114d6578063c30c7cae146114f4578063c552934d14611512576107d1565b8063bcbe1b52116101ef578063bcbe1b5214611444578063bdbfcd1c14611462578063be2134eb14611480578063bfaf715b1461149e576107d1565b8063b5487801146113d0578063b72d34c2146113ee578063b74d4f311461140a578063bb540c8214611426576107d1565b80638f32d59b1161031b578063a22db74c116102ae578063ad9a5e2a1161027d578063ad9a5e2a1461133e578063aee7b9f51461135a578063b0ac4a7314611378578063b35946e214611394578063b4650bf7146113b2576107d1565b8063a22db74c146112b4578063a37b9e7f146112d2578063a747b93b146112f0578063aa00070b14611320576107d1565b8063983a97d9116102ea578063983a97d91461123357806398df5ccf1461125c5780639be0587e14611278578063a0bc8d5314611296576107d1565b80638f32d59b146111bb57806390b2ee81146111d957806397c193c5146111f7578063982e2adf14611215576107d1565b8063832b4b0311610393578063878f760311610362578063878f76031461114357806389257a1214611161578063892eac131461117f5780638da5cb5b1461119d576107d1565b8063832b4b03146110bd57806383deca3d146110db578063857c3c9b1461110b57806386bbe19614611127576107d1565b80637f8e4f64116103cf5780637f8e4f641461102c578063812adb061461104a57806381d24d8d146110665780638264d10914611096576107d1565b806378226b4114610fb5578063793957b714610fd15780637cde729f14610fed5780637dc0d1d01461100e576107d1565b80633cf1e7b8116105f65780634e9a10161161050657806361fd719411610499578063693d42e711610468578063693d42e714610f35578063715018a614610f5357806372142a2b14610f5d5780637352023f14610f7b578063750da89914610f99576107d1565b806361fd719414610ebd578063638aa95b14610edb57806363a6894014610ef9578063653d8ad914610f17576107d1565b80635af11ab1116104d55780635af11ab114610e455780635ce7174f14610e635780635d5bd5d314610e815780636109682014610e9f576107d1565b80634e9a101614610dcf5780635309d95214610ded578063536f263214610e095780635ac1baa914610e27576107d1565b806344e258b4116105895780634907c7ae116105585780634907c7ae14610d3d57806349386b5e14610d5b5780634958aa0714610d795780634d2cd8ab14610d955780634e24f77f14610db3576107d1565b806344e258b414610cac578063454ccd7114610cdc578063470bce8014610d0157806348ccda3c14610d1f576107d1565b806341976e09116105c557806341976e0914610c24578063429db17a14610c5457806343d39c8b14610c7057806343e8b0be14610c8e576107d1565b80633cf1e7b814610bae5780633f8ce23f14610bcc5780633f9c862214610be85780633fa6a66b14610c06576107d1565b80631c384781116106f15780632e163c86116106845780633196ebbd116106535780633196ebbd14610b1a578063354c671714610b3857806335b3744214610b5657806338118ab414610b7457806339ec407b14610b90576107d1565b80632e163c8614610aa65780632e8d6e1814610ac25780632ea3546014610ade5780633151a24514610afc576107d1565b806325e079b1116106c057806325e079b114610a345780632735b36314610a5057806328e8452d14610a6c5780632c7a5ae414610a8a576107d1565b80631c384781146109be5780631cc16a6a146109da57806322cc56d9146109f8578063243d665914610a16576107d1565b8063101b56681161076957806318b3b5091161073857806318b3b50914610936578063192a9861146109545780631b1be2bf146109845780631b856311146109a2576107d1565b8063101b5668146108c0578063106f39d1146108dc57806310b4acdc146108fa5780631800298314610918576107d1565b80630b58a7a3116107a55780630b58a7a31461084a5780630dc00a37146108685780630dc0849a146108845780630e1e23e2146108a2576107d1565b8062e4768b146107d657806302c93c70146107f2578063033672b71461081057806304eb246d1461082c575b600080fd5b6107f060048036036107eb91908101906150e7565b61179c565b005b6107fa61183b565b604051610807919061550a565b60405180910390f35b61082a60048036036108259190810190615095565b61186d565b005b610834611908565b6040516108419190615468565b60405180910390f35b610852611ab9565b60405161085f9190615409565b60405180910390f35b610882600480360361087d9190810190615095565b611adf565b005b61088c611b7a565b6040516108999190615409565b60405180910390f35b6108aa611ba0565b6040516108b7919061550a565b60405180910390f35b6108da60048036036108d59190810190615095565b611bd2565b005b6108e4611c6d565b6040516108f1919061550a565b60405180910390f35b610902611c9f565b60405161090f919061550a565b60405180910390f35b610920611cd1565b60405161092d919061550a565b60405180910390f35b61093e611d03565b60405161094b9190615409565b60405180910390f35b61096e60048036036109699190810190615095565b611d29565b60405161097b919061550a565b60405180910390f35b61098c611e75565b6040516109999190615409565b60405180910390f35b6109bc60048036036109b79190810190615095565b611e9b565b005b6109d860048036036109d39190810190615095565b611f36565b005b6109e2611fd1565b6040516109ef9190615409565b60405180910390f35b610a00611ff7565b604051610a0d919061550a565b60405180910390f35b610a1e612029565b604051610a2b919061550a565b60405180910390f35b610a4e6004803603610a499190810190615095565b61205b565b005b610a6a6004803603610a659190810190615095565b6120f6565b005b610a74612191565b604051610a819190615409565b60405180910390f35b610aa46004803603610a9f9190810190615095565b6121b7565b005b610ac06004803603610abb9190810190615095565b612252565b005b610adc6004803603610ad7919081019061514c565b6122ed565b005b610ae6612453565b604051610af3919061550a565b60405180910390f35b610b04612485565b604051610b11919061550a565b60405180910390f35b610b226124b7565b604051610b2f919061550a565b60405180910390f35b610b406124e9565b604051610b4d9190615409565b60405180910390f35b610b5e61250f565b604051610b6b919061550a565b60405180910390f35b610b8e6004803603610b899190810190615095565b612541565b005b610b986125dc565b604051610ba59190615409565b60405180910390f35b610bb6612602565b604051610bc3919061550a565b60405180910390f35b610be66004803603610be19190810190615095565b612634565b005b610bf06126cf565b604051610bfd919061550a565b60405180910390f35b610c0e6126e0565b604051610c1b9190615409565b60405180910390f35b610c3e6004803603610c399190810190615095565b612706565b604051610c4b919061550a565b60405180910390f35b610c6e6004803603610c699190810190615095565b61274f565b005b610c786127ea565b604051610c859190615409565b60405180910390f35b610c96612810565b604051610ca3919061550a565b60405180910390f35b610cc66004803603610cc19190810190615095565b612842565b604051610cd3919061550a565b60405180910390f35b610ce46128de565b604051610cf898979695949392919061556a565b60405180910390f35b610d09612944565b604051610d16919061550a565b60405180910390f35b610d27612955565b604051610d349190615409565b60405180910390f35b610d4561297b565b604051610d52919061550a565b60405180910390f35b610d636129ad565b604051610d709190615409565b60405180910390f35b610d936004803603610d8e9190810190615095565b6129d3565b005b610d9d612a6e565b604051610daa919061550a565b60405180910390f35b610dcd6004803603610dc89190810190615095565b612aa0565b005b610dd7612b3b565b604051610de4919061550a565b60405180910390f35b610e076004803603610e029190810190615095565b612b6d565b005b610e11612c08565b604051610e1e9190615409565b60405180910390f35b610e2f612c2e565b604051610e3c919061550a565b60405180910390f35b610e4d612c60565b604051610e5a919061550a565b60405180910390f35b610e6b612c71565b604051610e78919061550a565b60405180910390f35b610e89612ca3565b604051610e96919061550a565b60405180910390f35b610ea7612cd5565b604051610eb49190615409565b60405180910390f35b610ec5612cfb565b604051610ed29190615409565b60405180910390f35b610ee3612d21565b604051610ef0919061550a565b60405180910390f35b610f01612d53565b604051610f0e9190615409565b60405180910390f35b610f1f612d79565b604051610f2c919061550a565b60405180910390f35b610f3d612d8a565b604051610f4a919061550a565b60405180910390f35b610f5b612dbc565b005b610f65612ec2565b604051610f729190615409565b60405180910390f35b610f83612ee8565b604051610f90919061550a565b60405180910390f35b610fb36004803603610fae9190810190615095565b612f1a565b005b610fcf6004803603610fca9190810190615095565b612fb5565b005b610feb6004803603610fe69190810190615095565b613050565b005b610ff56130eb565b6040516110059493929190615525565b60405180910390f35b61101661311f565b6040516110239190615409565b60405180910390f35b611034613145565b6040516110419190615409565b60405180910390f35b611064600480360361105f9190810190615095565b61316b565b005b611080600480360361107b91908101906151c7565b613206565b60405161108d919061550a565b60405180910390f35b61109e61352f565b6040516110b49a999897969594939291906155e8565b60405180910390f35b6110c56135ae565b6040516110d2919061550a565b60405180910390f35b6110f560048036036110f09190810190615095565b6135e0565b604051611102919061550a565b60405180910390f35b61112560048036036111209190810190615095565b61367a565b005b611141600480360361113c9190810190615095565b613715565b005b61114b6137b0565b6040516111589190615409565b60405180910390f35b6111696137d6565b6040516111769190615409565b60405180910390f35b6111876137fc565b6040516111949190615468565b60405180910390f35b6111a56139ad565b6040516111b29190615409565b60405180910390f35b6111c36139d6565b6040516111d0919061544d565b60405180910390f35b6111e1613a34565b6040516111ee9190615409565b60405180910390f35b6111ff613a5a565b60405161120c9190615409565b60405180910390f35b61121d613a80565b60405161122a9190615409565b60405180910390f35b61123b613aa6565b6040516112539c9b9a99989796959493929190615684565b60405180910390f35b61127660048036036112719190810190615095565b613b3e565b005b611280613bd9565b60405161128d9190615409565b60405180910390f35b61129e613bff565b6040516112ab919061550a565b60405180910390f35b6112bc613c31565b6040516112c99190615409565b60405180910390f35b6112da613c57565b6040516112e79190615409565b60405180910390f35b61130a60048036036113059190810190615095565b613c7d565b604051611317919061550a565b60405180910390f35b611328613cc6565b6040516113359190615409565b60405180910390f35b61135860048036036113539190810190615095565b613cec565b005b611362613d87565b60405161136f919061550a565b60405180910390f35b611392600480360361138d9190810190615095565b613db9565b005b61139c613e54565b6040516113a9919061550a565b60405180910390f35b6113ba613e86565b6040516113c7919061550a565b60405180910390f35b6113d8613eb8565b6040516113e59190615468565b60405180910390f35b61140860048036036114039190810190615095565b614069565b005b611424600480360361141f9190810190615095565b614104565b005b61142e61419f565b60405161143b919061550a565b60405180910390f35b61144c6141d1565b6040516114599190615409565b60405180910390f35b61146a6141f7565b604051611477919061550a565b60405180910390f35b611488614229565b604051611495919061550a565b60405180910390f35b6114b860048036036114b39190810190615095565b61425b565b005b6114d460048036036114cf9190810190615095565b6142f6565b005b6114de614391565b6040516114eb9190615409565b60405180910390f35b6114fc6143b7565b6040516115099190615409565b60405180910390f35b61152c60048036036115279190810190615095565b6143dd565b005b611536614478565b604051611543919061550a565b60405180910390f35b61156660048036036115619190810190615095565b6144aa565b005b611570614545565b60405161157d919061550a565b60405180910390f35b6115a0600480360361159b9190810190615095565b614577565b005b6115aa614612565b6040516115b79190615409565b60405180910390f35b6115da60048036036115d59190810190615095565b614638565b005b6115f660048036036115f19190810190615095565b6146d3565b005b61160061476e565b60405161160d9190615409565b60405180910390f35b61161e614794565b60405161162b9190615409565b60405180910390f35b61163c6147ba565b604051611649919061550a565b60405180910390f35b61165a6147ec565b6040516116679190615409565b60405180910390f35b61168a60048036036116859190810190615095565b614812565b005b6116946148ad565b6040516116a1919061550a565b60405180910390f35b6116c460048036036116bf9190810190615095565b6148df565b005b6116ce61497a565b6040516116db9190615409565b60405180910390f35b6116fe60048036036116f99190810190615095565b6149a0565b005b61171a600480360361171591908101906150e7565b6149f3565b005b61173660048036036117319190810190615095565b614a92565b005b611740614b2d565b005b61174a614c3a565b6040516117579190615409565b60405180910390f35b611768614ce1565b6040516117759190615409565b60405180910390f35b611786614d07565b6040516117939190615409565b60405180910390f35b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156118375780600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b5050565b6000611868601d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611d29565b905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156119055780600d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b606060008090506000611919612453565b905081811115611927578091505b6000611931611c6d565b90508281111561193f578092505b60006119496126cf565b905083811115611957578093505b60006119616124b7565b90508481111561196f578094505b60606040518060400160405280600481526020017f6e6f6e65000000000000000000000000000000000000000000000000000000008152509050848614156119ea576040518060400160405280600881526020017f436f6d706f756e6400000000000000000000000000000000000000000000000081525090505b83861415611a2b576040518060400160405280600781526020017f46756c6372756d0000000000000000000000000000000000000000000000000081525090505b82861415611a6c576040518060400160405280600481526020017f645964580000000000000000000000000000000000000000000000000000000081525090505b81861415611aad576040518060400160405280600481526020017f416176650000000000000000000000000000000000000000000000000000000081525090505b80965050505050505090565b601b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415611b775780602360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b601d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000611bcd601960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166135e0565b905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415611c6a5780601e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b6000611c9a601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166135e0565b905090565b6000611ccc600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612842565b905090565b6000611cfe600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612842565b905090565b601860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ed6ff7606040518163ffffffff1660e01b815260040160206040518083038186803b158015611d9457600080fd5b505afa158015611da8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611dcc91908101906150be565b9050611e6d633b9aca008273ffffffffffffffffffffffffffffffffffffffff1663c540148e866040518263ffffffff1660e01b8152600401611e0f9190615409565b60206040518083038186803b158015611e2757600080fd5b505afa158015611e3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611e5f91908101906151f0565b614d2d90919063ffffffff16565b915050919050565b601c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415611f335780601960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415611fce5780600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b601260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000612024601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166135e0565b905090565b6000612056601a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611d29565b905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156120f35780601d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561218e5780601860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561224f5780600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156122ea5780600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b6122f56139d6565b612334576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161232b906154ea565b60405180910390fd5b60008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161236f9190615409565b60206040518083038186803b15801561238757600080fd5b505afa15801561239b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506123bf91908101906151f0565b90508173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b81526004016123fc929190615424565b602060405180830381600087803b15801561241657600080fd5b505af115801561242a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061244e9190810190615123565b505050565b6000612480600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612842565b905090565b60006124b2602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611d29565b905090565b60006124e4601c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611d29565b905090565b602360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600061253c601b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611d29565b905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156125d95780600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b601a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600061262f601660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166135e0565b905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156126cc5780602460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b60006126db6002613206565b905090565b601460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156127e75780601560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b601960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600061283d601760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166135e0565b905090565b60006128d7622014808373ffffffffffffffffffffffffffffffffffffffff1663ae9d70b06040518163ffffffff1660e01b815260040160206040518083038186803b15801561289157600080fd5b505afa1580156128a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506128c991908101906151f0565b614d7790919063ffffffff16565b9050919050565b6000806000806000806000806128f2612a6e565b6128fa611c9f565b612902613d87565b61290a611cd1565b612912612d21565b61291a612453565b6129226147ba565b61292a614545565b975097509750975097509750975097509091929394959697565b60006129506003613206565b905090565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006129a8602360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611d29565b905090565b602760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415612a6b5780602860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b6000612a9b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612842565b905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415612b385780602760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b6000612b68602860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611d29565b905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415612c055780601b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b601e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000612c5b601e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611d29565b905090565b6000612c6c6000613206565b905090565b6000612c9e602060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611d29565b905090565b6000612cd0601860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166135e0565b905090565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000612d4e600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612842565b905090565b601160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000612d856001613206565b905090565b6000612db7602460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611d29565b905090565b612dc46139d6565b612e03576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612dfa906154ea565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000612f15601060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166135e0565b905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415612fb25780601460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561304d5780600e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156130e85780601f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b6000806000806130f9612d79565b613101612c60565b6131096126cf565b613111612944565b935093509350935090919293565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156132035780600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600080600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fd47eda6846040518263ffffffff1660e01b8152600401613264919061550a565b60206040518083038186803b15801561327c57600080fd5b505afa158015613290573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506132b4919081019061519e565b60000151905060006301e28500820290506000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cb04a34c866040518263ffffffff1660e01b8152600401613322919061550a565b604080518083038186803b15801561333957600080fd5b505afa15801561334d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506133719190810190615175565b600001516fffffffffffffffffffffffffffffffff1690506000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cb04a34c876040518263ffffffff1660e01b81526004016133e6919061550a565b604080518083038186803b1580156133fd57600080fd5b505afa158015613411573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506134359190810190615175565b602001516fffffffffffffffffffffffffffffffff16905060008160015484028161345c57fe5b0490506000600154600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e55202286040518163ffffffff1660e01b815260040160206040518083038186803b1580156134cc57600080fd5b505afa1580156134e0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250613504919081019061519e565b600001516001548488028161351557fe5b04028161351e57fe5b049050809650505050505050919050565b6000806000806000806000806000806135466141f7565b61354e612ee8565b61355661419f565b61355e613e54565b613566611c6d565b61356e611ff7565b613576612602565b61357e612810565b613586612ca3565b61358e611ba0565b995099509950995099509950995099509950995090919293949596979899565b60006135db601f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611d29565b905090565b600061367360648373ffffffffffffffffffffffffffffffffffffffff166309ec6b6b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561362d57600080fd5b505afa158015613641573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061366591908101906151f0565b614d2d90919063ffffffff16565b9050919050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156137125780601760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156137ad5780600c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606000809050600061380d613d87565b90508181111561381b578091505b6000613825611ff7565b905082811115613833578092505b600061383d612c60565b90508381111561384b578093505b60006138556148ad565b905084811115613863578094505b60606040518060400160405280600481526020017f6e6f6e65000000000000000000000000000000000000000000000000000000008152509050848614156138de576040518060400160405280600881526020017f436f6d706f756e6400000000000000000000000000000000000000000000000081525090505b8386141561391f576040518060400160405280600781526020017f46756c6372756d0000000000000000000000000000000000000000000000000081525090505b82861415613960576040518060400160405280600481526020017f645964580000000000000000000000000000000000000000000000000000000081525090505b818614156139a1576040518060400160405280600481526020017f416176650000000000000000000000000000000000000000000000000000000081525090505b80965050505050505090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16613a18614de7565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b601060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600080600080600080600080600080613ac0612029565b613ac861250f565b613ad06124b7565b613ad861183b565b613ae0612c2e565b613ae8612c71565b613af06148ad565b613af8612485565b613b0061297b565b613b08612d8a565b613b10613bff565b613b18612b3b565b9b509b509b509b509b509b509b509b509b509b509b509b50909192939495969798999a9b565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415613bd65780602960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b601560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000613c2c602760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611d29565b905090565b600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b602060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415613d845780600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b6000613db4600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612842565b905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415613e515780601060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b6000613e81601360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166135e0565b905090565b6000613eb3602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611d29565b905090565b606060008090506000613ec9612a6e565b905081811115613ed7578091505b6000613ee1612810565b905082811115613eef578092505b6000613ef9612944565b905083811115613f07578093505b6000613f11612029565b905084811115613f1f578094505b60606040518060400160405280600481526020017f6e6f6e6500000000000000000000000000000000000000000000000000000000815250905084861415613f9a576040518060400160405280600881526020017f436f6d706f756e6400000000000000000000000000000000000000000000000081525090505b83861415613fdb576040518060400160405280600781526020017f46756c6372756d0000000000000000000000000000000000000000000000000081525090505b8286141561401c576040518060400160405280600481526020017f645964580000000000000000000000000000000000000000000000000000000081525090505b8186141561405d576040518060400160405280600481526020017f416176650000000000000000000000000000000000000000000000000000000081525090505b80965050505050505090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156141015780602060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561419c5780600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b60006141cc601160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166135e0565b905090565b602660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000614224600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166135e0565b905090565b6000614256602560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611d29565b905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156142f35780602560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561438e5780602660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156144755780601660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b60006144a5602960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611d29565b905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156145425780600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b6000614572600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612842565b905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561460f5780602160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b601660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156146d05780602260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561476b5780601a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b601760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006147e7600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612842565b905090565b601f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156148aa5780601360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b60006148da602160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611d29565b905090565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156149775780601160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6149a86139d6565b6149e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016149de906154ea565b60405180910390fd5b6149f081614def565b50565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415614a8e5780600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b5050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415614b2a5780601c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b614b356139d6565b614b74576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614b6b906154ea565b60405180910390fd5b60003373ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1631604051614bb1906153f4565b60006040518083038185875af1925050503d8060008114614bee576040519150601f19603f3d011682016040523d82523d6000602084013e614bf3565b606091505b5050905080614c37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614c2e906154aa565b60405180910390fd5b50565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ed6ff7606040518163ffffffff1660e01b815260040160206040518083038186803b158015614ca457600080fd5b505afa158015614cb8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250614cdc91908101906150be565b905090565b600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b602260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000614d6f83836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250614f1d565b905092915050565b600080831415614d8a5760009050614de1565b6000828402905082848281614d9b57fe5b0414614ddc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614dd3906154ca565b60405180910390fd5b809150505b92915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415614e5f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614e569061548a565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008083118290614f64576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614f5b9190615468565b60405180910390fd5b506000838581614f7057fe5b049050809150509392505050565b600081359050614f8d81615882565b92915050565b600081519050614fa281615882565b92915050565b600081519050614fb781615899565b92915050565b600081359050614fcc816158b0565b92915050565b600060408284031215614fe457600080fd5b614fee604061573e565b90506000614ffe84828501615056565b600083015250602061501284828501615056565b60208301525092915050565b60006020828403121561503057600080fd5b61503a602061573e565b9050600061504a84828501615080565b60008301525092915050565b600081519050615065816158c7565b92915050565b60008135905061507a816158de565b92915050565b60008151905061508f816158de565b92915050565b6000602082840312156150a757600080fd5b60006150b584828501614f7e565b91505092915050565b6000602082840312156150d057600080fd5b60006150de84828501614f93565b91505092915050565b600080604083850312156150fa57600080fd5b600061510885828601614f7e565b92505060206151198582860161506b565b9150509250929050565b60006020828403121561513557600080fd5b600061514384828501614fa8565b91505092915050565b60006020828403121561515e57600080fd5b600061516c84828501614fbd565b91505092915050565b60006040828403121561518757600080fd5b600061519584828501614fd2565b91505092915050565b6000602082840312156151b057600080fd5b60006151be8482850161501e565b91505092915050565b6000602082840312156151d957600080fd5b60006151e78482850161506b565b91505092915050565b60006020828403121561520257600080fd5b600061521084828501615080565b91505092915050565b61522281615808565b82525050565b61523181615792565b82525050565b615240816157a4565b82525050565b60006152518261576b565b61525b8185615781565b935061526b81856020860161583e565b61527481615871565b840191505092915050565b600061528c602683615781565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006152f2601683615781565b91507f7472616e73666572206f6620455448206661696c6564000000000000000000006000830152602082019050919050565b6000615332602183615781565b91507f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f60008301527f77000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000615398602083615781565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b60006153d8600083615776565b9150600082019050919050565b6153ee816157fe565b82525050565b60006153ff826153cb565b9150819050919050565b600060208201905061541e6000830184615228565b92915050565b60006040820190506154396000830185615219565b61544660208301846153e5565b9392505050565b60006020820190506154626000830184615237565b92915050565b600060208201905081810360008301526154828184615246565b905092915050565b600060208201905081810360008301526154a38161527f565b9050919050565b600060208201905081810360008301526154c3816152e5565b9050919050565b600060208201905081810360008301526154e381615325565b9050919050565b600060208201905081810360008301526155038161538b565b9050919050565b600060208201905061551f60008301846153e5565b92915050565b600060808201905061553a60008301876153e5565b61554760208301866153e5565b61555460408301856153e5565b61556160608301846153e5565b95945050505050565b600061010082019050615580600083018b6153e5565b61558d602083018a6153e5565b61559a60408301896153e5565b6155a760608301886153e5565b6155b460808301876153e5565b6155c160a08301866153e5565b6155ce60c08301856153e5565b6155db60e08301846153e5565b9998505050505050505050565b6000610140820190506155fe600083018d6153e5565b61560b602083018c6153e5565b615618604083018b6153e5565b615625606083018a6153e5565b61563260808301896153e5565b61563f60a08301886153e5565b61564c60c08301876153e5565b61565960e08301866153e5565b6156676101008301856153e5565b6156756101208301846153e5565b9b9a5050505050505050505050565b60006101808201905061569a600083018f6153e5565b6156a7602083018e6153e5565b6156b4604083018d6153e5565b6156c1606083018c6153e5565b6156ce608083018b6153e5565b6156db60a083018a6153e5565b6156e860c08301896153e5565b6156f560e08301886153e5565b6157036101008301876153e5565b6157116101208301866153e5565b61571f6101408301856153e5565b61572d6101608301846153e5565b9d9c50505050505050505050505050565b6000604051905081810181811067ffffffffffffffff8211171561576157600080fd5b8060405250919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b600061579d826157de565b9050919050565b60008115159050919050565b60006157bb82615792565b9050919050565b60006fffffffffffffffffffffffffffffffff82169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006158138261581a565b9050919050565b60006158258261582c565b9050919050565b6000615837826157de565b9050919050565b60005b8381101561585c578082015181840152602081019050615841565b8381111561586b576000848401525b50505050565b6000601f19601f8301169050919050565b61588b81615792565b811461589657600080fd5b50565b6158a2816157a4565b81146158ad57600080fd5b50565b6158b9816157b0565b81146158c457600080fd5b50565b6158d0816157c2565b81146158db57600080fd5b50565b6158e7816157fe565b81146158f257600080fd5b5056fea365627a7a723158205b339a38c76beadfc541fcf06ed73d5cc41204ac5c9f47fe25622b25f17ad7c66c6578706572696d656e74616cf564736f6c634300050c0040"

// DeployApr deploys a new Ethereum contract, binding an instance of Apr to it.
func DeployApr(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Apr, error) {
	parsed, err := abi.JSON(strings.NewReader(AprABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AprBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Apr{AprCaller: AprCaller{contract: contract}, AprTransactor: AprTransactor{contract: contract}, AprFilterer: AprFilterer{contract: contract}}, nil
}

// Apr is an auto generated Go binding around an Ethereum contract.
type Apr struct {
	AprCaller     // Read-only binding to the contract
	AprTransactor // Write-only binding to the contract
	AprFilterer   // Log filterer for contract events
}

// AprCaller is an auto generated read-only Go binding around an Ethereum contract.
type AprCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AprTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AprTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AprFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AprFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AprSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AprSession struct {
	Contract     *Apr              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AprCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AprCallerSession struct {
	Contract *AprCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AprTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AprTransactorSession struct {
	Contract     *AprTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AprRaw is an auto generated low-level Go binding around an Ethereum contract.
type AprRaw struct {
	Contract *Apr // Generic contract binding to access the raw methods on
}

// AprCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AprCallerRaw struct {
	Contract *AprCaller // Generic read-only contract binding to access the raw methods on
}

// AprTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AprTransactorRaw struct {
	Contract *AprTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApr creates a new instance of Apr, bound to a specific deployed contract.
func NewApr(address common.Address, backend bind.ContractBackend) (*Apr, error) {
	contract, err := bindApr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Apr{AprCaller: AprCaller{contract: contract}, AprTransactor: AprTransactor{contract: contract}, AprFilterer: AprFilterer{contract: contract}}, nil
}

// NewAprCaller creates a new read-only instance of Apr, bound to a specific deployed contract.
func NewAprCaller(address common.Address, caller bind.ContractCaller) (*AprCaller, error) {
	contract, err := bindApr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AprCaller{contract: contract}, nil
}

// NewAprTransactor creates a new write-only instance of Apr, bound to a specific deployed contract.
func NewAprTransactor(address common.Address, transactor bind.ContractTransactor) (*AprTransactor, error) {
	contract, err := bindApr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AprTransactor{contract: contract}, nil
}

// NewAprFilterer creates a new log filterer instance of Apr, bound to a specific deployed contract.
func NewAprFilterer(address common.Address, filterer bind.ContractFilterer) (*AprFilterer, error) {
	contract, err := bindApr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AprFilterer{contract: contract}, nil
}

// bindApr binds a generic wrapper to an already deployed contract.
func bindApr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AprABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Apr *AprRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Apr.Contract.AprCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Apr *AprRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apr.Contract.AprTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Apr *AprRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Apr.Contract.AprTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Apr *AprCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Apr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Apr *AprTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Apr *AprTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Apr.Contract.contract.Transact(opts, method, params...)
}

// AAVE is a free data retrieval call binding the contract method 0x48ccda3c.
//
// Solidity: function AAVE() constant returns(address)
func (_Apr *AprCaller) AAVE(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "AAVE")
	return *ret0, err
}

// AAVE is a free data retrieval call binding the contract method 0x48ccda3c.
//
// Solidity: function AAVE() constant returns(address)
func (_Apr *AprSession) AAVE() (common.Address, error) {
	return _Apr.Contract.AAVE(&_Apr.CallOpts)
}

// AAVE is a free data retrieval call binding the contract method 0x48ccda3c.
//
// Solidity: function AAVE() constant returns(address)
func (_Apr *AprCallerSession) AAVE() (common.Address, error) {
	return _Apr.Contract.AAVE(&_Apr.CallOpts)
}

// ABAT is a free data retrieval call binding the contract method 0xaa00070b.
//
// Solidity: function ABAT() constant returns(address)
func (_Apr *AprCaller) ABAT(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "ABAT")
	return *ret0, err
}

// ABAT is a free data retrieval call binding the contract method 0xaa00070b.
//
// Solidity: function ABAT() constant returns(address)
func (_Apr *AprSession) ABAT() (common.Address, error) {
	return _Apr.Contract.ABAT(&_Apr.CallOpts)
}

// ABAT is a free data retrieval call binding the contract method 0xaa00070b.
//
// Solidity: function ABAT() constant returns(address)
func (_Apr *AprCallerSession) ABAT() (common.Address, error) {
	return _Apr.Contract.ABAT(&_Apr.CallOpts)
}

// ADAI is a free data retrieval call binding the contract method 0x39ec407b.
//
// Solidity: function ADAI() constant returns(address)
func (_Apr *AprCaller) ADAI(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "ADAI")
	return *ret0, err
}

// ADAI is a free data retrieval call binding the contract method 0x39ec407b.
//
// Solidity: function ADAI() constant returns(address)
func (_Apr *AprSession) ADAI() (common.Address, error) {
	return _Apr.Contract.ADAI(&_Apr.CallOpts)
}

// ADAI is a free data retrieval call binding the contract method 0x39ec407b.
//
// Solidity: function ADAI() constant returns(address)
func (_Apr *AprCallerSession) ADAI() (common.Address, error) {
	return _Apr.Contract.ADAI(&_Apr.CallOpts)
}

// AETH is a free data retrieval call binding the contract method 0x89257a12.
//
// Solidity: function AETH() constant returns(address)
func (_Apr *AprCaller) AETH(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "AETH")
	return *ret0, err
}

// AETH is a free data retrieval call binding the contract method 0x89257a12.
//
// Solidity: function AETH() constant returns(address)
func (_Apr *AprSession) AETH() (common.Address, error) {
	return _Apr.Contract.AETH(&_Apr.CallOpts)
}

// AETH is a free data retrieval call binding the contract method 0x89257a12.
//
// Solidity: function AETH() constant returns(address)
func (_Apr *AprCallerSession) AETH() (common.Address, error) {
	return _Apr.Contract.AETH(&_Apr.CallOpts)
}

// AKNC is a free data retrieval call binding the contract method 0x354c6717.
//
// Solidity: function AKNC() constant returns(address)
func (_Apr *AprCaller) AKNC(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "AKNC")
	return *ret0, err
}

// AKNC is a free data retrieval call binding the contract method 0x354c6717.
//
// Solidity: function AKNC() constant returns(address)
func (_Apr *AprSession) AKNC() (common.Address, error) {
	return _Apr.Contract.AKNC(&_Apr.CallOpts)
}

// AKNC is a free data retrieval call binding the contract method 0x354c6717.
//
// Solidity: function AKNC() constant returns(address)
func (_Apr *AprCallerSession) AKNC() (common.Address, error) {
	return _Apr.Contract.AKNC(&_Apr.CallOpts)
}

// ALEND is a free data retrieval call binding the contract method 0xe1f2aeff.
//
// Solidity: function ALEND() constant returns(address)
func (_Apr *AprCaller) ALEND(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "ALEND")
	return *ret0, err
}

// ALEND is a free data retrieval call binding the contract method 0xe1f2aeff.
//
// Solidity: function ALEND() constant returns(address)
func (_Apr *AprSession) ALEND() (common.Address, error) {
	return _Apr.Contract.ALEND(&_Apr.CallOpts)
}

// ALEND is a free data retrieval call binding the contract method 0xe1f2aeff.
//
// Solidity: function ALEND() constant returns(address)
func (_Apr *AprCallerSession) ALEND() (common.Address, error) {
	return _Apr.Contract.ALEND(&_Apr.CallOpts)
}

// ALINK is a free data retrieval call binding the contract method 0xff3c269e.
//
// Solidity: function ALINK() constant returns(address)
func (_Apr *AprCaller) ALINK(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "ALINK")
	return *ret0, err
}

// ALINK is a free data retrieval call binding the contract method 0xff3c269e.
//
// Solidity: function ALINK() constant returns(address)
func (_Apr *AprSession) ALINK() (common.Address, error) {
	return _Apr.Contract.ALINK(&_Apr.CallOpts)
}

// ALINK is a free data retrieval call binding the contract method 0xff3c269e.
//
// Solidity: function ALINK() constant returns(address)
func (_Apr *AprCallerSession) ALINK() (common.Address, error) {
	return _Apr.Contract.ALINK(&_Apr.CallOpts)
}

// AMANA is a free data retrieval call binding the contract method 0xbcbe1b52.
//
// Solidity: function AMANA() constant returns(address)
func (_Apr *AprCaller) AMANA(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "AMANA")
	return *ret0, err
}

// AMANA is a free data retrieval call binding the contract method 0xbcbe1b52.
//
// Solidity: function AMANA() constant returns(address)
func (_Apr *AprSession) AMANA() (common.Address, error) {
	return _Apr.Contract.AMANA(&_Apr.CallOpts)
}

// AMANA is a free data retrieval call binding the contract method 0xbcbe1b52.
//
// Solidity: function AMANA() constant returns(address)
func (_Apr *AprCallerSession) AMANA() (common.Address, error) {
	return _Apr.Contract.AMANA(&_Apr.CallOpts)
}

// AMKR is a free data retrieval call binding the contract method 0x7f8e4f64.
//
// Solidity: function AMKR() constant returns(address)
func (_Apr *AprCaller) AMKR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "AMKR")
	return *ret0, err
}

// AMKR is a free data retrieval call binding the contract method 0x7f8e4f64.
//
// Solidity: function AMKR() constant returns(address)
func (_Apr *AprSession) AMKR() (common.Address, error) {
	return _Apr.Contract.AMKR(&_Apr.CallOpts)
}

// AMKR is a free data retrieval call binding the contract method 0x7f8e4f64.
//
// Solidity: function AMKR() constant returns(address)
func (_Apr *AprCallerSession) AMKR() (common.Address, error) {
	return _Apr.Contract.AMKR(&_Apr.CallOpts)
}

// AREP is a free data retrieval call binding the contract method 0x97c193c5.
//
// Solidity: function AREP() constant returns(address)
func (_Apr *AprCaller) AREP(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "AREP")
	return *ret0, err
}

// AREP is a free data retrieval call binding the contract method 0x97c193c5.
//
// Solidity: function AREP() constant returns(address)
func (_Apr *AprSession) AREP() (common.Address, error) {
	return _Apr.Contract.AREP(&_Apr.CallOpts)
}

// AREP is a free data retrieval call binding the contract method 0x97c193c5.
//
// Solidity: function AREP() constant returns(address)
func (_Apr *AprCallerSession) AREP() (common.Address, error) {
	return _Apr.Contract.AREP(&_Apr.CallOpts)
}

// ASNX is a free data retrieval call binding the contract method 0x61fd7194.
//
// Solidity: function ASNX() constant returns(address)
func (_Apr *AprCaller) ASNX(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "ASNX")
	return *ret0, err
}

// ASNX is a free data retrieval call binding the contract method 0x61fd7194.
//
// Solidity: function ASNX() constant returns(address)
func (_Apr *AprSession) ASNX() (common.Address, error) {
	return _Apr.Contract.ASNX(&_Apr.CallOpts)
}

// ASNX is a free data retrieval call binding the contract method 0x61fd7194.
//
// Solidity: function ASNX() constant returns(address)
func (_Apr *AprCallerSession) ASNX() (common.Address, error) {
	return _Apr.Contract.ASNX(&_Apr.CallOpts)
}

// ASUSD is a free data retrieval call binding the contract method 0x536f2632.
//
// Solidity: function ASUSD() constant returns(address)
func (_Apr *AprCaller) ASUSD(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "ASUSD")
	return *ret0, err
}

// ASUSD is a free data retrieval call binding the contract method 0x536f2632.
//
// Solidity: function ASUSD() constant returns(address)
func (_Apr *AprSession) ASUSD() (common.Address, error) {
	return _Apr.Contract.ASUSD(&_Apr.CallOpts)
}

// ASUSD is a free data retrieval call binding the contract method 0x536f2632.
//
// Solidity: function ASUSD() constant returns(address)
func (_Apr *AprCallerSession) ASUSD() (common.Address, error) {
	return _Apr.Contract.ASUSD(&_Apr.CallOpts)
}

// ATUSD is a free data retrieval call binding the contract method 0x0b58a7a3.
//
// Solidity: function ATUSD() constant returns(address)
func (_Apr *AprCaller) ATUSD(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "ATUSD")
	return *ret0, err
}

// ATUSD is a free data retrieval call binding the contract method 0x0b58a7a3.
//
// Solidity: function ATUSD() constant returns(address)
func (_Apr *AprSession) ATUSD() (common.Address, error) {
	return _Apr.Contract.ATUSD(&_Apr.CallOpts)
}

// ATUSD is a free data retrieval call binding the contract method 0x0b58a7a3.
//
// Solidity: function ATUSD() constant returns(address)
func (_Apr *AprCallerSession) ATUSD() (common.Address, error) {
	return _Apr.Contract.ATUSD(&_Apr.CallOpts)
}

// AUSDC is a free data retrieval call binding the contract method 0x1b1be2bf.
//
// Solidity: function AUSDC() constant returns(address)
func (_Apr *AprCaller) AUSDC(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "AUSDC")
	return *ret0, err
}

// AUSDC is a free data retrieval call binding the contract method 0x1b1be2bf.
//
// Solidity: function AUSDC() constant returns(address)
func (_Apr *AprSession) AUSDC() (common.Address, error) {
	return _Apr.Contract.AUSDC(&_Apr.CallOpts)
}

// AUSDC is a free data retrieval call binding the contract method 0x1b1be2bf.
//
// Solidity: function AUSDC() constant returns(address)
func (_Apr *AprCallerSession) AUSDC() (common.Address, error) {
	return _Apr.Contract.AUSDC(&_Apr.CallOpts)
}

// AUSDT is a free data retrieval call binding the contract method 0x0dc0849a.
//
// Solidity: function AUSDT() constant returns(address)
func (_Apr *AprCaller) AUSDT(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "AUSDT")
	return *ret0, err
}

// AUSDT is a free data retrieval call binding the contract method 0x0dc0849a.
//
// Solidity: function AUSDT() constant returns(address)
func (_Apr *AprSession) AUSDT() (common.Address, error) {
	return _Apr.Contract.AUSDT(&_Apr.CallOpts)
}

// AUSDT is a free data retrieval call binding the contract method 0x0dc0849a.
//
// Solidity: function AUSDT() constant returns(address)
func (_Apr *AprCallerSession) AUSDT() (common.Address, error) {
	return _Apr.Contract.AUSDT(&_Apr.CallOpts)
}

// AWBTC is a free data retrieval call binding the contract method 0x982e2adf.
//
// Solidity: function AWBTC() constant returns(address)
func (_Apr *AprCaller) AWBTC(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "AWBTC")
	return *ret0, err
}

// AWBTC is a free data retrieval call binding the contract method 0x982e2adf.
//
// Solidity: function AWBTC() constant returns(address)
func (_Apr *AprSession) AWBTC() (common.Address, error) {
	return _Apr.Contract.AWBTC(&_Apr.CallOpts)
}

// AWBTC is a free data retrieval call binding the contract method 0x982e2adf.
//
// Solidity: function AWBTC() constant returns(address)
func (_Apr *AprCallerSession) AWBTC() (common.Address, error) {
	return _Apr.Contract.AWBTC(&_Apr.CallOpts)
}

// AZRX is a free data retrieval call binding the contract method 0x49386b5e.
//
// Solidity: function AZRX() constant returns(address)
func (_Apr *AprCaller) AZRX(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "AZRX")
	return *ret0, err
}

// AZRX is a free data retrieval call binding the contract method 0x49386b5e.
//
// Solidity: function AZRX() constant returns(address)
func (_Apr *AprSession) AZRX() (common.Address, error) {
	return _Apr.Contract.AZRX(&_Apr.CallOpts)
}

// AZRX is a free data retrieval call binding the contract method 0x49386b5e.
//
// Solidity: function AZRX() constant returns(address)
func (_Apr *AprCallerSession) AZRX() (common.Address, error) {
	return _Apr.Contract.AZRX(&_Apr.CallOpts)
}

// CBAT is a free data retrieval call binding the contract method 0x72142a2b.
//
// Solidity: function CBAT() constant returns(address)
func (_Apr *AprCaller) CBAT(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "CBAT")
	return *ret0, err
}

// CBAT is a free data retrieval call binding the contract method 0x72142a2b.
//
// Solidity: function CBAT() constant returns(address)
func (_Apr *AprSession) CBAT() (common.Address, error) {
	return _Apr.Contract.CBAT(&_Apr.CallOpts)
}

// CBAT is a free data retrieval call binding the contract method 0x72142a2b.
//
// Solidity: function CBAT() constant returns(address)
func (_Apr *AprCallerSession) CBAT() (common.Address, error) {
	return _Apr.Contract.CBAT(&_Apr.CallOpts)
}

// CDAI is a free data retrieval call binding the contract method 0x878f7603.
//
// Solidity: function CDAI() constant returns(address)
func (_Apr *AprCaller) CDAI(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "CDAI")
	return *ret0, err
}

// CDAI is a free data retrieval call binding the contract method 0x878f7603.
//
// Solidity: function CDAI() constant returns(address)
func (_Apr *AprSession) CDAI() (common.Address, error) {
	return _Apr.Contract.CDAI(&_Apr.CallOpts)
}

// CDAI is a free data retrieval call binding the contract method 0x878f7603.
//
// Solidity: function CDAI() constant returns(address)
func (_Apr *AprCallerSession) CDAI() (common.Address, error) {
	return _Apr.Contract.CDAI(&_Apr.CallOpts)
}

// CETH is a free data retrieval call binding the contract method 0x61096820.
//
// Solidity: function CETH() constant returns(address)
func (_Apr *AprCaller) CETH(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "CETH")
	return *ret0, err
}

// CETH is a free data retrieval call binding the contract method 0x61096820.
//
// Solidity: function CETH() constant returns(address)
func (_Apr *AprSession) CETH() (common.Address, error) {
	return _Apr.Contract.CETH(&_Apr.CallOpts)
}

// CETH is a free data retrieval call binding the contract method 0x61096820.
//
// Solidity: function CETH() constant returns(address)
func (_Apr *AprCallerSession) CETH() (common.Address, error) {
	return _Apr.Contract.CETH(&_Apr.CallOpts)
}

// CREP is a free data retrieval call binding the contract method 0xd9c75e37.
//
// Solidity: function CREP() constant returns(address)
func (_Apr *AprCaller) CREP(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "CREP")
	return *ret0, err
}

// CREP is a free data retrieval call binding the contract method 0xd9c75e37.
//
// Solidity: function CREP() constant returns(address)
func (_Apr *AprSession) CREP() (common.Address, error) {
	return _Apr.Contract.CREP(&_Apr.CallOpts)
}

// CREP is a free data retrieval call binding the contract method 0xd9c75e37.
//
// Solidity: function CREP() constant returns(address)
func (_Apr *AprCallerSession) CREP() (common.Address, error) {
	return _Apr.Contract.CREP(&_Apr.CallOpts)
}

// CSAI is a free data retrieval call binding the contract method 0xeb22a022.
//
// Solidity: function CSAI() constant returns(address)
func (_Apr *AprCaller) CSAI(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "CSAI")
	return *ret0, err
}

// CSAI is a free data retrieval call binding the contract method 0xeb22a022.
//
// Solidity: function CSAI() constant returns(address)
func (_Apr *AprSession) CSAI() (common.Address, error) {
	return _Apr.Contract.CSAI(&_Apr.CallOpts)
}

// CSAI is a free data retrieval call binding the contract method 0xeb22a022.
//
// Solidity: function CSAI() constant returns(address)
func (_Apr *AprCallerSession) CSAI() (common.Address, error) {
	return _Apr.Contract.CSAI(&_Apr.CallOpts)
}

// CUSDC is a free data retrieval call binding the contract method 0xa37b9e7f.
//
// Solidity: function CUSDC() constant returns(address)
func (_Apr *AprCaller) CUSDC(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "CUSDC")
	return *ret0, err
}

// CUSDC is a free data retrieval call binding the contract method 0xa37b9e7f.
//
// Solidity: function CUSDC() constant returns(address)
func (_Apr *AprSession) CUSDC() (common.Address, error) {
	return _Apr.Contract.CUSDC(&_Apr.CallOpts)
}

// CUSDC is a free data retrieval call binding the contract method 0xa37b9e7f.
//
// Solidity: function CUSDC() constant returns(address)
func (_Apr *AprCallerSession) CUSDC() (common.Address, error) {
	return _Apr.Contract.CUSDC(&_Apr.CallOpts)
}

// CWBTC is a free data retrieval call binding the contract method 0x28e8452d.
//
// Solidity: function CWBTC() constant returns(address)
func (_Apr *AprCaller) CWBTC(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "CWBTC")
	return *ret0, err
}

// CWBTC is a free data retrieval call binding the contract method 0x28e8452d.
//
// Solidity: function CWBTC() constant returns(address)
func (_Apr *AprSession) CWBTC() (common.Address, error) {
	return _Apr.Contract.CWBTC(&_Apr.CallOpts)
}

// CWBTC is a free data retrieval call binding the contract method 0x28e8452d.
//
// Solidity: function CWBTC() constant returns(address)
func (_Apr *AprCallerSession) CWBTC() (common.Address, error) {
	return _Apr.Contract.CWBTC(&_Apr.CallOpts)
}

// CZRX is a free data retrieval call binding the contract method 0xf8b5fff0.
//
// Solidity: function CZRX() constant returns(address)
func (_Apr *AprCaller) CZRX(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "CZRX")
	return *ret0, err
}

// CZRX is a free data retrieval call binding the contract method 0xf8b5fff0.
//
// Solidity: function CZRX() constant returns(address)
func (_Apr *AprSession) CZRX() (common.Address, error) {
	return _Apr.Contract.CZRX(&_Apr.CallOpts)
}

// CZRX is a free data retrieval call binding the contract method 0xf8b5fff0.
//
// Solidity: function CZRX() constant returns(address)
func (_Apr *AprCallerSession) CZRX() (common.Address, error) {
	return _Apr.Contract.CZRX(&_Apr.CallOpts)
}

// DYDX is a free data retrieval call binding the contract method 0xc043fca2.
//
// Solidity: function DYDX() constant returns(address)
func (_Apr *AprCaller) DYDX(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "DYDX")
	return *ret0, err
}

// DYDX is a free data retrieval call binding the contract method 0xc043fca2.
//
// Solidity: function DYDX() constant returns(address)
func (_Apr *AprSession) DYDX() (common.Address, error) {
	return _Apr.Contract.DYDX(&_Apr.CallOpts)
}

// DYDX is a free data retrieval call binding the contract method 0xc043fca2.
//
// Solidity: function DYDX() constant returns(address)
func (_Apr *AprCallerSession) DYDX() (common.Address, error) {
	return _Apr.Contract.DYDX(&_Apr.CallOpts)
}

// IBAT is a free data retrieval call binding the contract method 0x1cc16a6a.
//
// Solidity: function IBAT() constant returns(address)
func (_Apr *AprCaller) IBAT(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "IBAT")
	return *ret0, err
}

// IBAT is a free data retrieval call binding the contract method 0x1cc16a6a.
//
// Solidity: function IBAT() constant returns(address)
func (_Apr *AprSession) IBAT() (common.Address, error) {
	return _Apr.Contract.IBAT(&_Apr.CallOpts)
}

// IBAT is a free data retrieval call binding the contract method 0x1cc16a6a.
//
// Solidity: function IBAT() constant returns(address)
func (_Apr *AprCallerSession) IBAT() (common.Address, error) {
	return _Apr.Contract.IBAT(&_Apr.CallOpts)
}

// IDAI is a free data retrieval call binding the contract method 0xdada70dc.
//
// Solidity: function IDAI() constant returns(address)
func (_Apr *AprCaller) IDAI(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "IDAI")
	return *ret0, err
}

// IDAI is a free data retrieval call binding the contract method 0xdada70dc.
//
// Solidity: function IDAI() constant returns(address)
func (_Apr *AprSession) IDAI() (common.Address, error) {
	return _Apr.Contract.IDAI(&_Apr.CallOpts)
}

// IDAI is a free data retrieval call binding the contract method 0xdada70dc.
//
// Solidity: function IDAI() constant returns(address)
func (_Apr *AprCallerSession) IDAI() (common.Address, error) {
	return _Apr.Contract.IDAI(&_Apr.CallOpts)
}

// IETH is a free data retrieval call binding the contract method 0x9be0587e.
//
// Solidity: function IETH() constant returns(address)
func (_Apr *AprCaller) IETH(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "IETH")
	return *ret0, err
}

// IETH is a free data retrieval call binding the contract method 0x9be0587e.
//
// Solidity: function IETH() constant returns(address)
func (_Apr *AprSession) IETH() (common.Address, error) {
	return _Apr.Contract.IETH(&_Apr.CallOpts)
}

// IETH is a free data retrieval call binding the contract method 0x9be0587e.
//
// Solidity: function IETH() constant returns(address)
func (_Apr *AprCallerSession) IETH() (common.Address, error) {
	return _Apr.Contract.IETH(&_Apr.CallOpts)
}

// IKNC is a free data retrieval call binding the contract method 0x63a68940.
//
// Solidity: function IKNC() constant returns(address)
func (_Apr *AprCaller) IKNC(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "IKNC")
	return *ret0, err
}

// IKNC is a free data retrieval call binding the contract method 0x63a68940.
//
// Solidity: function IKNC() constant returns(address)
func (_Apr *AprSession) IKNC() (common.Address, error) {
	return _Apr.Contract.IKNC(&_Apr.CallOpts)
}

// IKNC is a free data retrieval call binding the contract method 0x63a68940.
//
// Solidity: function IKNC() constant returns(address)
func (_Apr *AprCallerSession) IKNC() (common.Address, error) {
	return _Apr.Contract.IKNC(&_Apr.CallOpts)
}

// ILINK is a free data retrieval call binding the contract method 0x18b3b509.
//
// Solidity: function ILINK() constant returns(address)
func (_Apr *AprCaller) ILINK(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "ILINK")
	return *ret0, err
}

// ILINK is a free data retrieval call binding the contract method 0x18b3b509.
//
// Solidity: function ILINK() constant returns(address)
func (_Apr *AprSession) ILINK() (common.Address, error) {
	return _Apr.Contract.ILINK(&_Apr.CallOpts)
}

// ILINK is a free data retrieval call binding the contract method 0x18b3b509.
//
// Solidity: function ILINK() constant returns(address)
func (_Apr *AprCallerSession) ILINK() (common.Address, error) {
	return _Apr.Contract.ILINK(&_Apr.CallOpts)
}

// IREP is a free data retrieval call binding the contract method 0x90b2ee81.
//
// Solidity: function IREP() constant returns(address)
func (_Apr *AprCaller) IREP(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "IREP")
	return *ret0, err
}

// IREP is a free data retrieval call binding the contract method 0x90b2ee81.
//
// Solidity: function IREP() constant returns(address)
func (_Apr *AprSession) IREP() (common.Address, error) {
	return _Apr.Contract.IREP(&_Apr.CallOpts)
}

// IREP is a free data retrieval call binding the contract method 0x90b2ee81.
//
// Solidity: function IREP() constant returns(address)
func (_Apr *AprCallerSession) IREP() (common.Address, error) {
	return _Apr.Contract.IREP(&_Apr.CallOpts)
}

// ISAI is a free data retrieval call binding the contract method 0xcda08710.
//
// Solidity: function ISAI() constant returns(address)
func (_Apr *AprCaller) ISAI(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "ISAI")
	return *ret0, err
}

// ISAI is a free data retrieval call binding the contract method 0xcda08710.
//
// Solidity: function ISAI() constant returns(address)
func (_Apr *AprSession) ISAI() (common.Address, error) {
	return _Apr.Contract.ISAI(&_Apr.CallOpts)
}

// ISAI is a free data retrieval call binding the contract method 0xcda08710.
//
// Solidity: function ISAI() constant returns(address)
func (_Apr *AprCallerSession) ISAI() (common.Address, error) {
	return _Apr.Contract.ISAI(&_Apr.CallOpts)
}

// ISUSD is a free data retrieval call binding the contract method 0x43d39c8b.
//
// Solidity: function ISUSD() constant returns(address)
func (_Apr *AprCaller) ISUSD(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "ISUSD")
	return *ret0, err
}

// ISUSD is a free data retrieval call binding the contract method 0x43d39c8b.
//
// Solidity: function ISUSD() constant returns(address)
func (_Apr *AprSession) ISUSD() (common.Address, error) {
	return _Apr.Contract.ISUSD(&_Apr.CallOpts)
}

// ISUSD is a free data retrieval call binding the contract method 0x43d39c8b.
//
// Solidity: function ISUSD() constant returns(address)
func (_Apr *AprCallerSession) ISUSD() (common.Address, error) {
	return _Apr.Contract.ISUSD(&_Apr.CallOpts)
}

// IUSDC is a free data retrieval call binding the contract method 0x3fa6a66b.
//
// Solidity: function IUSDC() constant returns(address)
func (_Apr *AprCaller) IUSDC(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "IUSDC")
	return *ret0, err
}

// IUSDC is a free data retrieval call binding the contract method 0x3fa6a66b.
//
// Solidity: function IUSDC() constant returns(address)
func (_Apr *AprSession) IUSDC() (common.Address, error) {
	return _Apr.Contract.IUSDC(&_Apr.CallOpts)
}

// IUSDC is a free data retrieval call binding the contract method 0x3fa6a66b.
//
// Solidity: function IUSDC() constant returns(address)
func (_Apr *AprCallerSession) IUSDC() (common.Address, error) {
	return _Apr.Contract.IUSDC(&_Apr.CallOpts)
}

// IWBTC is a free data retrieval call binding the contract method 0xc30c7cae.
//
// Solidity: function IWBTC() constant returns(address)
func (_Apr *AprCaller) IWBTC(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "IWBTC")
	return *ret0, err
}

// IWBTC is a free data retrieval call binding the contract method 0xc30c7cae.
//
// Solidity: function IWBTC() constant returns(address)
func (_Apr *AprSession) IWBTC() (common.Address, error) {
	return _Apr.Contract.IWBTC(&_Apr.CallOpts)
}

// IWBTC is a free data retrieval call binding the contract method 0xc30c7cae.
//
// Solidity: function IWBTC() constant returns(address)
func (_Apr *AprCallerSession) IWBTC() (common.Address, error) {
	return _Apr.Contract.IWBTC(&_Apr.CallOpts)
}

// IZRX is a free data retrieval call binding the contract method 0xa22db74c.
//
// Solidity: function IZRX() constant returns(address)
func (_Apr *AprCaller) IZRX(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "IZRX")
	return *ret0, err
}

// IZRX is a free data retrieval call binding the contract method 0xa22db74c.
//
// Solidity: function IZRX() constant returns(address)
func (_Apr *AprSession) IZRX() (common.Address, error) {
	return _Apr.Contract.IZRX(&_Apr.CallOpts)
}

// IZRX is a free data retrieval call binding the contract method 0xa22db74c.
//
// Solidity: function IZRX() constant returns(address)
func (_Apr *AprCallerSession) IZRX() (common.Address, error) {
	return _Apr.Contract.IZRX(&_Apr.CallOpts)
}

// GetABATAPR is a free data retrieval call binding the contract method 0x5ce7174f.
//
// Solidity: function getABATAPR() constant returns(uint256)
func (_Apr *AprCaller) GetABATAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getABATAPR")
	return *ret0, err
}

// GetABATAPR is a free data retrieval call binding the contract method 0x5ce7174f.
//
// Solidity: function getABATAPR() constant returns(uint256)
func (_Apr *AprSession) GetABATAPR() (*big.Int, error) {
	return _Apr.Contract.GetABATAPR(&_Apr.CallOpts)
}

// GetABATAPR is a free data retrieval call binding the contract method 0x5ce7174f.
//
// Solidity: function getABATAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetABATAPR() (*big.Int, error) {
	return _Apr.Contract.GetABATAPR(&_Apr.CallOpts)
}

// GetADAIAPR is a free data retrieval call binding the contract method 0x243d6659.
//
// Solidity: function getADAIAPR() constant returns(uint256)
func (_Apr *AprCaller) GetADAIAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getADAIAPR")
	return *ret0, err
}

// GetADAIAPR is a free data retrieval call binding the contract method 0x243d6659.
//
// Solidity: function getADAIAPR() constant returns(uint256)
func (_Apr *AprSession) GetADAIAPR() (*big.Int, error) {
	return _Apr.Contract.GetADAIAPR(&_Apr.CallOpts)
}

// GetADAIAPR is a free data retrieval call binding the contract method 0x243d6659.
//
// Solidity: function getADAIAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetADAIAPR() (*big.Int, error) {
	return _Apr.Contract.GetADAIAPR(&_Apr.CallOpts)
}

// GetAETHAPR is a free data retrieval call binding the contract method 0xe8d79148.
//
// Solidity: function getAETHAPR() constant returns(uint256)
func (_Apr *AprCaller) GetAETHAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getAETHAPR")
	return *ret0, err
}

// GetAETHAPR is a free data retrieval call binding the contract method 0xe8d79148.
//
// Solidity: function getAETHAPR() constant returns(uint256)
func (_Apr *AprSession) GetAETHAPR() (*big.Int, error) {
	return _Apr.Contract.GetAETHAPR(&_Apr.CallOpts)
}

// GetAETHAPR is a free data retrieval call binding the contract method 0xe8d79148.
//
// Solidity: function getAETHAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetAETHAPR() (*big.Int, error) {
	return _Apr.Contract.GetAETHAPR(&_Apr.CallOpts)
}

// GetAKNCAPR is a free data retrieval call binding the contract method 0x4907c7ae.
//
// Solidity: function getAKNCAPR() constant returns(uint256)
func (_Apr *AprCaller) GetAKNCAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getAKNCAPR")
	return *ret0, err
}

// GetAKNCAPR is a free data retrieval call binding the contract method 0x4907c7ae.
//
// Solidity: function getAKNCAPR() constant returns(uint256)
func (_Apr *AprSession) GetAKNCAPR() (*big.Int, error) {
	return _Apr.Contract.GetAKNCAPR(&_Apr.CallOpts)
}

// GetAKNCAPR is a free data retrieval call binding the contract method 0x4907c7ae.
//
// Solidity: function getAKNCAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetAKNCAPR() (*big.Int, error) {
	return _Apr.Contract.GetAKNCAPR(&_Apr.CallOpts)
}

// GetALENDAPR is a free data retrieval call binding the contract method 0x832b4b03.
//
// Solidity: function getALENDAPR() constant returns(uint256)
func (_Apr *AprCaller) GetALENDAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getALENDAPR")
	return *ret0, err
}

// GetALENDAPR is a free data retrieval call binding the contract method 0x832b4b03.
//
// Solidity: function getALENDAPR() constant returns(uint256)
func (_Apr *AprSession) GetALENDAPR() (*big.Int, error) {
	return _Apr.Contract.GetALENDAPR(&_Apr.CallOpts)
}

// GetALENDAPR is a free data retrieval call binding the contract method 0x832b4b03.
//
// Solidity: function getALENDAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetALENDAPR() (*big.Int, error) {
	return _Apr.Contract.GetALENDAPR(&_Apr.CallOpts)
}

// GetALINKAPR is a free data retrieval call binding the contract method 0x3151a245.
//
// Solidity: function getALINKAPR() constant returns(uint256)
func (_Apr *AprCaller) GetALINKAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getALINKAPR")
	return *ret0, err
}

// GetALINKAPR is a free data retrieval call binding the contract method 0x3151a245.
//
// Solidity: function getALINKAPR() constant returns(uint256)
func (_Apr *AprSession) GetALINKAPR() (*big.Int, error) {
	return _Apr.Contract.GetALINKAPR(&_Apr.CallOpts)
}

// GetALINKAPR is a free data retrieval call binding the contract method 0x3151a245.
//
// Solidity: function getALINKAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetALINKAPR() (*big.Int, error) {
	return _Apr.Contract.GetALINKAPR(&_Apr.CallOpts)
}

// GetAMANAAPR is a free data retrieval call binding the contract method 0xb4650bf7.
//
// Solidity: function getAMANAAPR() constant returns(uint256)
func (_Apr *AprCaller) GetAMANAAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getAMANAAPR")
	return *ret0, err
}

// GetAMANAAPR is a free data retrieval call binding the contract method 0xb4650bf7.
//
// Solidity: function getAMANAAPR() constant returns(uint256)
func (_Apr *AprSession) GetAMANAAPR() (*big.Int, error) {
	return _Apr.Contract.GetAMANAAPR(&_Apr.CallOpts)
}

// GetAMANAAPR is a free data retrieval call binding the contract method 0xb4650bf7.
//
// Solidity: function getAMANAAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetAMANAAPR() (*big.Int, error) {
	return _Apr.Contract.GetAMANAAPR(&_Apr.CallOpts)
}

// GetAMKRAPR is a free data retrieval call binding the contract method 0xbe2134eb.
//
// Solidity: function getAMKRAPR() constant returns(uint256)
func (_Apr *AprCaller) GetAMKRAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getAMKRAPR")
	return *ret0, err
}

// GetAMKRAPR is a free data retrieval call binding the contract method 0xbe2134eb.
//
// Solidity: function getAMKRAPR() constant returns(uint256)
func (_Apr *AprSession) GetAMKRAPR() (*big.Int, error) {
	return _Apr.Contract.GetAMKRAPR(&_Apr.CallOpts)
}

// GetAMKRAPR is a free data retrieval call binding the contract method 0xbe2134eb.
//
// Solidity: function getAMKRAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetAMKRAPR() (*big.Int, error) {
	return _Apr.Contract.GetAMKRAPR(&_Apr.CallOpts)
}

// GetAREPAPR is a free data retrieval call binding the contract method 0x693d42e7.
//
// Solidity: function getAREPAPR() constant returns(uint256)
func (_Apr *AprCaller) GetAREPAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getAREPAPR")
	return *ret0, err
}

// GetAREPAPR is a free data retrieval call binding the contract method 0x693d42e7.
//
// Solidity: function getAREPAPR() constant returns(uint256)
func (_Apr *AprSession) GetAREPAPR() (*big.Int, error) {
	return _Apr.Contract.GetAREPAPR(&_Apr.CallOpts)
}

// GetAREPAPR is a free data retrieval call binding the contract method 0x693d42e7.
//
// Solidity: function getAREPAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetAREPAPR() (*big.Int, error) {
	return _Apr.Contract.GetAREPAPR(&_Apr.CallOpts)
}

// GetASNXAPR is a free data retrieval call binding the contract method 0x4e9a1016.
//
// Solidity: function getASNXAPR() constant returns(uint256)
func (_Apr *AprCaller) GetASNXAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getASNXAPR")
	return *ret0, err
}

// GetASNXAPR is a free data retrieval call binding the contract method 0x4e9a1016.
//
// Solidity: function getASNXAPR() constant returns(uint256)
func (_Apr *AprSession) GetASNXAPR() (*big.Int, error) {
	return _Apr.Contract.GetASNXAPR(&_Apr.CallOpts)
}

// GetASNXAPR is a free data retrieval call binding the contract method 0x4e9a1016.
//
// Solidity: function getASNXAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetASNXAPR() (*big.Int, error) {
	return _Apr.Contract.GetASNXAPR(&_Apr.CallOpts)
}

// GetASUSDAPR is a free data retrieval call binding the contract method 0x5ac1baa9.
//
// Solidity: function getASUSDAPR() constant returns(uint256)
func (_Apr *AprCaller) GetASUSDAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getASUSDAPR")
	return *ret0, err
}

// GetASUSDAPR is a free data retrieval call binding the contract method 0x5ac1baa9.
//
// Solidity: function getASUSDAPR() constant returns(uint256)
func (_Apr *AprSession) GetASUSDAPR() (*big.Int, error) {
	return _Apr.Contract.GetASUSDAPR(&_Apr.CallOpts)
}

// GetASUSDAPR is a free data retrieval call binding the contract method 0x5ac1baa9.
//
// Solidity: function getASUSDAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetASUSDAPR() (*big.Int, error) {
	return _Apr.Contract.GetASUSDAPR(&_Apr.CallOpts)
}

// GetATUSDAPR is a free data retrieval call binding the contract method 0x35b37442.
//
// Solidity: function getATUSDAPR() constant returns(uint256)
func (_Apr *AprCaller) GetATUSDAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getATUSDAPR")
	return *ret0, err
}

// GetATUSDAPR is a free data retrieval call binding the contract method 0x35b37442.
//
// Solidity: function getATUSDAPR() constant returns(uint256)
func (_Apr *AprSession) GetATUSDAPR() (*big.Int, error) {
	return _Apr.Contract.GetATUSDAPR(&_Apr.CallOpts)
}

// GetATUSDAPR is a free data retrieval call binding the contract method 0x35b37442.
//
// Solidity: function getATUSDAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetATUSDAPR() (*big.Int, error) {
	return _Apr.Contract.GetATUSDAPR(&_Apr.CallOpts)
}

// GetAUSDCAPR is a free data retrieval call binding the contract method 0x3196ebbd.
//
// Solidity: function getAUSDCAPR() constant returns(uint256)
func (_Apr *AprCaller) GetAUSDCAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getAUSDCAPR")
	return *ret0, err
}

// GetAUSDCAPR is a free data retrieval call binding the contract method 0x3196ebbd.
//
// Solidity: function getAUSDCAPR() constant returns(uint256)
func (_Apr *AprSession) GetAUSDCAPR() (*big.Int, error) {
	return _Apr.Contract.GetAUSDCAPR(&_Apr.CallOpts)
}

// GetAUSDCAPR is a free data retrieval call binding the contract method 0x3196ebbd.
//
// Solidity: function getAUSDCAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetAUSDCAPR() (*big.Int, error) {
	return _Apr.Contract.GetAUSDCAPR(&_Apr.CallOpts)
}

// GetAUSDTAPR is a free data retrieval call binding the contract method 0x02c93c70.
//
// Solidity: function getAUSDTAPR() constant returns(uint256)
func (_Apr *AprCaller) GetAUSDTAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getAUSDTAPR")
	return *ret0, err
}

// GetAUSDTAPR is a free data retrieval call binding the contract method 0x02c93c70.
//
// Solidity: function getAUSDTAPR() constant returns(uint256)
func (_Apr *AprSession) GetAUSDTAPR() (*big.Int, error) {
	return _Apr.Contract.GetAUSDTAPR(&_Apr.CallOpts)
}

// GetAUSDTAPR is a free data retrieval call binding the contract method 0x02c93c70.
//
// Solidity: function getAUSDTAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetAUSDTAPR() (*big.Int, error) {
	return _Apr.Contract.GetAUSDTAPR(&_Apr.CallOpts)
}

// GetAWBTCAPR is a free data retrieval call binding the contract method 0xc6882b74.
//
// Solidity: function getAWBTCAPR() constant returns(uint256)
func (_Apr *AprCaller) GetAWBTCAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getAWBTCAPR")
	return *ret0, err
}

// GetAWBTCAPR is a free data retrieval call binding the contract method 0xc6882b74.
//
// Solidity: function getAWBTCAPR() constant returns(uint256)
func (_Apr *AprSession) GetAWBTCAPR() (*big.Int, error) {
	return _Apr.Contract.GetAWBTCAPR(&_Apr.CallOpts)
}

// GetAWBTCAPR is a free data retrieval call binding the contract method 0xc6882b74.
//
// Solidity: function getAWBTCAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetAWBTCAPR() (*big.Int, error) {
	return _Apr.Contract.GetAWBTCAPR(&_Apr.CallOpts)
}

// GetAZRXAPR is a free data retrieval call binding the contract method 0xa0bc8d53.
//
// Solidity: function getAZRXAPR() constant returns(uint256)
func (_Apr *AprCaller) GetAZRXAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getAZRXAPR")
	return *ret0, err
}

// GetAZRXAPR is a free data retrieval call binding the contract method 0xa0bc8d53.
//
// Solidity: function getAZRXAPR() constant returns(uint256)
func (_Apr *AprSession) GetAZRXAPR() (*big.Int, error) {
	return _Apr.Contract.GetAZRXAPR(&_Apr.CallOpts)
}

// GetAZRXAPR is a free data retrieval call binding the contract method 0xa0bc8d53.
//
// Solidity: function getAZRXAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetAZRXAPR() (*big.Int, error) {
	return _Apr.Contract.GetAZRXAPR(&_Apr.CallOpts)
}

// GetAaveAPR is a free data retrieval call binding the contract method 0x192a9861.
//
// Solidity: function getAaveAPR(address token) constant returns(uint256)
func (_Apr *AprCaller) GetAaveAPR(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getAaveAPR", token)
	return *ret0, err
}

// GetAaveAPR is a free data retrieval call binding the contract method 0x192a9861.
//
// Solidity: function getAaveAPR(address token) constant returns(uint256)
func (_Apr *AprSession) GetAaveAPR(token common.Address) (*big.Int, error) {
	return _Apr.Contract.GetAaveAPR(&_Apr.CallOpts, token)
}

// GetAaveAPR is a free data retrieval call binding the contract method 0x192a9861.
//
// Solidity: function getAaveAPR(address token) constant returns(uint256)
func (_Apr *AprCallerSession) GetAaveAPR(token common.Address) (*big.Int, error) {
	return _Apr.Contract.GetAaveAPR(&_Apr.CallOpts, token)
}

// GetAaveCore is a free data retrieval call binding the contract method 0xf82ce27d.
//
// Solidity: function getAaveCore() constant returns(address)
func (_Apr *AprCaller) GetAaveCore(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getAaveCore")
	return *ret0, err
}

// GetAaveCore is a free data retrieval call binding the contract method 0xf82ce27d.
//
// Solidity: function getAaveCore() constant returns(address)
func (_Apr *AprSession) GetAaveCore() (common.Address, error) {
	return _Apr.Contract.GetAaveCore(&_Apr.CallOpts)
}

// GetAaveCore is a free data retrieval call binding the contract method 0xf82ce27d.
//
// Solidity: function getAaveCore() constant returns(address)
func (_Apr *AprCallerSession) GetAaveCore() (common.Address, error) {
	return _Apr.Contract.GetAaveCore(&_Apr.CallOpts)
}

// GetAllAaveAPR is a free data retrieval call binding the contract method 0x983a97d9.
//
// Solidity: function getAllAaveAPR() constant returns(uint256 aDAI, uint256 aTUSD, uint256 aUSDC, uint256 aUSDT, uint256 aSUSD, uint256 aBAT, uint256 aETH, uint256 aLINK, uint256 aKNC, uint256 aREP, uint256 aZRX, uint256 aSNX)
func (_Apr *AprCaller) GetAllAaveAPR(opts *bind.CallOpts) (struct {
	ADAI  *big.Int
	ATUSD *big.Int
	AUSDC *big.Int
	AUSDT *big.Int
	ASUSD *big.Int
	ABAT  *big.Int
	AETH  *big.Int
	ALINK *big.Int
	AKNC  *big.Int
	AREP  *big.Int
	AZRX  *big.Int
	ASNX  *big.Int
}, error) {
	ret := new(struct {
		ADAI  *big.Int
		ATUSD *big.Int
		AUSDC *big.Int
		AUSDT *big.Int
		ASUSD *big.Int
		ABAT  *big.Int
		AETH  *big.Int
		ALINK *big.Int
		AKNC  *big.Int
		AREP  *big.Int
		AZRX  *big.Int
		ASNX  *big.Int
	})
	out := ret
	err := _Apr.contract.Call(opts, out, "getAllAaveAPR")
	return *ret, err
}

// GetAllAaveAPR is a free data retrieval call binding the contract method 0x983a97d9.
//
// Solidity: function getAllAaveAPR() constant returns(uint256 aDAI, uint256 aTUSD, uint256 aUSDC, uint256 aUSDT, uint256 aSUSD, uint256 aBAT, uint256 aETH, uint256 aLINK, uint256 aKNC, uint256 aREP, uint256 aZRX, uint256 aSNX)
func (_Apr *AprSession) GetAllAaveAPR() (struct {
	ADAI  *big.Int
	ATUSD *big.Int
	AUSDC *big.Int
	AUSDT *big.Int
	ASUSD *big.Int
	ABAT  *big.Int
	AETH  *big.Int
	ALINK *big.Int
	AKNC  *big.Int
	AREP  *big.Int
	AZRX  *big.Int
	ASNX  *big.Int
}, error) {
	return _Apr.Contract.GetAllAaveAPR(&_Apr.CallOpts)
}

// GetAllAaveAPR is a free data retrieval call binding the contract method 0x983a97d9.
//
// Solidity: function getAllAaveAPR() constant returns(uint256 aDAI, uint256 aTUSD, uint256 aUSDC, uint256 aUSDT, uint256 aSUSD, uint256 aBAT, uint256 aETH, uint256 aLINK, uint256 aKNC, uint256 aREP, uint256 aZRX, uint256 aSNX)
func (_Apr *AprCallerSession) GetAllAaveAPR() (struct {
	ADAI  *big.Int
	ATUSD *big.Int
	AUSDC *big.Int
	AUSDT *big.Int
	ASUSD *big.Int
	ABAT  *big.Int
	AETH  *big.Int
	ALINK *big.Int
	AKNC  *big.Int
	AREP  *big.Int
	AZRX  *big.Int
	ASNX  *big.Int
}, error) {
	return _Apr.Contract.GetAllAaveAPR(&_Apr.CallOpts)
}

// GetAllCompoundAPR is a free data retrieval call binding the contract method 0x454ccd71.
//
// Solidity: function getAllCompoundAPR() constant returns(uint256 cDAI, uint256 cBAT, uint256 cETH, uint256 cREP, uint256 cSAI, uint256 cUSDC, uint256 cWBTC, uint256 cZRC)
func (_Apr *AprCaller) GetAllCompoundAPR(opts *bind.CallOpts) (struct {
	CDAI  *big.Int
	CBAT  *big.Int
	CETH  *big.Int
	CREP  *big.Int
	CSAI  *big.Int
	CUSDC *big.Int
	CWBTC *big.Int
	CZRC  *big.Int
}, error) {
	ret := new(struct {
		CDAI  *big.Int
		CBAT  *big.Int
		CETH  *big.Int
		CREP  *big.Int
		CSAI  *big.Int
		CUSDC *big.Int
		CWBTC *big.Int
		CZRC  *big.Int
	})
	out := ret
	err := _Apr.contract.Call(opts, out, "getAllCompoundAPR")
	return *ret, err
}

// GetAllCompoundAPR is a free data retrieval call binding the contract method 0x454ccd71.
//
// Solidity: function getAllCompoundAPR() constant returns(uint256 cDAI, uint256 cBAT, uint256 cETH, uint256 cREP, uint256 cSAI, uint256 cUSDC, uint256 cWBTC, uint256 cZRC)
func (_Apr *AprSession) GetAllCompoundAPR() (struct {
	CDAI  *big.Int
	CBAT  *big.Int
	CETH  *big.Int
	CREP  *big.Int
	CSAI  *big.Int
	CUSDC *big.Int
	CWBTC *big.Int
	CZRC  *big.Int
}, error) {
	return _Apr.Contract.GetAllCompoundAPR(&_Apr.CallOpts)
}

// GetAllCompoundAPR is a free data retrieval call binding the contract method 0x454ccd71.
//
// Solidity: function getAllCompoundAPR() constant returns(uint256 cDAI, uint256 cBAT, uint256 cETH, uint256 cREP, uint256 cSAI, uint256 cUSDC, uint256 cWBTC, uint256 cZRC)
func (_Apr *AprCallerSession) GetAllCompoundAPR() (struct {
	CDAI  *big.Int
	CBAT  *big.Int
	CETH  *big.Int
	CREP  *big.Int
	CSAI  *big.Int
	CUSDC *big.Int
	CWBTC *big.Int
	CZRC  *big.Int
}, error) {
	return _Apr.Contract.GetAllCompoundAPR(&_Apr.CallOpts)
}

// GetAllDyDxAPR is a free data retrieval call binding the contract method 0x7cde729f.
//
// Solidity: function getAllDyDxAPR() constant returns(uint256 dSAI, uint256 dETH, uint256 dUSDC, uint256 dDAI)
func (_Apr *AprCaller) GetAllDyDxAPR(opts *bind.CallOpts) (struct {
	DSAI  *big.Int
	DETH  *big.Int
	DUSDC *big.Int
	DDAI  *big.Int
}, error) {
	ret := new(struct {
		DSAI  *big.Int
		DETH  *big.Int
		DUSDC *big.Int
		DDAI  *big.Int
	})
	out := ret
	err := _Apr.contract.Call(opts, out, "getAllDyDxAPR")
	return *ret, err
}

// GetAllDyDxAPR is a free data retrieval call binding the contract method 0x7cde729f.
//
// Solidity: function getAllDyDxAPR() constant returns(uint256 dSAI, uint256 dETH, uint256 dUSDC, uint256 dDAI)
func (_Apr *AprSession) GetAllDyDxAPR() (struct {
	DSAI  *big.Int
	DETH  *big.Int
	DUSDC *big.Int
	DDAI  *big.Int
}, error) {
	return _Apr.Contract.GetAllDyDxAPR(&_Apr.CallOpts)
}

// GetAllDyDxAPR is a free data retrieval call binding the contract method 0x7cde729f.
//
// Solidity: function getAllDyDxAPR() constant returns(uint256 dSAI, uint256 dETH, uint256 dUSDC, uint256 dDAI)
func (_Apr *AprCallerSession) GetAllDyDxAPR() (struct {
	DSAI  *big.Int
	DETH  *big.Int
	DUSDC *big.Int
	DDAI  *big.Int
}, error) {
	return _Apr.Contract.GetAllDyDxAPR(&_Apr.CallOpts)
}

// GetAllFulcrumAPR is a free data retrieval call binding the contract method 0x8264d109.
//
// Solidity: function getAllFulcrumAPR() constant returns(uint256 iZRX, uint256 iREP, uint256 iKNC, uint256 iWBTC, uint256 iUSDC, uint256 iETH, uint256 iSAI, uint256 iDAI, uint256 iLINK, uint256 iSUSD)
func (_Apr *AprCaller) GetAllFulcrumAPR(opts *bind.CallOpts) (struct {
	IZRX  *big.Int
	IREP  *big.Int
	IKNC  *big.Int
	IWBTC *big.Int
	IUSDC *big.Int
	IETH  *big.Int
	ISAI  *big.Int
	IDAI  *big.Int
	ILINK *big.Int
	ISUSD *big.Int
}, error) {
	ret := new(struct {
		IZRX  *big.Int
		IREP  *big.Int
		IKNC  *big.Int
		IWBTC *big.Int
		IUSDC *big.Int
		IETH  *big.Int
		ISAI  *big.Int
		IDAI  *big.Int
		ILINK *big.Int
		ISUSD *big.Int
	})
	out := ret
	err := _Apr.contract.Call(opts, out, "getAllFulcrumAPR")
	return *ret, err
}

// GetAllFulcrumAPR is a free data retrieval call binding the contract method 0x8264d109.
//
// Solidity: function getAllFulcrumAPR() constant returns(uint256 iZRX, uint256 iREP, uint256 iKNC, uint256 iWBTC, uint256 iUSDC, uint256 iETH, uint256 iSAI, uint256 iDAI, uint256 iLINK, uint256 iSUSD)
func (_Apr *AprSession) GetAllFulcrumAPR() (struct {
	IZRX  *big.Int
	IREP  *big.Int
	IKNC  *big.Int
	IWBTC *big.Int
	IUSDC *big.Int
	IETH  *big.Int
	ISAI  *big.Int
	IDAI  *big.Int
	ILINK *big.Int
	ISUSD *big.Int
}, error) {
	return _Apr.Contract.GetAllFulcrumAPR(&_Apr.CallOpts)
}

// GetAllFulcrumAPR is a free data retrieval call binding the contract method 0x8264d109.
//
// Solidity: function getAllFulcrumAPR() constant returns(uint256 iZRX, uint256 iREP, uint256 iKNC, uint256 iWBTC, uint256 iUSDC, uint256 iETH, uint256 iSAI, uint256 iDAI, uint256 iLINK, uint256 iSUSD)
func (_Apr *AprCallerSession) GetAllFulcrumAPR() (struct {
	IZRX  *big.Int
	IREP  *big.Int
	IKNC  *big.Int
	IWBTC *big.Int
	IUSDC *big.Int
	IETH  *big.Int
	ISAI  *big.Int
	IDAI  *big.Int
	ILINK *big.Int
	ISUSD *big.Int
}, error) {
	return _Apr.Contract.GetAllFulcrumAPR(&_Apr.CallOpts)
}

// GetCBATAPR is a free data retrieval call binding the contract method 0x10b4acdc.
//
// Solidity: function getCBATAPR() constant returns(uint256)
func (_Apr *AprCaller) GetCBATAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getCBATAPR")
	return *ret0, err
}

// GetCBATAPR is a free data retrieval call binding the contract method 0x10b4acdc.
//
// Solidity: function getCBATAPR() constant returns(uint256)
func (_Apr *AprSession) GetCBATAPR() (*big.Int, error) {
	return _Apr.Contract.GetCBATAPR(&_Apr.CallOpts)
}

// GetCBATAPR is a free data retrieval call binding the contract method 0x10b4acdc.
//
// Solidity: function getCBATAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetCBATAPR() (*big.Int, error) {
	return _Apr.Contract.GetCBATAPR(&_Apr.CallOpts)
}

// GetCDAIAPR is a free data retrieval call binding the contract method 0x4d2cd8ab.
//
// Solidity: function getCDAIAPR() constant returns(uint256)
func (_Apr *AprCaller) GetCDAIAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getCDAIAPR")
	return *ret0, err
}

// GetCDAIAPR is a free data retrieval call binding the contract method 0x4d2cd8ab.
//
// Solidity: function getCDAIAPR() constant returns(uint256)
func (_Apr *AprSession) GetCDAIAPR() (*big.Int, error) {
	return _Apr.Contract.GetCDAIAPR(&_Apr.CallOpts)
}

// GetCDAIAPR is a free data retrieval call binding the contract method 0x4d2cd8ab.
//
// Solidity: function getCDAIAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetCDAIAPR() (*big.Int, error) {
	return _Apr.Contract.GetCDAIAPR(&_Apr.CallOpts)
}

// GetCETHAPR is a free data retrieval call binding the contract method 0xaee7b9f5.
//
// Solidity: function getCETHAPR() constant returns(uint256)
func (_Apr *AprCaller) GetCETHAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getCETHAPR")
	return *ret0, err
}

// GetCETHAPR is a free data retrieval call binding the contract method 0xaee7b9f5.
//
// Solidity: function getCETHAPR() constant returns(uint256)
func (_Apr *AprSession) GetCETHAPR() (*big.Int, error) {
	return _Apr.Contract.GetCETHAPR(&_Apr.CallOpts)
}

// GetCETHAPR is a free data retrieval call binding the contract method 0xaee7b9f5.
//
// Solidity: function getCETHAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetCETHAPR() (*big.Int, error) {
	return _Apr.Contract.GetCETHAPR(&_Apr.CallOpts)
}

// GetCREPAPR is a free data retrieval call binding the contract method 0x18002983.
//
// Solidity: function getCREPAPR() constant returns(uint256)
func (_Apr *AprCaller) GetCREPAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getCREPAPR")
	return *ret0, err
}

// GetCREPAPR is a free data retrieval call binding the contract method 0x18002983.
//
// Solidity: function getCREPAPR() constant returns(uint256)
func (_Apr *AprSession) GetCREPAPR() (*big.Int, error) {
	return _Apr.Contract.GetCREPAPR(&_Apr.CallOpts)
}

// GetCREPAPR is a free data retrieval call binding the contract method 0x18002983.
//
// Solidity: function getCREPAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetCREPAPR() (*big.Int, error) {
	return _Apr.Contract.GetCREPAPR(&_Apr.CallOpts)
}

// GetCSAIAPR is a free data retrieval call binding the contract method 0x638aa95b.
//
// Solidity: function getCSAIAPR() constant returns(uint256)
func (_Apr *AprCaller) GetCSAIAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getCSAIAPR")
	return *ret0, err
}

// GetCSAIAPR is a free data retrieval call binding the contract method 0x638aa95b.
//
// Solidity: function getCSAIAPR() constant returns(uint256)
func (_Apr *AprSession) GetCSAIAPR() (*big.Int, error) {
	return _Apr.Contract.GetCSAIAPR(&_Apr.CallOpts)
}

// GetCSAIAPR is a free data retrieval call binding the contract method 0x638aa95b.
//
// Solidity: function getCSAIAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetCSAIAPR() (*big.Int, error) {
	return _Apr.Contract.GetCSAIAPR(&_Apr.CallOpts)
}

// GetCUSDCAPR is a free data retrieval call binding the contract method 0x2ea35460.
//
// Solidity: function getCUSDCAPR() constant returns(uint256)
func (_Apr *AprCaller) GetCUSDCAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getCUSDCAPR")
	return *ret0, err
}

// GetCUSDCAPR is a free data retrieval call binding the contract method 0x2ea35460.
//
// Solidity: function getCUSDCAPR() constant returns(uint256)
func (_Apr *AprSession) GetCUSDCAPR() (*big.Int, error) {
	return _Apr.Contract.GetCUSDCAPR(&_Apr.CallOpts)
}

// GetCUSDCAPR is a free data retrieval call binding the contract method 0x2ea35460.
//
// Solidity: function getCUSDCAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetCUSDCAPR() (*big.Int, error) {
	return _Apr.Contract.GetCUSDCAPR(&_Apr.CallOpts)
}

// GetCWBTCAPR is a free data retrieval call binding the contract method 0xdae180b8.
//
// Solidity: function getCWBTCAPR() constant returns(uint256)
func (_Apr *AprCaller) GetCWBTCAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getCWBTCAPR")
	return *ret0, err
}

// GetCWBTCAPR is a free data retrieval call binding the contract method 0xdae180b8.
//
// Solidity: function getCWBTCAPR() constant returns(uint256)
func (_Apr *AprSession) GetCWBTCAPR() (*big.Int, error) {
	return _Apr.Contract.GetCWBTCAPR(&_Apr.CallOpts)
}

// GetCWBTCAPR is a free data retrieval call binding the contract method 0xdae180b8.
//
// Solidity: function getCWBTCAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetCWBTCAPR() (*big.Int, error) {
	return _Apr.Contract.GetCWBTCAPR(&_Apr.CallOpts)
}

// GetCZRCAPR is a free data retrieval call binding the contract method 0xc9f82170.
//
// Solidity: function getCZRCAPR() constant returns(uint256)
func (_Apr *AprCaller) GetCZRCAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getCZRCAPR")
	return *ret0, err
}

// GetCZRCAPR is a free data retrieval call binding the contract method 0xc9f82170.
//
// Solidity: function getCZRCAPR() constant returns(uint256)
func (_Apr *AprSession) GetCZRCAPR() (*big.Int, error) {
	return _Apr.Contract.GetCZRCAPR(&_Apr.CallOpts)
}

// GetCZRCAPR is a free data retrieval call binding the contract method 0xc9f82170.
//
// Solidity: function getCZRCAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetCZRCAPR() (*big.Int, error) {
	return _Apr.Contract.GetCZRCAPR(&_Apr.CallOpts)
}

// GetCompoundAPR is a free data retrieval call binding the contract method 0x44e258b4.
//
// Solidity: function getCompoundAPR(address token) constant returns(uint256)
func (_Apr *AprCaller) GetCompoundAPR(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getCompoundAPR", token)
	return *ret0, err
}

// GetCompoundAPR is a free data retrieval call binding the contract method 0x44e258b4.
//
// Solidity: function getCompoundAPR(address token) constant returns(uint256)
func (_Apr *AprSession) GetCompoundAPR(token common.Address) (*big.Int, error) {
	return _Apr.Contract.GetCompoundAPR(&_Apr.CallOpts, token)
}

// GetCompoundAPR is a free data retrieval call binding the contract method 0x44e258b4.
//
// Solidity: function getCompoundAPR(address token) constant returns(uint256)
func (_Apr *AprCallerSession) GetCompoundAPR(token common.Address) (*big.Int, error) {
	return _Apr.Contract.GetCompoundAPR(&_Apr.CallOpts, token)
}

// GetDyDxAPR is a free data retrieval call binding the contract method 0x81d24d8d.
//
// Solidity: function getDyDxAPR(uint256 marketId) constant returns(uint256)
func (_Apr *AprCaller) GetDyDxAPR(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getDyDxAPR", marketId)
	return *ret0, err
}

// GetDyDxAPR is a free data retrieval call binding the contract method 0x81d24d8d.
//
// Solidity: function getDyDxAPR(uint256 marketId) constant returns(uint256)
func (_Apr *AprSession) GetDyDxAPR(marketId *big.Int) (*big.Int, error) {
	return _Apr.Contract.GetDyDxAPR(&_Apr.CallOpts, marketId)
}

// GetDyDxAPR is a free data retrieval call binding the contract method 0x81d24d8d.
//
// Solidity: function getDyDxAPR(uint256 marketId) constant returns(uint256)
func (_Apr *AprCallerSession) GetDyDxAPR(marketId *big.Int) (*big.Int, error) {
	return _Apr.Contract.GetDyDxAPR(&_Apr.CallOpts, marketId)
}

// GetDyDxDAIAPR is a free data retrieval call binding the contract method 0x470bce80.
//
// Solidity: function getDyDxDAIAPR() constant returns(uint256)
func (_Apr *AprCaller) GetDyDxDAIAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getDyDxDAIAPR")
	return *ret0, err
}

// GetDyDxDAIAPR is a free data retrieval call binding the contract method 0x470bce80.
//
// Solidity: function getDyDxDAIAPR() constant returns(uint256)
func (_Apr *AprSession) GetDyDxDAIAPR() (*big.Int, error) {
	return _Apr.Contract.GetDyDxDAIAPR(&_Apr.CallOpts)
}

// GetDyDxDAIAPR is a free data retrieval call binding the contract method 0x470bce80.
//
// Solidity: function getDyDxDAIAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetDyDxDAIAPR() (*big.Int, error) {
	return _Apr.Contract.GetDyDxDAIAPR(&_Apr.CallOpts)
}

// GetDyDxETHAPR is a free data retrieval call binding the contract method 0x5af11ab1.
//
// Solidity: function getDyDxETHAPR() constant returns(uint256)
func (_Apr *AprCaller) GetDyDxETHAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getDyDxETHAPR")
	return *ret0, err
}

// GetDyDxETHAPR is a free data retrieval call binding the contract method 0x5af11ab1.
//
// Solidity: function getDyDxETHAPR() constant returns(uint256)
func (_Apr *AprSession) GetDyDxETHAPR() (*big.Int, error) {
	return _Apr.Contract.GetDyDxETHAPR(&_Apr.CallOpts)
}

// GetDyDxETHAPR is a free data retrieval call binding the contract method 0x5af11ab1.
//
// Solidity: function getDyDxETHAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetDyDxETHAPR() (*big.Int, error) {
	return _Apr.Contract.GetDyDxETHAPR(&_Apr.CallOpts)
}

// GetDyDxSAIAPR is a free data retrieval call binding the contract method 0x653d8ad9.
//
// Solidity: function getDyDxSAIAPR() constant returns(uint256)
func (_Apr *AprCaller) GetDyDxSAIAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getDyDxSAIAPR")
	return *ret0, err
}

// GetDyDxSAIAPR is a free data retrieval call binding the contract method 0x653d8ad9.
//
// Solidity: function getDyDxSAIAPR() constant returns(uint256)
func (_Apr *AprSession) GetDyDxSAIAPR() (*big.Int, error) {
	return _Apr.Contract.GetDyDxSAIAPR(&_Apr.CallOpts)
}

// GetDyDxSAIAPR is a free data retrieval call binding the contract method 0x653d8ad9.
//
// Solidity: function getDyDxSAIAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetDyDxSAIAPR() (*big.Int, error) {
	return _Apr.Contract.GetDyDxSAIAPR(&_Apr.CallOpts)
}

// GetDyDxUSDCAPR is a free data retrieval call binding the contract method 0x3f9c8622.
//
// Solidity: function getDyDxUSDCAPR() constant returns(uint256)
func (_Apr *AprCaller) GetDyDxUSDCAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getDyDxUSDCAPR")
	return *ret0, err
}

// GetDyDxUSDCAPR is a free data retrieval call binding the contract method 0x3f9c8622.
//
// Solidity: function getDyDxUSDCAPR() constant returns(uint256)
func (_Apr *AprSession) GetDyDxUSDCAPR() (*big.Int, error) {
	return _Apr.Contract.GetDyDxUSDCAPR(&_Apr.CallOpts)
}

// GetDyDxUSDCAPR is a free data retrieval call binding the contract method 0x3f9c8622.
//
// Solidity: function getDyDxUSDCAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetDyDxUSDCAPR() (*big.Int, error) {
	return _Apr.Contract.GetDyDxUSDCAPR(&_Apr.CallOpts)
}

// GetFulcrumAPR is a free data retrieval call binding the contract method 0x83deca3d.
//
// Solidity: function getFulcrumAPR(address token) constant returns(uint256)
func (_Apr *AprCaller) GetFulcrumAPR(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getFulcrumAPR", token)
	return *ret0, err
}

// GetFulcrumAPR is a free data retrieval call binding the contract method 0x83deca3d.
//
// Solidity: function getFulcrumAPR(address token) constant returns(uint256)
func (_Apr *AprSession) GetFulcrumAPR(token common.Address) (*big.Int, error) {
	return _Apr.Contract.GetFulcrumAPR(&_Apr.CallOpts, token)
}

// GetFulcrumAPR is a free data retrieval call binding the contract method 0x83deca3d.
//
// Solidity: function getFulcrumAPR(address token) constant returns(uint256)
func (_Apr *AprCallerSession) GetFulcrumAPR(token common.Address) (*big.Int, error) {
	return _Apr.Contract.GetFulcrumAPR(&_Apr.CallOpts, token)
}

// GetIDAIAPR is a free data retrieval call binding the contract method 0x43e8b0be.
//
// Solidity: function getIDAIAPR() constant returns(uint256)
func (_Apr *AprCaller) GetIDAIAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getIDAIAPR")
	return *ret0, err
}

// GetIDAIAPR is a free data retrieval call binding the contract method 0x43e8b0be.
//
// Solidity: function getIDAIAPR() constant returns(uint256)
func (_Apr *AprSession) GetIDAIAPR() (*big.Int, error) {
	return _Apr.Contract.GetIDAIAPR(&_Apr.CallOpts)
}

// GetIDAIAPR is a free data retrieval call binding the contract method 0x43e8b0be.
//
// Solidity: function getIDAIAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetIDAIAPR() (*big.Int, error) {
	return _Apr.Contract.GetIDAIAPR(&_Apr.CallOpts)
}

// GetIETHAPR is a free data retrieval call binding the contract method 0x22cc56d9.
//
// Solidity: function getIETHAPR() constant returns(uint256)
func (_Apr *AprCaller) GetIETHAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getIETHAPR")
	return *ret0, err
}

// GetIETHAPR is a free data retrieval call binding the contract method 0x22cc56d9.
//
// Solidity: function getIETHAPR() constant returns(uint256)
func (_Apr *AprSession) GetIETHAPR() (*big.Int, error) {
	return _Apr.Contract.GetIETHAPR(&_Apr.CallOpts)
}

// GetIETHAPR is a free data retrieval call binding the contract method 0x22cc56d9.
//
// Solidity: function getIETHAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetIETHAPR() (*big.Int, error) {
	return _Apr.Contract.GetIETHAPR(&_Apr.CallOpts)
}

// GetIKNCAPR is a free data retrieval call binding the contract method 0xbb540c82.
//
// Solidity: function getIKNCAPR() constant returns(uint256)
func (_Apr *AprCaller) GetIKNCAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getIKNCAPR")
	return *ret0, err
}

// GetIKNCAPR is a free data retrieval call binding the contract method 0xbb540c82.
//
// Solidity: function getIKNCAPR() constant returns(uint256)
func (_Apr *AprSession) GetIKNCAPR() (*big.Int, error) {
	return _Apr.Contract.GetIKNCAPR(&_Apr.CallOpts)
}

// GetIKNCAPR is a free data retrieval call binding the contract method 0xbb540c82.
//
// Solidity: function getIKNCAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetIKNCAPR() (*big.Int, error) {
	return _Apr.Contract.GetIKNCAPR(&_Apr.CallOpts)
}

// GetILINKAPR is a free data retrieval call binding the contract method 0x5d5bd5d3.
//
// Solidity: function getILINKAPR() constant returns(uint256)
func (_Apr *AprCaller) GetILINKAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getILINKAPR")
	return *ret0, err
}

// GetILINKAPR is a free data retrieval call binding the contract method 0x5d5bd5d3.
//
// Solidity: function getILINKAPR() constant returns(uint256)
func (_Apr *AprSession) GetILINKAPR() (*big.Int, error) {
	return _Apr.Contract.GetILINKAPR(&_Apr.CallOpts)
}

// GetILINKAPR is a free data retrieval call binding the contract method 0x5d5bd5d3.
//
// Solidity: function getILINKAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetILINKAPR() (*big.Int, error) {
	return _Apr.Contract.GetILINKAPR(&_Apr.CallOpts)
}

// GetIREPAPR is a free data retrieval call binding the contract method 0x7352023f.
//
// Solidity: function getIREPAPR() constant returns(uint256)
func (_Apr *AprCaller) GetIREPAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getIREPAPR")
	return *ret0, err
}

// GetIREPAPR is a free data retrieval call binding the contract method 0x7352023f.
//
// Solidity: function getIREPAPR() constant returns(uint256)
func (_Apr *AprSession) GetIREPAPR() (*big.Int, error) {
	return _Apr.Contract.GetIREPAPR(&_Apr.CallOpts)
}

// GetIREPAPR is a free data retrieval call binding the contract method 0x7352023f.
//
// Solidity: function getIREPAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetIREPAPR() (*big.Int, error) {
	return _Apr.Contract.GetIREPAPR(&_Apr.CallOpts)
}

// GetISAIAPR is a free data retrieval call binding the contract method 0x3cf1e7b8.
//
// Solidity: function getISAIAPR() constant returns(uint256)
func (_Apr *AprCaller) GetISAIAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getISAIAPR")
	return *ret0, err
}

// GetISAIAPR is a free data retrieval call binding the contract method 0x3cf1e7b8.
//
// Solidity: function getISAIAPR() constant returns(uint256)
func (_Apr *AprSession) GetISAIAPR() (*big.Int, error) {
	return _Apr.Contract.GetISAIAPR(&_Apr.CallOpts)
}

// GetISAIAPR is a free data retrieval call binding the contract method 0x3cf1e7b8.
//
// Solidity: function getISAIAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetISAIAPR() (*big.Int, error) {
	return _Apr.Contract.GetISAIAPR(&_Apr.CallOpts)
}

// GetISUSDAPR is a free data retrieval call binding the contract method 0x0e1e23e2.
//
// Solidity: function getISUSDAPR() constant returns(uint256)
func (_Apr *AprCaller) GetISUSDAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getISUSDAPR")
	return *ret0, err
}

// GetISUSDAPR is a free data retrieval call binding the contract method 0x0e1e23e2.
//
// Solidity: function getISUSDAPR() constant returns(uint256)
func (_Apr *AprSession) GetISUSDAPR() (*big.Int, error) {
	return _Apr.Contract.GetISUSDAPR(&_Apr.CallOpts)
}

// GetISUSDAPR is a free data retrieval call binding the contract method 0x0e1e23e2.
//
// Solidity: function getISUSDAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetISUSDAPR() (*big.Int, error) {
	return _Apr.Contract.GetISUSDAPR(&_Apr.CallOpts)
}

// GetIUSDCAPR is a free data retrieval call binding the contract method 0x106f39d1.
//
// Solidity: function getIUSDCAPR() constant returns(uint256)
func (_Apr *AprCaller) GetIUSDCAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getIUSDCAPR")
	return *ret0, err
}

// GetIUSDCAPR is a free data retrieval call binding the contract method 0x106f39d1.
//
// Solidity: function getIUSDCAPR() constant returns(uint256)
func (_Apr *AprSession) GetIUSDCAPR() (*big.Int, error) {
	return _Apr.Contract.GetIUSDCAPR(&_Apr.CallOpts)
}

// GetIUSDCAPR is a free data retrieval call binding the contract method 0x106f39d1.
//
// Solidity: function getIUSDCAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetIUSDCAPR() (*big.Int, error) {
	return _Apr.Contract.GetIUSDCAPR(&_Apr.CallOpts)
}

// GetIWBTCAPR is a free data retrieval call binding the contract method 0xb35946e2.
//
// Solidity: function getIWBTCAPR() constant returns(uint256)
func (_Apr *AprCaller) GetIWBTCAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getIWBTCAPR")
	return *ret0, err
}

// GetIWBTCAPR is a free data retrieval call binding the contract method 0xb35946e2.
//
// Solidity: function getIWBTCAPR() constant returns(uint256)
func (_Apr *AprSession) GetIWBTCAPR() (*big.Int, error) {
	return _Apr.Contract.GetIWBTCAPR(&_Apr.CallOpts)
}

// GetIWBTCAPR is a free data retrieval call binding the contract method 0xb35946e2.
//
// Solidity: function getIWBTCAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetIWBTCAPR() (*big.Int, error) {
	return _Apr.Contract.GetIWBTCAPR(&_Apr.CallOpts)
}

// GetIZRXAPR is a free data retrieval call binding the contract method 0xbdbfcd1c.
//
// Solidity: function getIZRXAPR() constant returns(uint256)
func (_Apr *AprCaller) GetIZRXAPR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getIZRXAPR")
	return *ret0, err
}

// GetIZRXAPR is a free data retrieval call binding the contract method 0xbdbfcd1c.
//
// Solidity: function getIZRXAPR() constant returns(uint256)
func (_Apr *AprSession) GetIZRXAPR() (*big.Int, error) {
	return _Apr.Contract.GetIZRXAPR(&_Apr.CallOpts)
}

// GetIZRXAPR is a free data retrieval call binding the contract method 0xbdbfcd1c.
//
// Solidity: function getIZRXAPR() constant returns(uint256)
func (_Apr *AprCallerSession) GetIZRXAPR() (*big.Int, error) {
	return _Apr.Contract.GetIZRXAPR(&_Apr.CallOpts)
}

// GetLiquidity is a free data retrieval call binding the contract method 0xa747b93b.
//
// Solidity: function getLiquidity(address _token) constant returns(uint256)
func (_Apr *AprCaller) GetLiquidity(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getLiquidity", _token)
	return *ret0, err
}

// GetLiquidity is a free data retrieval call binding the contract method 0xa747b93b.
//
// Solidity: function getLiquidity(address _token) constant returns(uint256)
func (_Apr *AprSession) GetLiquidity(_token common.Address) (*big.Int, error) {
	return _Apr.Contract.GetLiquidity(&_Apr.CallOpts, _token)
}

// GetLiquidity is a free data retrieval call binding the contract method 0xa747b93b.
//
// Solidity: function getLiquidity(address _token) constant returns(uint256)
func (_Apr *AprCallerSession) GetLiquidity(_token common.Address) (*big.Int, error) {
	return _Apr.Contract.GetLiquidity(&_Apr.CallOpts, _token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _token) constant returns(uint256)
func (_Apr *AprCaller) GetPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "getPrice", _token)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _token) constant returns(uint256)
func (_Apr *AprSession) GetPrice(_token common.Address) (*big.Int, error) {
	return _Apr.Contract.GetPrice(&_Apr.CallOpts, _token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _token) constant returns(uint256)
func (_Apr *AprCallerSession) GetPrice(_token common.Address) (*big.Int, error) {
	return _Apr.Contract.GetPrice(&_Apr.CallOpts, _token)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Apr *AprCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Apr *AprSession) IsOwner() (bool, error) {
	return _Apr.Contract.IsOwner(&_Apr.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Apr *AprCallerSession) IsOwner() (bool, error) {
	return _Apr.Contract.IsOwner(&_Apr.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Apr *AprCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Apr *AprSession) Oracle() (common.Address, error) {
	return _Apr.Contract.Oracle(&_Apr.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_Apr *AprCallerSession) Oracle() (common.Address, error) {
	return _Apr.Contract.Oracle(&_Apr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Apr *AprCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Apr *AprSession) Owner() (common.Address, error) {
	return _Apr.Contract.Owner(&_Apr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Apr *AprCallerSession) Owner() (common.Address, error) {
	return _Apr.Contract.Owner(&_Apr.CallOpts)
}

// RecommendDAI is a free data retrieval call binding the contract method 0xb5487801.
//
// Solidity: function recommendDAI() constant returns(string)
func (_Apr *AprCaller) RecommendDAI(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "recommendDAI")
	return *ret0, err
}

// RecommendDAI is a free data retrieval call binding the contract method 0xb5487801.
//
// Solidity: function recommendDAI() constant returns(string)
func (_Apr *AprSession) RecommendDAI() (string, error) {
	return _Apr.Contract.RecommendDAI(&_Apr.CallOpts)
}

// RecommendDAI is a free data retrieval call binding the contract method 0xb5487801.
//
// Solidity: function recommendDAI() constant returns(string)
func (_Apr *AprCallerSession) RecommendDAI() (string, error) {
	return _Apr.Contract.RecommendDAI(&_Apr.CallOpts)
}

// RecommendETH is a free data retrieval call binding the contract method 0x892eac13.
//
// Solidity: function recommendETH() constant returns(string)
func (_Apr *AprCaller) RecommendETH(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "recommendETH")
	return *ret0, err
}

// RecommendETH is a free data retrieval call binding the contract method 0x892eac13.
//
// Solidity: function recommendETH() constant returns(string)
func (_Apr *AprSession) RecommendETH() (string, error) {
	return _Apr.Contract.RecommendETH(&_Apr.CallOpts)
}

// RecommendETH is a free data retrieval call binding the contract method 0x892eac13.
//
// Solidity: function recommendETH() constant returns(string)
func (_Apr *AprCallerSession) RecommendETH() (string, error) {
	return _Apr.Contract.RecommendETH(&_Apr.CallOpts)
}

// RecommendUSDC is a free data retrieval call binding the contract method 0x04eb246d.
//
// Solidity: function recommendUSDC() constant returns(string)
func (_Apr *AprCaller) RecommendUSDC(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Apr.contract.Call(opts, out, "recommendUSDC")
	return *ret0, err
}

// RecommendUSDC is a free data retrieval call binding the contract method 0x04eb246d.
//
// Solidity: function recommendUSDC() constant returns(string)
func (_Apr *AprSession) RecommendUSDC() (string, error) {
	return _Apr.Contract.RecommendUSDC(&_Apr.CallOpts)
}

// RecommendUSDC is a free data retrieval call binding the contract method 0x04eb246d.
//
// Solidity: function recommendUSDC() constant returns(string)
func (_Apr *AprCallerSession) RecommendUSDC() (string, error) {
	return _Apr.Contract.RecommendUSDC(&_Apr.CallOpts)
}

// InCaseETHGetsStuck is a paid mutator transaction binding the contract method 0xf60a15ed.
//
// Solidity: function inCaseETHGetsStuck() returns()
func (_Apr *AprTransactor) InCaseETHGetsStuck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "inCaseETHGetsStuck")
}

// InCaseETHGetsStuck is a paid mutator transaction binding the contract method 0xf60a15ed.
//
// Solidity: function inCaseETHGetsStuck() returns()
func (_Apr *AprSession) InCaseETHGetsStuck() (*types.Transaction, error) {
	return _Apr.Contract.InCaseETHGetsStuck(&_Apr.TransactOpts)
}

// InCaseETHGetsStuck is a paid mutator transaction binding the contract method 0xf60a15ed.
//
// Solidity: function inCaseETHGetsStuck() returns()
func (_Apr *AprTransactorSession) InCaseETHGetsStuck() (*types.Transaction, error) {
	return _Apr.Contract.InCaseETHGetsStuck(&_Apr.TransactOpts)
}

// InCaseTokenGetsStuck is a paid mutator transaction binding the contract method 0x2e8d6e18.
//
// Solidity: function inCaseTokenGetsStuck(address _TokenAddress) returns()
func (_Apr *AprTransactor) InCaseTokenGetsStuck(opts *bind.TransactOpts, _TokenAddress common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "inCaseTokenGetsStuck", _TokenAddress)
}

// InCaseTokenGetsStuck is a paid mutator transaction binding the contract method 0x2e8d6e18.
//
// Solidity: function inCaseTokenGetsStuck(address _TokenAddress) returns()
func (_Apr *AprSession) InCaseTokenGetsStuck(_TokenAddress common.Address) (*types.Transaction, error) {
	return _Apr.Contract.InCaseTokenGetsStuck(&_Apr.TransactOpts, _TokenAddress)
}

// InCaseTokenGetsStuck is a paid mutator transaction binding the contract method 0x2e8d6e18.
//
// Solidity: function inCaseTokenGetsStuck(address _TokenAddress) returns()
func (_Apr *AprTransactorSession) InCaseTokenGetsStuck(_TokenAddress common.Address) (*types.Transaction, error) {
	return _Apr.Contract.InCaseTokenGetsStuck(&_Apr.TransactOpts, _TokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Apr *AprTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Apr *AprSession) RenounceOwnership() (*types.Transaction, error) {
	return _Apr.Contract.RenounceOwnership(&_Apr.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Apr *AprTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Apr.Contract.RenounceOwnership(&_Apr.TransactOpts)
}

// SetLiquidity is a paid mutator transaction binding the contract method 0xf4069cba.
//
// Solidity: function setLiquidity(address _token, uint256 _liquidity) returns()
func (_Apr *AprTransactor) SetLiquidity(opts *bind.TransactOpts, _token common.Address, _liquidity *big.Int) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "setLiquidity", _token, _liquidity)
}

// SetLiquidity is a paid mutator transaction binding the contract method 0xf4069cba.
//
// Solidity: function setLiquidity(address _token, uint256 _liquidity) returns()
func (_Apr *AprSession) SetLiquidity(_token common.Address, _liquidity *big.Int) (*types.Transaction, error) {
	return _Apr.Contract.SetLiquidity(&_Apr.TransactOpts, _token, _liquidity)
}

// SetLiquidity is a paid mutator transaction binding the contract method 0xf4069cba.
//
// Solidity: function setLiquidity(address _token, uint256 _liquidity) returns()
func (_Apr *AprTransactorSession) SetLiquidity(_token common.Address, _liquidity *big.Int) (*types.Transaction, error) {
	return _Apr.Contract.SetLiquidity(&_Apr.TransactOpts, _token, _liquidity)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address _token, uint256 _price) returns()
func (_Apr *AprTransactor) SetPrice(opts *bind.TransactOpts, _token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "setPrice", _token, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address _token, uint256 _price) returns()
func (_Apr *AprSession) SetPrice(_token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Apr.Contract.SetPrice(&_Apr.TransactOpts, _token, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address _token, uint256 _price) returns()
func (_Apr *AprTransactorSession) SetPrice(_token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Apr.Contract.SetPrice(&_Apr.TransactOpts, _token, _price)
}

// SetNewAAVE is a paid mutator transaction binding the contract method 0x812adb06.
//
// Solidity: function set_new_AAVE(address _new_AAVE) returns()
func (_Apr *AprTransactor) SetNewAAVE(opts *bind.TransactOpts, _new_AAVE common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_AAVE", _new_AAVE)
}

// SetNewAAVE is a paid mutator transaction binding the contract method 0x812adb06.
//
// Solidity: function set_new_AAVE(address _new_AAVE) returns()
func (_Apr *AprSession) SetNewAAVE(_new_AAVE common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAAVE(&_Apr.TransactOpts, _new_AAVE)
}

// SetNewAAVE is a paid mutator transaction binding the contract method 0x812adb06.
//
// Solidity: function set_new_AAVE(address _new_AAVE) returns()
func (_Apr *AprTransactorSession) SetNewAAVE(_new_AAVE common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAAVE(&_Apr.TransactOpts, _new_AAVE)
}

// SetNewABAT is a paid mutator transaction binding the contract method 0xb72d34c2.
//
// Solidity: function set_new_ABAT(address _new_ABAT) returns()
func (_Apr *AprTransactor) SetNewABAT(opts *bind.TransactOpts, _new_ABAT common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_ABAT", _new_ABAT)
}

// SetNewABAT is a paid mutator transaction binding the contract method 0xb72d34c2.
//
// Solidity: function set_new_ABAT(address _new_ABAT) returns()
func (_Apr *AprSession) SetNewABAT(_new_ABAT common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewABAT(&_Apr.TransactOpts, _new_ABAT)
}

// SetNewABAT is a paid mutator transaction binding the contract method 0xb72d34c2.
//
// Solidity: function set_new_ABAT(address _new_ABAT) returns()
func (_Apr *AprTransactorSession) SetNewABAT(_new_ABAT common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewABAT(&_Apr.TransactOpts, _new_ABAT)
}

// SetNewADAI is a paid mutator transaction binding the contract method 0xd6a79ada.
//
// Solidity: function set_new_ADAI(address _new_ADAI) returns()
func (_Apr *AprTransactor) SetNewADAI(opts *bind.TransactOpts, _new_ADAI common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_ADAI", _new_ADAI)
}

// SetNewADAI is a paid mutator transaction binding the contract method 0xd6a79ada.
//
// Solidity: function set_new_ADAI(address _new_ADAI) returns()
func (_Apr *AprSession) SetNewADAI(_new_ADAI common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewADAI(&_Apr.TransactOpts, _new_ADAI)
}

// SetNewADAI is a paid mutator transaction binding the contract method 0xd6a79ada.
//
// Solidity: function set_new_ADAI(address _new_ADAI) returns()
func (_Apr *AprTransactorSession) SetNewADAI(_new_ADAI common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewADAI(&_Apr.TransactOpts, _new_ADAI)
}

// SetNewAETH is a paid mutator transaction binding the contract method 0xccc41681.
//
// Solidity: function set_new_AETH(address _new_AETH) returns()
func (_Apr *AprTransactor) SetNewAETH(opts *bind.TransactOpts, _new_AETH common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_AETH", _new_AETH)
}

// SetNewAETH is a paid mutator transaction binding the contract method 0xccc41681.
//
// Solidity: function set_new_AETH(address _new_AETH) returns()
func (_Apr *AprSession) SetNewAETH(_new_AETH common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAETH(&_Apr.TransactOpts, _new_AETH)
}

// SetNewAETH is a paid mutator transaction binding the contract method 0xccc41681.
//
// Solidity: function set_new_AETH(address _new_AETH) returns()
func (_Apr *AprTransactorSession) SetNewAETH(_new_AETH common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAETH(&_Apr.TransactOpts, _new_AETH)
}

// SetNewAKNC is a paid mutator transaction binding the contract method 0x0dc00a37.
//
// Solidity: function set_new_AKNC(address _new_AKNC) returns()
func (_Apr *AprTransactor) SetNewAKNC(opts *bind.TransactOpts, _new_AKNC common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_AKNC", _new_AKNC)
}

// SetNewAKNC is a paid mutator transaction binding the contract method 0x0dc00a37.
//
// Solidity: function set_new_AKNC(address _new_AKNC) returns()
func (_Apr *AprSession) SetNewAKNC(_new_AKNC common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAKNC(&_Apr.TransactOpts, _new_AKNC)
}

// SetNewAKNC is a paid mutator transaction binding the contract method 0x0dc00a37.
//
// Solidity: function set_new_AKNC(address _new_AKNC) returns()
func (_Apr *AprTransactorSession) SetNewAKNC(_new_AKNC common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAKNC(&_Apr.TransactOpts, _new_AKNC)
}

// SetNewALEND is a paid mutator transaction binding the contract method 0x793957b7.
//
// Solidity: function set_new_ALEND(address _new_ALEND) returns()
func (_Apr *AprTransactor) SetNewALEND(opts *bind.TransactOpts, _new_ALEND common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_ALEND", _new_ALEND)
}

// SetNewALEND is a paid mutator transaction binding the contract method 0x793957b7.
//
// Solidity: function set_new_ALEND(address _new_ALEND) returns()
func (_Apr *AprSession) SetNewALEND(_new_ALEND common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewALEND(&_Apr.TransactOpts, _new_ALEND)
}

// SetNewALEND is a paid mutator transaction binding the contract method 0x793957b7.
//
// Solidity: function set_new_ALEND(address _new_ALEND) returns()
func (_Apr *AprTransactorSession) SetNewALEND(_new_ALEND common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewALEND(&_Apr.TransactOpts, _new_ALEND)
}

// SetNewALINK is a paid mutator transaction binding the contract method 0xd2db39fe.
//
// Solidity: function set_new_ALINK(address _new_ALINK) returns()
func (_Apr *AprTransactor) SetNewALINK(opts *bind.TransactOpts, _new_ALINK common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_ALINK", _new_ALINK)
}

// SetNewALINK is a paid mutator transaction binding the contract method 0xd2db39fe.
//
// Solidity: function set_new_ALINK(address _new_ALINK) returns()
func (_Apr *AprSession) SetNewALINK(_new_ALINK common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewALINK(&_Apr.TransactOpts, _new_ALINK)
}

// SetNewALINK is a paid mutator transaction binding the contract method 0xd2db39fe.
//
// Solidity: function set_new_ALINK(address _new_ALINK) returns()
func (_Apr *AprTransactorSession) SetNewALINK(_new_ALINK common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewALINK(&_Apr.TransactOpts, _new_ALINK)
}

// SetNewAMANA is a paid mutator transaction binding the contract method 0xbfd5b402.
//
// Solidity: function set_new_AMANA(address _new_AMANA) returns()
func (_Apr *AprTransactor) SetNewAMANA(opts *bind.TransactOpts, _new_AMANA common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_AMANA", _new_AMANA)
}

// SetNewAMANA is a paid mutator transaction binding the contract method 0xbfd5b402.
//
// Solidity: function set_new_AMANA(address _new_AMANA) returns()
func (_Apr *AprSession) SetNewAMANA(_new_AMANA common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAMANA(&_Apr.TransactOpts, _new_AMANA)
}

// SetNewAMANA is a paid mutator transaction binding the contract method 0xbfd5b402.
//
// Solidity: function set_new_AMANA(address _new_AMANA) returns()
func (_Apr *AprTransactorSession) SetNewAMANA(_new_AMANA common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAMANA(&_Apr.TransactOpts, _new_AMANA)
}

// SetNewAMKR is a paid mutator transaction binding the contract method 0xbfaf715b.
//
// Solidity: function set_new_AMKR(address _new_AMKR) returns()
func (_Apr *AprTransactor) SetNewAMKR(opts *bind.TransactOpts, _new_AMKR common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_AMKR", _new_AMKR)
}

// SetNewAMKR is a paid mutator transaction binding the contract method 0xbfaf715b.
//
// Solidity: function set_new_AMKR(address _new_AMKR) returns()
func (_Apr *AprSession) SetNewAMKR(_new_AMKR common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAMKR(&_Apr.TransactOpts, _new_AMKR)
}

// SetNewAMKR is a paid mutator transaction binding the contract method 0xbfaf715b.
//
// Solidity: function set_new_AMKR(address _new_AMKR) returns()
func (_Apr *AprTransactorSession) SetNewAMKR(_new_AMKR common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAMKR(&_Apr.TransactOpts, _new_AMKR)
}

// SetNewAREP is a paid mutator transaction binding the contract method 0x3f8ce23f.
//
// Solidity: function set_new_AREP(address _new_AREP) returns()
func (_Apr *AprTransactor) SetNewAREP(opts *bind.TransactOpts, _new_AREP common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_AREP", _new_AREP)
}

// SetNewAREP is a paid mutator transaction binding the contract method 0x3f8ce23f.
//
// Solidity: function set_new_AREP(address _new_AREP) returns()
func (_Apr *AprSession) SetNewAREP(_new_AREP common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAREP(&_Apr.TransactOpts, _new_AREP)
}

// SetNewAREP is a paid mutator transaction binding the contract method 0x3f8ce23f.
//
// Solidity: function set_new_AREP(address _new_AREP) returns()
func (_Apr *AprTransactorSession) SetNewAREP(_new_AREP common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAREP(&_Apr.TransactOpts, _new_AREP)
}

// SetNewASNX is a paid mutator transaction binding the contract method 0x4958aa07.
//
// Solidity: function set_new_ASNX(address _new_ASNX) returns()
func (_Apr *AprTransactor) SetNewASNX(opts *bind.TransactOpts, _new_ASNX common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_ASNX", _new_ASNX)
}

// SetNewASNX is a paid mutator transaction binding the contract method 0x4958aa07.
//
// Solidity: function set_new_ASNX(address _new_ASNX) returns()
func (_Apr *AprSession) SetNewASNX(_new_ASNX common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewASNX(&_Apr.TransactOpts, _new_ASNX)
}

// SetNewASNX is a paid mutator transaction binding the contract method 0x4958aa07.
//
// Solidity: function set_new_ASNX(address _new_ASNX) returns()
func (_Apr *AprTransactorSession) SetNewASNX(_new_ASNX common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewASNX(&_Apr.TransactOpts, _new_ASNX)
}

// SetNewASUSD is a paid mutator transaction binding the contract method 0x101b5668.
//
// Solidity: function set_new_ASUSD(address _new_ASUSD) returns()
func (_Apr *AprTransactor) SetNewASUSD(opts *bind.TransactOpts, _new_ASUSD common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_ASUSD", _new_ASUSD)
}

// SetNewASUSD is a paid mutator transaction binding the contract method 0x101b5668.
//
// Solidity: function set_new_ASUSD(address _new_ASUSD) returns()
func (_Apr *AprSession) SetNewASUSD(_new_ASUSD common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewASUSD(&_Apr.TransactOpts, _new_ASUSD)
}

// SetNewASUSD is a paid mutator transaction binding the contract method 0x101b5668.
//
// Solidity: function set_new_ASUSD(address _new_ASUSD) returns()
func (_Apr *AprTransactorSession) SetNewASUSD(_new_ASUSD common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewASUSD(&_Apr.TransactOpts, _new_ASUSD)
}

// SetNewATUSD is a paid mutator transaction binding the contract method 0x5309d952.
//
// Solidity: function set_new_ATUSD(address _new_ATUSD) returns()
func (_Apr *AprTransactor) SetNewATUSD(opts *bind.TransactOpts, _new_ATUSD common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_ATUSD", _new_ATUSD)
}

// SetNewATUSD is a paid mutator transaction binding the contract method 0x5309d952.
//
// Solidity: function set_new_ATUSD(address _new_ATUSD) returns()
func (_Apr *AprSession) SetNewATUSD(_new_ATUSD common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewATUSD(&_Apr.TransactOpts, _new_ATUSD)
}

// SetNewATUSD is a paid mutator transaction binding the contract method 0x5309d952.
//
// Solidity: function set_new_ATUSD(address _new_ATUSD) returns()
func (_Apr *AprTransactorSession) SetNewATUSD(_new_ATUSD common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewATUSD(&_Apr.TransactOpts, _new_ATUSD)
}

// SetNewAUSDC is a paid mutator transaction binding the contract method 0xf53a0628.
//
// Solidity: function set_new_AUSDC(address _new_AUSDC) returns()
func (_Apr *AprTransactor) SetNewAUSDC(opts *bind.TransactOpts, _new_AUSDC common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_AUSDC", _new_AUSDC)
}

// SetNewAUSDC is a paid mutator transaction binding the contract method 0xf53a0628.
//
// Solidity: function set_new_AUSDC(address _new_AUSDC) returns()
func (_Apr *AprSession) SetNewAUSDC(_new_AUSDC common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAUSDC(&_Apr.TransactOpts, _new_AUSDC)
}

// SetNewAUSDC is a paid mutator transaction binding the contract method 0xf53a0628.
//
// Solidity: function set_new_AUSDC(address _new_AUSDC) returns()
func (_Apr *AprTransactorSession) SetNewAUSDC(_new_AUSDC common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAUSDC(&_Apr.TransactOpts, _new_AUSDC)
}

// SetNewAUSDT is a paid mutator transaction binding the contract method 0x25e079b1.
//
// Solidity: function set_new_AUSDT(address _new_AUSDT) returns()
func (_Apr *AprTransactor) SetNewAUSDT(opts *bind.TransactOpts, _new_AUSDT common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_AUSDT", _new_AUSDT)
}

// SetNewAUSDT is a paid mutator transaction binding the contract method 0x25e079b1.
//
// Solidity: function set_new_AUSDT(address _new_AUSDT) returns()
func (_Apr *AprSession) SetNewAUSDT(_new_AUSDT common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAUSDT(&_Apr.TransactOpts, _new_AUSDT)
}

// SetNewAUSDT is a paid mutator transaction binding the contract method 0x25e079b1.
//
// Solidity: function set_new_AUSDT(address _new_AUSDT) returns()
func (_Apr *AprTransactorSession) SetNewAUSDT(_new_AUSDT common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAUSDT(&_Apr.TransactOpts, _new_AUSDT)
}

// SetNewAWBTC is a paid mutator transaction binding the contract method 0x98df5ccf.
//
// Solidity: function set_new_AWBTC(address _new_AWBTC) returns()
func (_Apr *AprTransactor) SetNewAWBTC(opts *bind.TransactOpts, _new_AWBTC common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_AWBTC", _new_AWBTC)
}

// SetNewAWBTC is a paid mutator transaction binding the contract method 0x98df5ccf.
//
// Solidity: function set_new_AWBTC(address _new_AWBTC) returns()
func (_Apr *AprSession) SetNewAWBTC(_new_AWBTC common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAWBTC(&_Apr.TransactOpts, _new_AWBTC)
}

// SetNewAWBTC is a paid mutator transaction binding the contract method 0x98df5ccf.
//
// Solidity: function set_new_AWBTC(address _new_AWBTC) returns()
func (_Apr *AprTransactorSession) SetNewAWBTC(_new_AWBTC common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAWBTC(&_Apr.TransactOpts, _new_AWBTC)
}

// SetNewAZRX is a paid mutator transaction binding the contract method 0x4e24f77f.
//
// Solidity: function set_new_AZRX(address _new_AZRX) returns()
func (_Apr *AprTransactor) SetNewAZRX(opts *bind.TransactOpts, _new_AZRX common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_AZRX", _new_AZRX)
}

// SetNewAZRX is a paid mutator transaction binding the contract method 0x4e24f77f.
//
// Solidity: function set_new_AZRX(address _new_AZRX) returns()
func (_Apr *AprSession) SetNewAZRX(_new_AZRX common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAZRX(&_Apr.TransactOpts, _new_AZRX)
}

// SetNewAZRX is a paid mutator transaction binding the contract method 0x4e24f77f.
//
// Solidity: function set_new_AZRX(address _new_AZRX) returns()
func (_Apr *AprTransactorSession) SetNewAZRX(_new_AZRX common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewAZRX(&_Apr.TransactOpts, _new_AZRX)
}

// SetNewCBAT is a paid mutator transaction binding the contract method 0xb74d4f31.
//
// Solidity: function set_new_CBAT(address _new_CBAT) returns()
func (_Apr *AprTransactor) SetNewCBAT(opts *bind.TransactOpts, _new_CBAT common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_CBAT", _new_CBAT)
}

// SetNewCBAT is a paid mutator transaction binding the contract method 0xb74d4f31.
//
// Solidity: function set_new_CBAT(address _new_CBAT) returns()
func (_Apr *AprSession) SetNewCBAT(_new_CBAT common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewCBAT(&_Apr.TransactOpts, _new_CBAT)
}

// SetNewCBAT is a paid mutator transaction binding the contract method 0xb74d4f31.
//
// Solidity: function set_new_CBAT(address _new_CBAT) returns()
func (_Apr *AprTransactorSession) SetNewCBAT(_new_CBAT common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewCBAT(&_Apr.TransactOpts, _new_CBAT)
}

// SetNewCDAI is a paid mutator transaction binding the contract method 0xad9a5e2a.
//
// Solidity: function set_new_CDAI(address _new_CDAI) returns()
func (_Apr *AprTransactor) SetNewCDAI(opts *bind.TransactOpts, _new_CDAI common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_CDAI", _new_CDAI)
}

// SetNewCDAI is a paid mutator transaction binding the contract method 0xad9a5e2a.
//
// Solidity: function set_new_CDAI(address _new_CDAI) returns()
func (_Apr *AprSession) SetNewCDAI(_new_CDAI common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewCDAI(&_Apr.TransactOpts, _new_CDAI)
}

// SetNewCDAI is a paid mutator transaction binding the contract method 0xad9a5e2a.
//
// Solidity: function set_new_CDAI(address _new_CDAI) returns()
func (_Apr *AprTransactorSession) SetNewCDAI(_new_CDAI common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewCDAI(&_Apr.TransactOpts, _new_CDAI)
}

// SetNewCETH is a paid mutator transaction binding the contract method 0x2e163c86.
//
// Solidity: function set_new_CETH(address _new_CETH) returns()
func (_Apr *AprTransactor) SetNewCETH(opts *bind.TransactOpts, _new_CETH common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_CETH", _new_CETH)
}

// SetNewCETH is a paid mutator transaction binding the contract method 0x2e163c86.
//
// Solidity: function set_new_CETH(address _new_CETH) returns()
func (_Apr *AprSession) SetNewCETH(_new_CETH common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewCETH(&_Apr.TransactOpts, _new_CETH)
}

// SetNewCETH is a paid mutator transaction binding the contract method 0x2e163c86.
//
// Solidity: function set_new_CETH(address _new_CETH) returns()
func (_Apr *AprTransactorSession) SetNewCETH(_new_CETH common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewCETH(&_Apr.TransactOpts, _new_CETH)
}

// SetNewCREP is a paid mutator transaction binding the contract method 0xc6b0d25f.
//
// Solidity: function set_new_CREP(address _new_CREP) returns()
func (_Apr *AprTransactor) SetNewCREP(opts *bind.TransactOpts, _new_CREP common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_CREP", _new_CREP)
}

// SetNewCREP is a paid mutator transaction binding the contract method 0xc6b0d25f.
//
// Solidity: function set_new_CREP(address _new_CREP) returns()
func (_Apr *AprSession) SetNewCREP(_new_CREP common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewCREP(&_Apr.TransactOpts, _new_CREP)
}

// SetNewCREP is a paid mutator transaction binding the contract method 0xc6b0d25f.
//
// Solidity: function set_new_CREP(address _new_CREP) returns()
func (_Apr *AprTransactorSession) SetNewCREP(_new_CREP common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewCREP(&_Apr.TransactOpts, _new_CREP)
}

// SetNewCSAI is a paid mutator transaction binding the contract method 0x1c384781.
//
// Solidity: function set_new_CSAI(address _new_CSAI) returns()
func (_Apr *AprTransactor) SetNewCSAI(opts *bind.TransactOpts, _new_CSAI common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_CSAI", _new_CSAI)
}

// SetNewCSAI is a paid mutator transaction binding the contract method 0x1c384781.
//
// Solidity: function set_new_CSAI(address _new_CSAI) returns()
func (_Apr *AprSession) SetNewCSAI(_new_CSAI common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewCSAI(&_Apr.TransactOpts, _new_CSAI)
}

// SetNewCSAI is a paid mutator transaction binding the contract method 0x1c384781.
//
// Solidity: function set_new_CSAI(address _new_CSAI) returns()
func (_Apr *AprTransactorSession) SetNewCSAI(_new_CSAI common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewCSAI(&_Apr.TransactOpts, _new_CSAI)
}

// SetNewCUSDC is a paid mutator transaction binding the contract method 0x86bbe196.
//
// Solidity: function set_new_CUSDC(address _new_CUSDC) returns()
func (_Apr *AprTransactor) SetNewCUSDC(opts *bind.TransactOpts, _new_CUSDC common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_CUSDC", _new_CUSDC)
}

// SetNewCUSDC is a paid mutator transaction binding the contract method 0x86bbe196.
//
// Solidity: function set_new_CUSDC(address _new_CUSDC) returns()
func (_Apr *AprSession) SetNewCUSDC(_new_CUSDC common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewCUSDC(&_Apr.TransactOpts, _new_CUSDC)
}

// SetNewCUSDC is a paid mutator transaction binding the contract method 0x86bbe196.
//
// Solidity: function set_new_CUSDC(address _new_CUSDC) returns()
func (_Apr *AprTransactorSession) SetNewCUSDC(_new_CUSDC common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewCUSDC(&_Apr.TransactOpts, _new_CUSDC)
}

// SetNewCWBTC is a paid mutator transaction binding the contract method 0x033672b7.
//
// Solidity: function set_new_CWBTC(address _new_CWBTC) returns()
func (_Apr *AprTransactor) SetNewCWBTC(opts *bind.TransactOpts, _new_CWBTC common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_CWBTC", _new_CWBTC)
}

// SetNewCWBTC is a paid mutator transaction binding the contract method 0x033672b7.
//
// Solidity: function set_new_CWBTC(address _new_CWBTC) returns()
func (_Apr *AprSession) SetNewCWBTC(_new_CWBTC common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewCWBTC(&_Apr.TransactOpts, _new_CWBTC)
}

// SetNewCWBTC is a paid mutator transaction binding the contract method 0x033672b7.
//
// Solidity: function set_new_CWBTC(address _new_CWBTC) returns()
func (_Apr *AprTransactorSession) SetNewCWBTC(_new_CWBTC common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewCWBTC(&_Apr.TransactOpts, _new_CWBTC)
}

// SetNewCZRX is a paid mutator transaction binding the contract method 0x78226b41.
//
// Solidity: function set_new_CZRX(address _new_CZRX) returns()
func (_Apr *AprTransactor) SetNewCZRX(opts *bind.TransactOpts, _new_CZRX common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_CZRX", _new_CZRX)
}

// SetNewCZRX is a paid mutator transaction binding the contract method 0x78226b41.
//
// Solidity: function set_new_CZRX(address _new_CZRX) returns()
func (_Apr *AprSession) SetNewCZRX(_new_CZRX common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewCZRX(&_Apr.TransactOpts, _new_CZRX)
}

// SetNewCZRX is a paid mutator transaction binding the contract method 0x78226b41.
//
// Solidity: function set_new_CZRX(address _new_CZRX) returns()
func (_Apr *AprTransactorSession) SetNewCZRX(_new_CZRX common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewCZRX(&_Apr.TransactOpts, _new_CZRX)
}

// SetNewDYDX is a paid mutator transaction binding the contract method 0x2c7a5ae4.
//
// Solidity: function set_new_DYDX(address _new_DYDX) returns()
func (_Apr *AprTransactor) SetNewDYDX(opts *bind.TransactOpts, _new_DYDX common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_DYDX", _new_DYDX)
}

// SetNewDYDX is a paid mutator transaction binding the contract method 0x2c7a5ae4.
//
// Solidity: function set_new_DYDX(address _new_DYDX) returns()
func (_Apr *AprSession) SetNewDYDX(_new_DYDX common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewDYDX(&_Apr.TransactOpts, _new_DYDX)
}

// SetNewDYDX is a paid mutator transaction binding the contract method 0x2c7a5ae4.
//
// Solidity: function set_new_DYDX(address _new_DYDX) returns()
func (_Apr *AprTransactorSession) SetNewDYDX(_new_DYDX common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewDYDX(&_Apr.TransactOpts, _new_DYDX)
}

// SetNewIDAI is a paid mutator transaction binding the contract method 0x857c3c9b.
//
// Solidity: function set_new_IDAI(address _new_IDAI) returns()
func (_Apr *AprTransactor) SetNewIDAI(opts *bind.TransactOpts, _new_IDAI common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_IDAI", _new_IDAI)
}

// SetNewIDAI is a paid mutator transaction binding the contract method 0x857c3c9b.
//
// Solidity: function set_new_IDAI(address _new_IDAI) returns()
func (_Apr *AprSession) SetNewIDAI(_new_IDAI common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewIDAI(&_Apr.TransactOpts, _new_IDAI)
}

// SetNewIDAI is a paid mutator transaction binding the contract method 0x857c3c9b.
//
// Solidity: function set_new_IDAI(address _new_IDAI) returns()
func (_Apr *AprTransactorSession) SetNewIDAI(_new_IDAI common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewIDAI(&_Apr.TransactOpts, _new_IDAI)
}

// SetNewIETH is a paid mutator transaction binding the contract method 0x429db17a.
//
// Solidity: function set_new_IETH(address _new_IETH) returns()
func (_Apr *AprTransactor) SetNewIETH(opts *bind.TransactOpts, _new_IETH common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_IETH", _new_IETH)
}

// SetNewIETH is a paid mutator transaction binding the contract method 0x429db17a.
//
// Solidity: function set_new_IETH(address _new_IETH) returns()
func (_Apr *AprSession) SetNewIETH(_new_IETH common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewIETH(&_Apr.TransactOpts, _new_IETH)
}

// SetNewIETH is a paid mutator transaction binding the contract method 0x429db17a.
//
// Solidity: function set_new_IETH(address _new_IETH) returns()
func (_Apr *AprTransactorSession) SetNewIETH(_new_IETH common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewIETH(&_Apr.TransactOpts, _new_IETH)
}

// SetNewIKNC is a paid mutator transaction binding the contract method 0xe9885aeb.
//
// Solidity: function set_new_IKNC(address _new_IKNC) returns()
func (_Apr *AprTransactor) SetNewIKNC(opts *bind.TransactOpts, _new_IKNC common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_IKNC", _new_IKNC)
}

// SetNewIKNC is a paid mutator transaction binding the contract method 0xe9885aeb.
//
// Solidity: function set_new_IKNC(address _new_IKNC) returns()
func (_Apr *AprSession) SetNewIKNC(_new_IKNC common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewIKNC(&_Apr.TransactOpts, _new_IKNC)
}

// SetNewIKNC is a paid mutator transaction binding the contract method 0xe9885aeb.
//
// Solidity: function set_new_IKNC(address _new_IKNC) returns()
func (_Apr *AprTransactorSession) SetNewIKNC(_new_IKNC common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewIKNC(&_Apr.TransactOpts, _new_IKNC)
}

// SetNewILINK is a paid mutator transaction binding the contract method 0x2735b363.
//
// Solidity: function set_new_ILINK(address _new_ILINK) returns()
func (_Apr *AprTransactor) SetNewILINK(opts *bind.TransactOpts, _new_ILINK common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_ILINK", _new_ILINK)
}

// SetNewILINK is a paid mutator transaction binding the contract method 0x2735b363.
//
// Solidity: function set_new_ILINK(address _new_ILINK) returns()
func (_Apr *AprSession) SetNewILINK(_new_ILINK common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewILINK(&_Apr.TransactOpts, _new_ILINK)
}

// SetNewILINK is a paid mutator transaction binding the contract method 0x2735b363.
//
// Solidity: function set_new_ILINK(address _new_ILINK) returns()
func (_Apr *AprTransactorSession) SetNewILINK(_new_ILINK common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewILINK(&_Apr.TransactOpts, _new_ILINK)
}

// SetNewIREP is a paid mutator transaction binding the contract method 0xb0ac4a73.
//
// Solidity: function set_new_IREP(address _new_IREP) returns()
func (_Apr *AprTransactor) SetNewIREP(opts *bind.TransactOpts, _new_IREP common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_IREP", _new_IREP)
}

// SetNewIREP is a paid mutator transaction binding the contract method 0xb0ac4a73.
//
// Solidity: function set_new_IREP(address _new_IREP) returns()
func (_Apr *AprSession) SetNewIREP(_new_IREP common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewIREP(&_Apr.TransactOpts, _new_IREP)
}

// SetNewIREP is a paid mutator transaction binding the contract method 0xb0ac4a73.
//
// Solidity: function set_new_IREP(address _new_IREP) returns()
func (_Apr *AprTransactorSession) SetNewIREP(_new_IREP common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewIREP(&_Apr.TransactOpts, _new_IREP)
}

// SetNewISAI is a paid mutator transaction binding the contract method 0xc552934d.
//
// Solidity: function set_new_ISAI(address _new_ISAI) returns()
func (_Apr *AprTransactor) SetNewISAI(opts *bind.TransactOpts, _new_ISAI common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_ISAI", _new_ISAI)
}

// SetNewISAI is a paid mutator transaction binding the contract method 0xc552934d.
//
// Solidity: function set_new_ISAI(address _new_ISAI) returns()
func (_Apr *AprSession) SetNewISAI(_new_ISAI common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewISAI(&_Apr.TransactOpts, _new_ISAI)
}

// SetNewISAI is a paid mutator transaction binding the contract method 0xc552934d.
//
// Solidity: function set_new_ISAI(address _new_ISAI) returns()
func (_Apr *AprTransactorSession) SetNewISAI(_new_ISAI common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewISAI(&_Apr.TransactOpts, _new_ISAI)
}

// SetNewISUSD is a paid mutator transaction binding the contract method 0x1b856311.
//
// Solidity: function set_new_ISUSD(address _new_ISUSD) returns()
func (_Apr *AprTransactor) SetNewISUSD(opts *bind.TransactOpts, _new_ISUSD common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_ISUSD", _new_ISUSD)
}

// SetNewISUSD is a paid mutator transaction binding the contract method 0x1b856311.
//
// Solidity: function set_new_ISUSD(address _new_ISUSD) returns()
func (_Apr *AprSession) SetNewISUSD(_new_ISUSD common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewISUSD(&_Apr.TransactOpts, _new_ISUSD)
}

// SetNewISUSD is a paid mutator transaction binding the contract method 0x1b856311.
//
// Solidity: function set_new_ISUSD(address _new_ISUSD) returns()
func (_Apr *AprTransactorSession) SetNewISUSD(_new_ISUSD common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewISUSD(&_Apr.TransactOpts, _new_ISUSD)
}

// SetNewIUSDC is a paid mutator transaction binding the contract method 0x750da899.
//
// Solidity: function set_new_IUSDC(address _new_IUSDC) returns()
func (_Apr *AprTransactor) SetNewIUSDC(opts *bind.TransactOpts, _new_IUSDC common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_IUSDC", _new_IUSDC)
}

// SetNewIUSDC is a paid mutator transaction binding the contract method 0x750da899.
//
// Solidity: function set_new_IUSDC(address _new_IUSDC) returns()
func (_Apr *AprSession) SetNewIUSDC(_new_IUSDC common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewIUSDC(&_Apr.TransactOpts, _new_IUSDC)
}

// SetNewIUSDC is a paid mutator transaction binding the contract method 0x750da899.
//
// Solidity: function set_new_IUSDC(address _new_IUSDC) returns()
func (_Apr *AprTransactorSession) SetNewIUSDC(_new_IUSDC common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewIUSDC(&_Apr.TransactOpts, _new_IUSDC)
}

// SetNewIWBTC is a paid mutator transaction binding the contract method 0xe49f09e4.
//
// Solidity: function set_new_IWBTC(address _new_IWBTC) returns()
func (_Apr *AprTransactor) SetNewIWBTC(opts *bind.TransactOpts, _new_IWBTC common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_IWBTC", _new_IWBTC)
}

// SetNewIWBTC is a paid mutator transaction binding the contract method 0xe49f09e4.
//
// Solidity: function set_new_IWBTC(address _new_IWBTC) returns()
func (_Apr *AprSession) SetNewIWBTC(_new_IWBTC common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewIWBTC(&_Apr.TransactOpts, _new_IWBTC)
}

// SetNewIWBTC is a paid mutator transaction binding the contract method 0xe49f09e4.
//
// Solidity: function set_new_IWBTC(address _new_IWBTC) returns()
func (_Apr *AprTransactorSession) SetNewIWBTC(_new_IWBTC common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewIWBTC(&_Apr.TransactOpts, _new_IWBTC)
}

// SetNewIZRX is a paid mutator transaction binding the contract method 0x38118ab4.
//
// Solidity: function set_new_IZRX(address _new_IZRX) returns()
func (_Apr *AprTransactor) SetNewIZRX(opts *bind.TransactOpts, _new_IZRX common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "set_new_IZRX", _new_IZRX)
}

// SetNewIZRX is a paid mutator transaction binding the contract method 0x38118ab4.
//
// Solidity: function set_new_IZRX(address _new_IZRX) returns()
func (_Apr *AprSession) SetNewIZRX(_new_IZRX common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewIZRX(&_Apr.TransactOpts, _new_IZRX)
}

// SetNewIZRX is a paid mutator transaction binding the contract method 0x38118ab4.
//
// Solidity: function set_new_IZRX(address _new_IZRX) returns()
func (_Apr *AprTransactorSession) SetNewIZRX(_new_IZRX common.Address) (*types.Transaction, error) {
	return _Apr.Contract.SetNewIZRX(&_Apr.TransactOpts, _new_IZRX)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Apr *AprTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Apr.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Apr *AprSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Apr.Contract.TransferOwnership(&_Apr.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Apr *AprTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Apr.Contract.TransferOwnership(&_Apr.TransactOpts, newOwner)
}

// AprOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Apr contract.
type AprOwnershipTransferredIterator struct {
	Event *AprOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AprOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AprOwnershipTransferred)
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
		it.Event = new(AprOwnershipTransferred)
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
func (it *AprOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AprOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AprOwnershipTransferred represents a OwnershipTransferred event raised by the Apr contract.
type AprOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Apr *AprFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AprOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Apr.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AprOwnershipTransferredIterator{contract: _Apr.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Apr *AprFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AprOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Apr.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AprOwnershipTransferred)
				if err := _Apr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Apr *AprFilterer) ParseOwnershipTransferred(log types.Log) (*AprOwnershipTransferred, error) {
	event := new(AprOwnershipTransferred)
	if err := _Apr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
