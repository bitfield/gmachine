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
	var wantA uint64 = 0
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantMemValue uint64 = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
}

func TestHALT(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Run()
	wantP := uint64(1)
	if g.P != wantP {
		t.Fatalf("want P == %d, got %d", wantP, g.P)
	}
}

func TestNOOP(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = uint64(gmachine.OpNOOP)
	g.Run()
	wantP := uint64(2)
	if g.P != wantP {
		t.Fatalf("want P == %d, got %d", wantP, g.P)
	}
}

func TestINCA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.OpINCA
	g.Run()
	wantA := uint64(1)
	if wantA != g.A {
		t.Errorf("want A value %d, got %d", wantA, g.A)
	}
}

func TestDECA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.A = 2
	g.Memory[0] = gmachine.OpDECA
	g.Run()
	wantA := uint64(1)
	if wantA != g.A {
		t.Errorf("want A value %d, got %d", wantA, g.A)
	}
}

func TestSubtraction(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.OpSETA
	g.Memory[1] = 3
	g.Memory[2] = gmachine.OpDECA
	g.Memory[3] = gmachine.OpDECA
	g.Run()
	wantA := uint64(1)
	if wantA != g.A {
		t.Errorf("want A value %d, got %d", wantA, g.A)
	}
}

func TestSETA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.OpSETA
	g.Memory[1] = 5
	g.Run()
	wantA := uint64(5)
	if wantA != g.A {
		t.Errorf("want A value %d, got %d", wantA, g.A)
	}
}