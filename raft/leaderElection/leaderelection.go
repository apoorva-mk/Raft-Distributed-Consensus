package leaderelection

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	appendentries "github.com/SUMUKHA-PK/Raft-Distributed-Consensus/raft/appendEntries"
	requestvotes "github.com/SUMUKHA-PK/Raft-Distributed-Consensus/raft/requestVotes"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// LeaderElection implements Raft leader election
func LeaderElection(config types.Configuration, IP string) bool {
	// if this server became the leader
	if findLeader(config, IP) {
		appendentries.AppendEntries(config, IP)
		return true
	}
	// checking for a heartbeat from some other leader
	// first the server waits for a heartbeat with a
	// timeout. And later checks if there was an
	// appendEntries from a peer by checking its state.

	leaderFoundFlag := 0
	heartBeatTimeOut := getTimer(250, 400, IP)
	<-heartBeatTimeOut.C
	if types.ServerData[IP].Name == "follower" {
		leaderFoundFlag = 1
	} else {
		types.ServerData[IP].VotedFor = -2
		fmt.Printf("%s set value to -2\n", IP)
	}
	if leaderFoundFlag == 1 {
		return true
	}
	return false
}

func findLeader(config types.Configuration, IP string) bool {
	incrementTermAndBecomeCandidate(IP)
	log.Printf("%s incremented term to %d\n", IP, types.ServerData[IP].CurrentTerm)
	voteChan := make(chan int)
	timer := getTimer(150, 300, IP)
	requestvotes.RequestVotes(config, IP, timer, voteChan)
	votes := <-voteChan
	log.Printf("Votes for %s is %d\n", IP, votes)
	numberServers := len(config.Servers)
	if votes > numberServers/2 {
		return true
	}
	return false
}

func incrementTermAndBecomeCandidate(IP string) {
	state := types.ServerData[IP]
	state.CurrentTerm++
	state.Name = "candidate"
	state.VotedFor = -2
	types.ServerData[IP] = state
}

func getTimer(minTimeout, maxTimeout int, IP string) *time.Timer {
	rand.Seed(time.Now().UnixNano())
	timeOut := rand.Intn(maxTimeout-minTimeout) + minTimeout
	log.Printf("%s set timeout at %d\n", IP, timeOut)
	return time.NewTimer(time.Duration(timeOut) * time.Millisecond)
}

/*

	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// go func() {
	// case where this server isnt the leader
	for !findLeader(config, IP) {

		// log.Println("No leader found, repeating election")
	}
	// 	wg.Done()
	// }()
	// wg.Add(1)
	// go func() {
	if leaderFoundFlag == 0 {
		// case where this server is the leader
		appendentries.AppendEntries(config, IP)
	}
	// 	wg.Done()
	// }()
	// wg.Wait()
*/
