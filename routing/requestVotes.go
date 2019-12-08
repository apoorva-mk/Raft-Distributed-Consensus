package routing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// RequestVotes is the signal that triggers the raft
// behaviour in server clusters
func RequestVotes(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request from client in startRaft.go: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var newReq map[string]types.RaftServer
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Couldn't Unmarshal data in startRaft.go: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	state := types.ServerData[r.Host]
	vote := voting(state, r.Host)
	outJSON, err := json.Marshal(vote)
	if err != nil {
		log.Printf("Can't Marshall to JSON in startRaft.go: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(outJSON))
}

func updateVoting(state *types.State, IP string) {
	state.Lock.Lock()
	state.VotedFor = 1
	state.Lock.Unlock()
}

func voting(state *types.State, IP string) (vote int) {
	state.Lock.Lock()
	if state.VotedFor == -2 {
		fmt.Printf("I %s can vote\n", IP)
		vote = 1
		state.VotedFor = 1
	} else {
		fmt.Printf("I %s cant vote\n", IP)
		vote = 0
	}
	state.Lock.Unlock()
	return
}
