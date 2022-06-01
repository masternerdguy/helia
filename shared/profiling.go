package shared

import (
	"fmt"
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
	log.Println(
		fmt.Sprintf("trying to get lock at %v", string(debug.Stack())),
	)

	m.mutex.Lock()

	log.Println(
		fmt.Sprintf("lock obtained at %v", string(debug.Stack())),
	)
}

func (m *TaggedMutex) Unlock() {
	m.mutex.Unlock()
}
