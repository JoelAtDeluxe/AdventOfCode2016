package main

import "strconv"

func main() {
	logic()
}

func isAWall(x, y int) bool {
	v := x*x + 3*x + 2*x*y + y + y*y
	v = v + 1358 // favorite number
	binary := strconv.FormatInt(int64(v), 2)
	count := 0
	for _, v := range binary {
		if v == '1'{
			count++
		}
	}
	return (count % 2 == 1)
}

func logic() {
	grid := buildGrid(40, 49)
	targetX, targetY := 31, 39
	playerX, playerY := 1, 1
	printGrid(&grid, playerX, playerY, targetX, targetY )
}

func buildGrid(w, h int) [][]bool {
	grid := make([][]bool, h)
	for i := range grid {
		grid[i] = make([]bool, w)
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			grid[i][j] = isAWall(j, i)
		}
	}
	return grid
}

func printGrid(grid *[][]bool, x, y, targetX, targetY int) {
	for y := range *grid {
		for x := range y {

		}
	} 
}