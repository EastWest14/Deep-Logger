//Deeplogger is a package for logging, debugging and automated testing of concurrent systems.
package deeplogger

import (
	dispatcher "deeplogger/dispatcher"
	"encoding/json"
	"log"
)

//ConstructLoggerFromConfig returns input handlers, dispatcher and output handlers that can be used to construct the deep logger system.
func ConstructLoggerFromConfig(config string) (inputHandlers map[string]InputHandler, disp *dispatcher.Dispatcher, outputHandlers map[string]OutputHandler) {
	disp = dispatcher.New("")
	inputHandlers = map[string]InputHandler{}
	outputHandlers = map[string]OutputHandler{}
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
		handl := NewInputHandler(stringName)
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
		handl := NewOutputHandler(disp, stringName)
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
		disp.AddRule(dispatcher.NewRule(dispatcher.NewMatchCondition(input), output))
	}

	return inputHandlers, disp, outputHandlers
}

//CountWriter is a mock object that implements io.Writer. Used to count number of calls to Write.
type CountWriter struct {
	V int
}

//Write increments internal counter by one.
func (iw *CountWriter) Write(input []byte) (n int, err error) {
	iw.V++
	return 0, nil
}
