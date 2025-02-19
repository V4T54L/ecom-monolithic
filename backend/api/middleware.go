package api

import (
	"ecom-mono-backend/internals/app/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenValidatorFunc func(token string) (*models.AuthToken, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("ServerToken")
		if err != nil {
			authHeader := ctx.GetHeader("Authorization")
			if authHeader != "" {
				parts := strings.Split(authHeader, " ")
				if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
					token = parts[1]
				}
			}
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing token"})
			return
		}

		userDetails, err := tokenValidatorFunc(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
			return
		}

		ctx.Set("user", userDetails)

		ctx.Next()
	}
}
