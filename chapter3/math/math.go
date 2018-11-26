package math

import (
	"fmt"
	"math"
)

// Examples demonstrates some of the functions
// in the math package
func Examples() {
	//sqrt Examples
	i := 25

	// i is an int, so convert
	result := math.Sqrt(float64(i))

	// sqrt of 25 == 5
	fmt.Println(result)

	// ceil rounds up
	result = math.Ceil(9.5)
	fmt.Println(result)

	// floor rounds down
	result = math.Floor(9.5)
	fmt.Println(result)

	// math also stores some consts:
	fmt.Println("Pi:", math.Pi, "E:", math.E)
}
