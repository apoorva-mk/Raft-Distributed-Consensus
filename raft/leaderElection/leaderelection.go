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
	// for !findALeader(config, IP) {
	findALeader(config, IP)
	log.Println("No leader found, repeating election")
	// }
	log.Println("Leader found!")
}

func incrementTermAndBecomeCandidate(IP string) {
	state := types.ServerData[IP]
	state.CurrentTerm++
	state.Name = "candidate"
	types.ServerData[IP] = state
}

func getTimer(minTimeout, maxTimeout int, IP string) *time.Timer {
	rand.Seed(time.Now().UnixNano())
	timeOut := rand.Intn(maxTimeout-minTimeout) + minTimeout
	log.Printf("%s set timeout at %d\n", IP, timeOut)
	return time.NewTimer(time.Duration(timeOut) * time.Millisecond)
}

func findALeader(config types.Configuration, IP string) bool {
	incrementTermAndBecomeCandidate(IP)
	log.Printf("%s incremented term to %d\n", IP, types.ServerData[IP].CurrentTerm)
	voteChan := make(chan int)
	timer := getTimer(150, 650, IP)
	requestvotes.RequestVotes(config, IP, timer, voteChan)
	votes := <-voteChan
	log.Printf("Votes for %s is %d\n", IP, votes)
	numberServers := len(config.Servers)
	if votes > numberServers/2 {
		return true
	}
	return false
}
