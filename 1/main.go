package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	logic()

	// targetLine := Line{Pair{1, 0}, Pair{4, 0}}
	// testLines := []Line{
	// 	Line{Pair{5, 5}, Pair{5, 10}}, //No //right out
	// 	Line{Pair{2, 3}, Pair{2, -3}}, //Yes
	// 	Line{Pair{1, 3}, Pair{1, -3}}, //Yes
	// 	Line{Pair{4, 3}, Pair{4, -3}}, //Yes
	// 	Line{Pair{2, 3}, Pair{2, 0}},  //Yes
	// 	Line{Pair{2, -3}, Pair{2, 0}},  //Yes

	// 	//near miss
	// 	Line{Pair{2, 3}, Pair{2, 1}},  //No //on top
	// 	Line{Pair{2, -1}, Pair{2, -5}},  //No // on bottom
	// 	Line{Pair{0, 3}, Pair{0, -3}},  //No //on left
	// 	Line{Pair{5, 3}, Pair{5, -3}},  //No // on right
	// }

	// for i, line := range testLines {
	// 	x, y, intersected := doesIntersect(targetLine, line)
	// 	if !intersected {
	// 		fmt.Printf("Line %v does not intersect\n", i)
	// 	} else {
	// 		fmt.Printf("Line %v does intersect at (%v, %v)\n", i, x, y)
	// 	}
	// }
}

//Pair is
type Pair struct {
	x int
	y int
}

//Line is
type Line struct {
	Origin Pair
	End    Pair
}

func logic() {
	filename := "input.txt"
	var directions []string
	parse := func(data []byte) {
		asStr := string(data)
		directions = strings.Split(asStr, ", ")
	}

	err := readFile(filename, parse)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", filename)
		return
	}

	e, n := collapseRoute(directions)
	fmt.Printf("Total Distance is: %v\n", intAbs(n)+intAbs(e))
	fmt.Printf("Go North: %v // Go East: %v\n", n, e)

	e, n, intersected := checkRoute(directions)
	if intersected {
		fmt.Printf("Total Distance to Repeated Node is: %v\n", intAbs(n)+intAbs(e))
		fmt.Printf("Go North: %v // Go East: %v\n", n, e)
	} else {
		fmt.Println("No intersections found")
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

func checkRoute(directions []string) (int, int, bool) {
	steper := stepBuilder()
	paths := make([]Line, 0)
	lastCoord := Pair{0, 0}

	for _, s := range directions {
		x, y := steper(s)
		p := Pair{x, y}
		if lastCoord != (Pair{}) {
			paths = append(paths, Line{lastCoord, p})
		}
		lastCoord = p
		x, y, intersected := getIntersection(paths)
		if intersected {
			return x, y, true
		}
	}
	return 0, 0, false
}

func getIntersection(lines []Line) (int, int, bool) {
	if len(lines) < 4 { // based on the way the puzzle works, we don't need to worry about this until we have 4 lines
		return 0, 0, false
	}
	lastIndex := len(lines) - 1
	lastMod := lastIndex % 2
	lastLine := lines[lastIndex]
	crossingLines := make([]Line, 0, len(lines)/2+1)

	//Based on the puzzle, since we start facing north, we know that every even index is going to be a horizontal line (i.e. heading east or west)
	// while all other lines are north/south/vertical lines
	for i, line := range lines {
		if i%2 != lastMod {
			crossingLines = append(crossingLines, line)
		}
	}
	crossingLines = crossingLines[:len(crossingLines)-1] //The last one semi-intersects (at the point) In reality, it will never intersect, so we exclude it here
	xDiff := lastLine.End.x - lastLine.Origin.x
	yDiff := lastLine.End.y - lastLine.Origin.y

	sortLines(&crossingLines, xDiff, yDiff) //sort crossingLines from closest to origin to furthest from origin

	for _, line := range crossingLines {
		if xDiff != 0 { // Line is moving left or right, so Y isn't changing, but X is
			if betweenPoints(line.Origin.y, lastLine.Origin.y, line.End.y) &&
				betweenPoints(lastLine.Origin.x, line.Origin.x, lastLine.End.x) {
				return line.Origin.x, lastLine.Origin.y, true
			}
		} else { // Line is moving up or down, so X isn't changing, but Y is
			if betweenPoints(line.Origin.x, lastLine.Origin.x, line.End.x) &&
				betweenPoints(lastLine.Origin.y, line.Origin.y, lastLine.End.y) {
				return lastLine.Origin.x, line.Origin.y, true
			}
		}
	}
	return 0, 0, false
}

func doesIntersect(target, testVal Line) (int, int, bool) {
	isHorz := (target.Origin.x - target.End.x) != 0

	if isHorz { // Line is moving left or right, so Y isn't changing, but X is
		if betweenPoints(testVal.Origin.y, target.Origin.y, testVal.End.y) &&
			betweenPoints(target.Origin.x, testVal.Origin.x, target.End.x) {
			return testVal.Origin.x, target.Origin.y, true
		}
	} else { // Line is moving up or down, so X isn't changing, but Y is
		if betweenPoints(testVal.Origin.x, target.Origin.x, testVal.End.x) &&
			betweenPoints(target.Origin.y, testVal.Origin.y, target.End.y) {
			return target.Origin.x, testVal.Origin.y, true
		}
	}
	return 0, 0, false
}

func betweenPoints(min, value, max int) bool {
	if min > max {
		max, min = min, max
	}
	return min <= value && value <= max
}

func sortLines(lines *[]Line, xDiff, yDiff int) {
	smallYFirst := func(i, j int) bool { return (*lines)[i].Origin.y < (*lines)[j].Origin.y }
	largeYFirst := func(i, j int) bool { return (*lines)[i].Origin.y > (*lines)[j].Origin.y }
	smallXFirst := func(i, j int) bool { return (*lines)[i].Origin.x < (*lines)[j].Origin.x }
	largeXFirst := func(i, j int) bool { return (*lines)[i].Origin.x > (*lines)[j].Origin.x }

	if xDiff > 0 { // moving LEFT
		sort.Slice(*lines, smallYFirst)
	} else if xDiff < 0 { // moving RIGHT
		sort.Slice(*lines, largeYFirst)
	} else if yDiff > 0 { // moving UP
		sort.Slice(*lines, smallXFirst)
	} else { //Must be moving DOWN
		sort.Slice(*lines, largeXFirst)
	}
}

func collapseRoute(directions []string) (int, int) {
	x, y := 0, 0
	steper := stepBuilder()

	for _, s := range directions {
		x, y = steper(s)
	}
	return x, y
}

func stepBuilder() func(string) (int, int) {
	orientation := 0
	x := 0
	y := 0

	return func(s string) (int, int) {
		distance, _ := strconv.Atoi(s[1:])

		if spin := string(s[:1]); spin == "R" {
			orientation = mod(orientation+1, 4)
		} else {
			orientation = mod(orientation-1, 4)
		}
		x, y = move(x, y, orientation, distance)
		return x, y
	}
}

func move(x, y, dir, dist int) (int, int) {
	switch dir {
	case 0: // NORTH
		y += dist
	case 1: // EAST
		x += dist
	case 2: // SOUTH
		y -= dist
	case 3:
		x -= dist
	}
	return x, y
}

func linearSearch(needle Pair, haystack []Pair) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func mod(n int, modulo int) int {
	if n >= modulo {
		return n % modulo
	} else if n < 0 {
		offset := (-n) % modulo
		answer := modulo - offset
		if answer == modulo { // we could end up in a situation where offset is 0, then modulo - 0 = modulo = 0
			return 0
		}
		return answer
	}
	return n
}

func intAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
