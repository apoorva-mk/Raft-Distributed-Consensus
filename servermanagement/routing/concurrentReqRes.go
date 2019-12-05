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
func ConcurrentReqRes(servers types.Configuration, payload []byte, endPoint string, serverID string) error {
	wg := &sync.WaitGroup{}
	wg.Add(len(servers.Servers))
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
					log.Printf("Bad response in startSignal.go: %v\n", err)
					return err
				}
				defer res.Body.Close()
				wg.Done()
				return nil
			}()
		}
	}
	wg.Wait()
	return nil
}
