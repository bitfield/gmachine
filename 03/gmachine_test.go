package gmachine_test

import (
	"gmachine"
	"testing"
)

func TestNop(t *testing.T) {
	g := gmachine.New()
	g.Memory[0] = gmachine.OpNOP
	g.Run()
	if g.P != 2 {
		t.Errorf("want P == 2, got %d", g.P)
	}
}
