package main

import (
	"fmt"
	"math"
	"strconv"
)

type Position struct {
	X int
	Y int
}

type OpenCell struct {
	P         Position
	D         int
	Traversed bool
}

func main() {
	logic()
}

func makeIsWall(favNum int) func(int, int) bool {
	return func(x, y int) bool {
		v := x*x + 3*x + 2*x*y + y + y*y
		v = v + favNum // favorite number
		binary := strconv.FormatInt(int64(v), 2)
		count := 0
		for _, v := range binary {
			if v == '1' {
				count++
			}
		}
		return (count%2 == 1)
	}
}

func logic() {
	//test stuff
	// isWallFunc := makeIsWall(10)
	// grid := buildGrid(11, 8, isWallFunc) // jts: So, I either need a way to grow this dynamically (probably the better idea), or I need to just make a bigger grid... I'll go with the latter for now
	// player := Position{1, 1}
	// target := Position{7, 4}

	// real stuff
	isWallFunc := makeIsWall(1358)
	grid := buildGrid(50, 50, isWallFunc)
	player := Position{1, 1}
	target := Position{31, 39}

	dijkstraTranverse(&grid, player, target)
	// printGrid(&grid, player, target, []OpenCell{})
}

func buildGrid(columns, rows int, isAWall func(int, int) bool) [][]bool {
	grid := make([][]bool, rows)
	for i := range grid {
		grid[i] = make([]bool, columns)
	}

	for c := 0; c < columns; c++ {
		for r := 0; r < rows; r++ {
			grid[r][c] = isAWall(c, r)
		}
	}
	return grid
}

func printGrid(grid *[][]bool, startPos Position, targetPos Position, openSet []OpenCell) {
	rowHeader := func(max int) func() string {
		i := 0
		headerLength := len(strconv.Itoa(max))
		return func() string {
			header := make([]rune, headerLength)
			num := strconv.Itoa(i)
			for i := range num {
				header[len(header)-1-i] = rune(num[len(num)-1-i])
			}
			rtn := fmt.Sprintf("%v ", string(header))
			i++
			return rtn
		}
	}(len(*grid))

	colMax := len((*grid)[0])
	printRowHeader(colMax)

	traversedCells := make([]Position, 0)
	openCells := make([]Position, 0)
	for _, cell := range openSet {
		if cell.Traversed {
			traversedCells = append(traversedCells, cell.P)
		} else {
			openCells = append(openCells, cell.P)
		}
	}

	inSlice := func(p Position, possibilities []Position) bool {
		for _, maybe := range possibilities {
			if maybe == p {
				return true
			}
		}
		return false
	}

	for r, _ := range *grid {
		fmt.Print(asColor(rowHeader(), Yellow))
		for c, _ := range (*grid)[r] {
			var ch string
			switch {
			case (*grid)[r][c]: // wall
				ch = "#"
			case (startPos == Position{c, r}):
				ch = asColor("S", Green)
			case (targetPos == Position{c, r}):
				ch = asColor("E", Red)
			case inSlice(Position{c, r}, traversedCells):
				ch = asColor("?", Purple)
			case inSlice(Position{c, r}, openCells):
				ch = asColor(":", BrownOrange)
			default:
				ch = "."
			}
			fmt.Print(ch)
		}
		fmt.Println()
	}
}

func printRowHeader(numcols int) {
	headerRows := len(strconv.Itoa(numcols))
	for i := headerRows - 1; i >= 0; i-- {
		b10 := int(math.Pow10(i))
		line := make([]byte, numcols)
		for j := 0; j < len(line); j++ {
			something := (j / b10)
			offset := something % 10
			if offset == 0 && i != 0 && something == 0 {
				line[j] = ' '
			} else {
				line[j] = byte('0' + offset)
			}
		}
		fmt.Println(asColor(fmt.Sprintf("  %v", string(line)), Yellow))
	}

}

func dijkstraTranverse(grid *[][]bool, startPos Position, targetPos Position) {
	openSet := make([]OpenCell, 0, len(*grid)*len((*grid)[0]))
	openSet = append(openSet, OpenCell{P: Position{startPos.X, startPos.Y}, D: 0})

	isAWall := func(p Position) bool {
		if p.Y < 0 || p.X < 0 {
			return true
		}
		if p.Y >= len(*grid) || p.X >= len((*grid)[0]) {
			// or grow the grid
			return true
		}
		// fmt.Println(p)
		return (*grid)[p.Y][p.X]
	}

	inOpenSet := func(p Position) bool {
		for _, cell := range openSet {
			if cell.P == p {
				return true
			}
		}
		return false
	}

	findNext := func() int { // this probably not necessary -- they should all be in order by distance already
		minIndex := -1
		for i, item := range openSet {
			if !item.Traversed && (minIndex == -1 || item.D < openSet[minIndex].D) {
				minIndex = i
			}
		}
		return minIndex
	}

	mark := func(index int) {
		openSet[index].Traversed = true
		cell := openSet[index]
		p := cell.P
		newDist := cell.D + 1
		UP := Position{p.X, p.Y - 1}
		DOWN := Position{p.X, p.Y + 1}
		LEFT := Position{p.X - 1, p.Y}
		RIGHT := Position{p.X + 1, p.Y}

		if newDist > 50 { //solution 2 limit
			return
		}

		if !isAWall(UP) && !inOpenSet(UP) {
			openSet = append(openSet, OpenCell{P: UP, D: newDist})
		}
		if !isAWall(DOWN) && !inOpenSet(DOWN) {
			openSet = append(openSet, OpenCell{P: DOWN, D: newDist})
		}
		if !isAWall(LEFT) && !inOpenSet(LEFT) {
			openSet = append(openSet, OpenCell{P: LEFT, D: newDist})
		}
		if !isAWall(RIGHT) && !inOpenSet(RIGHT) {
			openSet = append(openSet, OpenCell{P: RIGHT, D: newDist})
		}
	}

	// reader := bufio.NewReader(os.Stdin)
	printGrid(grid, startPos, targetPos, openSet)
	fmt.Println("----------------")
	for {
		minIndex := findNext()
		if minIndex > -1 {
			cell := openSet[minIndex]
			// fmt.Println("Looking at:", cell.P.X, cell.P.Y)
			if cell.P == targetPos {
				fmt.Println("Found the path! Distance: ", cell.D) // solution 1
				break
			}
			mark(minIndex)
			// printGrid(grid, startPos, targetPos, openSet)
			// reader.ReadString('\n')
		} else {
			fmt.Println("Ran out of checks! Checked:", len(openSet), "cells")
			break // nothing else to do
		}
	}
}
