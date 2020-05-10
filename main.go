package main

import (
	"blockchainFromScratch/blockchain"
	"blockchainFromScratch/handler"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {

	// init blockchain
	chain := blockchain.Chain{}
	chain.Init()

	b, err := json.Marshal(chain)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	// Create a new router
	r := mux.NewRouter()

	blockchainHandler := &handler.BlockchainHandler{Chain: &chain}
	r.Handle("/blockchain", blockchainHandler).Methods("GET")

	transactionHandler := &handler.TransactionHandler{Chain: &chain}
	r.Handle("/transaction", transactionHandler).Methods("POST")

	broadcastHandler := &handler.BroadcastHandler{Chain: &chain}
	r.Handle("/transaction/broadcast", broadcastHandler).Methods("POST")

	//r.HandleFunc("/mine", ProfileHandler).Methods("GET")
	//r.HandleFunc("/receive-new-block", ProfileHandler).Methods("GET")
	//r.HandleFunc("/register-and-broadcast-node", ProfileHandler).Methods("GET")
	//r.HandleFunc("/register-node", ProfileHandler).Methods("GET")
	//r.HandleFunc("/register-nodes-bulk", ProfileHandler).Methods("GET")
	//r.HandleFunc("/consensus", ProfileHandler).Methods("GET")
	//r.HandleFunc("/block/{blockHash}", ProfileHandler).Methods("GET")
	//r.HandleFunc("/transaction/{transactionId}", ProfileHandler).Methods("GET")
	//r.HandleFunc("/address/{address}", ProfileHandler).Methods("GET")
	//r.HandleFunc("/block-explorer", ProfileHandler).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
