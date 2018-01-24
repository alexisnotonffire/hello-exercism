// Package letter provides functions for determining
// rune frequency within a given string, or list of strings.
package letter

// FreqMap maps letters to the number of times they appear
type FreqMap map[rune]int

// Frequency returns a FreqMap for a given string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calculates a FreqMap for a list
// of strings by calculating Frequency for each list item
// in parallel before totalling at the end.
func ConcurrentFrequency(list []string) FreqMap {
	m := FreqMap{}
	c := make(chan FreqMap)

	for _, s := range list {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	for _ = range list {
		for r, n := range <-c {
			m[r] += n
		}
	}

	return m
}
