package gmachine

const (
	OpHALT = iota
	OpNOP
)

type Machine struct {
	A, P   int64
	Memory []int64
}

func New() Machine {
	return Machine{
		Memory: make([]int64, DefaultMemSize),
	}
}

func (g *Machine) Run() {
	for {
		op := g.Memory[g.P]
		g.P++
		if op == OpHALT {
			return
		}
	}
}
