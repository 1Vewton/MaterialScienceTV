package user

import (
	"context"
	"errors"
	"time"

	"github.com/1Vewton/MaterialScienceTV/backend/pkg/security"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// User struct provides the user datamodel for database
type User struct {
	gorm.Model
	UserID       string `gorm:"primaryKey"`
	UserName     string
	HashValue    string
	Salt         string
	Email        string
	RegisteredAt time.Time
}

// Equals tests if two users are the same
func (user *User) Equals(
	another *User,
) bool {
	return user.UserID == another.UserID &&
		user.UserName == another.UserName &&
		user.Email == another.Email
}

// NewUser creates new user
func NewUser(
	userName string,
	password string,
	email string,
	id string,
) (*User, error) {
	saltValue, err := security.GenerateSalt()
	if err != nil {
		return nil, err
	}
	hash := security.ToHash(
		password,
		saltValue,
	)
	return &User{
		UserID:       id,
		UserName:     userName,
		HashValue:    hash,
		Salt:         saltValue,
		Email:        email,
		RegisteredAt: time.Now(),
	}, nil
}

// AddUser adds new user
func AddUser(
	ctx context.Context,
	db *gorm.DB,
	user *User,
) error {
	err := gorm.G[User](db).Create(
		ctx,
		user,
	)
	return err
}

// HasUser checks if a user exists
func HasUser(
	ctx context.Context,
	db *gorm.DB,
	userID string,
) (bool, error) {
	users, err := gorm.G[*User](db).Where(
		&User{
			UserID: userID,
		},
	).Find(ctx)
	if err != nil {
		return false, err
	}
	if len(users) == 0 {
		return false, nil
	}
	return true, nil
}

// HasUserName checks if user name exists
func HasUserName(
	ctx context.Context,
	db *gorm.DB,
	userName string,
) (bool, error) {
	users, err := gorm.G[*User](db).Where(
		&User{
			UserName: userName,
		},
	).Find(ctx)
	if err != nil {
		return false, err
	}
	if len(users) == 0 {
		return false, nil
	}
	return true, nil
}

// GetUser gets the user
func GetUser(
	ctx context.Context,
	db *gorm.DB,
	userID string,
) (*User, error) {
	user, err := gorm.G[*User](db).Where(
		&User{
			UserID: userID,
		},
	).First(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UploadLoginToken uploads login token
func UploadLoginToken(
	ctx context.Context,
	client *redis.Client,
	loginToken string,
	userID string,
	expireDuration time.Duration,
) error {
	_, err := client.Set(
		ctx,
		loginToken,
		userID,
		expireDuration,
	).Result()
	return err
}

// VerifyLoginToken checks if login token exists
func VerifyLoginToken(
	ctx context.Context,
	client *redis.Client,
	db *gorm.DB,
	loginToken string,
) (*User, bool, error) {
	userID, err := client.Get(
		ctx,
		loginToken,
	).Result()
	if err != nil {
		return nil, false, err
	}
	user, err := gorm.G[*User](db).Where(
		&User{
			UserID: userID,
		},
	).First(ctx)
	if err != nil {
		return user, false, err
	}
	result := true
	if user == nil {
		result = false
	}
	return user, result, nil
}

// Login manages the login
func Login(
	ctx context.Context,
	db *gorm.DB,
	userName *string,
	email *string,
	password string,
) (*User, bool, error) {
	var mainUser *User
	if userName == nil && email == nil {
		return nil, false, errors.New(
			"you must fill one of the fields in userName or email",
		)
	} else if userName != nil {
		user, err := gorm.G[*User](db).Where(
			&User{
				UserName: *userName,
			},
		).First(ctx)
		if err != nil {
			return nil, false, err
		}
		if user == nil {
			return nil, false, nil
		}
		mainUser = user
	} else {
		user, err := gorm.G[*User](db).Where(
			&User{
				Email: *email,
			},
		).First(ctx)
		if err != nil {
			return nil, false, err
		}
		if user == nil {
			return nil, false, nil
		}
		mainUser = user
	}
	isLoggedIn := security.ComparePassword(
		password,
		mainUser.Salt,
		mainUser.HashValue,
	)
	return mainUser, isLoggedIn, nil
}

// DeleteUser deletes the user
func DeleteUser(
	ctx context.Context,
	db *gorm.DB,
	userID string,
) error {
	_, err := gorm.G[*User](db).Where(
		&User{
			UserID: userID,
		},
	).Delete(ctx)
	return err
}
