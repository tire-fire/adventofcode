package main

import (
	"fmt"
	"github.com/tire-fire/adventofcode/2024/lib"
	"math"
	"strconv"
	"strings"
)

// Define the initial registers
var A, B, C int
var program []int
var output []int

// Operand Value Resolver for Combo Operands
func resolveCombo(operand int) int {
	switch operand {
	case 0, 1, 2, 3: // Literals 0-3
		return operand
	case 4: // Register A
		return A
	case 5: // Register B
		return B
	case 6: // Register C
		return C
	default:
		panic("Invalid combo operand")
	}
}

func executeProgram() {
	ip := 0 // Instruction Pointer
	for ip < len(program) {
		opcode := program[ip]   // Current opcode
		operand := program[ip+1] // Operand
		switch opcode {
		case 0: // adv: A = A / (2^operand)
			denominator := int(math.Pow(2, float64(resolveCombo(operand))))
			A = A / denominator
		case 1: // bxl: B = B XOR operand
			B = B ^ operand
		case 2: // bst: B = combo operand % 8
			B = resolveCombo(operand) % 8
		case 3: // jnz: if A != 0, jump to operand
			if A != 0 {
				ip = operand
				continue // Skip increment of ip
			}
		case 4: // bxc: B = B XOR C
			B = B ^ C
		case 5: // out: output combo operand % 8
			value := resolveCombo(operand) % 8
			output = append(output, value)
		case 6: // bdv: B = A / (2^operand)
			denominator := int(math.Pow(2, float64(resolveCombo(operand))))
			B = A / denominator
		case 7: // cdv: C = A / (2^operand)
			denominator := int(math.Pow(2, float64(resolveCombo(operand))))
			C = A / denominator
		default:
			panic(fmt.Sprintf("Unknown opcode: %d", opcode))
		}
		ip += 2 // Move to the next instruction
	}
}

func main() {
	lines, err := lib.ReadInput()
	if err != nil {
		panic("Failed to read input")
	}

	// Parse the input for registers and program
	for _, line := range lines {
		if strings.HasPrefix(line, "Register A:") {
			A, _ = strconv.Atoi(strings.TrimSpace(strings.Split(line, ":")[1]))
		} else if strings.HasPrefix(line, "Register B:") {
			B, _ = strconv.Atoi(strings.TrimSpace(strings.Split(line, ":")[1]))
		} else if strings.HasPrefix(line, "Register C:") {
			C, _ = strconv.Atoi(strings.TrimSpace(strings.Split(line, ":")[1]))
		} else if strings.HasPrefix(line, "Program:") {
			programStr := strings.TrimSpace(strings.Split(line, ":")[1])
			for _, val := range strings.Split(programStr, ",") {
				intVal, _ := strconv.Atoi(strings.TrimSpace(val))
				program = append(program, intVal)
			}
		}
	}

	// Run the program
	executeProgram()

	// Print the output joined by commas
	for i, val := range output {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(val)
	}
	fmt.Println()
}

