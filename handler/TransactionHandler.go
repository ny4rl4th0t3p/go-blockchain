package handler

import (
	"blockchainFromScratch/blockchain"
	"blockchainFromScratch/datastore"
	"encoding/json"
	"net/http"
)

type TransactionHandler struct {
	Chain *blockchain.Chain
}

func (th *TransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var transaction datastore.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	th.Chain.AddTransaction(transaction)
	//th.Chain.BroadcastTransaction(transaction)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
