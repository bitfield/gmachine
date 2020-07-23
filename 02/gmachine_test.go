package gmachine_test

import (
	"gmachine"
	"testing"
)

func TestHalt(t *testing.T) {
	g := gmachine.New()
	g.Run()
	if g.P != 1 {
		t.Errorf("want P == 1, got %d", g.P)
	}
	g.Run()
	if g.P != 2 {
		t.Errorf("want P == 2, got %d", g.P)
	}
}
