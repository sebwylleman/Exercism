package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	seenLetters := map[rune]struct{}{}
	lowercaseInput := strings.ToLower(input)

	for _, char := range lowercaseInput {
		if unicode.IsLetter(char) {
			seenLetters[char] = struct{}{}
		}
		// 26 characters in the English alphabet
		if len(seenLetters) == 26 {
			return true
		}
	}
	return false
}
