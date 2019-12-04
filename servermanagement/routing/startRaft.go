package routing

import (
	"encoding/json"
	"fmt"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
	"io/ioutil"
	"log"
	"net/http"
)

// StartRaft is the signal that triggers the raft
// behaviour in server clusters
func StartRaft(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request from client in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var newReq map[string]types.RaftServer
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Couldn't Unmarshal data in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("I %s am a %s\n", r.Host, newReq[r.Host].State)
	outJSON, err := json.Marshal("Started Servers")
	if err != nil {
		log.Printf("Can't Marshall to JSON in startRaft.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(outJSON))

	//begin leader election
}
