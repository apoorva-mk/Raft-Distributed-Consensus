package routing

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// AppendEntries implements the append entries end point
func AppendEntries(w http.ResponseWriter, r *http.Request) {
	types.ServerData[r.Host].Name = "follower"
	vote := 1
	// fmt.Println("In append entries")
	outJSON, err := json.Marshal(vote)
	if err != nil {
		log.Printf("Can't Marshall to JSON in requestVotes.go: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(outJSON))
}
