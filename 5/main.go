package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	numThreads := 1
	fns := genHashInput("ugkcyxxp", numThreads)
	// solvePartOne(fns[0])
	solvePartTwo(fns[0])
}

func solvePartOne(hashGen func()string)  {
	found := 0
	for i := 0; i < 100000000; i++ {
		hash := gimmeAHash(hashGen())
		if hash[0:5] == "00000" {
			fmt.Println(hash[5:6], " @ ", hash, " => ", i)
			found++
			if found >= 8 {
				fmt.Println("Done!")
				break
			}
		}
	}
}

func solvePartTwo(hashGen func() string ) {
	password := "________"
	for i := 0; i < 100000000; i++ {
		hash := gimmeAHash(hashGen())
		if hash[0:5] == "00000" {
			pos, err := strconv.Atoi(hash[5:6])
			if err == nil && pos < 8 {
				val := hash[6:7]

				if password[pos] == '_' {
					password = fmt.Sprintf("%v%v%v", password[:pos], val, password[pos+1:])
					fmt.Println(password)
					if !strings.Contains(password, "_") {
						break
					}
				}
			}
		}
	}
}

func magic(base string, start, inc int) func() string {
	return func() string {
		rtn := fmt.Sprintf("%v%v", base, start)
		start += inc
		return rtn
	}
}

func genHashInput(base string, numThreads int) []func() string {
	rtn := make([]func() string, numThreads)

	for i := range rtn {
		rtn[i] = magic(base, i, numThreads)
	}
	return rtn
}

func gimmeAHash(body string) string {
	hash := md5.Sum([]byte(body))

	rtn := hex.EncodeToString(hash[:])
	return rtn
}
