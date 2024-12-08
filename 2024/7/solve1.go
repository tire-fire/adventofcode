package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/tire-fire/adventofcode/2024/lib"
)

func main() {
	lines, err := lib.ReadInput()
	if err != nil {
		panic("Failed to read input")
	}

	totalCalibration := 0
	for _, line := range lines {
		target, numbers := parseEquation(line)
		if canAchieveTarget(target, numbers) {
			totalCalibration += target
		}
	}

	fmt.Println(totalCalibration)
}

func parseEquation(line string) (int, []int) {
	parts := strings.Split(line, ":")
	target, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	numberStrings := strings.Fields(strings.TrimSpace(parts[1]))
	numbers := make([]int, len(numberStrings))
	for i, num := range numberStrings {
		numbers[i], _ = strconv.Atoi(num)
	}
	return target, numbers
}

func canAchieveTarget(target int, numbers []int) bool {
	if len(numbers) == 0 {
		return false
	}
	operatorsCount := len(numbers) - 1
	operators := generateOperatorPermutations(operatorsCount)

	for _, ops := range operators {
		if evaluateExpression(numbers, ops) == target {
			return true
		}
	}

	return false
}

func generateOperatorPermutations(n int) [][]string {
	if n == 0 {
		return [][]string{{}}
	}

	subPermutations := generateOperatorPermutations(n - 1)
	permutations := [][]string{}

	for _, sub := range subPermutations {
		permutations = append(permutations, append([]string{"+"}, sub...))
		permutations = append(permutations, append([]string{"*"}, sub...))
	}

	return permutations
}

func evaluateExpression(numbers []int, operators []string) int {
	result := numbers[0]
	for i, op := range operators {
		if op == "+" {
			result += numbers[i+1]
		} else if op == "*" {
			result *= numbers[i+1]
		}
	}
	return result
}

