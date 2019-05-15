package tweak

import (
	"strings"
)

// StringTweaker is a type of string
// that can reverse itself
type StringTweaker struct{}

// Args are a list of options for how to tweak
// the string
type Args struct {
	String  string
	ToUpper bool
	Reverse bool
}

// Tweak conforms to the RPC library which require:
// - the method's type is exported.
// - the method is exported.
// - the method has two arguments, both exported (or builtin) types.
// - the method's second argument is a pointer.
// - the method has return type error.
func (s StringTweaker) Tweak(args *Args, resp *string) error {

	result := string(args.String)
	if args.ToUpper {
		result = strings.ToUpper(result)
	}
	if args.Reverse {
		runes := []rune(result)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		result = string(runes)

	}
	*resp = result
	return nil
}
