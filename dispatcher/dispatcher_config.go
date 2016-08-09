package dispatcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type dispatcherConfig struct {
	name           string
	isOn           bool
	inputHandlers  []string                //TODO: provide access methods
	outputHandlers []*OutputHandlerElement //TODO: provide access methods
	dispatchRules  []DispatchRule          //TODO: provide access methods
	//TODO: ShouldPanicOnInvalidInput
}

//TODO: locks on reads.
//TODO: locks on writes.

func configFromFile(filename string) *dispatcherConfig {
	rawContent, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to load config: " + filename + ".") //TODO: Don't panic here, return error
	}

	return configFromString(string(rawContent))
}

func configFromString(jsonStr string) *dispatcherConfig {
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
		dc.inputHandlers = append(dc.inputHandlers, stringCode)
	}

	outCodes := dat["outputHandlers"].([]interface{})
	for _, outCode := range outCodes {
		//TODO: check validity
		stringCode := outCode.(string)
		dc.outputHandlers = append(dc.outputHandlers, &OutputHandlerElement{stringCode, nil})
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

		dispatchRules = append(dispatchRules, *newDispatchRule(input, output, intersect))
	}
	return dispatchRules
}

func (dc *dispatcherConfig) IsOn() bool {
	return dc.isOn
}

func (dc *dispatcherConfig) TurnOn() {
	dc.beginChangingConfigState()
	dc.isOn = true
	dc.endChangingConfigState()
}

func (dc *dispatcherConfig) TurnOff() {
	dc.beginChangingConfigState()
	dc.isOn = false
	dc.endChangingConfigState()
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

type OutputHandlerElement struct {
	code        string
	eventOutput func(Event)
}

type DispatchRule struct {
	Input     string
	Output    OutputHandlerElement
	Intersect bool
}

//Input is input handler code or "ALL"
func newDispatchRule(input, output string, intersect bool) *DispatchRule {
	return &DispatchRule{Input: input, Output: OutputHandlerElement{output, nil}, Intersect: intersect}
}
