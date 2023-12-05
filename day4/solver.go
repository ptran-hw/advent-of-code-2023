package day4

import (
	"log"
	"math"
)

type Solver struct {}

type ScratchCard struct {
	id int
	winningNumbers []int
	cardNumbers []int
}

func (s Solver) Solve() {
	//scratchCards := getSampleScratchCards()
	scratchCards := readScratchCards()

	solveScratchCardPointSum(scratchCards)
	solveMultiplyingScratchCardPointSum(scratchCards)
}

func solveScratchCardPointSum(scratchCards []ScratchCard) {
	totalScore := 0

	for _, currScratchCard := range scratchCards {
		scoringNumberCount := countOverlappingValues(currScratchCard.cardNumbers, currScratchCard.winningNumbers)

		totalScore += calculateScore(scoringNumberCount)
	}

	log.Println("scratch cards total score:", totalScore)
}

func solveMultiplyingScratchCardPointSum(scratchCards []ScratchCard) {
	scratchCardCount := make(map[int]int)
	for _, currScratchCard := range scratchCards {
		scratchCardCount[currScratchCard.id]++
	}

	for _, currScratchCard := range scratchCards {
		scoringNumberCount := countOverlappingValues(currScratchCard.cardNumbers, currScratchCard.winningNumbers)

		for offset := 1; offset <= scoringNumberCount; offset++ {
			_, isFound := scratchCardCount[currScratchCard.id + offset]
			if !isFound {
				break
			}

			scratchCardCount[currScratchCard.id + offset] += scratchCardCount[currScratchCard.id]
		}
	}

	totalCount := 0
	for _, value := range scratchCardCount {
		totalCount += value
	}

	log.Println("multiplying scratch cards total count:", totalCount)
}

// assuming that a winning number can only show up at most once in card numbers
func countOverlappingValues(arrA, arrB []int) int {
	overlappingCount := 0

	for _, valueA := range arrA {
		for _, valueB := range arrB {
			if valueA == valueB {
				overlappingCount++
			}
		}
	}

	return overlappingCount
}

func calculateScore(matchingNumbers int) int {
	switch matchingNumbers {
	case 0:
		return 0
	case 1:
		return 1
	default:
		value := math.Pow(2, float64(matchingNumbers - 1))
		return int(value)
	}
}
