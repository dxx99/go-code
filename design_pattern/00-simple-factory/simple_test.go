package simple_factory

import "testing"

func TestNewAnimal(t *testing.T) {
	api := NewAnimal(1)
	if api != nil {
		NewAnimal(1).Say()
	}
	api2 := NewAnimal(2)
	if api2 != nil {
		api2.Say()
	}
	api3 := NewAnimal(1)
	if api3 != nil {
		api3.Say()
	}
}
