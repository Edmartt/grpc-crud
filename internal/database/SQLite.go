package database

import (
	"github.com/edmartt/grpc-test/internal/utils"
	db "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteDB struct {
	ZLogger utils.ILogger
}

func NewSQLiteDB() *SQLiteDB {
	sqliteDB := SQLiteDB{
		ZLogger: utils.NewZeroLoggerAdapter(),
	}
	return &sqliteDB
}

func (sqlite SQLiteDB) GetConnection() (*gorm.DB, error) {
	sqlite.ZLogger.Info("establishing sqlite connection")
	connection, conError := gorm.Open(db.Open("data.db"))

	if conError != nil {
		sqlite.ZLogger.Error(conError.Error())
		return nil, conError
	}
	return connection, nil
}
