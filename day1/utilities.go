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

func filterDigitChars(line string) string {
	stringBuilder := strings.Builder{}

	for _, runeValue := range line {
		if unicode.IsDigit(runeValue) {
			stringBuilder.WriteString(string(runeValue))
		}
	}

	return stringBuilder.String()
}

func filterDigitCharsWithTokens(line string, tokens []string) string {
	digitChars := strings.Builder{}

	buffer := strings.Builder{}
	for _, runeValue := range line {
		charValue := string(runeValue)
		buffer.WriteString(charValue)

		bufferStr := buffer.String()
	tokenLoop:
		for _, token := range tokens {
			if strings.Contains(bufferStr, token) {
				switch token {
				case "0", "zero":
					digitChars.WriteString("0")
					break
				case "1", "one":
					digitChars.WriteString("1")
					break
				case "2", "two":
					digitChars.WriteString("2")
					break
				case "3", "three":
					digitChars.WriteString("3")
					break
				case "4", "four":
					digitChars.WriteString("4")
					break
				case "5", "five":
					digitChars.WriteString("5")
					break
				case "6", "six":
					digitChars.WriteString("6")
					break
				case "7", "seven":
					digitChars.WriteString("7")
					break
				case "8", "eight":
					digitChars.WriteString("8")
					break
				case "9", "nine":
					digitChars.WriteString("9")
					break
				default:
					log.Panic("unable to handle unexpected token:", token)
				}

				// handles the edge case of overlapping words (eg. nineight)
				tailRune := []rune(buffer.String())[buffer.Len()-1]
				buffer.Reset()

				if !unicode.IsDigit(tailRune) {
					buffer.WriteRune(tailRune)
				}

				break tokenLoop
			}
		}
	}

	return digitChars.String()
}
