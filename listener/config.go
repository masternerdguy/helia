package listener

import (
	"os"
	"strconv"
)

type listenerConfig struct {
	ShutdownToken string
	Port          int
}

func loadConfiguration() (listenerConfig, error) {
	// get environment variables
	shutdownToken := os.Getenv("shutdownToken")
	port := os.Getenv("listenerPort")

	// parse port number
	portInt, err := strconv.ParseInt(port, 10, 32)

	return listenerConfig{
		ShutdownToken: shutdownToken,
		Port:          int(portInt),
	}, err
}
