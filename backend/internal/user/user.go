package user

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// User struct provides the user datamodel for database
type User struct {
	gorm.Model
	UserID       string `gorm:"primaryKey"`
	UserName     string
	Password     string
	Email        string
	RegisteredAt time.Time
}

// Equals tests if two users are the same
func (user *User) Equals(
	another *User,
) bool {
	return user.UserID == another.UserID &&
		user.UserName == another.UserName &&
		user.Password == another.Password &&
		user.Email == another.Email
}

// NewUser creates new user
func NewUser(
	userName string,
	password string,
	email string,
	id string,
) *User {
	return &User{
		UserID:       id,
		UserName:     userName,
		Password:     password,
		Email:        email,
		RegisteredAt: time.Now(),
	}
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
