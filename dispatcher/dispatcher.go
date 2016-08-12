package dispatcher

import (
	"deeplogger/event"
	"errors"
)

type DispatcherLog struct {
	*dispatcherConfig
}

func NewDispatcherWithFile(filename string) *DispatcherLog {
	return &DispatcherLog{dispatcherConfig: configFromFile(filename)}
}

func (dl *DispatcherLog) InputEvent(ev event.Event) {

	var outputE OutputHandlerElement
	ok, outputEFromRule := dl.matchOutputHandler(ev)

	if ok && outputEFromRule != nil {
		//Need to find a pointer, not the value.
		for _, outputEl := range dl.outputHandlers {
			if outputEl.name == outputEFromRule.name {
				outputE = *outputEl
			}
		}
		outputE.eventOutput(ev)
		return
	} else if ok {
		return
	} else {
		panic("No output function set for output handler:" + string(outputE.name))
		return
	}

}

func (dl *DispatcherLog) matchOutputHandler(ev event.Event) (ok bool, outputH *OutputHandlerElement) {
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

func checkEventValidity(event event.Event) bool {
	return checkInputNameValidity(event.InputHandlerName())
}

func (dl *DispatcherLog) routeEvent(ev event.Event) *OutputHandlerElement {
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

func (rule *DispatchRule) ruleMatch(ev event.Event) (matches, intersects bool) {
	if ev.InputHandlerName() != rule.Input {
		return false, false
	} else {
		matches = true
		intersects = rule.Intersect
		return
	}
}

//TODO: add locking.
func (dl *DispatcherLog) RegisterOutputHandler(outputHC string, handlerFunc func(event.Event)) error {
	for _, outputHE := range dl.outputHandlers {
		if outputHE.name == outputHC {
			outputHE.eventOutput = handlerFunc
			return nil
		}

	}
	return errors.New("Failed to register output handler.")
}
