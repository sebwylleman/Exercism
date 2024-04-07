package reverse

func Reverse(input string) string {
	runes := []rune(input)
	l := 0
	r := len(runes) - 1
	for l < r {
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}
	return string(runes)
}
