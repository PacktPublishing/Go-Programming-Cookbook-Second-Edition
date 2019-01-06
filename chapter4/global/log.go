package global

// UseLog demonstrates using our global
// log
func UseLog() error {
	if err := Init(); err != nil {
		return err
	}

	// if we were in another package these would be
	// global.WithField and
	// global.Debug
	WithField("key", "value").Debug("hello")
	Debug("test")

	return nil
}
