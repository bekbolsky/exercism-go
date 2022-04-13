// Package twofer is a Go package that provides a function ShareWith.
// It returns a string with the message "One for X, one for me."
package twofer

// ShareWith return a string with the message "One for X, one for me."
// where X is the name passed as argument.
// If no name is passed, the message should be "One for you, one for me."
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
