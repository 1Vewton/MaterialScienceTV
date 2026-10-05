package redismanager

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// RedisClient defines the client for redis
var RedisClient *redis.Client

// Close closes the redis client
func Close(
	client *redis.Client,
) error {
	err := client.Close()
	return err
}

// NewToken creates token for certain module
func NewToken(
	moduleName string,
) string {
	id := uuid.NewString()
	return fmt.Sprintf(
		"%s:%s",
		moduleName,
		id,
	)
}
