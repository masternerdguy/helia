package universe

//Meta Structure for generic metadata
type Meta map[string]interface{}

//GetFloat64 Gets a float64 value from the metadata - returns a bool indicating whether it exists
func (m Meta) GetFloat64(key string) (float64, bool) {
	v, e := m[key]

	if !e {
		return 0, e
	}

	return v.(float64), e
}
