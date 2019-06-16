package main

import (
	"context"
	"net/http"
	"os"
	"testing"

	"net/http/httptest"

	"cloud.google.com/go/datastore"
)

func TestController_handle(t *testing.T) {
	projectID := os.Getenv("GCLOUD_DATASET_ID")
	cli, err := datastore.NewClient(context.Background(), projectID)
	if err != nil {
		t.Error(err)
	}
	type fields struct {
		store *datastore.Client
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"base-case", fields{cli}, args{httptest.NewRecorder(), httptest.NewRequest("POST", "/", nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				store: tt.fields.store,
			}
			c.handle(tt.args.w, tt.args.r)
		})
	}
}
