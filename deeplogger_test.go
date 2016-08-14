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
	inpHandl, disp, outHandl := ConstructLoggerFromConfig(config)
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
}
