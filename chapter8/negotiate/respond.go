package negotiate

import "io"
import "github.com/unrolled/render"

// Respond switches on Content Type to determine
// the response
func (n *Negotiator) Respond(w io.Writer, status int, v interface{}) {
	switch n.ContentType {
	case render.ContentJSON:
		n.Render.JSON(w, status, v)
	case render.ContentXML:
		n.Render.XML(w, status, v)
	default:
		n.Render.JSON(w, status, v)
	}
}
