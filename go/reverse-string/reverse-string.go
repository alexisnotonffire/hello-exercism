// Package reverse provides a function
// to reverse a string
package reverse

// String reverses a string
func String(s string) string {
	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}
