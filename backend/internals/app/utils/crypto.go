package utils

import (
	"crypto/sha256"
	"ecom-mono-backend/internals/app/models"
	"encoding/base64"
	"errors"

	"github.com/google/uuid"
)

type sessionCrypto struct {
	tokens             map[string]*models.AuthToken
	newRandomBase64Str func() (string, error)
	hashSecret         string
}

func NewSessionCrypto(hashSecret string) ICrypto {
	return &sessionCrypto{
		make(map[string]*models.AuthToken),
		func() (string, error) {
			newUUID, err := uuid.NewRandom()
			if err != nil {
				return "", err
			}

			return base64.RawURLEncoding.EncodeToString([]byte(newUUID.String())), nil
		},
		hashSecret,
	}
}

func (c *sessionCrypto) GetTokenObj(token string) (*models.AuthToken, error) {
	val, ok := c.tokens[token]
	if !ok {
		return nil, errors.New("invalid token provided")
	}
	return val, nil
}

func (c *sessionCrypto) GetTokenStr(tokenObj *models.AuthToken) (string, error) {
	tokenStr, err := c.newRandomBase64Str()
	if err != nil {
		return "", err
	}

	c.tokens[tokenStr] = tokenObj
	return tokenStr, nil
}

func (c *sessionCrypto) Hash(value string) string {
	dataToHash := value + c.hashSecret

	hash := sha256.Sum256([]byte(dataToHash))

	return base64.StdEncoding.EncodeToString(hash[:])
}
