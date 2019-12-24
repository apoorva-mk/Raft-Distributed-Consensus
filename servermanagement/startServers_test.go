package servermanagement

import (
	"encoding/json"
	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
	"github.com/mozilla/mig/modules/netstat"
	"log"
	"os"
	"testing"
)

func TestStartServer(t *testing.T) {
	file, _ := os.Open("../server.config.json")
	decoder := json.NewDecoder(file)
	configuration := make(map[int]types.Server)
	_ = decoder.Decode(&configuration)
	// Iterating through the list of servers and ensuring
	// that specified IP addresses and ports are available
	var IP, port string
	for _, v := range configuration {
		IP = v.IP
		port = v.Port
		foundIP, elementIP, errIP := netstat.HasIPConnected(IP)
		foundPort, elementPort, errPort := netstat.HasListeningPort(port)
		if foundIP == true {
			t.Errorf("IP address %s already in use", IP)
		}
		if foundPort == true {
			t.Errorf("Port %s already in use", port)
		}
		log.Println(foundIP, elementIP, errIP)
		log.Println(foundPort, elementPort, errPort)
	}
	// Creating a test configuration with a single server
	// to ensure that start servers starts a server at the given IP and port
	testConfiguration := make(map[int]types.Server)
	testConfiguration[0] = types.Server{
		IP:   IP,
		Port: port,
	}
	go StartServers(types.Configuration{testConfiguration})
	foundIP, _, _ := netstat.HasIPConnected(IP)
	foundPort, _, _ := netstat.HasListeningPort(port)
	// TODO : fix the issue of why IP is shown as not being used.
	if foundPort == false {
		t.Errorf("Error in StartServer.go : unable to start server with given IP and port number")
		log.Println(foundIP, foundPort)
	}
}
