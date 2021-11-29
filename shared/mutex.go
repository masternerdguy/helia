package shared

import "sync"

// A mutex that contains information about what it locks
type LabeledMutex struct {
	Structure string
	UID       string
	mutex     sync.Mutex
}

func (m *LabeledMutex) Lock() {
	m.mutex.Lock()
}

func (m *LabeledMutex) Unlock() {
	m.mutex.Unlock()
}
