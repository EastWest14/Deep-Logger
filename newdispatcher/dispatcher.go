package newdispatcher

type Dispatcher struct {
	name           string
	on             bool
	inputHandlers  map[string]bool
	outputHandlers map[string]interface{}
}

func NewWithName(name string) *Dispatcher {
	return &Dispatcher{name: name, inputHandlers: map[string]bool{}}
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
