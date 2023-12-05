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
	seeds, conversionGroups := getSampleSeedValues(), getSampleConversionGroups()
	//seeds, conversionGroups := readGardenData()

	sortConversionRules(conversionGroups)

	solveLowestLocationNumber(append([]int{}, seeds...), conversionGroups)
	solveLowestLocationNumberWithSeedRanges(convertToRanges(seeds), conversionGroups)
}

func solveLowestLocationNumber(seeds []int, conversionGroups [][]ConversionRule) {
	values := seeds // does not prevent mutating seeds
	for _, conversionRules := range conversionGroups {
		for index, value := range seeds {
			// TODO: rewrite to check for overlap, then calculate, else keep same value
			// - this will make code reuse easier with part 2
			values[index] = calculate(value, conversionRules)
		}
	}

	minValue := values[0]
	for _, value := range values {
		minValue = getMin(minValue, value)
	}

	log.Println("lowest location number:", minValue)
}

func solveLowestLocationNumberWithSeedRanges(seedRanges [][]int, conversionGroups [][]ConversionRule) {
	values := seedRanges
	for _, conversionRules := range conversionGroups {
		valueBuffer := make([][]int, 0)

		valueRangeLoop:
		for _, valueRange := range values {

			for _, rule := range conversionRules {
				overlappingRange, nonOverlappingLowValueRange, nonOverlappingHighValueRange := partitionOverlappingRange(valueRange, rule)

				if nonOverlappingHighValueRange != nil {
					valueBuffer = append(valueBuffer, nonOverlappingHighValueRange)
				}

				if overlappingRange != nil {
					valueBuffer = append(valueBuffer, calculateWithRange(overlappingRange, rule))
				}

				if nonOverlappingLowValueRange == nil {
					break valueRangeLoop
				}

				valueRange = nonOverlappingLowValueRange
			}

			valueBuffer = append(valueBuffer, valueRange)
		}

		values = valueBuffer
	}

	minValue := values[0][0]
	for _, valueRange := range values {
		minValue = getMin(minValue, valueRange[0])
	}

	log.Println("lowest location number using ranges:", minValue)
}

func calculate(value int, rules []ConversionRule) int {
	for _, rule := range rules {
		if rule.start <= value && value <= rule.end {
			return value + rule.offset
		}
	}

	return value
}

// assume rule is applicable
func calculateWithRange(valueRange []int, rule ConversionRule) []int {
	return []int{valueRange[0] + rule.offset, valueRange[1] + rule.offset}
}

func sortConversionRules(conversionGroups [][]ConversionRule) {
	for _, group := range conversionGroups {
		sort.Slice(group, func(i, j int) bool {
			return group[i].start > group[j].start
		})
	}
}

func convertToRanges(seeds []int) [][]int {
	seedRanges := make([][]int, 0)

	for i := 0; i < len(seeds)/2; i++ {
		startIndex := i * 2
		endIndex := startIndex + 1

		startValue := seeds[startIndex]
		endValue := startValue + seeds[endIndex] - 1
		seedRanges = append(seedRanges, []int{startValue, endValue})
	}

	return seedRanges
}

// returns: overlapping range, prior non-overlapping range, after non-overlapping range
func partitionOverlappingRange(valueRange []int, rule ConversionRule) ([]int, []int, []int) {
	switch {
	case valueRange[1] < rule.start:
		return nil, valueRange, nil
	case valueRange[0] < rule.start && valueRange[1] <= rule.end:
		return []int{rule.start, valueRange[1]}, []int{valueRange[0], rule.start - 1}, nil
	case valueRange[0] < rule.start && rule.end < valueRange[1]:
		return []int{rule.start, valueRange[1]}, []int{valueRange[0], rule.start - 1}, []int{rule.end + 1, valueRange[1]}
	case rule.start <= valueRange[0] && valueRange[1] <= rule.end:
		return valueRange, nil, nil
	case rule.start <= valueRange[0] && rule.end < valueRange[1]:
		return []int{valueRange[0], rule.end}, nil, []int{rule.end + 1, valueRange[1]}
	case rule.end < valueRange[0]:
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
