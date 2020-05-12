package datastore

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Chain struct {
	Blocks              []Block       `json:"blocks"`
	PendingTransactions []Transaction `json:"pending_transactions"`
	Dif                 string        `json:"dificulty"`
}

func (chain *Chain) Init() {
	chain.PendingTransactions = []Transaction{}
	chain.Dif = "0000"
	hash := sha256.New()
	hash.Write([]byte("63N3515-8L0CK"))
	chain.Blocks = append(chain.Blocks, Block{
		Index:             1,
		Timestamp:         time.Now().Unix(),
		Transactions:      nil,
		Nonce:             0,
		Hash:              hex.EncodeToString(hash.Sum(nil)),
		PreviousBlockHash: "",
	})
}

func (chain *Chain) AddTransaction(transaction Transaction) Transaction {
	transaction.TransactionId = uuid.New()
	chain.PendingTransactions = append(chain.PendingTransactions, transaction)
	return transaction
}

func (chain *Chain) BroadcastTransaction(transaction Transaction, knownNodes []NetworkNode) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"amount":    transaction.Amount,
		"sender":    transaction.Sender,
		"recipient": transaction.Recipient,
		"id":        transaction.TransactionId,
	})
	if err != nil {
		log.Fatalln(err)
	}
	for _, node := range knownNodes {
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

func (chain *Chain) GetLastBlock() Block {
	return chain.Blocks[len(chain.Blocks)-1]
}

func (chain *Chain) AddNewBlock(block Block) {
	chain.Blocks = append(chain.Blocks, block)
}

func (chain *Chain) IsValid() bool {
	isValid := true
	for i := 1; i < len(chain.Blocks); i++ {
		if !strings.HasPrefix(chain.Blocks[i].Hash, chain.Dif) && chain.Blocks[i].BlockHash(chain.Blocks[i-1].Hash) != chain.Blocks[i].Hash || chain.Blocks[i].PreviousBlockHash != chain.Blocks[i-1].Hash {
			isValid = false
		}
	}
	return isValid
}
