package day4

import (
	"bufio"
	"github.com/agrison/go-commons-lang/stringUtils"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const scratchCardsFile = "./day4/scratch-cards.txt"

var scratchCardPattern = regexp.MustCompile("Card\\s+(\\d+): (.*) \\| (.*)")

func readScratchCards() []ScratchCard {
	file, err := os.Open(scratchCardsFile)
	if err != nil {
		log.Panicf("unable to read file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scratchCards := make([]ScratchCard, 0)
	for scanner.Scan() {
		line := scanner.Text()

		matches := scratchCardPattern.FindStringSubmatch(line)
		if len(matches) != 4 {
			log.Panicf("invalid format line: %s", line)
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Panicf("unable to parse id: %s", line)
		}

		scratchCards = append(scratchCards, ScratchCard{ id: id, winningNumbers: parseIntArray(matches[2]), cardNumbers: parseIntArray(matches[3]) })
	}


	return scratchCards
}

func parseIntArray(str string) []int {
	output := make([]int, 0)
	for _, value := range strings.Split(str, " ") {
		if stringUtils.IsBlank(value) {
			continue
		}

		intValue, err := strconv.Atoi(value)
		if err != nil {
			log.Panicf("unable to parse value: %s", value)
		}

		output = append(output, intValue)
	}

	return output
}

func getSampleScratchCards() []ScratchCard {
	return []ScratchCard{
		{ id: 1, winningNumbers: []int{41, 48, 83, 86, 17}, cardNumbers: []int {83, 86,  6, 31, 17,  9, 48, 53} },
		{ id: 2, winningNumbers: []int{13, 32, 20, 16, 61}, cardNumbers: []int {61, 30, 68, 82, 17, 32, 24, 19} },
		{ id: 3, winningNumbers: []int{ 1, 21, 53, 59, 44}, cardNumbers: []int {69, 82, 63, 72, 16, 21, 14,  1} },
		{ id: 4, winningNumbers: []int{41, 92, 73, 84, 69}, cardNumbers: []int {59, 84, 76, 51, 58,  5, 54, 83} },
		{ id: 5, winningNumbers: []int{87, 83, 26, 28, 32}, cardNumbers: []int {88, 30, 70, 12, 93, 22, 82, 36} },
		{ id: 6, winningNumbers: []int{31, 18, 13, 56, 72}, cardNumbers: []int {74, 77, 10, 23, 35, 67, 36, 11} },
	}
}
