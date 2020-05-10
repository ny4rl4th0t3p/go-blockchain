package handler

import (
	"blockchainFromScratch/blockchain"
	"blockchainFromScratch/datastore"
	"encoding/json"
	"net/http"
)

type BroadcastHandler struct {
	Chain *blockchain.Chain
}

type BroadcastRequest struct {
	Amount int
}

func (th *BroadcastHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var transaction datastore.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	th.Chain.BroadcastTransaction(th.Chain.AddTransaction(transaction))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
