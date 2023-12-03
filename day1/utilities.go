package day1

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const calibrationDocumentFile = "./day1/calibration-document.txt"

var tokenToDigitMap = map[string]string{
	"1":     "1",
	"one":   "1",
	"2":     "2",
	"two":   "2",
	"3":     "3",
	"three": "3",
	"4":     "4",
	"four":  "4",
	"5":     "5",
	"five":  "5",
	"6":     "6",
	"six":   "6",
	"7":     "7",
	"seven": "7",
	"8":     "8",
	"eight": "8",
	"9":     "9",
	"nine":  "9",
	"0":     "0",
	"zero":  "0",
}

// consider extracting the reader into a shared file,
// then we can define function to map []string to easier type to work with
func readCalibrationDocument() []string {
	file, err := os.Open(calibrationDocumentFile)
	if err != nil {
		log.Panicf("unable to read file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func getSampleCalibrationDocument() []string {
	return []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
}

func getSampleCalibrationDocumentWithWords() []string {
	return []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
}

func findFirst(str string, tokens []string) string {
	minIndex := len(str)
	result := ""

	for _, currToken := range tokens {
		currTokenIndex := strings.Index(str, currToken)
		if currTokenIndex != -1 && currTokenIndex < minIndex {
			minIndex = currTokenIndex
			result = currToken
		}
	}

	return result
}

func findLast(str string, tokens []string) string {
	maxIndex := -1
	result := ""

	for _, currToken := range tokens {
		currTokenIndex := strings.LastIndex(str, currToken)
		if currTokenIndex > maxIndex {
			maxIndex = currTokenIndex
			result = currToken
		}
	}

	return result
}

func convertToDigit(token string) string {
	digit, isFound := tokenToDigitMap[token]
	if !isFound {
		log.Panic("unable to handle unexpected token:", token)
	}

	return digit
}
