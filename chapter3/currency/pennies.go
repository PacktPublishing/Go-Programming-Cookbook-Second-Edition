package currency

import (
	"strconv"
)

// ConvertPenniesToDollarString takes a penny amount as
// an int64 and returns a dollar string representation
func ConvertPenniesToDollarString(amount int64) string {
	// parse the pennies as a base 10 int
	result := strconv.FormatInt(amount, 10)

	// check if negative, will set it back later
	negative := false
	if result[0] == '-' {
		result = result[1:]
		negative = true
	}

	// left pad with 0 if we're passed in value < 100
	for len(result) < 3 {
		result = "0" + result
	}
	length := len(result)

	// add in the decimal
	result = result[0:length-2] + "." + result[length-2:]

	// from the negative we stored earlier!
	if negative {
		result = "-" + result
	}

	return result
}
