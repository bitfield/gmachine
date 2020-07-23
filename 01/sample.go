package gmachine

type Machine struct {
	A, P   int64
	Memory []int64
}

func New() Machine {
	return Machine{
		Memory: make([]int64, DefaultMemSize),
	}
}
