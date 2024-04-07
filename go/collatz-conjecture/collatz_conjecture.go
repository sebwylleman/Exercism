package collatzconjecture

import "errors"

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
