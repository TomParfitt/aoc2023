package day01

import (
	"testing"

	"github.com/TomParfitt/aoc2023/internal/utils/file"
	"github.com/stretchr/testify/assert"
)

func TestGetClibrationValueExample(t *testing.T) {
	// Given
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	// When
	calibrationValue := getSumOfCalibrationValue(input)

	// Then
	assert.Equal(t, 142, calibrationValue)
}

func TestGetClibrationValue(t *testing.T) {
	// Given
	input := file.MustReadToString("day01.txt")

	// When
	calibrationValue := getSumOfCalibrationValue(input)

	// Then
	assert.Equal(t, 52834, calibrationValue)
}

func TestGetClibrationValueExample2(t *testing.T) {
	// Given
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	// When
	calibrationValue := getSumOfCalibrationValue(input)

	// Then
	assert.Equal(t, 281, calibrationValue)
}

func TestGetClibrationValue2(t *testing.T) {
	// Given
	input := file.MustReadToString("day01.txt")

	// When
	calibrationValue := getSumOfCalibrationValue(input)

	// Then
	assert.Equal(t, 52834, calibrationValue)
}
