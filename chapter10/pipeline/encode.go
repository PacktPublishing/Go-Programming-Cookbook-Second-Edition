package pipeline

import (
	"context"
	"encoding/base64"
	"fmt"
)

// Encode takes plain text as int
// and returns "string => <base64 string encoding>
// as out
func (w *Worker) Encode(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case val := <-w.in:
			w.out <- fmt.Sprintf("%s => %s", val, base64.StdEncoding.EncodeToString([]byte(val)))
		}
	}
}
