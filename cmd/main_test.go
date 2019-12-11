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
		t.Fatalf("Config file does not exist: %v",err)
	}
	//check for Empty config file
	if info.Size() == 0 {
		t.Fatalf("Config file is empty: %v",err)
	}
	//Check for opening of file
	file, err := os.Open("../server.config.json")
	if err != nil {
		t.Fatalf("Error opening file: %v",err)
	}
	decoder := json.NewDecoder(file)
	//check for Invalid config file
	configuration := make(map[int]types.Server)
	err = decoder.Decode(&configuration)
	if err != nil {
		t.Fatalf("Invalid config file : %v",err)
	}
	//check for invalid IP and port
	for _, v := range configuration {
		IP := v.IP
		_, err := regexp.MatchString(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`, IP)
		if err != nil {
			t.Fatalf("Invalid IP Address in config file : %v",err)
		}
		port := v.Port
		if _, err := strconv.Atoi(port); err != nil {
			t.Fatalf("Invalid Port Number in config file : %v", err)
		}
	}
}
