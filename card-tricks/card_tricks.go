package cards

// FavoriteCards returns a slice with cards in that order.
// favorite three cards of the deck are: 2, 6 and 9.
func FavoriteCards() []int {
	var favoriteCards []int
	favoriteCards = append(favoriteCards, 2, 6, 9)
	return favoriteCards
}

// GetItem return the card at position 'index' from the given stack.
// If the index is out of bounds (ie. if it is negative or after the end of the stack),
// we want to return '-1'
func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice) {
		return -1
	}
	return slice[index]
}

// SetItem exchanges the card at position index with the new card provided and return the adjusted stack.
// This will modify the input slice which is the expected behavior.
// If the index is out of bounds (ie. if it is negative or after the end of the stack),
// we want to append the new card to the end of the stack
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index >= len(slice) {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}

// PrependItems add the card(s) specified in the value parameter at the top of the stack.
// If no argument is given for the value parameter, then the result equals the original slice.
func PrependItems(slice []int, value ...int) []int {
	return append(value, slice...)
}

// RemoveItem removes the card at position index from the stack and return the stack.
// Note that this may modify the input slice which is ok.
// If the index is out of bounds (ie. if it is negative or after the end of the stack),
// leaves the stack unchanged.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
