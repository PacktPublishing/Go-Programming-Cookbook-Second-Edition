package docker

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestVersionHandler(t *testing.T) {
	type args struct {
		v *VersionInfo
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{&VersionInfo{Version: "1.0", BuildDate: time.Now()}, httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := VersionHandler(tt.args.v)
			h(tt.args.w, tt.args.r)
		})
	}
}
