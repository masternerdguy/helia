package shared

// An event queue for a client
type eventQueue struct {
	Events []Event
}

// An event raised by a client
type Event struct {
	Type int
	Body interface{}
}
