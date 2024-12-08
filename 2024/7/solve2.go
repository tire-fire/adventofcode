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
	memo := make(map[string]bool)
	return evaluateWithMemo(numbers, target, 0, 0, memo)
}

func evaluateWithMemo(numbers []int, target, currentIndex, currentResult int, memo map[string]bool) bool {
	if currentIndex == len(numbers) {
		return currentResult == target
	}

	memoKey := fmt.Sprintf("%d:%d", currentIndex, currentResult)
	if val, exists := memo[memoKey]; exists {
		return val
	}

	nextNumber := numbers[currentIndex]

	// Try addition
	if evaluateWithMemo(numbers, target, currentIndex+1, currentResult+nextNumber, memo) {
		memo[memoKey] = true
		return true
	}

	// Try multiplication
	if evaluateWithMemo(numbers, target, currentIndex+1, currentResult*nextNumber, memo) {
		memo[memoKey] = true
		return true
	}

	// Try concatenation
	if currentIndex > 0 {
		concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentResult, nextNumber))
		if evaluateWithMemo(numbers, target, currentIndex+1, concatenated, memo) {
			memo[memoKey] = true
			return true
		}
	}

	memo[memoKey] = false
	return false
}

