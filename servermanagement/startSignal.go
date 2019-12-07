package servermanagement

import (
	"encoding/json"
	"log"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// StartSignal initiates the servers to behave
// according to the raft protocol.
func StartSignal(config types.Configuration, RaftServers map[string]types.RaftServer) error {
	for i := range config.Servers {
		RaftServers[config.Servers[i].IP+":"+config.Servers[i].Port] =
			types.RaftServer{
				ServerState: types.State{
					Name:        "follower",
					ID:          i,
					CurrentTerm: 0,
					VotedFor:    -2,
					Log:         []types.LogData{},
					CommitIndex: -1,
					LastApplied: -1,
					NextIndex:   []int{},
					MatchIndex:  []int{},
				},
				Config: config,
			}
	}
	payload, err := json.Marshal(RaftServers)
	if err != nil {
		log.Printf("Can't Marshall to JSON in startSignal.go : %v\n", err)
		return err
	}
	_, err = ConcurrentReqRes(config, payload, "/startRaft", "-1")
	return err
}
