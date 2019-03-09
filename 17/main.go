package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	logic()
}

func logic() {
	// input := puzzleTestInput()
	input := puzzleInput()

	successes := []string{}
	evaluatePath(input, []int{0, 0}, &successes)
	success := successes[0]
	for _, s := range successes[1:] {
		if len(s) > len(success) {
			success = s
		}
	}

	// Solution to part 1
	// success := ""
	// evaluatePathSmallest(input, []int{0, 0}, &success)

	fmt.Println(success)
	fmt.Println(len(success))
}

func evaluatePath(key string, position []int, successes *[]string) {
	if position[0] == 3 && position[1] == 3 {
		*successes = append(*successes, key[8:])
		return
	}
	hash := gimmeAHash(key)
	paths := isOpen(hash[:4])                 // in order U, D, L, R
	paths[0] = paths[0] && (position[1] != 0) // up
	paths[1] = paths[1] && (position[1] != 3) // down
	paths[2] = paths[2] && (position[0] != 0) // left
	paths[3] = paths[3] && (position[0] != 3) // right

	if paths[0] {
		newPosition := []int{position[0], position[1] - 1}
		evaluatePath(fmt.Sprintf("%v%v", key, "U"), newPosition, successes)
	}
	if paths[1] {
		newPosition := []int{position[0], position[1] + 1}
		evaluatePath(fmt.Sprintf("%v%v", key, "D"), newPosition, successes)
	}
	if paths[2] {
		newPosition := []int{position[0] - 1, position[1]}
		evaluatePath(fmt.Sprintf("%v%v", key, "L"), newPosition, successes)
	}
	if paths[3] {
		newPosition := []int{position[0] + 1, position[1]}
		evaluatePath(fmt.Sprintf("%v%v", key, "R"), newPosition, successes)
	}
	return
}

func evaluatePathSmallest(key string, position []int, success *string) {
	if position[0] == 3 && position[1] == 3 {
		newRoute := key[8:]
		if len(newRoute) < len(*success) || *success == "" {
			(*success) = newRoute
		}
	}
	if len(key)-8 > len(*success) && *success != "" {
		return
	}
	hash := gimmeAHash(key)
	paths := isOpen(hash[:4])                 // in order U, D, L, R
	paths[0] = paths[0] && (position[1] != 0) // up
	paths[1] = paths[1] && (position[1] != 3) // down
	paths[2] = paths[2] && (position[0] != 0) // left
	paths[3] = paths[3] && (position[0] != 3) // right

	if paths[0] {
		newPosition := []int{position[0], position[1] - 1}
		evaluatePathSmallest(fmt.Sprintf("%v%v", key, "U"), newPosition, success)
	}
	if paths[1] {
		newPosition := []int{position[0], position[1] + 1}
		evaluatePathSmallest(fmt.Sprintf("%v%v", key, "D"), newPosition, success)
	}
	if paths[2] {
		newPosition := []int{position[0] - 1, position[1]}
		evaluatePathSmallest(fmt.Sprintf("%v%v", key, "L"), newPosition, success)
	}
	if paths[3] {
		newPosition := []int{position[0] + 1, position[1]}
		evaluatePathSmallest(fmt.Sprintf("%v%v", key, "R"), newPosition, success)
	}
	return
}

func isOpen(code string) []bool {
	rtn := make([]bool, 4)
	for i := range rtn {
		rtn[i] = strings.Contains("bcdef", code[i:i+1])
	}
	return rtn
}

// Opting to hard code here, as parsing is pointless for such a small dataset
func puzzleInput() string {
	return "vkjiggvb"
}

func puzzleTestInput() string {
	return "ihgpwlah" // DDRRRD
	// return "kglvqrro" // DDUDRLRRUDRD
	// return "ulqzkmiv" // DRURDRUDDLLDLUURRDULRLDUUDDDRR
}

func gimmeAHash(body string) string {
	hash := md5.Sum([]byte(body))

	rtn := hex.EncodeToString(hash[:])
	return rtn
}
