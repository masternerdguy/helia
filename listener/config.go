package listener

import (
	"encoding/json"
	"errors"
	"os"
)

type listenerConfig struct {
	ShutdownToken string `json:"shutdownToken"`
	Port          int    `json:"port"`
	AzureHacks    bool   `json:"azureHacks"`
}

func loadConfiguration() (listenerConfig, error) {
	var config listenerConfig

	configFile, err := os.Open("listener-configuration.json")

	if err != nil {
		return config, errors.New("unable to load listener configuration")
	}

	defer configFile.Close()

	if err != nil {
		return listenerConfig{}, err
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config, nil
}
