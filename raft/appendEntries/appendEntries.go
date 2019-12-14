package appendentries

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/general"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// AppendEntries implements the append entries of raft
func AppendEntries(config types.Configuration, IP string) error {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() error {
		payload, err := json.Marshal(config)
		if err != nil {
			log.Printf("Can't Marshall to JSON in requestVotes.go : %v\n", err)
			return err
		}
		_, err = general.ConcurrentReqRes(config, []byte(payload), "/appendEntries", types.ServerData[IP].ID)
		if err != nil {
			fmt.Println(err)
			return err
		}
		wg.Done()
		return nil
	}()
	wg.Wait()
	return nil
}
