package main

import (
	"fmt"
	"math/big"
	"github.com/tire-fire/adventofcode/2024/lib"
)

type Machine struct {
	Ax, Ay *big.Int
	Bx, By *big.Int
	Px, Py *big.Int
}

var machines []Machine

func main() {
	lines, err := lib.ReadInput()
	if err != nil {
		panic("Failed to read input")
	}

	machines = parseMachines(lines)
	fmt.Println(machines)

	winnableCosts := []*big.Int{}

	for _, m := range machines {
		cost, ok := solveMachine(m)
		if ok {
			winnableCosts = append(winnableCosts, cost)
		}
	}

	// So we should find exactly those two solutions.
	totalCost := big.NewInt(0)
	for _, c := range winnableCosts {
		totalCost.Add(totalCost, c)
	}

	fmt.Println(len(winnableCosts), totalCost)
}

func parseMachines(lines []string) []Machine {
	offset := big.NewInt(10000000000000)
	for i := 0; i < len(lines); i += 4 {
		if lines[i] == "" {
			continue
		}
		var aX, aY, bX, bY, pX, pY int64
		fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &aX, &aY)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &bX, &bY)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &pX, &pY)
		PxBig := big.NewInt(pX)
		PyBig := big.NewInt(pY)
		PxBig.Add(PxBig, offset)
		PyBig.Add(PyBig, offset)
		machines = append(machines, Machine{big.NewInt(aX), big.NewInt(aY), big.NewInt(bX), big.NewInt(bY), PxBig, PyBig})
	}
	return machines
}

// solveMachine tries to solve the linear system for the given machine.
// Returns (cost, true) if solvable and (0, false) if not.
func solveMachine(m Machine) (*big.Int, bool) {
	// D = Ax*By - Ay*Bx
	D := new(big.Int).Sub(
		new(big.Int).Mul(m.Ax, m.By),
		new(big.Int).Mul(m.Ay, m.Bx),
	)

	if D.Sign() == 0 {
		// No unique solution
		return big.NewInt(0), false
	}

	// a_num = Px*By - Py*Bx
	a_num := new(big.Int).Sub(
		new(big.Int).Mul(m.Px, m.By),
		new(big.Int).Mul(m.Py, m.Bx),
	)

	// b_num = -Px*Ay + Py*Ax
	b_num := new(big.Int).Add(
		new(big.Int).Mul(m.Py, m.Ax),
		new(big.Int).Neg(new(big.Int).Mul(m.Px, m.Ay)),
	)

	// Check divisibility for a and b
	if !divisible(a_num, D) || !divisible(b_num, D) {
		return big.NewInt(0), false
	}

	a := new(big.Int).Div(a_num, D)
	b := new(big.Int).Div(b_num, D)

	// a and b must be nonnegative
	if a.Sign() < 0 || b.Sign() < 0 {
		return big.NewInt(0), false
	}

	// cost = 3*a + b
	// a,b are possibly huge, but fit in int64 only if very large might cause issues.
	// We'll do the cost calculation in big.Int and then convert (assuming result fits in int).
	cost := new(big.Int).Mul(a, big.NewInt(3))
	cost.Add(cost, b)

	return cost, true
}

// divisible checks if x is divisible by y
func divisible(x, y *big.Int) bool {
	// x mod y == 0
	mod := new(big.Int).Mod(x, y)
	return mod.Sign() == 0
}
