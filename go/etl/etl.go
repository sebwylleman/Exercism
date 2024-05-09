package etl

import (
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	newValues := make(map[string]int)

	for score, letters := range in {
		for _, letter := range letters {
			newValues[strings.ToLower(letter)] = score
		}
	}
	return newValues
}
