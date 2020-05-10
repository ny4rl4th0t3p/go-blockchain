package blockchain

import (
	"blockchainFromScratch/datastore"
	"bytes"
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
		Hash:              "63N3515-8L0CK",
		PreviousBlockHash: "",
	})
}

func (chain *Chain) AddTransaction(transaction datastore.Transaction) datastore.Transaction {
	if transaction.TransactionId != "" {
		transaction.TransactionId = uuid.New()
	}
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
