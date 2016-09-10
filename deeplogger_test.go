package deeplogger

import (
	"testing"
)

const config = `{"dispatcher_name": "DISP1",
	"isOn": true,
	"inputHandlers": ["AAA", "BBB"],
	"outputHandlers": ["ZZZ", "YYY"],
	"dispatchRules": [
		{"input":"AAA", "output": "ZZZ"},
		{"input":"BBB", "output": "YYY"}
	]
}`

func TestConstructLoggerFromConfig(t *testing.T) {
	//Failed cases
	failConfigs := []string{"",
		"---",
		`{"a": "b"}`,
		`{"dispatcher_name": "D"}`,
		`{"dispatcher_name": "D", "isOn": false}`,
		`{"dispatcher_name": "D", "isOn": false, "inputHandlers": ["i"]}`,
		`{"dispatcher_name": "D", "isOn": false, "inputHandlers": ["i"], "outputHandlers": ["o"]}`,
		`{"dispatcher_name": "D", "isOn": false, "inputHandlers": ["i"], "outputHandlers": ["o"], "dispatchRules": [{}, {}]}`,
	}
	for i, fConf := range failConfigs {
		_, _, _, err := ConstructLoggerFromConfig(fConf)
		if err == nil {
			t.Errorf("In failed case %d expected error, but got no error.", i)
		}
	}

	//Working case
	inpHandl, disp, outHandl, err := ConstructLoggerFromConfig(config)
	if err != nil {
		t.Error("Failed to load config. Error: " + err.Error())
	}
	if disp.Name() != "DISP1" {
		t.Error("Dispatcher name read incorrectly.")
	}
	if disp.IsOn() != true {
		t.Error("Dispatcher should be on.")
	}
	if exists, _ := disp.HasInputHandler("AAA"); !exists {
		t.Error("Input handler not created in dispatcher.")
	}
	if exists, _ := disp.HasInputHandler("BBB"); !exists {
		t.Error("Input handler not created in dispatcher.")
	}
	if exists := disp.HasOutputHandler("YYY"); !exists {
		t.Error("Output handler not created in dispatcher.")
	}
	if exists := disp.HasOutputHandler("ZZZ"); !exists {
		t.Error("Output handler not created in dispatcher.")
	}

	if _, exists := inpHandl["AAA"]; !exists {
		t.Error("Input handler not created.")
	}
	if _, exists := inpHandl["BBB"]; !exists {
		t.Error("Input handler not created.")
	}
	if _, exists := inpHandl["ZZZ"]; exists {
		t.Error("Input handler created for wrong key.")
	}

	if _, exists := outHandl["ZZZ"]; !exists {
		t.Error("Output handler not created.")
	}
	if _, exists := outHandl["YYY"]; !exists {
		t.Error("Output handler not created.")
	}
	if _, exists := outHandl["nope"]; exists {
		t.Error("Output handler test false positive.")
	}
}
