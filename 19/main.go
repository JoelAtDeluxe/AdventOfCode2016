package main

import (
	"fmt"
	"math"
)

func main() {
	var answer int
	answer = solution1(3012210)
	answer = solution2(3012210)
	fmt.Println("Winning seat: ", answer)

	// solution finding:
	// for i := 1; i <= 81; i++ {
	// 	fmt.Print(i, "==> ")
	// 	solution2BruteForce(i)
	// }
}

func solution1(num int) int {
	var shift uint = 1

	for (num / (2 << shift)) > 0 {
		shift++
	}
	shift-- // get back to less than num

	remainder := num - (2 << shift)
	return (2 * remainder) + 1
}

func solution2(num int) int {
	var pow float64 = 1
	fNum := float64(num)
	for fNum > math.Pow(3.0, pow) {
		pow++
	}
	// 3^pow is now strictly greater than num
	upperBound := int(math.Pow(3.0, pow))
	lowerBound := int(math.Pow(3.0, pow-1))
	midpoint := upperBound - lowerBound
	winner := 0
	if num == upperBound || num == lowerBound { // if we're exactly 3^n
		winner = num
	} else if num <= midpoint {
		winner = num - lowerBound
	} else {
		winner = (num-midpoint)*2 + (midpoint - lowerBound)
	}
	return winner
}

func solution2BruteForce(num int) int {

	//set up seats
	elfs := make([]int, num)
	for i := range elfs {
		elfs[i] = (i + 1)
	}
	remaining := num

	for i := 0; remaining > 1; i++ {
		i = i % num
		if elfs[i] < 0 { // skip eliminated elfs
			continue
		}

		//take a turn
		stealFrom := (remaining / 2)
		//advance that number of (eligable) elfs
		k := 0
		var stealIndex int
		for j := 0; j < stealFrom; j++ {
			k++
			stealIndex = (k + i) % num
			if elfs[stealIndex] < 0 {
				j-- // don't count this iteration
			}
		}
		elfs[stealIndex] = elfs[stealIndex] * -1
		remaining--
	}
	//figure out the winner
	rtn := 0
	for _, v := range elfs {
		if v > 0 {
			rtn = v
			fmt.Println("Winner was at seat: ", v)
			break
		}
	}
	return rtn
}
