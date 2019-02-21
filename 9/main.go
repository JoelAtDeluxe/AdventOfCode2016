package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func main() {
	logic()
}

func logic() {
	filename := "real_input.txt"

	var encoded string

	parse := func(data []byte) {
		encoded = string(data)
	}
	err := readFile(filename, parse)

	if err != nil {
		fmt.Printf("Unable to read input: %v\n", filename)
		return
	}

	count := calcDecompressed(encoded)
	fmt.Println("Decoded length is: ", count)
}

func calcDecompressed(encoded string) int {
	// fmt.Println("Called with: ", encoded)
	quantityRegex, _ := regexp.Compile("(\\d+)x(\\d+)\\)")
	mode := "ECHO"
	count := 0
	for i := 0; i < len(encoded); i++ {
		ch := encoded[i]
		switch mode {
		case "ECHO":
			if ch == '(' {
				mode = "EVAL"
			} else {
				// fmt.Println("Ate: ", string(ch))
				count++
			}
		case "EVAL":
			nextCh := encoded[i]
			nextIndex := i
			for nextCh != ')' {
				nextCh = encoded[nextIndex]
				nextIndex++
			}
			matches := quantityRegex.FindStringSubmatch(encoded[i:nextIndex])
			charScan, repetitions := toInt(matches[1]), toInt(matches[2])
			// count += (charScan * repetitions)
			i = nextIndex + charScan - 1 // place cursor at end of repeated character
			partialCount := calcDecompressed(encoded[nextIndex : i + 1])
			// fmt.Println("adding: ", count, partialCount * repetitions)
			count += partialCount * repetitions
			mode = "ECHO"
		}
	}
	// fmt.Println("returning: ", count)
	return count
}

func toInt(someInt string) int {
	v, err := strconv.Atoi(someInt)
	if err != nil {
		fmt.Println("Error reading int")
		return 0
	}
	return v
}

func readFile(path string, parse func([]byte)) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	parse(data)
	return nil
}
