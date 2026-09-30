package redismanager

import (
	"testing"
)

// TestConnectAndClose tests connection and closing
func TestConnectAndClose(
	t *testing.T,
) {
	t.Parallel()
	testClient := NewRedisConfig(
		"localhost:6379",
		"",
	).WithDialTimeout(
		5,
	).WithReadTimeout(
		5,
	).WithWriteTimeout(
		5,
	).WithMaxRetries(
		3,
	).WithMaxRetryBackoff(
		5,
	).WithMinRetryBackoff(
		3,
	).ToClient()
	err := testClient.Close()
	if err != nil {
		t.Error(err)
	}
}
