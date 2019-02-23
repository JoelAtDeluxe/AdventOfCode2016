package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	logic()
}

type Bot struct {
	Name       int
	HighVal    int
	LowVal     int
	LowTarget  GiveTarget
	HighTarget GiveTarget
}

type GiveTarget struct {
	Name        int
	IsOutputBin bool
}

var RED = "\033[0;31m"
var NC = "\033[0m"

func logic() {
	filename := "input.txt"

	bots := make([]Bot, 0)
	outputs := make([][]int, 21) // After looking at the input.

	instRe, _ := regexp.Compile(`bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)`)
	inputRe, _ := regexp.Compile(`value (\d+) goes to bot (\d+)`)

	parse := func(data []byte) {
		asStr := string(data)
		lines := strings.Split(asStr, "\n")
		updateBots := func(botName int) {
			if botName >= len(bots) {
				newBots := make([]Bot, botName+1)
				for i := range bots {
					newBots[i] = bots[i]
				}
				bots = newBots
			}

			if bots[botName] == (Bot{}) {
				bots[botName] = Bot{Name: botName}
			}
		}

		for _, line := range lines {
			if instRe.MatchString(line) {
				matches := instRe.FindStringSubmatch(line)
				botNameStr, lowType, lowName, highType, highName := matches[1], matches[2], matches[3], matches[4], matches[5]

				botName := toInt(botNameStr)
				updateBots(botName)

				bots[botName].LowTarget = GiveTarget{toInt(lowName), lowType == "output"}
				bots[botName].HighTarget = GiveTarget{toInt(highName), highType == "output"}

			} else if inputRe.MatchString(line) {
				matches := inputRe.FindStringSubmatch(line)
				value, botNameStr := matches[1], matches[2]

				botName := toInt(botNameStr)
				updateBots(botName)
				(&bots[botName]).giveVal(toInt(value))

			} else {
				fmt.Println("Unable to find re match for line: ", line)
			}
		}

	}
	err := readFile(filename, parse)

	if err != nil {
		fmt.Printf("Unable to read input: %v\n", filename)
		return
	}

	fmt.Println("Evaluating bots...")
	whosGotWhat(&bots, &outputs)
	for i, b := range bots {
		if b.LowVal > 0 {
			followChain(&bots, &outputs, i)
			break
		}
	}
	whosGotWhat(&bots, &outputs)

	fmt.Println("Done!")
}

func followChain(pBots *[]Bot, pOutputs *[][]int, botIndex int) {
	outputs := *pOutputs
	bot := &(*pBots)[botIndex]

	if (*bot).HighVal == 0 || (*bot).LowVal == 0 {
		return
	}
	// fmt.Print(bot.Name, "-->")

	if (*bot).HighVal == 61 && (*bot).LowVal == 17 {
		fmt.Print(RED)
		fmt.Println("Bot", (*bot).Name, "compares", (*bot).LowVal, "and", (*bot).HighVal)
		fmt.Print(NC)
	} else {
		// fmt.Println("Bot", bot.Name, "compares", bot.LowVal, "and", bot.HighVal)
	}

	moveVal := func(target GiveTarget, value *int) {
		valCopy := *value
		(*value) = 0
		if target.IsOutputBin {
			fmt.Println(valCopy, "going to output:", target.Name)
			outputs[target.Name] = append(outputs[target.Name], valCopy)
			whosGotWhat(pBots, pOutputs)
		} else {
			fmt.Println(valCopy, "going to bot:", target.Name)
			(*pBots)[target.Name].giveVal(valCopy)
		}
	}

	moveVal((*bot).LowTarget, &(*bot).LowVal)
	followChain(pBots, pOutputs, (*bot).LowTarget.Name)

	moveVal((*bot).HighTarget, &(*bot).HighVal)
	followChain(pBots, pOutputs, (*bot).HighTarget.Name)
}

func whosGotWhat(pBots *[]Bot, outputs *[][]int) {
	fmtVal := func(i int) string {
		if i == 0 {
			return "_"
		}
		return strconv.Itoa(i)
	}

	fmt.Println("??????????????????????????")
	fmt.Println("Outputs --")
	for i, output := range *outputs {
		if len(output) > 0 {
			fmt.Printf("%v => %v \n", i, output)
		}
	}

	fmt.Println("--------------------------")
	fmt.Println("Bots --")
	for _, bot := range *pBots {
		if bot.HighVal != 0 || bot.LowVal != 0 {
			fmt.Printf("%v : %v | %v \n", bot.Name, fmtVal(bot.HighVal), fmtVal(bot.LowVal))
		}
	}
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!")
}

func readFile(path string, parse func([]byte)) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	parse(data)
	return nil
}

func toInt(someInt string) int {
	v, err := strconv.Atoi(someInt)
	if err != nil {
		fmt.Println("Error reading int")
		return 0
	}
	return v
}

func (b *Bot) giveVal(val int) {
	// fmt.Println("Bot:", b.Name, "Is getting: ", val)
	if b.HighVal == 0 {
		b.HighVal = val
	} else if val > b.HighVal {
		b.LowVal = b.HighVal
		b.HighVal = val
	} else if b.HighVal == 0 && b.LowVal == 0 {
		fmt.Println("Got tooo many values!")
	} else {
		b.LowVal = val
	}
}
