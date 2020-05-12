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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rbh.Chain)
}
