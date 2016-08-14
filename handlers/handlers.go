package handlers

import (
	"deeplogger/event"
	bih "deeplogger/handlers/basicinputhandler"
	boh "deeplogger/handlers/basicouthandler"
	dispatcher "deeplogger/newdispatcher"
	"io"
)

type InputHandler interface {
	SetDispatcher(*dispatcher.Dispatcher)
	LogEvent(event.Event)
}

//TODO: add type enumeration.
func CreateInputHandler(name string) InputHandler {
	return bih.NewWithDispatcherAndInputString(nil, name)
} //TODO: rename function.

type OutputHandler interface {
	TakeInEvent(event.Event)
	SetOutputWriter(io.Writer)
}

func NewOutputHandler(disp *dispatcher.Dispatcher, name string) OutputHandler {
	return boh.New(disp, name)
}
