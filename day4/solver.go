package day4

import (
	"log"
	"math"
)

type Solver struct{}

type ScratchCard struct {
	id             int
	winningNumbers []int
	cardNumbers    []int
}

func (s Solver) Solve() {
	// scratchCards := getSampleScratchCards()
	scratchCards := readScratchCards()

	solveScratchCardPointSum(scratchCards)
	solveMultiplyingScratchCardSum(scratchCards)
}

/*
Given scratch cards []ScratchCard, which has winning numbers and card numbers
Sum the score of each scratch card

Note: the scoring is based on amount of overlapping winning and card numbers
- the scoring sequence: 0, 1, 2, 4, 8, 16, ...
*/
func solveScratchCardPointSum(scratchCards []ScratchCard) {
	totalScore := 0

	for _, currScratchCard := range scratchCards {
		scoringNumberCount := countOverlappingValues(currScratchCard.cardNumbers, currScratchCard.winningNumbers)

		totalScore += calculateScore(scoringNumberCount)
	}

	log.Println("scratch cards total score:", totalScore)
}

/*
Given scratch cards []ScratchCard, which has winning numbers and card numbers
Sum the total cards in possession after processing the cards in order

Note: overlapping winning numbers grants additional scratch cards
- one overlapping number grants one card of the next scratch card
- two overlapping numbers grants one card each of the next two scratch cards
*/
func solveMultiplyingScratchCardSum(scratchCards []ScratchCard) {
	scratchCardCount := make(map[int]int)
	for _, currScratchCard := range scratchCards {
		scratchCardCount[currScratchCard.id]++
	}

	for _, currScratchCard := range scratchCards {
		scoringNumberCount := countOverlappingValues(currScratchCard.cardNumbers, currScratchCard.winningNumbers)

		for offset := 1; offset <= scoringNumberCount; offset++ {
			_, isFound := scratchCardCount[currScratchCard.id+offset]
			if !isFound {
				break
			}

			scratchCardCount[currScratchCard.id+offset] += scratchCardCount[currScratchCard.id]
		}
	}

	totalCount := 0
	for _, currScratchCardCount := range scratchCardCount {
		totalCount += currScratchCardCount
	}

	log.Println("multiplying scratch cards total count:", totalCount)
}

func calculateScore(matchingNumbers int) int {
	switch matchingNumbers {
	case 0:
		return 0
	case 1:
		return 1
	default:
		value := math.Pow(2, float64(matchingNumbers-1))
		return int(value)
	}
}
