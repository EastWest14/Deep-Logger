package basic_output_handler

import (
	"deeplogger/simpleevent"
	"os"
	"testing"
)

func TestTakeInEvent(t *testing.T) {
	boh := BasicOutputHandler{DispatchLog: nil, OutputHandlerCode: "ABC", OutputWriter: os.Stdout}

	ev := simpleevent.New("Hello world!")
	ev.SetInputHandlerCode("XYZ")
	boh.takeInEvent(ev)
}
