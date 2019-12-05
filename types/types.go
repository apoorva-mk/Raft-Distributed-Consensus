package types

// Server describes a single server instance in the cluster
type Server struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port string `json:"port"`
}

// Configuration is the entire config file description
type Configuration struct {
	Servers []Server `json:"servers"`
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
	Name        string
	CurrentTerm int
	// VotedFor maintains the ID of the voted
	// server; -1 if its leader, -2 at init
	VotedFor int
	// Log is the command received by the leader.
	// each entry contains the term and the command.
	Log []LogData
	// above 4 variables are persistent in the server
	// CommitIndex maintains the highest log entry
	// that is known to be committed.
	CommitIndex int
	// LastApplied is the highest log entry applied
	// to the state machine
	LastApplied int
	// above 2 variables are volatile on all servers
	// NextIndex maintains a list of the next log
	// entry to be sent to the followers.
	NextIndex []int
	// MatchIndex maintains the highest log entry
	// that is known to be replicated on the server
	MatchIndex []int
	// above 2 variables are volatile only int the
	// leader and for each follower. Its also
	// re-init after each election.
}

// RaftServer describes a single raft server
type RaftServer struct {
	ServerState State  `json:"serverState"`
	ID          int    `json:"id"`
	IP          string `json:"IP"`
	Port        string `json:"Port"`
	// the designation map holds the mapping to
	// the IP of the server to the state of the
	// server in the cluster. This is useful to
	// send out data to necessary servers.
	DesignationMap map[string]State `json:"designationMap"`
}
