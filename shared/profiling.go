package shared

import (
	"log"
	"os"
	"runtime/debug"
	"sync"
)

var CpuProfileFile *os.File
var HeapProfileFile *os.File

type TaggedMutex struct {
	mutex sync.Mutex
}

func (m *TaggedMutex) Lock() {
	log.Printf("trying to get lock at %v", string(debug.Stack()))

	m.mutex.Lock()

	log.Printf("lock obtained at %v", string(debug.Stack()))
}

func (m *TaggedMutex) Unlock() {
	m.mutex.Unlock()
}
