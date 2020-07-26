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
	g.RunProgram([]uint64{
		gmachine.OpHALT,
	})
	if g.P != 1 {
		t.Errorf("want P == 1, got %d", g.P)
	}
}

func TestNOOP(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.RunProgram([]uint64{
		gmachine.OpNOOP,
		gmachine.OpHALT,
	})
	if g.P != 2 {
		t.Errorf("want P == 2, got %d", g.P)
	}
}

func TestINCA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.RunProgram([]uint64{
		gmachine.OpINCA,
		gmachine.OpHALT,
	})
	if g.A != 1 {
		t.Errorf("want A == 1, got %d", g.A)
	}
}

func TestDECA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.A = 2
	g.RunProgram([]uint64{
		gmachine.OpDECA,
		gmachine.OpHALT,
	})
	if g.A != 1 {
		t.Errorf("want A == 1, got %d", g.A)
	}
}

func TestSETA(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.RunProgram([]uint64{
		gmachine.OpSETA, 5,
		gmachine.OpHALT,
	})
	if g.A != 5 {
		t.Errorf("want A == 5, got %d", g.A)
	}
	if g.P != 3 {
		t.Errorf("want P == 3, got %d", g.P)
	}
}

func TestSub3(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		input, want uint64
	}{
		{input: 3, want: 1},
		{input: 100, want: 98},
		{input: 5, want: 3},
	}
	for _, tc := range tcs {
		g := gmachine.New()
		g.RunProgram([]uint64{
			gmachine.OpSETA, tc.input,
			gmachine.OpDECA,
			gmachine.OpDECA,
			gmachine.OpHALT,
		})
		if g.A != tc.want {
			t.Errorf("want A == %d, got %d", tc.want, g.A)
		}
	}
}
