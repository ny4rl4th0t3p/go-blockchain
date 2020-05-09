package handler

import (
	"blockchainFromScratch/blockchain"
	"encoding/json"
	"net/http"
)

type BlockchainHandler struct {
	Chain *blockchain.Chain
}

func (bh *BlockchainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bh.Chain)
}
