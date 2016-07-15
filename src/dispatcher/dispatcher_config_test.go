package dispatcher

import (
	"testing"
)

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
