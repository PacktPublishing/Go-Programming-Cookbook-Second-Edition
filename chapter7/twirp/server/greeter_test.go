package main

import (
	"reflect"
	"testing"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter7/twirp/rpc/greeter"
	"golang.org/x/net/context"
)

func TestGreeter_Greet(t *testing.T) {
	type fields struct {
		Exclaim bool
	}
	type args struct {
		ctx context.Context
		r   *greeter.GreetRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *greeter.GreetResponse
		wantErr bool
	}{
		{"exclaim", fields{true}, args{context.Background(), &greeter.GreetRequest{Greeting: "Hello", Name: "there"}}, &greeter.GreetResponse{Response: "Hello there!"}, false},
		{"goodbye", fields{false}, args{context.Background(), &greeter.GreetRequest{Greeting: "Goodbye", Name: "there"}}, &greeter.GreetResponse{Response: "Goodbye there."}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Greeter{
				Exclaim: tt.fields.Exclaim,
			}
			got, err := g.Greet(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Greeter.Greet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Greeter.Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
