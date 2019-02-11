package proxy

import (
	"log"
	"net/http"
)

// Proxy holds our configured client
// and BaseURL to proxy to
type Proxy struct {
	Client  *http.Client
	BaseURL string
}

// ServeHTTP means that proxy implments the Handler interface
// It manipulates the request, forwards it to BaseURL, then returns
// the response
func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := p.ProcessRequest(r); err != nil {
		log.Printf("error occurred during process request: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := p.Client.Do(r)
	if err != nil {
		log.Printf("error occurred during client operation: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	CopyResponse(w, resp)
}
