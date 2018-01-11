// Package grains provides functions for calculating the value and running sums
// of powers of two.

package grains

import "fmt"

// Square relies on the equivalence between binary format and powers of two.
// By using a bitshift we can easily calculate the value of 2^i
func Square(i int) (uint64, error) {
    if !(0 < i && i < 65) {
        err := fmt.Errorf("input must be in range 1 - 64")
        return 0, err
    }
    square := 1 << (uint(i) - 1)
    return uint64(square), nil
}

// Total uses a similar idea only instead of just bitshifting, we set all bits to be 1.
func Total() uint64 {
    var total uint64 = (1 << 63) + ((1 << 63) - 1)
    return total
}
