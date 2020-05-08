package blockchain

import "time"

//type Chain struct {
//}

type Transaction struct {
}

type NetworkNode struct {
}

type Block struct {
	Index             int
	Timestamp         time.Time
	Transactions      []Transaction
	Nonce             int
	Hash              string
	PreviousBlockHash string
}

var nodeUrl string
var chain []Block
var pendingTransactions []Transaction
var networkNodes []NetworkNode

func Blockchain() {
	nodeUrl = ""
	chain = []Block{}
	pendingTransactions = []Transaction{}
	networkNodes = []NetworkNode{}
	// create genesis block
	createNewBlock(666, "0", "0")
}

func createNewBlock(nonce int, previousBlockHash string, hash string) Block {
	block := Block{
		Index:             len(chain),
		Timestamp:         time.Now(),
		Transactions:      pendingTransactions,
		Nonce:             nonce,
		Hash:              hash,
		PreviousBlockHash: previousBlockHash,
	}
	pendingTransactions = []Transaction{}
	chain = append(chain, block)
	return block
}

func getLastBlock() {

}

func createNewTransaction() {

}

func addTransactionToPendingTransactions() {

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
