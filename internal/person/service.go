package person

import (
	"context"

	"github.com/edmartt/grpc-test/internal/person/data"
	"github.com/edmartt/grpc-test/internal/person/models"
	pb "github.com/edmartt/grpc-test/internal/person/protos/bin"
	"github.com/google/uuid"
)

var DataAccess data.IUserDataAccess

func init() {
	DataAccess = data.UserDataAccess{}
}

type Service struct {
	pb.UnimplementedPersonServiceServer
}

func (s *Service) Create(ctx context.Context, person *pb.Person) (*pb.CreatePersonResponse, error) {
	id := uuid.New().String()

	newPerson := &models.Person{
		ID:        id,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Email:     person.Email,
	}

	status := DataAccess.Create(*newPerson)

	return &pb.CreatePersonResponse{Id: newPerson.ID, Response: status}, nil
}

func (s *Service) Get(ctx context.Context, request *pb.GetPersonRequest) (*pb.GetPersonResponse, error) {
	id := request.Id
	dbResponse, err := DataAccess.Read(id)

	if err != nil {
		return nil, err
	}

	person := &pb.Person{
		Id:        dbResponse.ID,
		FirstName: dbResponse.FirstName,
		LastName:  dbResponse.LastName,
		Email:     dbResponse.Email,
	}

	return &pb.GetPersonResponse{
		Person: person,
	}, nil
}

func (s *Service) Delete(ctx context.Context, request *pb.DeletePersonRequest) (*pb.DeletePersonResponse, error) {
	id := request.Id
	queryPerson, err := DataAccess.Read(id)

	if err != nil {
		return nil, err
	}

	if queryPerson.ID == "" {
		return &pb.DeletePersonResponse{
			Id:     id,
			Status: "not found",
		}, nil
	}

	DataAccess.Delete(queryPerson)

	delPersonResponse := &pb.DeletePersonResponse{
		Id:     queryPerson.ID,
		Status: "deleted",
	}

	return delPersonResponse, nil

}
