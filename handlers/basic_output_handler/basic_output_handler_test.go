package basic_output_handler

import (
	"os"
	"testing"
	"time"
)

func TestTakeInEvent(t *testing.T) {
	boh := BasicOutputHandler{DispatchLog: nil, OutputHandlerCode: "ABC", OutputWriter: os.Stdout}

	ev := &EventDummy{inputHandlerCode: "XYZ", message: "Hello world!"}
	boh.takeInEvent(ev)
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
