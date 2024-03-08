package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Statement struct {
	//XMLName     xml.Name `xml:"statement"`
	AccountInfo Account   `xml:"BkToCstmrStmt>Stmt>Acct"`
	Balances    []Balance `xml:"BkToCstmrStmt>Stmt>Bal"`
	Entries     []Entry   `xml:"BkToCstmrStmt>Stmt>Ntry"`
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

type Entry struct {
	EntryDetails []RemitanceInformation `xml:"NtryDtls>TxDtls>RmtInf"`
}

type RemitanceInformation struct {
	Structured struct {
		RmtdAmount struct {
			Value    float64 `xml:",chardata"`
			Currency string  `xml:"Ccy,attr"`
		} `xml:"RfrdDocAmt>RmtdAmt"`
		CdOrPrtry string `xml:"CdtrRefInf>Tp>CdOrPrtry>Cd"`
		Ref       int    `xml:"CdtrRefInf>Ref"`
	} `xml:"Strd"`
	Unstructured struct {
		Info string `xml:",chardata"`
	} `xml:"Ustrd"`
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

	jsonData, err := json.Marshal(statement)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Open a file for writing, create it if it doesn't exist, truncate it if it does
	file, err := os.OpenFile("output.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Make sure to close the file when done

	// Write data to the file
	_, err = file.WriteString(string(jsonData))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data has been written to output.json")
}
