package handler

import (
	"blockchainFromScratch/datastore"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type SearchBlockHandler struct {
	Blocks *[]datastore.Block
}

func (sbh *SearchBlockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	blockHash := params["blockHash"]
	for _, block := range *sbh.Blocks {
		if block.Hash == blockHash {
			json.NewEncoder(w).Encode(block)
			break
		}
	}
}
