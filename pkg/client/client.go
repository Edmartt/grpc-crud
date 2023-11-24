package client

import (
	"context"
	"fmt"
	"log"

	pb "github.com/edmartt/grpc-test/internal/person/protos/bin"
)

func CreatePerson(person *pb.Person) string {
	connection := grpcConnector()
	serviceClient := pb.NewPersonServiceClient(connection)

	serverResponse, err := serviceClient.Create(context.Background(), &pb.Person{
		Id:        person.Id,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Email:     person.Email,
	})

	if err != nil {
		return err.Error()
	}

	response := fmt.Sprintf("Person with %s id created", serverResponse.Id)
	serverResponse.Response = response

	return serverResponse.Response
}

func ReadPerson(request *pb.GetPersonRequest) *pb.GetPersonResponse {
	connection := grpcConnector()

	serviceClient := pb.NewPersonServiceClient(connection)

	serverResponse, err := serviceClient.Get(context.Background(), request)

	if err != nil {
		return nil
	}

	return serverResponse
}

func DeletePerson(request *pb.DeletePersonRequest) *pb.DeletePersonResponse {
	connection := grpcConnector()

	serviceClient := pb.NewPersonServiceClient(connection)

	serverResponse, err := serviceClient.Delete(context.Background(), request)

	if err != nil {
		log.Println("Error Delete", err.Error())
		return nil
	}

	return serverResponse
}
