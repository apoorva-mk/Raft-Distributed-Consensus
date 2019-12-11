package main

import (
	"encoding/json"
	"os"
	"regexp"
	"strconv"
	"testing"

	"github.com/SUMUKHA-PK/Raft-Distributed-Consensus/types"
)

func TestMain(t *testing.T) {
	//check if the config file exists
	info, err := os.Stat("../server.config.json")
	if os.IsNotExist(err) {
		t.Fatalf("Config file does not exist")
	}
	//check if there is no read permission to config file
	mask := info.Mode()
	if mask&(1<<2) == 0 {
		t.Fatalf("No read permission to config file")
	}
	//check for Empty config file
	if info.Size() == 0 {
		t.Fatalf("Config file is empty")
	}
	//check for Invalid config file
	file, err := os.Open("../server.config.json")
	decoder := json.NewDecoder(file)
	configuration := make(map[int]types.Server)
	err = decoder.Decode(&configuration)
	if err != nil {
		t.Errorf("Check if the config file is in the correct format")
		t.Fatalf("Invalid config file")
	}
	//check for invalid IP or Port
	for _, v := range configuration {
		IP := v.IP
		_, err := regexp.MatchString(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`, IP)
		if err != nil {
			t.Fatalf("Invalid IP Address in config file")
		}
		Port := v.Port
		port, err := strconv.Atoi(Port)
		if err != nil {
			t.Fatalf("Invalid Port")
		}
		if port < 1024 {
			t.Fatalf("Invalid Port number")
		}
	}
	main()
}
