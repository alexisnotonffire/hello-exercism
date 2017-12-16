// Package acronym provides the function Abbreviate which produces a
// suitable acronym from the provided inout string.
package acronym

import "strings"

// Abbreviate breaks the input string into 'phrases' where a phrase is
// either a new word or part of a hyphenated statement. The first letter
// of each phrase creates the acronym.
func Abbreviate(fullName string) string {
	var abbreviation string
	fullName = strings.Replace(fullName, "-", " ", -1)
	phrases := strings.Split(fullName, " ")

	// For each phrase, take the first char.
	for phraseNo := 0; phraseNo < len(phrases); phraseNo++ {
		abbreviation += strings.ToUpper(phrases[phraseNo][:1])
	}

	return abbreviation
}
