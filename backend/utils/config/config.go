package config

import (
	"github.com/1Vewton/MaterialScienceTV/backend/database/databasetype"
)

// config struct include the basic settings.
type config struct {
	serverPort           *string
	databaseURL          *string
	redisURL             *string
	redisPassword        *string
	redisDialTimeout     *int
	redisReadTimeout     *int
	redisWriteTimeout    *int
	redisMaxRetries      *int
	redisMinRetryBackoff *int
	redisMaxRetryBackoff *int
	databaseType         *databasetype.DBType
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

// GetRedisDialTimeout gets the Dial Timeout for the Redis
func (cfg *config) GetRedisDialTimeout() (int, error) {
	return SetConfigInteger(
		"REDIS_DIAL_TIMEOUT",
		10,
		&cfg.redisDialTimeout,
	)
}

// GetRedisReadTimeout gets the Read Timeout for the Redis
func (cfg *config) GetRedisReadTimeout() (int, error) {
	return SetConfigInteger(
		"REDIS_READ_TIMEOUT",
		5,
		&cfg.redisReadTimeout,
	)
}

// GetRedisWriteTimeout gets the Write Timeout for the Redis
func (cfg *config) GetRedisWriteTimeout() (int, error) {
	return SetConfigInteger(
		"REDIS_WRITE_TIMEOUT",
		5,
		&cfg.redisWriteTimeout,
	)
}

// GetRedisMaxRetries gets the max retries for the Redis
func (cfg *config) GetRedisMaxRetries() (int, error) {
	return SetConfigInteger(
		"REDIS_MAX_RETRIES",
		5,
		&cfg.redisMaxRetries,
	)
}

// GetRedisMaxRetryBackOff gets the max retry backoff for the Redis
func (cfg *config) GetRedisMaxRetryBackOff() (int, error) {
	return SetConfigInteger(
		"REDIS_MAX_RETRY_BACKOFF",
		100,
		&cfg.redisMaxRetryBackoff,
	)
}

// GetRedisMinRetryBackOff gets the min retry backoff for the Redis
func (cfg *config) GetRedisMinRetryBackOff() (int, error) {
	return SetConfigInteger(
		"REDIS_MIN_RETRY_BACKOFF",
		10,
		&cfg.redisMinRetryBackoff,
	)
}
