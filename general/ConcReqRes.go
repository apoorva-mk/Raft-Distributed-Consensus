package general

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
func ConcurrentReqRes(servers types.Configuration, payload []byte, endPoint string, serverID int) ([]types.URLResponse, error) {
	resCh := make(chan types.URLResponse, len(servers.Servers))
	wg := &sync.WaitGroup{}
	if serverID != -1 {
		wg.Add(len(servers.Servers) - 1)
	} else {
		wg.Add(len(servers.Servers))
	}
	for k, v := range servers.Servers {
		if k != serverID {
			URL := "http://" + v.IP + ":" + v.Port + endPoint
			go func() error {
				req, err := http.NewRequest("POST", URL, strings.NewReader(string(payload)))
				if err != nil {
					log.Printf("Bad request in startSignal.go : %v\n", err)
					return err
				}
				req.Header.Add("Content-Type", "application/json")

				res, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Printf("Bad response in ConcReqRes.go: %v\n", err)
					return err
				}
				resCh <- types.URLResponse{URL, res}
				// defer res.Body.Close()
				wg.Done()
				return nil
			}()
		}
	}
	var responses []types.URLResponse
	go func() { wg.Wait(); close(resCh) }()
	for x := range resCh {
		responses = append(responses, x)
	}
	return responses, nil
}
