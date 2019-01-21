package rest

import "net/http"

// APITransport does a SetBasicAuth
// for every request
type APITransport struct {
	*http.Transport
	username, password string
}

// RoundTrip does the basic auth before deferring to the
// default transport
func (t *APITransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.username, t.password)
	return t.Transport.RoundTrip(req)
}
