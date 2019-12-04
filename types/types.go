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

// RaftServer describes a single raft server
type RaftServer struct {
	// can be follower, leader or candidate
	// all servers start as a follower, if they
	// dont hear from a leader, they can become
	// candidates. Leaders are elected from the
	// leader election process.
	State string `json:"state"`
	IP    string `json:"IP"`
	Port  string `json:"Port"`
	// the designation map holds the mapping to
	// the IP of the server to the state of the
	// server in the cluster. This is useful to
	// send out data to necessary servers.
	DesignationMap map[string]string `json:"designationMap"`
}
