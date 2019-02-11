package controllers

import (
	"encoding/json"
	"net/http"
)

// SetValue modifies the underlying storage of the controller object
func (c *Controller) SetValue(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	value := r.FormValue("value")
	c.storage.Put(value)
	w.WriteHeader(http.StatusOK)
	p := Payload{Value: value}
	if payload, err := json.Marshal(p); err == nil {
		w.Write(payload)
	}

}
