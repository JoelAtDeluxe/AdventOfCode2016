package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type IntRange struct {
	Lower uint64
	Upper uint64
}

func main() {
	needed := true
	var blocks []IntRange
	var err error
	if needed {
		// do some pre-processing
		blocks, err = loadRangeFile("input.txt")
		if err != nil {
			fmt.Printf("Unable to read input\n")
			return
		}
		blocks = SortBlacklist(blocks)
		writeBlacklist(blocks, "sortedInput.txt")
		blocks = mergeSortedBlacklist(blocks)
		writeBlacklist(blocks, "consolodated.txt")
	} else {
		blocks, err = loadRangeFile("consolodated.txt")
		if err != nil {
			fmt.Printf("Unable to read input\n")
			return
		}
	}
	fmt.Println("First available IP: ", solution1(blocks))
	fmt.Println("Total available IP: ", solution2(blocks))
}

func loadRangeFile(filename string) ([]IntRange, error) {
	var blocks []IntRange

	err := readFile(filename, func(data []byte) {
		blocks = parseFile(data)
	})

	return blocks, err
}

func SortBlacklist(blocks []IntRange) []IntRange {
	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].Lower < blocks[j].Lower
	})
	return blocks
}

func writeBlacklist(blocks []IntRange, filename string) {
	contents := make([]string, len(blocks))
	for i, block := range blocks {
		contents[i] = fmt.Sprintf("%v-%v", block.Lower, block.Upper)
	}
	writeFile("sortedInput.txt", strings.Join(contents, "\n"))
}

func mergeSortedBlacklist(blocks []IntRange) []IntRange {
	trueBlackList := make([]IntRange, 0)
	trueBlackList = append(trueBlackList, blocks[0])
	for _, block := range blocks[1:] {
		lastBlock := &trueBlackList[len(trueBlackList)-1]
		if block.Lower <= (lastBlock.Upper + 1) {
			max := lastBlock.Upper
			if block.Upper > max {
				max = block.Upper
			}
			lastBlock.Upper = max
		} else {
			trueBlackList = append(trueBlackList, block)
		}
	}
	return trueBlackList
}

func parseFile(data []byte) []IntRange {
	asStr := string(data)
	ipAddrs := strings.Split(asStr, "\n")
	blocks := make([]IntRange, len(ipAddrs))
	for i, v := range ipAddrs {
		components := strings.Split(v, "-")
		low, _ := strconv.ParseUint(components[0], 10, 0)
		high, _ := strconv.ParseUint(components[1], 10, 0)
		if low > high { //make sure left is always the lower bound
			low, high = high, low
		}
		blocks[i] = IntRange{low, high}
	}
	return blocks
}

func solution1(blocks []IntRange) uint64 {
	return blocks[0].Upper + 1
}

func solution2(blocks []IntRange) uint64 {
	var sum uint64
	var upperRange uint64 = 4294967295
	for i := range blocks[1:] {
		trueIndex := i + 1
		sum += blocks[trueIndex].Lower - blocks[trueIndex-1].Upper - 1
	}
	if blocks[len(blocks)-1].Upper < upperRange {
		sum += upperRange - blocks[len(blocks)-1].Upper - 1
	}
	return sum
}

func readFile(path string, parse func([]byte)) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	parse(data)
	return nil
}

func writeFile(path string, content string) error {
	return ioutil.WriteFile(path, []byte(content), 0666)
}
