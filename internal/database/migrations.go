package database

import (
	"log"

	"github.com/edmartt/grpc-test/internal/person/models"
)

type Migrations struct {
	DB IConnection
}

func (m Migrations) MigrateData() {
	connection, conErr := m.DB.GetConnection()

	if conErr != nil {
		log.Println("Connection Error Migrations: ", conErr.Error())
	}

	connection.AutoMigrate(&models.Person{})
}

func InitMigrations() {
	migrations := Migrations{
		DB: SQLiteDB{},
	}

	migrations.MigrateData()
}
