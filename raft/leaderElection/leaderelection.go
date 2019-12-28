package leaderelection

import (
	"log"
	"math/rand"
	"time"

	appendentries "github.com/SUMUKHA-PK/Raft-Distributed-Consensus/raft/appendEntries"
	requestvotes "github.com/SUMUKHA-PK/Raft-Distributed-Consensus/raft/requestVotes"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// LeaderElection implements Raft leader election
// The election process is initially started by all servers
// concurrently. The fastest on
func LeaderElection(config types.Configuration, IP string) bool {
	if types.ServerData[IP].VotedFor == -2 {
		// if this server became the leader
		if findLeader(config, IP) {
			appendentries.AppendEntries(config, IP)
			return true
		}
	}
	// checking for a heartbeat from some other leader
	// first the server waits for a heartbeat with a
	// timeout. And later checks if there was an
	// appendEntries from a peer by checking its state.
	leaderFoundFlag := 0
	heartBeatTimeOut, _ := getTimer(150, 300)
	<-heartBeatTimeOut.C
	if types.ServerData[IP].Name == "follower" {
		leaderFoundFlag = 1
	} else {
		types.ServerData[IP].VotedFor = -2
		// fmt.Printf("%s set value to -2\n", IP)
	}
	if leaderFoundFlag == 1 {
		return true
	}
	return false

	// types.ServerData[IP] = incrementTermAndBecomeCandidate(types.ServerData, IP)
	// log.Printf("%s incremented term to %d\n", IP, types.ServerData[IP].CurrentTerm)

	// if majorityObtained(votes, config) {
	// 	log.Printf("%v is elected as a leader", IP)
	// }
}

func findLeader(config types.Configuration, IP string) bool {
	incrementTermAndBecomeCandidate(types.ServerData, IP)
	log.Printf("%s incremented term to %d\n", IP, types.ServerData[IP].CurrentTerm)
	voteChan := make(chan int)
	rand.Seed(time.Now().UnixNano())
	timer, duration := getTimer(150, 300)
	log.Printf("%s set timeout at %d\n", IP, duration)

	requestvotes.RequestVotes(config, IP, timer, voteChan)
	votes := <-voteChan
	// vote for one-self only if the vote hasn't been given
	if types.ServerData[IP].VotedFor == -2 {
		types.ServerData[IP].VotedFor = -1
		votes++
	}
	log.Printf("Votes for %s is %d\n", IP, votes)
	numberServers := len(config.Servers)
	if votes > numberServers/2 {
		return true
	}
	return false
}

func incrementTermAndBecomeCandidate(ServerData map[string]*types.State, IP string) *types.State {
	state := ServerData[IP]
	state.CurrentTerm++
	state.Name = "candidate"
	state.VotedFor = -2
	types.ServerData[IP] = state
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
