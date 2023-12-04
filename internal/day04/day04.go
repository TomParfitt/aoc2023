package day04

import (
	"strings"
	"sync"

	"github.com/TomParfitt/aoc2023/internal/utils/process"
	"github.com/TomParfitt/aoc2023/internal/utils/sum"
)

func sumGetPoints(input string) int {
	return process.SumResults(input, getPoints)
}

func getPoints(input string, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done()

	line := strings.Replace(input, "  ", " ", -1)
	parts := strings.FieldsFunc(line, split)

	winningNumbers := strings.Split(strings.Trim(parts[1], " "), " ")
	handNumbers := strings.Split(strings.Trim(parts[2], " "), " ")

	sum := 0
	for _, winningNumber := range winningNumbers {
		for _, handNumber := range handNumbers {
			if winningNumber == handNumber {
				sum = max(1, sum*2)
				break
			}
		}
	}

	resultChan <- sum
}

func getNumberOfCards(input string) int {
	lines := strings.Split(input, "\n")

	results := make([]int, len(lines))
	for i := range lines {
		results[i] = 1
	}

	for i, line := range lines {
		count := countWinningLines(line)

		for c := i + 1; c <= count+i; c++ {
			results[c] += results[i]
		}
	}

	return sum.IntS(results)
}

func countWinningLines(input string) int {
	line := strings.Replace(input, "  ", " ", -1)
	parts := strings.FieldsFunc(line, split)

	winningNumbers := strings.Split(strings.Trim(parts[1], " "), " ")
	handNumbers := strings.Split(strings.Trim(parts[2], " "), " ")

	count := 0
	for _, winningNumber := range winningNumbers {
		for _, handNumber := range handNumbers {
			if winningNumber == handNumber {
				count++
				break
			}
		}
	}
	return count
}

func split(r rune) bool {
	return r == ':' || r == '|'
}
