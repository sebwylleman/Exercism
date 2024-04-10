// Package luhn contains a checksum formula used to validate a variety of identification numbers, such as credit card numbers and Canadian Social Insurance Numbers.
package luhn

import (
	"strings"
	"unicode"
)

// Valid takes string input and returns a boolean determining whether or not it is a valid identification number according to the Luhn formula.
func Valid(cardNumber string) bool {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
	if len(cardNumber) < 2 {
		return false
	}

	// If even, luhn's doubling algorithm will proceed on the first character. Otherwise, it will start on the next one.
	isEven := len(cardNumber)%2 == 0
	sum, digit := 0, 0
	for _, character := range cardNumber {
		if !unicode.IsDigit(character) {
			return false
		}
		// Convert character (rune) to integer for arithmetic computations. Subtracting the current unicode decimal value by 48 (the ASCII decimal representation of '0')
		digit = int(character) - 48
		// If the length is even then we want to perform luhn's double formula.
		if isEven {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		// Switch the boolean value process the next digit differently.
		isEven = !isEven
	}
	// Returns true to validate the identification number.
	return sum%10 == 0
}
