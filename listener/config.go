package listener

import (
	"encoding/json"
	"os"
)

type listenerConfig struct {
	ShutdownToken string `json:"shutdownToken"`
}

func loadConfiguration() (listenerConfig, error) {
	var config listenerConfig

	configFile, err := os.Open("listener-configuration.json")
	defer configFile.Close()

	if err != nil {
		return listenerConfig{}, err
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config, nil
}
