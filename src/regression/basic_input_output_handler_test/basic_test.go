package basic_input_output_handler_test

import (
	"bytes"
	"dispatcher"
	inhandl "handlers/basic_input_handler"
	outhandl "handlers/basic_output_handler"
	"io"
	"strings"
	"testing"
)

const (
	INPUT_HANDLER_CODE  = "AAA"
	OUTPUT_HANDLER_CODE = "ZZZ"
)

func TestMain(t *testing.T) {
	var buffer bytes.Buffer
	inHandler := configureDispatcherAndHandlers(&buffer)

	const message = "Hello World!"
	inHandler.LogEvent(message)
	output := buffer.String()
	if !strings.Contains(output, message) {
		t.Errorf("Log output doesn't contain the message: %s, instead the message is: %s", message, output)
	}
	if !strings.Contains(output, INPUT_HANDLER_CODE) {
		t.Errorf("Log output doesn't contain the input code: %s, instead the message is: %s", INPUT_HANDLER_CODE, output)
	}
}

func configureDispatcherAndHandlers(writer io.Writer) *inhandl.BasicInputHandler {
	disp := dispatcher.NewDispatcherWithFile("../../../config/test_config.json")
	inHandler := inhandl.NewWithDispatcherAndInputString(disp, INPUT_HANDLER_CODE)
	_ = outhandl.NewWithDispatcherAndOutputString(disp, OUTPUT_HANDLER_CODE, writer)
	return inHandler
}
