package shared

import (
	"time"
)

var teeLogChannel chan teeLog
var teeLogInitialized = false
var teeLogHandlers []logTeeHandler

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

	// start watcher
	go func() {
		for {
			// wait for log
			log := <-teeLogChannel

			for _, h := range teeLogHandlers {
				// write message
				h(log.Message, log.EventTime)
			}

			// short sleep
			time.Sleep(20 * time.Millisecond)
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
