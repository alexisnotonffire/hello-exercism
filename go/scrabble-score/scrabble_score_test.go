package scrabble

import "testing"

func TestScore(t *testing.T) {
	for _, test := range scrabbleScoreTests {
		if actual := Score(test.input); actual != test.expected {
			t.Errorf("Score(%q) expected %d, Actual %d", test.input, test.expected, actual)
		}
	}
}

// This test extends the exercise to include score modifiers.
func TestPositionedScore(t *testing.T) {
	for _, test := range scrabblePositionedScoreTests {
		actual, err := PositionedScore(test.inputWord, test.inputLMods, test.inputWMod)
		if actual < 0 && err == nil {
			t.Fatalf("Score(%q) should have raised an error", test.inputWord)
		}
		if actual >= 0 && err != nil {
			t.Fatalf("Score(%q) should not have returned an error", test.inputWord)
		}
		if actual != test.expected {
			t.Fatalf("Score(%q) expected %d, actual %d", test.inputWord, test.expected, actual)
		}
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreTests {
			Score(test.input)
		}
	}
}
