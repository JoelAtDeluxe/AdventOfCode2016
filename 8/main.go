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
	filename := "input.txt"

	var instructions []string

	parse := func(data []byte) {
		asStr := string(data)
		instructions = strings.Split(asStr, "\n")
	}
	err := readFile(filename, parse)

	if err != nil {
		fmt.Printf("Unable to read input: %v\n", filename)
		return
	}

	grid := [6][50]bool{}
	// applyRect(&grid, 3, 2)
	// applyRowTranslate(&grid, 0, 10)
	// applyColTranslate(&grid, 10, 2)

	rectRegex, _ := regexp.Compile(`rect (\d+)x(\d+)`)
	rotRowRegex, _ := regexp.Compile(`rotate row y=(\d+) by (\d+)`)
	rotColRegex, _ := regexp.Compile(`rotate column x=(\d+) by (\d+)`)

	for _, step := range instructions {
		if rectRegex.MatchString(step) {
			matches := rectRegex.FindStringSubmatch(step)
			applyRect(&grid, toInt(matches[1]), toInt(matches[2]))
		}
		if rotRowRegex.MatchString(step) {
			matches := rotRowRegex.FindStringSubmatch(step)
			applyRowTranslate(&grid, toInt(matches[1]), toInt(matches[2]))
		}
		if rotColRegex.MatchString(step) {
			matches := rotColRegex.FindStringSubmatch(step)
			applyColTranslate(&grid, toInt(matches[1]), toInt(matches[2]))
		}
	}

	drawGrid(&grid)

	fmt.Println("num highlighted: ", countActive(&grid))
}

func countActive(grid *[6][50]bool) int {
	count := 0
	for i :=0 ; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == true {
				count++
			}
		}
	}
	return count
}

func toInt(someInt string) int {
	v, err := strconv.Atoi(someInt)
	if err != nil {
		fmt.Println("Error reading int")
		return 0
	}
	return v
}

func applyRowTranslate(grid *[6][50]bool, rowIdx, amt int) {
	line := (*grid)[rowIdx]
	absRotate := amt % len(line)
	copy := line[:]
	for i := 0; i < len(line); i++ {
		(*grid)[rowIdx][(i+absRotate)%len(line)] = copy[i]
	}
}

func applyColTranslate(grid *[6][50]bool, colIdx, amt int) {
	absRotate := amt % len(grid)
	copy := make([]bool, len(grid))
	for i := range grid {
		copy[i] = grid[i][colIdx]
	}
	for i := range grid {
		(*grid)[(i+absRotate)%len(grid)][colIdx] = copy[i]
	}
}

func applyRect(grid *[6][50]bool, numCols, numRows int) {
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			(*grid)[i][j] = true
		}
	}
}

func drawGrid(grid *[6][50]bool) {
	for _, line := range *grid {
		runes := make([]rune, len(line))
		for i, ch := range line {
			if ch == true {
				runes[i] = '#'
			} else {
				runes[i] = '.'
			}
		}
		fmt.Println(string(runes))
	}
}

func readFile(path string, parse func([]byte)) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	parse(data)
	return nil
}
