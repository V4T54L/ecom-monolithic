package repository

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRepository(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectPing()

	repo := NewRepository(db)
	assert.NotNil(t, repo)
	assert.NotNil(t, repo.User)
}

func TestRepository_Create(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name          string
		setupMock     func(mock sqlmock.Sqlmock)
		expectedID    int
		expectedError error
	}{
		{
			name: "success",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (name, username, email, password, role) VALUES ($1,$2,$3,$4,$5) RETURNING id;")).
					WithArgs("Test Name", "testuser", "test@example.com", "hashedpassword", "user").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(123))
			},
			expectedID:    123,
			expectedError: nil,
		},
		{
			name: "failure - query error",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (name, username, email, password, role) VALUES ($1,$2,$3,$4,$5) RETURNING id;")).
					WithArgs("Test Name", "testuser", "test@example.com", "hashedpassword", "user").
					WillReturnError(errors.New("database error"))
			},
			expectedID:    0,
			expectedError: errors.New("database error"),
		},
		{
			name: "timeout",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (name, username, email, password, role) VALUES ($1,$2,$3,$4,$5) RETURNING id;")).
					WithArgs("Test Name", "testuser", "test@example.com", "hashedpassword", "user").
					WillDelayFor(QueryDuration + time.Millisecond).
					WillReturnRows(sqlmock.NewRows([]string{"id"}))
			},
			expectedID:    0,
			expectedError: errors.New("canceling query due to user request"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()
			if tc.setupMock != nil {
				tc.setupMock(mock)
			}

			repo := NewRepository(db)
			id, err := repo.User.Create(context.Background(), "Test Name", "testuser", "test@example.com", "hashedpassword", "user")

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedID, id)
			}
		})
	}
}

func TestRepository_GetByCreds(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name              string
		setupMock         func(mock sqlmock.Sqlmock)
		expectedID        int
		expectedName      string
		expectedEmail     string
		expectedRole      string
		expectedCreatedAt string
		expectedUpdatedAt string
		expectedError     error
	}{
		{
			name: "success",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, role, created_at, updated_at FROM users WHERE username=$1 AND password=$2;")).
					WithArgs("testuser", "hashedpassword").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role", "created_at", "updated_at"}).
						AddRow(123, "Test Name", "test@example.com", "user", "2023-11-20 10:00:00", "2023-11-20 10:00:00"))
			},
			expectedID:        123,
			expectedName:      "Test Name",
			expectedEmail:     "test@example.com",
			expectedRole:      "user",
			expectedCreatedAt: "2023-11-20 10:00:00",
			expectedUpdatedAt: "2023-11-20 10:00:00",
			expectedError:     nil,
		},
		{
			name: "failure - query error",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, role, created_at, updated_at FROM users WHERE username=$1 AND password=$2;")).
					WithArgs("testuser", "hashedpassword").
					WillReturnError(errors.New("database error"))
			},
			expectedID:        0,
			expectedName:      "",
			expectedEmail:     "",
			expectedRole:      "",
			expectedCreatedAt: "",
			expectedUpdatedAt: "",
			expectedError:     errors.New("database error"),
		},
		{
			name: "failure - no rows",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, role, created_at, updated_at FROM users WHERE username=$1 AND password=$2;")).
					WithArgs("testuser", "hashedpassword").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role", "created_at", "updated_at"}))
			},
			expectedID:        0,
			expectedName:      "",
			expectedEmail:     "",
			expectedRole:      "",
			expectedCreatedAt: "",
			expectedUpdatedAt: "",
			expectedError:     sql.ErrNoRows,
		},
		{
			name: "timeout",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, role, created_at, updated_at FROM users WHERE username=$1 AND password=$2;")).
					WithArgs("testuser", "hashedpassword").
					WillDelayFor(QueryDuration + time.Millisecond).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role", "created_at", "updated_at"}))
			},
			expectedID:        0,
			expectedName:      "",
			expectedEmail:     "",
			expectedRole:      "",
			expectedCreatedAt: "",
			expectedUpdatedAt: "",
			expectedError:     errors.New("canceling query due to user request"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			if tc.setupMock != nil {
				tc.setupMock(mock)
			}

			repo := NewRepository(db)
			id, name, email, role, createdAt, updatedAt, err := repo.User.GetByCreds(context.Background(), "testuser", "hashedpassword")

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedID, id)
				assert.Equal(t, tc.expectedName, name)
				assert.Equal(t, tc.expectedEmail, email)
				assert.Equal(t, tc.expectedRole, role)
				assert.Equal(t, tc.expectedCreatedAt, createdAt)
				assert.Equal(t, tc.expectedUpdatedAt, updatedAt)
			}
		})
	}
}
