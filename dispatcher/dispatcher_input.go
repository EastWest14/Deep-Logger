package dispatcher

import (
	"errors"
	"strings"
)

type InputHandlerCode string

//TODO: add error codes?

func (dc *dispatcherConfig) AddInputHandlerCode(iHanCode InputHandlerCode) error {
	dc.beginChangingConfigState()
	if checkInputCodeValidity(iHanCode) != true {
		return errors.New("Invalid input handler code.")
	}
	if dc.checkInputCodeInsertability(iHanCode) != true {
		return errors.New("Input handler code cannot be inserted (duplicate).")
	}
	length := len(dc.inputHandlers)
	capacity := cap(dc.inputHandlers)
	if length == capacity {
		newSlice := make([]InputHandlerCode, length, capacity*2+1)
		copy(newSlice, dc.inputHandlers)
		dc.inputHandlers = newSlice
	}
	dc.inputHandlers = dc.inputHandlers[0 : length+1]
	dc.inputHandlers[length] = iHanCode
	dc.endChangingConfigState()
	return nil
}

func checkInputCodeValidity(inputCode InputHandlerCode) bool {
	ic := string(inputCode)
	if len(ic) != 3 {
		return false
	}
	if strings.ToUpper(ic) != ic {
		return false
	}
	return true
}

//check if this error code can be inserted into the slice (check duplicates)
func (dc *dispatcherConfig) checkInputCodeInsertability(iHanCode InputHandlerCode) bool {
	for _, code := range dc.inputHandlers {
		if code == iHanCode {
			return false
		}
	}
	return true
}
