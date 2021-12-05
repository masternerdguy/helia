package shared

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

var MutexFreeze bool
var ShutdownSignal bool

// A mutex that contains information about what it locks
type LabeledMutex struct {
	Structure       string
	UID             string
	lastCaller      string
	lastCallerStack string // only captured if in "aggressive" mode due to performance penalty
	lastLocker      string
	lastLockerStack string // only captured if in "aggressive" mode due to performance penalty
	lastLocked      int64
	lastUnlocked    int64
	isLocked        bool
	mutex           sync.Mutex
	lastCallerMutex sync.Mutex
	aggressiveMode  bool // if true, performance-intensive logging will be performed
}

// When set to true, performance intensive logging (eg: call stack) will be performed
func (m *LabeledMutex) SetAggressiveFlag(f bool) {
	// obtain lock
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// set flag
	m.aggressiveMode = f
}

func (m *LabeledMutex) Lock(caller string) {
	// store most recent caller (this will likely be the one causing the freeze)
	m.lastCallerMutex.Lock()
	m.lastCaller = caller

	if m.aggressiveMode {
		m.lastCallerStack = string(debug.Stack())
	}

	m.lastCallerMutex.Unlock()

	// obtain lock
	m.mutex.Lock()

	// store lock timestamp
	m.lastLocked = time.Now().UnixNano()
	m.isLocked = true

	// store last locker
	m.lastLocker = caller

	if m.aggressiveMode {
		m.lastLockerStack = string(debug.Stack())
	}

	// kill process if not labeled
	if len(m.Structure) == 0 || len(m.UID) == 0 {
		go func() {
			// give time for panic to print output
			time.Sleep(20 * time.Millisecond)
			os.Exit(2)
		}()

		panic(fmt.Sprintf("Unlabeled LabeledMutex! Make sure a label is assigned to every instance of the associated structure: %v", m))
	}

	// monitor for suspiciously long locks on a separate goroutine
	go func(m *LabeledMutex) {
		it := 0

		// wait ~5 seconds
		for {
			// exit goroutine if shutting down
			if ShutdownSignal {
				break
			}

			// exit goroutine if lock released
			if !m.isLocked {
				break
			}

			// exit goroutine if global freeze declared
			if MutexFreeze {
				break
			}

			if it > 500 {
				// are we still locked?
				if m.lastLocked >= m.lastUnlocked {
					MutexFreeze = true

					// this is a freeze - core will save the world state and shut down the system
					go func() {
						time.Sleep(10 * time.Second)
						log.Println(fmt.Sprintf("Mutex locked for a very suspicious amount of time, this was almost certainly a freeze: %v", m))
					}()

					log.Println(fmt.Sprintf("! Emergency shutdown - deadlock detected: %v", m))
					return
				} else {
					// lock released!
					break
				}
			}

			// sleep in small increments to avoid pegging cpu
			time.Sleep(10 * time.Millisecond)
			it++
		}
	}(m)
}

func (m *LabeledMutex) Unlock() {
	if !m.isLocked {
		go func() {
			// give time for panic to print output
			time.Sleep(20 * time.Millisecond)
			os.Exit(0)
		}()

		panic(fmt.Sprintf("Attempted to unlock a mutex that is already unlocked: %v", m))
	}

	// store unlock timestamp
	m.lastUnlocked = time.Now().UnixNano()
	m.isLocked = false

	// release lock
	m.mutex.Unlock()
}
