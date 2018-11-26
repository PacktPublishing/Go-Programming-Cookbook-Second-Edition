package dataconv

import (
	"fmt"
	"strconv"
)

// Strconv demonstrates some strconv
// functions
func Strconv() error {
	//strconv is a good way to convert to and from strings
	s := "1234"
	// we can specify the base (10) and precision
	// 64 bit
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	fmt.Println(res)

	// lets try hex
	res, err = strconv.ParseInt("FF", 16, 64)
	if err != nil {
		return err
	}

	fmt.Println(res)

	// we can do other useful things like:
	val, err := strconv.ParseBool("true")
	if err != nil {
		return err
	}

	fmt.Println(val)

	return nil
}
