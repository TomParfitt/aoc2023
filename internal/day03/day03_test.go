package day03

import (
	"fmt"
	"os"
	"testing"
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
	if sum != 4361 {
		t.Fatalf(`sumPartNumbers with example value %v != 4361`, sum)
	}
}

func TestSumPartNumbers(t *testing.T) {
	// Given
	b, err := os.ReadFile("day03.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)

	// When
	sum := sumPartNumbers(str)

	// Then
	if sum != 551094 {
		t.Fatalf(`sumPartNumbers with example value %v != 551094`, sum)
	}
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
	if sum != 467835 {
		t.Fatalf(`sumPartNumbers2 with example value %v != 467835`, sum)
	}
}

func TestSumPartNumbers2(t *testing.T) {
	// Given
	b, err := os.ReadFile("day03.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)

	// When
	sum := sumPartNumbers2(str)

	// Then
	if sum != 80179647 {
		t.Fatalf(`sumPartNumbers with example value %v != 80179647`, sum)
	}
}
