package database

import (
	"github.com/edmartt/grpc-test/internal/person/models"
	"github.com/edmartt/grpc-test/internal/utils"
)

type Migrations struct {
	DB      IConnection
	ZLogger utils.ILogger
}

func (m Migrations) MigrateData() {
	m.ZLogger.Info("trying DB connection")

	connection, conErr := m.DB.GetConnection()

	if conErr != nil {
		m.ZLogger.Fatal(conErr.Error())
	}

	err := connection.AutoMigrate(&models.Person{})

	if err != nil {
		m.ZLogger.Error(err.Error())
	}

}

func InitMigrations() {

	db := NewSQLiteDB()

	migrations := Migrations{
		DB:      db,
		ZLogger: db.ZLogger,
	}

	db.ZLogger.Trace("migrations initialized")

	migrations.MigrateData()
}
