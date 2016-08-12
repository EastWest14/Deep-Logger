package basicouthandler

import (
	"deeplogger/dispatcher"
	"deeplogger/event"
	"fmt"
	"io"
)

type BasicOutputHandler struct {
	DispatchLog       *dispatcher.DispatcherLog
	OutputHandlerName string
	OutputWriter      io.Writer
}

func NewWithDispatcherAndOutputString(dl *dispatcher.DispatcherLog, outputName string, outWriter io.Writer) *BasicOutputHandler {
	boh := BasicOutputHandler{DispatchLog: dl, OutputHandlerName: outputName, OutputWriter: outWriter}
	boh.registerWithDispatcher()
	return &boh
}

func (boh *BasicOutputHandler) takeInEvent(ev event.Event) {
	evString := fmt.Sprintln("[" + ev.InputHandlerName() + "]: " + ev.EventMessage())
	boh.outputData([]byte(evString))
}

func (boh *BasicOutputHandler) outputData(data []byte) {
	boh.OutputWriter.Write(data)
}

func (boh *BasicOutputHandler) registerWithDispatcher() {
	boh.DispatchLog.RegisterOutputHandler(boh.OutputHandlerName, boh.takeInEvent)
}
