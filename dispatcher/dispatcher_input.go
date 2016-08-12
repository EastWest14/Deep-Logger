package dispatcher

import (
	"strings"
)

//TODO: add error codes?

/*
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
}*/

func checkInputNameValidity(inputName string) bool {
	if len(inputName) != 3 {
		return false
	}
	if strings.ToUpper(inputName) != inputName {
		return false
	}
	return true
}

//check if this error code can be inserted into the slice (check duplicates)
func (dc *dispatcherConfig) checkInputCodeInsertability(iHanName string) bool {
	for _, name := range dc.inputHandlers {
		if name == iHanName {
			return false
		}
	}
	return true
}
