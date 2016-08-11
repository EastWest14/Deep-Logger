package basicinputhandler

import (
	dispatcher "deeplogger/dispatcher"
)

type BasicInputHandler struct {
	DispatchLog      *dispatcher.DispatcherLog
	InputHandlerCode string
}

func NewWithDispatcherAndInputString(dl *dispatcher.DispatcherLog, inputCode string) *BasicInputHandler {
	return &BasicInputHandler{DispatchLog: dl, InputHandlerCode: inputCode}
}

func (bih *BasicInputHandler) LogEvent(ev dispatcher.Event) {
	if bih.DispatchLog == nil {
		panic("No dispatcher registered.")
		return
	}
	ev.SetInputHandlerCode(bih.InputHandlerCode)
	bih.DispatchLog.InputEvent(ev)
}
