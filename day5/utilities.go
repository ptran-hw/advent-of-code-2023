package day5

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const seedDataFile = "./day5/initial-seed.txt"
const gardenDataFile = "./day5/garden-map.txt"

var rulePattern = regexp.MustCompile("^\\d+ \\d+ \\d+$")

func readSeedData() []int {
	file, err := os.Open(seedDataFile)
	if err != nil {
		log.Panicf("unable to read file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		log.Panic("invalid seed data file")
	}

	line := scanner.Text()

	seeds := make([]int, 0)
	for _, token := range strings.Split(line, " ") {
		intValue, err := strconv.Atoi(token)
		if err != nil {
			log.Panicf("unable to parse seed value: %s", token)
		}

		seeds = append(seeds, intValue)
	}

	return seeds
}

func readGardenData() [][]ConversionRule {
	file, err := os.Open(gardenDataFile)
	if err != nil {
		log.Panicf("unable to read file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	conversionGroups := make([][]ConversionRule, 0)
	for scanner.Scan() {
		for !rulePattern.MatchString(scanner.Text()) {
			scanner.Scan()
		}

		conversionRules := make([]ConversionRule, 0)
		for rulePattern.MatchString(scanner.Text()) {
			rawRule := strings.Split(rulePattern.FindString(scanner.Text()), " ")
			if len(rawRule) != 3 {
				log.Panicf("invalid rules format: %s", scanner.Text())
			}

			destinationStart, err := strconv.Atoi(rawRule[0])
			if err != nil {
				log.Panicf("unable to parse rule destinationStart value: %s", rawRule[0])
			}

			sourceStart, err := strconv.Atoi(rawRule[1])
			if err != nil {
				log.Panicf("unable to parse rule sourceStart value: %s", rawRule[1])
			}

			rangeLength, err := strconv.Atoi(rawRule[2])
			if err != nil {
				log.Panicf("unable to parse rule rangeLength value: %s", rawRule[2])
			}

			conversionRules = append(conversionRules, ConversionRule{start: sourceStart, end: sourceStart + rangeLength - 1, offset: destinationStart - sourceStart})
			scanner.Scan()
		}
		conversionGroups = append(conversionGroups, conversionRules)
	}

	return conversionGroups
}

func getSampleSeedValues() []int {
	return []int{79, 14, 55, 13}
}

func getSampleConversionGroups() [][]ConversionRule {
	return [][]ConversionRule{
		{
			{start: 98, end: 99, offset: -48},
			{start: 50, end: 97, offset: 2},
		},
		{
			{start: 52, end: 53, offset: -15},
			{start: 15, end: 51, offset: -15},
			{start: 0, end: 14, offset: 39},
		},
		{
			{start: 53, end: 60, offset: -4},
			{start: 11, end: 52, offset: -11},
			{start: 7, end: 10, offset: 50},
			{start: 0, end: 6, offset: 42},
		},
		{
			{start: 25, end: 94, offset: -7},
			{start: 18, end: 24, offset: 70},
		},
		{
			{start: 77, end: 99, offset: -32},
			{start: 64, end: 76, offset: 4},
			{start: 45, end: 63, offset: 36},
		},
		{
			{start: 69, end: 69, offset: -69},
			{start: 0, end: 68, offset: 1},
		},
		{
			{start: 93, end: 96, offset: -37},
			{start: 56, end: 92, offset: 4},
		},
	}
}

// returns overlapping range, prior non-overlapping range, after non-overlapping range
func splitOverlappingRange(valueRange, ruleRange []int) ([]int, []int, []int) {
	switch {
	case valueRange == nil:
		return nil, nil, nil
	case valueRange[1] < ruleRange[0]:
		return nil, valueRange, nil
	case valueRange[0] < ruleRange[0] && valueRange[1] <= ruleRange[1]:
		return []int{ruleRange[0], valueRange[1]}, []int{valueRange[0], ruleRange[0] - 1}, nil
	case valueRange[0] < ruleRange[0] && ruleRange[1] < valueRange[1]:
		return []int{ruleRange[0], valueRange[1]}, []int{valueRange[0], ruleRange[0] - 1}, []int{ruleRange[1] + 1, valueRange[1]}
	case ruleRange[0] <= valueRange[0] && valueRange[1] <= ruleRange[1]:
		return valueRange, nil, nil
	case ruleRange[0] <= valueRange[0] && valueRange[0] <= ruleRange[1] && ruleRange[1] < valueRange[1]:
		return []int{valueRange[0], ruleRange[1]}, nil, []int{ruleRange[1] + 1, valueRange[1]}
	case ruleRange[1] < valueRange[0]:
		return nil, nil, valueRange
	}

	log.Panic("unreachable statement")
	return nil, nil, nil
}

func getMin(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
