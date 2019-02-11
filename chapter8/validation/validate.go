package validation

import "errors"

// Verror is an error that occurs
// during validation, we can
// return this to a user
type Verror struct {
	error
}

// Payload is the value we
// process
type Payload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ValidatePayload is 1 implementation of
// the closure in our controller
func ValidatePayload(p *Payload) error {
	if p.Name == "" {
		return Verror{errors.New("name is required")}
	}

	if p.Age <= 0 || p.Age >= 120 {
		return Verror{errors.New("age is required and must be a value greater than 0 and less than 120")}
	}
	return nil
}
