package database

import "gorm.io/gorm"

type IConnection interface {
	GetConnection() (*gorm.DB, error)
}
