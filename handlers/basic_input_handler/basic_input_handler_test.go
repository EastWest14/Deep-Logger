package basic_input_handler

import (
	"deeplogger/dispatcher"
	"testing"
)

func TestNew(t *testing.T) {
	bih := NewWithDispatcherAndInputString(nil, "ABC")
	if bih.InputHandlerCode != dispatcher.InputHandlerCode("ABC") {
		t.Error("New nandler created incorrectly.")
	}
}

//TODO: implement.
/*
func TestLogEvent(t *testing.T) {

}*/
