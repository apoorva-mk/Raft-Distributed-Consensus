package servermanagement

import (
	"log"
	"sync"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/servermanagement/routing"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
	"github.com/gorilla/mux"
)

// StartServers starts the servers based on the config file
// All Servers will have same IP (i.e local host) and can share router
// Servers will differ in the port number
func StartServers(config types.Configuration) {
	r := mux.NewRouter()
	r = routing.SetupRouting(r)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	for _, s := range config.Servers {
		go startServer(r, s)
	}
	wg.Wait()
}

func startServer(r *mux.Router, s types.Server) {
	serverData := server.Data{
		Router: r,
		IP:     s.IP,
		Port:   s.Port,
		HTTPS:  false,
	}

	err := server.Server(&serverData)

	if err != nil {
		log.Fatalf("Could not start server %v: %v", s.URI(), err)
	}
}
