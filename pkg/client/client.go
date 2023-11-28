package client

import (
	"context"

	pb "github.com/edmartt/grpc-test/internal/person/protos/bin"
)

func CreatePerson(person *pb.Person) (string, error) {
	connection, err := grpcConnector()

	if err != nil {
		return "", err
	}
	serviceClient := pb.NewPersonServiceClient(connection)

	serverResponse, err := serviceClient.Create(context.Background(), &pb.Person{
		Id:        person.Id,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Email:     person.Email,
	})

	if err != nil {
		return "", err
	}

	response := serverResponse.Id
	serverResponse.Response = response

	return serverResponse.Response, nil
}

func ReadPerson(request *pb.GetPersonRequest) (*pb.GetPersonResponse, error) {
	connection, err := grpcConnector()

	if err != nil {
		return nil, err
	}

	serviceClient := pb.NewPersonServiceClient(connection)

	serverResponse, err := serviceClient.Get(context.Background(), request)

	if err != nil {
		return nil, err
	}

	return serverResponse, nil
}

func DeletePerson(request *pb.DeletePersonRequest) (*pb.DeletePersonResponse, error) {
	connection, err := grpcConnector()

	if err != nil {
		return nil, err
	}

	serviceClient := pb.NewPersonServiceClient(connection)

	serverResponse, err := serviceClient.Delete(context.Background(), request)

	if err != nil {
		return nil, err
	}

	return serverResponse, nil
}
