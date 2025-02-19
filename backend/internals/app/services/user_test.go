package services

import (
	"context"
	"ecom-mono-backend/internals/app/models"
	"ecom-mono-backend/internals/app/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newUserService(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name          string
		repo          repository.IUserRepo
		hashFunc      func(value string) string
		getTokenFunc  func(tokenObj *models.AuthToken) (string, error)
		expectedError error
	}{
		{
			name:          "success",
			repo:          &mockUserRepo{},
			hashFunc:      func(value string) string { return "" },
			getTokenFunc:  func(tokenObj *models.AuthToken) (string, error) { return "", nil },
			expectedError: nil,
		},
		{
			name:          "failure - nil repo",
			repo:          nil,
			hashFunc:      func(value string) string { return "" },
			getTokenFunc:  func(tokenObj *models.AuthToken) (string, error) { return "", nil },
			expectedError: errors.New("nil value received for IUserRepo"),
		},
		{
			name:          "failure - nil hashFunc",
			repo:          &mockUserRepo{},
			hashFunc:      nil,
			getTokenFunc:  func(tokenObj *models.AuthToken) (string, error) { return "", nil },
			expectedError: errors.New("nil value received for hashFunc"),
		},
		{
			name:          "failure - nil getTokenFunc",
			repo:          &mockUserRepo{},
			hashFunc:      func(value string) string { return "" },
			getTokenFunc:  nil,
			expectedError: errors.New("nil value received for getTokenFunc"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			userService, err := newUserService(tc.repo, tc.hashFunc, tc.getTokenFunc)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err)
				assert.Nil(t, userService)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, userService)
			}
		})
	}
}

func Test_userService_Signup(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name          string
		setupMock     func(repo *mockUserRepo, crypto *mockCrypto)
		expectedID    int
		expectedError error
		nameArg       string
		usernameArg   string
		emailArg      string
		passwordArg   string
		roleArg       string
	}{
		{
			name: "success",
			setupMock: func(repo *mockUserRepo, crypto *mockCrypto) {
				repo.createFn = func(ctx context.Context, name, username, email, hashedPassword, role string) (int, error) {
					assert.Equal(t, "Test Name", name)
					assert.Equal(t, "testuser", username)
					assert.Equal(t, "test@example.com", email)
					assert.Equal(t, "hashed_password", hashedPassword) // Verify the password is being hashed correctly.
					assert.Equal(t, "user", role)
					return 123, nil
				}
				crypto.hashFunc = func(value string) string {
					return "hashed_" + value
				}
			},
			expectedID:    123,
			expectedError: nil,
			nameArg:       "Test Name",
			usernameArg:   "testuser",
			emailArg:      "test@example.com",
			passwordArg:   "password",
			roleArg:       "user",
		},
		{
			name: "failure - repo create error",
			setupMock: func(repo *mockUserRepo, crypto *mockCrypto) {
				repo.createFn = func(ctx context.Context, name, username, email, hashedPassword, role string) (int, error) {
					return 0, errors.New("database error")
				}
				crypto.hashFunc = func(value string) string {
					return "hashed_" + value
				}
			},
			expectedID:    0,
			expectedError: errors.New("database error"),
			nameArg:       "Test Name",
			usernameArg:   "testuser",
			emailArg:      "test@example.com",
			passwordArg:   "password",
			roleArg:       "user",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			repo := &mockUserRepo{}
			crypto := &mockCrypto{}
			if tc.setupMock != nil {
				tc.setupMock(repo, crypto)
			}

			userService := &userService{repo: repo, hashFunc: crypto.Hash, getTokenFunc: crypto.GetTokenStr}
			id, err := userService.Signup(context.Background(), tc.nameArg, tc.usernameArg, tc.emailArg, tc.passwordArg, tc.roleArg)

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

func Test_userService_Login(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name                string
		setupMock           func(repo *mockUserRepo, crypto *mockCrypto)
		usernameArg         string
		passwordArg         string
		expectedToken       string
		expectedUserDetails *models.AuthToken
		expectedError       error
	}{
		{
			name: "success",
			setupMock: func(repo *mockUserRepo, crypto *mockCrypto) {
				repo.getByCredsFn = func(ctx context.Context, username, hashedPassword string) (int, string, string, string, string, string, error) {
					assert.Equal(t, "testuser", username)
					assert.Equal(t, "hashed_password", hashedPassword)
					return 123, "Test Name", "test@example.com", "user", "2023-11-20", "2023-11-20", nil
				}
				crypto.hashFunc = func(value string) string {
					return "hashed_" + value
				}
				crypto.getTokenStrFn = func(tokenObj *models.AuthToken) (string, error) {
					assert.Equal(t, 123, tokenObj.ID)
					assert.Equal(t, "Test Name", tokenObj.Name)
					assert.Equal(t, "testuser", tokenObj.Username)
					assert.Equal(t, "test@example.com", tokenObj.Email)
					assert.Equal(t, "user", tokenObj.Role)
					return "generated_token", nil
				}
			},
			usernameArg:         "testuser",
			passwordArg:         "password",
			expectedToken:       "generated_token",
			expectedUserDetails: &models.AuthToken{ID: 123, Name: "Test Name", Username: "testuser", Email: "test@example.com", Role: "user", CreatedAt: "2023-11-20", UpdatedAt: "2023-11-20"},
			expectedError:       nil,
		},
		{
			name: "failure - repo get by creds error",
			setupMock: func(repo *mockUserRepo, crypto *mockCrypto) {
				repo.getByCredsFn = func(ctx context.Context, username, hashedPassword string) (int, string, string, string, string, string, error) {
					return 0, "", "", "", "", "", errors.New("database error")
				}
				crypto.hashFunc = func(value string) string {
					return "hashed_" + value
				}
			},
			usernameArg:         "testuser",
			passwordArg:         "password",
			expectedToken:       "",
			expectedUserDetails: nil,
			expectedError:       errors.New("database error"),
		},
		{
			name: "failure - token generation error",
			setupMock: func(repo *mockUserRepo, crypto *mockCrypto) {
				repo.getByCredsFn = func(ctx context.Context, username, hashedPassword string) (int, string, string, string, string, string, error) {
					return 123, "Test Name", "test@example.com", "user", "2023-11-20", "2023-11-20", nil
				}
				crypto.hashFunc = func(value string) string {
					return "hashed_" + value
				}
				crypto.getTokenStrFn = func(tokenObj *models.AuthToken) (string, error) {
					return "", errors.New("token generation failed")
				}
			},
			usernameArg:         "testuser",
			passwordArg:         "password",
			expectedToken:       "",
			expectedUserDetails: nil,
			expectedError:       errors.New("token generation failed"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			repo := &mockUserRepo{}
			crypto := &mockCrypto{}
			if tc.setupMock != nil {
				tc.setupMock(repo, crypto)
			}

			userService := &userService{repo: repo, hashFunc: crypto.Hash, getTokenFunc: crypto.GetTokenStr}
			token, userDetails, err := userService.Login(context.Background(), tc.usernameArg, tc.passwordArg)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedToken, token)
				assert.Equal(t, tc.expectedUserDetails, userDetails)
			}
		})
	}
}
