package main

import (
	"fmt"
	"regexp"
	"strings"

	"./tooling"
)

func main() {
	Logic()
	// s := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	// fmt.Println("S: ", string(s))
	// tooling.Rotate(&s, 2, 0, len(s)-1, false)
	// tooling.Rotate(&s, 2, 0, len(s)-1, true)
	// tooling.Rotate(&s, 2, 0, len(s)-1, true)
	// tooling.Rotate(&s, 2, 0, len(s)-1, false)

	// fmt.Println("--------")
	// fmt.Println("S: ", string(s))

	// tooling.Rotate(&s, 2, 2, 5, false)
	// tooling.Rotate(&s, 2, 2, 5, true)
	// tooling.Rotate(&s, 2, 2, 5, true)
	// tooling.Rotate(&s, 2, 2, 5, false)
}

func Logic() {
	filename := "input.txt"

	program := make([]string, 0)

	parse := func(data []byte) {
		asStr := string(data)
		program = strings.Split(asStr, "\n")
	}
	err := tooling.ReadFile(filename, parse)

	if err != nil {
		fmt.Printf("Unable to read input: %v\n", filename)
		return
	}

	result := evaluate("abcdefgh", program)

	// Test values
	// program = []string{
	// 	"swap position 4 with position 0",
	// 	"swap letter d with letter b",
	// 	"reverse positions 0 through 4",
	// 	"rotate left 1 step",
	// 	"move position 1 to position 4",
	// 	"move position 3 to position 0",
	// 	"rotate based on position of letter b",
	// 	"rotate based on position of letter d",
	// }
	// result := evaluate("abcde", program)
	fmt.Println("Result:", result)
}

func evaluate(password string, program []string) string {
	swapPosRegex, _ := regexp.Compile(`swap position (\d+) with position (\d)`)
	swapLetRegex, _ := regexp.Compile(`swap letter ([a-z]) with letter ([a-z])`)
	rotLeftRegex, _ := regexp.Compile(`rotate left (\d+) steps?`)
	rotRightRegex, _ := regexp.Compile(`rotate right (\d+) steps?`)
	rotPosRegex, _ := regexp.Compile(`rotate based on position of letter ([a-z])`)
	revPosRegex, _ := regexp.Compile(`reverse positions (\d+) through (\d+)`)
	movPosRegex, _ := regexp.Compile(`move position (\d+) to position (\d+)`)

	scrambled := make([]rune, len(password))
	for i, v := range password {
		scrambled[i] = v
	}

	for i, stmt := range program {
		switch {
		case swapPosRegex.MatchString(stmt):
			matches := swapPosRegex.FindStringSubmatch(stmt)
			a, b := tooling.ToInt(matches[1]), tooling.ToInt(matches[2])
			tooling.SwapIdx(&scrambled, a, b)

		case swapLetRegex.MatchString(stmt):
			matches := swapLetRegex.FindStringSubmatch(stmt)
			a, b := matches[1], matches[2]
			aIndex, bIndex := tooling.FindLetter(&scrambled, a), tooling.FindLetter(&scrambled, b)
			tooling.SwapIdx(&scrambled, aIndex, bIndex)

		case rotLeftRegex.MatchString(stmt):
			matches := rotLeftRegex.FindStringSubmatch(stmt)
			amt := tooling.ToInt(matches[1])
			tooling.Rotate(&scrambled, amt, 0, len(scrambled)-1, true)

		case rotRightRegex.MatchString(stmt):
			matches := rotRightRegex.FindStringSubmatch(stmt)
			amt := tooling.ToInt(matches[1])
			tooling.Rotate(&scrambled, amt, 0, len(scrambled)-1, false)

		case rotPosRegex.MatchString(stmt):
			matches := rotPosRegex.FindStringSubmatch(stmt)
			letter := matches[1]
			letterIndex := tooling.FindLetter(&scrambled, letter)
			rotCount := 1
			if letterIndex >= 4 {
				rotCount++
			}
			rotCount += letterIndex
			tooling.Rotate(&scrambled, rotCount, 0, len(scrambled)-1, false)

		case revPosRegex.MatchString(stmt):
			matches := revPosRegex.FindStringSubmatch(stmt)
			start, end := tooling.ToInt(matches[1]), tooling.ToInt(matches[2])
			tooling.ReversePortion(&scrambled, start, end)

		case movPosRegex.MatchString(stmt):
			matches := movPosRegex.FindStringSubmatch(stmt)
			source, target := tooling.ToInt(matches[1]), tooling.ToInt(matches[2])

			if source < target {
				tooling.Rotate(&scrambled, 1, source, target, true)
			} else {
				tooling.Rotate(&scrambled, 1, source, target, false)
			}

		default:
			fmt.Println("Unknown operation: " + stmt)
		}
		fmt.Println("After", i, "pass:", string(scrambled))
	}

	return string(scrambled)
}
