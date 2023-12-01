package internal

import (
	"log/slog"
	"regexp"
	"strconv"
	"strings"
)

func getCalibrationValue(input string) int {

	r, _ := regexp.Compile("(one|two|three|four|five|six|seven|eight|nine|[0-9]){1}")
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {

		found := r.FindAllString(line, -1)

		first := convertToIntStr(found[0])
		last := convertToIntStr(found[len(found)-1])

		number, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}
		slog.Debug("getCalibrationValue", "number", number)

		sum += number
	}

	slog.Info("getCalibrationValue", "sum", sum)
	return sum
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
