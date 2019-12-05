package types

// Server describes a single server instance in the cluster
type Server struct {
	ID   int    `json:"id"`
	IP   string `json:"ip"`
	Port string `json:"port"`
}

func (s *Server) URI() string {
	return s.IP + ":" + s.Port
}

func (s *Server) URL(protocol, endpoint string) string {
	return protocol + s.URI() + endpoint
}
