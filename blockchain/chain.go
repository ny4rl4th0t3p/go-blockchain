package blockchain

import (
	"blockchainFromScratch/datastore"
	"time"
)

type Chain struct {
	Blocks              []datastore.Block       `json:"blocks"`
	PendingTransactions []datastore.Transaction `json:"pending_transactions"`
}

func (chain *Chain) Init() {
	pendingTransactions = []datastore.Transaction{}
	networkNodes = []datastore.NetworkNode{}
	networkNodes = append(networkNodes, datastore.NetworkNode{
		NodeUID: "f98hf9oph39",
		NodeURL: "http://localhost",
		Port:    8000,
	})

	chain.Blocks = append(chain.Blocks, datastore.Block{
		Index:             1,
		Timestamp:         time.Now().Unix(),
		Transactions:      nil,
		Nonce:             0,
		Hash:              "63N3515-8L0CK",
		PreviousBlockHash: "",
	})

}

func (chain Chain) AddTransaction(transaction datastore.Transaction) {
	//this.pendingTransactions.push(transactionObj);
	//return this.getLastBlock()['index'] + 1;

	pendingTransactions = append(pendingTransactions, transaction)

}
