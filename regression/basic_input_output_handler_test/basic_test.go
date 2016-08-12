package basic_input_output_handler_test

import (
	"bytes"
	"deeplogger"
	"deeplogger/dispatcher"
	"deeplogger/event"
	inhandl "deeplogger/handlers/basicinputhandler"
	outhandl "deeplogger/handlers/basicouthandler"
	"fmt"
	"io"
	"strings"
	"testing"
)

const (
	INPUT_HANDLER_NAME  = "AAA"
	OUTPUT_HANDLER_NAME = "ZZZ"
)

func TestA(t *testing.T) {
	var buffer bytes.Buffer
	inHandler := configureDispatcherAndHandlers("../../config/test_config.json", &buffer, INPUT_HANDLER_NAME, OUTPUT_HANDLER_NAME)

	const message = "Hello World!"
	inHandler.LogEvent(event.New(message))
	output := buffer.String()
	if !strings.Contains(output, message) {
		t.Errorf("Log output doesn't contain the message: %s, instead the message is: %s", message, output)
	}
	if !strings.Contains(output, INPUT_HANDLER_NAME) {
		t.Errorf("Log output doesn't contain the input handler name: %s, instead the message is: %s", INPUT_HANDLER_NAME, output)
	}
	fmt.Println(output)
}

func configureDispatcherAndHandlers(configFile string, writer io.Writer, inp_name, out_name string) *inhandl.BasicInputHandler {
	disp := dispatcher.NewDispatcherWithFile(configFile)
	inHandler := inhandl.NewWithDispatcherAndInputString(disp, INPUT_HANDLER_NAME)
	_ = outhandl.NewWithDispatcherAndOutputString(disp, OUTPUT_HANDLER_NAME, writer)
	return inHandler
}

func TestIsOff(t *testing.T) {
	iw := deeplogger.CountWriter{V: 0}
	inHandler := configureDispatcherAndHandlers("../../config/test_config_off.json", &iw, INPUT_HANDLER_NAME, OUTPUT_HANDLER_NAME)
	inHandler.LogEvent(event.New("Shouldn't go through"))

	if iw.V != 0 {
		t.Error("Got output although the dispatcher was off.")
	}

}
