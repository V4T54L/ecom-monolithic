package api

import (
	v1 "ecom-mono-backend/api/v1"
	"ecom-mono-backend/internals/app/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *v1.Handler, tokenValidatorFunc func(token string) (*models.AuthToken, error)) *gin.Engine {
	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/login", h.User.LoginHandler)
		v1.POST("/auth/signup", h.User.SignupHandler)

		userRoutes := v1.Group("")
		userRoutes.Use(AuthMiddleware(tokenValidatorFunc))

		userRoutes.GET("/profile", h.User.GetUserDetails)
		// v1.GET("/users", userHandler.GetAllUsers)
		// v1.GET("/users/:id", userHandler.GetUserByID)
		// v1.DELETE("/users/:id", userHandler.DeleteUser)

		// v1.Use(middleware.Auth()) // middleware for authentication

		// v1.POST("/todos", todoHandler.CreateTodo)
		// v1.GET("/todos", todoHandler.GetAllTodos)
		// v1.GET("/todos/:id", todoHandler.GetTodoByID)
		// v1.DELETE("/todos/:id", todoHandler.DeleteTodo)
		// v1.PUT("/todos/:id", todoHandler.UpdateTodo)
	}

	return r
}
