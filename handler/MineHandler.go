package handler

import (
	"blockchainFromScratch/datastore"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type MineHandler struct {
	Chain *datastore.Chain
	Nodes []datastore.NetworkNode
}

func (mh *MineHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	lastBlock := mh.Chain.GetLastBlock()

	newBlock := datastore.Block{
		Index:             lastBlock.Index + 1,
		Timestamp:         time.Now().Unix(),
		Transactions:      mh.Chain.PendingTransactions,
		PreviousBlockHash: lastBlock.Hash,
	}

	newBlock.Hash = newBlock.PoW(lastBlock.Hash, mh.Chain.Dif)

	mh.Chain.AddNewBlock(newBlock)
	mh.Chain.PendingTransactions = nil

	requestBody, err := json.Marshal(map[string]interface{}{
		"index":               newBlock.Index,
		"timestamp":           newBlock.Timestamp,
		"transactions":        newBlock.Transactions,
		"nonce":               newBlock.Nonce,
		"hash":                newBlock.Hash,
		"previous_block_hash": newBlock.PreviousBlockHash,
	})
	if err != nil {
		log.Fatalln(err)
	}

	for _, node := range mh.Nodes {
		resp, err := http.Post(node.NodeURL+":"+strconv.Itoa(node.Port)+"/receive-new-block", "application/json", bytes.NewBuffer(requestBody))
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mh.Chain)
}
