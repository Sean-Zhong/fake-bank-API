package main

import (
	"errors"
	"net/http"

	"github.com/Sean-Zhong/fake-bank-API/parser"
	"github.com/gin-gonic/gin"
)

const PORT string = "8081"

var statements parser.BankStatements = parser.Parse()

type AccountDetails struct {
	Account parser.Account
	Balance []parser.Balance
}

func listAccounts(context *gin.Context) {
	var accountList []parser.Account
	for _, statement := range statements.Statements {
		accountList = append(accountList, statement.AccountInfo)
	}
	context.IndentedJSON(http.StatusOK, accountList)
}

func getAccountById(accountId string) (*AccountDetails, error) {
	for _, stmt := range statements.Statements {
		if stmt.AccountInfo.AccountId == accountId {
			return &AccountDetails{Account: stmt.AccountInfo, Balance: stmt.Balances}, nil
		}
	}
	return nil, errors.New("account not found")
}

func getAccount(context *gin.Context) {
	accountId := context.Param("accountId")
	accountData, err := getAccountById(accountId)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Account not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, accountData)
}

func getStatementById(accountId string) (*parser.Statement, error) {
	for _, stmt := range statements.Statements {
		if stmt.AccountInfo.AccountId == accountId {
			return &stmt, nil
		}
	}
	return nil, errors.New("statement not found")
}

func listTransactions(context *gin.Context) {
	accountId := context.Param("accountId")
	statement, err := getStatementById(accountId)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Account not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, statement.Transactions)
}

func main() {
	router := gin.Default()
	router.GET("/listtransactions/:accountId", listTransactions)
	router.GET("/getaccount/:accountId", getAccount)
	router.GET("/listaccounts", listAccounts)
	router.Run("localhost:" + PORT)
}
