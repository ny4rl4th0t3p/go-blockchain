package handler

import (
	"blockchainFromScratch/blockchain"
	"blockchainFromScratch/datastore"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type TransactionHandler struct {
	Chain *blockchain.Chain
}

type TransactionRequest struct {
	Amount int
}

func (th *TransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var transaction datastore.Transaction

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	transaction.TransactionId = uuid.New()

	fmt.Printf("%v", transaction)
	th.Chain.PendingTransactions = append(th.Chain.PendingTransactions, transaction)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(th.Chain.PendingTransactions)
}
