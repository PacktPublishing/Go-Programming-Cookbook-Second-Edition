package pipeline

import "context"

// Worker have one role
// that is determined when
// Work is called
type Worker struct {
	in  chan string
	out chan string
}

// Job is a job a worker can do
type Job string

const (
	// Print echo's all input to
	// stdout
	Print Job = "print"
	// Encode base64 encodes input
	Encode Job = "encode"
)

// Work is how to dispatch a worker, they are assigned
// a job here
func (w *Worker) Work(ctx context.Context, j Job) {
	switch j {
	case Print:
		w.Print(ctx)
	case Encode:
		w.Encode(ctx)
	default:
		return
	}
}
