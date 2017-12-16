// Package isogram does something
package isogram

import (
	"strings"
	"unicode"
)

func dropPunctuation(r rune) rune {
	isPunctuation := (unicode.IsPunct(r) || unicode.IsSpace(r))
	if isPunctuation {
		return -1
	}
	return r
}

// IsIsogram checks if a word is an isogram or not by comparing the
// length of the word with the first instance of a letter removed to
// the length of the word with all instances removed.
func IsIsogram(word string) bool {

	// first normalise all words
	word = strings.ToLower(word)
	word = strings.Map(dropPunctuation, word)

	for _, letter := range word {
		word = strings.Replace(word, string(letter), "", 1)
		isRepeated := len(word) != len(strings.Replace(word, string(letter), "", -1))

		if isRepeated {
			return false
		}
	}
	return true
}
