package day04

import (
	"fmt"
	"os"
	"testing"

	"github.com/TomParfitt/aoc2023/internal/utils"
)

func TestGetPointsExample(t *testing.T) {
	// Given
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	// When
	valiIdsSum := utils.SumProcessResults(input, getPoints)

	// Then
	if valiIdsSum != 13 {
		t.Fatalf(`getValidIdsSum with example value %v != 13`, valiIdsSum)
	}
}

func TestGetPoints(t *testing.T) {
	// Given
	b, err := os.ReadFile("day04.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)

	// When
	pointsSum := utils.SumProcessResults(str, getPoints)

	// Then
	if pointsSum != 22488 {
		t.Fatalf(`getPoints with example value %v != 22488`, pointsSum)
	}
}

func TestGetNumberOfCardsExample(t *testing.T) {
	// Given
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	// When
	getNumberOfCards := getNumberOfCards(input)

	// Then
	if getNumberOfCards != 30 {
		t.Fatalf(`getSumOfMinCubes with example value %v != 30`, getNumberOfCards)
	}
}

func TestGetNumberOfCards(t *testing.T) {
	// Given
	b, err := os.ReadFile("day04.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)

	// When
	numberOfCards := getNumberOfCards(str)

	// Then
	if numberOfCards != 7013204 {
		t.Fatalf(`getNumberOfCards with example value %v != 7013204`, numberOfCards)
	}
}
