package shared

//eventQueue An event queue for a client
type eventQueue struct {
	Events []Event
}

//Event An event raised by a client
type Event struct {
	Type int
	Body interface{}
}
