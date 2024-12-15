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

	stoneStrings := strings.Fields(lines[0])

	stones := make([]int, len(stoneStrings))
	for i, s := range stoneStrings {
		num, _ := strconv.Atoi(s)
		stones[i] = num
	}

	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	fmt.Println(len(stones))
}

func blink(stones []int) []int {
	newStones := []int{}

	for _, stone := range stones {
		switch {
		case stone == 0:
			// 0 becomes 1
			newStones = append(newStones, 1)

		case isEvenDigitCount(stone):
			// split evens
			left, right := splitStone(stone)
			newStones = append(newStones, left, right)

		default:
			// * 2024
			newStones = append(newStones, stone*2024)
		}
	}

	return newStones
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
