package day3

import (
	"log"
	"strconv"
)

type Solver struct{}

type SchematicPosition struct {
	value       string
	xCoordinate int
	yCoordinate int
}

func (s Solver) Solve() {
	// positions := getSampleSchematicPositions()
	positions := readSchematicPositions()

	solvePartNumbersSum(positions)
	solveGearRatioSum(positions)
}

/*
Given positions []SchematicPosition, consisting of numeric/symbol positions
Determine the numeric positions that neighbours symbol positions
Return the sum of numeric position value
*/
func solvePartNumbersSum(positions []SchematicPosition) {
	numericPositions, symbolPositions := filterNumericPositions(positions)

	partPositions := make([]SchematicPosition, 0)
	for _, currNumericPosition := range numericPositions {
		for _, currSymbolPosition := range symbolPositions {
			if isNeighbouringPositions(currNumericPosition, currSymbolPosition) {
				partPositions = append(partPositions, currNumericPosition)
				break
			}
		}
	}

	total := 0
	for _, currPosition := range partPositions {
		value, err := strconv.Atoi(currPosition.value)
		if err != nil {
			log.Panicf("unable to parse part number value: %s", currPosition.value)
		}
		total += value
	}

	log.Println("part numbers sum:", total)
}

/*
Given positions []SchematicPosition, consisting of numeric/symbol positions
Find numeric positions where only 2 neighbours a single symbol position
Return the sum of gear ratio (product of neighbour numeric positions)
*/
func solveGearRatioSum(positions []SchematicPosition) {
	numericPositions, symbolPositions := filterNumericPositions(positions)

	gearRatioPositions := make([][]SchematicPosition, 0)
	for _, currSymbolPosition := range symbolPositions {
		neighbourNumericPositions := make([]SchematicPosition, 0)

		for _, currNumericPosition := range numericPositions {
			if isNeighbouringPositions(currNumericPosition, currSymbolPosition) {
				neighbourNumericPositions = append(neighbourNumericPositions, currNumericPosition)
			}
		}

		if len(neighbourNumericPositions) == 2 {
			gearRatioPositions = append(gearRatioPositions, neighbourNumericPositions)
		}
	}

	total := 0
	for _, neighbourPositions := range gearRatioPositions {
		valueA, err := strconv.Atoi(neighbourPositions[0].value)
		if err != nil {
			log.Panicf("unable to parse value for position: %v", neighbourPositions[0])
		}

		valueB, err := strconv.Atoi(neighbourPositions[1].value)
		if err != nil {
			log.Panicf("unable to parse value for position: %v", neighbourPositions[1])
		}

		gearRatio := valueA * valueB
		total += gearRatio
	}

	log.Println("gear ratio sum:", total)
}

func filterNumericPositions(positions []SchematicPosition) ([]SchematicPosition, []SchematicPosition) {
	numericPositions := make([]SchematicPosition, 0)
	otherPositions := make([]SchematicPosition, 0)

	for _, currPosition := range positions {
		if isNumeric(currPosition.value) {
			numericPositions = append(numericPositions, currPosition)
		} else {
			otherPositions = append(otherPositions, currPosition)
		}
	}

	return numericPositions, otherPositions
}

func isNeighbouringPositions(numericPosition, symbolPosition SchematicPosition) bool {
	length := len(numericPosition.value)
	for xIncrement := 0; xIncrement < length; xIncrement++ {
		if isNeighbour(numericPosition.xCoordinate+xIncrement, numericPosition.yCoordinate, symbolPosition.xCoordinate, symbolPosition.yCoordinate) {
			return true
		}
	}

	return false
}
