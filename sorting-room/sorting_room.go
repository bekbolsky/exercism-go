package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
// Output example: "This is the number -12.3"
func DescribeNumber(f float64) string {
	parsedFloat := strconv.FormatFloat(f, 'f', 1, 64)
	return fmt.Sprintf("This is the number %s", parsedFloat)
}

// NumberBox is a generic interface for a number.
type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
// Including one digit after the decimal.
// Output example: "This is a box containing the number 2.0"
func DescribeNumberBox(nb NumberBox) string {
	intToFloat := float64(nb.Number())
	formattedFloat := strconv.FormatFloat(intToFloat, 'f', 1, 64)
	return fmt.Sprintf("This is a box containing the number %s", formattedFloat)
}

// FancyNumber is a number that has a value.
type FancyNumber struct {
	n string
}

// Value returns the value of the FancyNumber.
func (i FancyNumber) Value() string {
	return i.n
}

// FancyNumberBox is a generic interface for a FancyNumber.
type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type) {
	case FancyNumber:
		extractedInt, _ := strconv.Atoi(fnb.Value())
		return extractedInt
	default:
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
// Output example: "This is a fancy box containing the number 10.0"
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	output := "This is a fancy box containing the number 0.0"
	extractedInt := ExtractFancyNumber(fnb)
	if extractedInt != 0 {
		intToFloat := float64(extractedInt)
		formattedFloat := strconv.FormatFloat(intToFloat, 'f', 1, 64)
		output = fmt.Sprintf("This is a fancy box containing the number %s", formattedFloat)
	}
	return output
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch i.(type) {
	case int:
		return DescribeNumber(float64(i.(int)))
	case float64:
		return DescribeNumber(i.(float64))
	case NumberBox:
		return DescribeNumberBox(i.(NumberBox))
	case FancyNumberBox:
		return DescribeFancyNumberBox(i.(FancyNumberBox))
	default:
		return "Return to sender"
	}
}
