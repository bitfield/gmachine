package gmachine_test

import (
	"gmachine"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	wantMemSize := gmachine.DefaultMemSize
	gotMemSize := len(g.Memory)
	if wantMemSize != gotMemSize {
		t.Errorf("want %d words of memory, got %d", wantMemSize, gotMemSize)
	}
	var wantP uint64 = 0
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	var wantMemValue uint64 = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
	var wantA uint64 = 0
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
}

func TestHALT(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Run()
	var wantP uint64 = 1
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestNOOP(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.NOOP
	g.Run()
	var wantP uint64 = 2
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestRunProgram(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.RunProgram([]uint64{
		gmachine.NOOP,
		gmachine.HALT,
	})
	if g.P != 2 {
		t.Errorf("want P == 2, got %d", g.P)
	}
}

func TestINCA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.INCA
	g.Run()
	var wantA uint64 = 1
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
}

func TestDECA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.A = 2
	g.Memory[0] = gmachine.DECA
	g.Run()
	var wantA uint64 = 1
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
}

func TestSubstract2From3(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.INCA
	g.Memory[1] = gmachine.INCA
	g.Memory[2] = gmachine.INCA
	g.Memory[3] = gmachine.DECA
	g.Memory[4] = gmachine.DECA
	g.Run()
	var wantA uint64 = 1
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
}

func TestSubstractTwo(t *testing.T) {
	testCases := []struct {
		desc                 string
		valueA, wantA, wantP uint64
	}{
		{
			desc:   "Substract 2 from 3",
			valueA: 3,
			wantA:  1,
			wantP:  5,
		},
		{
			desc:   "Substract 2 from 3",
			valueA: 200,
			wantA:  198,
			wantP:  5,
		},
	}
	for _, tC := range testCases {
		g := gmachine.New()
		t.Run(tC.desc, func(t *testing.T) {
			g.Memory[0] = gmachine.SETA
			g.Memory[1] = tC.valueA
			g.Memory[2] = gmachine.DECA
			g.Memory[3] = gmachine.DECA
			g.Run()
			if tC.wantA != g.A {
				t.Errorf("want A value %d, got %d", tC.wantA, g.A)
			}
			if tC.wantP != g.P {
				t.Errorf("want P value %d, got %d", tC.wantP, g.P)
			}
		})
	}
}
func TestSETA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.SETA
	g.Memory[1] = 5
	g.Run()
	var wantA uint64 = 5
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantP uint64 = 3
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}
