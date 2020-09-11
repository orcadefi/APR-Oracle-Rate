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

type Dydx struct {
	Markets []Market `json:"markets"`
}

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

type Currency struct {
	UUID            string `json:"uuid"`
	Symbol          string `json:"symbol"`
	ContractAddress string `json:"contractAddress"`
	Decimals        int64  `json:"decimals"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

func main() {
	response, _ := http.Get("https://api.dydx.exchange/v1/markets")
	body, _ := ioutil.ReadAll(response.Body)
	resp, _ := UnmarshalDydx(body)
	fmt.Println("Dydx % APR API")
	for _, token := range resp.Markets {
		fmt.Println("---------------------------")
		fmt.Println(token.Symbol)
		lendsecond, _ := strconv.ParseFloat(token.SupplyInterestRateSeconds, 64)
		lendyear := lendsecond * 31536000 * 100
		borrowsecond, _ := strconv.ParseFloat(token.BorrowInterestRateSeconds, 64)
		borrowyear := borrowsecond * 31536000 * 100
		fmt.Println("Lend Rate",lendyear,"%")
		fmt.Println("Borrow Rate: ",borrowyear,"%")
	}
}