package utils

import (
	"ecom-mono-backend/internals/app/models"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockNewRandomBase64Str func() (string, error)

func TestNewSessionCrypto(t *testing.T) {
	t.Parallel()
	hashSecret := "testSecret"
	crypto := NewSessionCrypto(hashSecret)
	assert.NotNil(t, crypto, "NewSessionCrypto should not return nil")

	sc, ok := crypto.(*sessionCrypto)
	assert.True(t, ok, "NewSessionCrypto should return a *sessionCrypto")
	assert.Equal(t, hashSecret, sc.hashSecret, "HashSecret should be set correctly")
}

func TestSessionCrypto_GetTokenObj(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name          string
		setup         func(c *sessionCrypto)
		token         string
		expectedToken *models.AuthToken
		expectedError error
	}{
		{
			name: "success - token exists",
			setup: func(c *sessionCrypto) {
				tokenStr, err := c.newRandomBase64Str()
				require.NoError(t, err)
				authToken := &models.AuthToken{ID: 1, Name: "Test User", Username: "testuser"}
				c.tokens[tokenStr] = authToken
			},
			token:         "existingToken",
			expectedToken: &models.AuthToken{ID: 1, Name: "Test User", Username: "testuser"},
			expectedError: nil,
		},
		{
			name: "failure - token does not exist",
			setup: func(c *sessionCrypto) {
			},
			token:         "nonExistentToken",
			expectedToken: nil,
			expectedError: errors.New("invalid token provided"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			crypto := &sessionCrypto{
				tokens: make(map[string]*models.AuthToken),
				newRandomBase64Str: func() (string, error) {
					return "existingToken", nil
				},
				hashSecret: "testSecret",
			}

			if tc.setup != nil {
				tc.setup(crypto)
			}

			token := tc.token
			if tc.name == "success - token exists" {
				for k := range crypto.tokens {
					token = k
					break
				}

			}

			tokenObj, err := crypto.GetTokenObj(token)

			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error(), "Error should match")
				assert.Nil(t, tokenObj, "Token object should be nil on error")
			} else {
				assert.NoError(t, err, "Should not return an error")
				assert.NotNil(t, tokenObj, "Token object should not be nil")
				assert.True(t, reflect.DeepEqual(tokenObj, tc.expectedToken), "Token objects should be equal")
			}
		})
	}
}

func TestSessionCrypto_GetTokenStr(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name               string
		setupNewRandom     func() (string, error)
		tokenObj           *models.AuthToken
		expectedTokenStr   string
		expectedError      error
		expectedTokensSize int // Check the number of tokens after GetTokenStr
	}{
		{
			name: "success",
			setupNewRandom: func() (string, error) {
				return "generatedToken", nil
			},
			tokenObj:           &models.AuthToken{ID: 1, Name: "Test User", Username: "testuser"},
			expectedTokenStr:   "generatedToken",
			expectedError:      nil,
			expectedTokensSize: 1,
		},
		{
			name: "failure - newRandomBase64Str fails",
			setupNewRandom: func() (string, error) {
				return "", errors.New("random string error")
			},
			tokenObj:           &models.AuthToken{ID: 1, Name: "Test User", Username: "testuser"},
			expectedTokenStr:   "",
			expectedError:      errors.New("random string error"),
			expectedTokensSize: 0, // The token shouldn't be added if there's an error
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			crypto := &sessionCrypto{
				tokens:             make(map[string]*models.AuthToken),
				newRandomBase64Str: tc.setupNewRandom,
				hashSecret:         "testSecret",
			}

			tokenStr, err := crypto.GetTokenStr(tc.tokenObj)

			assert.Equal(t, tc.expectedTokenStr, tokenStr, "Token string should match")
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error(), "Error should match")
			} else {
				assert.NoError(t, err, "Should not return an error")
				assert.Contains(t, crypto.tokens, tokenStr, "Token should be added to the map")
			}

			assert.Equal(t, tc.expectedTokensSize, len(crypto.tokens), "Tokens map size should match")
		})
	}
}

// Tests for Hash
func TestSessionCrypto_Hash(t *testing.T) {
	crypto := &sessionCrypto{
		newRandomBase64Str: nil,
		hashSecret:         "secret",
	}

	hash := crypto.Hash("value")

	assert.Greater(t, len(hash), 1)
}
