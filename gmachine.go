// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

const (
	OpHALT = iota
	OpNOOP
	OpINCA
	OpDECA
	OpSETA
)

type Machine struct{
	Memory []uint64
	P uint64
	A uint64
}

func New() Machine {
	return Machine{
		Memory: make([]uint64, DefaultMemSize),
	}
}

func (g *Machine) Run() {
	for {
		op := g.Memory[g.P]
		g.P++
		switch op {
		case OpHALT:
			return
		case OpNOOP:
		case OpINCA:
			g.A++
		case OpDECA:
			g.A--
		case OpSETA:
			g.A = g.Memory[g.P]
			g.P++
		}
	}
}