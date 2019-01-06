package global

import (
	"errors"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

// we make our global package level
// variable lower case
var (
	log     *logrus.Logger
	initLog sync.Once
)

// Init sets up the logger intially
// if run multiple times, it returns
// an error
func Init() error {
	err := errors.New("already initialized")
	initLog.Do(func() {
		err = nil
		log = logrus.New()
		log.Formatter = &logrus.JSONFormatter{}
		log.Out = os.Stdout
		log.Level = logrus.DebugLevel
	})
	return err
}

// SetLog sets the log
func SetLog(l *logrus.Logger) {
	log = l
}

// WithField exports the logs withfield connected
// to our global log
func WithField(key string, value interface{}) *logrus.Entry {
	return log.WithField(key, value)
}

// Debug exports the logs Debug connected
// to our global log
func Debug(args ...interface{}) {
	log.Debug(args...)
}
