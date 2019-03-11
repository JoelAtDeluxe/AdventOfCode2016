package tooling

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func ToInt(someInt string) int {
	v, err := strconv.Atoi(someInt)
	if err != nil {
		fmt.Println("Error reading int")
		return 0
	}
	return v
}

func ReadFile(path string, parse func([]byte)) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	parse(data)
	return nil
}

func Mod(n int, modulo int) int {
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
