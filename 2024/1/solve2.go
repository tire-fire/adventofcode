package main

import (
	"fmt"
	"log"
)


// calculateTotalDistance computes the total distance between two lists
func calculateSimilarity(leftList, rightList []int) int {
	log.Printf("Calculating similarity between two lists")

	rhsCounts := make(map[int]int)
	for _, num := range rightList {
		rhsCounts[num]++
	}

	// Calculate the similarity score
	similarityScore := 0
	for _, num := range leftList {
		count := rhsCounts[num]
		score := num * count
		log.Printf("Number: %d, Count in RHS: %d, Contribution to Similarity Score: %d", num, count, score)
		similarityScore += score
	}

	log.Printf("Final Similarity Score: %d", similarityScore)
	return similarityScore
}

// solve2 processes the input and computes the total distance between LHS and RHS.
func solve2(lines []string) int {
	log.Println("Splitting input into left-hand side (LHS) and right-hand side (RHS) lists")
	lhs, rhs, err := splitInput(lines)
	if err != nil {
		log.Fatalf("Error processing input: %v", err)
	}

	// Validate that both lists are of the same length
	if len(lhs) != len(rhs) {
		log.Fatalf("Left-hand side and right-hand side lists are of unequal length")
	}

	// Calculate total distance
	log.Println("Calculating total distance between LHS and RHS")
	similarity := calculateSimilarity(lhs, rhs)

	return similarity
}

func main() {
	log.Println("Starting program")

	// Read the input file
	lines, err := ReadLines("input")
	//lines, err := ReadLines("exampleinput")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	// Solve the problem using solve1
	result := solve2(lines)
	fmt.Printf("Total Similarity: %d\n", result)

	log.Println("Program completed successfully")
}

