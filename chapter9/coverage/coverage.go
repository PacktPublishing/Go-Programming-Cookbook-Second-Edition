package main

import "errors"

// Coverage is a simple function with some brancing conditions
func Coverage(condition bool) error {
	if condition {
		return errors.New("condition was set")
	}
	return nil
}
