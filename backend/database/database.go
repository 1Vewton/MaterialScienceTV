package database

import (
	"errors"

	"github.com/1Vewton/MaterialScienceTV/backend/database/databasetype"
	"github.com/1Vewton/MaterialScienceTV/backend/utils/config"
	"github.com/1Vewton/MaterialScienceTV/backend/utils/logger"
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
func InitDataBase() error {
	var err error
	switch config.Config.GetDatabaseType() {
	case databasetype.Sqlite:
		databaseLogger.Info("Use Sqlite")
		DataBase, err = gorm.Open(
			sqlite.Open(config.Config.GetDatabaseURL()),
			&gorm.Config{},
		)
	case databasetype.MySQL:
		databaseLogger.Info("Use MySQL")
		DataBase, err = gorm.Open(
			mysql.Open(config.Config.GetDatabaseURL()),
			&gorm.Config{},
		)
	case databasetype.PostgreSQL:
		databaseLogger.Info("Use Postgres")
		DataBase, err = gorm.Open(
			postgres.Open(config.Config.GetDatabaseURL()),
			&gorm.Config{},
		)
	default:
		return errors.New("Database type not support")
	}
	if err != nil {
		databaseLogger.Error(err.Error())
		return err
	}
	return nil
}
