package basic_output_handler

import (
	"dispatcher"
	"io"
)

type BasicOutputHandler struct {
	DispatchLog       *dispatcher.DispatcherLog
	OutputHandlerCode dispatcher.OutputHandlerCode
	outputWriter      io.Writer
}

func (boh *BasicOutputHandler) takeInEvent(ev dispatcher.Event) {

}

func (boh *BasicOutputHandler) RegisterWithDispatcher() {
	boh.DispatchLog.RegisterOutputHandler(boh.OutputHandlerCode, boh.takeInEvent)
}
