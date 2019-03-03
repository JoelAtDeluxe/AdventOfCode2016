package old

import (
	"fmt"
	"strings"
	"time"

	"../tooling"
)

func Logic() {
	filename := "input_part2.txt"

	instructions := make([]string, 0)

	parse := func(data []byte) {
		asStr := string(data)
		instructions = strings.Split(asStr, "\n")
	}
	err := tooling.ReadFile(filename, parse)

	if err != nil {
		fmt.Printf("Unable to read input: %v\n", filename)
		return
	}

	start := time.Now()
	state := evaluate(instructions)
	duration := time.Since(start)
	fmt.Println(state["a"])
	fmt.Println("Finished in: ", duration.Seconds())
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
			if tooling.IsNum(val) {
				state[reg] = tooling.ToInt(val)
			} else {
				state[reg] = state[val]
			}
		case "jnz":
			val := components[1]
			if (tooling.IsNum(val) && val != "0") || state[val] != 0 {
				i += tooling.ToInt(components[2]) - 1
			}
		case "inc":
			state[components[1]]++
		case "dec":
			state[components[1]]--
		}
	}
	return state
}
