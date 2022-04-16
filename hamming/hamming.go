package hamming

import (
	"fmt"
)

/*
Distance calculates the Hamming distance between two strands of DNA.
If we compare two strands of DNA and count the differences
between them we can see how many mistakes occurred.

The Hamming distance is only defined for sequences of equal length,
so an attempt to calculate it between sequences of different lengths should not work.
*/
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("hamming distance is only defined for sequences of equal length")
	}

	var differences int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			differences++
		}
	}
	return differences, nil
}
