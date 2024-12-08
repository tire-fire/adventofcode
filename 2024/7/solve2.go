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
		if evaluate(target, numbers, 0, 0) {
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

func evaluate(numbers []int, target, currentIndex, currentResult int) bool {
	if currentIndex == len(numbers) {
		return currentResult == target
	}

	if currentResult > target {
		return false
	}

	nextNumber := numbers[currentIndex]

	// Try addition
	if evaluate(numbers, target, currentIndex+1, currentResult+nextNumber) {
		return true
	}

	// Try multiplication
	if evaluate(numbers, target, currentIndex+1, currentResult*nextNumber) {
		return true
	}

	// Try concatenation
	if currentIndex > 0 {
		concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentResult, nextNumber))
		if evaluate(numbers, target, currentIndex+1, concatenated) {
			return true
		}
	}

	return false
}

