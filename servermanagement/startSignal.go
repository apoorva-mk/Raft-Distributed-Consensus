package servermanagement

import (
	"encoding/json"
	"log"
	"time"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/servermanagement/routing"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

// StartSignal initiates the servers to behave
// according to the raft protocol.
func StartSignal(config types.Configuration, delay_factor time.Duration) error {
	payload, err := json.Marshal(config.Servers)
	if err != nil {
		log.Printf("Can't Marshall to JSON in startSignal.go : %v\n", err)
		return err
	}

	time.Sleep(delay_factor * time.Duration(len(config.Servers)) * 500 * time.Microsecond)

	err = routing.ConcurrentReqRes(config.Servers, payload, "/startRaft", -1)
	if err == nil {
		log.Print("Raft servers initiated")
	} else {
		log.Fatal("Unable to initiate all raft servers")
	}
	return err
}
