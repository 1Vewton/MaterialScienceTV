package user

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/1Vewton/MaterialScienceTV/backend/internal/database/redismanager"
)

// TestUploadAndFetch tests the uploading and fetching of tmpUser
func TestUploadAndFetch(
	t *testing.T,
) {
	t.Parallel()
	ctx := t.Context()
	testClient := redismanager.NewRedisConfig(
		"localhost:6379",
		"",
	).ToClient()
	id := uuid.NewString()
	testUser := NewTmpUser(
		id,
		"test",
		"A114514",
		"@gmail.com",
	)
	err := testUser.UploadToRedis(
		ctx,
		testClient,
		id,
		5*time.Minute,
	)
	if err != nil {
		t.Error(err)
	}
	_, err = GetTmpUser(
		ctx,
		testClient,
		id,
	)
	err = testClient.Close()
	if err != nil {
		t.Error(err)
	}
}
