package grains

import (
	"errors"
)

// Square computes the number of grains on a square 2^(n-1)
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("Square Number, n, must be an integer between 1 and 64")
	}
	// each shift binary shift raises by power of two
	return 1 << uint64(n-1), nil
}

// Total returns the total grains on a board of 64 squares
func Total() uint64 {
	// takes the complement of all 0's.
	// Must write 0 inside of uint64, otherwise 0 defaults to signed int
	return ^uint64(0)
}
