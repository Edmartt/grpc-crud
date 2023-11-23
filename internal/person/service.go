package person

import (
	"context"

	"github.com/edmartt/grpc-test/internal/person/data"
	"github.com/edmartt/grpc-test/internal/person/models"
	pb "github.com/edmartt/grpc-test/internal/person/protos/bin"
	"github.com/google/uuid"
)

var dataAccess data.IUserDataAccess

func init() {
	dataAccess = data.UserDataAccess{}
}

type Service struct {
	pb.UnimplementedPersonServiceServer
}

func (s *Service) Create(ctx context.Context, person *pb.Person) (*pb.PersonCreatedResponse, error) {
	id := uuid.New().String()

	newPerson := &models.Person{
		ID:        id,
		FirstName: person.Name,
		LastName:  person.LastName,
		Email:     person.Email,
	}

	status := dataAccess.Create(*newPerson)

	return &pb.PersonCreatedResponse{Id: newPerson.ID, Response: status}, nil
}

func (s *Service) Get(ctx context.Context, request *pb.PersonRequest) (*pb.Person, error) {
	id := request.Id
	dbResponse, err := dataAccess.Read(id)

	if err != nil {
		return nil, err
	}

	personResponse := &pb.Person{
		Id:       dbResponse.ID,
		Name:     dbResponse.FirstName,
		LastName: dbResponse.LastName,
		Email:    dbResponse.Email,
	}

	return personResponse, nil
}