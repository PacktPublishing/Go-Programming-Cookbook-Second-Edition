package validation

// Controller holds our validation functions
type Controller struct {
	ValidatePayload func(p *Payload) error
}

// New initializes a controller with our
// local validation, it can be overwritten
func New() *Controller {
	return &Controller{
		ValidatePayload: ValidatePayload,
	}
}
