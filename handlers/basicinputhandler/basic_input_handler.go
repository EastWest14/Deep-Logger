package basicinputhandler

import (
	olddispatcher "deeplogger/dispatcher"
	"deeplogger/event"
	dispatcher "deeplogger/newdispatcher"
)

type BasicInputHandler struct {
	DispatchLog      *olddispatcher.DispatcherLog
	Dispatcher       *dispatcher.Dispatcher
	InputHandlerName string
}

func NewWithDispatcherAndInputString(dl *olddispatcher.DispatcherLog, inputName string) *BasicInputHandler {
	return &BasicInputHandler{DispatchLog: dl, InputHandlerName: inputName}
}

func (bih *BasicInputHandler) SetDispatcher(d *dispatcher.Dispatcher) {
	bih.Dispatcher = d
}

func (bih *BasicInputHandler) LogEvent(ev event.Event) {
	if bih.DispatchLog == nil {
		panic("No dispatcher registered.")
		return
	}
	ev.SetInputHandlerName(bih.InputHandlerName)
	bih.DispatchLog.InputEvent(ev)
	bih.Dispatcher.InputEvent(ev)
}
