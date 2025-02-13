package database

import "database/sql"

type PostgresConnector struct{}

func (r *PostgresConnector) Open(uri string) (*sql.DB, error) {
	return sql.Open("postgres", uri)
}
