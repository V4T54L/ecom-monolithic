package services

import (
	"context"
	"ecom-mono-backend/internals/app/models"
	"ecom-mono-backend/internals/app/repository"
	"ecom-mono-backend/internals/app/utils"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockUserRepo struct {
	createFn     func(ctx context.Context, name, username, email, hashedPassword, role string) (int, error)
	getByCredsFn func(ctx context.Context, username, hashedPassword string) (id int, name, email, role, created_at, updated_at string, err error)
}

func (m *mockUserRepo) Create(ctx context.Context, name, username, email, hashedPassword, role string) (int, error) {
	if m.createFn != nil {
		return m.createFn(ctx, name, username, email, hashedPassword, role)
	}
	return 0, errors.New("Create not implemented in mock")
}

func (m *mockUserRepo) GetByCreds(ctx context.Context, username, hashedPassword string) (id int, name, email, role, created_at, updated_at string, err error) {
	if m.getByCredsFn != nil {
		return m.getByCredsFn(ctx, username, hashedPassword)
	}
	return 0, "", "", "", "", "", errors.New("GetByCreds not implemented in mock")
}

type mockCrypto struct {
	hashFunc      func(value string) string
	getTokenStrFn func(tokenObj *models.AuthToken) (string, error)
}

func (m *mockCrypto) Hash(value string) string {
	if m.hashFunc != nil {
		return m.hashFunc(value)
	}
	return "hashed_" + value
}

func (m *mockCrypto) GetTokenObj(_ string) (*models.AuthToken, error) {
	return nil, nil
}

func (m *mockCrypto) GetTokenStr(tokenObj *models.AuthToken) (string, error) {
	if m.getTokenStrFn != nil {
		return m.getTokenStrFn(tokenObj)
	}
	return fmt.Sprintf("token_%d", tokenObj.ID), nil
}

func TestNewService(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name          string
		repo          *repository.Repository
		crypto        utils.ICrypto
		expectedError error
	}{
		{
			name:          "success",
			repo:          &repository.Repository{User: &mockUserRepo{}},
			crypto:        &mockCrypto{},
			expectedError: nil,
		},
		{
			name:          "failure - nil repo",
			repo:          nil,
			crypto:        &mockCrypto{},
			expectedError: errors.New("nil value provided for Repository"),
		},
		{
			name:          "failure - nil crypto",
			repo:          &repository.Repository{User: &mockUserRepo{}},
			crypto:        nil,
			expectedError: errors.New("nil value provided for ICrypto"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			service, err := NewService(tc.repo, tc.crypto)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err)
				assert.Nil(t, service)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, service)
				assert.NotNil(t, service.User)
			}
		})
	}
}
