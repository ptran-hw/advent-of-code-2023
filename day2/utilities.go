package day2

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const cubeGameLogFile = "./day2/cube-game.log"
const blueCubeKey = "blue"
const redCubeKey = "red"
const greenCubeKey = "green"

var gamePattern = regexp.MustCompile("Game (\\d+): (.*)")

func readGames() []Game {
	file, err := os.Open(cubeGameLogFile)
	if err != nil {
		log.Panicf("unable to read file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	games := make([]Game, 0)
	for scanner.Scan() {
		line := scanner.Text()

		matches := gamePattern.FindStringSubmatch(line)
		if len(matches) != 3 {
			log.Panicf("invalid format line: %s", line)
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Panicf("unable to parse game id: %s", matches[1])
		}

		currGame := Game{
			id: id,
			rounds: parseRounds(matches[2]),
		}
		games = append(games, currGame)
	}

	return games
}

func parseRounds(str string) []Round {
	rounds := make([]Round, 0)

	rawRounds := strings.Split(str, ";")
	for _, currRawRound := range rawRounds {
		currRawRound = strings.TrimSpace(currRawRound)
		rawCubes := strings.Split(currRawRound, ",")

		currRound := Round{}
		for _, currRawCube := range rawCubes {
			currRawCube = strings.TrimSpace(currRawCube)

			cubeInfo := strings.Split(currRawCube, " ")
			if len(cubeInfo) != 2 {
				log.Panicf("invalid format round: %s", currRawRound)
			}

			count, err := strconv.Atoi(cubeInfo[0])
			if err != nil {
				log.Panicf("unable to parse cube count: %v", err)
			}

			switch cubeInfo[1] {
			case blueCubeKey:
				currRound.blueCubes = count
				break
			case redCubeKey:
				currRound.redCubes = count
				break
			case greenCubeKey:
				currRound.greenCubes = count
				break
			default:
				log.Panicf("unable to handle cube key: %s", cubeInfo[1])
			}
		}

		rounds = append(rounds, currRound)
	}

	return rounds
}

func getSampleGames() []Game {
	return []Game{
		{
			id: 1,
			rounds: []Round{
				{ blueCubes: 3, redCubes: 4 },
				{ blueCubes: 6, redCubes: 1, greenCubes: 2 },
				{ greenCubes: 2 },
			},
		},
		{
			id: 2,
			rounds: []Round{
				{ blueCubes: 1, greenCubes: 2 },
				{ blueCubes: 4, redCubes: 1, greenCubes: 3 },
				{ blueCubes: 1, greenCubes: 1 },
			},
		},
		{
			id: 3,
			rounds: []Round{
				{ blueCubes: 6, redCubes: 20, greenCubes: 8 },
				{ blueCubes: 5, redCubes: 4, greenCubes: 13 },
				{ redCubes: 1, greenCubes: 5 },
			},
		},
		{
			id: 4,
			rounds: []Round{
				{ blueCubes: 6, redCubes: 1, greenCubes: 3 },
				{ redCubes: 6, greenCubes: 3 },
				{ blueCubes: 15, redCubes: 14, greenCubes: 3 },
			},
		},
		{
			id: 5,
			rounds: []Round{
				{ blueCubes: 1, redCubes: 6, greenCubes: 3 },
				{ blueCubes: 2, redCubes: 1, greenCubes: 2 },
			},
		},
	}
}


func getMax(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
