package mockgen

// Controller is a struct demonstrating
// one way to initialize interfaces
type Controller struct {
	GetSetter
}

// GetThenSet checks if a value is set. If not
// it sets it.
func (c *Controller) GetThenSet(key, value string) error {
	val, err := c.Get(key)
	if err != nil {
		return err
	}

	if val != value {
		return c.Set(key, value)
	}
	return nil
}
