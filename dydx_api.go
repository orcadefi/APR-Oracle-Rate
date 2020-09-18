package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
	"strconv"
)

// To parse and unparse this JSON data, add this code to your project and do:
//
//    dydx, err := UnmarshalDydx(bytes)
//    bytes, err = dydx.Marshal()
func UnmarshalDydx(data []byte) (Dydx, error) {
	var r Dydx
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Dydx) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
//Store the data of the markets
type Dydx struct {
	Markets []Market `json:"markets"`
}
//Store the all market data from each coin
type Market struct {
	ID                        int64       `json:"id"`
	Name                      string      `json:"name"`
	Symbol                    string      `json:"symbol"`
	SupplyIndex               string      `json:"supplyIndex"`
	BorrowIndex               string      `json:"borrowIndex"`
	SupplyInterestRateSeconds string      `json:"supplyInterestRateSeconds"`
	BorrowInterestRateSeconds string      `json:"borrowInterestRateSeconds"`
	TotalSupplyPar            string      `json:"totalSupplyPar"`
	TotalBorrowPar            string      `json:"totalBorrowPar"`
	LastIndexUpdateSeconds    string      `json:"lastIndexUpdateSeconds"`
	OraclePrice               string      `json:"oraclePrice"`
	CollateralRatio           string      `json:"collateralRatio"`
	MarginPremium             string      `json:"marginPremium"`
	SpreadPremium             string      `json:"spreadPremium"`
	CurrencyUUID              string      `json:"currencyUuid"`
	CreatedAt                 string      `json:"createdAt"`
	UpdatedAt                 string      `json:"updatedAt"`
	DeletedAt                 interface{} `json:"deletedAt"`
	Currency                  Currency    `json:"currency"`
	TotalSupplyAPR            string      `json:"totalSupplyAPR"`
	TotalBorrowAPR            string      `json:"totalBorrowAPR"`
	TotalSupplyAPY            string      `json:"totalSupplyAPY"`
	TotalBorrowAPY            string      `json:"totalBorrowAPY"`
	TotalSupplyWei            string      `json:"totalSupplyWei"`
	TotalBorrowWei            string      `json:"totalBorrowWei"`
}
//Store the details of each coin in the platform
type Currency struct {
	UUID            string `json:"uuid"`
	Symbol          string `json:"symbol"`
	ContractAddress string `json:"contractAddress"`
	Decimals        int64  `json:"decimals"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

func main() {
	//Request the lend and borrow data with a Post method to the Dydx API
	response, _ := http.Get("https://api.dydx.exchange/v1/markets")
	//Convert the response to readable format
	body, _ := ioutil.ReadAll(response.Body)
	//Convert Json to Struct
	resp, _ := UnmarshalDydx(body)
	fmt.Println("Dydx % APR API")
	for _, token := range resp.Markets {
		fmt.Println("---------------------------")
		fmt.Println(token.Symbol)
		//Convert to float for generate decimal format
		lendsecond, _ := strconv.ParseFloat(token.SupplyInterestRateSeconds, 64)
		//Convert to decimal format (Lend)
		lendyear := lendsecond * 31536000 * 100
		//Convert to float for generate decimal format
		borrowsecond, _ := strconv.ParseFloat(token.BorrowInterestRateSeconds, 64)
		//Convert to decimal format (Borrow)
		borrowyear := borrowsecond * 31536000 * 100
		//Print data in % format
		fmt.Println("Lend Rate",lendyear,"%")
		fmt.Println("Borrow Rate: ",borrowyear,"%")
	}
}