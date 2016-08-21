package handlers

import (
	bih "deeplogger/handlers/basicinputhandler"
	"testing"
)

var _ InputHandler = bih.NewWithDispatcherAndInputString(nil, "")

func TestBIH(t *testing.T) {

}
