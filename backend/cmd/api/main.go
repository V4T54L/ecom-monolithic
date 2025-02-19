package main

import (
	"ecom-mono-backend/api"
	v1 "ecom-mono-backend/api/v1"
	"ecom-mono-backend/config"
	"ecom-mono-backend/internals/app/repository"
	"ecom-mono-backend/internals/app/services"
	"ecom-mono-backend/internals/app/utils"
	"ecom-mono-backend/internals/database"
	"fmt"
	"log"
	"os"
)

func main() {
	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		err := config.LoadConfigurationFile(".env")
		if err != nil {
			log.Fatal("Failed to load configuration : ", err)
		}
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Failed to get Config : ", err)
	}

	dbUri := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	connector := database.PostgresConnector{}
	db, err := database.NewPostgreSQLDB(dbUri, cfg.MaxIdleConns, cfg.MaxOpenConns, &connector)
	if err != nil {
		log.Fatal("error connecting to database : ", err)
	}
	defer db.Close()

	repo := repository.NewRepository(db.GetConn())

	crypto := utils.NewSessionCrypto(string(cfg.HashSecret))

	service, err := services.NewService(repo, crypto)
	if err != nil {
		log.Fatal("error initializing service : ", err)
	}

	handler, err := v1.NewHandler(service)
	if err != nil {
		log.Fatal("error initializing handler : ", err)
	}

	r := api.SetupRouter(handler, crypto.GetTokenObj)
	if err = r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("error occurred in the server : ", err)
	}
}
