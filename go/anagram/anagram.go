package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	var result []string
	lowerCaseSubject := strings.ToLower(subject)

	for _, word := range candidates {
		if isAnagram([]rune(lowerCaseSubject), []rune(strings.ToLower(word))) {
			result = append(result, word)
		}
	}
	return result
}

func isAnagram(target, candidate []rune) bool {
	if string(target) == string(candidate) {
		return false
	}
	targetCount := runeCount(target)
	candidateCount := runeCount(candidate)

	if len(targetCount) != len(candidateCount) {
		return false
	}
	for char, count := range targetCount {
		if count != candidateCount[char] {
			return false
		}
	}
	return true
}

func runeCount(runeSlice []rune) map[rune]int {
	rc := map[rune]int{}

	for _, r := range runeSlice {
		rc[r]++
	}
	return rc
}
