package day1

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

const calibrationDocumentFile = "./day1/calibration-document.txt"

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

func filterDigitChars(line string) string {
	stringBuilder := strings.Builder{}

	for _, runeValue := range line {
		if unicode.IsDigit(runeValue) {
			stringBuilder.WriteString(string(runeValue))
		}
	}

	return stringBuilder.String()
}
