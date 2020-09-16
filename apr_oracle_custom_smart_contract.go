package main

import (
	apr "./apr_oracle_custom"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	//Storing apr rates
	// Ethereum client connector
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/928905f152fe4546b37cc136336fda96")
	if err != nil {
		log.Fatal(err)
	}
	// Contract address (APR ORACLE)
	address := common.HexToAddress("")
	//Instance connection to Smart Contract
	instance, err := apr.NewAprCustom(address, client)
	if err != nil {
		log.Fatal(err)
	}

	for true {
		//Get normal pairs data
		fmt.Print(instance.GetLendAaveAPR(nil,common.HexToAddress(USDT)))
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress(USDT)))
		fmt.Println(instance.GetLendCompoundAPR(nil,common.HexToAddress(USDT)))
		fmt.Println(instance.GetBorrowAPR(nil,common.HexToAddress(USDT)))
	}
	//fmt.Println("contract is loaded")
	_ = instance
}