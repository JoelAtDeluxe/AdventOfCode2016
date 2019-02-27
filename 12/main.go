package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	logic()
}

func logic() {
	filename := "input_part2.txt"

	instructions := make([]string, 0)

	parse := func(data []byte) {
		asStr := string(data)
		instructions = strings.Split(asStr, "\n")
	}
	err := readFile(filename, parse)

	if err != nil {
		fmt.Printf("Unable to read input: %v\n", filename)
		return
	}

	state := evaluate(instructions)
	fmt.Println(state["a"])
	fmt.Println("Done!")
}

func evaluate(instructions []string) map[string]int {
	state := make(map[string]int)
	parsedInstructions := make([][]string, len(instructions))
	for i := range instructions {
		parsedInstructions[i] = strings.Split(instructions[i], " ")
	}
	for i := 0; i < len(instructions); i++ {
		// fmt.Println("Executing step: ", i, "(", instructions[i], ") state =>", state)
		components := parsedInstructions[i]

		switch components[0] {
		case "cpy":
			val := components[1]
			reg := components[2]
			if isNum(val) {
				state[reg] = toInt(val)
			} else {
				state[reg] = state[val]
			}
		case "jnz":
			val := components[1]
			if (isNum(val) && val != "0") || state[val] != 0 {
				i += toInt(components[2]) - 1
			}
		case "inc":
			state[components[1]]++
		case "dec":
			state[components[1]]--
		}
	}
	return state
}

func readFile(path string, parse func([]byte)) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	parse(data)
	return nil
}

//isNum is _much_ faster, but abuses the puzzle input slightly
func isNum(s string) bool {
	return strings.Contains("0123456789-", s[0:1])
}

// Slow, but accurate
func isNumReal(s string) bool {
	rtn, err := regexp.Match(`\d+`, []byte(s))
	if err != nil {
		fmt.Println("Got an error: ", err)
	}
	return rtn
}

func toInt(someInt string) int {
	v, err := strconv.Atoi(someInt)
	if err != nil {
		fmt.Println("Error reading int")
		return 0
	}
	return v
}
