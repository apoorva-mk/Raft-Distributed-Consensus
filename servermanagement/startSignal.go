package servermanagement

import (
	"log"
	"net/http"
	"sync"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// StartSignal initiates the servers to behave
// according to the raft protocol.
func StartSignal(config types.Configuration) error {
	wg := &sync.WaitGroup{}
	wg.Add(len(config.Servers))
	for i := range config.Servers {
		URL := "http://" + config.Servers[i].IP + ":" + config.Servers[i].Port + "/startRaft"
		go func(i int) error {
			req, err := http.NewRequest("GET", URL, nil) //strings.NewReader(string(payload)))
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
