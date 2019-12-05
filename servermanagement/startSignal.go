package servermanagement

import (
	"encoding/json"
	"log"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/servermanagement/routing"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// StartSignal initiates the servers to behave
// according to the raft protocol.
func StartSignal(config types.Configuration, RaftServers map[string]types.RaftServer) error {
	designationMap := make(map[string]types.State)
	for i := range config.Servers {
		designationMap[config.Servers[i].IP+":"+config.Servers[i].Port] = types.State{
			Name:        "follower",
			ID:          i,
			CurrentTerm: 0,
			VotedFor:    -2,
			Log:         []types.LogData{},
			CommitIndex: -1,
			LastApplied: -1,
			NextIndex:   []int{},
			MatchIndex:  []int{},
		}
	}
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
				Config:         config,
				DesignationMap: designationMap,
			}
	}
	// RaftServers[config.Servers[0].IP+":"+config.Servers[0].Port] = types.RaftServer{"leader", config.Servers[0].IP, config.Servers[0].Port, designationMap}
	payload, err := json.Marshal(RaftServers)
	if err != nil {
		log.Printf("Can't Marshall to JSON in startSignal.go : %v\n", err)
		return err
	}
	return routing.ConcurrentReqRes(config, payload, "/startRaft", "-1")
}
