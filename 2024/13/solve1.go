package main

import (
	"fmt"
	"math"
	"github.com/tire-fire/adventofcode/2024/lib"
)


type Machine struct {
	aX, aY, bX, bY, pX, pY int
}

var machines []Machine

func main() {
	lines, err := lib.ReadInput()
	if err != nil {
		panic("Failed to read input")
	}

	machines = parseMachines(lines)
	fmt.Println(machines)

	maxPrizes, totalCost := calculateMaxPrizesAndCost()
	fmt.Println(totalCost)
}

func parseMachines(lines []string) []Machine {
	machines := []Machine{}
	for i := 0; i < len(lines); i += 4 { // Process 4 lines at a time (3 data lines + 1 blank)
		if lines[i] == "" {
			continue
		}
		var aX, aY, bX, bY, pX, pY int
		fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &aX, &aY)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &bX, &bY)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &pX, &pY)
		machines = append(machines, Machine{aX, aY, bX, bY, pX, pY})
	}
	return machines
}

func calculateMaxPrizesAndCost() (int, int) {
	maxPrizes := 0
	totalCost := 0

	for _, machine := range machines {
		cost := findMinCost(machine)
		if cost != -1 {
			maxPrizes++
			totalCost += cost
		}
	}

	return maxPrizes, totalCost
}

func findMinCost(machine Machine) int {
	minCost := math.MaxInt
	foundSolution := false

	for nA := 0; nA <= 100; nA++ {
		remX := machine.pX - nA*machine.aX
		remY := machine.pY - nA*machine.aY

		// Handle cases where bX or bY is 0
		if machine.bX == 0 && remX != 0 {
			continue
		}
		if machine.bY == 0 && remY != 0 {
			continue
		}

		// Check if remaining movement can be exactly covered by button B
		if (machine.bX == 0 || remX%machine.bX == 0) && (machine.bY == 0 || remY%machine.bY == 0) {
			nB := 0
			if machine.bX != 0 {
				nB = remX / machine.bX
			}
			if machine.bY != 0 {
				if machine.bX == 0 {
					nB = remY / machine.bY
				} else if nB != remY/machine.bY {
					continue // Mismatch in nB for X and Y
				}
			}

			if nB >= 0 {
				cost := 3*nA + nB
				if cost < minCost {
					minCost = cost
				}
				foundSolution = true
			}
		}
	}

	if foundSolution {
		return minCost
	}
	return -1 // No solution
}

