// Package twofer implements a single, boring function
package twofer

// ShareWith takes in an optional argument which is printed back in a stringg to the user
// If not argument is provided, the name defaults to "you"
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
