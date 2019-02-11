package middleware

import (
	"net/http"
)

// Handler is very basic
func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}
