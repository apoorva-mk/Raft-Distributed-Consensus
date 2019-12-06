package servermanagement

import (
	"log"
	"sync"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/servermanagement/routing"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
	"github.com/gorilla/mux"
)

// StartServers starts all the servers based on the config file
func StartServers(config types.Configuration) {
	// since all servers in the clusters are the same,
	// they'll have the same routing setup. Only the IP
	// of the HTTP call will differ.
	r := mux.NewRouter()
	r = routing.SetupRouting(r)
	// each server is started parallely using a wait group
	// the wait group enables the parallel execution due to
	// the blocking nature of the servers
	wg := &sync.WaitGroup{}
	wg.Add(len(config.Servers))
	for i := range config.Servers {
		go func(i string) {
			serverData := server.Data{
				Router: r,
				IP:     config.Servers[i].IP,
				Port:   config.Servers[i].Port,
				HTTPS:  false,
			}
			err := server.Server(&serverData)
			if err != nil {
				log.Fatalf("Could not start server : %v", err)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
