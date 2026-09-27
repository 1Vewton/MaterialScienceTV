package database

import (
	"errors"

	"github.com/1Vewton/MaterialScienceTV/backend/internal/database/databasetype"
	"github.com/1Vewton/MaterialScienceTV/backend/internal/user/userdata"
	"github.com/1Vewton/MaterialScienceTV/backend/pkg/logger"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DataBase connection
var DataBase *gorm.DB
var databaseLogger *logger.Logger

func init() {
	databaseLogger = logger.NewLogger(
		"DataBase",
		nil,
	)
}

// InitDataBase Initialize the database
func InitDataBase(
	dbType databasetype.DBType,
	databaseURL string,
) (*gorm.DB, error) {
	var err error
	var resDB *gorm.DB
	// Initialize the database connection
	switch dbType {
	case databasetype.Sqlite:
		databaseLogger.Info("Use Sqlite")
		resDB, err = gorm.Open(
			sqlite.Open(databaseURL),
			&gorm.Config{},
		)
	case databasetype.MySQL:
		databaseLogger.Info("Use MySQL")
		resDB, err = gorm.Open(
			mysql.Open(databaseURL),
			&gorm.Config{},
		)
	case databasetype.PostgreSQL:
		databaseLogger.Info("Use Postgres")
		resDB, err = gorm.Open(
			postgres.Open(databaseURL),
			&gorm.Config{},
		)
	default:
		return nil, errors.New("Database type not support")
	}
	if err != nil {
		databaseLogger.Error(err.Error())
		return nil, err
	}
	// Automigrate the data
	err = resDB.AutoMigrate(&userdata.User{})
	if err != nil {
		databaseLogger.Error(err.Error())
		return nil, err
	}
	return resDB, nil
}

// CloseDatabase closes the database
func CloseDatabase(
	db *gorm.DB,
) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	return err
}
