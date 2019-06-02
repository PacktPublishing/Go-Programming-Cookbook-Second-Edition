package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

	sarama "github.com/Shopify/sarama"
)

func TestKafkaController_Handler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAsyncProducer := NewMockAsyncProducer(ctrl)

	type fields struct {
		producer sarama.AsyncProducer
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
		{"no message", fields{mockAsyncProducer}, args{httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &KafkaController{
				producer: tt.fields.producer,
			}
			c.Handler(tt.args.w, tt.args.r)
		})
	}
}
