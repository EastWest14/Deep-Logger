package dispatcher

import (
	"deeplogger/simpleevent"
	"testing"
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
	dl := DispatcherLog{configFromString(sampleJSON)}

	v := false
	dummyF := func(ev Event) {
		v = true
		return
	}
	err := dl.RegisterOutputHandler("LOL", dummyF)
	if err != nil {
		t.Error(err.Error())
	}

	ev := simpleevent.New("")
	ev.SetInputHandlerName("XYZ")
	dl.InputEvent(ev)
	if v != true {
		t.Error("Event routed incorrectly.")
	}
}

func TestMatchOutputHandler(t *testing.T) {
	dl := DispatcherLog{configFromString(sampleJSON)}
	ev := simpleevent.New("")
	ev.SetInputHandlerName("ABC")
	ok, outputH := dl.matchOutputHandler(ev)
	if !ok || outputH.name != "WOW" {
		t.Error("Event routed incorrectly.")
	}
}

func TestCheckEventValidity(t *testing.T) {
	cases := []struct {
		inputName string
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
		ev := simpleevent.New(aCase.message)
		ev.SetInputHandlerName(aCase.inputName)
		if chVal := checkEventValidity(ev); chVal != aCase.valid {
			t.Errorf("Error in case %d. Expecting %v, got %v", i, aCase.valid, chVal)
		}
	}
}

func TestRouteEvent(t *testing.T) {
	dl := DispatcherLog{configFromString(sampleJSON)}
	ev := simpleevent.New("")
	ev.SetInputHandlerName("ABC")
	outputHandlerEl := dl.routeEvent(ev)
	if outputHandlerEl.name != "WOW" {
		t.Error("Event routed incorrectly.")
	}
	ev = simpleevent.New("")
	ev.SetInputHandlerName("XYZ")
	outputHandlerEl = dl.routeEvent(ev)
	if outputHandlerEl.name != "LOL" {
		t.Error("Event routed incorrectly.")
	}
}

func TestRegisterOutputHandler(t *testing.T) {
	dc := configFromFile("../config/little_config.json")
	dl := DispatcherLog{dispatcherConfig: dc}
	v := false
	dummyF := func(ev Event) {
		v = true
		return
	}
	err := dl.RegisterOutputHandler("YYY", dummyF)
	if err != nil {
		t.Error(err.Error())
	}

	ev := simpleevent.New("")
	ev.SetInputHandlerName("XYZ")
	dl.InputEvent(ev)
	if v != true {
		t.Error("Output handler registration did not lead to correct routing.")
	}
}
