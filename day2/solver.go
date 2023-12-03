package day2

import (
	"log"
)

const maxBlueCubes = 14
const maxRedCubes = 12
const maxGreenCubes = 13

type Solver struct{}

type Game struct {
	id     int
	rounds []Round
}

type Round struct {
	blueCubes  int
	redCubes   int
	greenCubes int
}

func (s Solver) Solve() {
	games := getSampleGames()
	//games := readGames()

	solveValidGamesIdSum(games)
	solveMinimumCubesPowerValue(games)
}

func solveValidGamesIdSum(games []Game) {
	validGames := filterValidGames(games)

	total := 0
	for _, currGame := range validGames {
		total += currGame.id
	}

	log.Println("total sum of valid game ids:", total)
}

func solveMinimumCubesPowerValue(games []Game) {
	gameMinimumRequiredCubes := calculateMinimumRequiredCubes(games)

	totalPower := 0
	for _, minimumCubeColours := range gameMinimumRequiredCubes {
		totalPower += minimumCubeColours[0] * minimumCubeColours[1] * minimumCubeColours[2]
	}

	log.Println("total power level:", totalPower)
}

func filterValidGames(games []Game) []Game {
	validGames := make([]Game, 0)
	for _, currGame := range games {
		isValid := true

		for _, currRound := range currGame.rounds {
			if currRound.blueCubes > maxBlueCubes || currRound.redCubes > maxRedCubes || currRound.greenCubes > maxGreenCubes {
				isValid = false
				break
			}
		}

		if isValid {
			validGames = append(validGames, currGame)
		}
	}

	return validGames
}

func calculateMinimumRequiredCubes(games []Game) [][]int {
	gameRequiredColours := make([][]int, 0)
	for _, currGame := range games {
		minRequiredBlue := 0
		minRequiredRed := 0
		minRequiredGreen := 0

		for _, currRound := range currGame.rounds {
			minRequiredBlue = getMax(minRequiredBlue, currRound.blueCubes)
			minRequiredRed = getMax(minRequiredRed, currRound.redCubes)
			minRequiredGreen = getMax(minRequiredGreen, currRound.greenCubes)
		}

		gameRequiredColours = append(gameRequiredColours, []int{minRequiredBlue, minRequiredRed, minRequiredGreen})
	}

	return gameRequiredColours
}
