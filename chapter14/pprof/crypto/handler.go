package crypto

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// GuessHandler checks if ?message=password
func GuessHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	msg := r.FormValue("message")

	// "password"
	real := []byte("$2a$10$2ovnPWuIjMx2S0HvCxP/mutzdsGhyt8rq/JqnJg/6OyC3B0APMGlK")

	if err := bcrypt.CompareHashAndPassword(real, []byte(msg)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("try again"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("you got it"))
	return
}
