package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	seenLetters := map[rune]bool{}
	lowercaseInput := strings.ToLower(input)

	for _, char := range lowercaseInput {
		if unicode.IsLetter(char) {
			seenLetters[char] = true
		}
		// 26 characters in the English alphabet
		if len(seenLetters) == 26 {
			return true
		}
	}
	return false
}
