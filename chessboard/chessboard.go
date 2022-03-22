package chessboard

// Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

/*
CountInRank returns how many squares are occupied in the chessboard,
within the given rank

Counts the total number of occupied squares by ranging over a map. Returns an integer.

Returns a count of zero (0) if the given rank cannot be found in the map.
*/
func CountInRank(cb Chessboard, rank string) int {
	rankValue, rankExists := cb[rank]
	if !rankExists {
		return 0
	}
	var squaresInRank int
	for _, squareIsOccupied := range rankValue {
		if squareIsOccupied {
			squaresInRank++
		}
	}
	return squaresInRank
}

/*
CountInFile returns how many squares are occupied in the chessboard,
within the given file

Counts the total number of occupied squares by ranging over the given file. Returns an integer.

Returns a count of zero (0) if the given file is not a valid one (not between 1 and 8, inclusive).
*/
func CountInFile(cb Chessboard, file int) int {
	if file < 1 || file > 8 {
		return 0
	}

	var squaresInFile int
	for _, rank := range cb {
		if rank[file-1] {
			squaresInFile++
		}
	}
	return squaresInFile
}

// CountAll counts how many squares are present in the chessboard and returns an integer.
// Since we don't need to check the content of the squares,
// consider using range omitting both index and value.
func CountAll(cb Chessboard) int {
	var squares int
	for _, rankValues := range cb {
		for range rankValues {
			squares++
		}
	}
	return squares
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var occupiedSquares int
	for rank := range cb {
		occupiedSquares += CountInRank(cb, rank)
	}
	return occupiedSquares
}
