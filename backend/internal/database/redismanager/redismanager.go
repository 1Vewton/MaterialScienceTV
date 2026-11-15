package redismanager

import (
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
