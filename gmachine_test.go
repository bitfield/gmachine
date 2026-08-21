package gmachine_test

import (
	"testing"

	"gmachine"
)

func TestNew(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	if g.CPU.PC != 0 {
		t.Errorf("after New, want pc == 0, got %d", g.CPU.PC)
	}
	got := g.Memory[0]
	if got != 0 {
		t.Errorf("after New, want Memory[0] == 0, got %d", got)
	}
}

// Uncomment this test once the previous test passes!
// 
// func TestHalt(t *testing.T) {
// 	t.Parallel()
// 	g := gmachine.New()
// 	g.Step()
// 	if g.CPU.PC != 1 {
// 		t.Errorf("after `halt`, want pc == 1, got %d", g.CPU.PC)
// 	}
// }
