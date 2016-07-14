package dispatcher

import (
	"testing"
	"time"
)

func TestCheckEventValidity(t *testing.T) {
	cases := []struct {
		inputCode string
		message   string
		valid     bool
	}{
		{"ABC", "hello", true},
		{"ABC", "", true},
		{"AX", "hello", false},
		{"", "x", false},
		{"ABCD", "hello", false},
		{"abc", "hello", false},
		{"ABc", "", false},
		{"ABC ", "hello", false},
		{" ABC", "hello", false},
	}

	for i, aCase := range cases {
		evD := EventDummy{inputHandlerCode: aCase.inputCode, message: aCase.message}
		if chVal := checkEventValidity(&evD); chVal != aCase.valid {
			t.Errorf("Error in case %d. Expecting %v, got %v", i, aCase.valid, chVal)
		}
	}
}

type EventDummy struct {
	inputHandlerCode string
	message          string
	time             time.Time
	evType           int
}

func (evD *EventDummy) InputHandlerCode() string {
	return evD.inputHandlerCode
}

func (evD *EventDummy) EventMessage() string {
	return evD.message
}

func (evD *EventDummy) EventTime() time.Time {
	return evD.time
}

func (evD *EventDummy) EventType() int {
	return evD.evType
}
