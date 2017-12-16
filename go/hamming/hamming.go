// Package hamming provides a function for calculating the hamming
// distance between two DNA sequences.
package hamming

import "fmt"

// Distance returns the Hamming distance between two provided DNA
// sequences. If the sequences are not the same length an error is
// raised.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("sequences must be of same length")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil

}
