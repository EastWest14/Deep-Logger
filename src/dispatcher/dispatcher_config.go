package dispatcher

type dispatcherConfig struct {
	name           string
	isOn           bool
	inputHandlers  []InputHandler  //TODO: provide access methods
	OutputHandlers []OutputHandler //TODO: provide access methods
	dispatchRules  []DispatchRule  //TODO: provide access methods
	//TODO: ShouldPanicOnInvalidInput
}

func (dc *dispatcherConfig) SetIsOn(on bool) {
	dc.isOn = on
}

func (dc *dispatcherConfig) IsOn() bool {
	return dc.isOn
}

type InputHandler int //TODO: temporaily int

type OutputHandler int //TODO: temporarily int

type DispatchRule int //TODO: temporarily int
