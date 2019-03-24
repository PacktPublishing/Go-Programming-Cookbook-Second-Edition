package pipeline

import (
	"context"
	"fmt"
)

// Print prints w.in and repalys it
// on w.out
func (w *Worker) Print(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case val := <-w.in:
			fmt.Println(val)
			w.out <- val
		}
	}
}
