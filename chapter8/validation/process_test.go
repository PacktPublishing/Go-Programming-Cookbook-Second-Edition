package validation

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_Process(t *testing.T) {
	type fields struct {
		ValidatePayload func(p *Payload) error
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
		{"base-case", fields{func(p *Payload) error { return nil }}, args{httptest.NewRecorder(), httptest.NewRequest("POST", "/", bytes.NewBufferString(`{"name": "123", "age": 3}`))}},
		{"Bad method", fields{func(p *Payload) error { return nil }}, args{httptest.NewRecorder(), httptest.NewRequest("GET", "/", bytes.NewBufferString(`{"name": "123", "age": 3}`))}},
		{"Bad payload", fields{func(p *Payload) error { return nil }}, args{httptest.NewRecorder(), httptest.NewRequest("POST", "/", bytes.NewBufferString(`{r}`))}},
		{"verror", fields{func(p *Payload) error { return Verror{errors.New("test")} }}, args{httptest.NewRecorder(), httptest.NewRequest("POST", "/", bytes.NewBufferString(`{"name": "123", "age": 3}`))}},
		{"regular error", fields{func(p *Payload) error { return errors.New("test") }}, args{httptest.NewRecorder(), httptest.NewRequest("POST", "/", bytes.NewBufferString(`{"name": "123", "age": 3}`))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				ValidatePayload: tt.fields.ValidatePayload,
			}
			c.Process(tt.args.w, tt.args.r)
		})
	}
}
