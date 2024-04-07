// Package twofer provides a function that generates a message on how to evenly split two cookies between people.
package twofer

import "fmt"

// ShareWith takes a name and returns a message by either mentioning the given name or by addressing as 'you' when no name is given.
func ShareWith(name string) string {
	// if a name isn't provided, its value will default to ""
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
