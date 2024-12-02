package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sort"
	"math"
)

// calculateTotalDistance computes the total distance between two lists
func calculateTotalDistance(leftList, rightList []int) int {
	log.Printf("Calculating total distance between two lists")

	// Sort both lists
	sort.Ints(leftList)
	sort.Ints(rightList)

	log.Printf("Left list (sorted): %v", leftList)
	log.Printf("Right list (sorted): %v", rightList)

	// Calculate the total distance
	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		distance := int(math.Abs(float64(leftList[i] - rightList[i])))
		log.Printf("Pair: (%d, %d), Distance: %d", leftList[i], rightList[i], distance)
		totalDistance += distance
	}

	log.Printf("Total distance calculated: %d", totalDistance)
	return totalDistance
}

// solve1 processes the input and computes the total distance between LHS and RHS.
func solve1(lines []string) int {
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
	totalDistance := calculateTotalDistance(lhs, rhs)

	return totalDistance
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
	result := solve1(lines)
	fmt.Printf("Total distance: %d\n", result)

	log.Println("Program completed successfully")
}

