package async

import "net/http"

// NewClient creates a new client and sets its appropriate channels
func NewClient(client *http.Client, bufferSize int) *Client {
	respch := make(chan *http.Response, bufferSize)
	errch := make(chan error, bufferSize)
	return &Client{
		Client: client,
		Resp:   respch,
		Err:    errch,
	}
}

// Client stores a client and has two channels to aggregate
// responses and errors
type Client struct {
	*http.Client
	Resp chan *http.Response
	Err  chan error
}

// AsyncGet performs a Get then returns
// the resp/error to the appropriate channel
func (c *Client) AsyncGet(url string) {
	resp, err := c.Get(url)
	if err != nil {
		c.Err <- err
		return
	}
	c.Resp <- resp
}
