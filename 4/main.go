package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"sort"
	"strings"
)

func main() {
	logic()

	// m := make(map[int]string)

	// m[1] = "A"

	// fmt.Println(m[1])
	// fmt.Println(m[2])

}

func logic() {
	filename := "input.txt"
	sectorSum := 0
	lines := []string{}
	parse := func(data []byte) {
		asStr := string(data)
		lines = strings.Split(asStr, "\n")
	}
	err := readFile(filename, parse)

	if err != nil {
		fmt.Printf("Unable to read input: %v\n", filename)
		return
	}

	validRooms := make([]Room, 0, len(lines))
	for _, line := range lines {
		words := strings.Split(line, "-")
		lastIndex := len(words) - 1
		roomName := words[:lastIndex]
		valid, sector := isValid(roomName, words[lastIndex])
		if valid {
			validRooms = append(validRooms, Room{strings.Join(roomName, " "), sector})
		}
	}

	for _, v := range validRooms {
		sectorSum += v.Sector
	}

	for _, v := range validRooms {
		(&v).decrypt()

		if strings.Contains( v.Name, "north") {
			fmt.Println(v.Name, v.Sector)
		}
	}

	fmt.Printf("Sector sum is: %v\n", sectorSum)
}

func (r *Room) decrypt() {
	offset := (*r).Sector % 26
	word := ""
	for _, ch := range (*r).Name {
		revised := ' '
		if ch != ' '  {
			base := int(ch - 96)
			revised = rune(((base + offset) % 26) + 96)
		}
		word = fmt.Sprintf("%v%v", word, string(revised))
	}
	(*r).Name = word
}

func isValid(words []string, checksumWithSector string) (bool, int) {
	letters := make(map[string] int)
	re, _ := regexp.Compile(`(\d+)\[([a-z]{5})\]`)
	matches := re.FindStringSubmatch(checksumWithSector)
	sector := matches[1]
	checksum := matches[2]

	for _, word := range words {
		countLetters(word, &letters)
	}

	calcChecksum := getChecksum(&letters)
	if calcChecksum == checksum {
		val, _ := strconv.Atoi(sector)
		return true, val
	}
	return false, 0
}

type LetterCount struct {
	Letter string
	Count int
}

type Room struct {
	Name string
	Sector int
}

func getChecksum(letters *map[string]int) string {	
	countLetters := make(map[int]string)
	for letter, count := range *letters {
		countLetters[count] = countLetters[count] + letter
	}
	unsorted := make([]LetterCount, len(countLetters))
	i := 0
	for count, letter := range countLetters {
		unsorted[i] = LetterCount{letter, count}
		i++
	}

	sort.Slice(unsorted, func(i, j int ) bool {
		return unsorted[i].Count > unsorted[j].Count
	})

	for i, _ := range unsorted {
		unnecessary := strings.Split(unsorted[i].Letter, "")
		sort.Slice(unnecessary, func(j, k int) bool {
			return rune(unnecessary[j][0]) < rune(unnecessary[k][0])
		})
		unsorted[i].Letter = strings.Join(unnecessary, "")
	}

	rtn := ""
	for _, s := range unsorted {
		rtn = fmt.Sprintf("%v%v", rtn, s.Letter)
		if len(rtn) >= 5 {
			break
		}
	}

	return rtn[:5]
	// Some letters are fixed, some are not
}

func countLetters(word string, letterMap *map[string]int) {
	for _, letter := range word {
		(*letterMap)[string(letter)] += 1
	}
}

func readFile(path string, parse func([]byte)) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	parse(data)
	return nil
}

