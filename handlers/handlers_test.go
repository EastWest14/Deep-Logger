package handlers

import (
	bih "deeplogger/handlers/basicinputhandler"
	boh "deeplogger/handlers/basicouthandler"
	"testing"
)

var _ InputHandler = bih.NewWithDispatcherAndInputString(nil, "")

var _ OutputHandler = boh.New(nil, "")

func TestBIH(t *testing.T) {

}
