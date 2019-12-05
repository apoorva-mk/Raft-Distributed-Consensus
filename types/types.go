package types

// Configuration is the entire config file description
type Configuration struct {
	Servers map[int]Server `json:"servers"`
}

// LogData is an instance of a single log
type LogData struct {
	Term    int    `json:"term"`
	Command string `json:"command"`
}
