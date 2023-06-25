package shared

import "os"

var config *sharedConfig

// Invokes loadConfiguration() to cache the shared configuration
func InitializeConfiguration() {
	// load shared configuration
	c, err := loadConfiguration()

	if err != nil {
		panic("unable to load shared configuration")
	}

	config = &c
}

// Structure representing the shared configuration
type sharedConfig struct {
	SendgridKey string
	FromEmail   string
}

// Reads the shared configuration from environment variables
func loadConfiguration() (sharedConfig, error) {
	// read environment variables
	sendgridKey := os.Getenv("sendgridKey")
	fromEmail := os.Getenv("fromEmail")

	// return configuration
	return sharedConfig{
		SendgridKey: sendgridKey,
		FromEmail:   fromEmail,
	}, nil
}
