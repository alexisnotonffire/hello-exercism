// Package raindrops provides a function which converts a number to
// a string based on its factors.
package raindrops

import "strconv"

// Convert returns the string expression of the input number based on the
// raindrop conversion.
func Convert(number int) string {
	raindrop := ""
	if number%3 == 0 {
		raindrop += "Pling"
	}
	if number%5 == 0 {
		raindrop += "Plang"
	}
	if number%7 == 0 {
		raindrop += "Plong"
	}
	if raindrop == "" {
		raindrop = strconv.Itoa(number)
	}
	return raindrop
}
