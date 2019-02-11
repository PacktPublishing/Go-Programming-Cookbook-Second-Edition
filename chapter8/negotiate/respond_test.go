package negotiate

import (
	"bytes"
	"testing"

	"github.com/unrolled/render"
)

func TestNegotiator_Respond(t *testing.T) {
	type args struct {
		status int
		v      interface{}
	}
	tests := []struct {
		name  string
		n     *Negotiator
		args  args
		wantW string
	}{
		{"json", &Negotiator{ContentType: render.ContentJSON, Render: render.New()}, args{200, map[string]string{"test": "test"}}, `{"test":"test"}`},
		{"xml", &Negotiator{ContentType: render.ContentXML, Render: render.New()}, args{200, map[string]string{"test": "test"}}, ``},
		{"none", &Negotiator{Render: render.New()}, args{200, map[string]string{"test": "test"}}, `{"test":"test"}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.n.Respond(w, tt.args.status, tt.args.v)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Negotiator.Respond() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
