package structured

import (
	"errors"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

// ThrowError throws an error that we'll trace
func ThrowError() error {
	err := errors.New("a crazy failure")
	log.WithField("id", "123").Trace("ThrowError").Stop(&err)
	return err
}

// CustomHandler splits to two streams
type CustomHandler struct {
	id      string
	handler log.Handler
}

// HandleLog adds a hook and does the emitting
func (h *CustomHandler) HandleLog(e *log.Entry) error {
	e.WithField("id", h.id)
	return h.handler.HandleLog(e)
}

// Apex has a number of useful tricks
func Apex() {
	log.SetHandler(&CustomHandler{"123", text.New(os.Stdout)})
	err := ThrowError()

	//With error convenience function
	log.WithError(err).Error("an error occurred")
}
