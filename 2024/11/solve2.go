package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
	"github.com/tire-fire/adventofcode/2024/lib"
)

func main() {
	lines, err := lib.ReadInput()
	if err != nil {
		panic("Failed to read input")
	}

	stones := parseInput(lines[0])
	totalStones := blink(stones, 75)

	fmt.Println(totalStones)
}

func parseInput(input string) []int {
	parts := strings.Fields(input)
	stones := make([]int, len(parts))
	for i, part := range parts {
		stones[i], _ = strconv.Atoi(part)
	}
	return stones
}

func blink(stones []int, numBlinks int) int {
	stoneCounts := make(map[int]int)
	for _, stone := range stones {
		stoneCounts[stone]++
	}

	for i := 0; i < numBlinks; i++ {
		newStoneCounts := make(map[int]int)
		for stone, count := range stoneCounts {
			switch {
			case stone == 0:
				// 0 becomes 1
				newStoneCounts[1] += count

			case isEvenDigitCount(stone):
				// split evens
				left, right := splitStone(stone)
				newStoneCounts[left] += count
				newStoneCounts[right] += count

			default:
				// * 2024
				newStoneCounts[stone*2024] += count
			}
		}
		stoneCounts = newStoneCounts
	}

	totalStones := 0
	for _, count := range stoneCounts {
		totalStones += count
	}
	return totalStones
}

func isEvenDigitCount(num int) bool {
	if num == 0 {
		return false
	}
	digitCount := int(math.Log10(float64(num))) + 1
	return digitCount%2 == 0
}

func splitStone(num int) (int, int) {
	if num == 0 {
		return 0, 0
	}

	digitCount := int(math.Log10(float64(num))) + 1
	half := digitCount / 2

	divisor := int(math.Pow10(half))
	left := num / divisor
	right := num % divisor

	return left, right
}
