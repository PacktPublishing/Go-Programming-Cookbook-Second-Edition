package structured

import "github.com/sirupsen/logrus"

// Hook will implement the logrus
// hook interface
type Hook struct {
	id string
}

// Fire will trigger whenever you log
func (hook *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["id"] = hook.id
	return nil
}

// Levels is what levels this hook will fire on
func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Logrus demonstrates some basic logrus functionality
func Logrus() {
	// we're emitting in json format
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(&Hook{"123"})

	fields := logrus.Fields{}
	fields["success"] = true
	fields["complex_struct"] = struct {
		Event string
		When  string
	}{"Something happened", "Just now"}

	x := logrus.WithFields(fields)
	x.Warn("warning!")
	x.Error("error!")
}
