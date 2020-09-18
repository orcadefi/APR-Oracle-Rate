package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// To parse and unparse this JSON data, add this code to your project and do:
//
//    fulcrum, err := UnmarshalFulcrum(bytes)
//    bytes, err = fulcrum.Marshal()
func UnmarshalFulcrum(data []byte) (Fulcrum, error) {
	var r Fulcrum
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Fulcrum) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
//Store the response data
type Fulcrum struct {
	Data    Data `json:"data"`
	Success bool `json:"success"`
}
//Store each coin lend/borrow data
type Data struct {
	Eth   string `json:"eth"`
	Usdc  string `json:"usdc"`
	Wbtc  string `json:"wbtc"`
	Knc   string `json:"knc"`
	Link  string `json:"link"`
	Dai   string `json:"dai"`
	Usdt  string `json:"usdt"`
	Lend  string `json:"lend"`
	Mkr   string `json:"mkr"`
	Bzrx  string `json:"bzrx"`
	Yfi   string `json:"yfi"`
	Ethv1 string `json:"ethv1"`
}

func main() {
	//Request the lend and borrow data with a Post method to the Fulcrum API
	lend, _ := http.Get("https://api.bzx.network/v1/supply-rate-apr")
	borrow, _ := http.Get("https://api.bzx.network/v1/borrow-rate-apr")
	//Convert the response to readable format
	lendbody, _ := ioutil.ReadAll(lend.Body)
	borrowbody, _ := ioutil.ReadAll(borrow.Body)
	//Convert Json to Struct
	lendresp, _ := UnmarshalFulcrum(lendbody)
	borrowresp, _ := UnmarshalFulcrum(borrowbody)
	//Print the collected data
	fmt.Println("Fulcrum % APR API")
	fmt.Println("---------------------------")
	fmt.Println("ETH")
	fmt.Println("Lend Rate: ",lendresp.Data.Eth,"%")
	fmt.Println("Borrow Rate",borrowresp.Data.Eth,"%")
	fmt.Println("---------------------------")
	fmt.Println("Usdc")
	fmt.Println("Lend Rate: ",lendresp.Data.Usdc,"%")
	fmt.Println("Borrow Rate",borrowresp.Data.Usdc,"%")
	fmt.Println("---------------------------")
	fmt.Println("Wbtc")
	fmt.Println("Lend Rate: ",lendresp.Data.Wbtc,"%")
	fmt.Println("Borrow Rate",borrowresp.Data.Wbtc,"%")
	fmt.Println("---------------------------")
	fmt.Println("Knc")
	fmt.Println("Lend Rate: ",lendresp.Data.Knc,"%")
	fmt.Println("Borrow Rate",borrowresp.Data.Knc,"%")
	fmt.Println("---------------------------")
	fmt.Println("Link")
	fmt.Println("Lend Rate: ",lendresp.Data.Link,"%")
	fmt.Println("Borrow Rate",borrowresp.Data.Link,"%")
	fmt.Println("---------------------------")
	fmt.Println("Dai")
	fmt.Println("Lend Rate: ",lendresp.Data.Dai,"%")
	fmt.Println("Borrow Rate",borrowresp.Data.Dai,"%")
	fmt.Println("---------------------------")
	fmt.Println("Usdt")
	fmt.Println("Lend Rate: ",lendresp.Data.Usdt,"%")
	fmt.Println("Borrow Rate",borrowresp.Data.Usdt,"%")
	fmt.Println("---------------------------")
	fmt.Println("Lend")
	fmt.Println("Lend Rate: ",lendresp.Data.Lend,"%")
	fmt.Println("Borrow Rate",borrowresp.Data.Lend,"%")
	fmt.Println("---------------------------")
	fmt.Println("Mkr")
	fmt.Println("Lend Rate: ",lendresp.Data.Mkr,"%")
	fmt.Println("Borrow Rate",borrowresp.Data.Mkr,"%")
	fmt.Println("---------------------------")
	fmt.Println("Bzrx")
	fmt.Println("Lend Rate: ",lendresp.Data.Bzrx,"%")
	fmt.Println("Borrow Rate",borrowresp.Data.Bzrx,"%")
	fmt.Println("---------------------------")
	fmt.Println("Yfi")
	fmt.Println("Lend Rate: ",lendresp.Data.Yfi,"%")
	fmt.Println("Borrow Rate",borrowresp.Data.Yfi,"%")
	fmt.Println("---------------------------")
	fmt.Println("Ethv1")
	fmt.Println("Lend Rate: ",lendresp.Data.Ethv1,"%")
	fmt.Println("Borrow Rate",borrowresp.Data.Ethv1,"%")
	fmt.Println("---------------------------")
}