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

func (bh *BroadcastHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var transaction datastore.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bh.Chain.BroadcastTransaction(bh.Chain.AddTransaction(transaction))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
