package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// To parse and unparse this JSON data, add this code to your project and do:
//
//    compound, err := UnmarshalCompound(bytes)
//    bytes, err = compound.Marshal()
func UnmarshalCompound(data []byte) (Compound, error) {
	var r Compound
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Compound) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Compound struct {
	CToken  []CToken    `json:"cToken"`
	Error   interface{} `json:"error"`
	Meta    interface{} `json:"meta"`
	Request Request     `json:"request"`
}

type CToken struct {
	BorrowRate               BorrowRate `json:"borrow_rate"`
	Cash                     BorrowRate `json:"cash"`
	CollateralFactor         BorrowRate `json:"collateral_factor"`
	CompBorrowApy            BorrowRate `json:"comp_borrow_apy"`
	CompSupplyApy            BorrowRate `json:"comp_supply_apy"`
	ExchangeRate             BorrowRate `json:"exchange_rate"`
	InterestRateModelAddress string     `json:"interest_rate_model_address"`
	Name                     string     `json:"name"`
	NumberOfBorrowers        int64      `json:"number_of_borrowers"`
	NumberOfSuppliers        int64      `json:"number_of_suppliers"`
	ReserveFactor            BorrowRate `json:"reserve_factor"`
	Reserves                 BorrowRate `json:"reserves"`
	SupplyRate               BorrowRate `json:"supply_rate"`
	Symbol                   string     `json:"symbol"`
	TokenAddress             string     `json:"token_address"`
	TotalBorrows             BorrowRate `json:"total_borrows"`
	TotalSupply              BorrowRate `json:"total_supply"`
	UnderlyingAddress        *string    `json:"underlying_address"`
	UnderlyingName           string     `json:"underlying_name"`
	UnderlyingPrice          BorrowRate `json:"underlying_price"`
	UnderlyingSymbol         string     `json:"underlying_symbol"`
}

type BorrowRate struct {
	Value string `json:"value"`
}

type Request struct {
	Addresses      []interface{} `json:"addresses"`
	BlockNumber    int64         `json:"block_number"`
	BlockTimestamp int64         `json:"block_timestamp"`
	Meta           bool          `json:"meta"`
	Network        string        `json:"network"`
}

func main() {
	response, _ := http.Get("https://api.compound.finance/api/v2/ctoken")
	body, _ := ioutil.ReadAll(response.Body)
	resp, _ := UnmarshalCompound(body)
	fmt.Println("Compound % APR API")
	for _, token := range resp.CToken {
		fmt.Println("---------------------------")
		fmt.Println(token.Symbol)
		lendRate, _ := strconv.ParseFloat(token.SupplyRate.Value, 64)
		borrowRate, _ := strconv.ParseFloat(token.BorrowRate.Value, 64)
		lendRate *= 100
		borrowRate *= 100
		fmt.Println("Lend Rate",lendRate,"%")
		fmt.Println("Borrow Rate: ",borrowRate,"%")
	}
}