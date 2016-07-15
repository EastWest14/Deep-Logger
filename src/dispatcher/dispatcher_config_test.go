package dispatcher

import (
	"testing"
)

func TestConfigSetName(t *testing.T) {
	dc := dispatcherConfig{}
	dc.SetName("ABC")
	if dc.name != "ABC" {
		t.Error("DispatchConfig name not setting correctly.")
	}
}

func TestConfigName(t *testing.T) {
	dc := dispatcherConfig{}
	dc.name = "ABC"
	if dc.Name() != "ABC" {
		t.Error("DispatchConfig name not being read correctly.")
	}
}

func TestConfigSetIsOn(t *testing.T) {
	dc := dispatcherConfig{}
	dc.isOn = false
	dc.SetIsOn(true)
	if dc.isOn != true {
		t.Error("DispatchConfig not turning on.")
	}
	dc.isOn = true
	dc.SetIsOn(false)
	if dc.isOn != false {
		t.Error("DispatchConfig not turning off.")
	}
}

func TestIsOn(t *testing.T) {
	dc := dispatcherConfig{}
	dc.isOn = false
	if dc.isOn != false {
		t.Error("DispatchConfig read method returns on, should be off.")
	}
	dc.isOn = true
	if dc.isOn != true {
		t.Error("DispatchConfig read method returns off, should be on.")
	}
}

//TODO: write tests for checking internal consistency
