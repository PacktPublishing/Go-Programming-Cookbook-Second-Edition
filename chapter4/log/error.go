package log

import "github.com/pkg/errors"
import "log"

// OriginalError returns the error original error
func OriginalError() error {
	return errors.New("error occurred")
}

// PassThroughError calls OriginalError and
// forwards the error along after wrapping.
func PassThroughError() error {
	err := OriginalError()
	// no need to check error
	// since this works with nil
	return errors.Wrap(err, "in passthrougherror")
}

// FinalDestination deals with the error
// and doesn't forward it
func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		// we log because an unexpected error occurred!
		log.Printf("an error occurred: %s\n", err.Error())
		return
	}
}
