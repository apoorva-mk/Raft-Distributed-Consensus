package leaderelection

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	requestvotes "github.com/SUMUKHA-PK/Raft-Distributed-Consensus/raft/requestVotes"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// LeaderElection implements Raft leader election
func LeaderElection(raftServers map[string]types.RaftServer, IP string) {
	incrementTermAndBecomeCandidate(raftServers, IP)
	log.Printf("%s incremented term to %d\n", IP, raftServers[IP].ServerState.CurrentTerm)

	votes := make(chan int)

	timer := getTimer(150, 650, IP)

	go func() {
		requestvotes.RequestVotes(raftServers, IP, timer, votes)
	}()
	// go func() {
	// 	fmt.Printf("dd%vdd", <-timer.C)
	// 	fmt.Println("Timer expired")
	// }()

	fmt.Println(<-votes)

}

func incrementTermAndBecomeCandidate(raftServer map[string]types.RaftServer, IP string) {
	state := raftServer[IP]
	state.ServerState.CurrentTerm++
	state.ServerState.Name = "candidate"
	raftServer[IP] = state
}

func getTimer(min, max int, IP string) *time.Timer {
	minTimeout := 150
	maxTimeout := 650
	rand.Seed(time.Now().UnixNano())
	timeOut := rand.Intn(maxTimeout-minTimeout) + minTimeout
	log.Printf("%s set timeout at %d\n", IP, timeOut)
	return time.NewTimer(time.Duration(timeOut) * time.Millisecond)
}
