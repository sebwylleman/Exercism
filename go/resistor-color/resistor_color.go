package resistorcolor

var resistorValue = make(map[string]int)

// Colors returns the list of all colors.
func Colors() []string {
	return []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {

	resistorValue := make(map[string]int)
	for i, v := range Colors() {
		resistorValue[v] = i
	}

	value, ok := resistorValue[color]
	if !ok {
		return 0
	}
	return value
}
