package thefarm

import (
	"errors"
	"fmt"
)

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, numCows int) (float64, error) {
	totalFodder, err1 := fc.FodderAmount(numCows)
	if err1 != nil {
		return 0, err1
	}
	Fatfactor, err2 := fc.FatteningFactor()
	if err2 != nil {
		return 0, err2
	}
	return (totalFodder / float64(numCows)) * Fatfactor, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64, error) {
	if numCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, numCows)
}

type InvalidCowsError struct {
	numCows int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numCows, e.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return &InvalidCowsError{numCows, "there are no negative cows"}
	}
	if numCows == 0 {
		return &InvalidCowsError{numCows, "no cows don't need food"}
	}
	return nil
}
