package rest

import "net/http"

// APIClient is our custom client
type APIClient struct {
	*http.Client
}

// NewAPIClient constructor initializes the client with our
// custom Transport
func NewAPIClient(username, password string) *APIClient {
	t := http.Transport{}
	return &APIClient{
		Client: &http.Client{
			Transport: &APITransport{
				Transport: &t,
				username:  username,
				password:  password,
			},
		},
	}
}

// GetGoogle is an API Call - we abstract away
// the REST aspects
func (c *APIClient) GetGoogle() (int, error) {
	resp, err := c.Get("http://www.google.com")
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}
