package requestvotes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/general"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// RequestVotes implements request votes rpc of raft
func RequestVotes(config types.Configuration, IP string, timer *time.Timer, finalVote chan int) {
	vote := 0
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		<-timer.C
		wg.Done()
		finalVote <- vote
	}()
	resCh := make(chan []types.URLResponse)
	var voteRes []types.URLResponse
	wg.Add(1)
	go func() {
		payload, err := json.Marshal(types.ReqVotesRequest{config, types.ServerData[IP].CurrentTerm, types.ServerData[IP].ID, -1, -1})
		if err != nil {
			log.Printf("Can't Marshall to JSON in requestVotes.go: %v\n", err)
		}
		voteRes, err = general.ConcurrentReqRes(config, []byte(payload), "/requestVotes", types.ServerData[IP].ID)
		if err != nil {
			fmt.Println(err)
		}
		resCh <- voteRes
		wg.Done()
	}()
	vote = getVotes(<-resCh)
	wg.Wait()
}

func getVotes(voteRes []types.URLResponse) int {
	count := 0
	for i := range voteRes {
		body, err := ioutil.ReadAll(voteRes[i].Res.Body)
		if err != nil {
			log.Printf("Bad request from client in requestVotes.go: %v\n", err)

		}
		var newReq int
		err = json.Unmarshal(body, &newReq)
		if err != nil {
			log.Printf("Couldn't Unmarshal data in requestVotes.go: %v\n", err)
		}
		if newReq == 1 {
			count++
		}
	}
	return count
}
