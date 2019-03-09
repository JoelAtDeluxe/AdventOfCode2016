package main

import "fmt"

func main() {
	var answer int
	//answer = solution1(3012210)
	answer = solution2(3012210)
	fmt.Println("Solution should be: ", answer)
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
	return 0
}
