package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	logic()
	// check := func(s string, exp string) {
	// 	_, abba := hasAbba(s)
	// 	if abba != exp {
	// 		fmt.Printf("Mismatch on \"%v\" expected: [%v] got: [%v]n", s, exp, abba)
	// 	}
	// }

	// check("abba", "abba")
	// check("zabba", "abba")
	// check("abbaz", "abba")
	// check("zabbaz", "abba")
	// check("rabbaz", "abba")
	// check("abc123abbaz", "abba")
	// check("abxba", "")
	// check("aba", "")
	// check("abc123", "")
	// check("mmmm", "")
	// check("nmmmmn", "")
	// check("manna", "anna")
	// check("abba123", "abba")
}

func logic() {
	filename := "real_input.txt"

	var ips []string

	parse := func(data []byte) {
		asStr := string(data)
		ips = strings.Split(asStr, "\n")
	}
	err := readFile(filename, parse)

	if err != nil {
		fmt.Printf("Unable to read input: %v\n", filename)
		return
	}

	abbaIPs := make([]string, 0, len(ips))
	abaBabIPs := make([]string, 0, len(ips))

	// g, b := splitIP(ips[1])
	// fmt.Println(g, b)
	// return

	for _, ip := range ips {
		good, bad := splitIP(ip)
		// fmt.Println(good, bad)
		// return
		if isAbbaCompatible(good, bad) {
			abbaIPs = append(abbaIPs, ip)
		}
		if isAbABaBCompatible(good, bad) {
			abaBabIPs = append(abaBabIPs, ip)
		}
	}

	fmt.Println("number of TLS ips is: ", len(abbaIPs))
	fmt.Println("number of SSL ips is: ", len(abaBabIPs))
}

func isAbbaCompatible(good, bad []string) bool {
	for _, b := range bad {
		has, _ := hasAbba(b)
		if has {
			return false
		}
	}

	for _, g := range good {
		has, _ := hasAbba(g)
		if has {
			return true
		}
	}
	return false
}

func isAbABaBCompatible(good, bad []string) bool {
	abas := findAbas(good)
	for _, aba := range abas {
		bab := fmt.Sprintf("%v%v%v", string(aba[1]), string(aba[0]), string(aba[1]))
		for _, b := range bad {
			if strings.Contains(b, bab) {
				return true
			}
		}
	}
	return false
}

func findAbas(good []string) []string {
	abas := make([]string, 0)
	for _, g := range good {
		for i, ch := range g[:len(g)-2] {
			if ch == rune(g[i+2]) && ch != rune(g[i+1]) {
				abas = append(abas, g[i:i+3])
			}
		}
	}
	return abas
}

func highlightAbba(s string) string {
	abbas := findAllAbba(s)
	RED := "\033[0;31m"
	NC := "\033[0m"
	rtn := strings.Replace(s, "[", fmt.Sprintf("%v%v", RED, "["), -1)
	rtn = strings.Replace(rtn, "]", fmt.Sprintf("%v%v", "]", NC), -1)

	for _, abba := range abbas {
		new := fmt.Sprintf("___%v___", abba)
		rtn = strings.Replace(rtn, abba, new, -1)
	}
	return rtn
}

func hasAbba(s string) (bool, string) {
	for i, ch := range s[:len(s)-3] {
		endCh := rune(s[i+3])
		if endCh == ch {
			if rune(s[i+1]) == rune(s[i+2]) && rune(s[i+1]) != ch {
				return true, s[i : i+4]
			}
		}
	}
	return false, ""
}

func findAllAbba(s string) []string {
	rtn := make([]string, 0)
	for i, ch := range s[:len(s)-3] {
		endCh := rune(s[i+3])
		if endCh == ch {
			if rune(s[i+1]) == rune(s[i+2]) && rune(s[i+1]) != ch {
				rtn = append(rtn, string(s[i:i+4]))
			}
		}
	}
	return rtn
}

func splitIP(ip string) ([]string, []string) {
	good := make([]string, 0)
	bad := make([]string, 0)

	parts := strings.Split(ip, "[")
	good = append(good, parts[0])
	for _, p := range parts[1:] {
		segments := strings.Split(p, "]") //assumes the last block does not end with a ], which my input does not have
		if len(segments) == 2 {
			bad = append(bad, segments[0])
			good = append(good, segments[1])
		} else {
			good = append(good, segments[0])
		}
	}
	return good, bad
}

func readFile(path string, parse func([]byte)) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	parse(data)
	return nil
}
