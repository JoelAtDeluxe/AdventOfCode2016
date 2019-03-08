package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
)

func main() {
	logic()
}

type FoundHash struct {
	Hash       string
	FoundIndex int
}

func logic() {
	salt := "ihaygndm"
	// salt := "abc"
	possibleHashes := make(map[rune][]FoundHash)
	knownKeys := make([]FoundHash, 0, 128)

	i := 0
	for stop := false; !stop; i++ { // We probably should start at 0, but my guess is that likely 0 is not going to work anyway
		toHashVal := fmt.Sprintf("%v%v", salt, i)
		// hash := gimmeAHash(toHashVal)  // Part 1 solution
		hash := stretchHash(toHashVal) // part 2 solution
		threeMatch, fiveMatch := analyzeHash(hash)

		if len(fiveMatch) > 0 {
			for _, match := range fiveMatch {
				matchingIndexes := possibleHashes[match]
				for _, v := range matchingIndexes {
					if i-v.FoundIndex <= 1000 {
						fmt.Printf("[%v] Adding hash as Key: %v (Found at: %v, verified at: %v with char: %v in %v)\n", len(knownKeys), v.Hash, v.FoundIndex, i, string(match), hash)
						knownKeys = append(knownKeys, v)
						if len(knownKeys) >= 64 {
							stop = true
						}
					}
				}
				possibleHashes[match] = []FoundHash{}
			}
		}
		if threeMatch != '_' {
			possibleHashes[threeMatch] = append(possibleHashes[threeMatch], FoundHash{hash, i})
		}
	}

	sort.Slice(knownKeys, func(i, j int) bool{ return knownKeys[i].FoundIndex < knownKeys[j].FoundIndex})

	fmt.Println("All keys found by index: ", knownKeys[63].FoundIndex)
}

func analyzeHash(hash string) (rune, []rune) {
	threeMatch := '_'
	fiveMatch := make([]rune, 0, 6)

	matchedIndex := 0

	for i, ch := range hash[1:] {
		chIndex := (i + 1)
		var repCount int
		if ch == rune(hash[matchedIndex]) {
			repCount = (chIndex + 1) - matchedIndex
		} else {
			matchedIndex = chIndex
		}

		if repCount == 3 && threeMatch == '_' {
			threeMatch = ch
		}
		if repCount == 5 {
			fiveMatch = append(fiveMatch, ch)
		}
	}

	return threeMatch, fiveMatch
}

func gimmeAHash(body string) string {
	hash := md5.Sum([]byte(body))

	rtn := hex.EncodeToString(hash[:])
	return rtn
}

func stretchHash(body string) string {
	copy := body
	for i := 0; i < 2017; i++ {
		copy = gimmeAHash(copy)
	}
	return copy
}
