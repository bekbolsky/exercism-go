package luhn

import (
	"strings"
)

// Valid determines whether or not given number is valid per the Luhn formula.
// Strings of length 1 or less are not valid. Spaces are allowed in the input,
// but they should be stripped before checking.
// All other non-digit characters are disallowed.
func Valid(id string) bool {
	id = strings.Replace(id, " ", "", -1)
	if len(id) <= 1 {
		return false
	}
	runeId := []rune(id)
	sum := 0
	length := len(runeId)
	for i := length - 1; i >= 0; i-- {
		if runeId[i] < '0' || runeId[i] > '9' {
			return false
		}
		runeToInt := int(runeId[i] - '0')
		if length%2 == 0 {
			if i%2 == 0 {
				if runeToInt*2 > 9 {
					sum += (runeToInt*2 - 9)
				} else {
					sum += runeToInt * 2
				}
			} else {
				sum += runeToInt
			}
		} else {
			if i%2 == 0 {
				sum += runeToInt
			} else {
				if runeToInt*2 > 9 {
					sum += (runeToInt*2 - 9)
				} else {
					sum += runeToInt * 2
				}
			}
		}
	}
	return sum%10 == 0
}
