package database

import (
	"testing"

	"github.com/1Vewton/MaterialScienceTV/backend/internal/database/databasetype"
)

// Test database initialization
func TestDataBaseInitialization(t *testing.T) {
	_, err := InitDataBase(
		databasetype.Sqlite,
		":memory:",
	)
	if err != nil {
		t.Error(err.Error())
	}
}
