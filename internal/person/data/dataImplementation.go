package data

import (
	"github.com/edmartt/grpc-test/internal/database"
	"github.com/edmartt/grpc-test/internal/person/models"
	"github.com/edmartt/grpc-test/internal/utils"
)

var dBConnectionObject database.IConnection

type UserDataAccess struct {
	DB      database.IConnection
	ZLogger utils.ILogger
	Person  models.Person
}

func init() {
	dBConnectionObject = database.SQLiteDB{
		ZLogger: utils.ZeroLoggerAdapter{},
	}
}

func (u UserDataAccess) Create(person models.Person) string {
	connection, err := dBConnectionObject.GetConnection()

	if err != nil {
		u.ZLogger.Error("db connection error")
		return "DB ERROR: " + err.Error()
	}

	connection.Create(&person)
	u.ZLogger.Info("person created")

	return "created"
}

func (u UserDataAccess) Read(id string) (*models.Person, error) {
	connection, err := dBConnectionObject.GetConnection()

	person := u.Person

	if err != nil {
		u.ZLogger.Error("db connection error")
		return nil, err
	}

	connection.First(&person, "id = ?", id)
	u.ZLogger.Info("person data fetched")

	return &person, nil
}

func (u UserDataAccess) Delete(person *models.Person) (*models.Person, error) {
	connection, err := dBConnectionObject.GetConnection()
	personModel := &models.Person{}

	if err != nil {
		u.ZLogger.Error("db connection error")
		return nil, err
	}

	connection.Delete(person)

	u.ZLogger.Info("person data deleted")

	return personModel, nil
}
