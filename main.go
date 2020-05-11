package main

import (
	"blockchainFromScratch/datastore"
	"blockchainFromScratch/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {

	// init blockchain
	chain := datastore.Chain{}
	chain.Init()

	node := datastore.NetworkNode{}
	node.NodeUID = ""
	node.NodeURL = "http://127.0.0.1"
	node.Port = 8000

	var knownNodes []datastore.NetworkNode
	knownNodes = append(knownNodes, datastore.NetworkNode{
		NodeUID: "f98hf9oph39",
		NodeURL: "http://server",
		Port:    8000,
	})

	// Create a new router
	r := mux.NewRouter()

	blockchainHandler := &handler.BlockchainHandler{Chain: &chain}
	r.Handle("/blockchain", blockchainHandler).Methods("GET")

	transactionHandler := &handler.TransactionHandler{Chain: &chain}
	r.Handle("/transaction", transactionHandler).Methods("POST")

	broadcastHandler := &handler.BroadcastHandler{Chain: &chain, KnownNodes: knownNodes}
	r.Handle("/transaction/broadcast", broadcastHandler).Methods("POST")

	mineHandler := &handler.MineHandler{Chain: &chain}
	r.Handle("/mine", mineHandler).Methods("GET")

	searchBlockHandler := &handler.SearchBlockHandler{Blocks: &chain.Blocks}
	r.Handle("/block/{blockHash}", searchBlockHandler).Methods("GET")

	searchTransactionHandler := &handler.SearchTransactionHandler{Blocks: &chain.Blocks}
	r.Handle("/transaction/{transactionId}", searchTransactionHandler).Methods("GET")

	//r.HandleFunc("/register-and-broadcast-node", ProfileHandler).Methods("GET")
	//r.HandleFunc("/register-node", ProfileHandler).Methods("GET")
	//r.HandleFunc("/register-nodes-bulk", ProfileHandler).Methods("GET")
	//r.HandleFunc("/consensus", ProfileHandler).Methods("GET")
	//r.HandleFunc("/address/{address}", ProfileHandler).Methods("GET")
	//r.HandleFunc("/block-explorer", ProfileHandler).Methods("GET")

	// TODO
	receiveBlockHandler := &handler.ReceiveBlockHandler{Chain: &chain}
	r.Handle("/receive-new-block", receiveBlockHandler).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
