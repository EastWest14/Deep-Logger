package newdispatcher

import (
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
	disp.AddOutputHandler("A")
	if _, ok := disp.outputHandlers["A"]; !ok {
		t.Error("Output handler not created correctly.")
	}
}
