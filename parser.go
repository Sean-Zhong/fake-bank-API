package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type Statement struct {
	//XMLName     xml.Name `xml:"statement"`
	AccountInfo Account   `xml:"BkToCstmrStmt>Stmt>Acct"`
	Balances    []Balance `xml:"BkToCstmrStmt>Stmt>Bal"`
	//Entries []Entry `xml:"Ntry"`
}

type Account struct {
	//XMLName xml.Name `xml:"Acct"`
	AccountId int    `xml:"Id>Othr>Id"`
	Currency  string `xml:"Ccy"`
	Owner     string `xml:"Ownr>Nm"`
}

type Balance struct {
	BalanceType string `xml:"Tp>CdOrPrtry>Cd"`
	Amount      struct {
		Value    float64 `xml:",chardata"`
		Currency string  `xml:"Ccy,attr"`
	} `xml:"Amt"`
}

func main() {
	camt053File, err := os.Open("camt053.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer camt053File.Close()

	fmt.Println("Successfully Opened camt053.xml")

	byteValue, _ := ioutil.ReadAll(camt053File)
	var statement Statement
	xml.Unmarshal(byteValue, &statement)
	fmt.Println(statement)
	fmt.Println(reflect.TypeOf(statement))
}
