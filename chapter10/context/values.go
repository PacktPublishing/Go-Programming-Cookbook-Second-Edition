package context

import "context"

type key string

const (
	timeoutKey  key = "TimeoutKey"
	deadlineKey key = "DeadlineKey"
)

// Setup sets some values
func Setup(ctx context.Context) context.Context {

	ctx = context.WithValue(ctx, timeoutKey, "timeout exceeded")
	ctx = context.WithValue(ctx, deadlineKey, "deadline exceeded")

	return ctx
}

// GetValue grabs a value given a key and
// returns a string representation of the
// value
func GetValue(ctx context.Context, k key) string {

	if val, ok := ctx.Value(k).(string); ok {
		return val
	}
	return ""

}
