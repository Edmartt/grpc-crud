package data

import (
	"github.com/edmartt/grpc-test/internal/database"
	"github.com/edmartt/grpc-test/internal/person/models"
)

var dBConnectionObject database.IConnection

type UserDataAccess struct {
	DB     database.IConnection
	person models.Person
}

func init() {
	dBConnectionObject = database.SQLiteDB{}
}

func (u UserDataAccess) Create(person models.Person) string {
	connection, err := dBConnectionObject.GetConnection()

	if err != nil {
		return "DB ERROR: " + err.Error()
	}

	connection.Create(&person)

	return "created"
}

func (u UserDataAccess) Read(id string) (*models.Person, error) {
	connection, err := dBConnectionObject.GetConnection()

	person := models.Person{}
	if err != nil {
		return nil, err
	}

	connection.First(&person, "id = ?", id)

	return &person, nil
}
