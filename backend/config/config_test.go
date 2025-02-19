package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setup(t *testing.T, setRequired, setOptional bool) []string {
	envVars := make([]string, 0, 10)
	if setRequired {
		t.Setenv("DB_HOST", "localhost")
		t.Setenv("DB_USER", "testuser")
		t.Setenv("DB_PASSWORD", "testpass")
		t.Setenv("DB_NAME", "testdb")
		t.Setenv("DB_PORT", "5432")
		t.Setenv("MAX_OPEN_CONNS", "10")
		t.Setenv("MAX_IDLE_CONNS", "5")
		envVars = append(envVars, "HOST", "USER", "PASSWORD", "DB", "PORT", "MAX_OPEN_CONNS", "MAX_IDLE_CONNS")
	}

	if setOptional {
		t.Setenv("SERVER_PORT", "8081")
		t.Setenv("TOKEN_SECRET", "test_token")
		t.Setenv("HASH_SECRET", "test_hash")
		envVars = append(envVars, "SERVER_PORT", "TOKEN_SECRET", "HASH_SECRET")
	}

	return envVars
}

func unsetEnvVars(t *testing.T, vars []string) {
	t.Helper()
	for _, v := range vars {
		err := os.Unsetenv(v)
		if err != nil {
			require.NoError(t, err, fmt.Sprintf("failed to unset environment variable %s", v))
		}
	}
}

func TestNewConfigSuccess(t *testing.T) {
	vars := setup(t, true, true)
	defer unsetEnvVars(t, vars)

	config, err := NewConfig()

	require.NoError(t, err)
	assert.Equal(t, "8081", config.ServerPort)
	assert.Equal(t, []byte("test_token"), config.TokenSecret)
	assert.Equal(t, []byte("test_hash"), config.HashSecret)
	assert.Equal(t, "localhost", config.DBHost)
	assert.Equal(t, "testuser", config.DBUser)
	assert.Equal(t, "testpass", config.DBPassword)
	assert.Equal(t, "testdb", config.DBName)
	assert.Equal(t, "5432", config.DBPort)
	assert.Equal(t, 10, config.MaxOpenConns)
	assert.Equal(t, 5, config.MaxIdleConns)
}

func TestNewConfigWithDefaults(t *testing.T) {
	vars := setup(t, true, false)
	defer unsetEnvVars(t, vars)

	config, err := NewConfig()

	require.NoError(t, err)
	assert.Equal(t, "8080", config.ServerPort)            // Default
	assert.Equal(t, []byte("secret"), config.TokenSecret) // Default
	assert.Equal(t, []byte("secret"), config.HashSecret)  // Default
	assert.Equal(t, "localhost", config.DBHost)
	assert.Equal(t, "testuser", config.DBUser)
	assert.Equal(t, "testpass", config.DBPassword)
	assert.Equal(t, "testdb", config.DBName)
	assert.Equal(t, "5432", config.DBPort)
	assert.Equal(t, 10, config.MaxOpenConns)
	assert.Equal(t, 5, config.MaxIdleConns)
}

func TestNewConfigWithErrorRequiredFields(t *testing.T) {
	config, err := NewConfig()

	sequentialTestCases := []struct {
		name  string
		key   string
		value string
	}{
		{"required HOST", "", ""},
		{"required USER", "DB_HOST", "tmp"},
		{"required PASSWORD", "DB_USER", "tmp"},
		{"required DB", "DB_PASSWORD", "tmp"},
		{"required PORT", "DB_NAME", "tmp"},
		{"required MAX_IDLE_CONNS", "DB_PORT", "tmp"},
		{"required MAX_OPEN_CONNS", "MAX_IDLE_CONNS", "12"},
	}

	vars := make([]string, 0, 10)

	for _, tc := range sequentialTestCases {
		if tc.key != "" {
			t.Setenv(tc.key, tc.value)
			vars = append(vars, tc.key)
		}

		config, err := NewConfig()

		require.Error(t, err)
		assert.Nil(t, config)
		t.Log("Err : ", err)
	}

	t.Log("Vars : ", vars)

	unsetEnvVars(t, vars)

	require.Nil(t, config)
	require.Error(t, err)
}

func TestGetStrSuccess(t *testing.T) {
	key := "SOME_KEY"
	val := "80"
	t.Setenv(key, val)
	defer unsetEnvVars(t, []string{key})

	res, err := getStr(key, nil)

	require.NoError(t, err)
	assert.Equal(t, res, val)
}

func TestGetStrFallback(t *testing.T) {
	key := "SOME_KEY"
	fallback := "15"

	res, err := getStr(key, &fallback)

	require.NoError(t, err)
	assert.Equal(t, res, fallback)
}

func TestGetStrError(t *testing.T) {
	key := "SOME_KEY"

	res, err := getStr(key, nil)

	require.Error(t, err)
	assert.Equal(t, res, "")
}

func TestGetIntSuccess(t *testing.T) {
	key := "SOME_KEY"
	val := 80
	t.Setenv(key, fmt.Sprintf("%d", val))
	defer unsetEnvVars(t, []string{key})

	res, err := getInt(key, nil)

	require.NoError(t, err)
	assert.Equal(t, res, val)
}

func TestGetIntFallback(t *testing.T) {
	key := "SOME_KEY"
	fallback := 15

	res, err := getInt(key, &fallback)

	require.NoError(t, err)
	assert.Equal(t, res, fallback)
}

func TestGetIntError(t *testing.T) {
	key := "SOME_KEY"

	res, err := getInt(key, nil)

	require.Error(t, err)
	assert.Equal(t, res, 0)
}

func TestLoadConfigurationFileError(t *testing.T) {
	err := LoadConfigurationFile("")

	require.Error(t, err)
}
