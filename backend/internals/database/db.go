package database

import "database/sql"

type DB interface {
	Ping() error
	Close() error
}

type DBConnector interface {
	Open(uri string) (*sql.DB, error)
}
