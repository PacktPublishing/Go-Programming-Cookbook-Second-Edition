package database

import (
	"database/sql"
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	type args struct {
		db *sql.DB
	}
	type rets struct {
		error1 error
		error2 error
	}
	tests := []struct {
		name    string
		args    args
		rets    rets
		wantErr bool
	}{
		{"All good", args{db}, rets{nil, nil}, false},
		{"Create error", args{db}, rets{errors.New("failed"), nil}, true},
		{"Insert error", args{db}, rets{nil, errors.New("failed")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectExec("CREATE TABLE").WillReturnError(tt.rets.error1).WillReturnResult(sqlmock.NewResult(0, 0))
			mock.ExpectExec("INSERT INTO").WillReturnError(tt.rets.error2).WillReturnResult(sqlmock.NewResult(0, 0))
			if err := Create(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
