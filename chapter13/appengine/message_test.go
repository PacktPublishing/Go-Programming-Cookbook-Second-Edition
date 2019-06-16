package main

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/datastore"
)

func TestController_storeMessage(t *testing.T) {
	projectID := os.Getenv("GCLOUD_DATASET_ID")
	cli, err := datastore.NewClient(context.Background(), projectID)
	if err != nil {
		t.Error(err)
	}

	type fields struct {
		store *datastore.Client
	}
	type args struct {
		ctx     context.Context
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"base-case", fields{cli}, args{context.Background(), "test"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				store: tt.fields.store,
			}
			if err := c.storeMessage(tt.args.ctx, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Controller.storeMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestController_queryMessages(t *testing.T) {
	projectID := os.Getenv("GCLOUD_DATASET_ID")
	cli, err := datastore.NewClient(context.Background(), projectID)
	if err != nil {
		t.Error(err)
	}

	type fields struct {
		store *datastore.Client
	}
	type args struct {
		ctx   context.Context
		limit int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"base-case", fields{cli}, args{context.Background(), 10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				store: tt.fields.store,
			}
			_, err := c.queryMessages(tt.args.ctx, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller.queryMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
