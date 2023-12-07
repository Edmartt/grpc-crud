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

func TestDeleteGRPC(t *testing.T) {
	service := person.Service{}

	request := &pb.DeletePersonRequest{
		Id: "1",
	}
	expectedResponse := &pb.DeletePersonResponse{
		Id:     "1",
		Status: "deleted",
	}

	resultDelete, err := service.Delete(context.Background(), request)

	if err != nil {
		t.Errorf("got %s", err)
	}

	if resultDelete.Status != expectedResponse.Status {
		t.Errorf("expected %s and got %s", expectedResponse.Status, resultDelete.Status)
	}
}

func TestNotFoundBeforeDelete(t *testing.T) {
	service := person.Service{}

	requestFilled := &pb.DeletePersonRequest{
		Id: "2",
	}

	expectedResponse := &pb.DeletePersonResponse{
		Id:     requestFilled.Id,
		Status: "not found",
	}

	resultNotFound, err := service.Delete(context.Background(), requestFilled)

	if err != nil {
		t.Errorf("got %s", err)
	}

	if resultNotFound.Status != expectedResponse.Status {
		t.Errorf("expected %s and got %s", expectedResponse.Status, resultNotFound)
	}

}
