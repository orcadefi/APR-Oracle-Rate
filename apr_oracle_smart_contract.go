package main

import (
	apr "./apr_oracle" // for demo
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

const (
	//Aave extended Smart Contract Addresses (Not Working in the smart contract)
	AWBTC = "0xFC4B8ED459e00e5400be803A9BB3954234FD50e3"
	AMKR = "0x7deB5e830be29F91E298ba5FF1356BB7f8146998"
	ABUSD = "0x6Ee0f7BB50a54AB5253dA0667B0Dc2ee526C30a8"
	AENJ = "0x712DB54daA836B53Ef1EcBb9c6ba3b9Efb073F40"
	AREN = "0x69948cC03f478B95283F7dbf1CE764d0fc7EC54C"
	//Aave extended Smart Contract Addresses Optionals (Not Working in the smart contract)
	AYFI = "0x12e51E77DAAA58aA0E9247db7510Ea4B46F9bEAd"
	ALEND = "0x7D2D3688Df45Ce7C552E19c27e007673da9204B8"
	//Compound extended Smart Contract Addresses
	CUSDT = "0xf650c3d88d12db855b8bf7d11be6c55a4e07dcc9"
	CZRX = "0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407"
	//Fulcrum extended Smart Contract Addresses
	IUSDT = "0x7e9997a38a439b2be7ed9c9c4628391d3e055d48"
)

//APR rate type
type Apr struct{
	Lending		Platform
	Borrow		Platform
}
//List of platforms
//Extended: is list of pairs that are not included in smart contract
type Platform struct{
	Aave				AaveAprList
	AaveExtended		AaveAprListExtended
	Compound			CompoundAprList
	CompoundExtended	CompoundAprListExtended
	Dydx				DydxAprList
	Fulcrum				FulcrumAprList
	FulcrumExtended		FulcrumAprListExtended
}

/*
Pair Types:
Used: A pair in DeFI rate list and in the smart contract.
Not Used: A pair different of the DeFI rate list.
Needed: A pair that the smart contract doesn´t have but is in the DeFI rate list.
Optional: A pair that is in the platform but isn´t in the DeFI rate list and smart contract.
 */

type AaveAprList struct {
	ADAI  	*big.Int //Used
	ATUSD 	*big.Int //Used
	AUSDC 	*big.Int //Used
	AUSDT 	*big.Int //Used
	ASUSD 	*big.Int //Used
	ABAT  	*big.Int //Used
	AETH  	*big.Int //Used
	ALINK 	*big.Int //Used
	AKNC  	*big.Int //Used
	AREP  	*big.Int //Used
	AZRX  	*big.Int //Used
	ASNX  	*big.Int //Used
}

type AaveAprListExtended struct {
	AWBTC	*big.Int //Needed
	AMKR	*big.Int //Needed
	ABUSD	*big.Int //Needed
	AENJ	*big.Int //Needed
	AREN	*big.Int //Needed
	//Optionals
	AYFI    *big.Int //Optional
	ALEND   *big.Int //Optional
}

type CompoundAprList struct {
	CDAI  	*big.Int //Used
	CBAT  	*big.Int //Used
	CETH  	*big.Int //Used
	CREP  	*big.Int //No used
	CSAI  	*big.Int //No used
	CUSDC 	*big.Int //Used
	CWBTC 	*big.Int //Used
	CZRC  	*big.Int //No used (Deprecated)
}

type CompoundAprListExtended struct {
	CUSDT	*big.Int //Needed
	CZRX	*big.Int //Needed
}

type DydxAprList struct {
	DSAI  	*big.Int //Used
	DETH  	*big.Int //Used
	DUSDC 	*big.Int //Used
	DDAI  	*big.Int //Used
}

type FulcrumAprList struct {
	IZRX  	*big.Int //No used
	IREP  	*big.Int //No used
	IKNC  	*big.Int //Used
	IWBTC 	*big.Int //No used
	IUSDC 	*big.Int //Used
	IETH  	*big.Int //Used
	ISAI  	*big.Int //No used
	IDAI  	*big.Int //Used
	ILINK 	*big.Int //Used
	ISUSD 	*big.Int //No used

}

type FulcrumAprListExtended struct {
	IUSDT	*big.Int //Needed
}

func main() {
	//Storing apr rates
	aprStore := Apr{}
	// Ethereum client connector
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/928905f152fe4546b37cc136336fda96")
	if err != nil {
		log.Fatal(err)
	}
	// Contract address (APR ORACLE)
	address := common.HexToAddress("0x97FF4A1b787ADe6b94cca95b61F79417c673331D")
	//Instance connection to Smart Contract
	instance, err := apr.NewApr(address, client)
	if err != nil {
		log.Fatal(err)
	}

	for true {
		//Get normal pairs data
		aprStore.Lending.Aave, _ = instance.GetAllAaveAPR(nil)
		aprStore.Lending.Compound, _ = instance.GetAllCompoundAPR(nil)
		aprStore.Lending.Dydx, _ = instance.GetAllDyDxAPR(nil)
		aprStore.Lending.Fulcrum, _ = instance.GetAllFulcrumAPR(nil)
		//Get extended pairs data
		aprStore.Lending.CompoundExtended.CUSDT, _ = instance.GetCompoundAPR(nil,common.HexToAddress(CUSDT))
		aprStore.Lending.CompoundExtended.CZRX, _ = instance.GetCompoundAPR(nil,common.HexToAddress(CZRX))
		aprStore.Lending.FulcrumExtended.IUSDT, _ = instance.GetFulcrumAPR(nil,common.HexToAddress(IUSDT))

		//Not Working in SmartContract
		aprStore.Lending.AaveExtended.AWBTC, _ = instance.GetAaveAPR(nil,common.HexToAddress(AWBTC))
		aprStore.Lending.AaveExtended.AMKR, _ = instance.GetAaveAPR(nil,common.HexToAddress(AMKR))
		aprStore.Lending.AaveExtended.ABUSD, _ = instance.GetAaveAPR(nil,common.HexToAddress(ABUSD))
		aprStore.Lending.AaveExtended.AENJ, _ = instance.GetAaveAPR(nil,common.HexToAddress(AENJ))
		aprStore.Lending.AaveExtended.AREN, _ = instance.GetAaveAPR(nil,common.HexToAddress(AREN))


		//Print retrieved data
		fmt.Println("% APR Lending Rate")
		fmt.Println("Aave:")
		fmt.Println("ABAT:",aprStore.Lending.Aave.ABAT)
		fmt.Println("ADAI:",aprStore.Lending.Aave.ADAI)
		fmt.Println("AETH:",aprStore.Lending.Aave.AETH)
		fmt.Println("AKNC:",aprStore.Lending.Aave.AKNC)
		fmt.Println("ALINK:",aprStore.Lending.Aave.ALINK)
		fmt.Println("AREP:",aprStore.Lending.Aave.AREP)
		fmt.Println("ASNX:",aprStore.Lending.Aave.ASNX)
		fmt.Println("ASUSD:",aprStore.Lending.Aave.ASUSD)
		fmt.Println("ATUSD:",aprStore.Lending.Aave.ATUSD)
		fmt.Println("AUSDC:",aprStore.Lending.Aave.AUSDC)
		fmt.Println("AUSDT:",aprStore.Lending.Aave.AUSDT)
		fmt.Println("AZRX:",aprStore.Lending.Aave.AZRX)
		fmt.Println("Aave Extended:")
		fmt.Println("AWBTC:",aprStore.Lending.AaveExtended.AWBTC)
		fmt.Println("AMKR:",aprStore.Lending.AaveExtended.AMKR)
		fmt.Println("ABUSD:",aprStore.Lending.AaveExtended.ABUSD)
		fmt.Println("AENJ:",aprStore.Lending.AaveExtended.AENJ)
		fmt.Println("AREN:",aprStore.Lending.AaveExtended.AREN)
		fmt.Println("Compound:")
		fmt.Println("CBAT:",aprStore.Lending.Compound.CBAT)
		fmt.Println("CDAI:",aprStore.Lending.Compound.CDAI)
		fmt.Println("CETH:",aprStore.Lending.Compound.CETH)
		fmt.Println("CUSDC:",aprStore.Lending.Compound.CUSDC)
		fmt.Println("CWBTC:",aprStore.Lending.Compound.CWBTC)
		fmt.Println("Compound Extended:")
		fmt.Println("CUSDT",aprStore.Lending.CompoundExtended.CUSDT)
		fmt.Println("CZRX", aprStore.Lending.CompoundExtended.CZRX)
		fmt.Println("DyDx:")
		fmt.Println("DDAI:",aprStore.Lending.Dydx.DDAI)
		fmt.Println("DETH:",aprStore.Lending.Dydx.DETH)
		fmt.Println("DSAI:",aprStore.Lending.Dydx.DSAI)
		fmt.Println("DUSDC:",aprStore.Lending.Dydx.DUSDC)
		fmt.Println("Fulcrum:")
		fmt.Println("IKNC:",aprStore.Lending.Fulcrum.IKNC)
		fmt.Println("IUSDC:",aprStore.Lending.Fulcrum.IUSDC)
		fmt.Println("IETH:",aprStore.Lending.Fulcrum.IETH)
		fmt.Println("IDAI:",aprStore.Lending.Fulcrum.IDAI)
		fmt.Println("ILINK:",aprStore.Lending.Fulcrum.ILINK)
		fmt.Println("Fulcrum Extended:")
		fmt.Println("IUSDT:",aprStore.Lending.FulcrumExtended.IUSDT)

		time.Sleep(60*time.Second)
	}
	//fmt.Println("contract is loaded")
	_ = instance
}

