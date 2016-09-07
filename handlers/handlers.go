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

type BlankInputHandler struct{}

//Panic
func (blih *BlankInputHandler) SetDispatcher(*dispatcher.Dispatcher) {
	panic("Attempting to set dispatcher on a blank Input Handler.")
}

//Do nothing
func (blih *BlankInputHandler) LogEvent(event.Event) {
	return
}

func NewBlankInputHandler() InputHandler {
	return &BlankInputHandler{}
}

type OutputHandler interface {
	TakeInEvent(event.Event)
	SetOutputWriter(io.Writer)
}

func NewOutputHandler(disp *dispatcher.Dispatcher, name string) OutputHandler {
	return boh.New(disp, name)
}
