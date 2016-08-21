package dispatcher

import (
	"deeplogger/event"
	"fmt"
)

type Dispatcher struct {
	name           string
	on             bool
	inputHandlers  map[string]bool
	outputHandlers map[string]func(ev event.Event)
	rules          []*DispatchRule
}

func New(name string) *Dispatcher {
	return &Dispatcher{name: name, inputHandlers: map[string]bool{}, outputHandlers: map[string]func(ev event.Event){}}
}

func (d *Dispatcher) Name() string {
	return d.name
}

func (d *Dispatcher) SetName(name string) {
	d.name = name
}

func (d *Dispatcher) IsOn() bool {
	return d.on
}

func (d *Dispatcher) TurnOn() {
	d.on = true
}

func (d *Dispatcher) TurnOff() {
	d.on = false
}

func (d *Dispatcher) AddInputHandler(name string, on bool) {
	if _, ok := d.inputHandlers[name]; ok {
		panic("Attempt to add a duplicate input handler.")
	} else {
		d.inputHandlers[name] = on
	}
}

func (d *Dispatcher) HasInputHandler(name string) (exists, isOn bool) {
	isOn, exists = d.inputHandlers[name]
	return
}

func (d *Dispatcher) AddOutputHandler(name string, handler func(ev event.Event)) {
	if _, ok := d.outputHandlers[name]; ok {
		panic("Attempt to add a duplicate output handler.")
	} else {
		d.outputHandlers[name] = handler
	}
}

func (d *Dispatcher) HasOutputHandler(name string) bool {
	_, exists := d.outputHandlers[name]
	return exists
}

func (d *Dispatcher) AddRule(rule *DispatchRule) {
	d.rules = append(d.rules, rule)
}

func (d *Dispatcher) InputEvent(ev event.Event) {
	//Check if input name is valid
	if exists, _ := d.HasInputHandler(ev.InputHandlerName()); !exists {
		panic("Message from unregistered input handler.")
	}

	for _, rule := range d.rules {
		if rule.matchesEvent(ev) {
			handlerFunc, ok := d.outputHandlers[rule.OutputHandlerName]
			if !ok {
				panic("Output handler not found.")
			} else if handlerFunc == nil {
				panic("No handler func for routed event.")
			}
			handlerFunc(ev)
			break
		}
	}
}

type MatchCondition struct {
	InputHandlerName string
}

func NewMatchCondition(inputHandlerName string) MatchCondition {
	return MatchCondition{InputHandlerName: inputHandlerName}
}

type DispatchRule struct {
	MatchCond         MatchCondition
	OutputHandlerName string
}

func (dr *DispatchRule) string() string {
	return fmt.Sprintln("Rule has input=" + dr.MatchCond.InputHandlerName + ", output=" + dr.OutputHandlerName)
}

func NewRule(matchCond MatchCondition, outHandlerName string) *DispatchRule {
	return &DispatchRule{MatchCond: matchCond, OutputHandlerName: outHandlerName}
}

func (dr *DispatchRule) matchesEvent(ev event.Event) bool {
	return ev.InputHandlerName() == dr.MatchCond.InputHandlerName
}
