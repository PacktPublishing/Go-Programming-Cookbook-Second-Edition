package database

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func TestQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name      string
		args      args
		queryTime interface{}
		queryErr  bool
		wantErr   bool
	}{
		{"base-case", args{db}, time.Now(), false, false},
		{"bad row", args{db}, 123, false, true},
		{"query error", args{db}, time.Now(), true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"name", "created"}).
				AddRow("test", tt.queryTime))

			if tt.queryErr {
				m.WillReturnError(errors.New("failed"))
			}
			if err := Query(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
