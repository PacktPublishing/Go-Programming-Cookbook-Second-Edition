package dbinterface

import (
	"errors"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T) {

	tests := []struct {
		name     string
		ExecErr1 bool
		ExecErr2 bool
		wantErr  bool
	}{
		{"base-case", false, false, false},
		{"1st exec error", true, false, true},
		{"2nd exec error", false, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockDB := NewMockDB(ctrl)
			m := mockDB.EXPECT().Exec(gomock.Any())
			if tt.ExecErr1 {
				m.Return(nil, errors.New("failed"))
			} else {
				m.Return(nil, nil)
			}
			m2 := mockDB.EXPECT().Exec(gomock.Any()).AnyTimes()
			if tt.ExecErr2 {
				m2.Return(nil, errors.New("failed"))
			} else {
				m2.Return(nil, nil)
			}
			if err := Create(mockDB); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
