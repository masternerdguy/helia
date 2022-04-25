package shared

import (
	"sync"
	"time"
)

var teeLogChannel chan teeLog
var teeLogInitialized = false
var teeLogHandlers []logTeeHandler
var teeLogWrite sync.Mutex // intentionally not using sync.Mutex to avoid dying if logging is slow

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
	teeLogWrite = sync.Mutex{}

	// start watcher
	go func() {
		for {
			// wait for log
			log := <-teeLogChannel

			// pass to logger functions on separate goroutines
			for _, h := range teeLogHandlers {
				go func(h logTeeHandler) {
					// obtain write lock
					teeLogWrite.Lock()
					defer teeLogWrite.Unlock()

					// write message
					h(log.Message, log.EventTime)
				}(h)
			}

			// short sleep
			time.Sleep(time.Millisecond * 5)
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
