package dispatcher

import (
	"encoding/json"
	"log"
)

type dispatcherConfig struct {
	name           string
	isOn           bool
	inputHandlers  []InputHandlerCode     //TODO: provide access methods
	outputHandlers []OutputHandlerElement //TODO: provide access methods
	dispatchRules  []DispatchRule         //TODO: provide access methods
	//TODO: ShouldPanicOnInvalidInput
}

//TODO: locks on reads.
//TODO: locks on writes.

//TODO: replace argument to filepath
func LoadConfigFromFile(jsonStr string) *dispatcherConfig {
	dc := dispatcherConfig{}

	var dat map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &dat)
	if err != nil {
		log.Fatal(err.Error())
	}
	dc.name = dat["name"].(string)
	dc.isOn = dat["isOn"].(bool)

	var inCodes []interface{}
	inCodes = dat["inputHandlers"].([]interface{})
	for _, inCode := range inCodes {
		//TODO: check validity
		stringCode := inCode.(string)
		dc.inputHandlers = append(dc.inputHandlers, InputHandlerCode(stringCode))
	}

	outCodes := dat["outputHandlers"].([]interface{})
	for _, outCode := range outCodes {
		//TODO: check validity
		stringCode := outCode.(string)
		dc.outputHandlers = append(dc.outputHandlers, OutputHandlerElement{OutputHandlerCode(stringCode), nil})
	}

	var dispatchRules []interface{}
	for _, dispRule := range dispatchRules {
		dRule := dispRule.(int)
		dc.dispatchRules = append(dc.dispatchRules, DispatchRule(dRule)) //TODO: modify type
	}
	return &dc
}

func (dc *dispatcherConfig) SetName(newName string) {
	//TODO: check name validity
	dc.beginChangingConfigState()
	dc.name = newName
	dc.endChangingConfigState()
}

func (dc *dispatcherConfig) Name() string {
	return dc.name
}

func (dc *dispatcherConfig) SetIsOn(on bool) {
	dc.beginChangingConfigState()
	dc.isOn = on
	dc.endChangingConfigState()
}

func (dc *dispatcherConfig) IsOn() bool {
	return dc.isOn
}

func (dc *dispatcherConfig) beginChangingConfigState() {

}

func (dc *dispatcherConfig) endChangingConfigState() {
	if dc.checkConfigStateConsistency() != true {
		panic("Dispatcher Config inconsistent.")
	}
}

func (dc *dispatcherConfig) checkConfigStateConsistency() bool {
	return true //TODO: write real method
}

type OutputHandlerCode string //TODO: temporarily int

type OutputHandlerElement struct {
	code        OutputHandlerCode
	eventOutput func(Event)
}

type DispatchRule int //TODO: temporarily int
