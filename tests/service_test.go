package tests

import (
	"context"
	"testing"

	"github.com/edmartt/grpc-test/internal/person"
	pb "github.com/edmartt/grpc-test/internal/person/protos/bin"
)

func init() {
	person.DataAccess = mockDataAccess{}
}

func TestCreateGRPC(t *testing.T) {

	pbPerson := &pb.Person{
		Id:        "1",
		FirstName: "protoMock1",
		LastName:  "protoLastMock1",
		Email:     "protoMock1Email",
	}

	service := person.Service{}

	responseGrpc, err := service.Create(context.Background(), pbPerson)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	if responseGrpc.Id == "" {
		t.Errorf("expected %v and got %v", &pbPerson.Id, responseGrpc.Id)
	}

}

func TestGetGRPC(t *testing.T) {
	request := &pb.GetPersonRequest{
		Id: "1",
	}

	rightPerson := &pb.Person{
		Id:        "1",
		FirstName: "ed",
		LastName:  "mart",
		Email:     "email",
	}

	wrongPerson := &pb.Person{
		Id:        "2",
		FirstName: "mock2name",
		LastName:  "mock2lastname",
		Email:     "no email",
	}

	service := person.Service{}

	result, err := service.Get(context.Background(), request)

	if err != nil {
		t.Errorf("Error %s", err)
	}

	if result.Person.Id != rightPerson.Id {
		t.Errorf("expected %s and got %s", rightPerson.Id, result.Person.Id)
	}

	if wrongPerson.Id == result.Person.Id {
		t.Errorf("expected %s and got %s", wrongPerson.Id, result.Person.Id)
	}

}
