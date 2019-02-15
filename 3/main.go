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
	// fmt.Println(isValidTriangle([]int{5, 10, 25}))
	// fmt.Println(isValidTriangle([]int{15, 10, 25}))
	// fmt.Println(isValidTriangle([]int{16, 10, 25}))
	// fmt.Println("---")
	// fmt.Println(isValidTriangle([]int{10, 25, 5}))
	// fmt.Println(isValidTriangle([]int{10, 25, 15}))
	// fmt.Println(isValidTriangle([]int{10, 25, 16}))
	// fmt.Println("---")
	// fmt.Println(isValidTriangle([]int{25, 10, 5}))
	// fmt.Println(isValidTriangle([]int{25, 10, 15}))
	// fmt.Println(isValidTriangle([]int{25, 10, 16}))
	// fmt.Println("---")
	// fmt.Println(isValidTriangle([]int{25, 25, 1}))
	// fmt.Println(isValidTriangle([]int{25, 1, 25}))
	// fmt.Println(isValidTriangle([]int{1, 25, 25}))
	// fmt.Println("---")
	// fmt.Println(isValidTriangle([]int{1, 1, 2}))
}

func logic() {
	filename := "input.txt"
	var triangles [][]int

	// parseRows := func(data []byte) {
	// 	asStr := string(data)
	// 	lines := strings.Split(asStr, "\n")
	// 	triangles = make([][]int, len(lines))
	// 	re, _ := regexp.Compile(`\s*(\d+)\s*(\d+)\s*(\d+)\s*`)

	// 	for i, line := range lines {
	// 		matches := re.FindStringSubmatch(line)
	// 		triangles[i] = strArrayToIntArray(matches[1:])
	// 	}
	// }
	// err := readFile(filename, parseRows)

	parseColumns := func(data []byte) {
		asStr := string(data)
		lines := strings.Split(asStr, "\n")
		triangles = make([][]int, len(lines))
		re, _ := regexp.Compile(`\s*(\d+)\s*(\d+)\s*(\d+)\s*`)

		groups := make([][]int, 3)
		for i, line := range lines {
			matches := re.FindStringSubmatch(line)
			groups[i%3] = strArrayToIntArray(matches[1:])

			if i%3 == 2 {
				triangles[i-2] = []int{groups[0][0], groups[1][0], groups[2][0]}
				triangles[i-1] = []int{groups[0][1], groups[1][1], groups[2][1]}
				triangles[i-0] = []int{groups[0][2], groups[1][2], groups[2][2]}
			}

		}
	}

	err := readFile(filename, parseColumns)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", filename)
		return
	}

	realTriangles := make([][]int, 0, len(triangles))

	for _, triangle := range triangles {
		if isValidTriangle(triangle) {
			realTriangles = append(realTriangles, triangle)
		}
	}

	fmt.Printf("I count %v real triangles\n", len(realTriangles))
}

func strArrayToIntArray(strings []string) []int {
	ints := make([]int, len(strings))
	for i, s := range strings {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func readFile(path string, parse func([]byte)) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	parse(data)
	return nil
}

func isValidTriangle(lengths []int) bool {
	// mi, mv, err := max(lengths)

	// if err != nil {
	// 	return false
	// }
	// for i, v := range lengths {
	// 	if i != mi {
	// 		mv = mv - v
	// 	}
	// }
	// return mv < 0
	return lengths[0] < lengths[1]+lengths[2] &&
		lengths[1] < lengths[0]+lengths[2] &&
		lengths[2] < lengths[1]+lengths[0]
}

func max(vals []int) (int, int, error) {
	if len(vals) == 0 {
		return 0, 0, fmt.Errorf("list isn't big enough")
	}

	mi, mv := 0, vals[0]
	for i, v := range vals[1:] {
		if v > mv {
			mi, mv = i+1, v
		}
	}
	return mi, mv, nil
}
