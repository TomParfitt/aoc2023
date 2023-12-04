package day01

import (
	"regexp"
	"strconv"
	"sync"

	"github.com/TomParfitt/aoc2023/internal/utils/process"
)

var r = regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine|[0-9]){1}")

func getSumOfCalibrationValue(input string) int {
	return process.SumResults(input, getCalibrationValue)
}

func getCalibrationValue(input string, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done()

	found := r.FindAllString(input, -1)

	first := convertToIntStr(found[0])
	last := convertToIntStr(found[len(found)-1])

	number, err := strconv.Atoi(first + last)
	if err != nil {
		panic(err)
	}

	resultChan <- number
}

func convertToIntStr(str string) string {
	switch str {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return str
	}
}
