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
//Store all the response API data
type Compound struct {
	CToken  []CToken    `json:"cToken"`
	Error   interface{} `json:"error"`
	Meta    interface{} `json:"meta"`
	Request Request     `json:"request"`
}
//Store all the CToken data (lend/borrow)
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
//Store the value of Borrow/Lend data
type BorrowRate struct {
	Value string `json:"value"`
}
//Store the request data
type Request struct {
	Addresses      []interface{} `json:"addresses"`
	BlockNumber    int64         `json:"block_number"`
	BlockTimestamp int64         `json:"block_timestamp"`
	Meta           bool          `json:"meta"`
	Network        string        `json:"network"`
}

func main() {
	//Request the lend and borrow data with a Post method to the Compound API
	response, _ := http.Get("https://api.compound.finance/api/v2/ctoken")
	//Convert the response to readable format
	body, _ := ioutil.ReadAll(response.Body)
	//Convert Json to Struct
	resp, _ := UnmarshalCompound(body)
	//Print the collected data
	fmt.Println("Compound % APR API")
	for _, token := range resp.CToken {
		fmt.Println("---------------------------")
		fmt.Println(token.Symbol)
		//Convert to float for generate decimal format
		lendRate, _ := strconv.ParseFloat(token.SupplyRate.Value, 64)
		borrowRate, _ := strconv.ParseFloat(token.BorrowRate.Value, 64)
		//Convert to decimal format (Lend/Borrow)
		lendRate *= 100
		borrowRate *= 100
		//Print data in % format
		fmt.Println("Lend Rate",lendRate,"%")
		fmt.Println("Borrow Rate: ",borrowRate,"%")
	}
}