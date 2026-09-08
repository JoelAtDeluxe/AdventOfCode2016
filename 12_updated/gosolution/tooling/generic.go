package tooling

import (
	"fmt"
	"os"
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
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	parse(data)
	return nil
}
