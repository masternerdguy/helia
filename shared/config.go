package shared

import (
	"encoding/json"
	"errors"
	"os"
)

var config *sharedConfig

func InitializeConfiguration() {
	// load shared configuration
	c, err := loadConfiguration()

	if err != nil {
		panic("unable to load shared configuration")
	}

	config = &c
}

type sharedConfig struct {
	SendgridKey string `json:"sendgridKey"`
	FromEmail   string `json:"fromEmail"`
}

func loadConfiguration() (sharedConfig, error) {
	var config sharedConfig

	configFile, err := os.Open("shared-configuration.json")

	if err != nil {
		return config, errors.New("unable to load shared configuration")
	}

	defer configFile.Close()

	if err != nil {
		return sharedConfig{}, err
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config, nil
}
