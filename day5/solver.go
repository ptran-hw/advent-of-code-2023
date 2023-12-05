package day5

import (
	"log"
	"sort"
)

type Solver struct {}

type ConversionRule struct {
	start int
	end int
	offset int
}

func (s Solver) Solve() {
	//seeds, conversionGroups := getSampleSeedValues(), getSampleConversionGroups()
	seeds, conversionGroups := readGardenData()

	sortConversionRules(conversionGroups)

	solveLowestLocationNumber(seeds, conversionGroups)
}

func solveLowestLocationNumber(seeds []int, conversionGroups [][]ConversionRule) {
	values := seeds
	for _, conversionRules := range conversionGroups {
		for index, value := range seeds {
			values[index] = calculate(value, conversionRules)
		}
	}

	minValue := values[0]
	for _, value := range values {
		minValue = getMin(minValue, value)
	}

	log.Println("lowest location number:", minValue)
}

func calculate(value int, rules []ConversionRule) int {
	for _, rule := range rules {
		if rule.start <= value && value <= rule.end {
			return value + rule.offset
		}
	}

	return value
}

func sortConversionRules(conversionGroups [][]ConversionRule) {
	for _, group := range conversionGroups {
		sort.Slice(group, func(i, j int) bool {
			return group[i].start > group[j].start
		})
	}
}

func getMin(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
