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
	_, disp, _ := ConstructLoggerFromConfig(config)
	if disp.Name() != "DISP1" {
		t.Error("Dispatcher name read incorrectly.")
	}
}
