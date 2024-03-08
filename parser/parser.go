package parser

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

const INPUTFILENAME string = "camt053.xml"
const OUTPUTFILENAME string = "output.json"

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

func Parse() []byte {
	camt053File, err := os.Open(INPUTFILENAME)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer camt053File.Close()

	fmt.Println("Successfully Opened " + INPUTFILENAME)

	byteValue, _ := ioutil.ReadAll(camt053File)
	var statement Statement
	xml.Unmarshal(byteValue, &statement)

	jsonData, err := json.Marshal(statement)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil
	}

	return jsonData
}
