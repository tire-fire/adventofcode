package main

import (
	"fmt"
	"github.com/tire-fire/adventofcode/2024/lib"
)

type Point struct {
	X, Y int
}

func main() {
	lines, err := lib.ReadInput()
	if err != nil {
		panic("Failed to read input")
	}

	uniqueAntinodes := calculateAntinodes(lines)
	fmt.Println(uniqueAntinodes)
}

func calculateAntinodes(input []string) int {
	antennas := parseMap(input)

	antinodes := make(map[Point]struct{})

	// Check pairs of antennas with the same frequency
	for _, points := range antennas {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				a, b := points[i], points[j]

				// Check for valid antinode positions
				dx, dy := b.X-a.X, b.Y-a.Y

				// Midpoints where antinodes could exist
				mid1 := Point{a.X - dx, a.Y - dy}
				mid2 := Point{b.X + dx, b.Y + dy}

				// Add to the antinode set if within bounds
				if inBounds(mid1, input) {
					antinodes[mid1] = struct{}{}
				}
				if inBounds(mid2, input) {
					antinodes[mid2] = struct{}{}
				}
			}
		}
	}

	return len(antinodes)
}

func parseMap(input []string) map[rune][]Point {
	antennas := make(map[rune][]Point)
	for y, line := range input {
		for x, char := range line {
			if char != '.' {
				antennas[char] = append(antennas[char], Point{X: x, Y: y})
			}
		}
	}
	return antennas
}

func inBounds(p Point, input []string) bool {
	return p.X >= 0 && p.Y >= 0 && p.Y < len(input) && p.X < len(input[0])
}

