package mocking

import "errors"

var ThrowError = func() error {
	return errors.New("always fails")
}

func DoSomeStuff(d DoStuffer) error {

	if err := d.DoStuff("test"); err != nil {
		return err
	}

	if err := ThrowError(); err != nil {
		return err
	}

	return nil
}
