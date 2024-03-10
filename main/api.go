package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Sean-Zhong/fake-bank-API/parser"
)

const PORT string = ":8081"

var statements parser.BankStatements = parser.Parse()

type AccountDetails struct {
	Account parser.Account
	Balance []parser.Balance
}

func convertToJson(input interface{}) []byte {
	var jsonData []byte
	jsonData, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil
	}
	return jsonData
}

func getStatementByAccountId(accountId int) parser.Statement {
	var stmt parser.Statement
	for _, elem := range statements.Statements {
		if elem.AccountInfo.AccountId == accountId {
			stmt = elem
		}
	}
	return stmt
}

func listAccounts() []byte {
	var accountList []parser.Account
	for _, statement := range statements.Statements {
		accountList = append(accountList, statement.AccountInfo)
	}
	return convertToJson(accountList)
}

func getAccount(accountId int) []byte {
	var stmt parser.Statement = getStatementByAccountId(accountId)
	var accountData = AccountDetails{Account: stmt.AccountInfo, Balance: stmt.Balances}

	// error handling needed
	return convertToJson(accountData)
}

func listTransactions(accountId int) []byte {
	var stmt parser.Statement = getStatementByAccountId(accountId)
	return convertToJson(stmt.Transactions)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func main() {
	fmt.Println(string(listAccounts()))
	fmt.Println(string(getAccount(54400001111)))
	fmt.Println(string(listTransactions(54400001111)))
	fmt.Println(statements)
	handleRequests()
}
