package negotiate

import (
	"encoding/xml"
	"net/http"
)

// Payload defines it's layout in xml and json
type Payload struct {
	XMLName xml.Name `xml:"payload" json:"-"`
	Status  string   `xml:"status" json:"status"`
}

// Handler gets a negotiator using the request,
// then renders a Payload
func Handler(w http.ResponseWriter, r *http.Request) {
	n := GetNegotiator(r)

	n.Respond(w, http.StatusOK, &Payload{Status: "Successful!"})
}
