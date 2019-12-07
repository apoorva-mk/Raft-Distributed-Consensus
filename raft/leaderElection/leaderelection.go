package leaderelection

import (
	"fmt"
	"math/rand"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// LeaderElection implements Raft leader election
func LeaderElection(raftServers map[string]types.RaftServer) {
	minTimeout := 150
	maxTimeout := 650

	timeOut := rand.Intn(maxTimeout-minTimeout) + minTimeout

	fmt.Println(timeOut)
}
