package main

import (
	"fmt"
	"strings"
	"strconv"

	"github.com/tire-fire/adventofcode/2024/lib"
)


func calculateSimilarity(leftList, rightList []int) int {
	rhsCounts := make(map[int]int)
	for _, num := range rightList {
		rhsCounts[num]++
	}

	similarityScore := 0
	for _, num := range leftList {
		count := rhsCounts[num]
		score := num * count
		fmt.Println("Number: %d, Count in RHS: %d, Contribution to Similarity Score: %d", num, count, score)
		similarityScore += score
	}

	return similarityScore
}

func getSimilarity(lines []string) int {
	lhs, rhs, err := splitInput(lines)
	if err != nil {
		panic("Error processing input")
	}

	if len(lhs) != len(rhs) {
		panic("Left-hand side and right-hand side lists are of unequal length")
	}

	similarity := calculateSimilarity(lhs, rhs)

	return similarity
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

	result := getSimilarity(lines)
	fmt.Println(result)
}

