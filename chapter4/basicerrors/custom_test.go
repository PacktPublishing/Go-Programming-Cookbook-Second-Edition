package basicerrors

import "testing"

func TestCustomError_Error(t *testing.T) {
	tests := []struct {
		name string
		c    CustomError
		want string
	}{
		{"base-case", CustomError{"test"}, "there was an error; test was the result"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Error(); got != tt.want {
				t.Errorf("CustomError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeFunc(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SomeFunc(); (err != nil) != tt.wantErr {
				t.Errorf("SomeFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
