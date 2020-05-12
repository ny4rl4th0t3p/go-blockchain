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
	node.Port = 8001

	var knownNodes []datastore.NetworkNode
	knownNodes = append(knownNodes, datastore.NetworkNode{
		NodeUID: "f98hf9oph39",
		NodeURL: "http://127.0.0.1",
		Port:    8000,
	})

	// Create a new router
	r := mux.NewRouter()

	blockchainHandler := &handler.BlockchainHandler{Chain: &chain}
	r.Handle("/blockchain", blockchainHandler).Methods("GET")

	transactionHandler := &handler.TransactionHandler{Chain: &chain}
	r.Handle("/transaction", transactionHandler).Methods("POST")

	broadcastHandler := &handler.BroadcastTransactionHandler{Chain: &chain, KnownNodes: knownNodes}
	r.Handle("/transaction/broadcast", broadcastHandler).Methods("POST")

	mineHandler := &handler.MineHandler{Chain: &chain, Nodes: knownNodes}
	r.Handle("/mine", mineHandler).Methods("GET")

	searchBlockHandler := &handler.SearchBlockHandler{Blocks: &chain.Blocks}
	r.Handle("/block/{blockHash}", searchBlockHandler).Methods("GET")

	searchTransactionHandler := &handler.SearchTransactionHandler{Blocks: &chain.Blocks}
	r.Handle("/transaction/{transactionId}", searchTransactionHandler).Methods("GET")

	addressHandler := &handler.AddressHandler{Blocks: &chain.Blocks}
	r.Handle("/address/{address}", addressHandler).Methods("GET")

	registerNodeHandler := &handler.RegisterNodeHandler{Nodes: knownNodes, LocalNode: node}
	r.Handle("/register-node", registerNodeHandler).Methods("POST")

	registerBroadcastNodeHandler := &handler.RegisterBroadcastNodeHandler{Nodes: knownNodes, LocalNode: node}
	r.Handle("/register-and-broadcast-node", registerBroadcastNodeHandler).Methods("POST")

	consensusNodeHandler := &handler.ConsensusHandler{Chain: &chain, Nodes: knownNodes}
	r.Handle("/consensus", consensusNodeHandler).Methods("GET")

	// TODO
	receiveBlockHandler := &handler.ReceiveBlockHandler{Chain: &chain}
	r.Handle("/receive-new-block", receiveBlockHandler).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
