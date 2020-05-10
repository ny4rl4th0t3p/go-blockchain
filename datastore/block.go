package datastore

type Block struct {
	Index             int           `json:"index"`
	Timestamp         int64         `json:"timestamp"`
	Transactions      []Transaction `json:"transactions"`
	Nonce             int           `json:"nonce"`
	Hash              []byte        `json:"hash"`
	PreviousBlockHash []byte        `json:"previous_block_hash"`
}
