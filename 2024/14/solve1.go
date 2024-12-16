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

	getSafetyFactor(robots)
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

func getSafetyFactor(robots []Robot) {
	// Dimensions of the space
	width := 101
	height := 103

	// Time to simulate
	t := 100

	// Advance each robot position by t seconds
	for i := range robots {
		r := &robots[i]
		r.x = (r.x + r.vx*t) % width
		if r.x < 0 {
			r.x += width
		}
		r.y = (r.y + r.vy*t) % height
		if r.y < 0 {
			r.y += height
		}
	}

	// Determine quadrants
	var q1, q2, q3, q4 int
	midX := width / 2   // 50 for width=101
	midY := height / 2  // 51 for height=103

	for _, r := range robots {
		x, y := r.x, r.y
		// Skip if on middle lines
		if x == midX || y == midY {
			continue
		}

		if x > midX && y < midY {
			q1++
		} else if x < midX && y < midY {
			q2++
		} else if x < midX && y > midY {
			q3++
		} else if x > midX && y > midY {
			q4++
		}
	}

	safetyFactor := q1 * q2 * q3 * q4
	fmt.Println(safetyFactor)
}

