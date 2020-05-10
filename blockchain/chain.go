package blockchain

import (
	"blockchainFromScratch/datastore"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
		NodeURL: "http://server",
		Port:    8000,
	})

	chain.Blocks = append(chain.Blocks, datastore.Block{
		Index:             1,
		Timestamp:         time.Now().Unix(),
		Transactions:      nil,
		Nonce:             0,
		Hash:              []byte("63N3515-8L0CK"),
		PreviousBlockHash: []byte(""),
	})
}

func (chain *Chain) AddTransaction(transaction datastore.Transaction) datastore.Transaction {
	transaction.TransactionId = uuid.New()
	chain.PendingTransactions = append(chain.PendingTransactions, transaction)
	return transaction
}

func (chain *Chain) BroadcastTransaction(transaction datastore.Transaction) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"amount":    transaction.Amount,
		"sender":    transaction.Sender,
		"recipient": transaction.Recipient,
		"id":        transaction.TransactionId,
	})
	if err != nil {
		log.Fatalln(err)
	}
	for _, node := range networkNodes {
		fmt.Println("entrando a tope")
		resp, err := http.Post(node.NodeURL+":"+strconv.Itoa(node.Port)+"/transaction", "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			log.Printf("Remote server connection error: [%s]\n", err)
			continue
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(body))
	}
}

func (chain *Chain) GetLastBlock() datastore.Block {
	return chain.Blocks[len(chain.Blocks)-1]
}

func (chain *Chain) BlockHash(previousBlockHash []byte, block datastore.Block) []byte {

	b, err := json.Marshal(block)
	if err != nil {
		fmt.Println(err)
	}
	var blockContent [][]byte
	blockContent = append(blockContent, previousBlockHash)
	blockContent = append(blockContent, []byte(string(b)))
	joinedBlockContent := bytes.Join(blockContent, []byte("|"))
	h := sha256.New()
	h.Write(joinedBlockContent)
	return h.Sum(nil)
}

func (chain *Chain) AddNewBlock(block datastore.Block) {

}
