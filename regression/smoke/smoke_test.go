package smoke_test

import (
	"deeplogger"
	"deeplogger/dispatcher"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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
	tDir, err := ioutil.TempDir("", "test")
	if err != nil {
		panic(fmt.Sprintf("Failed creating temp directory: %s", err.Error()))
	}
	tFile := filepath.Join(tDir, "test_file_to_load")
	err = ioutil.WriteFile(tFile, []byte(config), 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to write temporary file: %s", err.Error()))
	}
	defer os.RemoveAll(tDir)

	inpHandlers, d, outHandlers, err := deeplogger.ConstructLoggerFromFilepath(tFile)
	if err != nil {
		panic("Failed loading config. " + err.Error())
	}
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
	defer func() {
		writeC.hitCounter = 0
	}()
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
}

func TestLoggingMessages(t *testing.T) {
	defer func() {
		writeC.hitCounter = 0
	}()
	cases := []struct {
		inpHandlerCode string
		message        string
	}{
		{"Input", ""},
		{"Input", "Hello World!"},
		{"Input", "Hello Again!"},
	}

	for _, aCase := range cases {
		inpHandler.LogMessage(aCase.message)
	}

	if writeC.hitCounter != len(cases) {
		t.Errorf("Expected %d write hits, got %d hits.", len(cases), writeC.hitCounter)
	}
}

func TestIsOn(t *testing.T) {
	defer func() {
		writeC.hitCounter = 0
		disp.TurnOn()
	}()
	inpHandler.LogMessage("")
	if writeC.hitCounter != 1 {
		t.Errorf("Expected write hit, got %d hits. Failure in state IsOn.", writeC.hitCounter)
	}
	writeC.hitCounter = 0

	disp.TurnOff()
	inpHandler.LogMessage("")
	if writeC.hitCounter != 0 {
		t.Errorf("Expected 0 write hits, got %d hits. Failure in state IsOff.", writeC.hitCounter)
	}
}
