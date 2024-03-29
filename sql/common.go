package sql

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Structure for generic JSON metadata
type Meta map[string]interface{}

// Converts generic metadata into JSON
func (a Meta) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Converts JSON into generic metadata
func (a *Meta) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
