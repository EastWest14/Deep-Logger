package dispatcher

import (
	"errors"
	"time"
)

type DispatcherLog struct {
	*dispatcherConfig
}

func NewDispatcherWithFile(filename string) *DispatcherLog {
	return &DispatcherLog{dispatcherConfig: configFromFile(filename)}
}

type Event interface {
	InputHandlerCode() string
	EventMessage() string
	EventTime() time.Time
	EventType() int //TODO: will be an enumeration.
}

func (dl *DispatcherLog) InputEvent(ev Event) {

	var outputE OutputHandlerElement
	ok, outputEFromRule := dl.matchOutputHandler(ev)

	if ok && outputEFromRule != nil {
		//Need to find a pointer, not the value.
		for _, outputEl := range dl.outputHandlers {
			if outputEl.code == outputEFromRule.code {
				outputE = *outputEl
			}
		}
		outputE.eventOutput(ev)
		return
	} else if ok {
		return
	} else {
		panic("No output function set for output handler:" + string(outputE.code))
		return
	}

}

func (dl *DispatcherLog) matchOutputHandler(ev Event) (ok bool, outputH *OutputHandlerElement) {
	//TODO: begin
	//TODO: defer end
	if !checkEventValidity(ev) {
		panic("event is invalid!")
	}
	if !dl.isOn {
		ok = true
		outputH = nil
		return
	}
	return true, dl.routeEvent(ev)
}

func checkEventValidity(event Event) bool {
	return checkInputCodeValidity(event.InputHandlerCode())
}

func (dl *DispatcherLog) routeEvent(ev Event) *OutputHandlerElement {
	//TODO: begin read
	for _, rule := range dl.dispatchRules {
		matches, _ := rule.ruleMatch(ev)
		if !matches {
			continue
		} else {
			return &rule.Output
		}
		//TODO: non-intersecting rules not implemented yet.
	}
	panic("Can't route event")
	return nil
	//TODO: end read.
}

func (rule *DispatchRule) ruleMatch(ev Event) (matches, intersects bool) {
	if ev.InputHandlerCode() != rule.Input {
		return false, false
	} else {
		matches = true
		intersects = rule.Intersect
		return
	}
}

//TODO: add locking.
func (dl *DispatcherLog) RegisterOutputHandler(outputHC OutputHandlerCode, handlerFunc func(Event)) error {
	for _, outputHE := range dl.outputHandlers {
		if outputHE.code == outputHC {
			outputHE.eventOutput = handlerFunc
			return nil
		}

	}
	return errors.New("Failed to register output handler.")
}
