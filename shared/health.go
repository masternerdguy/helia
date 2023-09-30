package shared

import (
	"fmt"
	"sync"
)

const PHASE_STARTUP = "Starting up"
const PHASE_RUNNING = "System ready"
const PHASE_SHUTDOWN = "Shutting down"

var serverHealthMessage string
var serverHealthPhase string
var serverHealthLock sync.Mutex

// Updates the global server health message returned on ping
func SetServerHealth(phase string, msg string) {
	// obtain lock
	serverHealthLock.Lock()
	defer serverHealthLock.Unlock()

	// check for running phase
	if serverHealthPhase == PHASE_RUNNING {
		// overwrite whatever may be there
		msg = "Helia is running!"
	} else {
		// validate phase in range
		if serverHealthPhase != PHASE_STARTUP && serverHealthPhase != PHASE_RUNNING && serverHealthPhase != PHASE_SHUTDOWN {
			panic(fmt.Sprintf("Health phase out of range: %v", serverHealthPhase))
		}
	}

	// update message
	serverHealthPhase = phase
	serverHealthMessage = msg
}

// Gets the global server health status returned on ping
func GetServerHealth() (string, string) {
	// obtain lock
	serverHealthLock.Lock()
	defer serverHealthLock.Unlock()

	// check for running phase
	if serverHealthPhase == PHASE_RUNNING {
		// overwrite whatever may be there
		serverHealthMessage = "Helia is running!"
	} else {
		// validate phase in range
		if serverHealthPhase != PHASE_STARTUP && serverHealthPhase != PHASE_RUNNING && serverHealthPhase != PHASE_SHUTDOWN {
			panic(fmt.Sprintf("Health phase out of range: %v", serverHealthPhase))
		}
	}

	// return message
	return serverHealthPhase, serverHealthMessage
}
