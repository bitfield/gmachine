package gmachine_test

import (
	"testing"

	"gmachine"
)

func TestNew(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	var wantP gmachine.Word = 0
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	var wantMemValue gmachine.Word = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
}
