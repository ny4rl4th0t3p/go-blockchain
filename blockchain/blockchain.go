package blockchain

import (
	"blockchainFromScratch/datastore"
	"time"
)

var nodeUrl string
var chain []datastore.Block
var pendingTransactions []datastore.Transaction
var networkNodes []datastore.NetworkNode

func createNewBlock(nonce int, previousBlockHash string, hash string) datastore.Block {
	block := datastore.Block{
		Index:             len(chain),
		Timestamp:         time.Now().Unix(),
		Transactions:      pendingTransactions,
		Nonce:             nonce,
		Hash:              hash,
		PreviousBlockHash: previousBlockHash,
	}
	pendingTransactions = []datastore.Transaction{}
	chain = append(chain, block)
	return block
}

func getLastBlock() {

}

func hashBlock() {

}

func proofOfWork() {

}

func chainIsValid() {

}

func getBlock() {

}

func getTransaction() {

}

func getAddressData() {

}
