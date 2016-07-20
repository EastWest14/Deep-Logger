package dispatcher

import (
	"time"
)

type DispatcherLog struct {
	dispatcherConfig
}

type Event interface {
	InputHandlerCode() InputHandlerCode
	EventMessage() string
	EventTime() time.Time
	EventType() int //TODO: will be an enumeration
}

func (dl *DispatcherLog) InputEvent(ev Event) {

}

func (dl *DispatcherLog) matchOutputHandler(ev Event) (ok bool, outputH OutputHandlerCode) {
	//TODO: begin
	//TODO: defer end
	if !checkEventValidity(ev) {
		panic("event is invalid!")
	}
	if !dl.isOn {
		ok = false //TODO: is this right?
		outputH = OutputHandlerCode("")
		return
	}
	outputE := dl.routeEvent(ev)
	return true, outputE.code
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
	//TODO: end read
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
