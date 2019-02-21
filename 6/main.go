package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	logic()
}

func logic() {
	filename := "input.txt"

	var messages []string

	parse := func(data []byte) {
		asStr := string(data)
		messages = strings.Split(asStr, "\n")
	}
	err := readFile(filename, parse)

	if err != nil {
		fmt.Printf("Unable to read input: %v\n", filename)
		return
	}

	// junk := [8]map[rune] int{}

	// for _, msg := range messages {
	// 	for i, c := range msg {
	// 		junk[i][c] += 1
	// 	}
	// }

	longWords := transposeStringArray(messages)
	target := make([]rune, len(longWords))
	for i, word := range longWords {
		// target[i] = mostOrLeastOccurances(word, true)
		target[i] = mostOrLeastOccurances(word, false)
	}
	fmt.Println("Password is: ", string(target))
}

func mostOrLeastOccurances(word []rune, most bool) rune {
	charCounter := make(map[rune]int)
	for _, c := range word {
		charCounter[c]++
	}
	
	var max int
	var maxC rune
	first := true

	for char, count := range charCounter {
		if first {
			first = false
			max, maxC = count, char
			continue
		}

		if most && count > max {
			max, maxC = count, char
		} else if !most && count < max {
			max, maxC = count, char
		}
	
	}
	return maxC
}

func transposeStringArray(messages []string) [][]rune {
	tranposed := make([][]rune, len(messages[0]))
	for i := range tranposed {
		tranposed[i] = make([]rune, len(messages))
	}

	for i, msg := range messages {
		for j, c := range msg {
			tranposed[j][i] = c
		}
	}
	return tranposed
}

func readFile(path string, parse func([]byte)) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	parse(data)
	return nil
}
