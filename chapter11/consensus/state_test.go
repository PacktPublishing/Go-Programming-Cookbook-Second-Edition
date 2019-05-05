package consensus

import "testing"

func Test_state_CanTransition(t *testing.T) {
	type args struct {
		next state
	}
	tests := []struct {
		name string
		s    state
		args args
		want bool
	}{
		{"first to second", first, args{second}, true},
		{"first to third", first, args{third}, true},
		{"second to first", second, args{first}, false},
		{"second to third", second, args{third}, true},
		{"third to first", third, args{first}, true},
		{"third to second", third, args{second}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.CanTransition(tt.args.next); got != tt.want {
				t.Errorf("state.CanTransition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_state_Transition(t *testing.T) {
	type args struct {
		next state
	}
	tests := []struct {
		name string
		s    state
		args args
	}{
		{"base-case", first, args{second}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Transition(tt.args.next)
		})
	}
}
