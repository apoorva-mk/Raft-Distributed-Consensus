package routing

import (
	"encoding/json"
	"fmt"
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
	var newReq map[string]types.RaftServer
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Couldn't Unmarshal data in startRaft.go: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("I %s, am a %s\n", r.Host, newReq[r.Host].ServerState.Name)
	outJSON, err := json.Marshal("Started Servers")
	if err != nil {
		log.Printf("Can't Marshall to JSON in startRaft.go: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(outJSON))

	// if r.Host == "127.0.0.1:55550" {
	// 	for k := range newReq[r.Host].DesignationMap {
	// 		fmt.Println(k)
	// 	}
	// 	fmt.Println(newReq[r.Host].Config)
	// }
	err = ConcurrentReqRes(newReq[r.Host].Config, body, "/leaderElection", newReq[r.Host].DesignationMap[r.Host].ID)
	if err != nil {
		log.Printf("Couldn't create requests to cluster in startRaft.go: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
