package basic_input_handler

import (
	dispatcher "deeplogger/dispatcher"
	"deeplogger/simpleevent"
)

type BasicInputHandler struct {
	DispatchLog      *dispatcher.DispatcherLog
	InputHandlerCode string
}

func NewWithDispatcherAndInputString(dl *dispatcher.DispatcherLog, inputCode string) *BasicInputHandler {
	return &BasicInputHandler{DispatchLog: dl, InputHandlerCode: inputCode}
}

/*
type basicEvent struct {
	inputHandlerCode string
	eventMessage     string
	eventTime        time.Time
}

func (be *basicEvent) InputHandlerCode() string {
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
}*/

func (bih *BasicInputHandler) LogEvent(message string) {
	if bih.DispatchLog == nil {
		panic("No dispatcher registered.")
		return
	}

	//ev := &basicEvent{inputHandlerCode: bih.InputHandlerCode, eventMessage: message, eventTime: time.Now()}
	ev := simpleevent.New(message)
	ev.SetInputHandlerCode(bih.InputHandlerCode)
	bih.DispatchLog.InputEvent(ev)
}
