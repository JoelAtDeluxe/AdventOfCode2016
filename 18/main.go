package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	logic()
}

func logic() {
	filename := "input.txt"
	var rawRow string
	// numRows := 40 // solution 1
	numRows := 400000 // solution 2 //100 x 400,000 bool vals => 40M bools => 40M bytes => 40Mbytes. Big, but not _that_ big

	err := readFile(filename, func(data []byte) {
		rawRow = string(data)
	})

	if err != nil {
		fmt.Printf("Unable to read input: %v\n", filename)
		return
	}

	grid := buildGrid(rawRow, numRows)
	sum := sumSafeTiles(grid)

	fmt.Println("There are",sum,"safe tiles in the grid")
}

func sumSafeTiles(grid [][]bool) int {
	count := 0
	for i := range grid {
		for _, val := range grid[i] {
			if val {
				count++
			}
		}
	}
	return count
}

func printGrid(grid [][]bool) {
	for i := range grid {
		for _, safe := range grid[i] {
			ch := "^"
			if safe {
				ch = "."
			}
			fmt.Print(ch)
		}
		fmt.Println()
	}
}

func buildGrid(firstRow string, numRows int) [][]bool {
	grid := make([][]bool, numRows)
	for i := range grid {
		grid[i] = make([]bool, len(firstRow))
	}

	// populate first row
	for i, ch := range firstRow {
		grid[0][i] = (ch == '.')
	}

	isSafe := func(row int, col int) bool {

		lSafe := col == 0 || grid[row-1][col-1]
		cSafe := grid[row-1][col]
		rSafe := col == len(firstRow)-1 || grid[row-1][col+1]

		if (!lSafe && !cSafe && rSafe) || // rule 1: Its left and center tiles are traps, but its right tile is not
			(!cSafe && !rSafe && lSafe) || // rule 2: Its center and right tiles are traps, but its left tile is not
			(!lSafe && cSafe && rSafe) || // rule 3: Only its left tile is a trap
			(!rSafe && cSafe && lSafe) { // rule 4: Only its right tile is a trap
			return false
		}

		return true
	}

	// populate lower rows

	for i := range grid[1:] {
		gIndex := 1 + i
		for j := range grid[gIndex] {
			grid[gIndex][j] = isSafe(gIndex, j)
		}
	}

	return grid
}

func readFile(path string, parse func([]byte)) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	parse(data)
	return nil
}
