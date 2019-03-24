package state

type op string

const (
	// Add values
	Add op = "add"
	// Subtract values
	Subtract = "sub"
	// Multiply values
	Multiply = "mult"
	// Divide values
	Divide = "div"
)

// WorkRequest perform an op
// on two values
type WorkRequest struct {
	Operation op
	Value1    int64
	Value2    int64
}

// WorkResponse returns the result
// and any errors
type WorkResponse struct {
	Wr     *WorkRequest
	Result int64
	Err    error
}
