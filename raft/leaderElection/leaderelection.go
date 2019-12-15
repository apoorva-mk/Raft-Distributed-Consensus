package leaderelection

import (
	"log"
	"math/rand"
	"time"

	requestvotes "github.com/SUMUKHA-PK/Raft-Distributed-Consensus/raft/requestVotes"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// LeaderElection implements Raft leader election
func LeaderElection(config types.Configuration, IP string) {
	types.ServerData[IP] = incrementTermAndBecomeCandidate(types.ServerData, IP)
	log.Printf("%s incremented term to %d\n", IP, types.ServerData[IP].CurrentTerm)

	rand.Seed(time.Now().UnixNano())
	timer, duration := getTimer(150, 300)
	log.Printf("%s set timeout at %d\n", IP, duration)

	voteChan := make(chan int)
	requestvotes.RequestVotes(config, IP, timer, voteChan)
	votes := <-voteChan
	log.Printf("Votes for %s is %d\n", IP, votes)

	if majorityObtained(votes, config) {
		log.Printf("%v is elected as a leader", IP)
	}
}

func incrementTermAndBecomeCandidate(ServerData map[string]*types.State, IP string) *types.State {
	state := ServerData[IP]
	state.CurrentTerm++
	state.Name = "candidate"
	return state
}

// Assumes random value has been seeded
// Returns a timer and it's duration in milliseconds
func getTimer(minTimeout, maxTimeout int) (*time.Timer, int) {
	timeOut := rand.Intn(maxTimeout-minTimeout) + minTimeout
	return time.NewTimer(time.Duration(timeOut) * time.Millisecond), timeOut
}

func majorityObtained(votes int, config types.Configuration) bool {
	return len(config.Servers)/2 < votes
}
