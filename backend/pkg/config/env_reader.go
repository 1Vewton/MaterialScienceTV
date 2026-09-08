package config

import (
	"os"
	"strconv"

	"github.com/1Vewton/MaterialScienceTV/backend/internal/database/databasetype"
)

// GetEnvString returns the string value stored in the environment
func GetEnvString(
	key string,
	defaultValue string,
) *string {
	result := os.Getenv(key)
	if result == "" {
		return &defaultValue
	}
	return &result
}

// GetEnvInteger returns the string value stored in the environment
func GetEnvInteger(
	key string,
	defaultValue int,
) (*int, error) {
	result := os.Getenv(key)
	if result == "" {
		return &defaultValue, nil
	}
	num, err := strconv.Atoi(result)
	return &num, err
}

// GetEnvDatabaseType returns the database type value stored in the environment
func GetEnvDatabaseType(
	key string,
	defaultValue databasetype.DBType,
) *databasetype.DBType {
	result := os.Getenv(key)
	if result == "" {
		return &defaultValue
	}
	num, err := strconv.Atoi(result)
	if err != nil {
		panic(err.Error())
	}
	databaseTypeEnum := databasetype.DBType(num)
	return &databaseTypeEnum
}
