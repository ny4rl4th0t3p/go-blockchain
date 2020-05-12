package handler

import (
	"blockchainFromScratch/datastore"
	"encoding/json"
	"net/http"
)

type BroadcastTransactionHandler struct {
	Chain      *datastore.Chain
	KnownNodes []datastore.NetworkNode
}

func (bth *BroadcastTransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var transaction datastore.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bth.Chain.BroadcastTransaction(bth.Chain.AddTransaction(transaction), bth.KnownNodes)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
