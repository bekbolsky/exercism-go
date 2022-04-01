package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// SillyNephewError is a custom error type.
// It should be returned in case the number of 'cows' is negative.
// The error message should include the number of 'cows' that was passed as argument.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// ErrNegativeFodder occurs when amounts of fodder are negative.
var ErrNegativeFodder = errors.New("negative fodder")

// ErrDivisionByZero occurs to prevent division by zero.
var ErrDivisionByZero = errors.New("division by zero")

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()
	if cows < 0 {
		return 0, &SillyNephewError{cows}
	}

	if err == ErrScaleMalfunction && fodderAmount > 0 {
		return fodderAmount * 2 / float64(cows), nil
	} else if err == ErrScaleMalfunction && fodderAmount < 0 {
		return 0, ErrNegativeFodder
	} else if err != nil {
		return 0, err
	}

	if fodderAmount < 0 {
		return 0, ErrNegativeFodder
	}

	if cows == 0 {
		return 0, ErrDivisionByZero
	}

	return fodderAmount / float64(cows), nil
}
