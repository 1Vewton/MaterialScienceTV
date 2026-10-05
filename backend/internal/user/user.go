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
