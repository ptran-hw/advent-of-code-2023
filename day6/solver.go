package day6

import "log"

type Solver struct {}

type Race struct {
	duration int
	recordDistance int
}

func (s Solver) Solve() {
	partOneRaces, partTwoRaces := getPartOneSampleRaces(), getPartTwoSampleRaces()
	// partOneRaces, partTwoRaces := getPartOneRealRaceData(), getPartTwoRealRaceData()

	solveTotalWinningRaceApproachScore(partOneRaces)
	solveTotalWinningRaceApproachScore(partTwoRaces)
}

/*
   Given races []Race, which tracks duration and record distance
   Calculate the total score (product of each race possible winning ways)
*/
func solveTotalWinningRaceApproachScore(races []Race) {
	totalScore := 1
	for _, currRace := range races {
		totalScore *= getWinningApproachCount(currRace)
	}

	log.Println("total winning race approach score:", totalScore)
}

func getWinningApproachCount(race Race) int {
	count := 0

	for chargingTime := 1; chargingTime < race.duration; chargingTime++ {
		distanceCovered := calculateDistance(chargingTime, race.duration - chargingTime)
		if distanceCovered > race.recordDistance {
			count++
		}
	}

	return count
}

func calculateDistance(speed, duration int) int {
	return speed * duration
}

