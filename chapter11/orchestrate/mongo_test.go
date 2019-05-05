package orchestrate

import (
	"testing"

	mgo "gopkg.in/mgo.v2"
)

func TestConnectAndQuery(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			session, err := mgo.Dial("localhost")
			if err != nil {
				t.Error(err)
			}
			if err := ConnectAndQuery(session); (err != nil) != tt.wantErr {
				t.Errorf("ConnectAndQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
