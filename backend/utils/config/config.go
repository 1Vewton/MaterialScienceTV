package config

import (
	"github.com/1Vewton/MaterialScienceTV/backend/database"
)

// config struct include the basic settings.
type config struct {
	databaseURL  *string
	databaseType *database.DBType
}

// GetDatabaseURL method returns the database url to connect
func (cfg *config) GetDatabaseURL() string {
	return SetConfigString(
		"DATABASE_URL",
		":memory:",
		&cfg.databaseURL,
	)
}

// GetDatabaseType method returns the database url to connect
func (cfg *config) GetDatabaseType() database.DBType {
	return SetConfigDBType(
		"DATABASE_URL",
		database.Sqlite,
		&cfg.databaseType,
	)
}
