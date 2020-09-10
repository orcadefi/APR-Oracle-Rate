package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	apr "./apr_oracle" // for demo
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/e15b0a66457d4800b2ee7f62d46f0b3e")
	if err != nil {
		log.Fatal(err)
	}

	// Iern address (APR ORACLE)
	address := common.HexToAddress("0x97FF4A1b787ADe6b94cca95b61F79417c673331D")
	instance, err := apr.NewApr(address, client)
	if err != nil {
		log.Fatal(err)
	}
	for true {
		aave, _ := instance.GetAllAaveAPR(nil)
		fmt.Println(aave.ADAI)
	}
	fmt.Println("contract is loaded")
	_ = instance
}