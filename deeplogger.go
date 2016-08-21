package deeplogger

import (
	"deeplogger/handlers"
	dispatcher "deeplogger/newdispatcher"
	"encoding/json"
	"log"
)

func ConstructLoggerFromConfig(config string) (inputHandlers map[string]handlers.InputHandler, disp *dispatcher.Dispatcher, outputHandlers map[string]handlers.OutputHandler) {
	disp = dispatcher.NewWithName("")
	inputHandlers = map[string]handlers.InputHandler{}
	outputHandlers = map[string]handlers.OutputHandler{}
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(config), &dat)
	if err != nil {
		log.Fatal(err.Error())
	}

	disp.SetName(dat["dispatcher_name"].(string))

	isOn := dat["isOn"].(bool)
	if isOn {
		disp.TurnOn()
	} else {
		disp.TurnOff()
	}

	var inNames []interface{}
	inNames = dat["inputHandlers"].([]interface{})
	for _, inName := range inNames {
		//modifying dispatcher
		stringName := inName.(string)
		disp.AddInputHandler(stringName, true) //TODO: is on?

		//creating handlers
		handl := handlers.CreateInputHandler(stringName)
		handl.SetDispatcher(disp)
		if _, present := inputHandlers[stringName]; present {
			panic("Attempt to ad duplicate input handlers.")
		}
		inputHandlers[stringName] = handl
	}

	outNames := dat["outputHandlers"].([]interface{})
	for _, outName := range outNames {
		stringName := outName.(string)
		//Creating handlers
		handl := handlers.NewOutputHandler(disp, stringName)
		if _, present := outputHandlers[stringName]; present {
			panic("Attempt to ad duplicate output handlers.")
		}
		outputHandlers[stringName] = handl
	}

	//TODO: find a way to automatically test
	dispatchRulesData := dat["dispatchRules"].([]interface{})
	for _, dispRule := range dispatchRulesData {
		dRule := dispRule.(map[string]interface{})
		input := dRule["input"].(string)
		output := dRule["output"].(string)
		disp.AddDispatchRule(dispatcher.NewRule(dispatcher.NewMatchCondWithName(input), output))
	}

	return inputHandlers, disp, outputHandlers
}

/*
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
}*/

type CountWriter struct {
	V int
}

func (iw *CountWriter) Write(input []byte) (n int, err error) {
	iw.V++
	return 0, nil
}
