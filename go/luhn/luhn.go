package luhn

import (
	"unicode"
)

func Valid(id string) bool {
	var cardDigits []int

	for _, char := range id {
		if unicode.IsDigit(char) {
			// subtracting char's unicode decimal value by 48 (the ASCII decimal representation of '0') to get the conversion in integer
			cardDigits = append(cardDigits, int(char)-48)
		} else if char != ' ' {
			return false
		}
	}
	if len(cardDigits) <= 1 {
		return false
	}
	return sum(luhnDouble(cardDigits))%10 == 0
}

func sum(cardDigits []int) int {
	var total int
	for _, digit := range cardDigits {
		total += digit
	}
	return total
}

func luhnDouble(cardDigits []int) []int {
	var result []int
	var double int
	var newDigit int
	// two pointers
	left, right := len(cardDigits)-2, len(cardDigits)-1

	for left >= -1 {
		// when the length of cardDigits is odd, the left pointer's last value will be out of bounds
		if left == -1 {
			result = append(result, cardDigits[right])
		} else {
			double = cardDigits[left] * 2
			if double > 9 {
				newDigit = double - 9
			} else {
				newDigit = double
			}
			result = append(result, newDigit, cardDigits[right])
		}
		left -= 2
		right -= 2
	}
	return result
}
