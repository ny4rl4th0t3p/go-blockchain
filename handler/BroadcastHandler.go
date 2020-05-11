package handler

import (
	"blockchainFromScratch/datastore"
	"encoding/json"
	"net/http"
)

type BroadcastHandler struct {
	Chain      *datastore.Chain
	KnownNodes []datastore.NetworkNode
}

func (bh *BroadcastHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var transaction datastore.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bh.Chain.BroadcastTransaction(bh.Chain.AddTransaction(transaction), bh.KnownNodes)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
