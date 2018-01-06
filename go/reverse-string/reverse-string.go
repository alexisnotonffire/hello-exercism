// Package reverse provides a function
// to reverse a string
package reverse

// String reverses a string
func String(s string) string {
	var reversed []byte
	for i := len(s) - 1; i >= 0; i-- {
		reversed = append(reversed, s[i])
	}
	reversedString := string(reversed)
	return reversedString
}
