package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"fmt"
	"strconv"
)

// ReadLines reads the content of a file and returns a slice of strings, each representing a line.
func ReadLines(filename string) ([]string, error) {
	log.Printf("Reading file: %s", filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return nil, err
	}
	defer func() {
		log.Printf("Closing file: %s", filename)
		file.Close()
	}()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading file: %v", err)
		return nil, err
	}

	log.Printf("Successfully read %d lines", len(lines))
	return lines, nil
}

// splitInput splits all lines into two integer slices: left-hand side (LHS) and right-hand side (RHS).
func splitInput(lines []string) ([]int, []int, error) {
	var lhs, rhs []int

	for i, line := range lines {
		// Split the line into fields (columns)
		columns := strings.Fields(line)
		if len(columns) != 2 {
			return nil, nil, fmt.Errorf("invalid format on line %d: %s (expected 2 columns)", i+1, line)
		}

		// Parse the left-hand side and right-hand side values
		left, err := strconv.Atoi(columns[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in left-hand side on line %d: %s", i+1, columns[0])
		}
		right, err := strconv.Atoi(columns[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in right-hand side on line %d: %s", i+1, columns[1])
		}

		// Append to the respective lists
		lhs = append(lhs, left)
		rhs = append(rhs, right)
	}

	return lhs, rhs, nil
}
