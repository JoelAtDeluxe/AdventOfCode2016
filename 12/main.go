package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	// "./old"
	"./tooling"
)

func main() {
	// old.Logic()
	Logic()  //10x faster than the old solution (100 times faster than the python3 version)
}

func Logic() {
	filename := "input_part2.txt"

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

	engine := compile(program)

	start := time.Now()
	evaluate(&engine)
	duration := time.Since(start)
	fmt.Println(engine.Memory[engine.MemMap["a"]])
	fmt.Println("Finished in: ", duration.Seconds())
}

func compile(program []string) Engine {
	engine := Engine{
		Commands: make([][3]int, len(program)),
		Memory:   make([]int, 0),
		MemMap:   make(map[string]int),
	}

	getNextRegister := func(s string) int {
		rtn, ok := engine.MemMap[s]
		if !ok {
			rtn = len(engine.Memory)
			engine.MemMap[s] = rtn
			engine.Memory = append(engine.Memory, 0)
		}

		return rtn
	}

	for i, line := range program {
		var command [3]int

		components := strings.Split(line, " ")

		switch components[0] {
		case "inc":
			register := getNextRegister(components[1])
			command = [3]int{IncReg, register, 0}
		case "dec":
			register := getNextRegister(components[1])
			command = [3]int{DecReg, register, 0}
		case "cpy":
			from := components[1]
			toInx := getNextRegister(components[2])
			if val, err := strconv.Atoi(from); err == nil {
				command = [3]int{CpyVal, val, toInx}
			} else {
				command = [3]int{CpyReg, getNextRegister(from), toInx}
			}
		case "jnz":
			cmpVal := components[1]
			direction := tooling.ToInt(components[2])
			if val, err := strconv.Atoi(cmpVal); err == nil {
				command = [3]int{JnzVal, val, direction}
			} else {
				command = [3]int{JnzReg, getNextRegister(cmpVal), direction}
			}
		}

		engine.Commands[i] = command
	}

	return engine
}

func evaluate(eng *Engine) {
	pc := 0
	progLen := len(eng.Commands)

	for pc < progLen {
		npc := pc + 1

		command := eng.Commands[pc]
		switch command[0] {
		case IncReg:
			eng.Memory[command[1]]++
		case DecReg:
			eng.Memory[command[1]]--
		case CpyReg:
			eng.Memory[command[2]] = eng.Memory[command[1]]
		case CpyVal:
			eng.Memory[command[2]] = command[1]
		case JnzVal:
			if command[1] != 0 {
				npc = pc + command[2]
			}
		case JnzReg:
			if eng.Memory[command[1]] != 0 {
				npc = pc + command[2]
			}
		}

		pc = npc
	}
}

const (
	IncReg = iota + 1
	DecReg
	CpyReg
	CpyVal
	JnzReg
	JnzVal
)

type Engine struct {
	Memory   []int
	MemMap   map[string]int
	Commands [][3]int
}
