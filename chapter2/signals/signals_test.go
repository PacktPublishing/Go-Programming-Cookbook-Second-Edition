package main

import (
	"os"
	"syscall"
	"testing"
)

func TestCatchSig(t *testing.T) {
	signals := make(chan os.Signal)
	done := make(chan bool)
	type args struct {
		ch     chan os.Signal
		done   chan bool
		signal os.Signal
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{signals, done, syscall.SIGINT}},
		{"sigterm", args{signals, done, syscall.SIGTERM}},
		{"sigkill", args{signals, done, syscall.SIGKILL}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go CatchSig(tt.args.ch, tt.args.done)
			signals <- tt.args.signal
			<-done
		})
	}
}
