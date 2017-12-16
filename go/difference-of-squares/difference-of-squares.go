// Package diffsquares provides a function for calculating the
// difference between the sum of the first N squares and the square
// of the sum of the first N numbers; i.e:
//    Difference(N) == (1 + ... + N)^2 - (1^2 + ... + N^2)
package diffsquares

// SumOfSquares calculates (1^2 + ... + N^2)
func SumOfSquares(N int) int {
	sumOfSquares := 0
	for n := 1; n <= N; n++ {
		sumOfSquares += n * n
	}
	return sumOfSquares
}

// SquareOfSums calculates (1 + ... + N)^2
func SquareOfSums(N int) int {
	squareOfSums := 0
	for n := 1; n <= N; n++ {
		squareOfSums += n
	}
	return squareOfSums * squareOfSums
}

// Difference calculates the... difference.
func Difference(N int) int {
	difference := SquareOfSums(N) - SumOfSquares(N)
	return difference
}
