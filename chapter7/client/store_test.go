package client

import "testing"

func TestController_DoOps(t *testing.T) {
	c := Setup(true, true)
	tests := []struct {
		name    string
		c       *Controller
		wantErr bool
	}{
		{"base-case", &Controller{c}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.DoOps(); (err != nil) != tt.wantErr {
				t.Errorf("Controller.DoOps() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
