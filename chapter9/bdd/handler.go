package bdd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandlerRequest will be json decoded
// into by Handler
type HandlerRequest struct {
	Name string `json:"name"`
}

// Handler takes a request and renders a response
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	dec := json.NewDecoder(r.Body)
	var req HandlerRequest
	if err := dec.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("BDD testing %s", req.Name)))
}
