package day03

import (
	"testing"

	"github.com/TomParfitt/aoc2023/internal/utils/file"
	"github.com/stretchr/testify/assert"
)

func TestSumPartNumbersExample(t *testing.T) {
	// Given
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	// When
	sum := sumPartNumbers(input)

	// Then
	assert.Equal(t, 4361, sum)
}

func TestSumPartNumbers(t *testing.T) {
	// Given
	input := file.MustReadToString("day03.txt")

	// When
	sum := sumPartNumbers(input)

	// Then
	assert.Equal(t, 551094, sum)
}

func TestSumPartNumbers2Example(t *testing.T) {
	// Given
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	// When
	sum := sumPartNumbers2(input)

	// Then
	assert.Equal(t, 467835, sum)
}

func TestSumPartNumbers2(t *testing.T) {
	// Given
	input := file.MustReadToString("day03.txt")

	// When
	sum := sumPartNumbers2(input)

	// Then
	assert.Equal(t, 80179647, sum)
}
