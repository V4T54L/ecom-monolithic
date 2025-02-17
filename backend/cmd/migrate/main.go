package main

import (
	"context"
	"database/sql"
	"ecom-mono-backend/config"
	"ecom-mono-backend/internals/database"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	migrationUpPath   = "./database/migrations/up.%s.sql"
	migrationDownPath = "./database/migrations/down.%s.sql"
	seedPath          = "./database/seeds/%s.sql"
)

var (
	// Names should be saved such that versionNames[version_no-1] = version_name
	versionNames = []string{
		"001_init",
	}
)

func getCurrentVersion(db *sql.Tx) (int, error) {
	var version int
	err := db.QueryRow("SELECT version_no FROM versions WHERE id = 1").Scan(&version)
	if err != nil {
		return -1, fmt.Errorf("failed to query current version: %w", err)
	}
	return version, nil
}

func executeFile(db *sql.Tx, filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file %s: %w", filePath, err)
	}

	log.Printf("Executing SQL statement from %s:\n", filePath)
	_, err = db.Exec(string(content))
	if err != nil {
		return fmt.Errorf("failed to execute SQL statement from %s: %w", filePath, err)
	}

	return nil
}

func getFilesToExcute(currentVersion, targetVersion int) []string {
	log.Println("Current : ", currentVersion)
	log.Println("Target : ", targetVersion)
	if targetVersion < 0 {
		targetVersion = -1
	}
	if currentVersion < 0 {
		currentVersion = -1
	}

	isUp := targetVersion-currentVersion > 0

	filesToExecute := []string{}
	if isUp {
		for _, versionName := range versionNames[currentVersion+1 : targetVersion] {
			filesToExecute = append(filesToExecute, fmt.Sprintf(migrationUpPath, versionName))
			filesToExecute = append(filesToExecute, fmt.Sprintf(seedPath, versionName))
		}
	} else {
		files := []string{}
		for _, versionName := range versionNames[targetVersion+1 : currentVersion] {
			files = append(files, fmt.Sprintf(migrationDownPath, versionName))
		}
		for i := len(files) - 1; i >= 0; i-- {
			filesToExecute = append(filesToExecute, files[i])
		}
	}
	return filesToExecute
}

func main() {
	targetVersion := flag.Int("target", -2, "target migration version")
	flag.Parse()

	if *targetVersion == -2 {
		log.Fatal("target flag not provided")
	}

	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		err := config.LoadConfigurationFile(".env")
		if err != nil {
			log.Fatal("Failed to load configuration : ", err)
		}
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	connector := database.PostgresConnector{}

	dbUri := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	db, err := database.NewPostgreSQLDB(dbUri, cfg.MaxIdleConns, cfg.MaxOpenConns, &connector)
	if err != nil {
		log.Fatal("error connecting to database : ", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	txn, err := db.GetConn().BeginTx(ctx, nil)
	if err != nil {
		log.Fatal("[-] error starting first transaction: ", err)
	}
	version, err := getCurrentVersion(txn)
	if err != nil {
		log.Println("[-] Error fetching current version : ", err)
		log.Println("[-] Continuing to migrate")
	}

	files := getFilesToExcute(version, *targetVersion)
	log.Println("Files : ", files)

	txn, err = db.GetConn().BeginTx(ctx, nil)
	if err != nil {
		log.Fatal("[-] error starting second transaction: ", err)
	}
	defer txn.Rollback()

	for _, filepath := range files {
		err = executeFile(txn, filepath)
		if err != nil {
			log.Fatal("[-] error migrating file: ", err)
		}
	}

	err = txn.Commit()
	if err != nil {
		log.Fatal("[-] error committing the migration: ", err)
	}
	log.Println("[+] Database migrated successfully!")
}
