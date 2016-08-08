package basic_input_handler

import (
	"dispatcher"
	"time"
)

type BasicInputHandler struct {
	DispatchLog      *dispatcher.DispatcherLog
	InputHandlerCode dispatcher.InputHandlerCode
}

func NewWithDispatcherAndInputString(dl *dispatcher.DispatcherLog, inputCode string) *BasicInputHandler {
	return &BasicInputHandler{DispatchLog: dl, InputHandlerCode: dispatcher.InputHandlerCode(inputCode)}
}

type basicEvent struct {
	inputHandlerCode dispatcher.InputHandlerCode
	eventMessage     string
	eventTime        time.Time
}

func (be *basicEvent) InputHandlerCode() dispatcher.InputHandlerCode {
	return be.inputHandlerCode
}

func (be *basicEvent) EventMessage() string {
	return be.eventMessage
}

func (be *basicEvent) EventTime() time.Time {
	return be.eventTime
}

func (be *basicEvent) EventType() int {
	return 0
}

func (bih *BasicInputHandler) LogEvent(message string) {
	if bih.DispatchLog == nil {
		panic("No dispatcher registered.")
		return
	}

	ev := &basicEvent{inputHandlerCode: bih.InputHandlerCode, eventMessage: message, eventTime: time.Now()}
	bih.DispatchLog.InputEvent(ev)
}