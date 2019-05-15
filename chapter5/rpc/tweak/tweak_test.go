package tweak

import "testing"

func TestStringTweaker_Tweak(t *testing.T) {
	type args struct {
		args *Args
		resp string
	}
	tests := []struct {
		name    string
		s       StringTweaker
		args    args
		wantErr bool
	}{
		{"no tweaks", StringTweaker{}, args{&Args{String: "test"}, "test"}, false},
		{"reverse", StringTweaker{}, args{&Args{String: "test", Reverse: true}, "tset"}, false},
		{"upper", StringTweaker{}, args{&Args{String: "test", ToUpper: true}, "TEST"}, false},
		{"both", StringTweaker{}, args{&Args{String: "test", Reverse: true, ToUpper: true}, "TSET"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StringTweaker{}
			var result string
			if err := s.Tweak(tt.args.args, &result); (err != nil) != tt.wantErr {
				t.Errorf("StringTweaker.Tweak() error = %v, wantErr %v", err, tt.wantErr)
			}

			if result != tt.args.resp {
				t.Errorf("wanted: %s, got %s", tt.args.resp, result)
			}

		})
	}
}
