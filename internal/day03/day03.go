package day03

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

type part struct {
	row  int
	colS int
	colE int
}

type symbol struct {
	row int
	col int
}

func sumPartNumbers(grid string) int {

	rows := strings.Split(grid, "\n")

	var symbols []symbol
	for y, row := range rows {
		for x, char := range row {
			if !unicode.IsDigit(char) && char != '.' {
				symbols = append(symbols, symbol{y, x})
			}
		}
	}

	var partNumbers []int
	for y, row := range rows {
		partNum := ""
		for x, char := range row {
			if unicode.IsDigit(char) {
				partNum += string(char)
			}

			if (!unicode.IsDigit(char) || len(row)-1 == x) && partNum != "" {

				idx := x
				if len(row)-1 == x && unicode.IsDigit(char) {
					idx = x + 1
				}

				for _, symbol := range symbols {
					if math.Abs(float64(symbol.row-y)) <= 1 && idx-len(partNum)-1 <= symbol.col && symbol.col <= idx {
						partNumber, _ := strconv.Atoi(partNum)
						partNumbers = append(partNumbers, partNumber)
						break
					}
				}
				partNum = ""
			}
		}
	}

	sum := 0
	for _, partNumber := range partNumbers {
		sum += partNumber
	}

	return sum
}

type cog struct {
	row         int
	col         int
	partNumbers []int
}

func (c *cog) addPartNumber(partNumber int) {
	c.partNumbers = append(c.partNumbers, partNumber)
}

func sumPartNumbers2(grid string) int {

	rows := strings.Split(grid, "\n")

	var cogs []cog
	for y, row := range rows {
		for x, char := range row {
			if char == '*' {
				cogs = append(cogs, cog{row: y, col: x})
			}
		}
	}

	for y, row := range rows {
		partNum := ""
		for x, char := range row {
			if unicode.IsDigit(char) {
				partNum += string(char)
			}

			if (!unicode.IsDigit(char) || len(row)-1 == x) && partNum != "" {

				idx := x
				if len(row)-1 == x && unicode.IsDigit(char) {
					idx = x + 1
				}

				for i, cog := range cogs {
					if math.Abs(float64(cog.row-y)) <= 1 && idx-len(partNum)-1 <= cog.col && cog.col <= idx {

						partNumber, _ := strconv.Atoi(partNum)
						cog.addPartNumber(partNumber)
						cogs[i] = cog
						break
					}
				}
				partNum = ""
			}
		}
	}

	sum := 0
	for _, cog := range cogs {
		if len(cog.partNumbers) == 2 {
			sum += cog.partNumbers[0] * cog.partNumbers[1]
		}
	}

	return sum
}
