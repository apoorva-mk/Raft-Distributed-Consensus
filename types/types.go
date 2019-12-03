package types

// Server describes a single server instance in the cluster
type Server struct {
	Name string
	IP   string
	Port string
}

// Configuration is the entire config file description
type Configuration struct {
	Servers []Server
}
