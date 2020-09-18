

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

//Store the request data
type Aave struct {
	Data Data `json:"data"`
}
//Store each coin data
type Data struct {
	Reserves []Reserve `json:"reserves"`
}
//Store rates of the coin
type Reserve struct {
	LiquidityRate      string `json:"liquidityRate"`
	Symbol             string `json:"symbol"`
	VariableBorrowRate string `json:"variableBorrowRate"`
}

func main() {
	// GraphQL query of Aave API data
	petition := "{\"query\": \"{ reserves (where: {usageAsCollateralEnabled: true})  { symbol variableBorrowRate liquidityRate}}\" }"
	// GraphQL query in byte form to Post.
	var petitiondata = []byte(petition)
	//Request the query data with a Post method to the Aave API
	response, _ := http.Post("https://api.thegraph.com/subgraphs/name/aave/protocol-multy-raw","application/json",bytes.NewBuffer(petitiondata))
	//Convert the response to readable format
	body, _ := ioutil.ReadAll(response.Body)
	//Convert Json to Struct
	resp, _ := UnmarshalAave(body)
	//Print the collected data
	fmt.Println("Aave % APR API")
	for _, token := range resp.Data.Reserves{
		fmt.Println("---------------------------")
		fmt.Println(token.Symbol)
		//Convert to float for generate decimal format
		lendsecond, _ := strconv.ParseFloat(token.LiquidityRate,64)
		//Convert to decimal format (Lend)
		lendsecond = lendsecond / 10000000000000000000000000
		lendyear := fmt.Sprintf("%.10f", lendsecond)
		//Convert to float for generate decimal format
		borrowsecond, _ := strconv.ParseFloat(token.VariableBorrowRate,64)
		//Convert to decimal format (Borrow)
		borrowsecond = borrowsecond / 10000000000000000000000000
		borrowyear := fmt.Sprintf("%.10f", borrowsecond)
		//Print data in % format
		fmt.Println("Lend Rate",lendyear,"%")
		fmt.Println("Borrow Rate: ",borrowyear,"%")
	}
}