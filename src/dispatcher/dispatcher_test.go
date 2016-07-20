package dispatcher

import (
	"testing"
	"time"
)

var sampleJSON = `
	{"name": "LALALA",
	"isOn": true,
	"inputHandlers": ["ABC", "XYZ"],
	"outputHandlers": ["WOW", "LOL"],
	"dispatchRules": [
		{"input":"ABC", "output": "WOW", "intersect": false},
		{"input":"XYZ", "output": "LOL", "intersect": true}
	]}
	`

func TestMatchOutputHandler(t *testing.T) {
	dl := DispatcherLog{*LoadConfigFromFile(sampleJSON)}
	ev := &EventDummy{inputHandlerCode: "ABC", message: ""}
	ok, outputH := dl.matchOutputHandler(ev)
	if !ok || string(outputH) != "WOW" {
		t.Error("Event routed incorrectly")
	}
}

func TestCheckEventValidity(t *testing.T) {
	cases := []struct {
		inputCode InputHandlerCode
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
	inputHandlerCode InputHandlerCode
	message          string
	time             time.Time
	evType           int
}

func (evD *EventDummy) InputHandlerCode() InputHandlerCode {
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

func TestRouteEvent(t *testing.T) {
	dl := DispatcherLog{*LoadConfigFromFile(sampleJSON)}
	ev := EventDummy{inputHandlerCode: "ABC", message: ""}
	outputHandlerEl := dl.routeEvent(&ev)
	if string(outputHandlerEl.code) != "WOW" {
		t.Error("Event routed incorrectly.")
	}
	ev.inputHandlerCode = InputHandlerCode("XYZ")
	outputHandlerEl = dl.routeEvent(&ev)
	if string(outputHandlerEl.code) != "LOL" {
		t.Error("Event routed incorrectly.")
	}
}
