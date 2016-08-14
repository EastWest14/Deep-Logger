package basicouthandler

import (
	olddispatcher "deeplogger/dispatcher"
	"deeplogger/event"
	dispatcher "deeplogger/newdispatcher"
	"fmt"
	"io"
)

type BasicOutputHandler struct {
	DispatchLog       *olddispatcher.DispatcherLog
	Dispatcher        *dispatcher.Dispatcher
	OutputHandlerName string
	OutputWriter      io.Writer
}

//TODO: rename.
func NewWithDispatcherAndOutputString(dl *olddispatcher.DispatcherLog, outputName string, outWriter io.Writer) *BasicOutputHandler {
	boh := BasicOutputHandler{DispatchLog: dl, OutputHandlerName: outputName, OutputWriter: outWriter}
	boh.registerWithDispatcher()
	return &boh
}

func New(disp *dispatcher.Dispatcher, name string) *BasicOutputHandler {
	boh := BasicOutputHandler{Dispatcher: disp, OutputHandlerName: name}
	//boh.registerWithDispatcher()
	//TODO: uncomment!
	return &boh
}

func (boh *BasicOutputHandler) SetOutputWriter(writer io.Writer) {
	boh.OutputWriter = writer
}

func (boh *BasicOutputHandler) TakeInEvent(ev event.Event) {
	evString := fmt.Sprintln("[" + ev.InputHandlerName() + "]: " + ev.EventMessage())
	boh.outputData([]byte(evString))
}

func (boh *BasicOutputHandler) outputData(data []byte) {
	boh.OutputWriter.Write(data)
}

func (boh *BasicOutputHandler) registerWithDispatcher() {
	boh.DispatchLog.RegisterOutputHandler(boh.OutputHandlerName, boh.TakeInEvent)
}
