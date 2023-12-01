package internal

import (
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func getCalibrationValue(input string) int {
	r, _ := regexp.Compile("(one|two|three|four|five|six|seven|eight|nine|[0-9]){1}")
	lines := strings.Split(input, "\n")

	var wg sync.WaitGroup
	resultChan := make(chan int)

	for _, line := range lines {
		wg.Add(1)
		go process(line, r, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	sum := 0
	for r := range resultChan {
		sum += r
	}

	return sum
}

func process(input string, r *regexp.Regexp, wg *sync.WaitGroup, resultChan chan<- int) {
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
