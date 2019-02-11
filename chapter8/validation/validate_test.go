package validation

import "testing"

func TestValidatePayload(t *testing.T) {
	type args struct {
		p *Payload
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"no name", args{&Payload{}}, true},
		{"no age", args{&Payload{Name: "test"}}, true},
		{"all good", args{&Payload{Name: "test", Age: 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePayload(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePayload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
