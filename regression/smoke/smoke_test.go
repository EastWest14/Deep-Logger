package smoke_test

import (
	"deeplogger"
	"deeplogger/dispatcher"
	"os"
	"testing"
)

//Start with blank input handlers to prevent accidental nil pointer dereference
var inpHandler deeplogger.InputHandler = deeplogger.NewBlankInputHandler()
var inpHandlerToNowhere deeplogger.InputHandler = deeplogger.NewBlankInputHandler()
var disp *dispatcher.Dispatcher
var outHandler deeplogger.OutputHandler

var writeC writeCounter = writeCounter{0}

type writeCounter struct {
	hitCounter int
}

func (wc *writeCounter) Write(input []byte) (n int, err error) {
	wc.hitCounter++
	return 0, nil
}

const config = `{"dispatcher_name": "Dispatcher",
	"isOn": true,
	"inputHandlers": ["Input"],
	"outputHandlers": ["Output"],
	"dispatchRules": [
		{"input":"Input", "output": "Output"}
	]
}`

func setupWithConfigString() {
	inpHandlers, d, outHandlers := deeplogger.ConstructLoggerFromConfig(config)
	disp = d
	inpHandler = inpHandlers["Input"]
	inpHandlerToNowhere = deeplogger.NewInputHandler("Input2")
	inpHandlerToNowhere.SetDispatcher(disp)
	outHandler = outHandlers["Output"]
	outHandler.SetOutputWriter(&writeC)
}

func setupManual() {
	disp = dispatcher.New("Dispatcher")
	disp.AddInputHandler("Input", true)
	disp.AddRule(dispatcher.NewRule(dispatcher.NewMatchCondition("Input"), "Output"))
	inpHandler = deeplogger.NewInputHandler("Input")
	inpHandler.SetDispatcher(disp)
	inpHandlerToNowhere = deeplogger.NewInputHandler("Input2")
	inpHandlerToNowhere.SetDispatcher(disp)
	outHandler = deeplogger.NewOutputHandler(disp, "Output")
	outHandler.SetOutputWriter(&writeC)
}

func TestMain(m *testing.M) {
	setupWithConfigString()
	res1 := m.Run()
	setupManual()
	res2 := m.Run()
	if res1 == 0 && res2 == 0 {
		os.Exit(0)
	} else if res1 != 0 {
		os.Exit(res1)
	} else {
		os.Exit(res2)
	}
}

func TestName(t *testing.T) {
	const expectedName = "Dispatcher"
	if disp.Name() != expectedName {
		t.Errorf("Expected dispatcher name to be %s, got %s", expectedName, disp.Name())
	}
}

func TestLoggingEvents(t *testing.T) {
	cases := []struct {
		inpHandlerCode string
		message        string
	}{
		{"Input", ""},
		{"Input", "Hello World!"},
		{"Input", "Hello Again!"},
		{"Input2", "Miss!"},
	}

	panicCount := 0
	for _, aCase := range cases {
		if aCase.inpHandlerCode == "Input" {
			inpHandler.LogEvent(deeplogger.NewEvent(aCase.message))
		} else {
			func() {
				defer func() {
					if r := recover(); r != nil {
						panicCount++

					}
				}()
				inpHandlerToNowhere.LogEvent(deeplogger.NewEvent(aCase.message))
			}()
		}
	}

	const (
		expectedPanicNum = 1
		expectedHitNum   = 3
	)
	if panicCount != expectedPanicNum {
		t.Errorf("Expected %d panic(s) from invalid input handler. Got %d panics.", expectedPanicNum, panicCount)
	}
	if writeC.hitCounter != expectedHitNum {
		t.Errorf("Expected %d write hits, got %d hits.", expectedHitNum, writeC.hitCounter)
	}
	writeC.hitCounter = 0
}
