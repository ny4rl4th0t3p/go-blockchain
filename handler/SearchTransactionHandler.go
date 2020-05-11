package handler

import (
	"blockchainFromScratch/datastore"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type SearchTransactionHandler struct {
	Blocks *[]datastore.Block
}

func (sth *SearchTransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	transactionId := params["transactionId"]
	for _, block := range *sth.Blocks {
		for _, transaction := range block.Transactions {
			fmt.Println(transaction)
			if strings.Compare(transaction.TransactionId.(uuid.UUID).String(), transactionId) == 0 {
				json.NewEncoder(w).Encode(transaction)
				break
			}

		}
	}
}
