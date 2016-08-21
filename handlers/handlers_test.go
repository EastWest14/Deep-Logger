package handlers

import (
	bih "deeplogger/handlers/basicinputhandler"
	"testing"
)

var _ InputHandler = bih.New(nil, "")

func TestBIH(t *testing.T) {

}
