package simpleevent

import ()

type SimpleEvent struct {
	inputHandlerName string
	message          string
}

func (se *SimpleEvent) InputHandlerName() string {
	return se.inputHandlerName
}

func (se *SimpleEvent) EventMessage() string {
	return se.message
}

func (se *SimpleEvent) SetInputHandlerName(inputHandlerName string) {
	se.inputHandlerName = inputHandlerName
	return
}

func New(message string) *SimpleEvent {
	return &SimpleEvent{message: message}
}
