package v1

import (
	"ecom-mono-backend/internals/app/services"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	GetUserDetails(ctx *gin.Context)
	LoginHandler(ctx *gin.Context)
	SignupHandler(ctx *gin.Context)
}

type Handler struct {
	User IUserHandler
}

func NewHandler(service *services.Service) (*Handler, error) {
	userHandler, err := newUserHandler(service.User)
	if err != nil {
		return nil, err
	}

	return &Handler{User: userHandler}, nil
}
