// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

const HALT = 0

type Machine struct {
	CPU struct {
		PC uint16
	}
	Memory [65536]byte
}

func New() *Machine {
	return &Machine{}
}

func (m *Machine) Step() {
	// Over to you to implement `Step`!
}
