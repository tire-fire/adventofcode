package main

import (
	"fmt"
	"math/big"
)

// Machine structure
type Machine struct {
	Ax, Ay *big.Int
	Bx, By *big.Int
	Px, Py *big.Int
}

func main() {
	// Original machines (Part One input)
	// Now we add 10,000,000,000,000 to Px and Py for each machine as specified in part two
	offset := big.NewInt(10000000000000)

	machines := []Machine{
		newMachine(94, 34, 22, 67, 8400, 5400, offset),
		newMachine(26, 66, 67, 21, 12748, 12176, offset),
		newMachine(17, 86, 84, 37, 7870, 6450, offset),
		newMachine(69, 23, 27, 71, 18641, 10279, offset),
	}

	winnableCosts := []int{}

	for _, m := range machines {
		cost, ok := solveMachine(m)
		if ok {
			winnableCosts = append(winnableCosts, cost)
		}
	}

	// According to the puzzle statement, now only machines #2 and #4 are solvable.
	// So we should find exactly those two solutions.
	totalCost := 0
	for _, c := range winnableCosts {
		totalCost += c
	}

	fmt.Println("Fewest tokens to win all possible prizes:", totalCost)
}

// newMachine creates a Machine with given parameters and applies the offset to Px, Py.
func newMachine(Ax, Ay, Bx, By, Px, Py int64, offset *big.Int) Machine {
	PxBig := big.NewInt(Px)
	PyBig := big.NewInt(Py)
	PxBig.Add(PxBig, offset)
	PyBig.Add(PyBig, offset)
	return Machine{
		Ax: big.NewInt(Ax), Ay: big.NewInt(Ay),
		Bx: big.NewInt(Bx), By: big.NewInt(By),
		Px: PxBig, Py: PyBig,
	}
}

// solveMachine tries to solve the linear system for the given machine.
// Returns (cost, true) if solvable and (0, false) if not.
func solveMachine(m Machine) (int, bool) {
	// D = Ax*By - Ay*Bx
	D := new(big.Int).Sub(
		new(big.Int).Mul(m.Ax, m.By),
		new(big.Int).Mul(m.Ay, m.Bx),
	)

	if D.Sign() == 0 {
		// No unique solution
		return 0, false
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
		return 0, false
	}

	a := new(big.Int).Div(a_num, D)
	b := new(big.Int).Div(b_num, D)

	// a and b must be nonnegative
	if a.Sign() < 0 || b.Sign() < 0 {
		return 0, false
	}

	// cost = 3*a + b
	// a,b are possibly huge, but fit in int64 only if very large might cause issues.
	// We'll do the cost calculation in big.Int and then convert (assuming result fits in int).
	cost := new(big.Int).Mul(a, big.NewInt(3))
	cost.Add(cost, b)

	if cost.IsInt64() {
		return int(cost.Int64()), true
	} else {
		// If the cost doesn't fit in int (which would be extremely large),
		// we can return the big number as a string or handle differently.
		// Problem doesn't specify extremely large inputs beyond int64.
		// We'll assume it fits.
		return 0, false
	}
}

// divisible checks if x is divisible by y
func divisible(x, y *big.Int) bool {
	// x mod y == 0
	mod := new(big.Int).Mod(x, y)
	return mod.Sign() == 0
}

