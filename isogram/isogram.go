package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	for _, letter := range word {
		if unicode.IsSpace(letter) || unicode.IsPunct(letter) {
			continue
		}
		letter = unicode.ToUpper(letter)
		if seen[letter] {
			return false
		}
		seen[letter] = true
	}
	return true
}
