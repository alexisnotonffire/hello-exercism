// Package luhn provides a simple validation function for
// luhn compliance
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func stripSpaces(r rune) rune {
	if unicode.IsSpace(r) {
		return -1
	}
	return r
}

func stripNonDigits(r rune) rune {
	if !unicode.IsDigit(r) {
		return -1
	}
	return r
}

// converts values to their luhn defined single-digit value
func allowedValue(d int) int {
	if d > 9 {
		d -= 9
	}
	return d
}

// Valid performs a luhn validity check, calculating the value under
// idValue and evaluating with isValid
func Valid(id string) bool {
	// checks for valid input: 2+ digits and no punctuation
	id = strings.Map(stripSpaces, id)
	idLen := len(id)
	_, err := strconv.Atoi(id)
	// id must have at least 2 runes and no symbols
	if idLen < 2 || err != nil {
		return false
	}

	// the below switch statement calculates the validity of the id
	// using the idValue to calculate the sum of:
	//   - every second digit starting from the last and reading left
	//   - 2*x for the other digits under 5
	//   - 2*x - 9 for any remaining digits
	idValue := 0
	isValid := false
	switch idLen % 2 {
	case 0:
		// if even, the even indexed bytes should be doubled
		for i := 0; i < idLen; i += 2 {
			idValue += allowedValue(2 * int(id[i]-'0'))
		}
		for i := 1; i < idLen; i += 2 {
			idValue += int(id[i] - '0')
		}
		isValid = (idValue%10 == 0)

	case 1:
		// if odd, the odd indexed bytes are doubled
		for i := 0; i < idLen; i += 2 {
			idValue += int(id[i] - '0')
		}
		for i := 1; i < idLen; i += 2 {
			idValue += allowedValue(2 * int(id[i]-'0'))
		}
		isValid = (idValue%10 == 0)
	}
	return isValid
}
