// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024
const HALT = 0
const NOOP = 1
const INCA = 2
const DECA = 3

type GMachine struct {
	A      uint64
	P      uint64
	Memory [DefaultMemSize]uint64
}

func New() *GMachine {
	return &GMachine{
		A:      0,
		P:      0,
		Memory: [1024]uint64{},
	}
}

func (g *GMachine) Run() {
	for _, slot := range g.Memory {
		g.P++
		switch slot {
		case HALT:
			return
		case INCA:
			g.A++
		case DECA:
			g.A--
		}
	}

}
