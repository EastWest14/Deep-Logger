package dispatcher

import (
	"testing"
)

func TestConfigFromFile(t *testing.T) {
	dc := configFromFile("../config/little_config.json")
	if dc == nil {
		t.Error("Failed loading config from file.")
	}
}

func TestConfigFromString(t *testing.T) {
	caseZero := `
	{"name": "LALALA",
	"isOn": true,
	"inputHandlers": ["ABC", "XYZ"],
	"outputHandlers": ["WOW", "LOL"],
	"dispatchRules": [
		{"input":"ABC", "output": "QQQ", "intersect": false},
		{"input":"AAA", "output": "LLL", "intersect": true}
	]}
	`
	dc := configFromString(caseZero)
	if dc.name != "LALALA" {
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

	//Dispatch rules Unmarshal test
	if len(dc.dispatchRules) != 2 {
		t.Error("DispatchConfig number of rules is off.")
	}
	ruleZero := dc.dispatchRules[0]
	if string(ruleZero.Input) != "ABC" || string(ruleZero.Output.code) != "QQQ" || ruleZero.Intersect != false {
		t.Error("Rule zero read incorrectly.")
	}
	ruleOne := dc.dispatchRules[1]
	if string(ruleOne.Input) != "AAA" || string(ruleOne.Output.code) != "LLL" || ruleOne.Intersect != true {
		t.Error("Rule one read incorrectly.")
	}
}

func TestConfigSetOnOff(t *testing.T) {
	dc := dispatcherConfig{}
	dc.isOn = false
	dc.TurnOn()
	if dc.isOn != true {
		t.Error("DispatchConfig not turning on.")
	}
	dc.isOn = true
	dc.TurnOff()
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

//TODO: write tests for checking internal consistency.
func TestNewDispatchRule(t *testing.T) {
	dr := newDispatchRule("ALL", "AAA", true)
	if string(dr.Input) != "ALL" || string(dr.Output.code) != "AAA" || dr.Intersect != true {
		t.Error("newDispatchRule sets rule incorrectly.")
	}
}
