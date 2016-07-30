package basic_input_handler

import (
	"dispatcher"
)

type BasicInputHandler struct {
	DispatchLog      *dispatcher.DispatcherLog
	InputHandlerCode dispatcher.InputHandlerCode
}

func NewWithDispatcherAndInputString(dl *dispatcher.DispatcherLog, inputCode string) *BasicInputHandler {
	return &BasicInputHandler{DispatchLog: dl, InputHandlerCode: dispatcher.InputHandlerCode(inputCode)}
}

func (bih *BasicInputHandler) RegisterWithDispatcher() {

}
