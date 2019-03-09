package main

import (
	"fmt"
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

func isAWall(x, y int) bool {
	v := x*x + 3*x + 2*x*y + y + y*y
	v = v + 1358 // favorite number
	binary := strconv.FormatInt(int64(v), 2)
	count := 0
	for _, v := range binary {
		if v == '1' {
			count++
		}
	}
	return (count%2 == 1)
}

func logic() {
	grid := buildGrid(40, 49)
	player := Position{1, 1}
	target := Position{31, 39}
	// dijkstraTranverse(&grid, player, target)
	printGrid(&grid, player, target)
}

func buildGrid(columns, rows int) [][]bool {
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

func printGrid(grid *[][]bool, startPos Position, targetPos Position) {
	wall, open := "#", "."
	for r, _ := range *grid {
		for c, _ := range (*grid)[r] {
			ch := open
			if (*grid)[r][c] {
				ch = wall
			}
			fmt.Print(ch)
		}
		fmt.Println()
	}
}

func dijkstraTranverse(grid *[][]bool, startPos Position, targetPos Position) {
	openSet := make([]OpenCell, 0, len(*grid)*len((*grid)[0]))
	openSet = append(openSet, OpenCell{P: Position{startPos.X, startPos.Y}, D: 0})

	mark := func(index int) {
		openSet[index].Traversed = true
		cell := openSet[index]
		p := cell.P
		distToRoot := cell.D
		if p.Y > 0 && (*grid)[p.X][p.Y-1] == false { // Up
			openSet = append(openSet, OpenCell{P: Position{p.X + 0, p.Y - 1}, D: distToRoot + 1})
		}
		if (*grid)[p.X][p.Y+1] == false { // Down // todo: this might need to grow
			openSet = append(openSet, OpenCell{P: Position{p.X + 0, p.Y + 1}, D: distToRoot + 1})
		}
		if p.X > 0 && (*grid)[p.X-1][p.Y] == false { // Left
			openSet = append(openSet, OpenCell{P: Position{p.X - 1, p.Y + 0}, D: distToRoot + 1})
		}
		if (*grid)[p.X+1][p.Y] == false { // Right
			openSet = append(openSet, OpenCell{P: Position{p.X + 1, p.Y + 0}, D: distToRoot + 1})
		}
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

	for {
		minIndex := findNext()
		if minIndex > -1 {
			cell := openSet[minIndex]
			fmt.Println("Looking at:", cell.P.X, cell.P.Y)
			if cell.P == targetPos {
				fmt.Println("Found the path! Distance: ", cell.D)
				break
			}
			mark(minIndex)
		} else {
			fmt.Println("Ran out of checks!")
			break // nothing else to do
		}
	}
}
