package deeplogger

import (
	"deeplogger/dispatcher"
	"deeplogger/event"
	"deeplogger/handlers"
)

type Event interface {
	event.Event
}

func NewEvent(message string) Event {
	return event.New(message)
}

type InputHandler interface {
	handlers.InputHandler
}

func NewInputHandler(name string) InputHandler {
	return handlers.NewInputHandler(name)
}

//NewBlankInputHandler creates a dummy input handler, which prevents accidental nil pointer dereferencing.
func NewBlankInputHandler() InputHandler {
	return handlers.NewBlankInputHandler()
}

type OutputHandler interface {
	handlers.OutputHandler
}

func NewOutputHandler(disp *dispatcher.Dispatcher, name string) OutputHandler {
	return handlers.NewOutputHandler(disp, name)
}
