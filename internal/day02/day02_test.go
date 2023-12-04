package day02

import (
	"testing"

	"github.com/TomParfitt/aoc2023/internal/utils/file"
	"github.com/stretchr/testify/assert"
)

func TestGetValidIdsSumExample(t *testing.T) {
	// Given
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	// When
	valiIdsSum := getSumOfValidIds(input)

	// Then
	assert.Equal(t, 8, valiIdsSum)
}

func TestGetValidIdsSum(t *testing.T) {
	// Given
	input := file.MustReadToString("day02.txt")

	// When
	valiIdsSum := getSumOfValidIds(input)

	// Then
	assert.Equal(t, 1853, valiIdsSum)
}

func TestGetSumOfMinCubesExample(t *testing.T) {
	// Given
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	// When
	sumOfMinCubes := getSumOfGamePower(input)

	// Then
	assert.Equal(t, 2286, sumOfMinCubes)
}

func TestGetSumOfMinCubes(t *testing.T) {
	// Given
	input := file.MustReadToString("day02.txt")

	// When
	sumOfMinCubes := getSumOfGamePower(input)

	// Then
	assert.Equal(t, 72706, sumOfMinCubes)
}
