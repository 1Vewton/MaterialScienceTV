package config

import (
	"github.com/1Vewton/MaterialScienceTV/backend/database"
)

// SetConfigString set the config or return the value in String field directly
func SetConfigString(
	key string,
	defaultValue string,
	field **string,
) string {
	if field == nil {
		panic("You cannot give a nil pointer to field in SetConfigString!")
	}
	if *field == nil {
		*field = GetEnvString(
			key,
			defaultValue,
		)
	}
	return **field
}

// SetConfigDBType set the config or return the value in DBType field directly
func SetConfigDBType(
	key string,
	defaultValue database.DBType,
	field **database.DBType,
) database.DBType {
	if field == nil {
		panic("You cannot give a nil pointer to field in SetConfigDBType!")
	}
	if *field == nil {
		*field = GetEnvDatabaseType(
			key,
			defaultValue,
		)
	}
	return **field
}
