// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024
const (
	HALT = iota
	NOOP
	INCA
	DECA
	SETA
)

type Instruction struct {
	Opcode   Word
	Operands int
}

var TranslateTable = map[string]Instruction{
	"HALT": {Opcode: HALT, Operands: 0},
	"NOOP": {Opcode: NOOP, Operands: 0},
	"SETA": {Opcode: SETA, Operands: 1},
	"DECA": {Opcode: DECA, Operands: 0},
	"INCA": {Opcode: INCA, Operands: 0},
}

type Word uint64

type GMachine struct {
	A      Word
	P      Word
	Memory []Word
}

func New() *GMachine {
	return &GMachine{
		A:      0,
		P:      0,
		Memory: make([]Word, DefaultMemSize),
	}
}

func (g *GMachine) Run() {
	for {
		opcode := g.Memory[g.P]
		g.P++
		switch opcode {
		case NOOP:
		case HALT:
			return
		case INCA:
			g.A++
		case DECA:
			g.A--
		case SETA:
			g.A = g.Memory[g.P]
			g.P++
		}
	}

}

func (g *GMachine) RunProgram(instructions []Word) {
	for i := range instructions {
		g.Memory[i] = instructions[i]
	}
	g.Run()
}

func (g *GMachine) ExecuteBinary(binPath string) error {
	binFile, err := os.Open(binPath)
	if err != nil {
		return err
	}
	defer binFile.Close()
	return g.RunProgramFromReader(binFile)
}

func Assemble(code []string) ([]Word, error) {
	words := []Word{}
	for pos := 0; pos < len(code); pos++ {
		op, ok := TranslateTable[code[pos]]
		if !ok {
			return nil, fmt.Errorf("invalid instruction %q", code[pos])
		}
		words = append(words, op.Opcode)
		if op.Operands > 0 {
			if pos+1 >= len(code) {
				return nil, errors.New("missing operand")
			}
			temp, err := strconv.Atoi(code[pos+1])
			if err != nil {
				return nil, err
			}
			operand := Word(temp)
			words = append(words, operand)
			pos++
		}
	}
	return words, nil
}

func AssembleFromFile(path string) ([]Word, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	code := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		code = append(code, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	words, err := Assemble(code)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func AssembleFromFileToBinary(inPath, outPath string) error {
	data, err := AssembleFromFile(inPath)
	if err != nil {
		return err
	}
	outFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer outFile.Close()
	return WriteWords(outFile, data)
}

func (g *GMachine) RunProgramFromReader(r io.Reader) error {
	words, err := ReadWords(r)
	if err != nil {
		return err
	}
	g.RunProgram(words)
	return nil
}

func ReadWords(r io.Reader) ([]Word, error) {
	rawBytes := make([]byte, 8)
	words := []Word{}
	for {
		_, err := r.Read(rawBytes)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		b := binary.BigEndian.Uint64(rawBytes)
		words = append(words, Word(b))
	}
	return words, nil
}

func WriteWords(w io.Writer, data []Word) error {
	for _, word := range data {
		rawBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(rawBytes, uint64(word))
		w.Write(rawBytes)
	}
	return nil
}

func RunCLI(path string) error {
	g := New()
	return g.ExecuteBinary(path)
}
