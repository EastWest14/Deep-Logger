package basic_output_handler

import (
	"dispatcher"
	"fmt"
	"io"
)

type BasicOutputHandler struct {
	DispatchLog       *dispatcher.DispatcherLog
	OutputHandlerCode dispatcher.OutputHandlerCode
	OutputWriter      io.Writer
}

func NewWithDispatcherAndOutputString(dl *dispatcher.DispatcherLog, outputCode string, outWriter io.Writer) *BasicOutputHandler {
	boh := BasicOutputHandler{DispatchLog: dl, OutputHandlerCode: dispatcher.OutputHandlerCode(outputCode), OutputWriter: outWriter}
	boh.RegisterWithDispatcher()
	return &boh
}

func (boh *BasicOutputHandler) takeInEvent(ev dispatcher.Event) {
	evString := fmt.Sprintln(string(ev.InputHandlerCode()) + " - " + ev.EventMessage())
	boh.outputData([]byte(evString))
}

func (boh *BasicOutputHandler) outputData(data []byte) {
	boh.OutputWriter.Write(data)
}

func (boh *BasicOutputHandler) RegisterWithDispatcher() {
	boh.DispatchLog.RegisterOutputHandler(boh.OutputHandlerCode, boh.takeInEvent)
}
