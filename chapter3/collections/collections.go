package collections

// WorkWith is the struct we'll
// be implementing collections for
type WorkWith struct {
	Data    string
	Version int
}

// Filter is a functional filter. It takes a list of
// WorkWith and a WorkWith Function that returns a bool
// for each "true" element we return it to the resultant
// list
func Filter(ws []WorkWith, f func(w WorkWith) bool) []WorkWith {
	// depending on results, smalles size for result
	// is len == 0
	result := make([]WorkWith, 0)
	for _, w := range ws {
		if f(w) {
			result = append(result, w)
		}
	}
	return result
}

// Map is a functional map. It takes a list of
// WorkWith and a WorkWith Function that takes a WorkWith
// and returns a modified WorkWith. The end result is
// a list of modified WorkWiths
func Map(ws []WorkWith, f func(w WorkWith) WorkWith) []WorkWith {
	// the result should always be the same
	// length
	result := make([]WorkWith, len(ws))

	for pos, w := range ws {
		newW := f(w)
		result[pos] = newW
	}
	return result
}
