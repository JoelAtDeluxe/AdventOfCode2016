package main

import "fmt"

type Disc struct {
	CurrentPosition int
	TotalPositions  int
	ClearPosition   int
}

func (d *Disc) advance(amt int) {
	d.CurrentPosition = (d.CurrentPosition + amt) % d.TotalPositions
}

func main() {
	logic()
}

func logic() {
	// discs := puzzleTestInput()
	discs := puzzleInput()
	t := 0
	for ; ; t++ {
		clear, _, _ := isClear(discs, t)
		if clear {
			break
		}
		// advancement := 1
		// if collidedDisc > 0 {
		// 	advancement = fakeLCM(discs[:collidedDisc])
		// }

		// fmt.Println(colidedDiscIndex, discPosition)
		// guess next number:
		// Position + lcm(passed discs) +

	}
	fmt.Println("Drop in at t =", t)
}

func deepCopyDiscs(discs []Disc) []Disc {
	copy := make([]Disc, len(discs))
	for i := range discs {
		copy[i] = Disc{discs[i].CurrentPosition, discs[i].TotalPositions, discs[i].ClearPosition }
	}
	return copy
}

func isClear(discs []Disc, t int) (bool, int, int) {
	copy := deepCopyDiscs(discs)	

	for i := range copy {
		copy[i].advance(t)
	}
	for i := 0; i < len(copy); i++ {
		if copy[i].CurrentPosition != copy[i].ClearPosition {
			// fmt.Println("Will stop on disc", i+1)
			return false, i, copy[i].CurrentPosition
		}
	}
	return true, 0, 0
}

// All of my numbers are prime, so the LCM of prime number is going to be their product
func fakeLCM(nums []int) int {
	product := 1

	for _, v := range nums {
		product *= v
	}
	return product
}

// Opting to hard code here, as parsing is pointless for such a small dataset
func puzzleInput() []Disc {
	discs := []Disc{
		{CurrentPosition: 5, TotalPositions: 17},
		{CurrentPosition: 8, TotalPositions: 19},
		{CurrentPosition: 1, TotalPositions: 7},
		{CurrentPosition: 7, TotalPositions: 13},
		{CurrentPosition: 1, TotalPositions: 5},
		{CurrentPosition: 0, TotalPositions: 3},
		{CurrentPosition: 0, TotalPositions: 11},
	}
	for i := range discs {
		discs[i].ClearPosition = mod(discs[i].TotalPositions-1-i, discs[i].TotalPositions)
	}

	return discs
}

func puzzleTestInput() []Disc {
	discs := []Disc{
		{CurrentPosition: 4, TotalPositions: 5},
		{CurrentPosition: 1, TotalPositions: 2},
	}
	for i := range discs {
		discs[i].ClearPosition = mod(discs[i].TotalPositions-1-i, discs[i].TotalPositions)
	}
	return discs
}

func mod(n int, modulo int) int {
	if n >= modulo {
		return n % modulo
	} else if n < 0 {
		offset := (-n) % modulo
		answer := modulo - offset
		if answer == modulo { // we could end up in a situation where offset is 0, then modulo - 0 = modulo = 0
			return 0
		}
		return answer
	}
	return n
}
