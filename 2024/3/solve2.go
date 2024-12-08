package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"github.com/tire-fire/adventofcode/2024/lib"
)

func parseAndCalculate(lines []string) int {
	mulRe := regexp.MustCompile(`mul\(\s*(\d+)\s*,\s*(\d+)\s*\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	totalSum := 0
	enabled := true // Multiplications are enabled at the start

	for _, line := range lines {
		fmt.Println("Processing line: %s", line)

		// Split the line into possible segments
		segments := strings.Split(line, ")")
		for _, segment := range segments {
			segment = segment + ")" // Re-add the closing parenthesis

			// Handle `do()` and `don't()` instructions
			if doRe.MatchString(segment) {
				fmt.Println("do() found: Enabling multiplications")
				enabled = true
				continue
			}
			if dontRe.MatchString(segment) {
				fmt.Println("don't() found: Disabling multiplications")
				enabled = false
				continue
			}

			// Process each mul instruction if enabled
			if enabled {
				// Find valid mul(X,Y) instructions
				matches := mulRe.FindAllStringSubmatch(segment, -1)
				fmt.Println("Found mul matches: %v", matches)

				for _, match := range matches {
					// Extract X and Y from the match and convert to integers
					x, err1 := strconv.Atoi(match[1])
					y, err2 := strconv.Atoi(match[2])

					if err1 == nil && err2 == nil {
						product := x * y
						fmt.Println("mul(%d,%d) = %d", x, y, product)
						totalSum += product
					}
				}
			}
		}
	}

	return totalSum
}

func main() {
	lines, err := lib.ReadInput()
	if err != nil {
		panic("Failed to read input")
	}

	result := parseAndCalculate(lines)
	fmt.Println(result)
}

