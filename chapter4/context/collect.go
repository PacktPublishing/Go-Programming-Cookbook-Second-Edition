package context

import (
	"context"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

// Initialize calls 3 functions to set up, then
// logs before terminating
func Initialize() {
	// set basic log up
	log.SetHandler(text.New(os.Stdout))
	// initialize our context
	ctx := context.Background()
	// create a logger and link it to
	// the context
	ctx, e := FromContext(ctx, log.Log)

	// set a field
	ctx = WithField(ctx, "id", "123")
	e.Info("starting")
	gatherName(ctx)
	e.Info("after gatherName")
	gatherLocation(ctx)
	e.Info("after gatherLocation")
}

func gatherName(ctx context.Context) {
	ctx = WithField(ctx, "name", "Go Cookbook")
}

func gatherLocation(ctx context.Context) {
	ctx = WithFields(ctx, log.Fields{"city": "Seattle", "state": "WA"})
}
