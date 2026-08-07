// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

const HALT = 0

type Machine struct {
	CPU struct {
		PC byte
	}
	Memory [256]byte
}

func New() *Machine {
	return &Machine{}
}

func (m *Machine) Run() {
	// Over to you to implement `Run`!
}
