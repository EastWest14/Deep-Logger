package basicouthandler

import (
	"deeplogger/simpleevent"
	"os"
	"testing"
)

func TestTakeInEvent(t *testing.T) {
	boh := BasicOutputHandler{DispatchLog: nil, OutputHandlerName: "ABC", OutputWriter: os.Stdout}

	ev := simpleevent.New("Hello world!")
	ev.SetInputHandlerName("XYZ")
	boh.takeInEvent(ev)
}
