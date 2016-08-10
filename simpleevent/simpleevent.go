package simpleevent

import ()

type SimpleEvent struct {
	inputHandlerCode string
	message          string
}

func (se *SimpleEvent) InputHandlerCode() string {
	return se.inputHandlerCode
}

func (se *SimpleEvent) EventMessage() string {
	return se.message
}

func (se *SimpleEvent) SetInputHandlerCode(inputHandlerCode string) {
	se.inputHandlerCode = inputHandlerCode
	return
}

func New(message string) *SimpleEvent {
	return &SimpleEvent{message: message}
}
