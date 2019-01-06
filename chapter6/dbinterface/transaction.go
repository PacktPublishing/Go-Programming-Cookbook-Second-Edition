package dbinterface

import "database/sql"

// DB is an interface that is satisfied
// by an sql.DB or an sql.Transaction
type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// Transaction can do anything a Query can do
// plus Commit, Rollback, or Stmt
type Transaction interface {
	DB
	Commit() error
	Rollback() error
	Stmt(stmt *sql.Stmt) *sql.Stmt
}
