package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ten", "jack", "queen", "king":
		return 10
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	c1 := ParseCard(card1)
	c2 := ParseCard(card2)
	dc := ParseCard(dealerCard)

	switch {
	case c1 == 11 && c2 == 11:
		return "P"

	case c1+c2 > 11 && c1+c2 < 17:
		if dc >= 7 {
			return "H"
		}
		return "S"
	case (c1+c2 == 21):
		if 2 < dc && dc < 10 {
			return "W"
		}
		return "S"
	case c1+c2 <= 11:
		return "H"
	case c1+c2 > 16 && c1+c2 < 21:
		return "S"
	}
	return ""
}
