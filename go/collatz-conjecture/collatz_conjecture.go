// Package collatzconjecture provides functions for implementing Collatz' conjecture.
package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps the Collatz sequence needs to reach 1.
// An error is returned if the input is not a positive integer.
func CollatzConjecture(n int) (int, error) {
	var numSteps int
	for {
		if n < 1 {
			return 0, errors.New("invalid argument")
		}
		if n == 1 {
			return numSteps, nil
		}
		if n%2 == 0 {
			n /= 2
			numSteps++
		} else {
			n = n*3 + 1
			numSteps++
		}
	}
}
