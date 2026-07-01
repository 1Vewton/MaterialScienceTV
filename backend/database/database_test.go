package database

import (
	"testing"
)

// Test database initialization
func TestDataBaseInitialization(t *testing.T) {
	err := InitDataBase()
	if err != nil {
		t.Error(err.Error())
	}
}
