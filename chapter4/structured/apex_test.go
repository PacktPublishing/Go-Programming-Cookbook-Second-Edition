package structured

import (
	"testing"

	"github.com/apex/log"
	"github.com/apex/log/handlers/discard"
)

func TestThrowError(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ThrowError(); (err != nil) != tt.wantErr {
				t.Errorf("ThrowError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCustomHandler_HandleLog(t *testing.T) {
	type fields struct {
		id      string
		handler log.Handler
	}
	type args struct {
		e *log.Entry
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"base-case", fields{"test", discard.New()}, args{log.NewEntry(&log.Logger{})}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &CustomHandler{
				id:      tt.fields.id,
				handler: tt.fields.handler,
			}
			if err := h.HandleLog(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("CustomHandler.HandleLog() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApex(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"base-case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Apex()
		})
	}
}
