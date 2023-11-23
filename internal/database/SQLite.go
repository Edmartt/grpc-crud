package database

import (
	"log"

	db "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteDB struct {
}

func (sqlite SQLiteDB) GetConnection() (*gorm.DB, error) {
	connection, conError := gorm.Open(db.Open("data.db"))

	if conError != nil {
		log.Fatal("error direct db: ", conError)
		return nil, conError
	}
	return connection, nil
}
