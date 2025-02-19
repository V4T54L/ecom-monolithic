package v1

import (
	"ecom-mono-backend/internals/app/models"
	"ecom-mono-backend/internals/app/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service services.IUserService
}

func newUserHandler(userService services.IUserService) (IUserHandler, error) {
	if userService == nil {
		return nil, errors.New("nil value provided for IUserService")
	}
	return &userHandler{userService}, nil
}

func (u *userHandler) LoginHandler(ctx *gin.Context) {
	payload := models.LoginPayload{}
	err := ctx.ShouldBindBodyWithJSON(&payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	token, userDetails, err := u.service.Login(ctx, payload.Username, payload.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// TODO: Update the cookie domain if it doesn't work
	ctx.SetCookie("ServerToken", token, 3600, "", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"user": userDetails})
}

func (u *userHandler) SignupHandler(ctx *gin.Context) {
	payload := models.SignupPayload{}
	err := ctx.ShouldBindBodyWithJSON(&payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, err = u.service.Signup(ctx, payload.Name, payload.Username, payload.Email, payload.Password, "user")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user created"})
}

func (u *userHandler) GetUserDetails(ctx *gin.Context) {
	userDetails, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}

	user, ok := userDetails.(*models.AuthToken)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type in context"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}
