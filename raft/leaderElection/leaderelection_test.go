package leaderelection

import (
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
	"math/rand"
	"testing"
	"time"
)

func TestIncrementTermAndBecomeCandidate(t *testing.T) {
	config, _ := types.BuildConfigurationFromConfigFile("../../server.config.json")
	serverData := types.BuildServerData(config)

	server := config.Servers[0]
	host := server.IP + ":" + server.Port
	// Increment term
	state := incrementTermAndBecomeCandidate(serverData, host)

	if state.CurrentTerm != 1 {
		t.Error("Unexpected term after first increment:", state.CurrentTerm)
	}

	if state.Name != "candidate" {
		t.Error("Unexpected name after first increment:", state.Name)
	}

	serverData[host] = state

	// Increment term again
	state = incrementTermAndBecomeCandidate(serverData, host)

	if state.CurrentTerm != 2 {
		t.Error("Unexpected term after second increment:", state.CurrentTerm)
	}

	if state.Name != "candidate" {
		t.Error("Unexpected name after second increment:", state.Name)
	}

	// Increment term for an invalid IP
	// state := incrementTermAndBecomeCandidate(serverData, "0.0.0.1")

}

func TestGetTimer(t *testing.T) {
	minTimeout := 150
	maxTimeout := 300

	// Use constant seed to replicate results across tests
	rand.Seed(1)

	// Test if timer duration is within expected bounds
	timer, _ := getTimer(minTimeout, maxTimeout)
	timerDuration := getDuration(timer)

	if timerDuration < time.Duration(minTimeout)*time.Millisecond {
		t.Error("Timer runs shorter than expected")
	} else if timerDuration > time.Duration(maxTimeout)*time.Millisecond {
		t.Error("Timer runs longer than expected")
	}
}

func TestMajorityObtained(t *testing.T) {
	// Has five servers - Don't check for errors
	config, _ := types.BuildConfigurationFromConfigFile("../../server.config.json")

	votes := len(config.Servers)/2 + 1

	if !majorityObtained(votes, config) {
		t.Error("Majority incorrectly decided for ", votes)
	}

	votes = len(config.Servers)/2 - 1

	if majorityObtained(votes, config) {
		t.Error("Majority incorrectly decided for ", votes)
	}
}

func getDuration(t *time.Timer) time.Duration {
	startTime := time.Now()
	<-t.C
	endTime := time.Now()
	return endTime.Sub(startTime)
}
