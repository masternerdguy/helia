package shared

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

var teeLogChannel chan teeLog
var teeLogInitialized = false
var teeLogHandlers []logTeeHandler
var teeLogWrite LabeledMutex

type logTeeHandler func(string, time.Time)

type teeLog struct {
	Message   string
	EventTime time.Time
}

// Prepares shared tee logger
func InitializeTeeLog(fns ...logTeeHandler) {
	if teeLogInitialized {
		panic("logtee is already initialized!")
	}

	// store logger functions
	teeLogHandlers = fns

	// initialize channel
	teeLogChannel = make(chan teeLog)

	// initialize mutex
	teeLogWrite = LabeledMutex{
		Structure: "PlayerReputationSheet",
		UID:       fmt.Sprintf("%v :: %v :: %v", uuid.New(), time.Now(), rand.Float64()),
	}

	// start watcher
	go func() {
		for {
			// wait for log
			log := <-teeLogChannel

			// detach log handling
			go func() {
				// pass to logger functions on separate goroutines
				for _, h := range teeLogHandlers {
					go func(h logTeeHandler) {
						// obtain write lock
						teeLogWrite.Lock("InitializeTeeLog::watcher")
						defer teeLogWrite.Unlock()

						// write message
						h(log.Message, log.EventTime)
					}(h)
				}
			}()
		}
	}()

	// mark as initialized
	teeLogInitialized = true
}

// Queues a log string to be handled by the initialized logger functions
func TeeLog(log string) {
	if !teeLogInitialized {
		panic("logtee is not initialized!")
	}

	// get timestamp
	now := time.Now()

	// send log to channel
	teeLogChannel <- teeLog{
		Message:   log,
		EventTime: now,
	}
}
