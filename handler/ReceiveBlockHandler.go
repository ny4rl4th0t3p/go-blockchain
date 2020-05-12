package handler

import (
	"blockchainFromScratch/datastore"
	"encoding/json"
	"net/http"
)

type ReceiveBlockHandler struct {
	Chain *datastore.Chain
}

func (rbh *ReceiveBlockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var block datastore.Block
	err := json.NewDecoder(r.Body).Decode(&block)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if block.PreviousBlockHash == rbh.Chain.Blocks[len(rbh.Chain.Blocks)-1].Hash && block.Index == rbh.Chain.Blocks[len(rbh.Chain.Blocks)-1].Index {
		rbh.Chain.Blocks = append(rbh.Chain.Blocks, block)
		rbh.Chain.PendingTransactions = nil
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rbh.Chain)
}
