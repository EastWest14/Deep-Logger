package dispatcher

import (
	"strings"
	"time"
)

type Event interface {
	InputHandlerCode() string
	EventMessage() string
	EventTime() time.Time
	EventType() int //will be an enumeration
}

func checkEventValidity(event Event) bool {
	return checkInputCodeValidity(event.InputHandlerCode())
}

func checkInputCodeValidity(inputCode string) bool {
	if len(inputCode) != 3 {
		return false
	}
	if strings.ToUpper(inputCode) != inputCode {
		return false
	}
	return true
}
