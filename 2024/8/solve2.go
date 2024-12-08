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

				// Get all collinear points between and beyond the two antennas
				collinearPoints := getCollinearPoints(a, b, len(input[0]), len(input))
				for _, p := range collinearPoints {
					antinodes[p] = struct{}{}
				}
			}
		}

		// Every antenna itself is also an antinode
		for _, p := range points {
			antinodes[p] = struct{}{}
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

func getCollinearPoints(a, b Point, width, height int) []Point {
	var points []Point

	// Calculate the direction vector between the two points
	dx, dy := b.X-a.X, b.Y-a.Y
	gcd := gcd(abs(dx), abs(dy)) // Simplify the vector
	dx /= gcd
	dy /= gcd

	// Extend the line backward
	x, y := a.X, a.Y
	for {
		x -= dx
		y -= dy
		if x < 0 || y < 0 || x >= width || y >= height {
			break
		}
		points = append(points, Point{X: x, Y: y})
	}

	// Extend the line forward
	x, y = b.X, b.Y
	for {
		x += dx
		y += dy
		if x < 0 || y < 0 || x >= width || y >= height {
			break
		}
		points = append(points, Point{X: x, Y: y})
	}

	// Include all points between a and b
	x, y = a.X, a.Y
	for {
		if x == b.X && y == b.Y {
			break
		}
		points = append(points, Point{X: x, Y: y})
		x += dx
		y += dy
	}

	return points
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

