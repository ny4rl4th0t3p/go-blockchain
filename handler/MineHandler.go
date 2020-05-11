package handler

import (
	"blockchainFromScratch/datastore"
	"encoding/json"
	"net/http"
	"time"
)

type MineHandler struct {
	Chain *datastore.Chain
}

func (mh *MineHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	lastBlock := mh.Chain.GetLastBlock()

	newBlock := datastore.Block{
		Index:             lastBlock.Index + 1,
		Timestamp:         time.Now().Unix(),
		Transactions:      mh.Chain.PendingTransactions,
		PreviousBlockHash: lastBlock.Hash,
	}

	//
	// hacer PoW para meter el nonce en newBlock.Nonce
	//

	//newBlock.Hash = mh.Chain.BlockHash(lastBlock.Hash, newBlock)

	newBlock.Hash = newBlock.BlockHash(lastBlock.Hash, mh.Chain.Dif)

	mh.Chain.AddNewBlock(newBlock)
	mh.Chain.PendingTransactions = nil

	//
	// Broadcast block
	//

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mh.Chain)
}
