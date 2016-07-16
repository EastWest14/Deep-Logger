package dispatcher

import (
	"time"
)

type DispatcherLog struct {
	dispatcherConfig
}

type Event interface {
	InputHandlerCode() InputHandlerCode
	EventMessage() string
	EventTime() time.Time
	EventType() int //will be an enumeration
}

func checkEventValidity(event Event) bool {
	return checkInputCodeValidity(event.InputHandlerCode())
}
