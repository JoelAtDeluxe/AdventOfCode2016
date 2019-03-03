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

func IsNumAlt(s string) bool {
	if !strings.Contains("0123456789-", s[0:1]) {
		return false
	}
	if len(s) > 1 {
		for ch := range s[1:] {
			if !strings.Contains("0123456789", string(ch)) {
				return false
			}
		}
	}
	return true
}

// Slow, but accurate
func IsNumReal(s string) bool {
	rtn, err := regexp.Match(`\d+`, []byte(s))
	if err != nil {
		fmt.Println("Got an error: ", err)
	}
	return rtn
}
