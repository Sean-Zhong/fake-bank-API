# fake-bank-API
This is a simple API that takes camt053 xml data and parses into json that can be aquired from set endpoints following RESTful API principles.

## Prerequisites
* GO installation - https://go.dev/dl/ (version 1.18.1 was used for this project)
* Gin Web Framework - https://github.com/gin-gonic/gin
* (optional) Postman - used for sending http requests

## Instructions
1. Fork or download the project onto your computer
2. Run api.go from fake-bank-API/main with command: `go run api.go` in terminal/cmd
3. In Postman or directly in your browser, go to desired endpoint.

### Endpoints
* listaccounts: list all accounts - `localhost:8081/listaccounts`
* getaccount: get details of a specific account - `localhost:8081/getaccount/"accountNumber"`
* listtransactions: list all transactions of a specific account - `localhost:8081/listtransactions/"accountNumber"`

"accountNumber" needs to be substituted with an actual account number of the desired account.
