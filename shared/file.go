package shared

import (
	"io/ioutil"
	"os"
)

// Reads all bytes from a file and returns a bool indicating whether or not it was read successfully
func ReadFileBytes(f string) (*[]byte, bool) {
	// open file
	fo, err := os.Open(f)

	if err != nil {
		return nil, false
	}

	// defer close
	defer fo.Close()

	// read contents
	data, err := ioutil.ReadAll(fo)

	if err != nil {
		return nil, false
	}

	// return result
	return &data, true
}
