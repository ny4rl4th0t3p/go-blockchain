package handler

import (
	"blockchainFromScratch/datastore"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type ConsensusHandler struct {
	Chain *datastore.Chain
	Nodes []datastore.NetworkNode
}

func (ch *ConsensusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var blockchains []datastore.Chain
	for _, n := range ch.Nodes {
		resp, err := http.Get(n.NodeURL + ":" + strconv.Itoa(n.Port) + "/blockchain")
		if err != nil {
			log.Printf("Remote server connection error: [%s]\n", err)
			continue
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Remote server connection error: [%s]\n", err)
			continue
		}
		var chain datastore.Chain
		json.Unmarshal(body, &chain)
		blockchains = append(blockchains, chain)
	}

	for _, bc := range blockchains {
		if len(bc.Blocks) > len(ch.Chain.Blocks) {
			if bc.IsValid() {
				fmt.Println("valid")
				ch.Chain.Blocks = bc.Blocks
				ch.Chain.PendingTransactions = bc.PendingTransactions
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blockchains)
}
