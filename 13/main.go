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
	isWallFunc := makeIsWall(10)
	grid := buildGrid(10, 7, isWallFunc)
	player := Position{1, 1}
	target := Position{7, 4}

	// real stuff
	// isWallFunc := makeIsWall(1358)
	// grid := buildGrid(40, 49)
	// player := Position{1, 1}
	// target := Position{31, 39}

	dijkstraTranverse(&grid, player, target)
	// printGrid(&grid, player, target)
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

func printGrid(grid *[][]bool, startPos Position, targetPos Position) {
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

	for r, _ := range *grid {
		fmt.Print(asColor(rowHeader(), Yellow))
		for c, _ := range (*grid)[r] {
			ch := "."
			if (*grid)[r][c] {
				ch = "#"
			} else if (startPos == Position{c, r}) {
				ch = asColor("S", Green)
			} else if (targetPos == Position{c, r}) {
				ch = asColor("E", Red)
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
