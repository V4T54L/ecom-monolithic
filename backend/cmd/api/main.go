package main

import (
	"ecom-mono-backend/internals/config"
	"ecom-mono-backend/internals/controllers"
	"ecom-mono-backend/internals/database"
	"ecom-mono-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvs()
	database.ConnectDB()

}

func main() {
	router := gin.Default()

	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.Login)
	router.GET("/user/profile", middlewares.AuthMiddleware, controllers.GetUserProfile)
	router.Static("/fe", "./public")
	router.Run()
}
