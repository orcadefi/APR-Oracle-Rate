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
	address := common.HexToAddress("0xc08...1FAdAe584")
	//Instance connection to Smart Contract
	instance, err := apr.NewAprCustom(address, client)
	if err != nil {
		log.Fatal(err)
	}

	for true {
		//Get normal pairs data
		fmt.Println(instance.GetBorrowAaveAPR(nil,common.HexToAddress("0x71010a9d003445ac60c4e6a7017c1e89a477b438")))
	}
	//fmt.Println("contract is loaded")
	_ = instance
}