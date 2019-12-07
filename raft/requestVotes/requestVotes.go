package requestvotes

import (
	"log"
	"math/rand"
	"time"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// RequestVotes implements request votes rpc of raft
func RequestVotes(raftServers map[string]types.RaftServer, IP string, timer *time.Timer, finalVote chan int) {
	vote := 0
	go func() {
		<-timer.C
		finalVote <- vote
	}()
	random := rand.Intn(500) + 400
	log.Printf("%s reply received in %d\n", IP, random)
	time.Sleep(time.Duration(random) * time.Millisecond)
	vote++
	// return vote + 1

	// go servermanagement.ConcurrentReqRes(raftServers[IP].Config,[]byte(raftServers),"")
}
