package main

import (
	"fmt"
	"strings"
)

func main() {
	logic()
}

func logic() {
	// input := puzzleTestInput()
	// stopAt := 20
	input := puzzleInput()
	// stopAt := 272 // problem 1
	stopAt := 35651584 // problem 2

	for len(input) < stopAt {
		input = step(input)
	}
	input = input[:stopAt]

	fmt.Println(Checksum(input))
}

func step(data string) string {
	copy := strings.Map(func(letter rune) rune { return letter }, data)
	reversed := Reverse(copy)
	notted := strings.Map(func(letter rune) rune {
		if letter == '0' {
			return '1'
		}
		return '0'
	}, reversed)
	return fmt.Sprintf("%v0%v", data, notted)
}

func Checksum(data string) string {
	checksum := make([]rune, len(data)/2)
	recurse := true
	if (len(checksum) % 2) == 1 {
		recurse = false
	}
	for i := 0; i < len(checksum); i++ {
		if data[2*i] == data[2*i+1] {
			checksum[i] = '1'
		} else {
			checksum[i] = '0'
		}
	}

	if recurse {
		return Checksum(string(checksum))
	} else {
		return string(checksum)
	}
}

// borrowed from https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Opting to hard code here, as parsing is pointless for such a small dataset
func puzzleInput() string {
	return "01000100010010111"
}

func puzzleTestInput() string {
	return "10000"
}
