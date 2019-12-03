package servermanagement

import (
	"log"
	"sync"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
	"github.com/gorilla/mux"
)

// StartServers starts all the servers based on the config file
func StartServers(config types.Configuration) {
	r := mux.NewRouter()
	// m := make(map[string]int)
	// r = routing.SetupRouting(r)
	// counter := 0
	wg := &sync.WaitGroup{}
	for i := range config.Servers {
		wg.Add(1)
		go func(i int) {
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
