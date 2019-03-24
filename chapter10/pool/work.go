package pool

import "errors"

type op string

const (
	// Hash is the bcrypt work type
	Hash op = "encrypt"
	// Compare is bcrypt compare work
	Compare = "decrypt"
)

// WorkRequest is a worker req
type WorkRequest struct {
	Op      op
	Text    []byte
	Compare []byte // optional
}

// WorkResponse is a worker resp
type WorkResponse struct {
	Wr      WorkRequest
	Result  []byte
	Matched bool
	Err     error
}

// Process dispatches work to the worker pool channel
func Process(wr WorkRequest) WorkResponse {
	switch wr.Op {
	case Hash:
		return hashWork(wr)
	case Compare:
		return compareWork(wr)
	default:
		return WorkResponse{Err: errors.New("unsupported operation")}
	}
}
