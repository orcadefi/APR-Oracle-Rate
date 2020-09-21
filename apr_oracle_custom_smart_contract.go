package main

import (
	apr "./aprOracle"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	//Storing apr rates
	// Ethereum client connector
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/928905f152fe4546b37cc136336fda96")
	if err != nil {
		log.Fatal(err)
	}
	// Contract address (APR ORACLE)
	address := common.HexToAddress("0xc087DbF9E4849F38C9a3a5EAc5684981FAdAe584")
	//Instance connection to Smart Contract
	instance, err := apr.NewAprOracleCaller(address, client)
	if err != nil {
		log.Fatal(err)
	}

	for true {
		//Get normal pairs data
		fmt.Println("Aave Platform Lend/Borrow")

		fmt.Println("ABAT")
		fmt.Println(instance.GetLendAaveAPR(nil,common.HexToAddress("0x0d8775f648430679a709e98d2b0cb6250d2887ef")))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0x0d8775f648430679a709e98d2b0cb6250d2887ef")))
		fmt.Println("ADAI")
		fmt.Println(instance.GetLendAaveAPR(nil,common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f")))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f")))
		fmt.Println("AETH")
		fmt.Println(instance.GetLendAaveAPR(nil,common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")))
		fmt.Println("AKNC")
		fmt.Println(instance.GetLendAaveAPR(nil,common.HexToAddress("0xdd974d5c2e2928dea5f71b9825b8b646686bd200")))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0xdd974d5c2e2928dea5f71b9825b8b646686bd200")))
		fmt.Println("ALEND")
		fmt.Println(instance.GetLendAaveAPR(nil,common.HexToAddress("0x80fb784b7ed66730e8b1dbd9820afd29931aab03")))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0x80fb784b7ed66730e8b1dbd9820afd29931aab03")))
		fmt.Println("ALINK")
		fmt.Println(instance.GetLendAaveAPR(nil,common.HexToAddress("0x514910771af9ca656af840dff83e8264ecf986ca")))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0x514910771af9ca656af840dff83e8264ecf986ca")))
		fmt.Println("AMANA")
		fmt.Println(instance.GetLendAaveAPR(nil,common.HexToAddress("0x0f5d2fb29fb7d3cfee444a200298f468908cc942")))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0x0f5d2fb29fb7d3cfee444a200298f468908cc942")))
		fmt.Println("AMKR")
		fmt.Println(instance.GetLendAaveAPR(nil,common.HexToAddress("0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2")))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2")))
		fmt.Println("AMKR")
		fmt.Println(instance.GetLendAaveAPR(nil,common.HexToAddress("0x1985365e9f78359a9b6ad760e32412f4a445e862")))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0x1985365e9f78359a9b6ad760e32412f4a445e862")))
		fmt.Println("ASNX")
		fmt.Println(instance.GetLendAaveAPR(nil,common.HexToAddress("0xc011a73ee8576fb46f5e1c5751ca3b9fe0af2a6f")))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0xc011a73ee8576fb46f5e1c5751ca3b9fe0af2a6f")))
		fmt.Println("ASUSD")
		fmt.Println(instance.GetLendAaveAPR(nil,common.HexToAddress("0x57ab1ec28d129707052df4df418d58a2d46d5f51")))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0x57ab1ec28d129707052df4df418d58a2d46d5f51")))
		fmt.Println("ATUSD")
		fmt.Println(instance.GetLendAaveAPR(nil,common.HexToAddress("0x0000000000085d4780b73119b644ae5ecd22b376")))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0x0000000000085d4780b73119b644ae5ecd22b376")))
		fmt.Println("AUSDC")
		fmt.Println(instance.GetLendAaveAPR(nil,common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")))
		fmt.Println("AUSDT")
		fmt.Println(instance.GetLendAaveAPR(nil,common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")))
		fmt.Println("AWBTC")
		fmt.Println(instance.GetLendAaveAPR(nil,common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599")))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599")))
		fmt.Println("AZRX")
		fmt.Println(instance.GetLendAaveAPR(nil,common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")))

		fmt.Println("Compound Platform Lend/Borrow")

		fmt.Println("CBAT")
		fmt.Println(instance.GetLendCompoundAPR(nil,common.HexToAddress("0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e")))
		fmt.Println(instance.GetBorrowCompoundAPR(nil,common.HexToAddress("0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e")))
		fmt.Println("CDAI")
		fmt.Println(instance.GetLendCompoundAPR(nil,common.HexToAddress("0x5d3a536e4d6dbd6114cc1ead35777bab948e3643")))
		fmt.Println(instance.GetBorrowCompoundAPR(nil,common.HexToAddress("0x5d3a536e4d6dbd6114cc1ead35777bab948e3643")))
		fmt.Println("CETH")
		fmt.Println(instance.GetLendCompoundAPR(nil,common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5")))
		fmt.Println(instance.GetBorrowCompoundAPR(nil,common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5")))
		fmt.Println("CREP")
		fmt.Println(instance.GetLendCompoundAPR(nil,common.HexToAddress("0x158079ee67fce2f58472a96584a73c7ab9ac95c1")))
		fmt.Println(instance.GetBorrowCompoundAPR(nil,common.HexToAddress("0x158079ee67fce2f58472a96584a73c7ab9ac95c1")))
		fmt.Println("CSAI")
		fmt.Println(instance.GetLendCompoundAPR(nil,common.HexToAddress("0xf5dce57282a584d2746faf1593d3121fcac444dc")))
		fmt.Println(instance.GetBorrowCompoundAPR(nil,common.HexToAddress("0xf5dce57282a584d2746faf1593d3121fcac444dc")))
		fmt.Println("CUSDC")
		fmt.Println(instance.GetLendCompoundAPR(nil,common.HexToAddress("0x39aa39c021dfbae8fac545936693ac917d5e7563")))
		fmt.Println(instance.GetBorrowCompoundAPR(nil,common.HexToAddress("0x39aa39c021dfbae8fac545936693ac917d5e7563")))
		fmt.Println("CWBTC")
		fmt.Println(instance.GetLendCompoundAPR(nil,common.HexToAddress("0xc11b1268c1a384e55c48c2391d8d480264a3a7f4")))
		fmt.Println(instance.GetBorrowCompoundAPR(nil,common.HexToAddress("0xc11b1268c1a384e55c48c2391d8d480264a3a7f4")))
		fmt.Println("CZRX")
		fmt.Println(instance.GetLendCompoundAPR(nil,common.HexToAddress("0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407")))
		fmt.Println(instance.GetBorrowCompoundAPR(nil,common.HexToAddress("0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407")))

		fmt.Println("Dydx Platform Lend/Borrow")

		fmt.Println("0")
		fmt.Println(instance.GetLendDyDxAPR(nil,big.NewInt(0) ))
		fmt.Println(instance.GetBorrowDyDxAPR(nil,big.NewInt(0)))
		fmt.Println("2")
		fmt.Println(instance.GetLendDyDxAPR(nil,big.NewInt(0) ))
		fmt.Println(instance.GetBorrowDyDxAPR(nil,big.NewInt(0)))
		fmt.Println("3")
		fmt.Println(instance.GetLendDyDxAPR(nil,big.NewInt(0) ))
		fmt.Println(instance.GetBorrowDyDxAPR(nil,big.NewInt(0)))

		fmt.Println("Fulcrum Platform Lend/Borrow")

		fmt.Println("iETH")
		fmt.Println(instance.GetLendFulcrumAPR(nil,common.HexToAddress("0xB983E01458529665007fF7E0CDdeCDB74B967Eb6") ))
		fmt.Println(instance.GetBorrowFulcrumAPR(nil,common.HexToAddress("0xB983E01458529665007fF7E0CDdeCDB74B967Eb6")))
		fmt.Println("iDAI")
		fmt.Println(instance.GetLendFulcrumAPR(nil,common.HexToAddress("0x6b093998D36f2C7F0cc359441FBB24CC629D5FF0") ))
		fmt.Println(instance.GetBorrowFulcrumAPR(nil,common.HexToAddress("0x6b093998D36f2C7F0cc359441FBB24CC629D5FF0")))
		fmt.Println("iUSDC")
		fmt.Println(instance.GetLendFulcrumAPR(nil,common.HexToAddress("0x32E4c68B3A4a813b710595AebA7f6B7604Ab9c15") ))
		fmt.Println(instance.GetBorrowFulcrumAPR(nil,common.HexToAddress("0x32E4c68B3A4a813b710595AebA7f6B7604Ab9c15")))
		fmt.Println("iUSDT")
		fmt.Println(instance.GetLendFulcrumAPR(nil,common.HexToAddress("0x7e9997a38A439b2be7ed9c9C4628391d3e055D48") ))
		fmt.Println(instance.GetBorrowFulcrumAPR(nil,common.HexToAddress("0x7e9997a38A439b2be7ed9c9C4628391d3e055D48")))
		fmt.Println("iWBTC")
		fmt.Println(instance.GetLendFulcrumAPR(nil,common.HexToAddress("0x2ffa85f655752fB2aCB210287c60b9ef335f5b6E") ))
		fmt.Println(instance.GetBorrowFulcrumAPR(nil,common.HexToAddress("0x2ffa85f655752fB2aCB210287c60b9ef335f5b6E")))
		fmt.Println("iLINK")
		fmt.Println(instance.GetLendFulcrumAPR(nil,common.HexToAddress("0x463538705E7d22aA7f03Ebf8ab09B067e1001B54") ))
		fmt.Println(instance.GetBorrowFulcrumAPR(nil,common.HexToAddress("0x463538705E7d22aA7f03Ebf8ab09B067e1001B54")))
		fmt.Println("iYFI")
		fmt.Println(instance.GetLendFulcrumAPR(nil,common.HexToAddress("0x7F3Fe9D492A9a60aEBb06d82cBa23c6F32CAd10b") ))
		fmt.Println(instance.GetBorrowFulcrumAPR(nil,common.HexToAddress("0x7F3Fe9D492A9a60aEBb06d82cBa23c6F32CAd10b")))
		fmt.Println("bzrx")
		fmt.Println(instance.GetLendFulcrumAPR(nil,common.HexToAddress("0x18240BD9C07fA6156Ce3F3f61921cC82b2619157") ))
		fmt.Println(instance.GetBorrowFulcrumAPR(nil,common.HexToAddress("0x18240BD9C07fA6156Ce3F3f61921cC82b2619157")))
		fmt.Println("iMKR")
		fmt.Println(instance.GetLendFulcrumAPR(nil,common.HexToAddress("0x9189c499727f88F8eCC7dC4EEA22c828E6AaC015") ))
		fmt.Println(instance.GetBorrowFulcrumAPR(nil,common.HexToAddress("0x9189c499727f88F8eCC7dC4EEA22c828E6AaC015")))
		fmt.Println("iLEND")
		fmt.Println(instance.GetLendFulcrumAPR(nil,common.HexToAddress("0x9189c499727f88F8eCC7dC4EEA22c828E6AaC015") ))
		fmt.Println(instance.GetBorrowFulcrumAPR(nil,common.HexToAddress("0x9189c499727f88F8eCC7dC4EEA22c828E6AaC015")))
		fmt.Println("iKNC")
		fmt.Println(instance.GetLendFulcrumAPR(nil,common.HexToAddress("0x687642347a9282Be8FD809d8309910A3f984Ac5a") ))
		fmt.Println(instance.GetBorrowFulcrumAPR(nil,common.HexToAddress("0x687642347a9282Be8FD809d8309910A3f984Ac5a")))


		break
	}
	//fmt.Print("contract is loaded")
	_ = instance
}