package main

import (
	"log"

	"gopkg.in/ini.v1"
)

func readConfig(file string) (string, int, string, string) {
	cfg, err := ini.Load(file)
	if err != nil {
		log.Fatalf("Error reading settings: %v", err)
	}

	apiToken := cfg.Section("Weatherflow").Key("Token").String()

	deviceID, err := cfg.Section("Weatherflow").Key("DeviceID").Int()
	if err != nil {
		log.Fatalf("Error converting device ID to int: %v", err)
	}

	displayName := cfg.Section("Weatherflow").Key("DisplayName").String()

	position := cfg.Section("Weatherflow").Key("Position").String()

	return apiToken, deviceID, displayName, position
}
