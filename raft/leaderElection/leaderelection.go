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
func LeaderElection(config types.Configuration, IP string) {
	incrementTermAndBecomeCandidate(IP)
	log.Printf("%s incremented term to %d\n", IP, types.ServerData[IP].CurrentTerm)
	voteChan := make(chan int)
	timer := getTimer(150, 650, IP)
	requestvotes.RequestVotes(config, IP, timer, voteChan)
	votes := <-voteChan
	fmt.Printf("Votes for %s is %d\n", IP, votes)
}

func incrementTermAndBecomeCandidate(IP string) {
	state := types.ServerData[IP]
	state.CurrentTerm++
	state.Name = "candidate"
	types.ServerData[IP] = state
}

func getTimer(min, max int, IP string) *time.Timer {
	minTimeout := 150
	maxTimeout := 650
	rand.Seed(time.Now().UnixNano())
	timeOut := rand.Intn(maxTimeout-minTimeout) + minTimeout
	log.Printf("%s set timeout at %d\n", IP, timeOut)
	return time.NewTimer(time.Duration(timeOut) * time.Millisecond)
}
