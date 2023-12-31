package day02

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"sync"

	"github.com/TomParfitt/aoc2023/internal/utils/process"
)

func getSumOfValidIds(input string) int {
	return process.SumResults(input, getValidId)
}

func getValidId(input string, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done()

	maxs := map[string]int{"red": 12, "green": 13, "blue": 14}

	parts := strings.Split(input, ":")
	idStr := strings.Split(parts[0], " ")[1]

	rounds := strings.Split(parts[1], ";")
	for _, round := range rounds {

		cubes := strings.Split(round, ",")
		for _, cube := range cubes {

			countAndColour := strings.Split(strings.Trim(cube, " "), " ")
			count, _ := strconv.Atoi(countAndColour[0])
			colour := countAndColour[1]

			if count > maxs[colour] {
				slog.Info("process", "msg", fmt.Sprintf("Game %v not possible: %v %v", idStr, colour, count))
				return
			}
		}
	}

	idInt, _ := strconv.Atoi(idStr)
	resultChan <- idInt
}

func getSumOfGamePower(input string) int {
	return process.SumResults(input, getGamePower)
}

func getGamePower(input string, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done()

	mins := map[string]int{"red": 0, "blue": 0, "green": 0}

	gameAndRounds := strings.Split(input, ":")
	idStr := strings.Split(gameAndRounds[0], " ")[1]

	rounds := strings.Split(gameAndRounds[1], ";")
	for _, round := range rounds {

		cubes := strings.Split(round, ",")
		for _, cube := range cubes {

			countAndColour := strings.Split(strings.Trim(cube, " "), " ")
			count, _ := strconv.Atoi(countAndColour[0])
			colour := countAndColour[1]

			if min := mins[colour]; count > min {
				mins[colour] = count
			}
		}
	}

	product := 1
	for _, value := range mins {
		product *= value
	}
	slog.Info("process", "msg", fmt.Sprintf("Game %v power: %v", idStr, product))

	resultChan <- product
}
