package deeplogger

import (
	"deeplogger/dispatcher"
	"deeplogger/handlers"
	"encoding/json"
	"log"
)

func ConstructLoggerFromConfig(config string) (inputHandlers map[string]handlers.InputHandler, disp *dispatcher.DispatcherLog, outputHandlers map[string]handlers.OutputHandler) {
	disp = dispatcher.New()

	var dat map[string]interface{}
	err := json.Unmarshal([]byte(config), &dat)
	if err != nil {
		log.Fatal(err.Error())
	}
	disp.SetName(dat["dispatcher_name"].(string))
	/*
		dc.isOn = dat["isOn"].(bool)

		var inNames []interface{}
		inNames = dat["inputHandlers"].([]interface{})
		for _, inName := range inNames {
			//TODO: check validity
			stringName := inName.(string)
			dc.inputHandlers = append(dc.inputHandlers, stringName)
		}

		outNames := dat["outputHandlers"].([]interface{})
		for _, outName := range outNames {
			//TODO: check validity
			stringName := outName.(string)
			dc.outputHandlers = append(dc.outputHandlers, &OutputHandlerElement{stringName, nil})
		}

		dc.dispatchRules = loadDispatchRules(dat)

		//return &dc
	*/
	return nil, disp, nil
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
