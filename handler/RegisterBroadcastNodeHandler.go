package handler

import (
	"blockchainFromScratch/datastore"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type RegisterBroadcastNodeHandler struct {
	Nodes     []datastore.NetworkNode
	LocalNode datastore.NetworkNode
}

func (rbnh *RegisterBroadcastNodeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var node datastore.NetworkNode
	err := json.NewDecoder(r.Body).Decode(&node)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	insert := true
	for _, n := range rbnh.Nodes {
		if node.NodeURL == n.NodeURL {
			insert = false
			break
		}
	}
	if node.NodeUID != "" && node.NodeURL != "" && node.NodeURL != rbnh.LocalNode.NodeURL && insert {
		rbnh.Nodes = append(rbnh.Nodes, node)
	}
	for _, n := range rbnh.Nodes {
		requestBody, err := json.Marshal(map[string]interface{}{
			"id":   n.NodeUID,
			"url":  n.NodeURL,
			"port": n.Port,
		})
		if err != nil {
			log.Fatalln(err)
		}
		resp, err := http.Post(n.NodeURL+":"+strconv.Itoa(n.Port)+"/register-node", "application/json", bytes.NewBuffer(requestBody))
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
	json.NewEncoder(w).Encode(rbnh.Nodes)
}
