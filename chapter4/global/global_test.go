package global

import (
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Init(); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetLog(t *testing.T) {
	type args struct {
		l *logrus.Logger
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{logrus.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLog(tt.args.l)
		})
	}
}

func TestWithField(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want *logrus.Entry
	}{
		{"base-case", args{"test", "case"}, log.WithField("test", "case")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithField(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDebug(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.args...)
		})
	}
}
