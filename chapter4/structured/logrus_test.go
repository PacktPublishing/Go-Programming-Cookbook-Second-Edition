package structured

import (
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogrus(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"base-case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Logrus()
		})
	}
}

func TestHook_Fire(t *testing.T) {
	type args struct {
		entry *logrus.Entry
	}
	tests := []struct {
		name    string
		hook    *Hook
		args    args
		wantErr bool
	}{
		{"base-case", &Hook{}, args{logrus.NewEntry(&logrus.Logger{})}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hook := &Hook{}
			if err := hook.Fire(tt.args.entry); (err != nil) != tt.wantErr {
				t.Errorf("Hook.Fire() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHook_Levels(t *testing.T) {
	tests := []struct {
		name string
		hook *Hook
		want []logrus.Level
	}{
		{"base-case", &Hook{}, logrus.AllLevels},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hook := &Hook{}
			if got := hook.Levels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hook.Levels() = %v, want %v", got, tt.want)
			}
		})
	}
}
