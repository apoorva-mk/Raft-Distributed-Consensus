package routing

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// StartRaft is the signal that triggers the raft
// behaviour in server clusters
func StartRaft(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request from client in startRaft.go: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var servers map[int]types.Server

	err = json.Unmarshal(body, &servers)
	if err != nil {
		log.Printf("Couldn't Unmarshal data in startRaft.go: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	outJSON, err := json.Marshal("Started Servers")
	if err != nil {
		log.Printf("Can't Marshall to JSON in startRaft.go: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(outJSON))

	err = ConcurrentReqRes(newReq[r.Host].Config, body, "/leaderElection", newReq[r.Host].ServerState.ID)

	if err != nil {
		log.Printf("Couldn't create requests to cluster in startRaft.go: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
