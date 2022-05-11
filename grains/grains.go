package grains

import (
	"fmt"
	"math"
)

// Square returns the number of grains on the square
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, fmt.Errorf("invalid number: %d", number)
	}
	return uint64(math.Pow(2, float64(number-1))), nil
	// using bitwise operators
	// return 1 << (uint64(number) - 1), nil
}

// Total returns the total number of grains
func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		squareOfGrain, _ := Square(i)
		total += squareOfGrain
	}
	return total
	// using bitwise operators
	// return 1<<64 - 1
}
