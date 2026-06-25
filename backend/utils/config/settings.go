package config

import (
	"github.com/1Vewton/MaterialScienceTV/backend/database"
)

// Settings struct include the basic settings.
type Settings struct {
	databaseURL  *string
	databaseType *database.DBType
}
