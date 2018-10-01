package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	configPath := "/path"
	var cfgFile *os.File
	_, err := os.Stat(configPath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("could not get file info for config - %v\n", err)
		}
		cfgFile, err = os.Create(configPath)
		if err != nil {
			log.Fatalf("could not create config file - %v\n", err)
		}
	}
	var cfg config
	err = json.NewDecoder(cfgFile).Decode(&cfg)
	if err != nil {
		log.Fatalf("could not decode config file - %v\n", err)
	}
	fmt.Println(cfg)
}
