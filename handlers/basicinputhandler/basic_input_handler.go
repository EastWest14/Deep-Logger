package basicinputhandler

import (
	dispatcher "deeplogger/dispatcher"
)

type BasicInputHandler struct {
	DispatchLog      *dispatcher.DispatcherLog
	InputHandlerName string
}

func NewWithDispatcherAndInputString(dl *dispatcher.DispatcherLog, inputName string) *BasicInputHandler {
	return &BasicInputHandler{DispatchLog: dl, InputHandlerName: inputName}
}

func (bih *BasicInputHandler) LogEvent(ev dispatcher.Event) {
	if bih.DispatchLog == nil {
		panic("No dispatcher registered.")
		return
	}
	ev.SetInputHandlerName(bih.InputHandlerName)
	bih.DispatchLog.InputEvent(ev)
}