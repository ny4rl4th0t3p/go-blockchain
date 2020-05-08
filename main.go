package main

import (
	"blockchainFromScratch/blockchain"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {

	// init blockchain
	blockchain.Blockchain()

	// Create a new router
	r := mux.NewRouter()

	// Attach a path with handler
	r.HandleFunc("/blockchain", ProfileHandler).Methods("GET")
	r.HandleFunc("/transaction", ProfileHandler).Methods("GET")
	r.HandleFunc("/transaction/broadcast", ProfileHandler).Methods("GET")
	r.HandleFunc("/mine", ProfileHandler).Methods("GET")
	r.HandleFunc("/receive-new-block", ProfileHandler).Methods("GET")
	r.HandleFunc("/register-and-broadcast-node", ProfileHandler).Methods("GET")
	r.HandleFunc("/register-node", ProfileHandler).Methods("GET")
	r.HandleFunc("/register-nodes-bulk", ProfileHandler).Methods("GET")
	r.HandleFunc("/consensus", ProfileHandler).Methods("GET")
	r.HandleFunc("/block/{blockHash}", ProfileHandler).Methods("GET")
	r.HandleFunc("/transaction/{transactionId}", ProfileHandler).Methods("GET")
	r.HandleFunc("/address/{address}", ProfileHandler).Methods("GET")
	r.HandleFunc("/block-explorer", ProfileHandler).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
