package dataconv

import "fmt"

// CheckType will print based on the
// interface type
func CheckType(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println("It's a string!")
	case int:
		fmt.Println("It's an int!")
	default:
		fmt.Println("not sure what it is...")
	}
}

// Interfaces demonstrates casting
// from anonymous interfaces to types
func Interfaces() {
	CheckType("test")
	CheckType(1)
	CheckType(false)

	var i interface{}
	i = "test"

	// manually check an interface
	if val, ok := i.(string); ok {
		fmt.Println("val is", val)
	}

	// this one should fail
	if _, ok := i.(int); !ok {
		fmt.Println("uh oh! glad we handled this")
	}
}
