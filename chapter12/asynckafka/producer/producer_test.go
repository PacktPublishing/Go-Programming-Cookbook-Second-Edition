package main

import (
	"testing"

	"github.com/golang/mock/gomock"

	sarama "github.com/Shopify/sarama"
)

func TestProcessResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAsyncProducer := NewMockAsyncProducer(ctrl)

	type args struct {
		producer sarama.AsyncProducer
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{mockAsyncProducer}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			succ := make(chan *sarama.ProducerMessage)
			errs := make(chan *sarama.ProducerError)
			mockAsyncProducer.EXPECT().Successes().AnyTimes().Return(succ)
			mockAsyncProducer.EXPECT().Errors().AnyTimes().Return(errs)
			go ProcessResponse(tt.args.producer)
			succ <- &sarama.ProducerMessage{}
			errs <- &sarama.ProducerError{}
		})
	}
}
