package handler

import (
	"blockchainFromScratch/datastore"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type AddressHandler struct {
	Blocks *[]datastore.Block
}

func (ah *AddressHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	address := params["address"]
	transactions := []datastore.Transaction{}
	for _, block := range *ah.Blocks {
		for _, transaction := range block.Transactions {
			if transaction.Recipient == address || transaction.Sender == address {
				transactions = append(transactions, transaction)
			}
		}
	}
	json.NewEncoder(w).Encode(transactions)
}
