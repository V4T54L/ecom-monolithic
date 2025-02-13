package database

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

type postgreSQLDB struct {
	conn *sql.DB
}

func NewPostgreSQLDB(dbUri string, maxIdleConns, maxOpenConns int, connector DBConnector) (DB, error) {
	if connector == nil {
		return nil, errors.New("nil value provided for DBConnector")
	}
	dbConn, err := connector.Open(dbUri)
	if err != nil {
		return nil, err
	}

	dbConn.SetMaxOpenConns(maxOpenConns)
	dbConn.SetMaxIdleConns(maxIdleConns)

	if err = dbConn.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL database")
	return &postgreSQLDB{conn: dbConn}, nil
}

func (p *postgreSQLDB) GetConn() *sql.DB {
	return p.conn
}

func (p *postgreSQLDB) Ping() error {
	return p.conn.Ping()
}

func (p *postgreSQLDB) Close() error {
	if p.conn == nil {
		return errors.New("database connection is nil")
	}
	return p.conn.Close()
}

var _ DB = (*postgreSQLDB)(nil)
