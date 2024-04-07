package isogram

import "strings"

func IsIsogram(word string) bool {
	letterCount := make(map[rune]int)
	lowerCaseWords := strings.ToLower(word)

	for _, letter := range lowerCaseWords {
		if letterCount[letter] != 0 && letter != '-' && letter != ' ' {
			return false
		}
		letterCount[letter]++
	}
	return true
}
