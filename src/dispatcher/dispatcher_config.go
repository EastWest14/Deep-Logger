package dispatcher

type dispatcherConfig struct {
	name           string
	isOn           bool
	inputHandlers  []InputHandler  //TODO: provide access methods
	OutputHandlers []OutputHandler //TODO: provide access methods
	dispatchRules  []DispatchRule  //TODO: provide access methods
	//TODO: ShouldPanicOnInvalidInput
}

//TODO: locks on reads.
//TODO: locks on writes.

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

type InputHandler int //TODO: temporaily int

type OutputHandler int //TODO: temporarily int

type DispatchRule int //TODO: temporarily int
