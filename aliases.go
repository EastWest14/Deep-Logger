package deeplogger

import (
	"deeplogger/event"
)

type Event interface {
	event.Event
}

func NewEvent(message string) Event {
	return event.New(message)
}
