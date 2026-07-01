package database

import (
	"time"

	"gorm.io/gorm"
)

// User struct provides the user datamodel for database
type User struct {
	gorm.Model
	UserID       string
	UserName     string
	Password     string
	Email        string
	RegisteredAt time.Time
}
