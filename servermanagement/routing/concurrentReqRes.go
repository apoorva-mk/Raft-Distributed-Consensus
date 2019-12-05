package routing

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// ConcurrentReqRes triggers the end points of multiple
// servers. Takes in the config and payload.
// serverID ensures that the request isn't forwarded to
// itself; parameter set to -1 to forward to all.
func ConcurrentReqRes(servers map[int]types.Server, payload []byte, endpoint string, serverID int) error {
	wg := &sync.WaitGroup{}
	wg.Add(len(servers))

	for id, server := range servers {
		if id != serverID {
			url := server.URL("http://", endpoint)
			go func() error {
				return handleRequestResponseCycle(url, payload)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	return nil
}

func handleRequestResponseCycle(URL string, payload []byte) error {
	req, err := http.NewRequest("POST", URL, strings.NewReader(string(payload)))

	if err != nil {
		log.Printf("Bad request created in concurrentReqRes.go : %v\n", err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Bad response recieved in concurrentReqRes.go: %v\n", err)
		return err
	}

	defer res.Body.Close()
	return nil
}
