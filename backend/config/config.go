package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config represents the application configuration
type Config struct {
	ServerPort   string
	TokenSecret  []byte
	HashSecret   []byte
	DBHost       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBPort       string
	MaxOpenConns int
	MaxIdleConns int
}

// NewConfig returns a new instance of Config.
func NewConfig() (*Config, error) {
	instance := &Config{}

	serverPort := "8080" // Default value for server port
	if val, err := getStr("SERVER_PORT", &serverPort); err == nil {
		instance.ServerPort = val
	}

	tokenSecret := "secret" // Default value for token secret
	if val, err := getStr("TOKEN_SECRET", &tokenSecret); err == nil {
		instance.TokenSecret = []byte(val)
	}

	hashSecret := "secret" // Default value for hash secret
	if val, err := getStr("HASH_SECRET", &hashSecret); err == nil {
		instance.HashSecret = []byte(val)
	}

	if val, err := getStr("DB_HOST", nil); err == nil {
		instance.DBHost = val
	} else {
		return nil, err
	}

	if val, err := getStr("DB_USER", nil); err == nil {
		instance.DBUser = val
	} else {
		return nil, err
	}

	if val, err := getStr("DB_PASSWORD", nil); err == nil {
		instance.DBPassword = val
	} else {
		return nil, err
	}

	if val, err := getStr("DB_Name", nil); err == nil {
		instance.DBName = val
	} else {
		return nil, err
	}

	if val, err := getStr("DB_PORT", nil); err == nil {
		instance.DBPort = val
	} else {
		return nil, err
	}

	if val, err := getInt("MAX_IDLE_CONNS", nil); err == nil {
		instance.MaxIdleConns = val
	} else {
		return nil, err
	}

	if val, err := getInt("MAX_OPEN_CONNS", nil); err == nil {
		instance.MaxOpenConns = val
	} else {
		return nil, err
	}

	return instance, nil
}

// LoadConfigurationFile loads configuration from a .env file at the specified path.
func LoadConfigurationFile(filePath string) error {
	if err := godotenv.Load(filePath); err != nil {
		return fmt.Errorf("error loading .env file '%s': %v", filePath, err)
	}
	return nil
}

// getStr retrieves an environment variable by key; returns fallback if not found.
// retuns error if missing environment variable and fallback is nil.
func getStr(key string, fallback *string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		if fallback == nil {
			return "", fmt.Errorf("missing required environment variable: %s", key)
		} else {
			return *fallback, nil
		}
	}
	return value, nil
}

// getInt retrieves an environment variable by key; returns fallback if not found.
// retuns error if missing environment variable and fallback is nil.
func getInt(key string, fallback *int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		if fallback == nil {
			return 0, fmt.Errorf("missing required environment variable: %s", key)
		} else {
			return *fallback, nil
		}
	}

	return strconv.Atoi(value)
}
