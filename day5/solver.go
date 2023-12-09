package day5

import (
	"log"
	"sort"
)

type Solver struct{}

type ConversionRule struct {
	start  int
	end    int
	offset int
}

func (s Solver) Solve() {
	seeds, conversionGroups := getSampleSeedValues(), getSampleConversionGroups()
	// seeds, conversionGroups := readSeedData(), readGardenData()
	reverseSortConversionRules(conversionGroups)

	solveMinLocation(seeds, conversionGroups)
	solveMinLocationWithRanges(seeds, conversionGroups)
}

func solveMinLocation(seeds []int, conversionGroups [][]ConversionRule) {
	seedRanges := make([][]int, 0)

	for _, value := range seeds {
		seedRanges = append(seedRanges, []int{value, value})
	}

	minLocation := calculateMinLocation(seedRanges, conversionGroups)
	log.Println("min location:", minLocation)
}

func solveMinLocationWithRanges(seeds []int, conversionGroups [][]ConversionRule) {
	seedRanges := mapToRangesUsingOffset(seeds)

	minLocation := calculateMinLocation(seedRanges, conversionGroups)
	log.Println("[with ranges] min location:", minLocation)
}

func calculateMinLocation(seedRanges [][]int, conversionGroups [][]ConversionRule) int {
	values := seedRanges
	for _, conversionRules := range conversionGroups {
		valueBuffer := make([][]int, 0)

		for _, valueRange := range values {
			for _, rule := range conversionRules {
				overlappingRange, nonOverlappingLowValueRange, nonOverlappingHighValueRange := splitOverlappingRange(valueRange, []int{rule.start, rule.end})

				if nonOverlappingHighValueRange != nil {
					valueBuffer = append(valueBuffer, nonOverlappingHighValueRange)
				}

				if overlappingRange != nil {
					valueBuffer = append(valueBuffer, applyConversionRule(overlappingRange, rule))
				}

				valueRange = nonOverlappingLowValueRange
			}

			if valueRange != nil {
				valueBuffer = append(valueBuffer, valueRange)
			}
		}

		values = valueBuffer
	}

	minValue := values[0][0]
	for _, valueRange := range values {
		minValue = getMin(minValue, valueRange[0])
	}

	return minValue
}

func reverseSortConversionRules(conversionGroups [][]ConversionRule) {
	for _, group := range conversionGroups {
		sort.Slice(group, func(i, j int) bool {
			return group[i].start > group[j].start
		})
	}
}

func mapToRangesUsingOffset(seeds []int) [][]int {
	if len(seeds) == 0 {
		log.Panic("invalid seeds value")
	}

	seedRanges := make([][]int, 0)
	index := 0

	for index < len(seeds)-1 {
		rangeStart := seeds[index]
		offset := seeds[index+1]
		rangeEnd := rangeStart + offset - 1

		seedRanges = append(seedRanges, []int{rangeStart, rangeEnd})
		index += 2
	}

	return seedRanges
}

func applyConversionRule(valueRange []int, rule ConversionRule) []int {
	return []int{valueRange[0] + rule.offset, valueRange[1] + rule.offset}
}
