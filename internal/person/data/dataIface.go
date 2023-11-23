package data

import (
	"github.com/edmartt/grpc-test/internal/person/models"
)

var dataAccessIface IUserDataAccess

type IUserDataAccess interface {
	Create(person models.Person) string
	Read(id string) (*models.Person, error)
	Delete(person *models.Person) (*models.Person, error)
}
