package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/tire-fire/adventofcode/2024/lib"
)

// State represents the program state with registers, pointer, and output
type State struct {
	Registers [3]int
	Pointer   int
	Output    []int
}

func (s *State) Reset(a int) {
	s.Registers = [3]int{a, 0, 0}
	s.Output = []int{}
	s.Pointer = 0
}

// Instruction defines the interface for all instructions
type Instruction interface {
	Run(state *State)
}

// GenericInstruction handles common instruction functionality
type GenericInstruction struct {
	Opcode  int
	Operand int
}

func main() {
	lines, err := lib.ReadInput()
	if err != nil {
		panic("Failed to read input")
	}
	state, instructions := Load(lines)

	output := []int{}
	for _, instr := range instructions {
		gi := instr.(*GenericInstruction)
		output = append(output, gi.Opcode, gi.Operand)
	}

	lowerBound := 0
	for i := len(output) - 1; i >= 0; i-- {
		for a := lowerBound; a < lowerBound+int(math.Pow(8, float64(len(output)-i))); a++ {
			state.Reset(a)
			for state.Pointer < len(instructions) {
				instructions[state.Pointer].Run(state)
			}
			if compareSlices(state.Output, output[i:]) {
				if len(state.Output) == len(output) {
					fmt.Println(a)
					return
				}
				lowerBound = a * 8
				break
			}
		}
	}
	return
}

func (gi *GenericInstruction) ComboOperand(state *State) int {
	mapping := map[int]int{4: 0, 5: 1, 6: 2}
	if regIndex, ok := mapping[gi.Operand]; ok {
		return state.Registers[regIndex]
	}
	return gi.Operand
}

func (gi *GenericInstruction) Run(state *State) {
	switch gi.Opcode {
	case 0: // AdvInstruction
		state.Registers[0] /= int(math.Pow(2, float64(gi.ComboOperand(state))))
	case 1: // BxlInstruction
		state.Registers[1] ^= gi.Operand
	case 2: // BstInstruction
		state.Registers[1] = gi.ComboOperand(state) % 8
	case 3: // JnzInstruction
		if state.Registers[0] != 0 {
			state.Pointer = gi.Operand
			return
		}
	case 4: // BxcInstruction
		state.Registers[1] ^= state.Registers[2]
	case 5: // OutInstruction
		state.Output = append(state.Output, gi.ComboOperand(state)%8)
	case 6: // BdvInstruction
		state.Registers[1] = state.Registers[0] / int(math.Pow(2, float64(gi.ComboOperand(state))))
	case 7: // CdvInstruction
		state.Registers[2] = state.Registers[0] / int(math.Pow(2, float64(gi.ComboOperand(state))))
	}
	state.Pointer++
}

func Load(lines []string) (*State, []Instruction) {
	state := &State{Registers: [3]int{0, 0, 0}}
	instructions := []Instruction{}

	for _, line := range lines {
		if strings.HasPrefix(line, "Register") {
			parts := strings.Fields(line)
			num, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				panic(fmt.Sprintf("Failed to parse register value in line: %s", line))
			}
			if strings.Contains(line, "A") {
				state.Registers[0] = num
			} else if strings.Contains(line, "B") {
				state.Registers[1] = num
			} else if strings.Contains(line, "C") {
				state.Registers[2] = num
			}
		} else if strings.HasPrefix(line, "Program") {
			program := strings.Fields(line)[1]
			progParts := strings.Split(program, ",")
			for i := 0; i < len(progParts); i += 2 {
				opcode, err1 := strconv.Atoi(progParts[i])
				operand, err2 := strconv.Atoi(progParts[i+1])
				if err1 != nil || err2 != nil {
					panic(fmt.Sprintf("Invalid program input at position %d: %s", i, progParts[i]))
				}
				instructions = append(instructions, &GenericInstruction{Opcode: opcode, Operand: operand})
			}
		}
	}
	return state, instructions
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
