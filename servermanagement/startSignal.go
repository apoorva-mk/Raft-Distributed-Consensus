package servermanagement

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// StartSignal initiates the servers to behave
// according to the raft protocol.
func StartSignal(config types.Configuration, RaftServers map[string]types.RaftServer) error {
	wg := &sync.WaitGroup{}
	wg.Add(len(config.Servers))
	designationMap := make(map[string]string)
	for i := range config.Servers {
		designationMap[config.Servers[i].IP+":"+config.Servers[i].Port] = "follower"
	}
	for i := range config.Servers {
		RaftServers[config.Servers[i].IP+":"+config.Servers[i].Port] = types.RaftServer{"follower", config.Servers[i].IP, config.Servers[i].Port, designationMap}
	}
	payload, err := json.Marshal(RaftServers)
	if err != nil {
		log.Printf("Can't Marshall to JSON in startSignal.go : %v\n", err)
		return err
	}
	for i := range config.Servers {
		URL := "http://" + config.Servers[i].IP + ":" + config.Servers[i].Port + "/startRaft"
		go func(i int) error {
			req, err := http.NewRequest("POST", URL, strings.NewReader(string(payload))) //strings.NewReader(string(payload)))
			if err != nil {
				log.Printf("Bad request in startSignal.go : %v\n", err)
				return err
			}
			req.Header.Add("Content-Type", "application/json")

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("Bad response in startSignal.go: %v\n", err)
				return err
			}
			defer res.Body.Close()
			wg.Done()
			return nil
		}(i)
	}
	wg.Wait()
	return nil
}
