package user

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

// TmpUser struct provides the user datamodel for temporarily storing in redis
type TmpUser struct {
	UserID       string `redis:"user_id"`
	UserName     string `redis:"user_name"`
	Password     string `redis:"password"`
	Email        string `redis:"email"`
	RegisteredAt int    `redis:"registered_at"`
}

// UploadToRedis uploads tmp user to redis.
// Use processed token instead of raw token
func (user *TmpUser) UploadToRedis(
	ctx context.Context,
	client *redis.Client,
	registerToken string,
	lifeTime time.Duration,
) error {
	encodedJson, err := json.Marshal(
		user,
	)
	if err != nil {
		return err
	}
	stringJson := string(encodedJson)
	_, err = client.Set(
		ctx,
		registerToken,
		stringJson,
		lifeTime,
	).Result()
	return err
}

// GetTmpUser gets the tmpuser from the redis for registration
func GetTmpUser(
	ctx context.Context,
	client *redis.Client,
	registerToken string,
) (*TmpUser, error) {
	stringResult, err := client.Get(
		ctx,
		registerToken,
	).Result()
	if err != nil {
		return nil, err
	}
	result := TmpUser{}
	err = json.Unmarshal(
		[]byte(stringResult),
		&result,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
