package handler

import (
	"blockchainFromScratch/datastore"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type BlockHandler struct {
	Blocks *[]datastore.Block
}

func (bh *BlockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	blockHash := params["blockHash"]
	fmt.Println(blockHash)
	for _, block := range *bh.Blocks {
		if block.Hash == blockHash {
			json.NewEncoder(w).Encode(block)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("")
}
