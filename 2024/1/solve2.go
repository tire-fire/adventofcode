package main

import (
	"fmt"
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
		fmt.Printf("Number: %d, Count in RHS: %d, Contribution to Similarity Score: %d", num, count, score)
		similarityScore += score
	}

	return similarityScore
}

func getSimilarity(lines []string) int {
	lhs, rhs, err := splitInput(lines)
	if err != nil {
		fmt.Fatalf("Error processing input: %v", err)
	}

	if len(lhs) != len(rhs) {
		fmt.Fatalf("Left-hand side and right-hand side lists are of unequal length")
	}

	similarity := calculateSimilarity(lhs, rhs)

	return similarity
}

func main() {
	lines, err := ReadInput()
	if err != nil {
		fmt.Fatalf("Failed to read input: %v", err)
	}

	result := getSimilarity(lines)
	fmt.Println(result)
}

