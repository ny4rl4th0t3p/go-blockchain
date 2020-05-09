package datastore

type Transaction struct {
	Amount        int         `json:"amount"`
	Sender        string      `json:"sender"`
	Recipient     string      `json:"recipient"`
	TransactionId interface{} `json:"id"`
}
