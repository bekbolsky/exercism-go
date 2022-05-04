package isogram

// IsIsogram returns true if the given word is an isogram.
// An isogram is a word or phrase without a repeating letter.
func IsIsogram(word string) bool {
	// 'seen' is a slice of type bool
	seen := make([]bool, 26)
	for _, letter := range word {
		// 'letter' is a rune
		if letter == '-' || letter == ' ' {
			continue
		}
		// make letter uppercase
		if !(letter >= 'A' && letter <= 'Z') {
			letter = letter - ('a' - 'A')
		}
		if seen[letter-'A'] {
			return false
		}
		seen[letter-'A'] = true
	}
	return true
}
