package client

import (
	"fmt"
	"net/http"
)

// Controller embeds an http.Client
// and uses it internally
type Controller struct {
	*http.Client
}

// DoOps with a controller object
func (c *Controller) DoOps() error {
	resp, err := c.Client.Get("http://www.google.com")
	if err != nil {
		return err
	}
	fmt.Println("results of client.DoOps", resp.StatusCode)
	return nil
}
