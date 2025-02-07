package main

import (
	"ecom-mono-backend/internals/config"
	"ecom-mono-backend/internals/database"
	"ecom-mono-backend/models"
)

func init() {
	config.LoadEnvs()
	database.ConnectDB()

}

func main() {

	database.DB.AutoMigrate(&models.User{})
}
