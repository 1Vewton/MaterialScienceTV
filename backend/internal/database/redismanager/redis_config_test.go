package redismanager

import (
	"testing"
	"time"
)

// TestSetConfig tests the config setting
func TestSetConfig(t *testing.T) {
	t.Parallel()
	newRedisConfig := NewRedisConfig(
		"localhost:6379",
		"",
	).WithDialTimeout(
		5,
	)
	if newRedisConfig.RedisOptions.DialTimeout != time.Duration(5)*time.Second {
		t.Errorf(
			"expected %d, got %d",
			newRedisConfig.RedisOptions.DialTimeout,
			time.Duration(5)*time.Second,
		)
	}
}
