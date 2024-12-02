package main

import (
	"fmt"
	"sort"
	"math"
	"github.com/tire-fire/adventofcode/2024/lib"
)

func calculateTotalDistance(leftList, rightList []int) int {
	fmt.Printf("Calculating total distance between two lists")

	sort.Ints(leftList)
	sort.Ints(rightList)

	fmt.Printf("Left list (sorted): %v", leftList)
	fmt.Printf("Right list (sorted): %v", rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		distance := int(math.Abs(float64(leftList[i] - rightList[i])))
		fmt.Printf("Pair: (%d, %d), Distance: %d", leftList[i], rightList[i], distance)
		totalDistance += distance
	}

	return totalDistance
}

func getDistance(lines []string) int {
	lhs, rhs, err := splitInput(lines)
	if err != nil {
		fmt.Fatalf("Error processing input: %v", err)
	}

	if len(lhs) != len(rhs) {
		fmt.Fatalf("Left-hand side and right-hand side lists are of unequal length")
	}

	totalDistance := calculateTotalDistance(lhs, rhs)

	return totalDistance
}

func main() {
	lines, err := ReadInput()
	if err != nil {
		fmt.Fatalf("Failed to read input: %v", err)
	}

	result := getDistance(lines)
	fmt.Printf(result)
}

