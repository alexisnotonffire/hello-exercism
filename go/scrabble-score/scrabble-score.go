// Package scrabble provides two calculators for working out your
// scrabble score with or without a defined word placement.
package scrabble

import (
	"fmt"
	"strconv"
	"strings"
)

// Points for all letters worth more than 1
var letterPoints = map[string]int{
	"D": 2, "G": 2,
	"B": 3, "C": 3, "M": 3, "P": 3,
	"F": 4, "H": 4, "V": 4, "W": 4, "Y": 4,
	"K": 5,
	"J": 8, "X": 8,
	"Q": 10, "Z": 10,
}

// Score calculates the number of points a word would be worth in
// a game of scrabble without any other modifiers.
func Score(word string) int {
	score := 0
	for _, letter := range strings.ToUpper(word) {
		letterScore := 0
		if letterPoints[string(letter)] == 0 {
			letterScore = 1 // letters not in letterPoints have value 1
		} else {
			letterScore = letterPoints[string(letter)]
		}
		score += letterScore
	}
	return score
}

// PositionedScore extends the above calculator to include modifiers.
// Modifiers must be single digit, non-negative numbers.
// The input should now look like:
//     word      : "helloworld" --the word you played
//     letterMods: "1112113211" --a string of modifiers for each letter
//     wordMod   : 3            --the word modifier
func PositionedScore(word string, letterMods string, wordMod int) (int, error) {
	if len(word) != len(letterMods) { // fail if strings are different lengths
		difference := len(word) - len(letterMods)
		err := fmt.Errorf("word arg has %q letter difference from letterMods", difference)
		return -1, err
	}
	if wordMod < 0 { // fail negative scores
		err := fmt.Errorf("wordMod must be non-negative")
		return -1, err
	}
	score := 0
	for i, letter := range strings.ToUpper(word) {
		letterScore := 0
		if letterPoints[string(letter)] == 0 {
			letterScore = 1
		} else {
			letterScore = letterPoints[string(letter)]
		}
		modifier, err := strconv.Atoi(string(letterMods[i]))
		if err != nil {
			return -1, err
		}
		letterScore *= modifier
		score += letterScore
	}
	score *= wordMod
	return score, nil
}
