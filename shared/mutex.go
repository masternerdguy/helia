package shared

import (
	"fmt"
	"sync"
)

// A mutex that contains information about what it locks
type LabeledMutex struct {
	Structure string
	UID       string
	mutex     sync.Mutex
}

func (m *LabeledMutex) Lock() {
	m.mutex.Lock()

	// panic if not labeled
	if len(m.Structure) == 0 || len(m.UID) == 0 {
		panic(fmt.Sprintf("Unlabeled LabeledMutex! Make sure a label is assigned to every instance of the associated structure: %v", m))
	}
}

func (m *LabeledMutex) Unlock() {
	m.mutex.Unlock()
}
