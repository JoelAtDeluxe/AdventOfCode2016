package tooling

import (
	"fmt"
	"regexp"
	"strings"
)

//isNum is _much_ faster, but abuses the puzzle input slightly
func IsNum(s string) bool {
	return strings.Contains("0123456789-", s[0:1])
}


// Slow, but accurate
func IsNumReal(s string) bool {
	rtn, err := regexp.Match(`\d+`, []byte(s))
	if err != nil {
		fmt.Println("Got an error: ", err)
	}
	return rtn
}
