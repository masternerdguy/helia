package shared

import "sync"

var serverHealthMessage string
var serverHealthPhase string
var serverHealthLock sync.Mutex

// Updates the global server health message returned on ping
func SetServerHealth(phase string, msg string) {
	// obtain lock
	serverHealthLock.Lock()
	defer serverHealthLock.Unlock()

	// update message
	serverHealthPhase = phase
	serverHealthMessage = msg
}

// Gets the global server health status returned on ping
func GetServerHealth() (string, string) {
	// obtain lock
	serverHealthLock.Lock()
	defer serverHealthLock.Unlock()

	// return message
	return serverHealthPhase, serverHealthMessage
}
