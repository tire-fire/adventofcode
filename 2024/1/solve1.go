package main

import (
	"fmt"
	"sort"
	"math"
	"strings"
	"strconv"

	"github.com/tire-fire/adventofcode/2024/lib"
)

func calculateTotalDistance(leftList, rightList []int) int {
	sort.Ints(leftList)
	sort.Ints(rightList)

	fmt.Println("Left list (sorted): %v", leftList)
	fmt.Println("Right list (sorted): %v", rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		distance := int(math.Abs(float64(leftList[i] - rightList[i])))
		fmt.Println("Pair: (%d, %d), Distance: %d", leftList[i], rightList[i], distance)
		totalDistance += distance
	}

	return totalDistance
}

func getDistance(lines []string) int {
	lhs, rhs, err := splitInput(lines)
	if err != nil {
		panic("Error processing input")
	}

	if len(lhs) != len(rhs) {
		panic("Left-hand side and right-hand side lists are of unequal length")
	}

	totalDistance := calculateTotalDistance(lhs, rhs)

	return totalDistance
}

func splitInput(lines []string) ([]int, []int, error) {
	var lhs, rhs []int

	for i, line := range lines {
		// Split the line into fields (columns)
		columns := strings.Fields(line)
		if len(columns) != 2 {
			return nil, nil, fmt.Errorf("invalid format on line %d: %s (expected 2 columns)", i+1, line)
		}

		left, err := strconv.Atoi(columns[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in left-hand side on line %d: %s", i+1, columns[0])
		}
		right, err := strconv.Atoi(columns[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in right-hand side on line %d: %s", i+1, columns[1])
		}

		lhs = append(lhs, left)
		rhs = append(rhs, right)
	}

	return lhs, rhs, nil
}

func main() {
	lines, err := lib.ReadInput()
	if err != nil {
		panic("Failed to read input")
	}

	result := getDistance(lines)
	fmt.Println(result)
}

