package parser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

const INPUTFILE string = "camt053.xml"

type BankStatements struct {
	Statements []Statement `xml:"BkToCstmrStmt>Stmt"`
}

type Statement struct {
	AccountInfo  Account   `xml:"Acct"`
	Balances     []Balance `xml:"Bal"`
	Transactions []Entry   `xml:"Ntry"`
}

type Account struct {
	AccountId string `xml:"Id>Othr>Id"`
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
	Reference string `xml:"NtryRef"`
	Amount    struct {
		Value    float64 `xml:",chardata"`
		Currency string  `xml:"Ccy,attr"`
	} `xml:"Amt"`
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

func readFile() []byte {
	camt053File, err := os.Open(INPUTFILE)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer camt053File.Close()

	fmt.Println("Successfully Opened " + INPUTFILE)

	byteValue, _ := ioutil.ReadAll(camt053File)
	return byteValue
}

func Parse() BankStatements {
	var statements BankStatements
	var fileData []byte = readFile()

	err := xml.Unmarshal(fileData, &statements)
	if err != nil {
		fmt.Println("xml unmarshal error: ", err)
	}

	return statements
}
