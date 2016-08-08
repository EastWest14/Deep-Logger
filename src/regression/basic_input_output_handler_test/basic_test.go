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

func TestA(t *testing.T) {
	var buffer bytes.Buffer
	inHandler := configureDispatcherAndHandlers("../../../config/test_config.json", &buffer, INPUT_HANDLER_CODE, OUTPUT_HANDLER_CODE)

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

func configureDispatcherAndHandlers(configFile string, writer io.Writer, inp_code, out_code string) *inhandl.BasicInputHandler {
	disp := dispatcher.NewDispatcherWithFile(configFile)
	inHandler := inhandl.NewWithDispatcherAndInputString(disp, INPUT_HANDLER_CODE)
	_ = outhandl.NewWithDispatcherAndOutputString(disp, OUTPUT_HANDLER_CODE, writer)
	return inHandler
}

type incrementWriter struct {
	v int
}

func (iw *incrementWriter) Write(input []byte) (n int, err error) {
	iw.v++
	return 0, nil
}

func TestIsOff(t *testing.T) {
	iw := &incrementWriter{v: 0}
	inHandler := configureDispatcherAndHandlers("../../../config/test_config_off.json", iw, INPUT_HANDLER_CODE, OUTPUT_HANDLER_CODE)
	inHandler.LogEvent("Shouldn't go through")

	if iw.v != 0 {
		t.Error("Got output although the dispatcher was off.")
	}

}
