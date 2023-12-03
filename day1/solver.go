package day1

import (
	"fmt"
	"github.com/agrison/go-commons-lang/stringUtils"
	"log"
	"strconv"
)

var digitTokens = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var wordTokens = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "zero"}

type Solver struct{}

func (s Solver) Solve() {
	//linesA, linesB := getSampleCalibrationDocument(), getSampleCalibrationDocumentWithWords()
	linesA, linesB := readCalibrationDocument(), readCalibrationDocument()

	solveCalibrationValueSum(linesA)
	solveCalibrationValueSumIncludingWords(linesB)
}

/*
Given []string lines,
Combine the first digit (eg. 1, 2, 3, etc) and last digit to form calibration value
Sum the calibration value for each line

Note:
- when there is one digit in line, it will be used as both first and last digit
*/
func solveCalibrationValueSum(lines []string) {
	numberTokens := digitTokens

	sum := 0
	for _, line := range lines {
		value := calculateCalibrationValueWithTokens(line, numberTokens)
		sum += value
	}

	log.Println("using digits only, calibration value sum:", sum)
}

/*
Given []string lines,
Combine the first number (eg. 1, 2, one, two, etc) and last number to form calibration value
Sum the calibration value for each line

Note:
- numbers are 0..9 and zero..nine
- there may be overlapping words (eg. threeight)
- when there is one number in line, it will be used as both first and last number
*/
func solveCalibrationValueSumIncludingWords(lines []string) {
	numberTokens := append(digitTokens, wordTokens...)

	sum := 0
	for _, line := range lines {
		sum += calculateCalibrationValueWithTokens(line, numberTokens)
	}

	log.Printf("using digits and words, calibration value sum: %d", sum)
}

func calculateCalibrationValueWithTokens(line string, tokens []string) int {
	firstToken := findFirst(line, tokens)
	if stringUtils.IsEmpty(firstToken) {
		log.Panicf("unable to find first token %v in line: %s", tokens, line)
	}

	lastToken := findLast(line, tokens)
	if stringUtils.IsEmpty(lastToken) {
		log.Panicf("unable to find last token %v in line: %s", tokens, line)
	}

	calibrationValue := fmt.Sprintf("%s%s", convertToDigit(firstToken), convertToDigit(lastToken))

	value, err := strconv.Atoi(calibrationValue)
	if err != nil {
		log.Panicf("unable to parse calculation value from: %s, %v", line, err)
	}

	return value
}
