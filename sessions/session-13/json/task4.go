package json

import (
	"encoding/json"
	"fmt"
	"os"
)

//Task 4: JSON Configuration File Reader
//
//Create a configuration JSON file, config.json, with fields such as app_name, version, debug_mode, and an array of services (e.g., ["service1", "service2"]).
//Write a program that reads the JSON file and prints out the configuration details in a formatted manner.

type Config struct {
	AppName   string   `json:"app_name"`
	Version   string   `json:"version"`
	DebugMode bool     `json:"debug_mode"`
	Services  []string `json:"services"`
}

func ReadConfig(path string) error {
	fmt.Println("#================================#")

	inputFile, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var config Config
	unmarshalError := json.Unmarshal(inputFile, &config)
	if unmarshalError != nil {
		return unmarshalError
	}

	fmt.Println("Configuration Details:")

	fmt.Println("- App Name:", config.AppName)
	fmt.Println("- Version:", config.Version)
	fmt.Println("- Debug Mode:", config.DebugMode)
	fmt.Println("- Services:", config.Services)

	return nil
}
