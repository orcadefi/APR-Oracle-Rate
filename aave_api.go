

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// To parse and unparse this JSON data, add this code to your project and do:
//
//    aave, err := UnmarshalAave(bytes)
//    bytes, err = aave.Marshal()
func UnmarshalAave(data []byte) (Aave, error) {
	var r Aave
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Aave) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Aave struct {
	Data Data `json:"data"`
}

type Data struct {
	Reserves []Reserve `json:"reserves"`
}

type Reserve struct {
	LiquidityRate      string `json:"liquidityRate"`
	Symbol             string `json:"symbol"`
	VariableBorrowRate string `json:"variableBorrowRate"`
}

func main() {
	petition := "{\"query\": \"{ reserves (where: {usageAsCollateralEnabled: true})  { symbol variableBorrowRate liquidityRate}}\" }"
	var petitiondata = []byte(petition)
	response, _ := http.Post("https://api.thegraph.com/subgraphs/name/aave/protocol-multy-raw","application/json",bytes.NewBuffer(petitiondata))
	body, _ := ioutil.ReadAll(response.Body)
	resp, _ := UnmarshalAave(body)
	fmt.Println("Aave % APR API")
	for _, token := range resp.Data.Reserves{
		fmt.Println("---------------------------")
		fmt.Println(token.Symbol)
		lendsecond, _ := strconv.ParseFloat(token.LiquidityRate,64)
		lendsecond = lendsecond / 10000000000000000000000000
		lendyear := fmt.Sprintf("%.10f", lendsecond)
		borrowsecond, _ := strconv.ParseFloat(token.VariableBorrowRate,64)
		borrowsecond = borrowsecond / 10000000000000000000000000
		borrowyear := fmt.Sprintf("%.10f", borrowsecond)
		fmt.Println("Lend Rate",lendyear,"%")
		fmt.Println("Borrow Rate: ",borrowyear,"%")
	}
}