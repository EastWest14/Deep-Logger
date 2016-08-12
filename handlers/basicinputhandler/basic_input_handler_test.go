package basicinputhandler

import (
	"testing"
)

func TestNew(t *testing.T) {
	bih := NewWithDispatcherAndInputString(nil, "ABC")
	if bih.InputHandlerName != "ABC" {
		t.Error("New nandler created incorrectly.")
	}
}

//TODO: implement.
/*
func TestLogEvent(t *testing.T) {

}*/
