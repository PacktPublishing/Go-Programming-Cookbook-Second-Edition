package channels

import "testing"
import "time"

func TestSender(t *testing.T) {
	type args struct {
		ch   chan string
		done chan bool
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{make(chan string, 10), make(chan bool)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go Sender(tt.args.ch, tt.args.done)
			time.Sleep(100 * time.Millisecond)
			tt.args.done <- true
		})
	}
}
