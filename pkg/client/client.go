package client

import (
	"context"
	"fmt"

	pb "github.com/edmartt/grpc-test/internal/person/protos/bin"
)

func CreatePerson(person *pb.Person) string {
	connection := grpcConnector()
	serviceClient := pb.NewPersonServiceClient(connection)

	serverResponse, err := serviceClient.Create(context.Background(), &pb.Person{
		Id:       person.Id,
		Name:     person.Name,
		LastName: person.LastName,
		Email:    person.Email,
	})

	if err != nil {
		return err.Error()
	}

	response := fmt.Sprintf("Person with %s id created", serverResponse.Id)
	serverResponse.Response = response

	return serverResponse.Response
}

func ReadPerson(request *pb.PersonRequest) *pb.Person {
	connection := grpcConnector()

	serviceClient := pb.NewPersonServiceClient(connection)

	serverResponse, err := serviceClient.Get(context.Background(), request)

	if err != nil {
		return nil
	}

	return serverResponse
}
