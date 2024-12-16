package main

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/tire-fire/adventofcode/2024/lib"
)

func main() {
	lines, err := lib.ReadInput()
	if err != nil {
		panic("Failed to read input")
	}
	robots := getRobots(lines)

	// Dimensions of the space
	width := 101
	height := 103

	// The pattern of a "Christmas tree"
	treePattern := []struct{ x, y int }{
		{2, 0},
		{1, 1}, {2, 1}, {3, 1},
		{0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2},
	}

	seconds := 0
	for {
		// Check if current arrangement matches the tree pattern
		if isChristmasTree(robots, width, height, treePattern) {
			fmt.Println(seconds)
			break
		}
		// Advance time by one second
		tick(robots, width, height)
		seconds++
	}
}

type Robot struct {
	x, y   int
	vx, vy int
}

func getRobots(lines []string) []Robot {
	var robots []Robot
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			continue
		}
		pPart := parts[0] // p=0,4
		vPart := parts[1] // v=3,-3

		pxStr := strings.Split(strings.TrimPrefix(pPart, "p="), ",")[0]
		pyStr := strings.Split(strings.TrimPrefix(pPart, "p="), ",")[1]

		xPos, _ := strconv.Atoi(pxStr)
		yPos, _ := strconv.Atoi(pyStr)

		vxStr := strings.Split(strings.TrimPrefix(vPart, "v="), ",")[0]
		vyStr := strings.Split(strings.TrimPrefix(vPart, "v="), ",")[1]

		vx, _ := strconv.Atoi(vxStr)
		vy, _ := strconv.Atoi(vyStr)

		robots = append(robots, Robot{x: xPos, y: yPos, vx: vx, vy: vy})
	}

	return robots
}

// tick advances all robot positions by one second and wraps them
func tick(robots []Robot, width, height int) {
	for i := range robots {
		r := &robots[i]
		r.x = (r.x + r.vx) % width
		if r.x < 0 {
			r.x += width
		}
		r.y = (r.y + r.vy) % height
		if r.y < 0 {
			r.y += height
		}
	}
}

func isChristmasTree(robots []Robot, width, height int, pattern []struct{ x, y int }) bool {
	if len(robots) < len(pattern) {
		return false
	}

	// Convert robot positions to a map for quick lookup
	posSet := make(map[[2]int]bool, len(robots))
	for _, r := range robots {
		posSet[[2]int{r.x, r.y}] = true
	}

	// To check pattern, try to align the tree based on each robot as a potential "anchor".
	// This could be optimized, but is straightforward.
	for anchorPos := range posSet {
		// anchorPos is where the pattern[0] would be placed
		dx := anchorPos[0] - pattern[0].x
		dy := anchorPos[1] - pattern[0].y

		match := true
		for _, p := range pattern {
			tx := p.x + dx
			ty := p.y + dy
			if !posSet[[2]int{tx, ty}] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}

	return false
}
