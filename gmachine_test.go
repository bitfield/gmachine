package gmachine_test

import (
	"gmachine"
	"testing"
)

func TestNew(t *testing.T) {
	g := gmachine.New()
	wantMemSize := gmachine.DefaultMemSize
	gotMemSize := len(g.Memory)
	if wantMemSize != gotMemSize {
		t.Errorf("want %d words of memory, got %d", gotMemSize, wantMemSize)
	}
	var wantP int64 = 0
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	var wantA int64 = 0
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantMemValue int64 = 0
	gotMemValue := g.Memory[0]
	if wantMemValue != gotMemValue {
		t.Errorf("want memory location 0 to contain %d, got %d", wantMemValue, gotMemValue)
	}
}
