package datastore

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Block struct {
	Index             int           `json:"index"`
	Timestamp         int64         `json:"timestamp"`
	Transactions      []Transaction `json:"transactions"`
	Nonce             int           `json:"nonce"`
	Hash              string        `json:"hash"`
	PreviousBlockHash string        `json:"previous_block_hash"`
}

func (block *Block) BlockHash(previousBlockHash string, dif string) string {
	t, err := json.Marshal(block.Transactions)
	if err != nil {
		fmt.Println(err)
	}
	nonce := 0
	sum := []byte("xxxx")
	for !strings.HasPrefix(hex.EncodeToString(sum), dif) {
		joinedBlockContent := previousBlockHash + string(t) + strconv.Itoa(nonce)
		h := sha1.New()
		h.Write([]byte(joinedBlockContent))
		sum = h.Sum(nil)
		nonce++
	}
	return hex.EncodeToString(sum)
}
