package handlers

import (
	"deeplogger/dispatcher"
	"deeplogger/event"
	bih "deeplogger/handlers/basicinputhandler"
	boh "deeplogger/handlers/basicouthandler"
	"io"
)

type InputHandler interface {
	SetDispatcher(*dispatcher.Dispatcher)
	LogEvent(event.Event)
}

//TODO: add type enumeration.

func NewInputHandler(name string) InputHandler {
	return bih.New(nil, name)
}

type OutputHandler interface {
	TakeInEvent(event.Event)
	SetOutputWriter(io.Writer)
}

func NewOutputHandler(disp *dispatcher.Dispatcher, name string) OutputHandler {
	return boh.New(disp, name)
}
