package day3

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const engineSchematicFile = "./day3/engine-schematic.txt"
const emptySymbol = "."

func readSchematicPositions() []SchematicPosition {
	file, err := os.Open(engineSchematicFile)
	if err != nil {
		log.Panicf("unable to read file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	positions := make([]SchematicPosition, 0)

	yIndex := 0
	for scanner.Scan() {
		line := scanner.Text()

		positions = append(positions, parseSchematicPositions(line, yIndex)...)
		yIndex++
	}

	return positions
}

// use sliding window to avoid using a buffer and conditional nesting
func parseSchematicPositions(line string, yIndex int) []SchematicPosition {
	headIndex := 0

	positions := make([]SchematicPosition, 0)
	for headIndex < len(line) {
		headChar := string(line[headIndex])

		switch {
		case headChar == emptySymbol:
			headIndex++
			continue
		case !isNumeric(headChar): // isSymbol
			positions = append(positions, SchematicPosition{value: headChar, xCoordinate: headIndex, yCoordinate: yIndex})
			headIndex++
			break
		case isNumeric(headChar):
			tailIndex := headIndex
			for tailIndex < len(line) && isNumeric(string(line[tailIndex])) {
				tailIndex++
			}

			positions = append(positions, SchematicPosition{value: line[headIndex:tailIndex], xCoordinate: headIndex, yCoordinate: yIndex})
			headIndex = tailIndex
			break
		default:
			log.Panicf("unable to handle head character: %s", headChar)
		}
	}

	return positions
}

func isNumeric(value string) bool {
	_, err := strconv.Atoi(value)

	return err == nil
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
}

func getSampleSchematicPositions() []SchematicPosition {
	return []SchematicPosition{
		{value: "467", xCoordinate: 0, yCoordinate: 0},
		{value: "114", xCoordinate: 5, yCoordinate: 0},
		{value: "*", xCoordinate: 3, yCoordinate: 1},
		{value: "35", xCoordinate: 2, yCoordinate: 2},
		{value: "633", xCoordinate: 6, yCoordinate: 2},
		{value: "#", xCoordinate: 6, yCoordinate: 3},
		{value: "617", xCoordinate: 0, yCoordinate: 4},
		{value: "*", xCoordinate: 3, yCoordinate: 4},
		{value: "+", xCoordinate: 5, yCoordinate: 5},
		{value: "58", xCoordinate: 7, yCoordinate: 5},
		{value: "592", xCoordinate: 2, yCoordinate: 6},
		{value: "755", xCoordinate: 6, yCoordinate: 7},
		{value: "$", xCoordinate: 3, yCoordinate: 8},
		{value: "*", xCoordinate: 5, yCoordinate: 8},
		{value: "664", xCoordinate: 1, yCoordinate: 9},
		{value: "598", xCoordinate: 5, yCoordinate: 9},
	}
}
