package newdispatcher

import (
	"deeplogger/event"
	"testing"
)

func TestNewWithName(t *testing.T) {
	const name = "D1"
	disp := NewWithName(name)
	if disp.name != name {
		t.Errorf("Expected name %s, got %s", name, disp.name)
	}
}

func TestAddInputHandler(t *testing.T) {
	disp := NewWithName("Test")
	disp.AddInputHandler("A", true)
	disp.AddInputHandler("B", false)
	if disp.inputHandlers["A"] != true {
		t.Error("Input hanlder state set incorrectly.")
	}
	if disp.inputHandlers["B"] == true {
		t.Error("Input hanlder state set incorrectly.")
	}
	//TODO: more complete testing.
}

func TestHasInputHandler(t *testing.T) {
	disp := NewWithName("Test")
	if exists, isOn := disp.HasInputHandler("nope"); exists != false || isOn != false {
		t.Error("HasInputHandler returns false positive or isOn is true for non-existing handler.")
	}
	disp.AddInputHandler("Test", false)
	if exists, isOn := disp.HasInputHandler("Test"); exists != true || isOn != false {
		t.Error("HasInputHandler returns false answer or isOn is true.")
	}
	disp.AddInputHandler("Test2", true)
	if exists, isOn := disp.HasInputHandler("Test2"); exists != true || isOn == false {
		t.Error("HasInputHandler returns false answer or isOn is false.")
	}
}

func TestAddOutputHandler(t *testing.T) {
	disp := NewWithName("Test")
	disp.AddOutputHandler("A", func(ev event.Event) {
	})
	if _, funcVar := disp.outputHandlers["A"]; !funcVar {
		t.Error("Output handler not created correctly.")
	}
}

func TestInputEvent(t *testing.T) {
	disp := NewWithName("Test")
	disp.AddInputHandler("input1", true)
	disp.AddInputHandler("input2", false)
	hit := false
	disp.AddOutputHandler("output1", func(ev event.Event) {
		hit = true
	})
	dr := NewRule(NewMatchCondWithName("input1"), "output1")
	disp.AddDispatchRule(dr)
	dr = NewRule(NewMatchCondWithName("input2"), "output1")
	disp.AddDispatchRule(dr)

	ev := event.New("")
	ev.SetInputHandlerName("input1")
	disp.InputEvent(ev)
	if !hit {
		t.Error("Event didn't route.")
	}

	hit = false
	ev.SetInputHandlerName("input2")
	disp.InputEvent(ev)
	if !hit {
		t.Error("Event didn't route.")
	}
}

func TestMatchesEvent(t *testing.T) {
	dr := NewRule(NewMatchCondWithName("input"), "output")
	ev := event.New("")
	ev.SetInputHandlerName("input")
	ev2 := event.New("")
	ev2.SetInputHandlerName("wrong")
	match := dr.matchesEvent(ev)
	if match != true {
		t.Error("Rule doesn't match, should match.")
	}
	match = dr.matchesEvent(ev2)
	if match == true {
		t.Error("Rule does match, shouldn't match.")
	}
}
