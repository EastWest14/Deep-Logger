package dispatcher

import (
	"encoding/json"
	"log"
	"io/ioutil"
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

func ConfigFromFile(filename string) *dispatcherConfig {
	rawContent, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Failed to load config.") //TODO: Don't panic here, return error
	}

	return LoadConfig(string(rawContent))
}

func LoadConfig(jsonStr string) *dispatcherConfig {
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

	dc.dispatchRules = loadDispatchRules(dat)

	return &dc
}

func loadDispatchRules(dat map[string]interface{}) []DispatchRule {
	var dispatchRules []DispatchRule

	var dispatchRulesData []interface{}
	dispatchRulesData = dat["dispatchRules"].([]interface{})
	for _, dispRule := range dispatchRulesData {
		dRule := dispRule.(map[string]interface{})
		input := dRule["input"].(string)
		output := dRule["output"].(string)
		intersect := dRule["intersect"].(bool)

		dispatchRules = append(dispatchRules, *NewDispatchRule(InputHandlerCode(input), OutputHandlerCode(output), intersect))
	}
	return dispatchRules
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

type DispatchRule struct {
	Input     InputHandlerCode
	Output    OutputHandlerElement
	Intersect bool
}

//Input is input handler code or "ALL"
func NewDispatchRule(input InputHandlerCode, output OutputHandlerCode, intersect bool) *DispatchRule {
	return &DispatchRule{Input: input, Output: OutputHandlerElement{output, nil}, Intersect: intersect}
}
