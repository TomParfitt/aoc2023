package internal

import (
	"fmt"
	"log/slog"
	"os"
	"testing"
)

func TestGetClibrationValueExample(t *testing.T) {
	// Given
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	// When
	calibrationValue := getCalibrationValue(input)

	// Then
	if calibrationValue != 142 {
		t.Fatalf(`getCalibrationValue with example value %v != 142`, calibrationValue)
	}
}

func TestGetClibrationValue(t *testing.T) {
	// Given
	b, err := os.ReadFile("day1.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)

	// When
	calibrationValue := getCalibrationValue(str)

	// Then
	slog.Info("TestGetClibrationValue", "calibrationValue", calibrationValue)
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
	calibrationValue := getCalibrationValue(input)

	// Then
	if calibrationValue != 281 {
		t.Fatalf(`getCalibrationValue with example value %v != 281`, calibrationValue)
	}
}

func TestGetClibrationValue2(t *testing.T) {
	// Given
	b, err := os.ReadFile("day1.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)

	// When
	calibrationValue := getCalibrationValue(str)

	// Then
	slog.Info("TestGetClibrationValue", "calibrationValue", calibrationValue)
}
