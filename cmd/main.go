package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/servermanagement"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

func main() {
	file, err := os.Open("server.config.json")
	if err != nil {
		log.Panic("Error reading from config file!")
	}
	decoder := json.NewDecoder(file)
	var configuration types.Configuration
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Panic("Error decoding config file!")
	}
	RaftServers := make(map[string]types.RaftServer)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go servermanagement.StartServers(configuration)
	go func() {
		time.Sleep(1 * time.Millisecond)
		servermanagement.StartSignal(configuration, RaftServers)
		wg.Done()
	}()
	wg.Wait()
}
