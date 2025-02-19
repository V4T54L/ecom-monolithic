package utils

import "ecom-mono-backend/internals/app/models"

type ICrypto interface {
	Hash(value string) string
	GetTokenStr(tokenObj *models.AuthToken) (string, error)
	GetTokenObj(token string) (*models.AuthToken, error)
}
