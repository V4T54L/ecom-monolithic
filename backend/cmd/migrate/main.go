package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

const (
	migrationUpPath   = "./database/migrations/up.%s.sql"
	migrationDownPath = "./database/migrations/down.%s.sql"
	seedPath          = "./database/seeds/%s.sql"
)

var (
	// Names should be saved such that versionNames[version_no-1] = version_name
	versionNames = []string{
		"001_init.sql",
	}
)

func getCurrentVersion(db *sql.DB) (int, error) {
	var version int
	err := db.QueryRow("SELECT version_no FROM versions WHERE id = 1").Scan(&version)
	if err != nil {
		return -1, fmt.Errorf("failed to query current version: %w", err)
	}
	return version, nil
}

func executeFile(db *sql.DB, filePath string) error {
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
	if targetVersion < 0 {
		targetVersion = -1
	}
	if currentVersion < 0 {
		currentVersion = -1
	}

	isUp := targetVersion-currentVersion < 0

	filesToExecute := []string{}
	if isUp {
		for _, versionName := range versionNames[currentVersion:] {
			filesToExecute = append(filesToExecute, fmt.Sprintf(migrationUpPath, versionName))
			filesToExecute = append(filesToExecute, fmt.Sprintf(seedPath, versionName))
		}
	} else {
		for _, versionName := range versionNames[currentVersion:] {
			filesToExecute = append(filesToExecute, fmt.Sprintf(migrationDownPath, versionName))
		}
	}
	return filesToExecute
}

func main() {
	// connector := database.PostgresConnector{}
	// database.NewPostgreSQLDB()
}
