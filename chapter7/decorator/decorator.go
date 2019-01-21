package decorator

import "net/http"

// TransportFunc implements the RountTripper interface
type TransportFunc func(*http.Request) (*http.Response, error)

// RoundTrip just calls the original function
func (tf TransportFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return tf(r)
}

// Decorator is a convenience function to represent our
// middleware inner function
type Decorator func(http.RoundTripper) http.RoundTripper

// Decorate is a helper to wrap all the middleware
func Decorate(t http.RoundTripper, rts ...Decorator) http.RoundTripper {
	decorated := t
	for _, rt := range rts {
		decorated = rt(decorated)
	}
	return decorated
}
