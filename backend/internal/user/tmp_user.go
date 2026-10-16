package user

import (
	"context"
	"encoding/json"
	"time"

	"github.com/1Vewton/MaterialScienceTV/backend/pkg/security"
	"github.com/redis/go-redis/v9"
)

// TmpUser struct provides the user datamodel for temporarily storing in redis
type TmpUser struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	HashValue    string `json:"hash_value"`
	Salt         string `json:"salt"`
	Email        string `json:"email"`
	RegisteredAt int64  `json:"registered_at"`
}

// NewTmpUser creates new tmp user
func NewTmpUser(
	userID string,
	userName string,
	password string,
	email string,
) (*TmpUser, error) {
	saltValue, err := security.GenerateSalt()
	if err != nil {
		return nil, err
	}
	hash := security.ToHash(
		password,
		saltValue,
	)
	return &TmpUser{
		UserID:       userID,
		UserName:     userName,
		Salt:         saltValue,
		HashValue:    hash,
		Email:        email,
		RegisteredAt: time.Now().UnixNano(),
	}, nil
}

// UploadToRedis uploads tmp user to redis.
// Use processed token instead of raw token
func (user *TmpUser) UploadToRedis(
	ctx context.Context,
	client *redis.Client,
	registerToken string,
	lifeTime time.Duration,
) error {
	encodedJSON, err := json.Marshal(
		user,
	)
	if err != nil {
		return err
	}
	stringJSON := string(encodedJSON)
	_, err = client.Set(
		ctx,
		registerToken,
		stringJSON,
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

// Equals checks if two users are the same
func (user *TmpUser) Equals(
	another *TmpUser,
) bool {
	return user.Email == another.Email &&
		user.RegisteredAt == another.RegisteredAt
}

// ToUserData converts tmp user to user data that stores in the database
func (user *TmpUser) ToUserData() *User {
	newUserData := &User{
		UserID:    user.UserID,
		UserName:  user.UserName,
		HashValue: user.HashValue,
		Salt:      user.Salt,
		Email:     user.Email,
		RegisteredAt: time.Unix(
			0,
			user.RegisteredAt,
		),
	}
	return newUserData
}
