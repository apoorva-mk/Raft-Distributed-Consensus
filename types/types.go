package types

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"
)

// ServerData is the per server data that contains
// the most updated data from each server. This is
// maintained for its own needs. Maps from the
// servers IP to its state.
var ServerData map[string]*State = make(map[string]*State)

// Server describes a single server instance in the cluster
type Server struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

// Configuration is the entire config file description
type Configuration struct {
	Servers map[int]Server `json:"servers"`
}

// LogData is an instance of a single log
type LogData struct {
	Term    int    `json:"term"`
	Command string `json:"command"`
}

// State has all the data on the
// servers state.
type State struct {
	// can be follower, leader or candidate
	// all servers start as a follower, if they
	// dont hear from a leader, they can become
	// candidates. Leaders are elected from the
	// leader election process.
	Name        string `json:"name"`
	ID          int    `json:"ID"`
	CurrentTerm int    `json:"currentTerm"`
	// VotedFor maintains the ID of the voted
	// server; -1 if its leader, -2 at init
	VotedFor int `json:"votedFor"`
	// Log is the command received by the leader.
	// each entry contains the term and the command.
	Log []LogData `json:"log"`
	// above 4 variables are persistent in the server
	// CommitIndex maintains the highest log entry
	// that is known to be committed.
	CommitIndex int `json:"commitIndex"`
	// LastApplied is the highest log entry applied
	// to the state machine
	LastApplied int `json:"lastApplied"`
	// above 2 variables are volatile on all servers
	// NextIndex maintains a list of the next log
	// entry to be sent to the followers.
	NextIndex []int `json:"nextIndex"`
	// MatchIndex maintains the highest log entry
	// that is known to be replicated on the server
	MatchIndex []int `json:"matchIndex"`
	// above 2 variables are volatile only int the
	// leader and for each follower. Its also
	// re-init after each election.
	Lock sync.Mutex
}

// RaftServer describes a single raft server
type RaftServer struct {
	ServerState State         `json:"serverState"`
	Config      Configuration `json:"config"`
}

// URLResponse facilitates responses for ConcurrentReqRes
type URLResponse struct {
	URL string         `json:"URL"`
	Res *http.Response `json:"res"`
}

// ReqVotesRequest represents a requestVote request structure
type ReqVotesRequest struct {
	Data         Configuration `json:"data"`
	Term         int           `json:"term"`
	CandidateID  int           `json:"candidateID"`
	LastLogIndex int           `json:"lastLogIndex"`
	LastLogTerm  int           `json:"lastLogTerm"`
}

// AppendEntriesReq represents an AppendEntires request
type AppendEntriesReq struct {
	Data     LogData
	ServerID int
}

// BuildConfigurationFromConfigFile builds a Configuration using filepath
func BuildConfigurationFromConfigFile(filepath string) (Configuration, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return Configuration{nil}, err
	}

	servers := make(map[int]Server)

	err = json.NewDecoder(file).Decode(&servers)
	if err != nil {
		return Configuration{nil}, err
	}
	return Configuration{servers}, nil
}

// BuildServerData maps IP to Current State of Server
func BuildServerData(config Configuration) map[string]*State {
	serverData := make(map[string]*State)
	for i, server := range config.Servers {
		serverData[server.IP+":"+server.Port] = &State{
			Name:        "follower",
			ID:          i,
			CurrentTerm: 0,
			VotedFor:    -2,
			Log:         []LogData{},
			CommitIndex: -1,
			LastApplied: -1,
			NextIndex:   []int{},
			MatchIndex:  []int{},
		}
	}
	return serverData
}
