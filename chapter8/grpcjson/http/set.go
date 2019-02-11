package main

import (
	"encoding/json"
	"net/http"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter8/grpcjson/internal"
	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter8/grpcjson/keyvalue"
	"github.com/apex/log"
)

// Controller holds an internal KeyValueObject
type Controller struct {
	*internal.KeyValue
}

// SetHandler wraps or GRPC Set
func (c *Controller) SetHandler(w http.ResponseWriter, r *http.Request) {
	var kv keyvalue.SetKeyValueRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&kv); err != nil {
		log.Errorf("failed to decode: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	gresp, err := c.Set(r.Context(), &kv)
	if err != nil {
		log.Errorf("failed to set: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(gresp)
	if err != nil {
		log.Errorf("failed to marshal: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
