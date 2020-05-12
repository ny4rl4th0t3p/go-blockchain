package datastore

type NetworkNode struct {
	NodeUID string `json:"id"`
	NodeURL string `json:"url"`
	Port    int    `json:"port"`
}
