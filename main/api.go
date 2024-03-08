package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sean-Zhong/fake-bank-API/parser"
)

const FILENAME string = "output.json"
const PORT string = ":8082"

var parsedData []byte = parser.Parse()

type StoredAccount struct {
}

func getAccounts() {
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func main() {
	fmt.Println(string(parsedData))
	handleRequests()
}
