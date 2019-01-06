package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

// Exec takes a new connection
// creates tables, and later drops them
// and issues some queries
func Exec(db *sql.DB) error {
	// uncaught error on cleanup, but we always
	// want to cleanup
	defer db.Exec("DROP TABLE example")

	if err := Create(db); err != nil {
		return err
	}

	if err := Query(db); err != nil {
		return err
	}
	return nil
}
