package routing

import "github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"

// ReqVotesRequest represents a requestVote request structure
type ReqVotesRequest struct {
	Data     map[string]types.RaftServer
	ServerID int
}
