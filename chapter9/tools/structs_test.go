package tools

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_c_example3(t *testing.T) {
	type fields struct {
		Branch bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"no branch", fields{false}, false},
		{"branch", fields{true}, true},
	}
	for _, tt := range tests {
		Convey(tt.name, t, func() {
			c := &c{
				Branch: tt.fields.Branch,
			}
			So((c.example3() != nil), ShouldEqual, tt.wantErr)
		})
	}
}
