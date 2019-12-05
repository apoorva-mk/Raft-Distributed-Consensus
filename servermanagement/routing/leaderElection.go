package routing

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// StartLeaderElection initiates leader election.
func StartLeaderElection(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request from client in leaderElection.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var newReq map[string]types.RaftServer
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Couldn't Unmarshal data in leaderElection.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
