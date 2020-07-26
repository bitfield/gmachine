// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

// Opcodes understood by the G-machine.
const (
	OpHALT = iota
	OpNOOP
	OpINCA
	OpDECA
	OpSETA
)

// Machine represents an instance of the G-machine, with memory and register
// state.
type Machine struct {
	A, P   uint64
	Memory []uint64
}

// New returns a pointer to a new Machine, initialised to its default state.
func New() *Machine {
	return &Machine{
		Memory: make([]uint64, DefaultMemSize),
	}
}

// Run starts the machine's fetch-execute cycle, fetching instructions from
// memory and executing them until told to stop (or encountering an error).
func (g *Machine) Run() {
	for {
		op := g.Memory[g.P]
		g.P++
		switch op {
		case OpDECA:
			g.A--
		case OpHALT:
			return
		case OpINCA:
			g.A++
		case OpSETA:
			g.A = g.Memory[g.P]
			g.P++
		}
	}
}

func (g *Machine) RunProgram(program []uint64) {
	copy(g.Memory, program)
	g.Run()
}
