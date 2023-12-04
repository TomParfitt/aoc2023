package internal

import (
	"fmt"
	"log/slog"
	"os"
	"testing"
)

func TestGetValidIdsSumExample(t *testing.T) {
	// Given
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	// When
	valiIdsSum := sumProcessResults(input, getValidId)

	// Then
	if valiIdsSum != 8 {
		t.Fatalf(`getValidIdsSum with example value %v != 8`, valiIdsSum)
	}
}

func TestGetValidIdsSum(t *testing.T) {
	// Given
	b, err := os.ReadFile("day2.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)

	// When
	valiIdsSum := sumProcessResults(str, getValidId)

	// Then
	slog.Info("getValidIdsSum", "valiIdsSum", valiIdsSum)
}

func TestGetSumOfMinCubesExample(t *testing.T) {
	// Given
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	// When
	sumOfMinCubes := sumProcessResults(input, getGamePower)

	// Then
	if sumOfMinCubes != 2286 {
		t.Fatalf(`getSumOfMinCubes with example value %v != 2286`, sumOfMinCubes)
	}
}

func TestGetSumOfMinCubes(t *testing.T) {
	// Given
	b, err := os.ReadFile("day2.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)

	// When
	sumOfMinCubes := sumProcessResults(str, getGamePower)

	// Then
	slog.Info("getSumOfMinCubes", "sumOfMinCubes", sumOfMinCubes)
}
