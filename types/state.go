package types

// State has all the data on the
// servers state.
type State struct {
	// can be follower, leader or candidate
	// all servers start as a follower, if they
	// dont hear from a leader, they can become
	// candidates. Leaders are elected from the
	// leader election process.
	CurrentTerm int `json:"currentTerm"`
	// Server.ID that received vote during current term. 0 if none
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
}

func BuildInitialState() State {
	return State{
		CurrentTerm: 0,
		VotedFor:    0,
		Log:         []LogData{},
		CommitIndex: -1,
		LastApplied: -1,
		NextIndex:   []int{},
		MatchIndex:  []int{},
	}
}
