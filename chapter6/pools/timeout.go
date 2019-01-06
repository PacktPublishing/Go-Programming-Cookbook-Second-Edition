package pools

import (
	"context"
	"time"
)

// ExecWithTimeout will timeout trying
// to get the current time
func ExecWithTimeout() error {
	db, err := Setup()
	if err != nil {
		return err
	}

	ctx := context.Background()

	// we want to timeout immediately
	ctx, can := context.WithDeadline(ctx, time.Now())

	// call cancel after we complete
	defer can()

	// our transaction is context aware
	_, err = db.BeginTx(ctx, nil)
	return err
}
