package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	logic()
}

var simpleGrid = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

var wall = '❌'

var complexGrid = [][]rune{
	{'❌', '❌', '1', '❌', '❌'},
	{'❌', '2', '3', '4', '❌'},
	{'5', '6', '7', '8', '9'},
	{'❌', 'A', 'B', 'C', '❌'},
	{'❌', '❌', 'D', '❌', '❌'},
}

func logic() {
	filename := "input.txt"
	var directions []string
	parse := func(data []byte) {
		asStr := string(data)
		directions = strings.Split(asStr, "\n")
	}

	err := readFile(filename, parse)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", filename)
		return
	}

	code := getSimplePasscode(directions)
	fmt.Printf("Passcode was: %v\n", code)

	complexCode := getComplexPasscode(directions)
	fmt.Printf("Passcode is: %v\n", complexCode)
}

func readFile(path string, parse func([]byte)) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	parse(data)
	return nil
}

func getSimplePasscode(directions []string) []int {
	code := make([]int, len(directions))
	startX, startY := 1, 1 // starting at 5

	for i, d := range directions {
		startX, startY = moveSimple(startX, startY, d)
		code[i] = simpleGrid[startY][startX]
	}
	return code
}

func getComplexPasscode(directions []string) []string {
	code := make([]string, len(directions))
	startX, startY := 0, 2 // starting at 5

	for i, d := range directions {
		startX, startY = moveComplex(startX, startY, d, complexGrid)
		code[i] = string(complexGrid[startY][startX])
	}
	return code
}

func moveSimple(curX, curY int, directions string) (int, int) {
	max := len(simpleGrid) - 1

	moveOnce := func(point, max int) int {
		if point < 0 {
			return 0
		} else if point > max {
			return max
		}
		return point
	}

	for _, c := range directions {
		// fmt.Printf("%v --%v-> ", grid[curY][curX], string(c))
		actX, actY := interpretDirection(c)
		curX = moveOnce(curX+actX, max)
		curY = moveOnce(curY+actY, max)
	}
	// fmt.Printf("(%v)\n", grid[curY][curX])
	return curX, curY
}

func interpretDirection(letter rune)(int, int) {
	actX, actY := 0, 0
	switch letter {
	case 'U':
		actY = -1
	case 'D':
		actY = 1
	case 'L':
		actX = -1
	case 'R':
		actX = 1
	}
	return actX, actY
}

func moveComplex(startX, startY int, directions string, grid [][]rune) (int, int) {
	moveOnce := func(oldX, oldY, xMove, yMove int) (int, int) {
		newX := oldX + xMove
		newY := oldY + yMove
		if newX < 0 || newX == len(grid) || newY < 0 || newY == len(grid) {
			return oldX, oldY
		}
		if grid[newY][newX] == wall {
			return oldX, oldY
		}

		return newX, newY
	}

	for _, c := range directions {
		actX, actY := interpretDirection(c)
		startX, startY = moveOnce(startX, startY, actX, actY)
	}
	return startX, startY
}
