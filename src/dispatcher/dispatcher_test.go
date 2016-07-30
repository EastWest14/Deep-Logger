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

func TestInputEvent(t *testing.T) {
	dl := DispatcherLog{*LoadConfig(sampleJSON)}

	v := false
	dummyF := func(ev Event) {
		v = true
		return
	}
	err := dl.RegisterOutputHandler(OutputHandlerCode("LOL"), dummyF)
	if err != nil {
		t.Error(err.Error())
	}

	ev := &EventDummy{inputHandlerCode: "XYZ", message: ""}
	dl.InputEvent(ev)
	if v != true {
		t.Error("Event routed incorrectly.")
	}
}

func TestMatchOutputHandler(t *testing.T) {
	dl := DispatcherLog{*LoadConfig(sampleJSON)}
	ev := &EventDummy{inputHandlerCode: "ABC", message: ""}
	ok, outputH := dl.matchOutputHandler(ev)
	if !ok || string(outputH.code) != "WOW" {
		t.Error("Event routed incorrectly.")
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
	dl := DispatcherLog{*LoadConfig(sampleJSON)}
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

func TestRegisterOutputHandler(t *testing.T) {
	dc := ConfigFromFile("../../config/little_config.json")
	dl := DispatcherLog{dispatcherConfig: *dc}
	v := false
	dummyF := func(ev Event) {
		v = true
		return
	}
	err := dl.RegisterOutputHandler(OutputHandlerCode("YYY"), dummyF)
	if err != nil {
		t.Error(err.Error())
	}

	evD := EventDummy{inputHandlerCode: InputHandlerCode("XYZ"), message: ""}
	dl.InputEvent(&evD)
	if v != true {
		t.Error("Output handler registration did not lead to correct routing.")
	}
}
