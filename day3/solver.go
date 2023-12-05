package day3

import (
	"log"
	"strconv"
)

type Solver struct{}

type SchematicPosition struct{
	value string
	xCoordinate int
	yCoordinate int
}

func (s Solver) Solve() {
	//positions := getSampleSchematicPositions()
	positions := readSchematicFile()

	solvePartNumbersSum(positions)
	solveGearRatioSum(positions)
}

func solvePartNumbersSum(positions []SchematicPosition) {
	numericPositions, symbolPositions := filterNumericPositions(positions)

	partNumbers := make([]int, 0)
	for _, currNumericPosition := range numericPositions {
		isPartNumber := false

		checkPartNumberLoop:
		for _, currSymbolPosition := range symbolPositions {
			if isNeighbouringPositions(currNumericPosition, currSymbolPosition) {
				isPartNumber = true
				break checkPartNumberLoop
			}
		}

		if isPartNumber {
			value, err := strconv.Atoi(currNumericPosition.value)
			if err != nil {
				log.Panicf("unable to parse part number value: %s", currNumericPosition.value)
			}

			partNumbers = append(partNumbers, value)
		}
	}

	total := 0
	for _, partNumber := range partNumbers {
		total += partNumber
	}

	log.Println("part numbers sum:", total)
}

func solveGearRatioSum(positions []SchematicPosition) {
	numericPositions, symbolPositions := filterNumericPositions(positions)

	total := 0
	for _, currSymbolPosition := range symbolPositions {
		neighbourNumericPositions := make([]SchematicPosition, 0)

		for _, currNumericPosition := range numericPositions {
			if isNeighbouringPositions(currNumericPosition, currSymbolPosition) {
				neighbourNumericPositions = append(neighbourNumericPositions, currNumericPosition)
			}
		}

		if len(neighbourNumericPositions) == 2 {
			valueA, err := strconv.Atoi(neighbourNumericPositions[0].value)
			if err != nil {
				log.Panicf("unable to parse value for position: %v", neighbourNumericPositions[0])
			}

			valueB, err := strconv.Atoi(neighbourNumericPositions[1].value)
			if err != nil {
				log.Panicf("unable to parse value for position: %v", neighbourNumericPositions[1])
			}

			gearRatio := valueA * valueB
			total += gearRatio
		}
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

func isNumeric(value string) bool {
	_, err := strconv.Atoi(value)

	return err == nil
}

func isNeighbouringPositions(numericPosition, symbolPosition SchematicPosition) bool {
	length := len(numericPosition.value)
	for xIncrement := 0; xIncrement < length; xIncrement++ {
		if isNeighbour(numericPosition.xCoordinate + xIncrement, numericPosition.yCoordinate, symbolPosition.xCoordinate, symbolPosition.yCoordinate) {
			return true
		}
	}

	return false
}

func isNeighbour(x1, y1, x2, y2 int) bool {
	return (x1-1 == x2 && y1-1 == y2) ||
		(x1-1 == x2 && y1 == y2) ||
		(x1-1 == x2 && y1+1 == y2) ||
		(x1 == x2 && y1-1 == y2) ||
		(x1 == x2 && y1+1 == y2) ||
		(x1+1 == x2 && y1-1 == y2) ||
		(x1+1 == x2 && y1 == y2) ||
		(x1+1 == x2 && y1+1 == y2)

	//case x1 - 1 == x2 && y1 - 1== y2:
	//case x1 - 1 == x2 && y1 == y2:
	//case x1 - 1 == x2 && y1 + 1== y2:
	//case x1 == x2 && y1 - 1== y2:
	//case x1 == x2 && y1 == y2:
	//case x1 == x2 && y1 + 1== y2:
	//case x1 + 1 == x2 && y1 - 1== y2:
	//case x1 + 1 == x2 && y1 == y2:
	//case x1 + 1 == x2 && y1 + 1== y2:
}
