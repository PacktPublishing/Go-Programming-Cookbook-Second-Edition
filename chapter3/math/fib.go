package math

import "math/big"

// global to memoize fib
var memoize map[int]*big.Int

func init() {
	// initialize the map
	memoize = make(map[int]*big.Int)
}

// Fib prints the nth digit of the fibonacci sequence
// it will return 1 for anything < 0 as well...
// it's calculated recursively and use big.Int since
// int64 will quickly overflow
func Fib(n int) *big.Int {
	if n < 0 {
		return nil
	}

	// base case
	if n < 2 {
		memoize[n] = big.NewInt(1)
	}

	// check if we stored it before
	// if so return with no calculation
	if val, ok := memoize[n]; ok {
		return val
	}

	// initialize map then add previous 2 fib values
	memoize[n] = big.NewInt(0)
	memoize[n].Add(memoize[n], Fib(n-1))
	memoize[n].Add(memoize[n], Fib(n-2))

	// return result
	return memoize[n]
}
