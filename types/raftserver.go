package types

// RaftServer describes a single raft server
type RaftServer struct {
	SelfState State         `json:"serverState"`
	Config    Configuration `json:"config"`
	// the designation map holds the mapping to
	// the IP of the server to the state of the
	// server in the cluster. This is useful to
	// send out data to necessary servers.
	StatesByURI map[string]State `json:"statesByURI"`
}

func BuildRaftServer(index int, config Configuration, statesByURI map[string]State) RaftServer {
	return RaftServer{
		SelfState:   BuildInitialState(),
		Config:      config,
		StatesByURI: statesByURI,
	}
}
