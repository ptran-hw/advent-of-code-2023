package day1

import (
	"fmt"
	"log"
	"strconv"
)

type Solver struct{}

func (s Solver) Solve() {
	//lines := getSampleCalibrationDocument()
	lines := readCalibrationDocument()

	solveCalibrationValueSum(lines)
}

func solveCalibrationValueSum(lines []string) {
	sum := 0

	for _, line := range lines {
		value := calculateCalibrationValue(line)
		sum += value
	}

	log.Println("calibration value sum:", sum)
}

func calculateCalibrationValue(line string) int {
	digits := filterDigitChars(line)

	calibrationValue := fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1])

	value, err := strconv.Atoi(calibrationValue)
	if err != nil {
		log.Panicf("unable to parse calculation value from: %s, %v", line, err)
	}

	return value
}
