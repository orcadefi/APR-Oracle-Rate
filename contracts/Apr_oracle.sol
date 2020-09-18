pragma solidity ^0.5.0;
library SafeMath {
    //Safe multiplication method
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        if (a == 0) {
            return 0;
        }
        uint256 c = a * b;
        require(c / a == b, "overflow");
        return c;
    }
    //Safe division method
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        return div(a, b, "SafeMath: division by zero");
    }
    function div(uint256 a, uint256 b, string memory errorMessage) internal pure returns (uint256) {
        // Solidity only automatically asserts when dividing by 0
        require(b > 0, errorMessage);
        uint256 c = a / b;
        // assert(a == b * c + a % b); // There is no case in which this doesn't hold

        return c;
    }
}
pragma solidity >=0.4.21 <0.7.0;
pragma experimental ABIEncoderV2;
//Interfaces to retrieve data from the smart contracts of each platform
interface Compound {
    function supplyRatePerBlock() external view returns (uint);
    function borrowRatePerBlock() external view returns (uint);
}
interface DyDx {
    struct val {
        uint256 value;
    }

    struct set {
        uint128 borrow;
        uint128 supply;
    }

    function getEarningsRate() external view returns (val memory);
    function getMarketInterestRate(uint256 marketId) external view returns (val memory);
    function getMarketTotalPar(uint256 marketId) external view returns (set memory);
}
interface Fulcrum {
    function supplyInterestRate() external view returns (uint256);
    function borrowInterestRate() external view returns (uint256);
}
interface LendingPoolAddressesProvider {
    function getLendingPoolCore() external view returns (address);
}

interface LendingPoolCore  {
    function getReserveCurrentLiquidityRate(address _reserve) external view returns (uint256 liquidityRate);
    function getReserveCurrentVariableBorrowRate(address _reserve) external view returns (uint256 variableBorrowRate);
}
contract APROracle {
    //Importing SafeMath library for safe arithmetical operations
    using SafeMath for uint256;
    uint256 DECIMAL = 10 ** 18;
    address public owner;
    address public DYDX;
    address public AAVE;
    constructor() public {
        owner = msg.sender;
        //Dydx on-chain data address
        DYDX = address(0x1E0447b19BB6EcFdAe1e4AE1694b0C3659614e4e);
        //Aave on-chain data address
        AAVE = address(0x24a42fD28C976A61Df5D00D0599C34c4f90748c8);
    }
    //Get lend APR data from Compound from Compound contract address
    function getLendCompoundAPR(address token) public view returns (uint256) {
        return Compound(token).supplyRatePerBlock().mul(2102400);
    }
    //Get borrow APR data from Compound from Compound contract address
    function getBorrowCompoundAPR(address token) public view returns (uint256) {
        return Compound(token).borrowRatePerBlock().mul(2102400);
    }
    //Get lend APR data from Fulcrum from Fulcrum contract address
    function getLendFulcrumAPR(address token) public view returns(uint256) {
        return Fulcrum(token).supplyInterestRate().div(100);
    }
    //Get borrow APR data from Fulcrum from Fulcrum contract address
    function getBorrowFulcrumAPR(address token) public view returns(uint256) {
        return Fulcrum(token).borrowInterestRate().div(100);
    }
    //Get lend APR data from Dydx from Dydx contract address
    function getLendDyDxAPR(uint256 marketId) public view returns(uint256) {
        uint256 rate      = DyDx(DYDX).getMarketInterestRate(marketId).value;
        uint256 aprBorrow = rate * 31622400;
        uint256 borrow    = DyDx(DYDX).getMarketTotalPar(marketId).borrow;
        uint256 supply    = DyDx(DYDX).getMarketTotalPar(marketId).supply;
        uint256 usage     = (borrow * DECIMAL) / supply;
        uint256 apr       = (((aprBorrow * usage) / DECIMAL) * DyDx(DYDX).getEarningsRate().value) / DECIMAL;
        return apr;
    }
    //Get borrow APR data from Dydx from Dydx contract address
    function getBorrowDyDxAPR(uint256 marketId) public view returns(uint256) {
        uint256 rate      = DyDx(DYDX).getMarketInterestRate(marketId).value;
        uint256 aprBorrow = rate * 31622400;
        return aprBorrow;
    }
    //Get lend APR data from Aave from Aave contract address
    function getLendAaveAPR(address token) public view returns (uint256) {
        LendingPoolCore core = LendingPoolCore(LendingPoolAddressesProvider(AAVE).getLendingPoolCore());
        return core.getReserveCurrentLiquidityRate(token).div(1e9);
    }
    //Get borrow APR data from Aave from Aave contract address
    function getBorrowAaveAPR(address token) public view returns (uint256) {
        LendingPoolCore core = LendingPoolCore(LendingPoolAddressesProvider(AAVE).getLendingPoolCore());
        return core.getReserveCurrentVariableBorrowRate(token).div(1e9);
    }
}