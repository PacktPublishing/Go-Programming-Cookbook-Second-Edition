package consensus

import (
	"reflect"
	"testing"

	"github.com/hashicorp/raft"
)

func TestNewFSM(t *testing.T) {
	tests := []struct {
		name string
		want *FSM
	}{
		{"base-case", &FSM{state: first}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFSM(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFSM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFSM_Apply(t *testing.T) {
	type fields struct {
		state state
	}
	type args struct {
		r *raft.Log
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{"base-case", fields{first}, args{&raft.Log{}}, "first"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FSM{
				state: tt.fields.state,
			}
			if got := f.Apply(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FSM.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}
