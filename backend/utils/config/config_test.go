package config

import (
	"testing"

	"github.com/1Vewton/MaterialScienceTV/backend/database/databasetype"
)

// Test the env reading of string var
func TestEnvReadingString(t *testing.T) {
	var res string
	res = *GetEnvString("test", "T")
	if res != "T" {
		t.Errorf("Expected T, got %s", res)
	}
	t.Setenv("test", "KK3S")
	res = *GetEnvString("test", "T")
	if res != "KK3S" {
		t.Errorf("Expected KK3S, got %s", res)
	}
}

// Test the env setting of string var
func TestEnvSettingString(t *testing.T) {
	var tSetting *config = &config{}
	t.Setenv("DATABASE_URL", "test.db")
	res := tSetting.GetDatabaseURL()
	if res != "test.db" {
		t.Errorf("Expected test.db, got %s", res)
	}
}

func TestEnvReadingDBType(t *testing.T) {
	res := GetEnvDatabaseType(
		"DATABASE_TYPE",
		databasetype.Sqlite,
	)
	if *res != databasetype.Sqlite {
		t.Errorf(
			"Expected %d, got %d",
			databasetype.Sqlite,
			*res,
		)
	}
	t.Setenv("DATABASE_TYPE", "0")
	if *res != databasetype.Sqlite {
		t.Errorf(
			"Expected %d, got %d",
			databasetype.Sqlite,
			*res,
		)
	}
}

// Test the env setting of database type var
func TestEnvSettingDatabaseType(t *testing.T) {
	var tSetting *config = &config{}
	t.Setenv("DATABASE_TYPE", "0")
	res := tSetting.GetDatabaseType()
	if res != databasetype.Sqlite {
		t.Errorf(
			"Expected %d, got %d",
			databasetype.Sqlite,
			res,
		)
	}
}
