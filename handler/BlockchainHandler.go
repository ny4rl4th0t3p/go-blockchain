package handler

import (
	"blockchainFromScratch/datastore"
	"encoding/json"
	"net/http"
)

type BlockchainHandler struct {
	Chain *datastore.Chain
}

func (bh *BlockchainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bh.Chain)
}
