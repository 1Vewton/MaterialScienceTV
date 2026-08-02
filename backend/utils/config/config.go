package config

import (
	"github.com/1Vewton/MaterialScienceTV/backend/database/databasetype"
)

// config struct include the basic settings.
type config struct {
	serverPort    *string
	databaseURL   *string
	redisURL      *string
	redisPassword *string
	databaseType  *databasetype.DBType
}

// GetDatabaseURL method returns the database url to connect
func (cfg *config) GetDatabaseURL() string {
	return SetConfigString(
		"DATABASE_URL",
		":memory:",
		&cfg.databaseURL,
	)
}

// GetDatabaseType method returns the database url to connect
func (cfg *config) GetDatabaseType() databasetype.DBType {
	return SetConfigDBType(
		"DATABASE_TYPE",
		databasetype.Sqlite,
		&cfg.databaseType,
	)
}

// GetServerPort gets the port the server is running
func (cfg *config) GetServerPort() string {
	return SetConfigString(
		"SERVER_PORT",
		"8080",
		&cfg.serverPort,
	)
}

// GetRedisURL gets the url for the Redis
func (cfg *config) GetRedisURL() string {
	return SetConfigString(
		"REDIS_URL",
		"localhost:6379",
		&cfg.redisURL,
	)
}

// GetRedisPassword gets the url for the Redis
func (cfg *config) GetRedisPassword() string {
	return SetConfigString(
		"REDIS_PASSWORD",
		"",
		&cfg.redisPassword,
	)
}
