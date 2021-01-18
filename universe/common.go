package universe

//Epsilon Approximate value of Number.Epsilon in JS
const Epsilon float64 = 2.2204460492503130808472633361816e-16

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

//GetInt Gets an int value from the metadata - returns a bool indicating whether it exists
func (m Meta) GetInt(key string) (int, bool) {
	v, e := m[key]

	if !e {
		return 0, e
	}

	return int(v.(float64)), e
}

//GetBool Gets a bool value from the metadata - returns a bool indicating whether it exists
func (m Meta) GetBool(key string) (bool, bool) {
	v, e := m[key]

	if !e {
		return false, e
	}

	return v.(bool), e
}

//GetString Gets a string value from the metadata - returns a bool indicating whether it exists
func (m Meta) GetString(key string) (string, bool) {
	v, e := m[key]

	if !e {
		return "", e
	}

	return v.(string), e
}

//Any Alias for a generic interface for any type
type Any interface{}
