package tools

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_example(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"base-case"},
	}
	for _, tt := range tests {
		Convey(tt.name, t, func() {
			res := example()
			So(res, ShouldBeNil)
		})
	}
}

func Test_example2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"base-case"},
	}
	for _, tt := range tests {
		Convey(tt.name, t, func() {
			res := example2()
			So(res, ShouldBeGreaterThanOrEqualTo, 1)
		})
	}
}
