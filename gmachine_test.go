package gmachine_test

import (
	"bytes"
	"gmachine"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	wantMemSize := gmachine.DefaultMemSize
	gotMemSize := len(g.Memory)
	if wantMemSize != gotMemSize {
		t.Errorf("want %d words of memory, got %d", wantMemSize, gotMemSize)
	}
	var wantP gmachine.Word = 0
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	var wantMemValue gmachine.Word = 0
	gotMemValue := g.Memory[gmachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
	var wantA gmachine.Word = 0
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
}

func TestHALT(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Run()
	var wantP gmachine.Word = 1
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestNOOP(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.Memory[0] = gmachine.NOOP
	g.Run()
	var wantP gmachine.Word = 2
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestRunProgram(t *testing.T) {
	t.Parallel()
	g := gmachine.New()
	g.RunProgram([]gmachine.Word{
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
	var wantA gmachine.Word = 1
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
	var wantA gmachine.Word = 1
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
	var wantA gmachine.Word = 1
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
}

func TestSubstractTwo(t *testing.T) {
	testCases := []struct {
		desc                 string
		valueA, wantA, wantP gmachine.Word
	}{
		{
			desc:   "Substract 2 from 3",
			valueA: 3,
			wantA:  1,
			wantP:  5,
		},
		{
			desc:   "Substract 2 from 200",
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
	var wantA gmachine.Word = 5
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantP gmachine.Word = 3
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestAssemble(t *testing.T) {
	t.Parallel()

	input := []string{"HALT", "NOOP"}
	want := []gmachine.Word{gmachine.HALT, gmachine.NOOP}
	got, err := gmachine.Assemble(input)
	if err != nil {
		t.Error(err)
	}
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestAssembleInvalid(t *testing.T) {
	t.Parallel()
	input := []string{""}
	_, err := gmachine.Assemble(input)
	if err == nil {
		t.Errorf("An error is expected but not found")
	}
}

func TestAssembleOperand(t *testing.T) {
	t.Parallel()
	input := []string{"SETA", "5"}
	want := []gmachine.Word{gmachine.SETA, 5}
	got, err := gmachine.Assemble(input)
	if err != nil {
		t.Error(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAssembleOperandInvalid(t *testing.T) {
	t.Parallel()
	input := []string{"SETA", "DECA"}
	_, err := gmachine.Assemble(input)
	if err == nil {
		t.Error("Expecting error but found")
	}
}

func TestAssembleFromFile(t *testing.T) {
	t.Parallel()

	want := []gmachine.Word{gmachine.HALT}
	got, err := gmachine.AssembleFromFile("testdata/local.gasm")
	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAssembleFromFileSetA(t *testing.T) {
	t.Parallel()

	words, err := gmachine.AssembleFromFile("testdata/seta.gasm")
	if err != nil {
		t.Error(err)
	}
	g := gmachine.New()
	g.RunProgram(words)
	var wantA gmachine.Word = 5
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantP gmachine.Word = 3
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestAssembleFromFileSetADeca(t *testing.T) {
	t.Parallel()

	words, err := gmachine.AssembleFromFile("testdata/setadeca.gasm")
	if err != nil {
		t.Error(err)
	}
	g := gmachine.New()
	g.RunProgram(words)
	var wantA gmachine.Word = 3
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}
	var wantP gmachine.Word = 5
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestRunProgramFromFile(t *testing.T) {
	t.Parallel()
	// SETA 258
	// DECA
	program := bytes.NewReader([]byte{
		0, 0, 0, 0, 0, 0, 0, gmachine.SETA,
		0, 0, 0, 0, 0, 0, 1, 2,
		0, 0, 0, 0, 0, 0, 0, gmachine.DECA,
	})
	g := gmachine.New()
	err := g.RunProgramFromReader(program)
	if err != nil {
		t.Fatal(err)
	}
	var wantA gmachine.Word = 257
	if wantA != g.A {
		t.Errorf("want initial A value %d, got %d", wantA, g.A)
	}

	var wantP gmachine.Word = 4
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
}

func TestReadWords(t *testing.T) {
	t.Parallel()
	want := []gmachine.Word{gmachine.SETA, math.MaxUint64, gmachine.DECA}
	input := bytes.NewReader([]byte{
		0, 0, 0, 0, 0, 0, 0, gmachine.SETA,
		255, 255, 255, 255, 255, 255, 255, 255,
		0, 0, 0, 0, 0, 0, 0, gmachine.DECA,
	})
	got, err := gmachine.ReadWords(input)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestWriteWords(t *testing.T) {
	t.Parallel()
	input := []gmachine.Word{gmachine.SETA, 10, gmachine.DECA}
	want := []byte{
		0, 0, 0, 0, 0, 0, 0, gmachine.SETA,
		0, 0, 0, 0, 0, 0, 0, 10,
		0, 0, 0, 0, 0, 0, 0, gmachine.DECA,
	}
	output := &bytes.Buffer{}
	err := gmachine.WriteWords(output, input)
	if err != nil {
		t.Fatal(err)
	}
	got := output.Bytes()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}
