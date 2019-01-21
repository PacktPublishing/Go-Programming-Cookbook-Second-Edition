package client

import (
	"crypto/tls"
	"net/http"
)

// Setup configures our client and redefines
// the global DefaultClient
func Setup(isSecure, nop bool) *http.Client {
	c := http.DefaultClient

	// Sometimes for testing, we want to
	// turn off SSL verification
	if !isSecure {
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		}
	}
	if nop {
		c.Transport = &NopTransport{}
	}
	http.DefaultClient = c
	return c
}

// NopTransport is a No-Op Transport
type NopTransport struct {
}

// RoundTrip Implements RoundTripper interface
func (n *NopTransport) RoundTrip(*http.Request) (*http.Response, error) {
	// note this is an unitialized Response
	// if you're looking at headers etc
	return &http.Response{StatusCode: http.StatusTeapot}, nil
}
