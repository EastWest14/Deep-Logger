package dispatcher

import (
	"testing"
)

func TestConfigUnmarshal(t *testing.T) {
	caseZero := `
	{"name": "LALALA",
	"isOn": true,
	"inputHandlers": ["ABC", "XYZ"],
	"outputHandlers": ["WOW", "LOL"],
	"dispatchRules": 5}
	`
	dc := LoadConfigFromFile(caseZero)
	if dc.Name() != "LALALA" {
		t.Error("DispatchConfig name not being read from JSON correctly.")
	}
	if dc.isOn != true {
		t.Error("DispatchConfig off, but JSON specifies on.")
	}
	//Input handlers Unmarshal test
	if len(dc.inputHandlers) != 2 {
		t.Error("DispatchConfig number of input codes read is off.")
		return
	}
	valid := true
	if dc.inputHandlers[0] == "ABC" {
		if dc.inputHandlers[1] != "XYZ" {
			valid = false
		}
	} else if dc.inputHandlers[0] == "XYZ" {
		if dc.inputHandlers[1] != "ABC" {
			valid = false
		}
	} else {
		valid = false
	}
	if valid == false {
		t.Error("DispatchConfig input codes read incorrectly from JSON.")
	}

	//Output handlers Unmarshal test
	if len(dc.outputHandlers) != 2 {
		t.Error("DispatchConfig number of output codes read is off.")
		return
	}
	valid = true
	if dc.outputHandlers[0].code == "WOW" {
		if dc.outputHandlers[1].code != "LOL" {
			valid = false
		}
	} else if dc.outputHandlers[0].code == "LOL" {
		if dc.outputHandlers[1].code != "WOW" {
			valid = false
		}
	} else {
		valid = false
	}
	if valid == false {
		t.Error("DispatchConfig output codes read incorrectly from JSON.")
	}
}

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
