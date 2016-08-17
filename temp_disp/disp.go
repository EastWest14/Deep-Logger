package disp

type MatchCondition struct {
	InputHandlerName string
}

func NewMatchCondWithName(inputHandlerName string) MatchCondition {
	return MatchCondition{InputHandlerName: inputHandlerName}
}

type DispatchRule struct {
	MatchCond         MatchCondition
	OutputHandlerName string
}

func NewRule(matchCond MatchCondition, outHandlerName string) *DispatchRule {
	return &DispatchRule{MatchCond: matchCond, OutputHandlerName: outHandlerName}
}
