package context

import (
	"context"

	"github.com/apex/log"
)

type key int

// logFields is a key we use
// for our context logging
const logFields key = 0

func getFields(ctx context.Context) *log.Fields {
	fields, ok := ctx.Value(logFields).(*log.Fields)
	if !ok {
		f := make(log.Fields)
		fields = &f
	}
	return fields
}

// FromContext takes an entry and a context
// then returns an entry populated from the context object
func FromContext(ctx context.Context, l log.Interface) (context.Context, *log.Entry) {
	fields := getFields(ctx)
	e := l.WithFields(fields)
	ctx = context.WithValue(ctx, logFields, fields)
	return ctx, e
}

// WithField adds a log field to the context
func WithField(ctx context.Context, key string, value interface{}) context.Context {
	return WithFields(ctx, log.Fields{key: value})
}

// WithFields adds many log fields to the context
func WithFields(ctx context.Context, fields log.Fielder) context.Context {
	f := getFields(ctx)
	for key, val := range fields.Fields() {
		(*f)[key] = val
	}
	ctx = context.WithValue(ctx, logFields, f)
	return ctx
}
