package shared

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// A mutex that contains information about what it locks
type LabeledMutex struct {
	Structure    string
	UID          string
	lastLocked   int64
	lastUnlocked int64
	isLocked     bool
	mutex        sync.Mutex
}

func (m *LabeledMutex) Lock() {
	// obtain lock
	m.mutex.Lock()

	// store lock timestamp
	m.lastLocked = time.Now().UnixNano()
	m.isLocked = true

	// kill process if not labeled
	if len(m.Structure) == 0 || len(m.UID) == 0 {
		go func() {
			// give time for panic to print output
			time.Sleep(20 * time.Millisecond)
			os.Exit(0)
		}()

		panic(fmt.Sprintf("Unlabeled LabeledMutex! Make sure a label is assigned to every instance of the associated structure: %v", m))
	}

	// monitor for suspiciously long locks on a separate goroutine
	go func(m *LabeledMutex) {
		it := 0

		// wait ~5 seconds
		for {
			// exit goroutine if lock released
			if !m.isLocked {
				break
			}

			if it > 500 {
				// are we still locked?
				if m.lastLocked >= m.lastUnlocked {
					// this is a freeze - dump and exit
					go func() {
						// give time for panic to print output
						time.Sleep(20 * time.Millisecond)
						os.Exit(0)
					}()

					panic(fmt.Sprintf("Mutex slept for a very suspicious amount of time, this is likely a freeze: %v", m))
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
