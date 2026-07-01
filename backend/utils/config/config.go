package config

import (
	"github.com/1Vewton/MaterialScienceTV/backend/database/databasetype"
)

// config struct include the basic settings.
type config struct {
	databaseURL  *string
	databaseType *databasetype.DBType
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
func (cfg *config) GetDatabaseType() databasetype.DBType {
	return SetConfigDBType(
		"DATABASE_TYPE",
		databasetype.Sqlite,
		&cfg.databaseType,
	)
}
