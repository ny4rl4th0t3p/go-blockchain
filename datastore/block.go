package datastore

type Block struct {
	Index             int           `json:"index"`
	Timestamp         int64         `json:"timestamp"`
	Transactions      []Transaction `json:"transactions"`
	Nonce             int           `json:"nonce"`
	Hash              string        `json:"hash"`
	PreviousBlockHash string        `json:"previous_block_hash"`
}
