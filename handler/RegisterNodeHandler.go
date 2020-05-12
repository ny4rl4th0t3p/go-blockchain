package handler

import (
	"blockchainFromScratch/datastore"
	"encoding/json"
	"net/http"
)

type RegisterNodeHandler struct {
	Nodes     []datastore.NetworkNode
	LocalNode datastore.NetworkNode
}

func (rnh *RegisterNodeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var node datastore.NetworkNode
	err := json.NewDecoder(r.Body).Decode(&node)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	insert := true
	for _, n := range rnh.Nodes {
		if node.NodeURL == n.NodeURL {
			insert = false
			break
		}
	}
	if node.NodeUID != "" && node.NodeURL != "" && node.NodeURL != rnh.LocalNode.NodeURL && insert {
		rnh.Nodes = append(rnh.Nodes, node)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rnh.Nodes)
}
