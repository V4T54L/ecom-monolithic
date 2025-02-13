package database

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPostgreSQLDB(t *testing.T) {
	t.Run("returns nil and error when Open fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer db.Close()

		mockConnector := &MockDBConnector{MockDB: db, MockError: errors.New("connection error")}

		database, err := NewPostgreSQLDB("test_uri", 10, 10, mockConnector)
		require.Error(t, err)
		require.Nil(t, database)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("returns DB instance on successful connection", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer db.Close()

		mockConnector := &MockDBConnector{MockDB: db, MockError: nil}

		database, err := NewPostgreSQLDB("test_uri", 10, 10, mockConnector)
		require.NoError(t, err)
		require.NotNil(t, database)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("returns error on Ping failure", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		require.NoError(t, err)
		defer db.Close()

		mockConnector := &MockDBConnector{MockDB: db, MockError: nil}

		mock.ExpectPing().WillReturnError(errors.New("ping error"))

		database, err := NewPostgreSQLDB("test_uri", 10, 10, mockConnector)
		require.Error(t, err)
		require.Nil(t, database)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("returns error when database connection is nil", func(t *testing.T) {
		db := &postgreSQLDB{conn: nil}
		err := db.Close()
		assert.Equal(t, errors.New("database connection is nil"), err)
	})

	t.Run("Close succeeds when connection is valid", func(t *testing.T) {
		dbMock, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer dbMock.Close()
		mock.ExpectClose()

		db := &postgreSQLDB{conn: dbMock}
		assert.NoError(t, db.Close())
	})
}

func TestPing(t *testing.T) {
	dbMock, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	require.NoError(t, err)
	defer dbMock.Close()

	mock.ExpectPing().WillReturnError(nil)

	db := &postgreSQLDB{conn: dbMock}
	assert.NoError(t, db.Ping())

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetConn(t *testing.T) {
	dbMock, _, err := sqlmock.New()
	require.NoError(t, err)
	defer dbMock.Close()

	db := &postgreSQLDB{conn: dbMock}
	assert.NotNil(t, db.GetConn())
}

type MockDBConnector struct {
	MockDB    *sql.DB
	MockError error
}

func (m *MockDBConnector) Open(uri string) (*sql.DB, error) {
	return m.MockDB, m.MockError
}
