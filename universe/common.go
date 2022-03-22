package universe

// Approximate value of Number.Epsilon in JS
const Epsilon float64 = 2.2204460492503130808472633361816e-16

// Structure for generic metadata
type Meta map[string]interface{}

// Gets a float64 value from the metadata - returns a bool indicating whether it exists
func (m Meta) GetFloat64(key string) (float64, bool) {
	v, e := m[key]

	if !e {
		return 0, e
	}

	return v.(float64), e
}

// Gets an int value from the metadata - returns a bool indicating whether it exists
func (m Meta) GetInt(key string) (int, bool) {
	v, e := m[key]

	if !e {
		return 0, e
	}

	return int(v.(float64)), e
}

// Gets an int value from the metadata - returns a bool indicating whether it exists (special case where int is not a float under the hood)
func (m Meta) GetPureInt(key string) (int, bool) {
	v, e := m[key]

	if !e {
		return 0, e
	}

	return v.(int), e
}

// Gets a bool value from the metadata - returns a bool indicating whether it exists
func (m Meta) GetBool(key string) (bool, bool) {
	v, e := m[key]

	if !e {
		return false, e
	}

	return v.(bool), e
}

// Gets a string value from the metadata - returns a bool indicating whether it exists
func (m Meta) GetString(key string) (string, bool) {
	v, e := m[key]

	if !e {
		return "", e
	}

	return v.(string), e
}

// Gets a map value from the metadata - returns a bool indicating whether it exists
func (m Meta) GetMap(key string) (Meta, bool) {
	v, e := m[key]

	if !e {
		return nil, e
	}

	return Meta(v.(map[string]interface{})), e
}

// Alias for a generic interface for any type
type Any interface{}
