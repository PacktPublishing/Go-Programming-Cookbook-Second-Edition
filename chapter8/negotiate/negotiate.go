package negotiate

import (
	"net/http"

	"github.com/unrolled/render"
)

// Negotiator wraps render and does
// some switching on ContentType
type Negotiator struct {
	ContentType string
	*render.Render
}

// GetNegotiator takes a request, and figures
// out the ContentType from the Content-Type header
func GetNegotiator(r *http.Request) *Negotiator {
	contentType := r.Header.Get("Content-Type")

	return &Negotiator{
		ContentType: contentType,
		Render:      render.New(),
	}
}
