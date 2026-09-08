package database

import (
	"testing"

	"github.com/1Vewton/MaterialScienceTV/backend/database/databasetype"
)

// Test database initialization
func TestDataBaseInitialization(t *testing.T) {
	err := InitDataBase(
		databasetype.Sqlite,
		":memory:",
	)
	if err != nil {
		t.Error(err.Error())
	}
}
