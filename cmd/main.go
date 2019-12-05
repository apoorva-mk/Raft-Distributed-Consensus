package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/servermanagement"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

func main() {
	serverConfig := initializeConfiguration("server.config.json")

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go servermanagement.StartServers(serverConfig)
	go servermanagement.StartSignal(serverConfig, 1)
	wg.Wait()
}

func initializeConfiguration(filepath string) types.Configuration {
	file, err := os.Open(filepath)
	if err != nil {
		log.Panic("Error reading from config file: ", err.Error())
	}

	configuration := make(map[int]types.Server)
	err = json.NewDecoder(file).Decode(&configuration)
	if err != nil {
		log.Panic("Error decoding config file: ", err.Error())
	}

	return types.Configuration{configuration}
}
