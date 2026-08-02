package config

import (
	"github.com/1Vewton/MaterialScienceTV/backend/database/databasetype"
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
	defaultValue databasetype.DBType,
	field **databasetype.DBType,
) databasetype.DBType {
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

// SetConfigInteger set the config or return value in integer field directly
func SetConfigInteger(
	key string,
	defaultValue int,
	field **int,
) (int, error) {
	if field == nil {
		panic("You cannot give a nil pointer to field in SetConfigDBType!")
	}
	if *field == nil {
		num, err := GetEnvInteger(
			key,
			defaultValue,
		)
		if err != nil {
			return defaultValue, err
		}
		*field = num
	}
	return **field, nil
}
