package scrabble

import "fmt"

// Source: exercism/problem-specifications
// Commit: 11ed503 scrabble-score: Make canonical-data.json compliant
// Problem Specifications Version: 1.0.0

type scrabbleTest struct {
	input    string
	expected int
}

var scrabbleScoreTests = []scrabbleTest{
	{"a", 1},                // lowercase letter
	{"A", 1},                // uppercase letter
	{"f", 4},                // valuable letter
	{"at", 2},               // short word
	{"zoo", 12},             // short, valuable word
	{"street", 6},           // medium word
	{"quirky", 22},          // medium, valuable word
	{"OxyphenButazone", 41}, // long, mixed-case word
	{"pinata", 8},           // english-like word
	{"", 0},                 // empty input
	{"abcdefghijklmnopqrstuvwxyz", 87}, // entire alphabet available
}

// These test cases extend the package to include score modifiers
type scrabblePositionedTest struct {
	inputWord  string
	inputLMods string
	inputWMod  int
	expected   int
	err        error
}

var scrabblePositionedScoreTests = []scrabblePositionedTest{
	{"easy", "1111", 1, 7, nil},
	{"abcdefghijklmnopqrstuvwxyz", "11111111111111111111111111", 1, 87, nil},
	{"aaaa", "2111", 1, 5, nil},
	{"aaaz", "2113", 1, 34, nil},
	{"aazz", "1111", 2, 44, nil},
	{"azzz", "2211", 2, 84, nil},
	{"zzzz", "1", 1, -1, fmt.Errorf("word arg has 3 letters than letterMods")},
	{"negative", "11111111", -1, -1, fmt.Errorf("wordMod must be non-negative")},
}
