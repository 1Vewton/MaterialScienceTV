package config

import (
	"testing"
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

// Test the env setting
func TestEnvSettingString(t *testing.T) {
	var tSetting *config = &config{}
	t.Setenv("DATABASE_URL", "test.db")
	res := tSetting.GetDatabaseURL()
	if res != "test.db" {
		t.Errorf("Expected test.db, got %s", res)
	}
}
