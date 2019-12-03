package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type server struct {
	Name string
	IP   string
	Port string
}

// Configuration is the entire config file description
type Configuration struct {
	Servers []server
}

func main() {
	file, err := os.Open("server.config.json")
	if err != nil {
		log.Panic("Error reading from config file!")
	}
	decoder := json.NewDecoder(file)
	var configuration Configuration
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Panic("Error decoding config file!")
	}
	fmt.Println(configuration)
}
